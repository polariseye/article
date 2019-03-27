actor模型开篇--ET框架详解
-------------------------

在介绍actor之前，大家先回忆下，windows的消息通知机制。Windows为每个应用程序都有分配一个消息队列，应用程序不断从队列中读取消息，并进行处理。其中涉及到的两个主要函数为`SendMessage`和`PostMessage`——用于给一个窗口发送消息。整个流程而言，就可以称之为actor模型。整体流程如下图所示：
![](windowsxiaoxiduilie.jpg)

说明

后续单独的actor单词指一个具体的参与对象，相当于接受和处理消息的windows窗体对象

actor模型由三部分组成
1. **状态(state)** : 每个内部的状态信息，比如，窗体标题，窗口尺寸

2. **行为(behavior)**：具体业务处理逻辑，比如改变窗口标题这个动作

3. **邮箱(mailbox)**: 用于缓存和调度所有请求，可以理解为，调用PostMessage后，消息暂存的处理方缓存（不是很恰当，但是很形象）

以actor1查询actor2的年龄为例，整体如下图所示（从左往右看）：

![](actorliuchengtu.png)

流程细说：

1. actor1获取到actor2的唯一标记，并发出`QueryAge`的请求到actor系统内部
2. actor系统把actor1的请求推送到actor所拥有的mailbox中
3. mailbox根据actor2的情况，把`QueryAge`请求发送给actor2对象处理
4. actor2把处理结果推送给actor系统
5. actor系统把结果推送给actor1的mailbox
6. mailbox根据actor1的情况，把`QueryAge`请求发送给actor1对象处理

说明

* 整个流程中，只有actor相关逻辑由我们自己实现，其他由actor系统内部实现
* 消息从mailbox到actor时，可以根据实际情况，可以mailbox推送给actor，也可以actor主动去拉取
* actor2的地址一般是由actor系统通过actor2的唯一标识推断出来，也就是说，actor1实际是需要提前知道actor的唯一标识的
* actor2和actor1实际可能在同一个进程，也可能是在不同进程



akka与akka.net
# 参考资料
* [Actor模型](https://github.com/egametang/ET/blob/master/Book/5.4Actor模型.md)
* [Actor模型原理 ](https://www.cnblogs.com/MOBIN/p/7236893.html)
* [当多线程并发遇到Actor](https://mp.weixin.qq.com/s/mzZatZ10Rh19IEgQvbhGUg)
* [Windows消息队列](https://www.cnblogs.com/lidabo/p/3695265.html)