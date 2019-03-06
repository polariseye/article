protocol buffer教程
-------------------------
## protocol buffer安装及使用

1. 安装protocol buffer代码编译工具:[点击下载](https://github.com/google/protobuf/releases)（直接下载可执行程序即可），把可执行程序放置在系统path目录中，以便通过命令protoc使用

2. 安装protocol buffer生成golang代码需要安装对应的golang插件，把可执行程序放置在系统的path目录中即可。
 * 插件1 [官方版](https://github.com/golang/protobuf)
 * 插件2 [修改版](https://github.com/gogo/protobuf)，这个是基于官方版进行的修改，听说其编码效率要高很多
 推荐使用修改版中的[protoc-gen-gofast](https://github.com/gogo/protobuf/tree/master/protoc-gen-gofast)。
<br/>
3. 下载依赖库[protobuf](https://github.com/golang/protobuf)，生成的代码会直接引用这个库

4. 编写proto脚本
		// 文件：tmp.proto
		//指定版本
		//注意proto3与proto2的写法有些不同
		syntax = "proto3";
		 
		//包名，通过protoc生成时go文件时，这个会直接变成包名
		package test;
		 
		//手机类型
		//枚举类型第一个字段必须为0
		enum PhoneType {
		    HOME = 0;
		    WORK = 1;
		}
		 
		//手机
		message Phone {
		    PhoneType type = 1;
		    string number = 2;
		}
		 
		//人
		message Person {
		    //后面的数字表示标识号
		    int32 id = 1;
		    string name = 2;
		    //repeated表示可重复
		    //可以有多个手机
		    repeated Phone phones = 3;
		}
		 
		// 联系簿
		message ContactBook {
		    repeated Person persons = 1;
		}

5. 把编写的proto文件生成golang代码。
 * 插件1使用命令:使用命令 **protoc --go_out=. tmp.proto** 生成代码
 * 插件2使用命令:使用命令 **protoc --gofast_out=. tmp.proto** 生成代码

参考文档：

* [在 Golang 中使用 Protobuf](https://studygolang.com/articles/4872)
* [golang使用protobuf简易教程](https://blog.csdn.net/qq_15437667/article/details/78425151)

## golang使用proto buf
protobuf提供的接口和json一致，直接使用``func Marshal(pb Message) ([]byte, error)``进行序列化，使用``func Unmarshal(buf []byte, pb Message) error``进行反序列化。示例如下：


proto文件定义：

	syntax = "proto3";
	
	package tstPkg;
	 
	//人
	message Person {
	    //后面的数字表示标识号
	    int32 id = 1;
	    string name = 2;
	    //repeated表示可重复
	    //可以有多个手机
	    repeated string phones = 3;
	}

使用代码:

	package main
	
	import (
		"io/ioutil"
		"log"
		"os"
		"testCode/protobufTst/tstPkg"
	
		"github.com/golang/protobuf/proto" // 引用对应的proto操作库
	)
	
	func main() {
		a := tstPkg.Person{} // 创建obj
		a.Id = 1
		a.Name = "今天天气好"
		a.Phones = append(a.Phones, "1")
		a.Phones = append(a.Phones, "2")
	
		data, err := proto.Marshal(&a) //// 序列化
		if err != nil {
			log.Fatal("Marshal error: ", err)
			return
		}

		err = proto.Unmarshal(data, &a) //// 反序列化
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
			return
		}
	
		ioutil.WriteFile("data.txt", data, os.ModePerm)
		print("数据长度：", len(data))
	}


## protocol buffer语法
以下只针对protobuf3进行使用介绍

参考文档：

* [Protobuf3语言指南](http://blog.csdn.net/u011518120/article/details/54604615#DefiningAMessageType)

## protocol buffer内部原理
参考文档:

* [Google Protocol Buffer 的使用和原理](https://www.ibm.com/developerworks/cn/linux/l-cn-gpb/)