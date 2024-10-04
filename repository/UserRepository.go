package repository

import (
	"context"

	"github.com/handarudwiki/raport-app-be/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	GetAll(ctx context.Context, query string, cursor int64, num int64) ([]model.User, error)

	GetByID(ctx context.Context, id int) (model.User, error)

	GetByEmail(ctx context.Context, email string) (model.User, error)

	Create(ctx context.Context, user model.User) (model.User, error)

	Update(ctx context.Context, user model.User) (model.User, error)

	Delete(ctx context.Context, id int) (bool, error)

	DeleteMany(ctx context.Context, ids []string) (bool, error)

	CreateMany(ctx context.Context, users []model.User) ([]model.User, error)
}
type repository struct {
	Db *gorm.DB
}

func NewAuthRepository(Db *gorm.DB) AuthRepository {
	return &repository{
		Db: Db,
	}
}

func (r *repository) GetAll(ctx context.Context, query string, cursor int64, num int64) (users []model.User, err error) {
	query = "%" + query + "%"
	err = r.Db.Where("role = ?", model.Teacher).Where("name LIKE ?", query).Where("email LIKE ?", query).Where("id > ?", cursor).Limit(int(num)).Find(&users).Error

	if err != nil {
		return
	}
	return
}

func (r *repository) GetByID(ctx context.Context, id int) (model.User, error) {

}

func (r *repository) GetByEmail(ctx context.Context, email string) (model.User, error) {

}

func (r *repository) Create(ctx context.Context, user model.User) (model.User, error) {

}

func (r *repository) Update(ctx context.Context, user model.User) (model.User, error) {

}

func (r *repository) Delete(ctx context.Context, id int) (bool, error) {

}

func (r *repository) DeleteMany(ctx context.Context, ids []string) (bool, error) {

}

func (r *repository) CreateMany(ctx context.Context, users []model.User) ([]model.User, error) {

}
