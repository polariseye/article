flutter_rust_bridge环境搭建笔记
-------------------------
flutter_rust_bridge是一个dart调用rust的代码生成工具。
官方项目地址:[https://github.com/fzyzcjy/flutter_rust_bridge](https://github.com/fzyzcjy/flutter_rust_bridge)
官方文档地址:[http://cjycode.com/flutter_rust_bridge/](http://cjycode.com/flutter_rust_bridge/)
搭建环境:centos

#搭建步骤
1. 安装rust:`curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh`
2. 去[官网](https://developer.android.google.cn/ndk/downloads/index.html)下载ndk:[https://dl.google.com/android/repository/android-ndk-r23c-linux.zip](https://dl.google.com/android/repository/android-ndk-r23c-linux.zip) 
3. 下载flutter,下载安装步骤可参考:[https://docs.flutter.dev/get-started/install](https://docs.flutter.dev/get-started/install).并把[flutter/bin]添加到环境变量:`export PATH=$PATH:/usr/local/flutter/bin`
4. 添加cargo的android编译工具:
	````
	rustup target add \
	    aarch64-linux-android \
	    armv7-linux-androideabi \
	    x86_64-linux-android \
	    i686-linux-android`
	````
5. 安装cargo ndk编译工具:`cargo install cargo-ndk`
6. 下载代码生成工具以及相关依赖:`cargo install flutter_rust_bridge_codegen just`
7. 下载ffigen工具:`dart pub global activate ffigen`
8. 使用flutter_rust_bridge_codegen时，指定参数`llvm_path`包含以下路径:
	````
	- '/usr/local/app/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64'
	- '/usr/local/app/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/bin'
	- '/usr/local/app/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/include'
	- '/usr/local/app/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/lib'
	- '/usr/local/app/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/lib64'
	- '/usr/local/app/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/libexec'
	- '/usr/local/app/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/share'
	- '/usr/lib64/clang-private'
	- '/usr/local/app/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/lib64/libclang.so'
	````

参照文档:[carg ndk问题解决笔记](./carg ndk问题解决笔记.md)解决编译过程中遇到的其他问题