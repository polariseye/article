---
title: "任务管理"
date: 2017-08-03 11:13:55
draft: true
---

ubuntu 在用命令整理
-------------------------

## 控制台管理
* 当前所有任务查看

  ``` jobs -l ```

* 让任务到后台执行

  ```command+&```
 
  比如：```./start.sh& ```

* 切换正在运行的程序到后台

  如果程序正在前台运行，可以使用 Ctrl+z 选项把程序暂停，然后用 bg %[number] 命令把这个程序放到后台运行
 
  1. ctrl + z可以将一个正在前台执行的命令放到后台，并且暂停
  2. 察看jobs使用jobs或ps命令可以察看正在执行的jobs
  3. bg将一个在后台暂停的命令，变成继续执行如果后台中有多个命令，可以用bg %jobnumber将选中的命令调出

* 把后台任务调整为执行状态

  用 fg %[number] 指令把一个程序掉到前台运行

* 把任务脱离终端，一直在后台执行

  ```nohup 执行打命令 &``` 执行这个指令后，会让程序脱离控制终端执行

  如果要查看脱离终端执行打任务，可以使用指令：```ps -ef|grep "执行打命令名" ```

* 终止任务执行

  也可以直接终止后台运行的程序，使用 kill 命令

  ``` [oracle@isgis121 ~]$ kill %1 ```

  但是如果任务被终止了(kill)，shell 从当前的shell环境已知的列表中删除任务的进程标识;也就是说，jobs命令显示的是当前shell环境中所起的后台正在运行或者被挂起的任务信息。


