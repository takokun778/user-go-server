package usecases

import (
	"github.com/takokun778/user-go-server/domains/models"
	"github.com/takokun778/user-go-server/domains/repositories"
)

type UserCreateUseCase struct {
	userRepository repositories.UserRepository
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

func NewUserCreateUseCaseOutput(user *models.User) (output *UserCreateUseCaseOutput) {
	output = new(UserCreateUseCaseOutput)
	output.user = user
	return
}

func NewUserCreateUseCase(userRepository repositories.UserRepository) *UserCreateUseCase {
	return &UserCreateUseCase{
		userRepository: userRepository,
	}
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
