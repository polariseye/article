NEAR学习笔记
----------------------------------

# 基本信息

官方地址:[https://near.org/](https://near.org/)<br>
区块浏览器:[http://explorer.testnet.near.org/](http://explorer.testnet.near.org/)<br>
社区地址:[https://near.org/community/](https://near.org/community/)<br>
文档地址:[https://docs.near.org/docs/develop/basics/getting-started](https://docs.near.org/docs/develop/basics/getting-started)<br>
测试网络地址:[https://wallet.testnet.near.org](https://wallet.testnet.near.org)<br>
主网络地址:[https://wallet.near.org](https://wallet.near.org)

nodejs开发包:[https://docs.near.org/docs/api/javascript-library](https://docs.near.org/docs/api/javascript-library)<br>
接口文档:[https://docs.near.org/docs/api/rpc](https://docs.near.org/docs/api/rpc)

**概要**<br>
* 是一个对开发人员友好的,第一层，分片、股权证明的公共区块链
* 采用分片设计的方式实现
* 使用[Project Aurora](http://www.aurora.dev/)来兼容太坊EVM协议
* 可与Octopus进行互操作。Octopus 是一个应用链网络，可与 NEAR 互操作，包括比特币、以太坊、Polkadot 和 Cosmos。
* 支持智能合约开发与DAPP应用开发
* 区块产生快，约1.2秒产生一个区块
* 交易手续费低
* 可以使用[Rainbow Bridge](https://rainbowbridge.app/)与以太坊互操作.即与以太坊进行资产转换
* 使用人类可读的账户Id而不使用公钥，也可以使用公钥当作账户地址
* 区块链的数据只在链上存储5个epoch（大约2.5天）。2.5天后数据被归档到单独的存储中。一个epoch大约12小时左右

**获取测试币方法**

在测试网注册账号即可获取到测试币：[https://wallet.testnet.near.org/](https://wallet.testnet.near.org/)<br>
我的测试号ID:qqnihao.testnet。<br>
测试号助记词:ski suspect speed daughter will produce spell bus ozone penalty fatigue enact<br>
测试网RPC地址:[https://rpc.testnet.near.org](https://rpc.testnet.near.org)

**主网信息**

主网钱包：[https://wallet.near.org/](https://wallet.near.org/)<br>
主网测试助记词:curious position laugh foster resist praise artwork grace february curve debris year<br>
主网RPC地址:[https://rpc.mainnet.near.org](https://rpc.mainnet.near.org)

## 账户
near可以使用人类可读的帐户 ID 而不是公钥哈希。需要使用钱进行账号创建

账户Id长度范围：[2,64]<br>
账户Id验证规则：`^(([a-z\d]+[\-_])*[a-z\d]+\.)*([a-z\d]+[\-_])*[a-z\d]+$`


每个账户都可以且仅可以创建一个合约

**[隐式账户](https://docs.near.org/docs/roles/integrator/implicit-accounts)**<br>
资料参考地址:[https://docs.near.org/docs/roles/integrator/implicit-accounts](https://docs.near.org/docs/roles/integrator/implicit-accounts)<br>
也可以使用隐式账户，隐式账户使用**ED25519**签名方法生成账户。公钥信息存储使用base58进行编码。但是账户Id使用公钥的16进制编码作为账户地址
* 它们允许您在创建帐户 ID 之前通过在本地生成 ED25519 密钥对来保留帐户 ID。
* 此密钥对具有映射到帐户 ID 的公钥。
* 帐户 ID 是公钥的小写十六进制表示。
* ED25519 公钥包含 32 个字节，映射到 64 个字符的帐户 ID。
* 一旦在链上创建相应的密钥，您就可以代表该帐户签署交易

## 交易

**交易费用**<br>
* 用户在调用应用程序时会立即收取gas费用 。而还是开发人员承担所有服务的成本。也可以让用户支付gas费用
* 交易费用不直接以NEAR代币计算，而是先计算得到一个固定值的gas计量单位进行计算。同一笔交易在任何时候，计算得到的gas费用是固定的
* gas单价会随着网络情况进行动态调整
* 交易的gas费用是确定性的，不能通过支付额外费用来提前计算
* 交易的具体费用很难预测。多余的手续费会返回到用户账户中
* 程序可以通过获取最近一个区块的gas价格来确定当前大概gas价格


## 其他
NEAR的彩虹桥使用了基于[aurora](https://aurora.dev/)实现的与ETH的互操作