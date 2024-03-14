package router

import (
	"context"
	"fmt"
	"net"
	"strconv"

	pb "github.com/muhammadfarhankt/grpc-gateway-grpcui/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
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

func (s *Server) mustEmbedUnimplementedInternServer() {}

func (s *Server) ServeRestApi() {
	conn, err := grpc.DialContext(
		context.Background(),
		":9901",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("error=", err.Error())
	}

	grpcMux := runtime.NewServeMux()
	err = pb.RegisterInternHandler(context.Background(), grpcMux, conn)
	if err != nil {
		fmt.Println("error=", err.Error())
		return
	}

	restAddress := ":9902"
	gwServer := &http.Server{Addr: restAddress, Handler: grpcMux}
	if err := gwServer.ListenAndServe(); err != nil {
		fmt.Println("error=", err.Error())
		return
	}

}
