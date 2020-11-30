package TransPort

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	Book "go-kit-gRpc-learn/gRpc/PB"
)

type BookService struct {
	BookListHandler grpc.Handler
	BookInfoHandler grpc.Handler
}

//通过grpc调用GetBookInfo时,GetBookInfo只做数据透传, 调用BookServer中对应Handler.ServeGRPC转交给go-kit处理
func (s *BookService) GetBookInfo(ctx context.Context, in *Book.BookInfoParams) (*Book.BookInfo, error) {
	_, rsp, err := s.BookInfoHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*Book.BookInfo), err
}

//通过grpc调用GetBookList时,GetBookList只做数据透传, 调用BookServer中对应Handler.ServeGRPC转交给go-kit处理
func (s *BookService) GetBookList(ctx context.Context, in *Book.BookListParams) (*Book.BookList, error) {
	_, rsp, err := s.BookListHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*Book.BookList), err
}
