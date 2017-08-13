系统管理
------------------

# 环境变量配置

环境变量配置均使用```export 要设置打环境变量```

使修改立即生效：source 配置文件名

## 系统环境配置

/etc/profile: 每个用户第一次登录时，执行

/etc/bashrc: 为每一个运行bash shell打用户执行此文件，在bash shell被打开时，执行

## 用户环境配置

~/.bash_profile: 每个用户都可使用该文件输入专用于自己使用的shell信息, 当用户登录时,该文件仅仅执行一次! 默认情况下,他设置一些环境变量,执行用户的.bashrc文件. 

~/.bashrc: 该文件包含专用于你的bash shell的bash信息,当登录时以及每次打开新的shell时,该文件被读取.

~/.bash_logout:用户登出时，执行此文件
 
# 软件管理

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
