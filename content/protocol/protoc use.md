protocol buffer教程
-------------------------
## protocol buffer安装及使用

1. 安装protocol buffer代码生成工具:[点击下载](https://github.com/google/protobuf/releases)（直接下载可执行程序即可），把可执行程序放置在系统path目录中，以便通过命令protoc使用

2. protocol buffer生成golang代码需要安装对应的golang插件，把可执行程序放置在系统的path目录中即可。[点击下载](https://github.com/gogo/protobuf),推荐使用其中的插件[protoc-gen-gofast](https://github.com/gogo/protobuf/tree/master/protoc-gen-gofast)

3. [protoc-gen-gofast](https://github.com/gogo/protobuf/tree/master/protoc-gen-gofast)生成的代码需要依赖对应protocol buffer的库protobuf，[点击下载](https://github.com/golang/protobuf)

4. 使用命令 **protoc --gofast_out=. myproto.proto** 生成代码

参考文档：

* [在 Golang 中使用 Protobuf](https://studygolang.com/articles/4872)

## protocol buffer语法
以下只针对protobuf3进行使用介绍

参考文档：

* [Protobuf3语言指南](http://blog.csdn.net/u011518120/article/details/54604615#DefiningAMessageType)