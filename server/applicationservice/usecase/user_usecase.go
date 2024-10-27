package usecase

import (
	"com.graffity/mission-sample/server/applicationservice/component"
	"com.graffity/mission-sample/server/applicationservice/component/mission"
	dtomission "com.graffity/mission-sample/server/applicationservice/dto/mission"
	"com.graffity/mission-sample/server/domain/entity"
	"com.graffity/mission-sample/server/domain/repository"
	"com.graffity/mission-sample/server/domain/value"
	"context"
	"errors"
	"log"
)

type UserUsecase struct {
	userRepository   repository.UserRepository
	missionProcessor *component.MissionProcessor
}

func NewUserUsecase(userRepository repository.UserRepository, missionProcessor *component.MissionProcessor) *UserUsecase {
	countReporter := mission.NewCountReporter(missionProcessor.MissionRepo, missionProcessor.MissionProgressRepo, missionProcessor.UserMissionRepo)
	reachReporter := mission.NewReachReporter(missionProcessor.MissionRepo, missionProcessor.MissionProgressRepo, missionProcessor.UserMissionRepo)
	missionProcessor.AddReporter(
		mission.Info{
			MissionType: value.MissionTypeLoginCount,
			Reporter:    countReporter,
		},
		mission.Info{
			MissionType: value.MissionTypeUserCreateReach,
			Reporter:    reachReporter,
		},
	)
	return &UserUsecase{
		userRepository:   userRepository,
		missionProcessor: missionProcessor,
	}
}

func (u *UserUsecase) Save(ctx context.Context) (*entity.User, dtomission.Results, error) {
	user := &entity.User{}
	err := u.userRepository.Create(ctx, user)
	if err != nil {
		return nil, nil, err
	}

	forms := dtomission.Forms{
		{
			MissionType: value.MissionTypeLoginCount,
			Targets: entity.Targets{
				{Progress: 1},
			},
		},
		{
			MissionType: value.MissionTypeUserCreateReach,
			Targets: entity.Targets{
				{Progress: 1},
			},
		},
	}
	missionData, err := u.missionProcessor.UpdateMissions(ctx, user.ID, forms)
	if err != nil {
		// addReporterのエラーハンドリング
		if errors.Is(err, component.ErrAddReporter) {
			// 特定のエラー処理
			log.Printf("addReporter error: %v", err)
			return user, nil, err
		}
	}

	return user, missionData, nil
}
