package services

import (
	"context"

	"github.com/handarudwiki/raport-app-be/model"
)

type UserService interface {
	GetAll(ctx context.Context, query string, cursor int64, num int64) (res []model.User, nextCursor string, err error)

	GetByID(ctx context.Context, id int) (model.User, error)

	Create(ctx context.Context, user model.User) (model.User, error)

	Update(ctx context.Context, user model.User) (model.User, error)

	Delete(ctx context.Context, id int) (bool, error)

	DeleteMany(ctx context.Context, ids []string) (bool, error)

	ImportFromExcel(ctx context.Context, file string) ([]model.User, error)

	UserCurrentLogin(ctx context.Context, email string) (model.User, error)

	Login(ctx context.Context, ur model.User) (*model.Token, error)

	RefreshToken(ctx context.Context, ur *model.Token) error
}
