package grpc

import (
	"net"

	service "server/pb"

	"google.golang.org/grpc"
)

var rpcServer *grpc.Server

// 初始化将服务储存到grpc的server上
func Init() {
	// new一个grpc的server
	rpcServer = grpc.NewServer()

	// 将刚刚我们新建的BookManager注册进去
	service.RegisterBookManagerServer(rpcServer, new(service.BookService))
}

// 运行rpcServer，传入listener
func Run(listener net.Listener) error {
	return rpcServer.Serve(listener)
}
