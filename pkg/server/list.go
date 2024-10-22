package server

import (
	"context"
	"infoBlox/pkg/gen/proto"
	"infoBlox/utils"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *AddressBookServer) ListUsers(ctx context.Context, empty *emptypb.Empty) (*proto.ListUsersResponse, error) {
	modelUsers, err := a.uc.List()
	if err != nil {
		log.Printf("Failed to list users: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to list users from the database: %v", err)
	}

	pbUsers := make([]*proto.User, len(modelUsers))
	for i, modelUser := range modelUsers {
		pbUsers[i] = utils.ModelToPb(modelUser)
	}

	return &proto.ListUsersResponse{Users: pbUsers}, nil
}
