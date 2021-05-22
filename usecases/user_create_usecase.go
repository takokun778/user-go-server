package usecases

import (
	"github.com/takokun778/user-go-server/domains/models"
	"github.com/takokun778/user-go-server/domains/repositories"
)

type UserCreateUseCase struct {
	userRepository repositories.UserRepository
}

func NewUserCreateUseCase(userRepository repositories.UserRepository) (usecase *UserCreateUseCase) {
	usecase = new(UserCreateUseCase)
	usecase.userRepository = userRepository
	return
}

type UserCreateUseCaseInput struct {
	name string
}

func NewUserCreateUseCaseInput(name string) (input *UserCreateUseCaseInput) {
	input = new(UserCreateUseCaseInput)
	input.name = name
	return
}

type UserCreateUseCaseOutput struct {
	user *models.User
}

func (output *UserCreateUseCaseOutput) User() *models.User {
	return output.user
}

func NewUserCreateUseCaseOutput(user *models.User) (output *UserCreateUseCaseOutput) {
	output = new(UserCreateUseCaseOutput)
	output.user = user
	return
}

func (u *UserCreateUseCase) Handle(input *UserCreateUseCaseInput) (output *UserCreateUseCaseOutput, err error) {
	user := models.CreateNewUser(input.name)
	result, err := u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	output = NewUserCreateUseCaseOutput(result)
	return output, nil
}
