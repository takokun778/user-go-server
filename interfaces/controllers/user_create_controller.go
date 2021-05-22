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
	response = new(pb.CreateResponse)
	response.User = new(pb.User)
	response.User.Id = output.User().Id()
	response.User.Name = output.User().Name()
	return response, nil
}
