llvm源码编译笔记
--------------------------------
编译环境: centos 7

**下载编译工具:**
* `yum install centos-release-scl`
* `yum install devtoolset-10-toolchain.x86_64`。 如果这个版本不合适要求，可以使用` yum list devtoolset* `来查看当前的最新版本
* 设置这个版本为默认编译器`echo "source /opt/rh/devtoolset-10/enable">>/etc/profile`
* 重启系统或`source /etc/profile`

**源码编译**
* 下载项目源码: `git clone https://github.com/llvm/llvm-project`
* 切换到目标tag: `git checkout llvmorg-14.0.0-rc4`
* 在项目目录llvm-project下面新建构建目录:`mkdir build`
* 生成编译配置:`cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=/usr/local/llvm -DLLVM_ENABLE_RTTI=ON -DLLVM_ENABLE_PROJECTS="llvm;clang;clang-tools-extra;compiler-rt;libcxx;libcxxabi" -G "Unix Makefiles" ../llvm`
* 执行编译:`make`
* 安装:`make install`

# 问题解决记录

**$'\r': command not found**<br>
原因: 是由于项目在windows下clone导致。<br>
解决办法:删除附.git目录外的所有文件，然后在linux下执行:`git checkout .`<br>

#通过其他源安装llvm
* 添加源: `yum -y install epel-release`
* 更新源: `yum -y update`
* 清理:`yum clean all`
* 缓存:`yum makecache`
* 查看存在的llvm版本:`yum list *llvm*`
* 下载:`yum install llvm11`


# 参考资料
* [llvm官网](http://llvm.org/)
* [llvm-project](https://github.com/llvm/llvm-project)
* [Centos7上源码编译安装llvm 11.0.0](https://zhuanlan.zhihu.com/p/350595463)