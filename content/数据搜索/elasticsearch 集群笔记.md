elasticsearch集群笔记
---------------------------------------------
elasticsearch 把数据以分片的方式存储数据。一个分片对应一个底层的工作单元。也就是一个分片对应一个lucene实例。

一个分片 可以是主分片或者副分片。索引内的任意一个文档 都归属于一个主分片，所以主分片的数量决定索引能够保存的最大文档数量。*一个主分片可以保存int.MAX_VALUE-128个文档*

一个副本分片只是一个主分片的拷贝。副本分片作为硬件故障时保护数据不丢失的冗余备份，并为搜索和返回文档等读操作提供服务。

**在索引建立的时候就已经确定了主分片数（索引主分片数不可修改），但是副本分片数可以随时修改。**

**多节点的集群环境下，elasticsearch要求必须设置访问的用户名与密码，否则无法访问**


**查询集群状态**
```
curl -X GET "localhost:9200/_cluster/health?pretty"
```
返回数据
```
{
  "cluster_name" : "elasticsearch",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 5,
  "active_shards" : 5,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 2,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 71.42857142857143
}

```
说明<br />
* status字段代表当前集群在总体上工作是否正常
  * green:所有的主分片和副本分片都正常运行。
  * yellow:所有的主分片都正常运行，但不是所有的副本分片都正常运行。
  * red:有主分片没能正常运行。


**变更副本数量**
```
curl -X PUT "localhost:9200/blogs/_settings?pretty" -H 'Content-Type: application/json' -d'
{
   "number_of_replicas" : 2
}
'
```

# 集群安全配置
安全配置是多节点环境下的强制要求。配置步骤如下:
1. 关闭正在运行的elasticsearch与kibana。
2. 修改每个节点的配置文件elasticsearch.yml。添加配置:**xpack.security.enabled: true**。修改完成后启动所有节点
3. 使用工具`./bin/elasticsearch-setup-passwords`创建密码:
 * 使用参数auto自动创建密码:**./bin/elasticsearch-setup-passwords auto**
 * 使用参数interactive主动设置密码:**./bin/elasticsearch-setup-passwords interactive**

说明:
 * 如果集群没有设置密码，则请求会返回状态码400
 * 密码设置后，如果没有使用用户名密码访问，则会返回状态码401

# 集群管理

**添加节点**<br/>
本机环境下，可以直接启动一个节点。只要集群名相关就会自动组成一个集群。 

* **查看排除列表**
```
curl -X GET "http://localhost:9200/_cluster/state?filter_path=metadata.cluster_coordination.voting_config_exclusions&pretty"
```
* 排除指定节点:
```
 # 按名字排除
 curl -X POST "http://localhost:9200/_cluster/voting_config_exclusions?node_names=<node_names>&wait_for_removal=false"
 # 按id排除
 curl -X POST "http://localhost:9200/_cluster/voting_config_exclusions?node_ids=<node_ids>&wait_for_removal=false"
```
* 清空排除列表
```
curl -X DELETE "http://localhost:9200/_cluster/voting_config_exclusions&wait_for_removal=false"
```

# 问题解决
* **master not discovered or elected yet, an election requires at least 2 nodes with ids from [eAVM79bzRbWA413GAswUfg, Xh0uqNORTG2FXSpfgxpPtA, ZBeR5v65RheQDQLDn6ylXA], have only discovered non-quorum [{node-3}{ZBeR5v65RheQDQLDn6ylXA}{YzgDFM63Scmu15BbRTtexw}{127.0.0.1}{127.0.0.1:9205}{cdfhilmrstw}]; discovery will continue using []**

 原因：是因为重命名了所有节点名。而之前的数据没有删除，这导致了这个问题。按照网上的说法。暂时只能清空所有节点的数据（data目录与logs目录）才能恢复。

* **This node previously joined a cluster with UUID [I5pMbI-9SFi7u_kvxIqcKg] and is now trying to join a different cluster with UUID [8P26aTS9QaO2mDmrFS4LYg]. This is forbidden and usually indicates an incorrect discovery or cluster bootstrapping configuration. Note that the cluster UUID persists across restarts and can only be changed by deleting the contents of the node's data paths [] which will also remove any data held by this node.**

  原因：两个独立的集群加入被拒绝。我测试过程中修改了`network.host`然后重新启动时出现了此问题
  
  解决办法:删除data目录与logs目录即可。

* **http状态码400且，提示Elasticsearch built-in security features are not enabled. Without authentication, your cluster could be accessible to anyone**

  原因:是因为Elastic Stack 7.13以上版本，在集群模式下，默认不开启Elastic安全功能时，会进行安全阻止。只需要修改文件`elasticsearch.yml`添加禁止安全配置设置即可

	```
	cluster.name: "hello"
	network.host: 127.0.0.1
	
	xpack.security.enabled: false
	```

# 参考资料
* [集群的安全配置](https://www.elastic.co/guide/en/elasticsearch/reference/7.16/security-minimal-setup.html)
* [手把手教你搭建一个Elasticsearch集群](https://www.cnblogs.com/tianyiliang/p/10291305.html)