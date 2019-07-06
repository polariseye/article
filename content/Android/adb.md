ADB使用指南
-------------------------------

Android 调试桥 (adb) 是一个通用命令行工具，其允许您与模拟器实例或连接的 Android 设备进行通信。说到 ADB 大家应该都不陌生，即 Android Debug Bridge，Android调试桥，身为 Android 开发的我们，熟练使用 ADB 命令将会大大提升我们的开发效率， ADB 的命令有很多，今天就来总结下我在开发常用到的一些 ADB 命令。

* Android模拟设备启动完成后查看连接到本地计算机上的Android设备列表:  adb devices (重点)

* 查看Android 版本  ：adb version

* 启动 adb server ：adb start-server

* 停止 adb server ：adb kill-server

* 列出手机装的所有app的包名：adb shell pm list packages

* 列出系统应用的所有包名：adb shell pm list packages -s

* 列出除了系统应用的第三方应用包名：adb shell pm list packages -3

* 安装一个apk，执行以下命令：adb install  直接卸载：adb uninstall

* 查看ADB帮助：adb help

* 参考文章：常用ADB命令