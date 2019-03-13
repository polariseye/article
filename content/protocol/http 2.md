HTTP 详解
---------------------
HTTP 是`Hypertext Transfer Protocol`的缩写，中文译为超文本传输协议。当前存在4个HTTP版本，前三个已经处于使用阶段，第四个才定调。以下主要讨论一下HTTP前三个版本。
# HTTP 1.0 概要
HTTP 1.0规定每个TCP连接只处理一次HTTP请求，处理完成后TCP连接立即断开。以通过这种方式节省服务器资源。这在纯文本或少量其他资源的情况下，没有什么问题。但现在，一个网页少则10多个关联文件，多则上百个。而且TCP建立连接需要3次握手，断开连接需要4次挥手并且建立连接过程的RTT(往返时延 Round-Trip Time)。导致建立TCP连接代价太慢。如果淘宝仍然使用HTTP1.0，估计没有人会有购物的欲望了吧。

导致HTTP 1.0如此慢的原因有如下几个因素:
1. 底层的TCP连接无法复用
2. 线头阻塞（head of line blocking），一般PC端浏览器只会对单个域名建立6到8个连接，手机端的连接数一般控制在4到6个。超过这个连接后，请求会被阻塞

HTTP 1.0存在的其他问题
1. 同服务器的一个端口只能有一个站点，由于默认只有80对外提供HTTP服务。这导致要么要么用其他端口，要么用其他服务器。不管哪种，都是比较要命的作法。
2. 文件不能断点续传


# HTTP 1.1 概要
HTTP 1.1在HTTP1.0的基础上做了很多优化：
1. 新增了更多的请求头和呼应，以对供身份认证，状态管理，Cache缓存，文件断点续传提供支持
2. 新增Host请求头，以对一个端口绑定多个站点提供支持
3. 提供对持久连接的支持（请求时，添加Keep-Alive头），提供HTTP 管线化(HTTP pipelining)技术，把多个HTTP请求放到同一个TCP连接中一一发送，发送过程中不需要等待服务器返回，但服务器端会对请求进行排除，按照先进先出(FIFO)的方式返回请求内容。这个过程中，某个请求不能临时更换TCP连接。

HTTP1.1存在的问题：
1. 线头阻塞（head of line blocking），一个TCP连接里如果前面需要传输一个较大文件，会直接阻塞后面排队的请求，也就是说线头阻塞问题并没有得到解决。
2. 在网络质量较差的情况下,管线化(HTTP pipelining)技术会让延迟更高（TCP的重传机制导致）。这导致浏览器一般没有打开这个功能。

为了减少TCP的连接建立数量，这个过程中诞生了一些有意思的技术：
1. CSS 文件内联(inline)， 直接把图片数据内嵌到css文件中
```
<imgsrc="data:image/png;base64,iVAGRw0KGDCFGNSUhEUgACBBQAVGADCAIATYJ7ljmRGGAAGElEVQQIW2P4DwcMDAxAfBvMAhEQMYgcACEHG8ELxtbPACCCTElFTEVBQmGA" />
```

  格式：data:{文件名};{数据编码方式},{编码后的数据内容}
2. Spriting 图片合并技术，把多个图片内容合并到一张大图中，然后使用css进行截图，就像这样
![](./tupianhebin.png)
3. JS文件合并 Concatenation，将多个JS文件合并为一个文件
4. 服务分片 Sharding，把资源分散在不同域名下面，从而可以让前端建立更多的连接，以提高呼应速度
# HTTP/2.0 概要


# 参考资料
* [HTTP/2 官网 https://http2.akamai.com](https://http2.akamai.com)
* [HTTP/2 github站点 https://http2.github.io](https://http2.github.io)
* [HTTP、HTTP2.0、SPDY、HTTPS 你应该知道的一些事 https://www.cnblogs.com/wujiaolong/p/5172e1f7e9924644172b64cb2c41fc58.html](https://www.cnblogs.com/wujiaolong/p/5172e1f7e9924644172b64cb2c41fc58.html)
* [What is HTTP/2 – The Ultimate Guide  https://kinsta.com/learn/what-is-http2/](https://kinsta.com/learn/what-is-http2/)
* [Http 2.0协议简介 https://www.jdon.com/dl/http2.html](https://www.jdon.com/dl/http2.html)
* [HTTP2.0 原理详解 http://www.cnblogs.com/zhuimengzhe/p/7290156.html](http://www.cnblogs.com/zhuimengzhe/p/7290156.html)
* [HTTP,HTTP2.0,SPDY,HTTPS你应该知道的一些事 http://www.alloyteam.com/2016/07/httphttp2-0spdyhttps-reading-this-is-enough/](http://www.alloyteam.com/2016/07/httphttp2-0spdyhttps-reading-this-is-enough/)
* [http2讲解 https://legacy.gitbook.com/book/ye11ow/http2-explained/details](https://legacy.gitbook.com/book/ye11ow/http2-explained/details)
* [HTTP 2.0的那些事 https://www.cnblogs.com/zlingh/p/5887143.html](https://www.cnblogs.com/zlingh/p/5887143.html)