package injector

import (
	"github.com/takokun778/user-go-server/domains/repositories"
	"github.com/takokun778/user-go-server/interfaces/controllers"
	"github.com/takokun778/user-go-server/interfaces/gateways"
	"github.com/takokun778/user-go-server/usecases"
)

/*
 DI
 依存性を解決するクラス
 serverクラスが肥大化するのを避けるため利用
*/
type Injector struct {
	Controller *Controller
}

func NewInjector() (injector *Injector) {
	injector = new(Injector)
	// repositoryの依存性解決
	// この部分でリポジトリーを付け替えられるようにしている
	// TODO 埋め込むのではなく引数に宣言すべき？ -> テストで使える
	usecase := newUseCase(gateways.NewUserGateway())
	injector.Controller = newController(usecase)
	return
}

// コントローラを集約したクラス
type Controller struct {
	Create *controllers.UserCreateController
}

// コントローラコンストラクタ
func newController(usecase *UseCase) (controller *Controller) {
	controller = new(Controller)
	controller.Create = controllers.NewUserCreateController(*usecase.create)
	return
}

// ユースケースを集約したクラス
type UseCase struct {
	create *usecases.UserCreateUseCase
}

// ユースケースコンストラクタ
func newUseCase(userRepository repositories.UserRepository) (usecase *UseCase) {
	usecase = new(UseCase)
	usecase.create = usecases.NewUserCreateUseCase(userRepository)
	return
}
