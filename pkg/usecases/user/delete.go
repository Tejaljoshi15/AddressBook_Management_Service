package user

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (uc *User) Delete(userID string) error {
	if userID == "" {
		return status.Error(codes.InvalidArgument, "userID cannot be empty")
	}

	_, err := uc.db.GetByUserID(userID)
	if err != nil {
		return status.Errorf(codes.NotFound, "User not found: %v", err)
	}

	err = uc.db.Delete(userID)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to delete user: %v", err)
	}

	return nil
}
