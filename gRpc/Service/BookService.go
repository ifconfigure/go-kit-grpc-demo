package Service

import (
	"context"
	Book "go-kit-gRpc-learn/gRpc/PB"
)

//GRPC开始
type BookService struct {}
func (s *BookService) GetBookInfo(ctx context.Context, in *Book.BookInfoParams) (*Book.BookInfo, error) {
	//请求详情时返回 书籍信息
	b := new(Book.BookInfo)
	b.BookId = in.BookId
	b.BookName = "21天精通php"
	return b,nil
}

func (s *BookService) GetBookList(ctx context.Context, in *Book.BookListParams) (*Book.BookList, error) {
	//请求列表时返回 书籍列表
	bl := new(Book.BookList)
	bl.BookList = append(bl.BookList, &Book.BookInfo{BookId:1,BookName:"21天精通php"})
	bl.BookList = append(bl.BookList, &Book.BookInfo{BookId:2,BookName:"21天精通java"})
	return bl,nil
}