prometheus学习笔记
----------------------------------------------------------------
# 概要
Prometheus 是由 SoundCloud 开源监控告警解决方案。官方地址:[点击前往](https://prometheus.io/docs/alerting/latest/alertmanager/)，代码仓库地址:[点击前往](https://github.com/prometheus/prometheus)。既kubernete之后，于2016年加入[云本地计算基金(Cloud Native Computing Foundation)](https://cncf.io/)。具有如下特点

* 存储的时序数据，通过使用指标(metric)与键值对(label)来定义时间序列数据

* 使用prometheus自己开发的的数据查询语言PromQL来进行数据查询

* 不依赖于分布式存储，可以单节点部署

* 使用基于HTTP的PULL拉取模型收集时间序列的数据

* 通过配套的网关节点，也可以主动推送时序数据

* 监控节点可以通过服务发现服务或者静态配置加入到监控
* 支持多模型的图形化界面与数据分析界面

主要包含组件如下：
* [主服务节点 prometheus server](https://github.com/prometheus/prometheus)，主要功能包含提供收集与存储时序数据

* 提供进程监控的[客户端库](https://prometheus.io/docs/instrumenting/clientlibs/)

* 一个用来提供[主动推送的网关](https://github.com/prometheus/pushgateway)

* 把被监控组件的数据输出到prometheus的组件称为[exporter](https://prometheus.io/docs/instrumenting/exporters/),prometheus提供各种组件的[exporter](https://prometheus.io/docs/instrumenting/exporters/)，比如:HAProxy,StatsD,Graphite等

* 一个用于处理报警的[报警管理服务](https://github.com/prometheus/alertmanager)

当然，prometheus还拥有其他丰富的组件，具体可以去[github](https://github.com/prometheus/prometheus)或[官网](https://prometheus.io)里面寻找

整体架构如下:
![](./prometheus_files/architecture.png)

# prometheus数据模型
prometheus按照时间序列去存储数据，相同的metric(指标，一个float64值)和多个label(key/value数据对)组成一条时间序列数据。

**metric(指标)**

指标名是指需要监控的对象的名字，一个指标由多个键值对构成，官方把这些键值对称呼为label。一般使用监控对象的类型名开头，如:`http_requests_total`，必须满足正则表达式 [a-zA-Z_:][a-zA-Z0-9_:]+

 * 指标命名规则:可以包含字母、数字、下划线，冒号

 * 最终形成的时间序列数据形如：http_requests_total{method="POST",endpoint="/api/tracks"}

与关系型数据库类比，指标名相当于表名，键值对(label)相当于字段,timestamp相当于主键,另外，**每个指标还有一个float64的值，表示指标的实际值**

# prometheus数据类型

**Counter**<br />
用于累计值，如计数类。**一直增加**，不会减少进程重启后会被重置

**Gauge**<br />
常规数值，**可加可减**,其为瞬时的，与时间没有关系的，可以任意变化的数据。重启进程后，会被重置

**Histogram**<br />
Histogram（直方图）可以理解为柱状图的意思，常用于跟踪事件发生的规模,例如：请求耗时、响应大小。它特别之处是可以对记录的内容进行分组，提供count和sum全部值的功能。**其主要用于表示一段时间内对数据的采样，并能够对其指定区间及总数进行统计。根据统计区间计算**

**Summary**<br />
与Histogram相似，用于表示一段时间内数据采样结果，而不是根据统计区间计算出来的。它提供一个quantiles的功能，可以按%比划分跟踪的结果。例如：quantile取值0.95，表示取采样值里面的95%数据。**不需要计算，直接存储结果**

# 安装
以监控本机为例，安装流程如下:

1. **下载镜像**
````
docker pull prom/prometheus
docker pull prom/node-exporter
````
2. **启动exporter**
````
docker run -d -p 9100:9100 \
  -v "/proc:/host/proc:ro" \
  -v "/sys:/host/sys:ro" \
  -v "/:/rootfs:ro" \
  --net="host" \
  prom/node-exporter
```` 
3. **启动prometheus**
 1. 创建配置文件`/etc/prometheus/prometheus.yml`,内容如下:
 ````
	global:
	  scrape_interval:     60s
	  evaluation_interval: 60s
	 
	scrape_configs:
	  - job_name: prometheus
	    static_configs:
	      - targets: ['localhost:9090']
	        labels:
	          instance: prometheus
	 
	  - job_name: linux
	    static_configs:
	      - targets: ['192.168.1.14:9100']
	        labels:
	          instance: localhost
 ````
 2. 启动
 ````
	docker run  -d \
	  -p 9090:9090 \
	  -v /etc/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml  \
	  prom/prometheus
 ````
 3. 浏览器中查看:`http://192.168.1.14:9090`

如有需要，也可以安装一个prometheus的可视化工具:[Grafana](https://grafana.com/),安装教程参考[基于docker 搭建Prometheus+Grafana](https://www.cnblogs.com/xiao987334176/p/9930517.html)

## 报警管理进程安装

## 推送网关安装

1. **下载镜像**
````
docker pull prom/alertmanager
````
2. **启动运行**
````
docker run -d \
  --name=pushgateway \
  -p 9091:9091 \
  prom/pushgateway 
````
说明:
 * ``"-persistence.file=push_file"``指定网关数据持久化方式,**但是经过验证，添加这个参数后，程序无法正常启动**

# 参考资料
* [官方网站](https://prometheus.io)
* [Prometheus 简书](https://www.jianshu.com/p/93c840025f01) 
* [基于docker 搭建Prometheus+Grafana](https://www.cnblogs.com/xiao987334176/p/9930517.html)
* [Prometheus监控系统在搜索服务场景中的应用与实践
](https://zhuanlan.zhihu.com/p/87041558?from_voters_page=true)