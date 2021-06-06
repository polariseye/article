elasticsearch学习笔记
------------------------------------
# 概要
[Elasticsearch](https://www.elastic.co/cn/)是一款基于开源库[Lucene](https://lucene.apache.org)的全文搜索引擎。它可以快速地储存、搜索和分析海量数据。其开发语言为`java`,提供了RESTful API接口,返回的数据格式为普通的`json`格式，因此，其他语言都可以简单地与之通信。


# 概念扫盲
Elasticsearch是面向文档的，它不仅存储文档，而且还对整个文档进行索引，排序和过滤。从而可以检索所有文档的所有内容。名词解释可参考此[链接](https://www.elastic.co/guide/cn/elasticsearch/guide/current/_indexing_employee_documents.html)内容。

**索引（名词）**：相当于关系型数据库中的数据库。

**索引（动词）**：存储一个文档到elasticsearch中的过程称之为索引。如果这个文档已经存在，会进行替换，不存在，则添加。

**类型**：结构相同的数据的集合，相当于关系型数据库的表。

**文档**：一条需要存储到elasticsearch中的数据，相当于关系型数据库的一行数据，每个文档都有一个唯一Id。

**倒排索引**：类似关系型数据库中的 B树（B-tree）索引。elasticsearch使用此算法来给文档添加索引，以提高查询效率。


相同结构的的文档的集合称之为**类型**，类型的集合称之为**索引**。一个elasticsearch可以有多个索引，每个索引可以有多个类型。

# 安装

1. 安装较新的java版本
2. 下载elasticsearch版本:[点击这里](https://www.elastic.co/cn/downloads/elasticsearch)
3. 运行elasticsearch:`bin/elasticsearch.bat`,如果需要在后台运行则需要添加参数:`-d`.默认运行端口为:'9200'


# 基本使用
* **检查elasticsearch运行状态**:`curl 'http://localhost:9200/?pretty'`

````
{
  "name" : "93K6M7P", // 节点名
  "cluster_name" : "elasticsearch", // 集群名
  "cluster_uuid" : "BbIPsUYkRi2-eVrExi7tBQ", // 集群的唯一Id
  "version" : {
    "number" : "7.3.2", // 使用的elasticsearch版本信息
    "build_flavor" : "default",
    "build_type" : "zip",
    "build_hash" : "1c1faf1",
    "build_date" : "2019-09-06T14:40:30.409026Z",
    "build_snapshot" : false,
    "lucene_version" : "8.1.0",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}

````
* **查看集群中文档数量**: `curl -XGET 'http://localhost:9200/_count?pretty'`

````
{
  "count" : 4,
  "_shards" : {
    "total" : 7,
    "successful" : 7,
    "skipped" : 0,
    "failed" : 0
  }
}

````
* **添加文档**:使用`PUT`方式进行请求：`curl -X PUT 地址:端口/索引名/类型名/Id值`,提交的数据内容为json数据

````
请求内容：
curl -X PUT "localhost:9200/megacorp/employee/1?pretty" -H 'Content-Type: application/json' -d'
{
    "first_name" : "John",
    "last_name" :  "Smith",
    "age" :        25,
    "about" :      "I love to go rock climbing",
    "interests": [ "sports", "music" ]
}
'

应答内容:
{
  "_index" : "megacorp", // 索引名
  "_type" : "employee", // 类型名
  "_id" : "1", // 唯一Id
  "_version" : 1, // 版本号，每进行一次操作，版本+1
  "result" : "created", // 操作结果，如果数据不存在，则是created,如果数据存在，则为updated
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 3,
  "_primary_term" : 3
}

````

* **检索文档**: 检索Id为1的文档 `curl -X GET "localhost:9200/megacorp/employee/1?pretty"`

````
应答内容:
{
  "_index" : "megacorp",
  "_type" : "employee",
  "_id" : "1",
  "_version" : 2,
  "_seq_no" : 3,
  "_primary_term" : 3,
  "found" : true,
  "_source" : {
    "first_name" : "John",
    "last_name" : "Smith",
    "age" : 25,
    "about" : "I love to go rock climbing",
    "interests" : [
      "sports",
      "music"
    ]
  }
}

````

* **搜索出所有文档**:`GET "localhost:9200/{索引名}/{类型名}/_search"`

* **单字段简单搜索**:`GET "localhost:9200/{索引名}/{类型名}/_search?q={字段名}:{字段值}"`

* **查询表达式搜索**:查询表达式是一个elasticsearch自定义的查询语法

````
查询出last_name包含Smith的员工
curl -X GET "localhost:9200/megacorp/employee/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "query" : {
        "match" : {
            "last_name" : "Smith"
        }
    }
}
'

查询出last_name包含Smith 且年龄大于30的员工
curl -X GET "localhost:9200/megacorp/employee/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "query" : {
        "bool": {
            "must": {
                "match" : {
                    "last_name" : "smith" 
                }
            },
            "filter": {
                "range" : {
                    "age" : { "gt" : 30 } 
                }
            }
        }
    }
}
'

精准匹配：查询出about为中包含"rock climbing"这个短语的项
curl -X GET "localhost:9200/megacorp/employee/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "query" : {
        "match_phrase" : {
            "about" : "rock climbing"
        }
    }
}
'

匹配信息高亮展示：
curl -X GET "localhost:9200/megacorp/employee/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "query" : {
        "match_phrase" : {
            "about" : "rock climbing"
        }
    },
    "highlight": {
        "fields" : {
            "about" : {}
        }
    }
}
'
结果:
{
   ...
   "hits": {
      "total":      1,
      "max_score":  0.23013961,
      "hits": [
         {
            ...
            "_score":         0.23013961,
            "_source": {
               "first_name":  "John",
               "last_name":   "Smith",
               "age":         25,
               "about":       "I love to go rock climbing",
               "interests": [ "sports", "music" ]
            },
            "highlight": {
               "about": [
                  "I love to go <em>rock</em> <em>climbing</em>" 
               ]
            }
         }
      ]
   }
}
````
说明：
  * match是查询搜索，搜索的结果默认按照相关性得分排序，即每个文档跟查询的匹配程度。
  * filter是查询过滤器

# C#使用elasticsearch

1. 通过nuget安装`NEST`,它是一个经过封装的elasticsearch的客户端
2. 创建elastic客户端
````
	// 单节点连接方式
	var node = new Uri("http://localhost:9200");
    var settings = new ConnectionSettings(node);
    var client = new ElasticClient(settings);

	// 多节点连接方式
	var nodes = new Uri[]
	{
	    new Uri("http://host1:9200"),
	    new Uri("http://host2:9200"),
	    new Uri("http://host3:9200")
	};
	
	var pool = new StaticConnectionPool(nodes);
	var settingsForPool = new ConnectionSettings(pool);
	var client = new ElasticClient(settingsForPool);
````


#参考资料
* [全文搜索引擎 Elasticsearch 入门教程](http://www.ruanyifeng.com/blog/2017/08/elasticsearch.html)
* [C#如何使用Elasticsearch](https://www.cnblogs.com/yinzhou/p/7479315.html)
* [elasticsearch官方教程](https://www.elastic.co/guide/cn/elasticsearch/guide/current/foreword_id.html)
* [NEST包的使用教程-elasticsearch官方文档](https://www.elastic.co/guide/en/elasticsearch/client/net-api/current/nest-getting-started.html)