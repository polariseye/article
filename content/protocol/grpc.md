GRPC 在golang中的使用
---------------------------------------

## 服务定义
rpc定义格式为:

	service tstService {
		rpc HelloWorld(Person) returns (Person) {}
	}

关键字说明:
1. ``rpc``代表这是一个rpc函数
2. ``returns`` 用于指定返回值，**所有函数必须有返回值，否则会编译不过**
3. 如果需要流式传输，则使用``stream``关键字修饰，标记为流式传输后，可以传输多个对象

重要说明
1. 所有函数必须有返回值，否则会编译不过
2. rpc涉及到的所有类型必须为message定义的类型，不能是基本数据类型，否则会提示对应类型不无法识别
3. 不能有函数重载，所有函数名都不能一样

整体服务定义方式如下:
````
// SampleService.proto
message Person {
	string Name = 1;
	int32 age = 2;
}

message CommonResponse {
	string Msg = 1;
}

service SampleService {
	// 简单RPC的定义
	rpc Hello(Person) returns (CommonResponse){}

	// 简单RPC的定义
	rpc Invoke(Person) returns (CommonResponse){}

	// 服务端流式RPC的定义,服务端传入多个请求对象，客户端返回一个响应结果
	rpc GetList(repeated string) returns (stream Person){} 
	
	// 客户端流式RPC的定义,客户端传入多个请求对象，服务端返回一个响应结果
	rpc UpdateList(stream Person) return (long) {}

	// 双向RPC的定义，客户端传入多个请求对象，服务端传入多个请求对象
	rpc Exchange(stream Person) return (stream Person) {}
}

```

关于proto buf的其他语法，请见文档[protobuf use](./protobuf use.md)

生成对应代码到当前的上层目录:``protoc --gofast_out=plugins=grpc:../ SampleService.proto``。如果用的是官方插件，则命令为``protoc --go_out=plugins=grpc:../ SampleService.proto``

## 服务端使用


## 参考资料
* [gRPC 官方文档中文版](http://doc.oschina.net/grpc)
* [gRPC](https://blog.csdn.net/weiwangchao_/article/details/82023191)