package main

import (
	"fmt"

	"github.com/muhammadfarhankt/grpc-gateway-grpcui/pkg/router"
)

func main() {
	fmt.Println("Welcome to GRPC Server")
	s := router.NewServer()
	s.ServeGrpc()
}
