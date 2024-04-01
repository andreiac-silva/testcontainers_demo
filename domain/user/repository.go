package user

import (
	"context"

	"github.com/andreiac-silva/testcontainers_demo/domain/model"

	"github.com/uptrace/bun"
)

type Repository interface {
	save(ctx context.Context, user model.User) (*int64, error)
	find(ctx context.Context, id int64) (model.User, error)
	findAll(ctx context.Context) ([]model.User, error)
}

type repositoryDB struct {
	db bun.IDB
}

func NewRepository(db bun.IDB) Repository {
	return &repositoryDB{db: db}
}

func (r repositoryDB) save(ctx context.Context, user model.User) (*int64, error) {
	_, err := r.db.NewInsert().Model(&user).Returning("id").Exec(ctx)
	if err != nil {
		return nil, err
	}
	return user.ID, err
}

func (r repositoryDB) find(ctx context.Context, id int64) (model.User, error) {
	var user model.User
	err := r.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	return user, err
}

func (r repositoryDB) findAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.db.NewSelect().Model(&users).Scan(ctx)
	return users, err
}
