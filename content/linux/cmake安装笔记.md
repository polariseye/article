cmake安装笔记
---------------------------
centos7 下安装cmake
* 下载cmake:`https://cmake.org/download/`
* 解压到`/usr/local/bin/cmake`
* 修改文件:`/etc/profile`,添加到环境变量:
	 ````
	export CMAKE_HOME=/usr/local/bin/cmake
	export PATH=$PATH:$CMAKE_HOME/bin
	 ````
