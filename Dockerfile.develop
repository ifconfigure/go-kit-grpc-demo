# 拉取 Go 语言最新的基础镜像
FROM golang:latest

#设置环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    PORT=8972

# 在容器内设置 /app 为当前工作目录
WORKDIR /app

# 把文件复制到当前工作目录
COPY . .
# 设置 GOPROXY 环境变量
ENV GOPROXY="https://goproxy.cn"
# 编译项目
RUN go build -o ./bin/rpc-server ./go-kit/rpc-server.go

EXPOSE 8972

ENTRYPOINT ["./bin/rpc-server"]