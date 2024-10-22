package usecases

import (
	pb "infoBlox/pkg/gen/proto"

	"github.com/stretchr/testify/mock"
)

type MockDBService struct {
	mock.Mock
}

func (m *MockDBService) Create(user *pb.User) (int64, error) {
	args := m.Called(user)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDBService) GetByUsername(username string) (*pb.User, bool) {
	args := m.Called(username)
	user, ok := args.Get(0).(*pb.User) // Safe type assertion with check
	return user, args.Bool(1) && ok    // Return false if type assertion fails
}

func (m *MockDBService) Update(user *pb.User) (int64, error) {
	args := m.Called(user)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockDBService) GetAllUsers() []*pb.User {
	args := m.Called()
	return args.Get(0).([]*pb.User)
}

func (m *MockDBService) Delete(username string) error {
	args := m.Called(username)
	return args.Error(0)
}
