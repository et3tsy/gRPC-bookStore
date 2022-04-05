package interact

import (
	"bufio"
	service "client/pb"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// 清空屏幕
func clear() error {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// 停顿
func pause(sec int64) {
	time.Sleep(time.Second * time.Duration(sec))
}

// 读入字符串
func readLine(s *string) (err error) {
	fmt.Scan()
	// 从stdin中取内容直到遇到换行符，停止
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return err
	}
	*s = str[:len(str)-2]

	return nil
}

// 读入整型
func readInt(i *int32) (err error) {
	var s string
	readLine(&s)
	j, err := strconv.Atoi(s)
	*i = int32(j)
	return
}

// 读入整型
func readInt64(i *int64) (err error) {
	var s string
	readLine(&s)
	j, err := strconv.Atoi(s)
	*i = int64(j)
	return
}

// 读入浮点型
func readFloat(i *float32) (err error) {
	var s string
	readLine(&s)
	j, err := strconv.ParseFloat(s, 32)
	*i = float32(j)
	return
}

// 客户端采用命令行窗口进行交互
func ConsoleService(rpc service.BookManagerClient) {

	for {
		if err := clear(); err != nil {
			fmt.Println("CMD fails to work...")
			break
		}
		fmt.Printf("1---Add new book\n")
		fmt.Printf("2---Query by ID\n")
		fmt.Printf("3---Query by name\n")
		fmt.Printf("4---Delete book\n")
		fmt.Printf("Please type in your choice:")

		// 选择分支
		var op string
		readLine(&op)

		switch op {
		case "1":
			{
				_, err := InsertNewBook(rpc)
				if err != nil {
					fmt.Printf("Insert book failed...[%v]", err)
				} else {
					fmt.Printf("Insert book finished...")
				}
				pause(3)
			}
		case "2":
			{
				err := QueryBookByID(rpc)
				if err != nil {
					fmt.Printf("Fail to fetch the record...[%v]", err)
				}
				pause(10)
			}
		case "3":
			{
				err := QueryBookByName(rpc)
				if err != nil {
					fmt.Printf("Query book by name failed...[%v]", err)
				}
				pause(10)
			}
		case "4":
			{
				_, err := DeleteBook(rpc)
				if err != nil {
					fmt.Printf("Delete book failed...[%v]", err)
				} else {
					fmt.Printf("Delete book finished...")
				}
				pause(3)
			}
		default:
			{
				fmt.Printf("Illegal value...\n")
				pause(3)
			}
		}

	}
}

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
