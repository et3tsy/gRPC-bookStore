package main

import (
	"fmt"
	"net"
	"server/dao/mysql"
	rpc "server/grpc"
	"server/logger"
	"server/settings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {

	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}

	// 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()

	// 初始化MySQL
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	// 初始化gRPC
	rpc.Init()

	// 设置监听
	listener, err := net.Listen(
		viper.GetString("grpc.network"),
		viper.GetString("grpc.addr"),
	)
	if err != nil {
		fmt.Printf("init listener failed, err:%v\n", err)
		return
	}

	// 启动gRPC服务
	if err := rpc.Run(listener); err != nil {
		fmt.Printf("start gRPC failed, err:%v\n", err)
		return
	}
}
