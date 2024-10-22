package storage

import (
	"fmt"
	"infoBlox/pkg/db/model"
)

type DB struct {
	storage map[string]*model.User
}

func New() *DB {
	return &DB{
		storage: make(map[string]*model.User),
	}
}

func (db *DB) Create(user *model.User) error {
	if _, exists := db.storage[user.UserID]; exists {
		return fmt.Errorf("user %s already exists", user.UserID)
	}
	db.storage[user.UserID] = user
	return nil
}

func (db *DB) Delete(userID string) error {
	if _, exists := db.storage[userID]; !exists {
		return fmt.Errorf("user %s does not exist", userID)
	}
	delete(db.storage, userID)
	return nil
}

func (db *DB) List() ([]*model.User, error) {
	users := make([]*model.User, 0, len(db.storage))
	for _, user := range db.storage {
		users = append(users, user)
	}
	return users, nil
}

func (db *DB) GetByUserID(userID string) (*model.User, error) {
	if user, exists := db.storage[userID]; exists {
		return user, nil
	}
	return nil, fmt.Errorf("user %s does not exist", userID)
}

func (db *DB) Update(user *model.User) error {
	if _, exists := db.storage[user.UserID]; !exists {
		return fmt.Errorf("user %s does not exist", user.UserID)
	}
	db.storage[user.UserID] = user
	return nil
}
