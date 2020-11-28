package main

import (
	"context"
	"fmt"
	Book "go-kit-gRpc-learn/gRpc/PB"
	"google.golang.org/grpc"
)

func main()  {
	serviceAddress := "127.0.0.1:8972"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()
	bookClient := Book.NewBookServiceClient(conn)
	bi, _ := bookClient.GetBookInfo(context.Background(), &Book.BookInfoParams{BookId: 1})
	fmt.Print("获取书籍信息:\t")
	fmt.Println("bookInfo:", bi.BookName)

	bl, _ := bookClient.GetBookList(context.Background(), &Book.BookListParams{Page: 1, Limit: 10})
	fmt.Println("获取书籍列表:\t")
	for _, b := range bl.BookList {
		fmt.Println("bookId:", b.BookId, " => ", "bookName:", b.BookName)
	}

}