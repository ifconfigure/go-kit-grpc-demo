package main

import (
	_ "github.com/opentracing/opentracing-go"
	_ "github.com/opentracing/opentracing-go/ext"
	_ "github.com/uber/jaeger-client-go"
	_ "github.com/uber/jaeger-client-go/config"
	_ "github.com/uber/jaeger-lib/metrics"
	Book "go-kit-gRpc-learn/gRpc/PB"
	"go-kit-gRpc-learn/go-kit/Handler"
	"go-kit-gRpc-learn/util"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/grpclog"
	_ "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	_ "io"
	_ "log"
	"net"
	_ "time"
)

func main() {
	//初始化jaeger
	servOpts := util.InitGrpcServerJaeger()

	//建立grpc服务
	bookServiceHandler := Handler.BookServiceHandler()
	serviceAddress := ":8972"
	ls, _ := net.Listen("tcp", serviceAddress)
	gs := grpc.NewServer(servOpts...)

	//注册反射
	reflection.Register(gs)

	//grpc注册服务
	Book.RegisterBookServiceServer(gs, bookServiceHandler)
	_ = gs.Serve(ls)
}
