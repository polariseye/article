cargo ndk 问题解决笔记
-------------------------

* 编译项目报:`Failed to find tool. Is '/usr/local/android-ndk-r23c/toolchains/llvm/prebuilt/linux-x86_64/bin/arm-linux-androideabi-ar' installed?`
 添加一个文件链接到llvm-ar即可:
````
	ln -s /usr/local/android-ndk-r23c/toolchains/llvm/prebuilt/linux-x86_64/bin/llvm-ar /usr/local/android-ndk-r23c/toolchains/llvm/prebuilt/linux-x86_64/bin/arm-linux-androideabi-ar
````
 

* 编译项目报:`ld: error: unable to find library -lgcc`
 可以参照[https://github.com/bbqsrc/cargo-ndk/issues/22](https://github.com/bbqsrc/cargo-ndk/issues/22).是由于新版本生命名了libgcc为libunwind。只需要添加软链接即可:
  ````
	ln -s /usr/local/android-ndk-r23c/toolchains/llvm/prebuilt/linux-x86_64/lib64/clang/12.0.9/lib/linux/arm/libuunwind.a /usr/local/android-ndk-r23c/toolchains/llvm/prebuilt/linux-x86_64/lib64/clang/12.0.9/lib/linux/aarch64/libgcc.a

	ln -s /usr/local/android-ndk-r23c/toolchains/llvm/prebuilt/linux-x86_64/lib64/clang/12.0.9/lib/linux/arm/libuunwind.a /usr/local/android-ndk-r23c/toolchains/llvm/prebuilt/linux-x86_64/lib64/clang/12.0.9/lib/linux/i386/libgcc.a

	ln -s /usr/local/android-ndk-r23c/toolchains/llvm/prebuilt/linux-x86_64/lib64/clang/12.0.9/lib/linux/arm/libuunwind.a /usr/local/android-ndk-r23c/toolchains/llvm/prebuilt/linux-x86_64/lib64/clang/12.0.9/lib/linux/x86_64/libgcc.a

	ln -s /usr/local/android-ndk-r23c/toolchains/llvm/prebuilt/linux-x86_64/lib64/clang/12.0.9/lib/linux/arm/libuunwind.a /usr/local/android-ndk-r23c/toolchains/llvm/prebuilt/linux-x86_64/lib64/clang/12.0.9/lib/linux/arm/libgcc.a
  ````

* 使用flutter_rust_bridge_codegen 提示:`ffigen could not find LLVM.`<br>
  是由于未识别到llvm或未找到clang.so文件。指定llvm_path需包含NDK相关目录:
````
 - '/usr/local/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64'
 - '/usr/local/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/bin'
 - '/usr/local/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/include'
 - '/usr/local/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/lib'
 - '/usr/local/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/lib64'
 - '/usr/local/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/libexec'
 - '/usr/local/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/share'
 - '/usr/local/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/lib64/libclang.so'
````
另外，由于clang默认名字为:`libclang.so.13`，所以需要添加一个到libclang.so的软链接:`ln -s /usr/local/android-ndk-r24/toolchains/llvm/prebuilt/linux-x86_64/lib64/libclang.so.13 libclang.so`

* 编译报:`'main' panicked at 'called `Result::unwrap()` on an `Err` value: Os { code: 2, kind: NotFound, message: "No such file or directory" }', /usr/local/cargo/registry/src/github.com-1ecc6299db9ec823/openssl-src-111.15.0+1.1.1k/src/lib.rs:469:39`<br>
 先执行`sudo apt install pkg-config`<br>
 参考官方issue[1455](https://github.com/sfackler/rust-openssl/issues/1455)。是由于make没有安装导致.执行`apt install make`即可

* ubuntu 下执行`cargo check`.提示:`could not find system library 'openssl' required by the 'openssl-sys'`<br>
 是由于本地未正确安装openssl导致，具体可以参考文档[https://blog.csdn.net/Javin_L/article/details/94740996](https://blog.csdn.net/Javin_L/article/details/94740996)。操作步骤如下:
````
# 安装aptitude 包管理工具
sudo apt install aptitude
# 使用aptitude安装libssl.并进行包降级. 注意，第一次提示是否接受该解决方案时，输入n。其他时候输入y
sudo aptitude install libssl-dev
````
  