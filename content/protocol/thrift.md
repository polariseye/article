Thrift教程
-------------------------------
是一个CS结构的RPC框架，以使用中间语言IDL（Interface Description Language）定义RPC接口和数据并生成具体具体语言代码的方式实现跨语言。

# 下载安装
下载地址：[http://www.apache.org/dyn/closer.cgi?path=/thrift/0.12.0/thrift-0.12.0.exe](http://www.apache.org/dyn/closer.cgi?path=/thrift/0.12.0/thrift-0.12.0.exe)

代码生成:`thrift --gen java Test.thrift`

# 基本类型

* byte:有符号字节
* i16:16位有符号整数
* i32:32位有符号整数
* i64:64位有符号整数
* double:64位浮点数

不支持无符号类型的数值

# 常量
````
const i32 MAX_REQUEST_TIME = 10;
````
* 末尾可以有分号，也可以没有
* 支持16进制

# 类型定义

````
typedef i32 Integer
````
类型定义末尾没有逗号或者分号

# 注释

````
# #号开头的注释
// //开头的注释
/*
 多行注释
*/
````
支持以上三种注释

# 集合类型

* list&lt;T&gt;:列表
* set&lt;T&gt;:无序集合,元素不能重复
* map&lt;Key,Value&gt;:字典结构

# 结构体
````
struct Human {
	1: required string name; //// 代表这个字段为必填字段
	2: optional i32 sex; //// optional代表这个字段为可选字段
}

````

# 枚举
````
enum Sex{
	Man,
	Woman
}
````

# 异常(exception)

支持自定义异常

````
exception CustException{
	1: i32 Code;
	2: string Msg;
}

service HelloworldService{
	string Hi(1:string name) throws CustException; //// 使用关键字throws指定要抛出的异常
}
````

# 服务(service)
````
service HelloworldService {
	string Hello(1: string name);
}
````

# 命名空间
`namespace 语言名 路径`
````
namespace java com.baidu.hi
````

* 末尾不能有分号