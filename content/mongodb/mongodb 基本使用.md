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

**数据库查看**
````
// 查看当前数据库服务中存在的数据库
show dbs; 

// 查看当前使用的数据库
db.getName();
````

**切换到数据库或创建数据库**
````
// 切换到数据库tstDb或创建数据库tstDb数据库
use tstDb; 
````
mongodb不存在创建数据库的说法。 也就是说,use语句并不会创建数据库。只有在往数据库里写数据时，数据库才会被创建 

**拷贝数据库**
````
db.copyDatabase(fromdb1,todb1,fromhost1[,username,password]) // 将fromhost1的数据库fromdb1复制到本地的todb1
````

**删除数据库**
````
db.dropDatabase() //// 删除当前数据库
````

**备份数据库**
````
// 把服务器127.0.0.1的数据库sourcedb备份到d:/backup
mongodump -h 127.0.0.1:6379 -d sourcedb -o d:/backup
 
// 反服务器127.0.0.1:6379的数据库sourcedb中的集合user备份到d:/backup
mongodump --host 127.0.0.1 --port 6379 --db sourcedb --collection user --out d:/backup 
````

**还原数据库**
````
// 把目录tst1下的数据恢复到服务器127.0.0.1的数据库targetdb中
mongorestore -h 127.0.0.1:6379 -d targetdb --drop ./tst1 
````

# 集合(collection)操作 #

**查看存在的集合**
````
// 使用show 语句查看
show collections; 

// 直接使用db对象上的函数查看,返回的是一个数组
db.getCollectionNames();
````

**创建集合**
````
// 创建集合collection1
db.createCollection("collection1");

// 创建一个带有指定选项的集合collection2
db.createCollection(collection2,{ capped:true, autoIndexId:true, size:6142800, max:10000}); 
````

数说明：
* capped: 如果为 true，则创建固定集合。固定集合是指有着固定大小的集合，当达到最大值时，它会自动覆盖最早的文档。 **当该值为 true 时，必须指定 size 参数**

* autoIndexId:如为 true，自动在 _id 字段创建索引。默认为 false
* size:为固定集合指定一个最大值（以字节计）。如果 capped 为 true，也需要指定该字段。
* max:指定固定集合中包含文档的最大数量

其他说明：

* 在写数据时，如果集合不存在会自动创建

**删除集合**
````
// 删除集合collection1
db.collection1.drop()
````

**本地集合复制**
````
//// 把集合source_collection拷贝到集合target_collection中
db.source_collection.find().forEach(function(x){db.target_collection.insert(x)})
````

# 文档操作 #

**数据写入**
````
// 往集合connection1写入一条数据
db.connection1.insert({name:"nihao",sex:"男",age:127})
````

**数据更新**
`db.collection_name.update(query,update,option)`
````
// 更新一条数据
db.collection1.update({name:"nihao"},{age:128,city:"chengdu"},{upsert:false,multi:false})
````
说明：
* query:查询条件，具体见数据查询章节
* update:要更新的数据
* upsert：可选，这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入
* multi : 可选，mongodb 默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新

**写入或更新**
````
// 数据存在则更新，不存在则添加
db.collection1.save({
  "_id" : ObjectId("5d285422e1f4f55e3328b1e5"),
  "name": "nihao2"
})
````

**删除文档**
````
// 从集合collection1中删除一条name=nihao的文档
db.collection1.remove({name:nihao},{justOne:true})

db.collection1.remove({name:nihao},1) // 与上面的等效

// 清空集合collection1
db.collection1.remove()
````
说明:
* justOne:如果设为 true 或 1，则只删除一个文档，如果不设置该参数，或使用默认值 false，则删除所有匹配条件的文档

**文档查询**
```
// 查询所有匹配项
db.collection.find(query, projection)

// 查询出一条匹配项
db.collection1.findone(query, projection)
````
示例
````
// 查询出collection1的所有记录,并使用格式化的json格式展示
db.collection1.find().pretty()

// 查询出collection1中name=1的所有记录,只返回记录中的字段name,age,city,并按照nam1升序，age1降序排列
db.collection1.find({name:"1"},{name:1,age:1,city:1}).sort({name:1,age:-1})

// 使用正则表达式搜索name包括ni 或 age在(10,100]的所有记录
db.collection1.find($or:{{name:{$regex:"ni"},$and:{age:{$gt:10},age:{$lte:100}}}})

// 取出结果中的100条，从第11条开始取
db.collection1.find().sort({name:1}).limit(100).skip(10)
````

# 用户管理 #
mongodb默认是没有开启权限验证的。如需开启权限验证，可以使在命令行添加参数`--auth`以开启权限验证。
也可以通过修改配置文件实现，windows里配置文件默认位置`C:\Program Files\MongoDB\Server\4.0\bin\mongod.cfg`，整个配置文件内容如下：
````
# Where and how to store data.
storage:
  dbPath: D:\Program Files\MongoDB\Server\4.0\data
  journal:
    enabled: true


# where to write logging data.
systemLog:
  destination: file
  logAppend: true
  path:  D:\Program Files\MongoDB\Server\4.0\log\mongod.log

# network interfaces
net:
  port: 27017
  bindIp: 127.0.0.1

# 这句就是开启权限验证
security:
  authorization: enabled

````

mongo用户相关管理语句如下
````
// 查看系统当前所有用户
use admin;
db.system.users.find();

//查看当前库内已有用户
show users

//查看当前库内可用的roles，默认只有built-in roles
show roles

// 添加普通用户
use tstDb1;
db.createUser({user:"tstDb1User",pwd:"pwd",roles:[{role:"read",db:"tstDb1User"},{role:"readWrite",db:"tstDb1User"}]})

// 添加超级用户
use admin;
db.createUser({user:"tstDb1User",pwd:"pwd",roles:[{role:"root",db:"admin"}]})

//修改密码
db.changeUserPassword('root','rootNew'); 

// 给已有用户新增权限
db.grantRolesToUser('tstDb1User', [{ role: 'write', db: 'tstDb1User' }])

//删除用户命令如下，虽然所有库的用户信息全存在admin的system.users中，删用户时还是要use <库名>才能删除
use db_name
db.dropUser("<username>")

// 连接到数据库后，使用root进行授权
db.auth("root","rootNew")

````
说明：
1. 创建用户时，必须use到对应库，就相当于此用户和此库绑定了，连接时，也必须连接到对应库，否则会提示连接失败
2. use到admin库创建的用户为超级用户
3. 角色分类：
 1. 数据库用户角色：read、readWrite;
 2. 数据库管理角色：dbAdmin、dbOwner、userAdmin；
 3. 集群管理角色：clusterAdmin、clusterManager、clusterMonitor、hostManager；
 4. 备份恢复角色：backup、restore；
 5. 所有数据库角色：readAnyDatabase、readWriteAnyDatabase、userAdminAnyDatabase、dbAdminAnyDatabase
 6. 超级用户角色：root
 7. 内部角色：__system 这里还有几个角色间接或直接提供了系统超级用户的访问（dbOwner 、userAdmin、userAdminAnyDatabase）
4. 不同的库下可以存在相同的用户名，即admin.root用户和test.root用户是可以同时存在的