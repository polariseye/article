使用typescript开发nodejs
-----------------------

# 环境准备
前提说明：
  已经认为本机电脑已经安装了node开发环境和npm包管理程序

安装步骤：
1. 安装typescript环境
`` npm install -g typescript``

2. 添加node的提示和支持。（经过对比，发现typings的代码版本非常老，所以采用此方式）
``npm install --save @types/node``
``npm install --save @types/electron``

备注：
 * 如果下载很慢，可以使用淘宝的镜像:`--registry=https://registry.npm.taobao.org`
 * 使用命令`tsc`查看是否可用，如果不可用则把目录`%NODEJSPATH%/nodejs/node_global`加入到环境变量

# 项目初始化

1. 初始化npm环境
``npm init`` 
执行此命令后，就会在项目目录下面创建一个名为package.json文件，用于配置包相关信息

2. 初始化typescript配置
``tsc --init`` 
执行此命令后，就会在项目目录下面创建一个名为tsconfig.json的文件，用于配置typescript相关配置,配置模版可以参考下面内容：
````
{
  // 编译器的配置
  "compilerOptions": {
    // 指定生成哪个模块系统代码
    "module": "commonjs",
    // 目标代码类型
    "target": "es5",
    // emitDecoratorMetadata 和 experimentalDecorators 是与装饰器相关的
    // 在编译的时候保留装饰器里面的原数据
    "emitDecoratorMetadata": true,
    "experimentalDecorators": true,
    // 在表达式和声明上有隐含的'any'类型时报错。
    "noImplicitAny": false,
    // 用于debug
    "sourceMap": false,
    // 仅用来控制输出的目录结构--outDir。
    "rootDir": "./src",
    // 编译完后要放在哪个文件夹里面
    "outDir": "./build",
    // 在监视模式下运行编译器。会监视输出文件，在它们改变时重新编译。
    "watch": true,
    // 开发的时候要使用 es6 的语法
    "lib": ["es6"]
  },
  "include": [
    "./src"
  ],
  // 排除编译的时候哪些个文件要排除掉
  "exclude": [
    "node_modules"
    "views",
    "static"
  ]
}
````
# 编码

````
// index.ts
let uName : string ="nihao"
console.log("hi--",uName)
````
# 项目编译
1. typescript代码转换为js代码：`tsc index.ts`
2. 执行：`node index.js`，考虑到方便，可以反此行配置在`package.json`文件中,然后使用命令:'npm start'执行：

```
{
  "name": "jsonlog",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "start": "tsc main.ts&&electron index.js"
  },
  "author": "",
  "license": "ISC",
  "devDependencies": {
    "electron": "5.0.8"
  },
  "dependencies": {
    "@types/jquery": "^3.3.30",
    "jquery": "^3.4.1",
    "jsdom": "^15.1.1"
  }
}

```
