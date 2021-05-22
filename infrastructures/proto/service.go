package proto

import (
	"context"

	pb "github.com/takokun778/user-go-grpc/user"
	"github.com/takokun778/user-go-server/injector"
)

// grpcのサービスを実装するクラス
type UserService struct {
	injector *injector.Injector
	// 未実装のServiceがあるがコンパイラエラーが発生しないように以下を宣言
	pb.UnimplementedUserServiceServer
}

// ユーザーサービスコンストラクタ
func NewUserService(injector *injector.Injector) (userServise *UserService) {
	userServise = new(UserService)
	userServise.injector = injector
	return
}

func (s *UserService) Create(context context.Context, request *pb.CreateRequest) (response *pb.CreateResponse, err error) {
	response, err = s.injector.Controller.Create.Execute(request)
	return
}
