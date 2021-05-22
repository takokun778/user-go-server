package usecases

import (
	"github.com/takokun778/user-go-server/domains/models"
	"github.com/takokun778/user-go-server/domains/repositories"
)

// ユーザー作成ユースケース
type UserCreateUseCase struct {
	userRepository repositories.UserRepository
}

// コンストラクタ
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

// ゲッター
func (output *UserCreateUseCaseOutput) User() *models.User {
	return output.user
}

func NewUserCreateUseCaseOutput(user *models.User) (output *UserCreateUseCaseOutput) {
	output = new(UserCreateUseCaseOutput)
	output.user = user
	return
}

/*
 ユーザー作成ユースケースの実処理
 エラー処理以外のif文が発生した場合はドメイン層への移動を検討すること
*/
func (u *UserCreateUseCase) Handle(input *UserCreateUseCaseInput) (output *UserCreateUseCaseOutput, err models.IBaseError) {
	user := models.CreateNewUser(input.name)
	result, err := u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	output = NewUserCreateUseCaseOutput(result)
	return output, nil
}
