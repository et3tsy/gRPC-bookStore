package main

import (
	"client/interact"
	service "client/pb"
	"client/settings"
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}

	// 新建连接并且添加证书 WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(viper.GetString("grpc.addr"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("Cannot connect to server")
		return
	}

	// 退出时关闭链接
	defer conn.Close()

	// 调用 Book.pb.go 中的 NewBookManagerClient 方法
	rpc := service.NewBookManagerClient(conn)

	interact.ConsoleService(rpc)

}
