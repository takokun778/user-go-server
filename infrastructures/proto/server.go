package proto

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	pb "github.com/takokun778/user-go-grpc/user"
	"github.com/takokun778/user-go-server/injector"
	"google.golang.org/grpc"
)

func Run() {

	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}

	authInterceptor := grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authenticate))

	server := grpc.NewServer(authInterceptor)

	injector := injector.NewInjector()

	userService := NewUserService(injector)

	pb.RegisterUserServiceServer(server, userService)

	server.Serve(listenPort)
}

func authenticate(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	fmt.Println(token)
	if err != nil {
		return nil, err
	}
	if token != "testtoken" {
		return nil, errors.New("unauthorized")
	}
	return ctx, nil
}
