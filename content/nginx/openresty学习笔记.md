openresty学习笔记
------------------------------------
#简介
openresty是一个基于nginx与lua的高性能web平台.官网地址:[http://openresty.org/cn/](http://openresty.org/cn/)

# centos开发环境搭建

** 安装openresty **
1. 添加openresty的仓库
	````
	# add the yum repo:
	wget https://openresty.org/package/centos/openresty.repo
	sudo mv openresty.repo /etc/yum.repos.d/
	
	# update the yum index:
	sudo yum check-update
	````
2. 安装
	````
	sudo yum install -y openresty

	// 安装命令行工具
	sudo yum install -y openresty-resty
	````

**nginx+lua 开发环境配置**

1. 修改`nginx.conf`配置文件，在配置文件的http节点下添加相关模块:`/usr/local/openresty/nginx/conf/nginx.conf`
	````
	#lua模块路径，多个之间”;”分隔，其中”;;”表示默认搜索路径
	
	lua_package_path "/opt/openresty/lualib/?.lua;;"; #lua 模块	
	lua_package_cpath "/opt/openresty/lualib/?.so;;"; #c 模块
	```` 
    添加后例如:
	````
	http {
	  include mime.types;
	  default_type application/octet-stream;
	  //.....
	  lua_package_path "/opt/openresty/lualib/?.lua;;"; #lua 模块	
	  lua_package_cpath "/opt/openresty/lualib/?.so;;"; #c 模块	
	}
	````
2. 建议把lua配置存放在单独的配置文件（此步骤可选）:`vim lua.conf`
	````
	server {
        listen       80;
        server_name  localhost;

        location /lua {
            default_type 'text/html';
            content_by_lua 'ngx.say("hello world")';
        }
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }

	````
	修改nginx.conf:`vim nginx.conf`
	````
	#user  nobody;
	worker_processes  1;
	
	events {
	    worker_connections  1024;
	}
		
	http {
	    include       mime.types;
	    default_type  application/octet-stream;
	    #access_log  logs/access.log  main;
	    lua_package_path "/usr/local/openresty/lualib/?.lua;;"; #lua 模块
	    lua_package_cpath "/usr/local/openresty/lualib/?.so;;"; #c 模块
	    sendfile        on;
	    #tcp_nopush     on;
	    #keepalive_timeout  0;
	    keepalive_timeout  65;
	    #gzip  on;
	    include lua.conf; #引入lua的配置
	}
	````
	注意，如果是基于默认的配置调整，则需要把之前80端口的相关配置注释掉，否则会端口冲突
3. 测试配置是否正常:
	````
	cd /usr/local/openresty/bin
	./openresty -t
	````
	如果成功，则返回如下:
	````
	nginx: the configuration file /usr/local/openresty/nginx/conf/nginx.conf syntax is ok
	nginx: configuration file /usr/local/openresty/nginx/conf/nginx.conf test is successful
	````
4. 配置正常后即可访问:[http://127.0.0.1/lua](http://127.0.0.1/lua)

**开发建议**
1. 开发时，关闭lua代码缓存。即每次代码变更时不用执行`nginx -s reload`:修改配置文件`lua.conf`的内容为:
	````
	server {
	        listen       80;
	        server_name  localhost;
			lua_code_cache off; # 默认是打开了lua脚本缓存的，通过此配置关闭缓存
			
	        location /lua {
	            default_type 'text/html';
	            content_by_lua 'ngx.say("hello world you are sb")';
	        }

	        error_page   500 502 503 504  /50x.html;
	        location = /50x.html {
	            root   html;
	        }
	}

	````

2. 我们的Lua代码一般会比较长，如果放在lua.conf文件会让这个配置很难维护，因此，我们放在单独的文件`/usr/local/openresty/lua/test.lua`中:
	````
	ngx.say("hello world you are sb");
	````
	此时，lua.conf内容如下:
	````
		server {
	        listen       80;
	        server_name  localhost;
			lua_code_cache off; # 默认是打开了lua脚本缓存的，通过此配置关闭缓存
			
	        location /lua {
	            default_type 'text/html';
	            content_by_lua_file /usr/local/openresty/lua/test.lua; # 注意，这里是content_by_lua_file 
	        }

	        error_page   500 502 503 504  /50x.html;
	        location = /50x.html {
	            root   html;
	        }
	}
	````
3. 代码分离。以上步骤要求我们自定义的代码需要放在openresty。这会让业务代码不方便维护。我们可以把代码放在单独的目录中。例如:
	````
	~/testsite/conf
	~/testsite/conf/nginx.conf
	~/testsite/conf/lua.conf
	~/testsite/logs
	~/testsite/lua
	~/testsite/lua/test.lua	
	````
	启动基于以上目录结构的项目:`/usr/local/openresty/nginx/sbin/nginx -p ~/testsite/ -c ~/testsite/conf/nginx.conf`<br/>
	其中nginx.conf与lua.conf与之前的内内基本一致

# 参考资料

* [openresty官网](http://openresty.org/cn/)
* [OpenResty运行环境搭建](https://www.cnblogs.com/babycomeon/p/11109501.html)