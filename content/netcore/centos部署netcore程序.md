centos下部署netcore程序
--------------------------------------
# 安装netcore基本环境
1. 添加netcore的源
````
sudo rpm -Uvh https://packages.microsoft.com/config/rhel/7/packages-microsoft-prod.rpm
````
2. 安装netcore
````
yum install dotnet-sdk-3.1
````
3. 安装netcore依赖组件
````
yum install libgdiplus
````
备注:
* 网上有说还需要安装`libc6-dev`，但是我在安装完`libgdiplus`后程序能够正常运行（System.Drawing会使用这个组件）。
 
# 部署netcore程序到centos并以服务的方式运行

以程序在目录`/var/www/hellowrd`为例:

1. 执行`vim /usr/lib/systemd/system/hellowrld.service`,拷贝内容

````
[Unit]
Description=my helloworld service

[Service]
WorkingDirectory=/var/www/helloworld
ExecStart=/usr/bin/dotnet helloworld.dll
Restart=always
RestartSec=10  # Restart service after 10 seconds if dotnet service crashes
SyslogIdentifier=dotnet-helloworld
User=root
Environment=ASPNETCORE_ENVIRONMENT=Production
Environment=DOTNET_PRINT_TELEMETRY_MESSAGE=false

[Install]
WantedBy=multi-user.target
````
2. 执行命令:`systemctl enable hellowrld.service`
3. 执行命令:`systemctl start helloworld`,使helloworld程序运行

# 参考资料
* [centos7下.netcore环境搭建](https://www.cnblogs.com/hanfeige/p/11389158.html)
* [Centos7部署.net core 控制台app为后台服务](https://blog.csdn.net/qq_32688731/article/details/87872710)
* [Linux CentOS 使用service运行.NET Core项目](https://blog.csdn.net/haitaodoit/article/details/86037021)