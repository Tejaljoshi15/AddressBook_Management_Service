package user

import (
	"fmt"
	"infoBlox/pkg/db/model"
)

func (uc *User) Create(user *model.User) error {
	err := uc.db.Create(user)
	if err != nil {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}

	return nil
}
