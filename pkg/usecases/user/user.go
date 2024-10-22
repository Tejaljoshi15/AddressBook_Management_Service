package user

import (
	"infoBlox/pkg/db/model"
)

type DBService interface {
	Create(*model.User) error
	Delete(string) error
	List() ([]*model.User, error)
	GetByUserID(string) (*model.User, error)
	Update(*model.User) error
}

type User struct {
	db DBService
}

func New(db DBService) *User {
	return &User{
		db: db,
	}
}
