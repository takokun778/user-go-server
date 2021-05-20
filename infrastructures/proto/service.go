package proto

import (
	"context"

	pb "github.com/takokun778/user-go-grpc/user"
	"github.com/takokun778/user-go-server/interfaces/controllers"
)

type UserService struct {
	userPostController *controllers.UserPostController
}

func (s *UserService) Create(context context.Context, request *pb.CreateRequest) (response *pb.CreateResponse, err error) {
	response, err = s.userPostController.Post(request)
	return
}
