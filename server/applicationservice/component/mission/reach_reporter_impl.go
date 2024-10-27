package mission

import (
	"com.graffity/mission-sample/server/applicationservice/dto/mission"
	"com.graffity/mission-sample/server/domain/entity"
	"com.graffity/mission-sample/server/domain/repository"
	"com.graffity/mission-sample/server/domain/service"
	"context"
	"fmt"
)

type reachReporter struct {
	mRepo  repository.MissionRepository
	mpRepo repository.MissionProgressRepository
	umRepo repository.UserMissionRepository
}

func NewReachReporter(
	missionRepository repository.MissionRepository,
	missionProgressRepository repository.MissionProgressRepository,
	userMissionRepository repository.UserMissionRepository,
) Reporter {
	return &reachReporter{
		mRepo:  missionRepository,
		mpRepo: missionProgressRepository,
		umRepo: userMissionRepository,
	}
}

func (r *reachReporter) Report(ctx context.Context, userID entity.UserID, document *mission.Document) (mission.Results, error) {
	ret := make(mission.Results, 0)
	for _, dm := range document.DataList {
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
		p := document.Form.GetMaxProgress(m)
		rm := service.NewReflectReachMissionStatus(um, mps, p)
		updated := rm.UpdateUserMission(ctx)
		if !updated {
			continue
		}
		if err = r.umRepo.Save(ctx, um); err != nil {
			return nil, err
		}

		ret = append(ret, &mission.Result{
			MissionData: &mission.Data{
				Mission:     m,
				UserMission: um,
			},
		})
	}
	return ret, nil
}
