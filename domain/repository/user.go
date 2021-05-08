package repository

import (
	"context"

	"danglingmind.com/ddd/domain/entity"
)

type UserRepository interface {
	GetById(ctx context.Context, id uint64) (*entity.User, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, id uint64, values map[string]interface{}) error
	Save(ctx context.Context, user entity.User) error
	GetByEmailPassword(ctx context.Context, us *entity.User) (*entity.User, error)
}
