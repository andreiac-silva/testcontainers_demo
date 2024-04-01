package user

import (
	"context"

	"github.com/andreiac-silva/testcontainers_demo/domain/model"
)

type Service interface {
	Create(ctx context.Context, user model.User) (*int64, error)
	Get(ctx context.Context, id int64) (model.User, error)
	List(ctx context.Context) ([]model.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s service) Create(ctx context.Context, user model.User) (*int64, error) {
	return s.repository.save(ctx, user)
}

func (s service) Get(ctx context.Context, id int64) (model.User, error) {
	return s.repository.find(ctx, id)
}

func (s service) List(ctx context.Context) ([]model.User, error) {
	return s.repository.findAll(ctx)
}
