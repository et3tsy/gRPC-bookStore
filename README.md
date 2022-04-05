# gRPC-bookStore



### 简单程序设计：利用gRPC作为中间件的数据管理系统

<a href="https://imgtu.com/i/qLs2l9"><img src="https://s1.ax1x.com/2022/04/05/qLs2l9.png" alt="qLs2l9.png" border="0" /></a>

### 设计说明——服务端

一般方便管理各类包和函数，我们将所有的包和函数分层。

```
|--dao 与操作数据库，与数据库相关
    |--book.go 与数据库交互，绑定书本对象
    |--mysql.go 完成mysql初始化
|--grpc 初始化gRPC服务
    |--grpc.go 初始化gRPC服务
|--logger 日志记录
    |--logger.go 设置日志记录
|--models 定义模型
    |--book.go 定义书本结构体
|--pb gRPC中间件方法
    |--Book.proto proto3源码
    |--Book.pb.go 中间件生成的go调用代码
    |--book.go 实现gRPC中指定api方法
|--settings
    |--config.yaml 配置环境属性
    |--settings.go 读取环境配置
|--server.go 服务端主函数
```



### 设计说明——服务端

一般方便管理各类包和函数，我们将所有的包和函数分层。

```
|--interact DOS交互设计
    |--console.go CMD窗口应用交互设计
|--models 定义模型
    |--book.go 定义书本结构体
|--pb gRPC中间件方法
    |--Book.proto proto3源码
    |--Book.pb.go 中间件生成的go调用代码
|--settings
    |--config.yaml 配置环境属性
    |--settings.go 读取环境配置
|--client.go 客户端主函数
```



