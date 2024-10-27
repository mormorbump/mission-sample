package repository

import (
	"com.graffity/mission-sample/server/domain/entity"
	"context"
)

type UserRepository interface {
	SelectByPKs(ctx context.Context, pks entity.UserPKs) (entity.Users, error)
	Create(ctx context.Context, user *entity.User) error
}
