python通用配置
--------------------
# 依赖库管理
python 一般使用pip 进行依赖库管理

由于国内使用python会很慢，所以建议使用国内镜像，更改国内镜像的方式如下：

* Unix:$HOME/.config/pip/pip.conf
* Mac: $HOME/Library/Application Support/pip/pip.conf
* Windows：%APPDATA%\pip\pip.ini，%APPDATA%的实际路径我电脑上是C:\Users\user_xxx\AppData\Roaming，可在cmd里执行echo %APPDATA%命令查看
