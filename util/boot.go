package util

import (
	"fmt"
	wrapper "github.com/moxiaomomo/grpc-jaeger"
	"google.golang.org/grpc"
	"os"
)

/**
负责框架启动时的一些初始化调用，如：Jaeger链路追踪
*/
func InitGrpcJaeger() []grpc.ServerOption {
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
