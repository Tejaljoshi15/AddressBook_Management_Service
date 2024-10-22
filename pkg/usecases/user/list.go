package user

import (
	"fmt"
	"infoBlox/pkg/db/model"
)

func (uc *User) List() ([]*model.User, error) {
	users, err := uc.db.List()
	if err != nil {

		return nil, fmt.Errorf("error listing users: %w", err)
	}
	return users, nil
}
