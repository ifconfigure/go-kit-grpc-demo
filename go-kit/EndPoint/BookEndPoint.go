package EndPoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	Book "go-kit-gRpc-learn/gRpc/PB"
	"go-kit-gRpc-learn/gRpc/Service"
)

//创建bookList的EndPoint
func MakeGetBookListEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//请求列表时返回 书籍列表
		bl := new(Book.BookList)
		bl.BookList = append(bl.BookList, &Book.BookInfo{BookId: 1, BookName: "Go入门到精通"})
		bl.BookList = append(bl.BookList, &Book.BookInfo{BookId: 2, BookName: "go-kit入门到精通"})
		bl.BookList = append(bl.BookList, &Book.BookInfo{BookId: 2, BookName: "go-micro入门到精通"})
		return bl, nil
	}
}

//创建bookInfo的EndPoint
func MakeGetBookInfoEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//请求详情时返回 书籍信息
		req := request.(*Book.BookInfoParams)
		//b := new(Book.BookInfo)
		//b.BookId = req.BookId
		//b.BookName = "Go与微服务"
		//return b, nil

		a := Service.BookService{}
		b, _ := a.GetBookInfo(ctx, req)
		return b, nil
	}
}

func DecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func EncodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}
