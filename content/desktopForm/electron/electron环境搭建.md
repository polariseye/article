electron 开发环境搭建（windows篇）
----------------------------------------
electron 是一个使用html+css+javascript开发桌面应用的框架。javascript采用nodejs，目的是使js能够访问本地资源。ui呈现采用Chrome浏览器进行呈现，考虑大小问题，有对此浏览器进行裁剪。

electron官网地址:[传送门](https://electronjs.org)
electron源码地址:[传送门](https://github.com/electron/electron)

# 安装步骤
官方安装教程:[传送门](https://electronjs.org/docs/tutorial/first-app)

### NodeJS安装
安装包下载地址:[传送门](http://nodejs.cn/download/)
由于墙的原因，建议把npm源切换到国内，比如淘宝的
```
npm config set registry "https://registry.npm.taobao.org/" 
```

## Electron安装
直接使用npm命令安装即可：
```
npm install -g electron
``` 
也可以在安装时，指定源
```
npm install -g cnpm --registry=https://registry.npm.taobao.org
```
其中,```-g```是指把electron安装到全局中,

## 打包工具安装
处于方便考虑，可以再安装一个打包工具[electron-packager](https://github.com/electron-userland/electron-packager)。使用npm方式安装如下：
```
npm install -g electron-packager
```
