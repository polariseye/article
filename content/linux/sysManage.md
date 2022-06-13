---
title: "系统管理"
date: 2021-09-26 12:17:11
draft: true
---

## 系统环境配置

/etc/profile: 每个用户第一次登录时，执行

/etc/bashrc: 为每一个运行bash shell打用户执行此文件，在bash shell被打开时，执行

## 用户环境配置

~/.bash_profile: 每个用户都可使用该文件输入专用于自己使用的shell信息, 当用户登录时,该文件仅仅执行一次! 默认情况下,他设置一些环境变量,执行用户的.bashrc文件. 

~/.bashrc: 该文件包含专用于你的bash shell的bash信息,当登录时以及每次打开新的shell时,该文件被读取.

~/.bash_logout:用户登出时，执行此文件

# 用户管理
 * **usermod -G 用户组名 用户名**：把用户移动到指定用户组中。会退出之前的用户组
 * **usermod -G -a 用户组名 用户名**：把用户添加到指定用户组中。之前的用户组会保留
 * **groups 用户名**：查看指定用户现在在哪些组里面
 

# 软件管理(ubuntu)

* 下载安装软件 apt-get install <软件名>
* 更新软件源列表 apt-get update
* 更新所有软件 apt-get dist-upgrade
* 删除软件和配置文件 apt-get --purge remove <软件名> 如果，要保留配置文件，则不需要--purge

# 进程管理

## 进程后台运行的几种方式

* 命令行后面加"&"

  比如: "rm -rf &"，作用是把当前任务放置在后台执行，可以使用命令"jobs"进行查看

* 格式：(命令 &)

  新开一个shell执行作业，所提交的作业并不在作业列表中，是无法通过jobs来查看的

* 格式：nohup 命令&

  忽略hangup信号，防止shell关闭时程序停掉

* 格式：setsid 命令

  脱离当前shell的进程关系,放置在后台执行。

* 格式：disown -h job名

  使某个正在运行打作业忽略HUP信号
  
# 磁盘管理
1. 列出当前系统中所有已挂载文件系统的类型：
```
sudo blkid
```

2. 显示指定设备 UUID：
```
sudo blkid -s UUID /dev/sda5
```

3. 显示所有设备 UUID：
```
sudo blkid -s UUID
```

4. 显示指定设备 LABEL：
```
sudo blkid -s LABEL /dev/sda5
```

5. 显示所有设备 LABEL：
```
sudo blkid -s LABEL
```

6. 显示所有设备文件系统：
```
sudo blkid -s TYPE
```

7. 显示所有设备：
````
sudo blkid -o device
````

8. 以列表方式查看详细信息：
```
sudo blkid -o list
```

9. 查看当前磁盘情况
```
fdisk -l
```

详细参见资料:
* [linux 磁盘挂载及查看磁盘](https://www.cnblogs.com/mangoVic/p/7161548.html)

## 服务管理
* 启动服务
```
systemctl start kafka
```

* **手动添加服务(以kafka.service为例)**
	1. 添加文件:/lib/systemd/system/kafka.service 内容如下:
````
	[Unit]
	Description=kafka
	After=network.target remote-fs.target nss-lookup.target zookeeper.service
	 
	[Service]
	Type=forking
	Environment="PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/jdk1.8.0_201/bin"
	ExecStart=/usr/local/kafka_2.12-2.1.1/bin/kafka-server-start.sh -daemon /usr/local/kafka_2.12-2.1.1/config/server.properties
	ExecReload=/bin/kill -s HUP $MAINPID
	ExecStop=/usr/local/kafka_2.12-2.1.1/bin/kafka-server-stop.sh
	PrivateTmp=true
	 
	[Install]
	WantedBy=multi-user.target
````
	2. 刷新配置: ``sudo systemctl daemon-reload``
	3. 启动服务
````
#启动
systemctl start kafka
#查看状态
systemctl status kafka -l
#停止
systemctl stop kafka
````

* 设置服务自启动(以服务kafka.service为例)
````
[root@localhost ~]# systemctl is-enabled kafka.service
disabled
[root@localhost ~]# systemctl enable crond.service
[root@localhost ~]# systemctl is-enabled crond.service
enabled
````
参考资料
* [CentOS7 配置kafka为服务](https://blog.csdn.net/yeshengming2/article/details/88848071)

# 防火墙配置
* 查看firewall服务状态
````
systemctl status firewalld
````
* 查看firewall服务状态
````
firewall-cmd --state
````
* firewall服务管理
````
# 开启
service firewalld start
# 重启
service firewalld restart
# 关闭
service firewalld stop
````
* 查看防火墙规则
````
firewall-cmd --list-all 
````
* 查询，开放，关闭端口
````
	# 查询端口是否开放
	firewall-cmd --query-port=8080/tcp
	# 开放80端口
	firewall-cmd --permanent --add-port=80/tcp
	# 移除端口
	firewall-cmd --permanent --remove-port=8080/tcp
	
	#重启防火墙(修改配置后要重启防火墙)
	firewall-cmd --reload
	
	# 参数解释
	1、firwall-cmd：是Linux提供的操作firewall的一个工具；
	2、--permanent：表示设置为持久；
	3、--add-port：标识添加的端口；
````
# 网络相关
* 系统端口占用查看：
````
netstat -tunlp
````

# VMware虚拟机管理
* 与主机之间共享文件夹
````
# 查看当前已经共享的目录
vmware-hgfsclient
# 挂载所有共享到hgfs目录
vmhgfs-fuse .host:/ /mnt/hgfs -o subtype=vmhgfs-fuse,allow_other
````
* ubuntu无法自适应窗口的解决方案
````
第一步： sudo apt-get autoremove open-vm-tools
第二步：sudo apt-get install open-vm-tools-desktop
````

# 其他问题解决
* **安装的minial版本的centos,不能使用ifconfig与netstat**
````
yum install net-tools
````
* 才安装完成后没有自动分配IP
````
# 重新分配IP
dhclient
# 开机自启动
systemctl enable network.service
````