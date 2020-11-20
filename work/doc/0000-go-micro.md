



# 介绍

[go-micro](https://github.com/micro/go-micro)最主要部分是微服务工具库[Micro](https://github.com/micro/micro) 

## Micro

micro工具库由以下几个部分组成：

- [**`api`**](https://micro.mu/docs/api.html) - API Gateway 网关。它是独立的HTTP入口，基于服务发现机制实现动态路由。
- [**`web`**](https://micro.mu/docs/web.html) - Web Dashboard web控制台。 提供可视化的发现与管理监控界面。
- [**`cli`**](https://micro.mu/docs/cli.html) - Command line interface 命令行接口。提供描述、查询终端服务的交互入口。
- [**`bot`**](https://micro.mu/docs/bot.html) - Slack与hipchat bot消息通知工具。也就是通过消息传递的CLI。
- [**`new`**](https://micro.mu/docs/new.html) - 新服务构建模板。

Micro依赖[go-micro](https://github.com/micro/go-micro)，通过它来使其变成可插拨的工具库。

## 相关资源

- [示例](https://github.com/micro/examples)，上面有相关如何使用micro的信息。
- [搜索](https://micro.mu/explore/)搜索可使用的相关micro开源项目。
- [博客](https://micro.mu/blog/)，深入理解micro，了解更多的服务设计思路。
- [视频](https://www.youtube.com/watch?v=xspaDovwk34) 2016年在英国Golong会议上，关于micro的简单介绍。
- [PPT](https://speakerdeck.com/asim)，上面有一些PPT，可供查阅

## Go Micro

Go Micro可以帮你编写微服务。

- Go Micro抽象出分布式系统
- 集成服务发布、RPC、分发/订阅机制、消息编码
- 超时容错、重试机制、负载均衡
- 功能可扩展
- 可插拔的后台交换技术

### 特点

`Go Micro` 是一个 `golang` 编写的用于构建微服务的插件化的 `RPC` 框架。它实现了服务创建、服务发现、服务间通信需要的功能需求。任何优秀的微服务架构都需要解决这三个基础问题：服务发现、同步通信和异步通信。

`Go Micro` 包括以下这些包和功能：

- Registry：客户端的服务发现
- Transport：同步通信
- Broker：异步通信
- Selector：节点过滤、负载均衡
- Codec：消息编解码
- Server：基于此库构建RPC服务端
- Client：基于此库构建RPC客户端

插件化

举个例子，默认的服务发现的机制是通过 `Consul`，但是如果想切换成 `etcd` 或者 `zookeeper` 或者任何你实现的方案，都是非常便利的。官方实现的插件可以在这个地址看到：[github.com/micro/go-plugins]

## [Go Config](https://github.com/micro/go-config)  

Go Config可以管理复杂的配置

- 动态管理 - 加载配置不需要重启
- 可插拔 - 可以选择从哪个源加载配置：文件、环境变量、consul。
- 可合并 - 针对多个配置源可以合并并重写。
- 回退 - 可以指定当key不存在时设置值。
- 可观察 - 可以查看配置的变动。

## [Go Plugins](https://github.com/micro/go-plugins)

- go-micro与micro的插件集
- 包含了绝大多数的后端技术
- grpc, kubernetes, etcd, kafka等等
- 经过生产环境验证







# [micro常见问题](https://www.jianshu.com/p/3d2dccfa055b) 



Micro由开源的库与工具组成，旨在辅助微服务开发。

- **go-micro** - 基于Go语言的可插拔RPC微服务开发框架；包含服务发现、RPC客户/服务端、广播/订阅机制等等。
- **go-plugins** - go-micro的插件有etcd、kubernetes、nats、rabbitmq、grpc等等。
- **micro** - 微服务工具集包含传统的入口点（entry point）；API 网关、CLI、Slack Bot、代理及Web UI。

还有其它相关的库和服务可以参考 [github.com/micro](https://github.com/micro)。



# [微服务工具箱](https://www.jianshu.com/p/25ee7550c15c) 



## sidecar

`Sidecar` 是一个轻量级的组装服务，概念上来说就是将 `Micro` 的库提供的功能，依附于其他语言的主程序中。`Sidecar` 本质上是一个单独运行的服务，通过 `http` 提供接口，其他语言通过接口使用 `Go Micro` 提供的功能。



`Sidecar` 的特性：

- 在服务发现系统进行注册
- 发现其他服务
- 与主程序进行健康检查
- 作为代理与 RPC 系统通信
- 通过 websocket 订阅











# 简单micro service实例

https://github.com/micro/services/tree/master/helloworld 

```go
package main
import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/services/helloworld/handler"
)
func main() {
	// Create service
	srv := service.New(
		service.Name("helloworld"),
	)
	// Register Handler
	srv.Handle(new(handler.Helloworld))

	// Run the service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
```



```shell
## call the service
micro call helloworld Helloworld.Call '{"name": "Alice"}'
```





# 问题





warning/srv/handler/warning.go 里面Get参数异常

micro是一个独立的程序，并没有服务与服务之间通信。

整个结构是怎样的，user和warning warning中srv和api什么关系，





token怎么产生

用户登录随机产生，产生流程

        "user_id": 3,
        "token": "204b0d2b9c6cbc1ae00071f7ac4b2fb5"




# 参考

[micro-china](https://microhq.cn/index-cn) 

[micro-china文档](https://microhq.cn/docs/micro/introduce-cn) 

[micro-china github](https://github.com/micro-in-cn/tutorials) 

[*Micro*/*Go*-*Micro* 中文示例、教程、资料，源码解读](https://github.com/micro-in-cn/tutorials) 

[简书go-micro介绍](https://www.jianshu.com/p/751cd31302e7) 

[go-micro github](https://github.com/micro/go-micro) 





