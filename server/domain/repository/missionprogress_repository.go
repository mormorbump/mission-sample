package repository

import (
	"com.graffity/mission-sample/server/domain/entity"
	"context"
)

type MissionProgressRepository interface {
	SelectAll(ctx context.Context) (entity.MissionProgresses, error)
	SelectByMissionID(ctx context.Context, missionID entity.MissionID) (entity.MissionProgresses, error)
}
