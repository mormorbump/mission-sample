package component

import (
	"com.graffity/mission-sample/server/applicationservice/dto"
	"com.graffity/mission-sample/server/domain/entity"
	"com.graffity/mission-sample/server/domain/repository"
	"com.graffity/mission-sample/server/domain/service"
	"context"
	"fmt"
)

type countReporter struct {
	mRepo  repository.MissionRepository
	mpRepo repository.MissionProgressRepository
	umRepo repository.UserMissionRepository
}

func NewCountReporter(
	missionRepository repository.MissionRepository,
	missionProgressRepository repository.MissionProgressRepository,
	userMissionRepository repository.UserMissionRepository,
) MissionReporter {
	return &countReporter{
		mRepo:  missionRepository,
		mpRepo: missionProgressRepository,
		umRepo: userMissionRepository,
	}
}

func (r *countReporter) Report(ctx context.Context, userID entity.UserID, document *dto.Document) (dto.Results, error) {
	ret := make(dto.Results, 0)
	for _, dm := range document.MissionDataList {
		m := dm.Mission
		mps, err := r.mpRepo.SelectByMissionID(ctx, m.ID)
		if err != nil {
			return nil, fmt.Errorf("MissionProgresses not found")
		}

		// MissionDataからUserMissionを取得
		um := dm.UserMission
		// nilだったら最初のmissionなので、firstThresholdを取得
		if um == nil {
			threshold := mps.GetFirstThreshold()
			um = entity.NewUserMission(userID, m.ID, threshold)
		}
		// aggregate: 集計
		p := document.Form.GetAggregateProgress(m)

		updated := service.UpdateUserMission(ctx, um, mps, p)
		if !updated {
			continue
		}
		if err = r.umRepo.Save(ctx, um); err != nil {
			return nil, err
		}

		ret = append(ret, &dto.Result{
			MissionData: &dto.MissionData{
				Mission:     m,
				UserMission: um,
			},
		})
	}
	return ret, nil
}
