package main

import (
	Book "go-kit-gRpc-learn/gRpc/PB"
	"go-kit-gRpc-learn/go-kit/Handler"
	"google.golang.org/grpc"
	"net"
)

func main() {
	bookServiceHandler := Handler.BookServiceHandler()
	//启动grpc服务
	serviceAddress := ":8972"
	ls, _ := net.Listen("tcp", serviceAddress)
	gs := grpc.NewServer()
	Book.RegisterBookServiceServer(gs, bookServiceHandler)
	_ = gs.Serve(ls)
}
