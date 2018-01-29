Python在windows下版本共存
-----------------------------
# 使用py实现
在安装python3(>=3.3)时，Python的安装包实际上在系统中安装了一个启动器py.exe，默认放置在文件夹C:\Windows\下面。这个启动器允许我们指定使用Python2还是Python3来运行代码（当然前提是你已经成功安装了Python2和Python3）

## 执行脚本
```
py -2/-3 待执行的文件.py
```
* -2代表使用python2执行
* -3代表使用python3执行

也可以在脚本中指定到底使用python2执行还是python3执行
具体：
```
#! python2
# coding: utf-8
```

## 下载依赖库
* 使用python2下载

```
py -2 -m pip install 依赖库名称
```
* 使用python3下载

```
py -3 -m pip install 依赖库名称
```