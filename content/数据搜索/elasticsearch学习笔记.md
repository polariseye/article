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


**文档元数据**<br />
一个文档由三个必须的元数据组成:
* _index:文档所在的索引名。**索引名必须小写，且不能以下划线开头，不能包含逗号**
* _type:文档的对象类型名。**可以是大写或者小写，但是不能以下划线或者句号开头，不应该包含逗号， 并且长度限制为256个字符.**
* _id:文档唯一标识。id是一个字符串，如果未填写，则elasticsearch会自动生成。**这个Id必须在同一索引同一类型下唯一**

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
## 文档变更操作

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
	
	也可以使用
	curl -X PUT "localhost:9200/megacorp/employee/1/_create?pretty" -H 'Content-Type: application/json' -d'
	{
	   "first_name" : "John",
	    "last_name" :  "Smith",
	    "age" :        25,
	    "about" :      "I love to go rock climbing",
	    "interests": [ "sports", "music" ]
	}
	'
	
	````
	如果没有Id值，则可以使用POST方式提交，这种方式会自动生成一个ID

* **更新文档**

	可以使用相同Id进行再次我加文档进行更新。实际上，elasticsearch文档是不可改变的。使用相同的Id直接使用新文档去覆盖之前的文档会导致重建索引，而之前的文档仍然存在。之前文档的数据会在elasticsearch垃圾回收时清理 
	
	更新文档时，可以使用参考version来进行基于乐观式的版本控制,如下
	```
	curl -X PUT "localhost:9200/website/blog/1?version=1&pretty" -H 'Content-Type: application/json' -d'
	{
	  "title": "My first blog entry",
	  "text":  "Starting to get the hang of this..."
	}
	'
	```
	其中的`version=1`指定为是针对版本号为1的进行操作

* **文档的部分更新**

	部分更新仍然是基于*检索-修改-重建索引*的处理过程。只是此过程只发生在分片内部。这样就可以减少很多网络请求
	```
	curl -X POST "localhost:9200/megacorp/employee/1/_update?pretty" -H 'Content-Type: application/json' -d'
	{
	   "doc" : {
			"first_name" : "hello",
	    	"last_name" :  "world",
	   }
	}
	'
	```
	需要更新的字段作为**doc**的值。它代表这些字段会与整个文档合并然后再覆盖整个文档。
	
	也可以使用脚本来进行文档的更新操作。elasticsearch采用了*Groovvy*脚本语言。**但是Groovy脚本引擎存在漏洞。所以如无必要，不建议开启脚本引擎**
	
	如果需要文档不存在就创建的更新操作，则可以使用**upsert**参数：
	```
	curl -X POST "localhost:9200/megacorp/employee/1/_update?pretty" -H 'Content-Type: application/json' -d'
	{
	   "doc" :{
			"first_name" : "hello",
	    	"last_name" :  "world",
		},
	   "upsert": {
			"first_name" : "hello2",
	    	"last_name" :  "world2",
	   }
	}
	'
	```


* **删除文档**：`curl -X DELETE "localhost:9200/{索引名}/{类型名}/{Id值}?pretty"`。如果文档不存在，则会返回404.此操作会更新版本号，而不管文档是否存在

## 查找文档

查询分为查询阶段与取数据阶段。
    
* 查询阶段： 外部请求的接收节点会以协调节点的身份，把请求查询阶段会把请求转给每个分片节点（可能是副本节点也可能是主节点）。每个节点会分配size*from的空间来进行查询结果的缓存。每个节点查询完成后，再把文档Id与排序字段信息发送给协调节点，节点再对所有数据进行排序。

* 取数据阶段：协调节点从查询返回的所有文档基本信息中找到需要返回的文档，然后再用Id去相关节点中取回文档的具体数据并返回给外部调用者。

查询
* **检索所有文档**: 检索Id为1的文档 `curl -X GET "localhost:9200/megacorp/employee/1?pretty"`
	
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
* **通过Id查找文档**:`GET "localhost:9200/{索引名}/{类型名}/{Id}`
* **通过Id查找文档且只返回`source`的内容**:`GET "localhost:9200/{索引名}/{类型名}/{Id}/_source`
* **一次检索多个文档**:使用 multi-get 或者 mget实现一次检索多个文档
	```
	curl -X GET "localhost:9200/_mget?pretty" -H 'Content-Type: application/json' -d'
	{
	   "docs" : [
	      {
	         "_index" : "website",
	         "_type" :  "blog",
	         "_id" :    2
	      },
	      {
	         "_index" : "website",
	         "_type" :  "pageviews",
	         "_id" :    1,
	         "_source": "views"
	      }
	   ]
	}
	'
	```
	如果是同一个索引，同一个类型。可以简写如下
	```
	curl -X GET "localhost:9200/website/blog/_mget?pretty" -H 'Content-Type: application/json' -d'
	{
	   "ids" : [ "2", "1" ]
	}
	'
	```

