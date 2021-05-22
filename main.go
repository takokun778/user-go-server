package main

import (
	"fmt"

	"github.com/takokun778/user-go-server/infrastructures/proto"
)

func main() {
	fmt.Println("User Server Run!")
	proto.Run()
}
