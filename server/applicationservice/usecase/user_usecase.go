package usecase

import (
	"com.graffity/mission-sample/server/applicationservice/component"
	"com.graffity/mission-sample/server/applicationservice/dto/mission"
	"com.graffity/mission-sample/server/domain/entity"
	"com.graffity/mission-sample/server/domain/repository"
	"com.graffity/mission-sample/server/domain/value"
	"context"
)

type UserUsecase struct {
	userRepository   repository.UserRepository
	missionProsessor *component.MissionProcessor
}

func NewUserUsecase(userRepository repository.UserRepository, missionProsessor *component.MissionProcessor) *UserUsecase {
	return &UserUsecase{
		userRepository:   userRepository,
		missionProsessor: missionProsessor,
	}
}

func (u *UserUsecase) Save(ctx context.Context) (*entity.User, mission.Results, error) {
	user := &entity.User{}
	err := u.userRepository.Create(ctx, user)
	if err != nil {
		return nil, nil, err
	}

	// TODO registryで定義したaddReporterにミスがあると例外出る
	forms := mission.Forms{
		{
			MissionType: value.MissionTypeLoginCount,
			Targets: mission.Targets{
				{Progress: 1},
			},
		},
		{
			MissionType: value.MissionTypeUserCreateReach,
			Targets: mission.Targets{
				{Progress: 1},
			},
		},
	}
	missionData, err := u.missionProsessor.UpdateMissions(ctx, user.ID, forms)
	if err != nil {
		return user, nil, err
	}

	return user, missionData, nil
}
