package component

import (
	"com.graffity/mission-sample/server/applicationservice/component/mission"
	missiondto "com.graffity/mission-sample/server/applicationservice/dto/mission"
	"com.graffity/mission-sample/server/domain/entity"
	"com.graffity/mission-sample/server/domain/repository"
	"com.graffity/mission-sample/server/domain/value"
	"context"
	"fmt"
	"github.com/scylladb/go-set/strset"
)

type MissionProcessor struct {
	MissionRepo         repository.MissionRepository
	MissionProgressRepo repository.MissionProgressRepository
	UserMissionRepo     repository.UserMissionRepository
	reporterMap         map[value.MissionType]mission.Reporter
}

func NewMissionProcessor(
	missionRepository repository.MissionRepository,
	missionProgressRepository repository.MissionProgressRepository,
	userMissionRepository repository.UserMissionRepository,
) *MissionProcessor {
	m := make(map[value.MissionType]mission.Reporter)
	return &MissionProcessor{
		MissionRepo:         missionRepository,
		MissionProgressRepo: missionProgressRepository,
		UserMissionRepo:     userMissionRepository,
		reporterMap:         m,
	}
}

// AddReporter MissionTypeとReporterの組み合わせ情報(mission.Info)を受け取り、reporterMapに格納していくメソッド
// Reporterの実装をMissionTypeで制御する感じ。ポリモーフィズムの準備
// 利用するときはinfoListに代入するためにReporterインスタンスを作らねばならず、registry的な使い方をするはず。
// TODO 上記の解釈は適切か
func (p *MissionProcessor) AddReporter(infoList ...mission.Info) {
	for _, info := range infoList {
		p.reporterMap[info.MissionType] = info.Reporter
	}
}

func (p *MissionProcessor) UpdateMissions(ctx context.Context, userID entity.UserID, forms missiondto.Forms) (missiondto.Results, error) {
	missionIDSet, formMap, err := p.getMissionIDSetAndFormMap(ctx, forms)
	if err != nil {
		return nil, err
	}

	ms, ums, err := p.getMissionsAndUserMissions(
		ctx,
		userID,
		toMissionIDs(missionIDSet),
	)
	if err != nil {
		return nil, err
	}
	missionMap := ms.ToMapByMissionType()
	userMissionMap := ums.ToMapByMissionID()

	ret, err := p.updateMissions(ctx, userID, missionMap, userMissionMap, formMap)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// getMissionIDSetAndFormMap
// MissionIDのユニークであるIDSetと、MissionTypeをキーとし、キーでグルーピングしたFormMapを取得するメソッド
// MissionマスタにはMissionTypeがあるので、それぞれ別々で取得してもあとで紐付けられる(なんならMissionTypeが同じMissionIDがあるので別じゃないとだめ)
func (p *MissionProcessor) getMissionIDSetAndFormMap(ctx context.Context, forms missiondto.Forms) (*strset.Set, map[value.MissionType]*missiondto.Form, error) {
	missions, err := p.MissionRepo.SelectAll(ctx)
	if err != nil {
		return nil, nil, err
	}
	// 元のFormを変更しないようにコピー
	copyForms := make(missiondto.Forms, len(forms))
	copy(copyForms, forms)

	idSet := strset.New()
	formMap := make(map[value.MissionType]*missiondto.Form)

	for _, form := range copyForms {
		// form情報からマスターミッションIDsを特定する。
		// strset.Addを用いることで、ユニークをとっている。これにより重複進捗を防ぐ
		idSet.Add(missions.
			FilterByMissionType(form.MissionType).
			FilterByTargets(form.Targets).
			GetIDs()..., //スライスの展開
		)

		// すでにformMapに同一のMissionTypeが存在していたらTargets(対象IDと進捗値のstructでできたスライス)を統合し、formMapに格納し直す
		f, ok := formMap[form.MissionType]
		if ok {
			f.Targets = append(f.Targets, form.Targets...)
		} else {
			f = form
		}
		formMap[form.MissionType] = f
	}
	return idSet, formMap, nil
}

func (p *MissionProcessor) getMissionsAndUserMissions(ctx context.Context, userID entity.UserID, missionIDs entity.MissionIDs) (entity.Missions, entity.UserMissions, error) {
	mPKs := make(entity.MissionPKs, 0, len(missionIDs))
	umPKs := make(entity.UserMissionPKs, 0, len(missionIDs))

	// クエリ作成のためそれぞれのPKリストを作成
	for _, missionID := range missionIDs {
		mPKs = append(mPKs, &entity.MissionPK{
			MissionID: missionID,
		})
		umPKs = append(umPKs, &entity.UserMissionPK{
			UserID:    userID,
			MissionID: missionID,
		})
	}

	ms, err := p.MissionRepo.SelectByPKs(ctx, mPKs)
	if err != nil {
		return nil, nil, err
	}
	ums, err := p.UserMissionRepo.SelectByPKs(ctx, umPKs)

	return ms, ums, nil
}

// updateMissions
// Mission更新処理を行うメソッド。対応するReporterを実行したいのでまずDocumentを作成する。
func (p *MissionProcessor) updateMissions(
	ctx context.Context,
	userID entity.UserID,
	mMap map[value.MissionType]entity.Missions,
	umMap map[entity.MissionID]*entity.UserMission,
	fMap map[value.MissionType]*missiondto.Form,
) (missiondto.Results, error) {
	ret := make(missiondto.Results, 0, len(fMap))

	for missionType, form := range fMap {
		// missionTypeに対応したmissionsを取得。なければスキップ
		ms, ok := mMap[missionType]
		if !ok {
			continue
		}
		// MissionsとUserMissionMapを用いて、Documentを作成
		mdl := missiondto.CreateMissionDataList(ms, umMap)
		doc := missiondto.Document{
			Form:     form,
			DataList: mdl,
		}
		// AddReporterでmissionTypeに紐づいたRepoterの実装をreporterMapに入れているので、それを順番に呼び出して実行
		repo, ok := p.reporterMap[missionType]
		// missionは存在するのにreporterが存在しなかったら例外
		if !ok {
			return nil, fmt.Errorf("reporter not found")
		}

		// missionTypeごとにmissionが複数あるので、doc.DataListも複数になってResultも複数になる
		results, err := repo.Report(ctx, userID, &doc)
		if err != nil {
			return nil, err
		}

		ret = append(ret, results...)
	}
	return ret, nil
}

func toMissionIDs(IDSet *strset.Set) entity.MissionIDs {
	IDs := IDSet.List()
	ret := make(entity.MissionIDs, len(IDs))
	for _, id := range IDs {
		ret = append(ret, entity.MissionID(id))
	}
	return ret
}
