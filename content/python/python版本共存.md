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

## python 虚拟环境
虚拟环境是基于virtualenv进行环境隔离。每次创建新的环境时，会进行一次基础库的download

### virtualenv安装
```
pip install virtualenv
``` 
### 为自己的项目创建独立的环境
初始化环境
```
virtualenv --no-site-packages venv
```
其中**venv**为自己想要的名称

重要参数说明:
* –no-site-packages: 已经安装到系统Python环境中的所有第三方包都不会复制过来，这样，我们就得到了一个不带任何第三方包的“干净”的Python运行环境
* --system-site-packages: 仍然沿用系统已有的所有三方库

### 进入独立环境
虚拟环境需要先进入虚拟环境，然后即可直接使用

linux里面:
```
source venv/bin/activate
```

windows里面：

执行文件:``` venv/Scripts/activate.bat ``` 或 ```venv/Scripts/activate.ps1```，此时就会在命令行首出现特舒字符串**(venv)**

### 使用独立环境
直接使用相关指令即可

### 退出独立环境
```
deactive
```