系统管理
------------------

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
