package server

import (
	"context"
	"infoBlox/pkg/gen/proto"
	"infoBlox/utils"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (a *AddressBookServer) AddUser(ctx context.Context, req *proto.AddUserRequest) (*proto.AddUserResponse, error) {

	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request cannot be nil")
	}

	if err := req.GetUser().Validate(); err != nil {
		log.Printf("Validation failed: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Validation failed for user: %v", err)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("Failed to extract metadata from context")
		return nil, status.Errorf(codes.Internal, "Failed to extract metadata")
	}
	if val, exists := md["x-your-custom-header"]; exists {
		log.Printf("Received custom header: %v", val)
	}

	modelUser := utils.PbToModel(req.GetUser())
	if err := a.uc.Create(modelUser); err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to add user to the database: %v", err)
	}

	pbUser := utils.ModelToPb(modelUser)

	return &proto.AddUserResponse{User: pbUser}, nil
}
