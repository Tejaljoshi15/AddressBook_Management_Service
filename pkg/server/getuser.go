package server

import (
	"context"
	pb "infoBlox/pkg/gen/proto"
	"infoBlox/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *AddressBookServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	userID := req.GetUserId()

	modelUser, err := a.uc.Get(userID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "An error occurred while retrieving user data: %v", err)
	}

	pbUser := utils.ModelToPb(modelUser)
	return &pb.GetUserResponse{User: pbUser}, nil
}
