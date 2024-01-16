package main

import (
	"context"
	"fmt"
	"grpc/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
}

func (s *Server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	return &proto.Response{Result: a + b}, nil
}
func (s *Server) Subtract(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	return &proto.Response{Result: a - b}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	proto.RegisterMathServiceServer(server, &Server{})
	reflection.Register(server)

	if e := server.Serve(listener); e != nil {
		panic(e)
	} else {
		fmt.Println("Listening on 8080")
	}
}
