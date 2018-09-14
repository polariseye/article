js中的CommonJS，AMD，CMD，UMD
------------------------------
js按照模块加载方式产生了3种规范。分别为CommonJS，AMD，CMD，UMD.

# CommonJS规范
官方地址：http://www.commonjs.org/
规范分为：模块定义，模块引用，模块标识
## 模块定义
使用定义exports导出函数,module.exports导出模块

````
// 文件名:hello.js
exports.helloWorld=function (name){
    alert("helloWorld "+name);
}

function SayModule () {

    this.hello = function () {

        console.log('hello');

    };

 

    this.goodbye = function () {

        console.log('goodbye');

    };

} 

module.exports = SayModule;
````

## 模块引用
使用require函数导入一个模块或文件，然后使用返回值调用模块函数，例如：
`` 
// 导出函数的使用
var hello=require("./hello.js");
hello.helloWorld("haha");

// 导出模块的使用
var obj=new hello();
obj.goodbye();
``
## 模块标识
模块标识指的是传递给require方法的参数，文件相对路径作为参数即可

## 特点
1. 模块定义简洁
2. 通过规范引入导出导入，可以流畅地连接各个模块

common.js本身也实现了基于这个的模块加载器
现在支持es6的各种js引擎都支持这种方式了。当然典型的是node.js

## 缺点
这个规范是最开始出现的，为解决js进行复杂功能开发而设计的。明显有点类似C语言代码文件的导入规则。
1. 更加偏向于服务端
2. 所有模块同步加载，并在第一次使用时加载。如果时放在浏览器这就有点要命了，因为浏览器还要去远端下载文件

# AMD规范
AMD 是 Asynchronous Module Definition 的简称，即“异步模块定义”，是从 CommonJS 讨论中诞生的。AMD 优先照顾浏览器的模块加载场景，使用了异步加载和回调的方式。

使用define函数定义模块实现和说明需要引用的模块
````
// file lib/sayModule.js
define(function (){
    return {
        sayHello: function () {
            console.log('hello');
        }
    };
});

 

//file main.js
define(['./lib/sayModule'], function (say){
    say.sayHello(); //hello
})
````


## 特点
1. 允许非同步加载模块，也可以按需动态加载模块
2. 可以简单方便地移植基于commonjs的代码
基于此实现的模块加载库是：RequireJS

# CMD规范
CMD（Common Module Definition），在CMD中，一个模块就是一个文件。这个是国内淘宝的玉伯大牛搞出来的东西。
全局函数define，用来定义模块。 
参数 factory 可以是一个函数，也可以为对象或者字符串。 
当 factory 为对象、字符串时，表示模块的接口就是该对象、字符串。


基于此实现的模块加载库是：SeaJS

同样Seajs也是预加载依赖js跟AMD的规范在预加载这一点上是相同的，明显不同的地方是调用，和声明依赖的地方。AMD和CMD都是用difine和require，但是CMD标准倾向于在使用过程中提出依赖，就是不管代码写到哪突然发现需要依赖另一个模块，那就在当前代码用require引入就可以了，规范会帮你搞定预加载，你随便写就可以了。但是AMD标准让你必须提前在头部依赖参数部分写好（没有写好？ 倒回去写好咯）。这就是最明显的区别。

# UMD规范
统一模块定义（UMD：Universal Module Definition ）就是将 AMD 和 CommonJS 合在一起的一种尝试，常见的做法是将CommonJS 语法包裹在兼容 AMD 的代码中。

# AMD和UMD的区别
1. 对于依赖的模块，AMD 是提前执行，CMD 是延迟执行。不过 RequireJS 从 2.0 开始，也改成可以延迟执行（根据写法不同，处理方式不同）。CMD 推崇 as lazy as possible.
2. CMD 推崇依赖就近，AMD 推崇依赖前置。
 
3. AMD 的 API 默认是一个当多个用，CMD 的 API 严格区分，推崇职责单一。比如 AMD 里，require 分全局 require 和局部 require，都叫 require。CMD 里，没有全局 require，而是根据模块系统的完备性，提供 seajs.use 来实现模块系统的加载启动。CMD 里，每个 API 都简单纯粹。

4. 还有一些细节差异，具体看这个规范的定义就好，就不多说了。

# 参考资料
* [该如何理解AMD ，CMD，CommonJS规范--javascript模块化加载学习总结 ](https://www.cnblogs.com/qianshui/p/5216580.html)
* [很全很全的JavaScript的模块讲解(CommonJS,AMD,CMD,ES6模块)](https://blog.csdn.net/arsaycode/article/details/78959780)
* [commonjs官网](http://www.commonjs.org/)
* [简要理解CommonJS规范](https://blog.csdn.net/u012443286/article/details/78825917)