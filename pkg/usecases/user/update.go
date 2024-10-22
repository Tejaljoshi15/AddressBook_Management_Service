package user

import (
	"infoBlox/pkg/db/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (uc *User) Update(user *model.User) error {
	if user == nil {
		return status.Error(codes.InvalidArgument, "user cannot be nil")
	}

	userFromDB, err := uc.db.GetByUserID(user.UserID)
	if err != nil {
		return status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	if userFromDB == nil {
		return status.Error(codes.Internal, "unexpected nil user from database")
	}
	userFromDB.Username = user.Username
	userFromDB.Address = user.Address
	userFromDB.PhoneNumber = user.PhoneNumber

	err = uc.db.Update(userFromDB)
	if err != nil {
		return status.Errorf(codes.Internal, "update failed: %v", err)
	}
	return nil
}
