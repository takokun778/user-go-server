package controllers

import (
	pb "github.com/takokun778/user-go-grpc/user"
	"github.com/takokun778/user-go-server/usecases"
)

type UserPostController struct {
	userCreateUseCase usecases.UserCreateUseCase
}

func NewUserPostController(usecase usecases.UserCreateUseCase) (controller *UserPostController) {
	controller = new(UserPostController)
	controller.userCreateUseCase = usecase
	return
}

func (c *UserPostController) Post(request *pb.CreateRequest) (response *pb.CreateResponse, err error) {
	input := usecases.NewUserCreateUseCaseInput(request.GetName())

	output, err := c.userCreateUseCase.Handle(input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		User: &pb.User{
			Id:   output.User().Id(),
			Name: output.User().Name(),
		},
	}, nil
}
