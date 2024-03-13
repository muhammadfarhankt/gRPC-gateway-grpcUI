package router

import (
	"context"
	"fmt"
	"net"
	"strconv"

	pb "github.com/muhammadfarhankt/grpc-gateway-grpcui/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.InternServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ServeGrpc() {
	listner, err := net.Listen("tcp", ":"+strconv.Itoa(9901))
	if err != nil {
		fmt.Println("error=", err.Error())
	}

	g := grpc.NewServer()

	pb.RegisterInternServer(g, s)
	reflection.Register(g)
	if err := g.Serve(listner); err != nil {
		fmt.Println("GRPC server error", err.Error())
	}
}

func (s *Server) GetInternInfo(c context.Context, req *pb.GetInternInfoReq) (*pb.GetInternInfoResponse, error) {
	res := &pb.GetInternInfoResponse{
		Name:   req.GetName(),
		Domain: "Golang",
		Hub:    "Calicut",
	}
	return res, nil
}

func (s *Server) mustEmbedUnimplementedInternServer() {

}
