GTK学习笔记
---------------------------------------
[GTK](https://www.gtk.org/)是使用C主语编写的跨平台小部件工具包，它可以轻松地被各种编程语言使用。当前最新版本是GTK4。

**WINDOWS下安装GTK**
1. 设置工具链使用msvc:`rustup default stable-msvc
`
2. 通过visual studio安装 C++桌面开发套件
3. 安装git与python
4. 安装python工具包:`pip install meson ninja`。这两个工具包用于源码编译gtk
5. 下载[pkg-config-lite](https://sourceforge.net/projects/pkgconfiglite/).并解压到:`c:\pkg-config-lite`
6. 配置环境变量Path:`C:\pkg-config-lite\bin
`与`C:\gnome\bin`后面这个目录暂时不存在，后面用于保存源码编译输出
7. 编译并安装gtk4。搜索并打开命令行窗口:`x64 Native Tools Command Prompt`,然后执行以下命名:
```
git clone https://gitlab.gnome.org/GNOME/gtk.git --depth 1
cd gtk
meson setup builddir --prefix=C:/gnome -Dbuild-tests=false -Dmedia-gstreamer=disabled
meson install -C builddir
```
8. 设置环境变量:`PKG_CONFIG_PATH:C:\gnome\lib\pkgconfig
`

**使用**
1. 创建项目:`cargo new gtk_test1`
2. 在cargo.toml中添加依赖: gtk = { version = "0.4.7", package = "gtk4" }



# 参考资料

* rust的GTK教程: https://gtk-rs.org/gtk4-rs/stable/latest/book/introduction.html
* GTK官方RUST教程: https://www.gtk.org/docs/language-bindings/rust/
* GTk4-rs源码: https://github.com/gtk-rs/gtk4-rs/
* GTK环境安装教程: https://zhuanlan.zhihu.com/p/429344225


meson setup builddir --prefix=D:/ProgramFiles/gnome -Dbuild-tests=false -Dmedia-gstreamer=disabled
