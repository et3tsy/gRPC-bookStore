package interact

import (
	service "client/pb"
	"context"
	"fmt"
)

// 插入新书
func InsertNewBook(rpc service.BookManagerClient) (int32, error) {
	pb := new(service.ParamBook)
	fmt.Printf("Please type in the ID:")
	readInt64(&pb.Id)
	fmt.Printf("Please type in the Price:")
	readFloat(&pb.Price)
	fmt.Printf("Please type in the Pages:")
	readInt(&pb.Pages)
	fmt.Printf("Please type in the Title:")
	readLine(&pb.Title)
	fmt.Printf("Please type in the Author:")
	readLine(&pb.Author)
	fmt.Printf("Please type in the Publisher:")
	readLine(&pb.Publisher)
	fmt.Printf("Please type in the ISBN:")
	readLine(&pb.ISBN)

	resp, err := rpc.Add(context.Background(), pb)

	if err != nil {
		return -1, err
	}

	return resp.Code, nil
}

// 删除书籍
func DeleteBook(rpc service.BookManagerClient) (int32, error) {
	pb := new(service.ParamInt)
	fmt.Printf("Please type in the ID of the book you want to delete:")
	readInt64(&pb.Param)

	resp, err := rpc.DeleteById(context.Background(), pb)

	if err != nil {
		return -1, err
	}

	return resp.Code, nil
}

// 根据 ID 查询书籍
func QueryBookByID(rpc service.BookManagerClient) error {
	pb := new(service.ParamInt)
	fmt.Printf("Please type in the ID of the book you want to query:")
	readInt64(&pb.Param)

	resp, err := rpc.QueryByID(context.Background(), pb)

	if err != nil {
		return err
	}

	fmt.Printf("%v", resp)

	return nil
}

// 根据名字进行模糊查询书籍
func QueryBookByName(rpc service.BookManagerClient) error {
	pb := new(service.ParamString)
	fmt.Printf("Please type in the name of the book you want to query:")
	readLine(&pb.Param)

	resp, err := rpc.QueryByName(context.Background(), pb)

	if err != nil {
		return err
	}

	fmt.Printf("%v", resp)

	return nil
}