* **返回文档的一部分**：`GET /{索引名}/{类型名}/{Id}?_source={字段名1},{字段名2}`

* **查看文档是否存在**：`curl -i -XHEAD {索引名}/{类型名}/{Id}`,如果文档存在，则返回http状态码200，不存在则返回404。由于多节点原因。不存在只是代表那一时刻不存在，或许这文档正在各节点同步过程中

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
	在多个字段上进行相同的查询
	curl -X GET "localhost:9200/megacorp/employee/_search?pretty" -H 'Content-Type: application/json' -d'
	{
	    "query" : {
		    "multi_match": {
		        "query":    "hello",
		        "fields":   [ "name", "last_name" ]
		    }
		}
	}
	'
	
	使用range进行范围查询 查询出last_name包含Smith 且年龄大于30的员工
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
	
	多条件查询
	curl -X GET "localhost:9200/megacorp/employee/_search?pretty" -H 'Content-Type: application/json' -d'
	{
	    "bool": {
	        "must":     { "match": { "title": "how to make millions" }},
	        "must_not": { "match": { "tag":   "spam" }},
	        "should": [
	            { "match": { "tag": "starred" }}
	        ],
	        "filter": {
	          "range": { "date": { "gte": "2014-01-01" }} 
	        }
	    }
	}
	
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
	  * filter是查询过滤器,fiter是精确查找，不会进行相关度计算。执行速度非常快
	  * match_phase是短语完整匹配
	  * bool用于组合多个条件进行查询
	  * multi_match:用于在多个字段上进行相同的查询
	  * must:文档必须满足这些条件才能包含进来
	  * must_not:文档必须不匹配这些条件才能包含进来
	  * should:如果满足这些语句中的任意语句，将增加 _score ，否则，无任何影响。它们主要用于修正每个文档的相关性得分
	  * range进行范围查询
	    * gt:大于
	    * gte:大于等于
	    * lt:小于
	    * lte:小于等于
	   * const_score:将一个不变的常量评分应用于所有匹配的文档。经常用于只有filter而没有其他查询的情况 

* **精确查找**

	进行精确查找时，需要使用过滤器(filters)。**过滤器不会进行相关度计算**（直接跳过评分阶段），所以执行速度非常快。因此，平时应该尽可能多使用过滤器。  
	
	**使用term进行精确匹配**
	```
	curl -X GET "localhost:9200/my_store/products/_search?pretty" -H 'Content-Type: application/json' -d'
	{
	    "query" : {
	        "constant_score" : { 
	            "filter" : {
	                "term" : { 
	                    "price" : 20
	                }
	            }
	        }
	    }
	}
	'
	```
	term会查找指定的精确查询，可以是数值与文本。也可以使用`terms`进行多值匹配的查询
	```
	curl -X GET "localhost:9200/my_store/products/_search?pretty" -H 'Content-Type: application/json' -d'
	{
	    "query" : {
	        "constant_score" : { 
	            "filter" : {
	                "terms" : { 
	                    "price" : [20,30]
	                }
	            }
	        }
	    }
	}
	'
	```
* **指定查询返回字段**    
	使用`_source`指定返回字段
	````
	curl -X GET "localhost:9200/my_store/products/_search?pretty" -H 'Content-Type: application/json' -d'
	{
	    "query" : {
	        "constant_score" : { 
	            "filter" : {
	                "terms" : { 
	                    "price" : [20,30]
	                }
	            }
	        }
	    },
		"_source":['price','id']
	}
	'
	````
	使用`_source`指定包含与排除的字段
	````
	curl -X GET "localhost:9200/my_store/products/_search?pretty" -H 'Content-Type: application/json' -d'
	{
	    "query" : {
	        "constant_score" : { 
	            "filter" : {
	                "terms" : { 
	                    "price" : [20,30]
	                }
	            }
	        }
	    },
		"_source":{
			"includes":['price','id'],
			"excludes":["desc"]
		}
	}
	'
	````
	使用查询参数指定,查询参数中指定时，可以使用`_source`与`fields`来指定，效果相同
	````
	curl -X GET "localhost:9200/my_store/products/_search?pretty&_source=price,id" -H 'Content-Type: application/json' -d'
	{
	    "query" : {
	        "constant_score" : { 
	            "filter" : {
	                "terms" : { 
	                    "price" : [20,30]
	                }
	            }
	        }
	    }
	}
	'
	````

