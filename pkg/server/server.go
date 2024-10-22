package server

import (
	"infoBlox/pkg/db/model"
	pb "infoBlox/pkg/gen/proto"
)

type Usecase interface {
	Create(*model.User) error
	Delete(string) error
	List() ([]*model.User, error)
	Get(string) (*model.User, error)
	Update(*model.User) error
}

type AddressBookServer struct {
	pb.UnimplementedAddressBookServer
	uc Usecase
}

func NewAddressBookServer(uc Usecase) *AddressBookServer {
	return &AddressBookServer{
		uc: uc,
	}
}
