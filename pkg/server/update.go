package server

import (
	"context"
	pb "infoBlox/pkg/gen/proto"
	"infoBlox/utils"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *AddressBookServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if err := req.GetUser().Validate(); err != nil {
		log.Printf("Validation failed: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Validation failed for user: %v", err)
	}

	modelUser := utils.PbToModel(req.GetUser())

	err := a.uc.Update(modelUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user: %v", err)
	}

	updatedPbUser := utils.ModelToPb(modelUser)
	return &pb.UpdateUserResponse{User: updatedPbUser}, nil
}
