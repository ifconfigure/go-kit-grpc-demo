package Handler

import (
	"github.com/go-kit/kit/transport/grpc"
	"go-kit-gRpc-learn/go-kit/EndPoint"
	"go-kit-gRpc-learn/go-kit/TransPort"
)

func BookServiceHandler()  *TransPort.BookService{
	BookServer := new(TransPort.BookService)
	//创建bookList的Handler
	bookListHandler := grpc.NewServer(
		EndPoint.MakeGetBookListEndpoint(),
		EndPoint.DecodeRequest,
		EndPoint.EncodeResponse,
	)
	//bookServer 增加 go-kit流程的 bookList处理逻辑
	BookServer.BookListHandler = bookListHandler

	//创建bookInfo的Handler
	bookInfoHandler := grpc.NewServer(
		EndPoint.MakeGetBookInfoEndpoint(),
		EndPoint.DecodeRequest,
		EndPoint.EncodeResponse,
	)
	//bookServer 增加 go-kit流程的 bookInfo处理逻辑
	BookServer.BookInfoHandler = bookInfoHandler

	return BookServer
}