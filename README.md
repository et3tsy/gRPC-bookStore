## 简单程序设计：利用gRPC作为中间件的数据管理系统

<a href="https://imgtu.com/i/qLs2l9"><img src="https://s1.ax1x.com/2022/04/05/qLs2l9.png" alt="qLs2l9.png" border="0" /></a>

### 数据库部分

```sql
create table book(
	book_id BIGINT PRIMARY KEY,
    price FLOAT NOT NULL,
    pages INT NOT NULL CHECK(pages>0),
    title VARCHAR(20) NOT NULL UNIQUE,
    author VARCHAR(20) NOT NULL,
    publisher VARCHAR(20) NOT NULL,
    book_isbn VARCHAR(20) NOT NULL
);
```



### 服务端

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



### 客户端

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



### proto3 生成源代码（定义接口）

```
syntax = "proto3";

option go_package="./;service";
package service;

// Msg: to response with the code
message ProdMsg{
  int32 code = 1; // 1
}

// ParamInt: to post with an integer
message ProdParamInt{
  int32 param = 1; // 1
}

// ParamString: to post with a String
message ProdParamString{
  string param = 1; // 1
}

// Book: define the model of book 
message ProdBook{
  int64 id = 1;       // the id
  string title = 2;   // the title
  string author = 3;  // the author
  string publisher=4; // the publisher
}

// Book: define a list of books
message ProdBookList
{
  repeated ProdBook book_list = 1;
}

service ProdService{
  rpc Add(ProdBook) returns(ProdMsg);                       // to add a new book into database
  rpc QueryByID(ProdParamInt) returns(ProdBook);            // to query the book by its id
  rpc QueryByName(ProdParamString) returns(ProdBookList);   // to query the book by its name
  rpc DeleteById(ProdParamInt) returns(ProdMsg);            // to delete a bool by its id
}

```

生成服务端、客户端目的代码

```
protoc --go_out=plugins=grpc:./ Book.proto
```



### 设计成果展示

```
https://www.bilibili.com/video/BV1K94y1Z7td
```

