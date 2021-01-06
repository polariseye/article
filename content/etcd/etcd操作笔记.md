ETCD学习笔记
-----------------------------------
#概要
ETCD是一个与zookeeper相似的分布式基础组件，用于共享配置和服务发现的分布式，一致性的KV存储系统

# 特点
1. 简单：基于HTTP+JSON的API让你用curl就可以轻松使用
2. 安全：可选SSL客户认证机制。
3. 快速：每个实例每秒支持一千次写操作。
4. 可信：使用Raft算法充分实现了分布式。

# ETCD vs ZK
* 一致性协议： ETCD使用[Raft]协议， ZK使用ZAB（类PAXOS协议），前者容易理解，方便工程实现；
* 运维方面：ETCD方便运维，ZK难以运维；
* 项目活跃度：ETCD社区与开发活跃，ZK已经快死了；
* API：ETCD提供HTTP+JSON, gRPC接口，跨平台跨语言，ZK需要使用其客户端；
* 访问安全方面：ETCD支持HTTPS访问，ZK在这方面缺失；

# 功能
ETCD提供的功能有:
1. key-value 操作:get,set,update,rm,设置节点过期时间
2. 目录操作：setdir,updatedir,ls
3. 监听节点变化:watch,exec-watch
4. 用户管理:
5. 集群节点管理:add,list,remove
6. 备份etcd数据

# 应用场景
* 场景0：配置管理
* 场景一：服务注册于发现（Service Discovery）
* 场景二：消息发布与订阅
* 场景三：负载均衡
* 场景四：分布式通知与协调
* 场景五：分布式锁、分布式队列
* 场景六：集群监控与Leader竞选


# golang使用etcd
golang的SDK地址:[etcd github](https://godoc.org/github.com/coreos/etcd/clientv3)
```
import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"sync"
	"time"
)

func main() {
	// 连接到etcd服务器
	configObj := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	clientObj, err := clientv3.New(configObj)
	if err != nil {
		fmt.Println("new client error:", err.Error())
		return
	}
	defer clientObj.Close()

	// 开一个协程监听变化
	pathVal := "/tst/1"
	go func() {
		watchChan := clientObj.Watcher.Watch(context.Background(), pathVal, clientv3.WithFragment())
		for {
			select {
			case obj := <-watchChan:
				for _, item := range obj.Events {
					fmt.Println(fmt.Sprintf("IsCreate:%v IsModify:%v Type:%v detail:%v ", item.IsCreate(), item.IsModify(), item.Type.String(), item.Kv.String()))
				}
			}
		}
	}()

	// 创建一个租约
	time.Sleep(500 * time.Millisecond)
	leaseObj, e := clientObj.Lease.Grant(context.TODO(), 10) // 单位：秒
	if e != nil {
		fmt.Println("grant lease err:", e.Error())
		return
	}
	// 设置过期的项
	_, err = clientObj.KV.Put(context.Background(), pathVal, "tsts1", clientv3.WithLease(leaseObj.ID))
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	// 添加或者更新项
	_, err = clientObj.KV.Put(context.Background(), pathVal+"/a1", "tst2")
	if err != nil {
		fmt.Println("error2:", err.Error())
		return
	}

	// 设置自动续期
	respObj, err := clientObj.KeepAlive(context.Background(), leaseObj.ID)
	if err != nil {
		fmt.Println("error3:", err.Error())
		return
	}
	go func() {
		for {
			tmpObj := <-respObj
			fmt.Println("完成一次续期 ttl:", tmpObj.TTL)
		}
	}()

	// 不断读取项的值
	for {
		r, e := clientObj.KV.Get(context.Background(), pathVal)
		if e != nil {
			fmt.Println("get err:", e.Error())
			return
		}

		if len(r.Kvs) <= 0 {
			break
		}

		fmt.Println("Count", r.Count, " 	", r.Kvs[0].String())
		time.Sleep(1 * time.Second)
	}
}

```

# 参考资料
* [ETCD相关介绍--整体概念及原理方面](https://www.cnblogs.com/softidea/p/6517959.html)
* [golang etcd简明教程](https://segmentfault.com/a/1190000020868242?utm_source=tag-newest)
* [etcd常用操作介绍](https://segmentfault.com/a/1190000020787391)