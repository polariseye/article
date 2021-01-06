Docker笔记
----------------------------------
* docker images ：列出当前的所有容器
* docker search 镜像名 ：查找远端仓库是否存在指定容器
* docker pull 镜像名：载入镜像，从远端拉取容器 
* docker run -i -t 镜像名:TAG 命名相关参数：运行交互式的容器,例如 docker run -i -t ubuntu /bin/bash
		-i 在新容器内指定一个伪终端或终端
		-t 允许你对容器内的标准输入 (STDIN) 进行交互
* docker run -d 镜像名:TAG 命名相关参数：创建一个后台的容器，并执行一个命名。当运行容器时，使用的镜像如果在本地中不存在，docker 就会自动从 docker 镜像仓库中下载，默认是从 Docker Hub 公共镜像源下载。
		-d 以后台的方式运行一个容器
* docker run -d -p 内部IP:内部端口:主机端口/协议 镜像名:TAG 命名相关参数 以指定的端口映射关系运行指定容器
		-P 内部端口随机映射到主机的高端口
		-p 容器内部端口绑定到指定的主机端口
		说明：
			1. 内部IP部分非必填，此时为: docker run -d -p 内部端口:主机端口/协议 镜像名:TAG 命名相关参数
			2. 默认为tcp，udp协议需要指定，例如: docker run -d -p 3000:5000/udp centos /bin/echo "helloworld"
* docker run -v [主机目录或文件:]容器目录或文件 镜像名:TAG 命名相关参数 以目录映射的方式实现数据持久化
* docker run --volumes-from 数据容器名 镜像名:TAG 命名相关参数 以指定数据容器来运行
* docker ps :查看当前有哪些容器进程
* docker logs 容器Id/容器名:查看指定容器的标准输出
* docker logs 容器名字或者 ID 2>&1 | grep xxx 。 `2>&1`代表 把标准错误（文件描述符2）重定向（>）到标准输出（文件描述符 1）的位置（&）
* docker stop 容器Id/容器名: 停止容器
* docker port 容器Id/容器名:查看指定容器的端口映射
* docker top 容器Id/容器名:查看指定容器内部运行的进程
* docker rm 容器名:删除容器，删除时，容器必须是停止状态
* docker tag 容器Id 原镜像名:新的标签名
* docker save -o 要保存到的文件路径 镜像名 ：将镜像保存为tar文件
* docker import file|URL 镜像名:TAG :从tar文件创建镜像
* docker export -o 目标文件路径.tar 容器Id|容器名 将容器的文件系统作为一个tar归档文件导出到STDOUT
* docker attach 参数 容器Id|容器名 ：连接到正在运行中的容器
		* 可以同时连接上同一个container来共享屏幕
		* detach: CTRL+C 会导致退出容器，并且stop容器，如果想不想stop，则添加参数: --sig-proxy=false
* docker exec [选项] 容器Id|容器名 命名相关参数 ：在指定的容器内执行命令
	* -d:分离模式，在后台运行
	* -i:即使没有附加也保持STDIN 打开
	* -t:分配一个伪终端
* docker inspect 查看容器的元数据信息
* docker volume ls : 列出当前所有的数据容器
* docker volume inspect 数据容器名 :查看数据容器的元数据 
* docker volume create 数控容器名 :使用默认的driver来创建数据容器
* docker create -v [主机目录或文件:]容器目录或文件 --name 数据容器名 镜像名:TAG 命名相关参数(/bin/true) 使用指定镜像来创建一个数据容器

## 自定义镜像
修改已有镜像

	1. 运行已有镜像的bash: docker run -i -t ubuntu:15.10 /bin/bash
	2. 做相关操作，比如: apt-get update
	3. 退出镜像：exit
	4. 查看进程：docker ps
	5. 提交镜像: docker commit -m="提交信息" -a="作者" 容器Id 容器名:TAG
			比如 ：docker commit -m="create httpserver" -a="polariseye" c10cde23a98f polariseye/httpserver:v1

从零开始构建
docker build -t 目标镜像名:TAG DockerFile文件路径
DockerFile格式如下
````
FROM 基于的哪个镜像

````

## docker run参数说明
* -p {内部IP}:{内部端口}:{主机端口}/协议 指定映射和协议
* -d 后台方式运行
* -i 在新容器内指定一个伪终端或终端
* -t 允许你对容器内的标准输入 (STDIN) 进行交互
* --name {进程名} 以指定的名字运行容器
* -v [主机目录或文件]:容器目录或文件 指定目录或文件的映射关系，实现数据持久化
* --volumes-from 数据容器名 ：以指定数据容器来运行容器

## docker的持久存储模式
1. 默认模式 不支持任何持久性
在不指定时，将采用这个模式。在这个模式下，数据容器进程退出后，相关文件系统操作都将会消失。容器进程还在时，操作将持续有效。对应的数据会临时存储在目录/var/lib/docker/volumes

2. 数据卷模式（Docker volume） 容器持久性
指定方式: docker run -v [主机目录:]容器中持久化的目录名 镜像名 命令参数
说明：
	* 如果未指定主机目录，则数据会保留在/var/lib/docker/volumes下面
	* 容器进程退出后，数据不会被删除
	* 保存在默认目录的数据只是一次性的，容器进程退出后将不能再被使用。因为上层目录名是随机生成的

3. 只含数据的容器(data container) 容器持久性
查看容器列表：docker volume list
数据容器在使用前，需要先创建:docker create -v [主机目录或文件]:容器目录或文件 --name 容器名 源镜像 执行的命令。例如 ：docker create -v /home --name socketserver centos /bin/true ,创建一个名为socketserver的数据容器
使用此容器：docker run --volumes-from 数据容器名 镜像名:TAG 命名相关参数 例如：docker run -i -t --volumes-from datacontainer centos /bin/bash

说明：
	* 创建数据容器也可以使用: docker volume create --name 容器名 来创建，这个将使用默认的 `local` driver来创建
	* 查看数据容器的元数据，可以使用指令：docker volume inspect 数据容器名
	* 删除容器时，并删除卷：docker rm -v 容器Id|容器名
	* 单独执行时，会产生很多孤单的卷

4. 从主机映射而得的数据卷 容器持久性
5. 从主机映射而得的数据卷且存储后端是共享存储 主机持久性
6. Convoy 存储插件 主机持久性

说明:
1. 默认模式 不支持任何持久性
1. 第 2-4 种支持容器持久性，即升级容器也不会移除数据
2. 5-6种支持主机持久性，即主机失效也不会引起数据丢失

## docker容器升级

## 相关其他指令
* `cat /etc/os-release` :查看系统版本信息
* `uname -a`:查看系统内核的版本信息
* `cat /proc/version`:查看当前系统内核的版本信息
* `docker ps -a --no-trunc` :查看容器列表详细信息，显示内容不截取，包含完整的command命令信息

## 参考资料
* [Docker容器的持久存储模式 ](http://dockone.io/article/1283)
* [理解Docker（8）：Docker 存储之卷（Volume）](https://www.cnblogs.com/sammyliu/p/5932996.html)
* [Docker数据管理：data container](https://blog.csdn.net/liumiaocn/article/details/52040414)