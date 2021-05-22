package controllers

import (
	pb "github.com/takokun778/user-go-grpc/user"
	"github.com/takokun778/user-go-server/usecases"
)

type UserCreateController struct {
	userCreateUseCase usecases.UserCreateUseCase
}

func NewUserCreateController(usecase usecases.UserCreateUseCase) (controller *UserCreateController) {
	controller = new(UserCreateController)
	controller.userCreateUseCase = usecase
	return
}

func (c *UserCreateController) Execute(request *pb.CreateRequest) (response *pb.CreateResponse, err error) {
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
