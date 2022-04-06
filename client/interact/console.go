package interact

import (
	service "client/pb"
	"fmt"
)

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
