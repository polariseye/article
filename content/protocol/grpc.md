GRPC 在golang中的使用
---------------------------------------
Grpc是由谷歌开发的基于HTTP/2的RPC框架。该框架主要使用protobuf对数据进行序列化和反序列化(可以替换为json或者其他的，但貌似很少有这么干的)。其特点如下：
1. 多语言支持，可以使用protobuf工具生成其他语言的代码
2. 由于使用了HTTP/2，所以提供了以流的方式发送或接收数据的功能
3. 支持自定义的数据加密方式，或者使用tls
4. 二进制帧传输，HTTP/2使用HPACK对header进行压缩，header每次只会发送改变量，而不是全量

##HTTP 2.0概述
HTTP/2是HTTP/1的升级版。相比HTTP/1来说，主要变更点如下:
1. 多路复用(Multiplexing)
![](./compare_http1_http2.png)
HTTP2.0使用了多路复用的技术，做到同一个连接并发处理多个请求，而且并发请求的数量比HTTP1.1大了好几个数量级。具体是通过在应用层和传输层之间增加一个二进制分帧层的方式来实现的。
![](./http2fenzhen.jpg)
当然HTTP1.1也可以多建立几个TCP连接，来支持处理更多并发的请求，但是创建TCP连接本身也是有开销的。并且浏览器客户端在同一时间，针对同一域名下的请求有一定数量限制。超过限制数目的请求会被阻塞

2. header压缩
HTTP1.1不支持header数据的压缩，HTTP2.0使用HPACK算法对header的数据进行压缩，并且只会传输变更的header项
3. 服务端push
在client请求时，server可以额外地多传输一部分数据，以便减少client的请求次数


## 服务定义
grpc定义使用的protobuf语法，相关具体语法可参考protobuf相关文章，具体格式为:

	syntax = "proto3"; // 定义使用的protobuf协议版本，都在建议使用3版本，尽量不要使用2		
	package tstPkg; // 定义包名

	message Person { //// 使用message定义一个消息,可以理解为Model的定义
		string Name = 1;
	}

	message ReturnVal { //// 使用message定义一个消息,可以理解为Model的定义
		int Code = 1;
	}
	
	service TstService { // 具体的服务定义
		// 以rpc为标记说明为rpc接口
		rpc HelloWorld(Person) returns (ReturnVal) {} 

		// rpc中被标记为stream时，则可以以流的方式传输多条数据,如果参数和返回值都有标记，则可以用来实现双向通信
		rpc ExChange(stream Person) return (stream Person) {}
	}

关键字说明:
1. ``rpc``代表这是一个rpc函数
		rpc HelloWorld(Person) returns (ReturnVal) {} 
2. ``returns`` 用于指定返回值，**所有函数必须指定返回值，否则会编译不过**
3. 如果需要流式传输，则使用``stream``关键字修饰，标记为流式传输后，可以传输多个对象。由于这个关键字，所以把流传输分成了三种：
	* 服务端流式 RPC，即客户端发送一个请求给服务端，可获取一个数据流用来读取一系列消息。客户端从返回的数据流里一直读取直到没有更多消息为止
			rpc GetList() return (stream Person) {}
	* 客户端流式 RPC，即客户端用提供的一个数据流写入并发送一系列消息给服务端。一旦客户端完成消息写入，就等待服务端读取这些消息并返回应答
			rpc UpdateList(stream Person) return (ReturnVal) {}
	* 双向流式 RPC，即两边都可以分别通过一个读写数据流来发送一系列消息。这两个数据流操作是相互独立的，所以客户端和服务端能按其希望的任意顺序读写。**可以在不同的协程分别操作这两个流**
			rpc ExChange(stream Person) return (stream ReturnVal) {}

重要说明
1. 所有函数必须有返回值，否则会编译不过
2. rpc涉及到的所有类型必须为message定义的类型，不能是基本数据类型，否则会提示对应类型不无法识别
3. 不能有函数重载，所有函数名都不能一样

关于protobuf的其他语法，请见文档[protobuf use](./protobuf use.md)

生成对应golang代码:``protoc --gofast_out=plugins=grpc:../ TstService.proto``。如果用的是官方插件，则命令为``protoc --go_out=plugins=grpc:../ TstService.proto``

**生成的代码对外包含四部分:{服务名}Client接口，创建Client代理对象，{服务名}Server函数，Register{服务名}Server函数。如果有使用stream，则会生成相应的 {服务名}_{函数名}Server接口和实现**