* **exists查询与missing查询**<br/>
    exists用于查询文档中指定字段是否有值，与sql中的NOT IS_NULL()一致

    missing用于查询文档中指定字段是否无值，与sql中的IS_NULL()一致
	```
	curl -X GET "localhost:9200/my_store/products/_search?pretty" -H 'Content-Type: application/json' -d'
	{
	    "query" : {
	        "constant_score" : { 
	            "filter" : {
	                "exists" : { 
	                    "field" : "title"
	                }
	            }
	        }
	    }
	}
	'
	```
* **使用_validate验证查询语句的正确性**
	```
	curl -X GET "localhost:9200/my_store/products/_validate/query?explain" -H 'Content-Type: application/json' -d'
		{
		    "query" : {
		        "constant_score" : { 
		            "filter" : {
		                "terms" : { 
		                    "price" : [20,30]
		                }
		            }
		        }
		    }
		}
		'
	```
	说明:<br/>
	 * explain参数用于得到不正确的原因


* **查询排序**   
	一般查询结果是按每项的`_score`来进行**降序排序**，也称为按照相关性排序。我们也可以使用`sort`参数进行指定排序。
	```
按照创建时间降序排序
curl -X GET "localhost:9200/my_store/products/_validate/query?explain" -H 'Content-Type: application/json' -d'
{
    "query" : {
        "bool" : {
            "filter" : { "term" : { "user_id" : 1 }}
        }
    },
    "sort": [
		{ "create_time": { "order": "desc" }},
		{ "_score": { "order": "desc" }}
	]
}
	```
如果指定了`sort`但排序字段中没有指定`_score`字段，则相关性分数不会被计算

* **查询分页**    
	```
	curl -X GET "localhost:9200/_search?size=5&pretty"
	curl -X GET "localhost:9200/_search?size=5&from=5&pretty"
	curl -X GET "localhost:9200/_search?size=5&from=10&pretty"
	```
	分页使用参数`size`与`from`。`size`指定返回的结果数量，默认值10.`from`指定需要跳过的记录数。<br/>
	分页注意事项:
	1. 一个请求经常跨越多个分片，每个分片都产生自己的排序结果。对结果排序的成本随分页的深度成指数上升。所以建议不要超过1000个结果。


* **使用scroll进行大量数据的获取**
   
	由于使用普通分页方式查询会导致资源消耗极大。所以，在进行大数据量查询时，使用游标查询来提升性能。    
	查询使用方式:    
	 1. scroll查询只需要添加请求参数:`scroll=1m`。其中`1m`代表游标查询结果缓存的时长为1分钟。
	 2. 查询会返回一个名为`scroll_id`的字段。接着取数据时，使用之前返回的`scroll_id`接着取数据。
	 ````
	GET /_search/scroll
	{
	    "scroll": "1m", // 此处会重置游标过期时间 
	    "scroll_id" : "cXVlcnlUaGVuRmV0Y2g7NTsxMDk5NDpkUmpiR2FjOFNhNnlCM1ZDMWpWYnRROzEwOTk1OmRSamJHYWM4U2E2eUIzVkMxalZidFE7MTA5OTM6ZFJqYkdhYzhTYTZ5QjNWQzFqVmJ0UTsxMTE5MDpBVUtwN2lxc1FLZV8yRGVjWlI2QUVBOzEwOTk2OmRSamJHYWM4U2E2eUIzVkMxalZidFE7MDs="
	}
	````	


## 批量操作
批量操作使用`bulk`API。 它允许在单个步骤中进行多次 create 、 index 、 update 或 delete 请求。请求格式：
```
{ action: { metadata }}\n
{ request body        }\n
{ action: { metadata }}\n
{ request body        }\n
```
说明:
 1. 每行一定要以换行符(\n)结尾， 包括最后一行 
 2. 这些行不能包含未转义的换行符，因为他们将会对解析造成干扰

action取值:
* **create**:如果文档不存在，那么就创建它
* **index**:创建一个新文档或者替换一个现有的文档
* **update**:部分更新一个文档
* **delete**:删除一个文档。删除操作不需要request body行

具体示例如下:
```
curl -X POST "localhost:9200/_bulk?pretty" -H 'Content-Type: application/json' -d'
{ "delete": { "_index": "website", "_type": "blog", "_id": "123" }} 
{ "create": { "_index": "website", "_type": "blog", "_id": "123" }}
{ "title":    "My first blog post" }
{ "index":  { "_index": "website", "_type": "blog" }}
{ "title":    "My second blog post" }
{ "update": { "_index": "website", "_type": "blog", "_id": "123", "_retry_on_conflict" : 3} }
{ "doc" : {"title" : "My updated blog post"} }
'
```
## 类型映射
类型映射信息存储了每个索引中的类型的每个字段的数据类型。

查看类型映射
```
curl -X GET "localhost:9200/{索引名}/_mapping/{类型名}"
```

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
* [elasticsearch中文社区](https://elasticsearch.cn/) 