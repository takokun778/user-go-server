package injector

import (
	"github.com/takokun778/user-go-server/domains/repositories"
	"github.com/takokun778/user-go-server/interfaces/controllers"
	"github.com/takokun778/user-go-server/interfaces/gateways"
	"github.com/takokun778/user-go-server/usecases"
)

type UseCase struct {
	create *usecases.UserCreateUseCase
}

func newUseCase(userRepository repositories.UserRepository) (usecase *UseCase) {
	usecase = new(UseCase)
	usecase.create = usecases.NewUserCreateUseCase(userRepository)
	return
}

type Controller struct {
	Create *controllers.UserCreateController
}

func newController(usecase *UseCase) (controller *Controller) {
	controller = new(Controller)
	controller.Create = controllers.NewUserCreateController(*usecase.create)
	return
}

type Injector struct {
	Controller *Controller
}

func NewInjector() (injector *Injector) {
	injector = new(Injector)
	// repositoryの依存性解決
	usecase := newUseCase(gateways.NewUserGateway())
	injector.Controller = newController(usecase)
	return
}
