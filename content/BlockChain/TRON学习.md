TRON学习
--------------------
TRON波场公链是基于波场协议实现的去中心化的区块链公有网络，是波场生态中最核心的部分。

官网: https://tron.network/<br>
开发者文档: https://developers.tron.network/

#接入参考资料
* API参考文档: https://cn.developers.tron.network/reference/full-node-api-overview
* TRON接入点信息: https://cn.developers.tron.network/docs/networks
* TRON其他文档: https://andelf.gitbook.io/tron/

**主网**

* 浏览器： https://tronscan.org
* TronGrid API： https://api.trongrid.io


**Shasta测试网**

Shasta测试网各个参数与主网保持一致，目前Shasta测试网不支持运行一个节点加入其中。

* 官网: https://www.trongrid.io/shasta
* 水龙头: https://www.trongrid.io/faucet
* 浏览器: https://shasta.tronscan.org
* HTTP API: https://api.shasta.trongrid.io
* grpc fullnode API: grpc.shasta.trongrid.io:50051
* grpc solidity API: grpc.shasta.trongrid.io:50061
* json-rpc API: https://api.shasta.trongrid.io/jsonrpc


**Nile测试网**

Nile测试网用于测试TRON新特性，代码版本一般会领先于主网。

* 官网：http://nileex.io
* 水龙头: http://nileex.io/join/getJoinPage
* 浏览器: https://nile.tronscan.org
* 状态: http://nileex.io/status/getStatusPage
* http API: https://api.nileex.io/
* Trongrid http AP: https://nile.trongrid.io/
* grpc fullnode API: grpc.nile.trongrid.io:50051
* grpc solidity API: grpc.nile.trongrid.io:50061
* json-rpc API: https://nile.trongrid.io/jsonrpc
* 数据库状态备份：http://47.90.243.177

**Tronex测试网**

Tronnex主要用于sun-network测试。

* 官网: http://testnet.tronex.io
* 水龙头: http://testnet.tronex.io/join/getJoinPage
* 浏览器: http://3.14.14.175:9000
* 状态: http://testnet.tronex.io/status/getStatusPage
* Full Node API: https://testhttpapi.tronex.io
* Event API: https://testapi.tronex.io
* Public Fullnode：
	* 47.252.87.28
	* 47.252.85.13
* 数据库状态备份: http://47.252.81.247

## 代币标准

**TRX**<br>
TRX是TRON网络上最主要的加密货币，应用场景广泛，TRON网络上的奖励以TRX的形式发放，用户只能通过质押TRX来获得资源以及投票权

* 燃烧TRX: TRON网络上的每笔交易都需要消耗带宽或者能量。当账户中带宽或能量不足时，就需要通过燃烧TRX来支付交易所需的资源。TRX的燃烧不但可以有助于降低TRX的通胀，而且还可以防止意外或者恶意的交易占用TRON网络资源。燃烧的TRX将被永久减除

**TRC-10**<br>
TRC-10是TRON网络支持的一种代币标准, 没有使用TRON虚拟机(TVM)，是基于链自身的一种代币标准。 在TRON网络中，每个帐户都能够发行TRC-10代币，需要支付1024TRX的发行费用

**TRC020**<BR>
TRC-20是为发行通证资产而制定的一套合约标准，即遵守这一标准编写的合约都被认为是一个TRC-20合约。当各类钱包、交易所在对接TRC-20合约的资产时，从这套合约标准中就可以知道这个合约定义了哪些函数、事件，从而方便的进行对接。