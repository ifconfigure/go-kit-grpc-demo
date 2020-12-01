package util

import (
	"fmt"
	wrapper "github.com/moxiaomomo/grpc-jaeger"
	"google.golang.org/grpc"
	"os"
)

/**
负责框架启动时的一些初始化调用，如：Jaeger链路追踪服务端
*/
func InitGrpcServerJaeger() []grpc.ServerOption {
	//jaeger链路追踪
	var servOpts []grpc.ServerOption
	tracer, _, err := wrapper.NewJaegerTracer("testSrv", "127.0.0.1:6831")
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}
	if tracer != nil {
		servOpts = append(servOpts, wrapper.ServerOption(tracer))
	}

	return servOpts
}

/**
负责框架启动时的一些初始化调用，如：Jaeger链路追踪客户端
*/
func InitGrpcClientJaeger() []grpc.DialOption {
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	tracer, _, err := wrapper.NewJaegerTracer("testCli", "127.0.0.1:6831")
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}

	if tracer != nil {
		dialOpts = append(dialOpts, wrapper.DialOption(tracer))
	}
	return dialOpts
}
