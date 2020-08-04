kafka环境搭建与常用命令行
------------------------------------------------------------------------------------
[官方地址: http://kafka.apache.org](http://kafka.apache.org)

[kafka2.4.0下载地址](https://www.apache.org/dyn/closer.cgi?path=/kafka/2.4.0/kafka_2.13-2.4.0.tgz)
[zookeeper下载地址](http://mirror.bit.edu.cn/apache/zookeeper/stable/)

# 单机配置
需要修改config/server.properties的项:

````
# 监听地址，提供给client连接使用
listeners=PLAINTEXT://127.0.0.1:9092
# 集群决策用的地址
advertised.listeners=PLAINTEXT://127.0.0.1:9092
# kafka数据的记录地址
log.dirs=/tmp/kafka-logs
# zookeeper连接地址,多个用, 分隔
zookeeper.connect=localhost:2181,127.0.0.1:3000
# 每个topic的默认分区数量
num.partitions=3
````

# 集群配置
需要修改config/server.properties的项:

````
# 每个节点的唯一Id，需要为整数int32.且不能重复
broker.id=1
# 监听地址，提供给client连接使用
listeners=PLAINTEXT://127.0.0.1:9092
# 集群决策用的地址
advertised.listeners=PLAINTEXT://127.0.0.1:9092
# kafka数据的记录地址
log.dirs=/tmp/kafka-logs
# zookeeper连接地址,多个用, 分隔
zookeeper.connect=localhost:2181,127.0.0.1:3000
````


# 参考资料
* [kafka基本原理概述——patition与replication分配](https://www.cnblogs.com/xjh713/p/7388262.html)
* [kafka入门及配置详解](https://blog.csdn.net/hweixing123/article/details/86536641)
* [Kafka深度解析](http://www.jasongj.com/2015/01/02/Kafka深度解析/)