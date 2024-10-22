package user

import (
	"fmt"
	"infoBlox/pkg/db/model"
	"log"
)

func (uc *User) Get(userID string) (*model.User, error) {
	log.Printf("Attempting to get user by userID: %s", userID)
	user, err := uc.db.GetByUserID(userID)
	if err != nil {
		log.Printf("Error fetching user from database: %s", err)
		return nil, err
	}
	if user == nil {
		log.Println("No user found")
		return nil, fmt.Errorf("no user found with userID: %s", userID)
	}
	log.Printf("User found: %+v", user)
	return user, nil
}
