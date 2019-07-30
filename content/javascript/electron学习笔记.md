electron学习笔记
--------------------------------

# 打包
本次打包使用了打包工具:`electron-packager`,具体流程如下

1. 安装打包工具:
````
npm install -g electron-packager
````
2. 配置打包命令：修改文件`package.json`添加打包命令，修改后整个文件内容如下图所示:
````
{
  "name": "Helloworld",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "start": "electron index.js",
    "pack": "electron-packager . jsonlog --win --out ../jsonlogRelease --icon=./jsonlog.ico --arch=x64 --app-version=0.0.1 --electron-version=5.0.8 --overwrite"
  },
  "author": "",
  "license": "ISC",
  "devDependencies": {
    "electron": "5.0.8"
  },
  "dependencies": {
    "@types/jquery": "^3.3.30",
    "jquery": "^3.4.1",
    "jsdom": "^15.1.1",
    "layui": "0.0.1",
    "layui-src": "^2.5.4"
  }
}
````
参数说明:
 * “.”：需要打包的应用目录（即当前目录
 * “Helloworld”：应用名称，
 * “--win”：打包平台（以Windows为例），
 * “--out ../myClient”：输出目录，
 * “--arch=64”：64位，
 * “--icon=./HelloWorld.ico”：指定可执行程序的ico图片，
 * “--app-version=0.0.1”：应用版本，
 * “--electron-version=5.0.8”：electron版本
 * “--overwrite”：删除输出目录后，重写输出内容
3. 执行打包
```
npm run pack
```
打包完成后，就会在对应的输出目录生成对应程序

# 参考资料

* [electron安装+运行+打包成桌面应用+打包成安装文件+开机自启动](https://www.cnblogs.com/kakayang/p/9559777.html)