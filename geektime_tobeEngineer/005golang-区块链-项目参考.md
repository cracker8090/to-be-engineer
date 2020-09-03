[toc]

# 网站


https://0xzx.com/



[golang构建一个简单区块链](http://blog.hubwiz.com/2019/01/17/blockchain-build-go/)  [源码](https://github.com/codehakase/blockchain-golang-tutorial/blob/master/blockchain.go) [区块链个人编写](http://blog.hubwiz.com) 



- [EOS教程](http://xc.hubwiz.com/course/5b52c0a2c02e6b6a59171ded?affid=0117blog)，本课程帮助你快速入门EOS区块链去中心化应用的开发，内容涵盖EOS工具链、账户与钱包、发行代币、智能合约开发与部署、使用代码与智能合约交互等核心知识点，最后综合运用各知识点完成一个便签DApp的开发。
- [java以太坊开发教程](http://xc.hubwiz.com/course/5b2b6e82c02e6b6a59171de2?affid=0117blog)，主要是针对java和android程序员进行区块链以太坊开发的web3j详解。
- [python以太坊](http://xc.hubwiz.com/course/5b40462cc02e6b6a59171de4?affid=0117blog)，主要是针对python工程师使用web3.py进行区块链以太坊开发的详解。
- [php以太坊](http://xc.hubwiz.com/course/5b36629bc02e6b6a59171de3?affid=0117blog)，主要是介绍使用php进行智能合约开发交互，进行账号创建、交易、转账、代币开发以及过滤器和交易等内容。
- [以太坊入门教程](http://xc.hubwiz.com/course/5a952991adb3847553d205d1?affid=0117blog)，主要介绍智能合约与dapp应用开发，适合入门。
- [以太坊开发进阶教程](http://xc.hubwiz.com/course/5abbb7acc02e6b6a59171dd6?affid=0117blog)，主要是介绍使用node.js、mongodb、区块链、ipfs实现去中心化电商DApp实战，适合进阶。
- [C#以太坊](http://xc.hubwiz.com/course/5b6048c3c02e6b6a59171dee?affid=0117blog)，主要讲解如何使用C#开发基于.Net的以太坊应用，包括账户管理、状态与交易、智能合约开发与交互、过滤器和交易等。
- [java比特币开发教程](http://xc.hubwiz.com/course/5bb35c90c02e6b6a59171df0?affid=0117blog)，本课程面向初学者，内容即涵盖比特币的核心概念，例如区块链存储、去中心化共识机制、密钥与脚本、交易与UTXO等，同时也详细讲解如何在Java代码中集成比特币支持功能，例如创建地址、管理钱包、构造裸交易等，是Java工程师不可多得的比特币开发学习课程。
- [php比特币开发教程](http://xc.hubwiz.com/course/5b9e779ac02e6b6a59171def?affid=0117blog)，本课程面向初学者，内容即涵盖比特币的核心概念，例如区块链存储、去中心化共识机制、密钥与脚本、交易与UTXO等，同时也详细讲解如何在Php代码中集成比特币支持功能，例如创建地址、管理钱包、构造裸交易等，是Php工程师不可多得的比特币开发学习课程。
- [tendermint区块链开发详解](http://xc.hubwiz.com/course/5bdec63ac02e6b6a59171df3?affid=0117blog)，本课程适合希望使用tendermint进行区块链开发的工程师，课程内容即包括tendermint应用开发模型中的核心概念，例如ABCI接口、默克尔树、多版本状态库等，也包括代币发行等丰富的实操代码，是go语言工程师快速入门区块链开发的最佳选择。



# bitcoin

借鉴了来自数字货币、密码学、博弈论、分布式系统、控制论等多个领域的 技术成果。本章将介绍比特币项目的来源、核心原理设计、相关的工具，以及关键的技术话题。    

比特币网络在功能上具有如下特点： 

- 去中心化：意味着没有任何独立个体可对网络中交易进行破坏，任何交易请求都需要大多数参与者的共识； 

- 匿名性：比特币网络中账户地址是匿名的，无法从交易信息关联到具体个体，但这也意味着很难进行审计；

- 通胀预防：比特币的发行需要通过挖矿计算来进行，发行量每四年减半,总量上限为 2100 万枚，无法被超发。    



pgp  gunpg,rsa区别？metzdowd的加密邮件列表。



# 以太坊



# Hyperledger

[gerrit路径](https://gerrit.hyperledger.org),[github路径](https://github.com/hyperledger/)

目前主要包括三大账本平台项目和若干其它项目。 

账本平台项目： 

Fabric：包括 Fabric、Fabric CA、Fabric SDK（包括 Node.Js、Python 和 Java 等语 言） 和 fabric-api、fabric-sdk-node、fabric-sdk-py 等，目标是区块链的基础核心平台， 支持 pbft 等新的 consensus 机制，支持权限管理，最早由 IBM 和 DAH 发起； 

SawToothLake：包括 arcade、core、dev-tools、validator、mktplace 等。是 Intel 主要 发起和贡献的区块链平台，支持全新的基于硬件芯片的共识机制 Proof of Elapsed Time（PoET） 。 

Iroha：账本平台项目，基于 C++ 实现，带有不少面向 Web 和 Mobile 的特性，主要由 Soramitsu 发起和贡献。    

其它项目： 

Blockchain Explorer：提供 Web 操作界面，通过界面快速查看查询绑定区块链的状态 （区块个数、交易历史） 信息等。 

Cello：提供"Blockchain as a Service" 功能，使用 Cello，管理员可以轻松获取和管理多 条区块链；应用开发者可以无需关心如何搭建和维护区块链。 目前，所有项目均处于孵化（Incubation） 状态。    

项目原则 

项目约定共同遵守的 基本原则 为： 重视模块化设计，包括交易、合同、一致性、身份、存储等技术场景； 

代码可读性，保障新功能和模块都可以很容易添加和扩展； 

演化路线，随着需求的深入和更多的应用场景，不断增加和演化新的项目。 

如果你对 Hyperledger 的源码实现感兴趣，可以参考 [Hyperledger 源码分析之 Fabric](https://github.com/yeasy/hyperledger_code_fabric)。    

https://wiki.hyperledger.org/groups/twgc

# EOS

Dan Larimer（Bitshares，Graphene和Steem / Steemit的发明人）与eos.io 团队一起宣布开发 EOS，这是一个共识区块链操作系统，提供数据库，帐户权限，日程安排，身份验证和互联网应用通信。EOS 将为开发人员提供他们需要的工具，以便他们可以专注于其应用程序的特定业务逻辑，而不用担心密码实现或与去中心化计算机的通信（即区块链）。此外，EOS 将使用并行化来增强区块链可伸缩性，以便达到每秒数百万次交易。 

https://sdk.cn/news/8101(EOS和以太坊区别，比特币，三代的区别)

鉴于以太坊使用 PoW（即将转而采用混合 PoW/PoS），EOS 将使用石墨烯技术，该技术利用 DPOS 共识机制。 EOS 的 DPOS 共识机制在硬分叉期间不会产生多个竞争链。 Steem 网络所经历的 18 次成功的硬分叉证明了这一点，背后采用的就是 Graphene 石墨烯算法。 

[区块链3.0：拥抱EOS](https://www.cnblogs.com/Evsward/p/eos-intro.html)



