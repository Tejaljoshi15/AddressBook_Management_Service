package db

import (
	"infoBlox/pkg/db/model"
)

type AddressBookStorage interface {
	Create(*model.User) error
	GetByUserID(string) (*model.User, error)
	Delete(string) error
	List() ([]*model.User, error)
	Update(*model.User) error
}

type DBInteractor struct {
	storage AddressBookStorage
}

func New(storage AddressBookStorage) *DBInteractor {
	return &DBInteractor{
		storage: storage,
	}
}

func (dbI *DBInteractor) Create(user *model.User) error {
	return dbI.storage.Create(user)
}

func (dbI *DBInteractor) GetByUserID(userID string) (*model.User, error) {
	return dbI.storage.GetByUserID(userID)
}

func (dbI *DBInteractor) Update(user *model.User) error {
	return dbI.storage.Update(user)
}

func (dbI *DBInteractor) Delete(userID string) error {
	return dbI.storage.Delete(userID)
}

func (dbI *DBInteractor) List() ([]*model.User, error) {
	return dbI.storage.List()
}
