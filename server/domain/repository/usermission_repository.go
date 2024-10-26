package repository

import (
	"com.graffity/mission-sample/server/domain/entity"
	"context"
)

type UserMissionRepository interface {
	SelectByPKs(ctx context.Context, pks entity.MissionProgressPKs) (entity.MissionProgresses, error)
	Save(ctx context.Context, e *entity.UserMission) error
}
