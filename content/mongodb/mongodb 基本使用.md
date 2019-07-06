mongodb基本使用
-----------------------------

# 连接到数据库 #

````
mongo 数据库连接字符串
````

数据库连接字符串格式:
````
mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]
```` 

* **mongodb://** 这是固定的格式，必须要指定。

* **username:password@** 可选项，如果设置，在连接数据库服务器之后，驱动都会尝试登陆这个数据库

* **host1** 必须的指定至少一个host, host1 是这个URI唯一要填写的。它指定了要连接服务器的地址。**如果要连接复制集，请指定多个主机地址**。

* **portX** 可选的指定端口，如果不填，默认为27017

* **/database** 如果指定username:password@，连接并验证登陆指定数据库。若不指定，默认打开 test 数据库。

* **?options** 是连接选项。如果不使用/database，则前面需要加上/。所有连接选项都是键值对name=value，键值对之间通过&或;（分号）隔开。

常见选项

* **safe=true|false** true: 在执行更新操作之后，驱动都会发送getLastError命令来确保更新成功。(还要参考 wtimeoutMS). false: 在每次更新之后，驱动不会发送getLastError来确保更新成功
* **connectTimeoutMS=ms** 可以打开连接的时间。
* **socketTimeoutMS=ms** 发送和接受sockets的时间。
* **ssl=true|false** 是否使用TLS/SSL连接

mongodb才安装时没有用户，所以可以直接使用:`mongo`或`mongo mongodb://localhost` 连接到本地数据库

mongodb是基于javascript V8引擎的操作。所以，所有操作都是调用函数而已

# 数据库操作 #

数据库查看
````
show dbs; //// 查看当前数据库服务中存在的数据库
db.getName(); //// 查看当前使用的数据库
````

创建数据库
````
use tstDb; //// 创建tstDb数据库
````
mongodb不存在创建数据库的说法。 也就是说,use语句并不会创建数据库。只有在往数据库里写数据时，数据库才会被创建 

拷贝数据库
````
db.cloneDatabase()
````

删除数据库
````
db.dropDatabase() //// 删除当前数据库
````

# 集合(collection)操作 #

查看存在的集合
````
show dbs; // 使用show 语句查看
db.getCollectionNames(); // 直接使用db对象上的函数查看
````

*  

# 数据操作 #

# 用户管理 #