## 服务端使用
下载依赖库,对应go.mod内容（需求根据实际调整replace信息）:
	
	require (
		github.com/golang/protobuf v0.0.0-c823c79ea1570fb5ff454033735a8e68575d1d0f
		google.golang.org/grpc v1.9.2
		golang.org/x/net v0.0.0-20180826012351-8a410e7b638d // indirect
		google.golang.org/appengine v1.1.0 // indirect
		golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3 // indirect
	    golang.org/x/text v0.3.0 // indirect
	    golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f // indirect
	    google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 // indirect
	    golang.org/x/tools v0.0.0-20190114222345-bf090417da8b // indirect
	    cloud.google.com/go v0.26.0 // indirect
	    golang.org/x/sys v0.0.0-20180830151530-49385e6e1522 // indirect
	    golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be // indirect
	)
	
	replace (
	    golang.org/x/text v0.3.0 => ../../github.com/golang/text
		google.golang.org/grpc v1.9.2 => ../../github.com/grpc/grpc-go
		golang.org/x/net v0.0.0-20180826012351-8a410e7b638d => ../../github.com/golang/net
		google.golang.org/appengine v1.1.0 => ../../github.com/golang/appengine
		golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3 => ../../github.com/golang/lint
	    golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => ../../github.com/golang/sync
	    google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => ../../github.com/golang/go-genproto
	    golang.org/x/tools v0.0.0-20190114222345-bf090417da8b => ../../github.com/golang/tools
	    cloud.google.com/go v0.26.0 => ../../github.com/googleapis/go-cloud/google-cloud-go
	    golang.org/x/sys v0.0.0-20180830151530-49385e6e1522 => ../../github.com/golang/syc
	    golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => ../../github.com/golang/oauth2
	)

服务具体实现：
	
	type TstServiceServerImpl struct {
	}
	
	// 简单RPC的定义
	func (this *TstServiceServerImpl) HelloWorld(contextObj context.Context, person *Person) (*ReturnVal, error) {
		var returnVal ReturnVal
		returnVal.Code = 1
		fmt.Println("收到一个请求")
	
		return &returnVal, nil
	}
	
	// 服务端流式 RPC，即客户端发送一个请求给服务端，可获取一个数据流用来读取一系列消息。客户端从返回的数据流里一直读取直到没有更多消息为止
	func (this *TstServiceServerImpl) GetList(str *tstPkg.CommonString, svr TstServiceServer_GetListServer) error {
		fmt.Println("收到的数据为:", str.Data)
	
		for {
			time.Sleep(1 * time.Second)
			svr.Send(&tstPkg.Person{
				Name:   fmt.Sprintf("%d", count),
			})
		}
	
		return nil
	}
	
	// 客户端流式 RPC，即客户端用提供的一个数据流写入并发送一系列消息给服务端。一旦客户端完成消息写入，就等待服务端读取这些消息并返回应答
	func (this *TstServiceServerImpl) UpdateList(svr tstPkg.SampleService_UpdateListServer) error {
		return nil
	}
	
	// 双向流式 RPC，即两边都可以分别通过一个读写数据流来发送一系列消息。这两个数据流操作是相互独立的，所以客户端和服务端能按其希望的任意顺序读写
	func (this *TstServiceServerImpl) Exchange(svr tstPkg.SampleService_ExchangeServer) error {
		contxtObj := svr.Context()

		var waitObj sync.WaitGroup
		waitObj.Add(2)

		// 接收数据	
		go func() {
			defer waitObj.Done()
			for {
				select {
				case <-contxtObj.Done():
					break
				default:
					item, err := svr.Recv()
					if err == io.EOF {
						break
					} else if err != nil {
						fmt.Println("出错了：", err.Error())
						break
					}
	
					data, _ := json.Marshal(item)
					fmt.Println("收到数据：", string(data))
				}
			}
		}()
	
		// 发送数据
		go func() {
			defer waitObj.Done()
			var i int32 = 0
			for {
				time.Sleep(1 * time.Second)
	
				svr.Send(&tstPkg.Person{
					Name: fmt.Sprintf("server_%v", i),
				})
				i += 1
			}
		}()
	
		waitObj.Wait()
	
		return nil
	}

主函数实现

	func main() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 20004)) //// 监听端口
		if err != nil {
			fmt.Printf("failed to listen: %v", err)
			return
		}
		grpcServer := grpc.NewServer() //// 定义grpc服务，根据实际可以传相关配置参数
		tstPkg.RegisterTstServiceServer(grpcServer, new(TstServiceServerImpl)) //// 注册服务
	
		grpcServer.Serve(lis) //// 开始处理。此处会卡住
	}

