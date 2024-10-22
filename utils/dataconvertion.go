package utils

import (
	"infoBlox/pkg/db/model"
	pb "infoBlox/pkg/gen/proto"
)

func ModelToPb(user *model.User) *pb.User {
	return &pb.User{
		UserId:   user.UserID,
		Username: user.Username,
		Address:  user.Address,
		Phone:    user.PhoneNumber,
	}
}

func PbToModel(user *pb.User) *model.User {
	if user == nil {
		return nil
	}
	return &model.User{
		UserID:      user.UserId,
		Username:    user.Username,
		Address:     user.Address,
		PhoneNumber: user.Phone,
	}
}
