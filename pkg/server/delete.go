package server

import (
	"context"
	pb "infoBlox/pkg/gen/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *AddressBookServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	userID := req.GetUserId()

	if userID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "userId must not be empty")
	}

	err := a.uc.Delete(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete user: %v", err)
	}

	return &emptypb.Empty{}, nil
}
