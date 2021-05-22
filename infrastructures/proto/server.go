package proto

import (
	"log"
	"net"

	pb "github.com/takokun778/user-go-grpc/user"
	"github.com/takokun778/user-go-server/injector"
	"google.golang.org/grpc"
)

func Run() {

	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()

	injector := injector.NewInjector()

	userService := NewUserService(injector)

	pb.RegisterUserServiceServer(server, userService)

	server.Serve(listenPort)
}
