### 安装
>git clone https://github.com/ifconfigure/go-kit-grpc-demo.git

---

### 简介
> 结合了go-kit和grpc的学习使用，go-kit因为分层的关系，会使得代码生涩难读，再加上go-kit一般是配合gRpc使用，所以便有了这个demo。

> 可以使用它作为项目的初始化框架，在分层的体系上尽量保持了精简，只需要稍微花点时间理解一下每层的意义，便可以快速进行微服务开发。

> TODO: 后期会引入服务注册/发现，熔断，降级，链路追踪，感谢朋友star💖一波

----

### 如果想运行纯GRPC的DEMO

##### 1.启动gRPC Server
1) ```cd gRpc```

2) ```go build grpc-server.go```

3) ```./grpc-server```

##### 2.启动gRPC Client
1) ```go build  grpc-client.go```

2) ```./grpc-client```

---


### 如果想运行go-kit + gRpc的demo
1) ```cd go-kit```

2) ```go build rpc-server.go```

3) ```./rpc-server```

4) 切换到grpc目录下编译并启动grpc-client,例：```./grpc-client```


---

### 目录结构描述

├── go-kit
`go-kit的主要目录， go-kit的分层思想和微服务都应该写在这里面`

│   ├── EndPoint
`负责业务逻辑，并且返回给Handler`

│   ├── Handler
`负责调用EndPoint，也就是把EndPoint和TransPort层封装起来，并且返回给grpc调用，可以理解为中间人`

│   ├── TransPort
`跟go-kit的思想一样，就是负责数据传入的处理，比如传入检验等`

├── gRpc
`grpc的业务逻辑`

│   ├── PB
`生成的PB文件放这里`

│   ├── Proto
`Proto3原始文件，为什么和PB要分开放是因为可以单独共享proto文件给其它团队生成调用`

│   ├── Service
`这里可以单独写grpc里的业务逻辑，是和go-kit分开的，如果在go-kit的endpoint写了业务就不用在这里写，当然也可以在这里写业务通过endpoint调用`
├── main.go
`暂无用`


