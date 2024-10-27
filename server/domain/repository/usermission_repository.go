package repository

import (
	"com.graffity/mission-sample/server/domain/entity"
	"context"
)

type UserMissionRepository interface {
	SelectByPKs(ctx context.Context, pks entity.UserMissionPKs) (entity.UserMissions, error)
	Save(ctx context.Context, e *entity.UserMission) error
}