## 客户端使用
客户端实现相对简单一点，直接使用`grpc.Dial`即可创建连接,然后使用生成的相关`New{服务名}Client`即可获取对应本地代理对象。具体实现如下:
	
	func main() {
		conn, err := grpc.Dial("127.0.0.1:20004", grpc.WithInsecure()) //// 连接到Server，并设置为不需要安全验证
		if err != nil {
			fmt.Println("error:", err.Error())
			return
		}
		defer conn.Close()
	
		// 通过连接获取一个TstServer的本地代理对象
		clientObj := NewTstServiceClient(conn)
	
		//SimpleRpcTst(clientObj)
		// DoubleStreamTst(clientObj)
	}
	
	// 普通RPC测试，
	func SimpleRpcTst(clientObj TstServiceClient) {
		// 简单RPC测试
		res, err := clientObj.HelloWorld(context.Background(), &Person{
			Name:   "2",
		})
		if err != nil {
			println("接口调用失败:", err.Error())
			return
		}
	
		d, _ := json.Marshal(res)
		println("输出结果为:", string(d))
	}
	
	// 服务端流测试
	func ReturnStreamTst(clientObj TstServiceClient) {
		contxtObj := context.Background()
	
		result, err := clientObj.GetList(contxtObj)
		if err != nil {
			fmt.Println("error :", err.Error())
			return
		}
	
		count := 0
		for {
			itemResult, err := result.Recv()
			if err != nil {
				fmt.Println("error :", err.Error())
				return
			}
	
			count++
			bytesData, _ := json.Marshal(itemResult)
			fmt.Println("收到一项:", string(bytesData))
		}
	}
	
	// 双向流测试
	func DoubleStreamTst(clientObj TsteServiceClient) {
		contextObj := context.Background()
		svr, err := clientObj.Exchange(contextObj)
		if err != nil {
			fmt.Println("调用出错:", err.Error())
			return
		}
	
		waitObj := sync.WaitGroup{}
		waitObj.Add(2)

		// 收数据
		go func() {
			defer waitObj.Done()
			for {
				item, err := svr.Recv()
				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Println("出错了：", err.Error())
					break
				}
	
				data, _ := json.Marshal(item)
				fmt.Println("收到数据：", string(data))
			}
		}()
	
		// 发数据
		go func() {
			defer waitObj.Done()
			var i int32 = 0
			for {
				time.Sleep(1 * time.Second)
	
				svr.Send(&Person{
					Name: fmt.Sprintf("client_%v", i),
				})
				i += 1
			}
		}()
	
		waitObj.Wait()
	}

## 总结
 Grpc封装很完整使用是一件十分轻松的事。对它我是又爱又恨。具体如下：

1. 相比传统基于HTTP/1的restful api，确实效率高了很多。但毕竟是HTTP，没办法和裸的TCP相比
2. 可以把Grpc当做是底层链路和协议解析，基于双向stream实现上层的具体映射。这在HTTP的大环境下堪比websocket。 大概定义如下:

		 message Request {
			int MethodId = 1;
			repeated byte RequestData = 2; ////消息的具体字节数组
		}
		message Response {
			int MethodId = 1;
			int Code = 2;
			repeated byte RequestData = 3; //// 应答的字节数组
		}
		
		service TransformService {
			rpc Transform(stream Request) returns (stream Response)
		}
如果公司存在多种开发语言时，使用这种方式也算是极大简化了socket编程，整体使用也很简单
3. HTTP/2只允许单个链接传输10亿流数据。原因在于： 
HTTP/2使用31位整形标示流，服务端使用奇数，客户端使用偶数，所以总共10亿可用
解决思路：超过一定数量的流，需要重启链接。

## 参考资料
* [HTTP1.0 HTTP 1.1 HTTP 2.0主要区别 https://blog.csdn.net/linsongbin1/article/details/54980801/](https://blog.csdn.net/linsongbin1/article/details/54980801/)
* [gRPC 官方文档中文版 http://doc.oschina.net/grpc](http://doc.oschina.net/grpc)
* [gRPC https://blog.csdn.net/weiwangchao_/article/details/82023191](https://blog.csdn.net/weiwangchao_/article/details/82023191)
* [gRPC实践总结 https://blog.csdn.net/phantom_111/article/details/78797502](https://blog.csdn.net/phantom_111/article/details/78797502)
* [gRPC最佳实践 https://www.colabug.com/4248015.html](https://www.colabug.com/4248015.html)