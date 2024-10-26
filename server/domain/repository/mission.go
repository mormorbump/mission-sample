package repository

import (
	"com.graffity/mission-sample/server/domain/entity"
	"context"
)

type MissionRepository interface {
	SelectAll(ctx context.Context) (entity.Missions, error)
	SelectByPKs(ctx context.Context, pks entity.MissionPKs) (entity.Missions, error)
}
