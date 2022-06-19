

go-kit,go-micro,kite,gizmo,go-zero





# linux升级golang



go-micro

插件化RPC微服务框架，该框架提供了服务发现、负载均衡、同步传输、异步通信以及事件驱动等机制，尝试简化分布式系统间的通信，专注于自身业务逻辑的开发。

![]()



# [微服务](http://www.topgoer.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/GoMicro%E5%85%A5%E9%97%A8.html) 



## 概述

Micro是一个[框架](https://github.com/micro/go-micro)

Micro是一个[工具包](https://github.com/micro/micro)

Micro是一个[社区](http://slack.micro.mu/)

Micro是一个[生态系统](https://micro.mu/explore/) 

Micro由开放源码库和工具组成，以帮助微服务开发。

- **go-micro** - 用于编写微服务的可插入Go RPC框架; 服务发现，客户端/服务器rpc，pub/sub等。
- **go-plugins** - go-micro的插件，包括etcd，kubernetes，nats，rabbitmq，grpc等
- **micro** - 一个包含传统入口点的微服务工具包; API网关，CLI，Slack Bot，Sidecar和Web UI。



提供的主要软件是[Micro](https://github.com/micro/micro)，一个微服务工具包。

该工具包由以下组件组成：

- **Go Micro** - 用于在Go中编写微服务的插件式RPC框架。它提供了用于服务发现，客户端负载平衡，编码，同步和异步通信库。
- **API** - 提供并将HTTP请求路由到相应微服务的API网关。它充当单个入口点，可以用作反向代理或将HTTP请求转换为RPC。
- **Sidecar** - 一种对语言透明的RPC代理，具有go-micro作为HTTP端点的所有功能。虽然Go是构建微服务的伟大语言，但您也可能希望使用其他语言，因此Sidecar提供了一种将其他应用程序集成到Micro世界的方法。
- **Web** - 用于Micro Web应用程序的仪表板和反向代理。我们认为应该基于微服务建立web应用，因此被视为微服务领域的一等公民。它的行为非常像API反向代理，但也包括对web sockets的支持。
- **CLI** - 一个直接的命令行界面来与你的微服务进行交互。它还使您可以利用Sidecar作为代理，您可能不想直接连接到服务注册表。
- **Bot** - Hubot风格的bot，位于您的微服务平台中，可以通过Slack，HipChat，XMPP等进行交互。它通过消息传递提供CLI的功能。可以添加其他命令来自动执行常见的操作任务。

*注意：Go-micro是一个独立的库，可以独立于其他工具包使用*。



该工具包是可插入式并运行时不感知。在笔记本电脑基于docker，使用kubernetes上运行micro或者AWS等等。

![](I:\md_pictures\micro-tools.png)



https://lixiangyun.gitbook.io/go-micro/overview/introduction



## 什么是微服务

来自行业专家的微服务的定义。

1. 松散耦合的面向服务的体系结构 Adrian Cockcroft
2. 一种将单个应用程序开发为一套小型服务的方法，每个小型服务都运行在自己的进程中，并与轻量级机制进行通信 Martin Fowler

采用了一种更加全面的方式与unix进程和管道对齐的方法。

Netfix，Twitter，Gilt和Hailo都是最好的例子。所有最终都建立了自己的微服务平台。

因为微服务是一种架构模式，所以Micro通过工具在逻辑上拆分。

查看体系结构上的[博客文章](https://micro.mu/blog/2016/04/18/micro-architecture.html)，获取详细的概述。

这部分应该很多详解，解释Micro是如何构建的，以及各种lib/仓库之间是如何相互关联的。



## go-micro

Go-micro是微服务的独立RPC框架。它是该工具包的核心，并受到上述所有组件的影响。在这里，我们将看看go-micro的每个特征。

### Registry

注册表提供可插入的服务发现库，来查找正在运行的服务。当前的实现是consul，etcd，内存和kubernetes。如果您的喜欢不一样，该接口很容易实现。

### Selector

选择器通过选择提供负载均衡机制。当客户端向服务器发出请求时，它将首先查询服务的注册表。这通常会返回一个表示服务的正在运行的节点列表。选择器将选择这些节点中的一个用于查询。多次调用选择器将允许使用平衡算法。目前的方法是循环法，随机哈希和黑名单。

### Broker

Broker是发布和订阅的可插入接口。微服务是一个事件驱动的架构，发布和订阅事件应该是一流的公民。目前的实现包括nats，rabbitmq和http（用于开发）

### Transport

传输是通过点对点传输消息的可插拔接口。目前的实现是http，rabbitmq和nats。通过提供这种抽象，运输可以无缝地换出。

### Client

客户端提供了一种制作RPC查询的方法。它结合了注册表，选择器，代理和传输。它还提供重试，超时，使用上下文等

### Server

服务器是构建正在运行的微服务的接口。它提供了一种提供RPC请求的方法。

### Plugins

提供go-micro的[micro/go-plugins](https://github.com/micro/go-plugins)插件。































