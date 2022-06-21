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

* 使用flutter_rust_bridge_codegen 提示:`ffigen could not find LLVM.`
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

* 编译报:`'main' panicked at 'called `Result::unwrap()` on an `Err` value: Os { code: 2, kind: NotFound, message: "No such file or directory" }', /usr/local/cargo/registry/src/github.com-1ecc6299db9ec823/openssl-src-111.15.0+1.1.1k/src/lib.rs:469:39`
* sudo apt install pkg-config
* 参考官方issue[1455](https://github.com/sfackler/rust-openssl/issues/1455)。是由于make没有安装导致.执行`apt install make`即可