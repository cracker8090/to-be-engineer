# RPC原理及应用

远程过程调用(Remote Procedure Call，简称RPC)，在微服务大行其道的今天，得到了广泛的应用 ，分布式系统服务群中开发应用 

[TOC]

## 原理介绍

RPC 可基于 HTTP 或 TCP 协议，Web Service 就是基于 HTTP 协议的 RPC，它具有良好的跨平台性，但其性能却不如基于 TCP 协议的 RPC。会两方面会直接影响 RPC 的性能，一是传输方式，二是序列化。 

在一般情况下，TCP 一定比 HTTP 快。就序列化而言，Java 提供了默认的序列化方式，但在高并发的情况下，这种方式将会带来一些性能上的瓶颈，于是市面上出现了一系列优秀的序列化框架，比如：Protobuf、Kryo、Hessian、Jackson 等，它们可以取代 Java 默认的序列化，从而提供更高效的性能。 

**RPC框架原理** 

在RPC框架中主要有三个角色：Provider、Consumer和Registry 

![](../../../md_pictures/rpc框架1.png)

- Server: 暴露服务的服务提供方。
- Client: 调用远程服务的服务消费方。
- Registry: 服务注册与发现的注册中心。

![](../../../md_pictures/rpc基本流程图.png)

一次完整的RPC调用流程（同步调用，异步另说）如下：
 1）服务消费方（client）调用以本地调用方式调用服务；
 2）client stub接收到调用后负责将方法、参数等组装成能够进行网络传输的消息体；
 3）client stub找到服务地址，并将消息发送到服务端；
 4）server stub收到消息后进行解码；
 5）server stub根据解码结果调用本地的服务；
 6）本地服务执行并将结果返回给server stub；
 7）server stub将返回结果打包成消息并发送至消费方；
 8）client stub接收到消息，并进行解码；
 9）服务消费方得到最终结果。

![rpc注册服务和发现](../../../md_pictures/rpc注册服务和发现.png)

服务提供者启动后主动向注册中心注册机器ip、port以及提供的服务列表； 

服务消费者启动时向注册中心获取服务提供方地址列表，可实现软负载均衡和Failover； 























1.动态代理——生成 client stub和server stub需要用到 **Java 动态代理技术 **，我们可以使用JDK原生的动态代理机制，可以使用一些开源字节码工具框架 如：CgLib、Javassist等。 

2.序列化技术——protobuf、Thrift、hessian、Kryo、Msgpack 

3.NIO ——基于netty这一IO通信框架，比如阿里巴巴HSF、dubbo，Hadoop Avro，推荐使用Netty 作为底层通信框架。 

4.服务注册中心——可选技术：Redis，Zookeeper，Consul，Etcd

基于 Netty4 + Zookeeper + protostuff + Spring 实现了一个简单、高效的RPC框架[Mango](https://link.jianshu.com/?t=https://github.com/TiFG/mango) 

## 优秀的开源rpc框架 

**Java中的RPC框架比较多，各有特色，广泛使用的有RMI、Hessian、Dubbo等**  

rpc协议有很多，最早的CORBA，java RMI，webservice的rpc风格，hessian，thrift，甚至rest API。

### Netty

只是网络通信框架，把java socket的api又封装了一次，最少代码完成网络通信任务。

### 分布式RPC框架性能大比拼

dubbo、motan、rpcx、gRPC、thrift的性能比较

### [阿里巴巴 Dubbo](https://link.jianshu.com?t=https://github.com/alibaba/dubbo)

[https://github.com/alibaba/dubbo](https://link.jianshu.com?t=https://github.com/alibaba/dubbo) 

可以和 Spring框架无缝集成，dubbo由于跟淘宝另一个类似的框架HSF（非开源）有竞争关系，导致dubbo团队已经解散 。后又更新。Dubbo 包含远程通讯、集群容错和自动发现三个核心部分。 

### 阿里Dubbo疯狂更新，关Spring Cloud什么事？

http://www.ityouknow.com/springcloud/2017/11/20/dubbo-update-again.html

###  [新浪微博 Motan](https://link.jianshu.com?t=https://github.com/weibocom/motan)

[https://github.com/weibocom/motan](https://link.jianshu.com?t=https://github.com/weibocom/motan) 

新浪java，Motan 在微博平台中已经广泛应用，每天为数百个服务完成近千亿次的调用。 

###  [gRPC](https://link.jianshu.com?t=https://github.com/grpc/grpc)

[https://github.com/grpc/grpc](https://link.jianshu.com?t=https://github.com/grpc/grpc) 

由Google主要面向移动应用开发并基于HTTP/2协议标准而设计，基于ProtoBuf(Protocol Buffers)序列化协议开发，且支持众多开发语言。本身它不是分布式的，所以要实现上面的框架的功能需要进一步的开发。 

不仅仅是因为它出身名门，它对移动端、多语言以及 HTTP/2 的支持也是它能脱颖而出的重要原因。 

###  [rpcx ](https://link.jianshu.com?t=https://github.com/smallnest/rpcx)

[https://github.com/smallnest/rpcx](https://link.jianshu.com?t=https://github.com/smallnest/rpcx) 

Go语言生态圈的Dubbo， 比Dubbo更轻量，实现了Dubbo的许多特性，借助于Go语言优秀的并发特性和简洁语法，可以使用较少的代码实现分布式的RPC服务。 

###  [Apache Thrift ](https://link.jianshu.com?t=https://thrift.apache.org/)

[https://thrift.apache.org/](https://link.jianshu.com?t=https://thrift.apache.org/)

### 百度brpc

[RPC框架brpc](https://github.com/brpc/brpc) 

dubbo应该是3这中最简单易用的了，但dubbo只支持Java语言，thrift和gprc都是支持跨语言的，并且dubbo内部帮你实现了对分布式注册中心zookeeper的使用，另外两个仍需自己实现对注册中心的操作。 thrift grpc等之所以支持跨语言，是因为他们自己定义了一套结构化数据存储格式，如Google的protobuf，用于编解码对象，作为各个语言通信的中间协议。 [brpc中的读和写都是wait-free的，这是最高程度的并发](https://github.com/brpc/brpc/blob/master/docs/cn/bthread.md) ，[定位问题的便利性](https://github.com/brpc/brpc/blob/master/docs/cn/builtin_service.md) 

| Dubbo            | Montan | rpcx                               | gRPC | Thrift              |                |
| ---------------- | ------ | ---------------------------------- | ---- | ------------------- | -------------- |
| 开发语言         | Java   | Java                               | Go   | 跨语言              | 跨语言         |
| 分布式(服务治理) | √      | √                                  | √    | ×                   | ×              |
| 多序列化框架支持 | √      | √  (当前支持Hessian2、Json,可扩展) | √    | ×  (只支持protobuf) | × (thrift格式) |
| 多种注册中心     | √      | √                                  | √    | ×                   | ×              |
| 管理中心         | √      | √                                  | √    | ×                   | ×              |
| 跨编程语言       | ×      | × (支持php client和C server)       | ×    | √                   | √              |

 



## 深入浅出rpc

### LPC&IPC

本地过程调用在不同的操作系统中，叫法不同，使用方式也不太一样。在Windows编程中，称为LPC；在linux编程中，更习惯称之为IPC，即进程间通信。 

消息队列：Linux提供的消息队列和各种分布式MQ不同，它是在内核中使用链表结构来保持消息的队列，然后其他进程从内核的消息队列中获取消息。 

信号：在控制台输入的`CTRL + C`来向执行的进程发送kill信号来结束该进程。对于信号，一般我们再终端交互窗口中使用比较多，在服务端开发中很少涉及。 

管道：管道命令，其实内部实现就是使用的linux管道接口，每个命令其实是一个进程，各个进程的标准输出STDOUT，作为下一个进程的标准输入STDIN。 

Linux管道包含：匿名管道和命名管道。 

（1）匿名管道：只能父子进程间通信。使用pipe()方法来创建：

```c
    #include <unistd.h>
    int pipe(int filedis[2]);
```

> 参数filedis返回两个文件描述符：filedes[0]为读而打开，filedes[1]为写而打开。filedes[1]的输出是filedes[0]的输入

（2）命名管道：可以在单台机器内的任何一组进程间进行通信。一般我们使用mkfifo()来创建命名管道：

```c
    #include <sys/types.h>
    #include <sys/stat.h>
    int mkfifo(const char * pathname,mode_t mode)
```

> 成功返回0，失败返回-1。成功返回之后，pathname其实就可以看着一个管道文件操作(当然并没有真实文件在磁盘存在)，对于文件操作的方法例如open,read,write都适用于fifo命名通道。

信号量Semaphore：信号量的主要作用就是同步，所以我们一般是使用共享内存方式完成进程间通信，而在此过程中通过信号量来完成多进程间的同步协调机制。 信号量其实就是一个比较特殊的变量，然后对它的操作都是原子进行的，并且一般只提供两种方法：P和V操作(在java中为wait()和notify())。 

- P(sv)：如果sv的值大于零，就给它减1；如果它的值为零，就挂起该进程的执行；
- V(sv)：如果有其他进程因等待sv而被挂起，就让它恢复运行，如果没有进程因等待sv而挂起，就给它加1。

linux对外提供的API接口方法如下所示:

```c
    struct sem {
      short sempid;/* pid of last operaton */
      ushort semval;/* current value */
      ushort semncnt;/* num procs awaiting increase in semval */
      ushort semzcnt;/* num procs awaiting semval = 0 */
    }
　　 #include <sys/types.h>
　　 #include <sys/ipc.h>
　　 #include <sys/sem.h>
    //首先获取一个信号量,只有该方法可以才能直接使用key，其他方法必须先semget然后才能使用信号量
　　 int semget(key_t key, int nsems, int flag);
    //对信号量进行操作，直接控制信号量信息，比如删除信号量
    int semctl(int semid, int semnum, int cmd, union semun arg);
    //改变信号量的值，P,V操作都是通过该方法
    int semop(int sem_id, struct sembuf *sem_opa, size_t num_sem_ops);
```

共享内存：

Linux对外提供了共享内存的方法来完成进程间通信 

```c
#include <sys/types.h>
　   #include <sys/ipc.h>
　　 #include <sys/shm.h>
    //创建共享内存空间，大小为size
    int shmget(key_t key, size_t size, int shmflg);
    //所有需要使用共享内存通信的进程，映射到自身的内存地址空间中
    void *shmat(int shmid, void *addr, int flag);
    //从当前进程地址空间中分离该共享内存
    int shmdt(const void *shmaddr);
    //控制共享内存的，比如删除该共享内存空间等
    int shmctl(int shm_id, int command, struct shmid_ds *buf);
```

> 从上面的方法可以很显然的看出，进程间的内存地址空间是独立隔离的(内核地址空间由于虚拟地址和物理地址是一致的，所以在进程间这块地址空间也是一致的，不过我们操作的都是用户空间的内存，所以不考虑这块)。当我们想要共享操作，必须要把物理内存分别绑定到对应进程的地址空间，才能共享操作。
>
> 使用的时候，很简单。`shmat`方法返回一个`void *`就可以强转某个指定的struct，然后直接操作该对象结构体即可。由于共享，所以需要考虑多线程同步安全问题。

Socket套接字



### Web Service技术

> `Web Service`一般有两种定义：
>
> 1. 特指 W3C组织制定的`web service`规范技术。其包括SOAP(一个基于XML的可扩展消息信封格式，需同时绑定一个网络传输协议。这个协议通常是HTTP或HTTPS，但也可能是SMTP或XMPP)、WSDL(一个XML格式文档，用以描述服务端口访问方式和使用协议的细节。通常用来辅助生成服务器和客户端代码及配置信息)和UDDI(一个用来发布和搜索WEB服务的协议，应用程序可借由此协议在设计或运行时找到目标WEB服务)。从上面三个定义就可以看出，这种规范技术是一个重量级的协议。
> 2. 泛指网络系统对外提供web服务所使用的技术。这里，我们主要是基于该定义来理解。

一般而言，**技术体系，必然是服务于架构体系的**。不同的架构，所约定的技术结构设计还是有些区别的。

因此，要了解web服务技术，必然要先了解其服务于哪个架构体系；也就是说，先去了解技术产生的架构背景。

Web Service 就是基于 HTTP 协议的 RPC，它具有良好的跨平台性，但其性能却不如基于 TCP 协议的 RPC。会两方面会直接影响 RPC 的性能，一是传输方式，二是序列化。 

Java 提供了 NIO 的解决方案，Java 7 也提供了更优秀的 NIO.2 支持，用 Java 实现 NIO 并不是遥不可及的事情，只是需要我们熟悉 NIO 的技术细节。 

#### SOA & 微服务

在分布式网络服务架构体系中，最火的莫过于 SOA(面向服务架构，Service-Oriented Architecture)和微服务。 

多机部署，通过Nginx或者其他代理/均衡软件来分发请求到相同服务的不同机器上 。**SOA**，也就是基于服务的架构设计理念。SOA的设计理念，就是把所有的服务都对外以HTTP或者其他协议方式对外暴露，绝对`不允许`相同的服务在不同的业务系统独立一套，然后共用底层数据库。服务化的设计系统，所有拆分的业务，彼此之间都通过暴露的服务接口通信，操作对方的数据。这样，各个业务系统之间开始独立自主的向着美好的方向发展了。 对于内部业务系统的架构设计，就出现了**微服务Micro-Service**了。我们把单个业务系统中一些功能细节的结构封装成服务，大的对外业务系统，组装各个微服务的接口数据，然后提供SOA服务。 

SOA其实和微服务，从我的视角来看，其实就是 业务外部和内部服务的不同架构设计而已，其技术框架很大程度上都可以通用。其区别如下图： 

![微服务SOA](../../../md_pictures/微服务SOA.png)

SOA一般使用SOAP或者REST方式来提供服务，这样外部业务系统可以使用通用网络协议来处理请求和响应，而微服务，还可以有一些私有的协议方式来提供服务，例如基于自定义协议的RPC框架。RPC使得调用服务简单，但是需要一些其他耗时间的交流协调工作，这适合微服务的场景，但是不一定适合SOA场景了。 

#### web服务技术结构

![](../../../md_pictures/web服务技术体系结构图.png)

`web service`被W3C设立规范之初，SOAP方案就被提出来。但是，随着服务化技术和架构的发展，SOAP多少有点过于复杂，因此就出现了简化版的REST方案。此后，由于分布式服务应用越来越大，对性能和易用性上面要求越来越大，因此就出现了RPC框架(很多时候，**RPC**并不被当做一种web service方案。在绝大部分博客中，介绍web service 只会讨论 **SOAP和REST**，主要是其基本上都是基于SOA来介绍服务方案)。 

##### SOAP

SOAP，全称为 Simple Object Access Protocol，也就是 简单对象访问协议。跟着`web service`一起出来的，可能会被淘汰。与SOAP相关的配套协议是WSDL (Web Service Description Language)，用来描述哪个服务器提供什么服务，怎样找到它，以及该服务使用怎样的接口规范，类似我们现在聊服务治理中的服务发现功能。 

SOAP服务整体流程是：首先，获得该服务的WSDL描述，根据WSDL构造一条格式化的SOAP请求发送给服务器，然后接收一条同样SOAP格式的应答，最后根据先前的WSDL解码数据。绝大多数情况下，请求和应答使用HTTP协议传输，那么发送请求就使用HTTP的POST方法。 

##### REST

REST，全称 REpresentational State Transfort，也就是 表示性状态转移。由于SOAP方案过于庞大复杂，在很多简单的web服务应用场景中，**轻量级的REST就出现替代SOAP方案**了。 和SOAP相比，REST只是对URI做了一些规范，数据才有**JSON格式**，底层传输使用**HTTP/HTTPS**来通信，因此，所有web服务器都可以快速支持该方案；开发人员也可以快速学习和使用。 

##### SOAP & REST

从命名来看，SOAP是一种协议，而REST只是一种方案。协议的设计很多时候，从上而下一整套都是新的，需要设计开发专门的工具支持；而方案相对就是基于目前以后的工具来做一些设计和约束，这就是为什么REST快速替换了SOAP的地位。

REST特点：

- 由于数据返回格式是自定义的，绝大部分使用JSON，这种数据结构节省带宽，并且前端JavaScript能天生支持。
- 无状态，基于HTTP协议，所以只能适应无状态场景。

SOAP特点：

- 协议有安全性的一些规范。
- 基于xml的标签约束，而且也不要去底层是HTTP传输，所以支持有状态的场景。

##### RPC家族

**RMI**是Java制定的远程通信协议。**Thrift**这种基于IDL来跨语言的RPC组件就出现了，按照Thrift官方规定的方式来写API结构，然后生成对应语言的API接口，继而就可以跨语言完成远程过程调用了。 作为服务化的组件，如果没有服务治理来完成大规模应用集群中服务调用管理工作，则运维工作则是非常繁重的，因此类似**dubbo**这种包含服务治理的RPC组件出现了。 

### RPC介绍

> RMI作为Java自带的官方RPC组件，单独介绍；然后我们来看看通用RPC实现结构。 

#### RMI介绍

RMI，全称是Remote Method Invocation，也就是远程方法调用。在JDK 1.2的时候，引入到Java体系的。当应用比较小，性能要求不高的情况下，使用RMI还是挺方便快捷的。 

RMI的调用流程 

![](../../../md_pictures/RMI调用流程.png)

> stub(桩)：stub实际上就是远程过程在客户端上面的一个代理proxy。当我们的客户端代码调用API接口提供的方法的时候，RMI生成的stub代码块会将请求数据序列化，交给远程服务端处理，然后将结果反序列化之后返回给客户端的代码。这些处理过程，对于客户端来说，基本是透明无感知的。
>
> remote：这层就是底层网络处理了，RMI对用户来说，屏蔽了这层细节。stub通过remote来和远程服务端进行通信。
>
> skeleton(骨架)：和stub相似，skeleton则是服务端生成的一个代理proxy。当客户端通过stub发送请求到服务端，则交给skeleton来处理，其会根据指定的服务方法来反序列化请求，然后调用具体方法执行，最后将结果返回给客户端。
>
> **registry(服务发现)**：rmi服务，在服务端实现之后需要注册到rmi server上，然后客户端从指定的rmi地址上lookup服务，调用该服务对应的方法即可完成远程方法调用。registry是个很重要的功能，当服务端开发完服务之后，要对外暴露，如果没有服务注册，则客户端是无从调用的，即使服务端的服务就在那里。

#### 通用RPC架构

一般，远程过程调用RPC就是本地动态代理隐藏通信细节，通过组件序列化请求，走网络到服务端，执行真正的服务代码，然后将结果返回给客户端，反序列化数据给调用方法的过程。 

![](../../../md_pictures/rpc调用流程.png)

> 1. serviceClient：这个模块主要是封装服务端对外提供的API，让客户端像使用本地API接口一样调用远程服务。一般，我们使用动态代理机制，当客户端调用api的方法时，serviceClient会走代理逻辑，去远程服务器请求真正的执行方法，然后将响应结果作为本地的api方法执行结果返回给客户端应用。类似RMI的stub模块。
> 2. processor：在服务端存在很多方法，当客户端请求过来，服务端需要定位到具体对象的具体方法，然后执行该方法，这个功能就由processor模块来完成。一般这个操作需要使用反射机制来获取用来执行真实处理逻辑的方法，当然，有的RPC直接在server初始化的时候，将一定规则写进Map映射中，这样直接获取对象即可。类似RMI的skeleton模块。
> 3. protocol：协议层，这是每个RPC组件的核心技术所在。一般，协议层包括编码/解码，或者说序列化和反序列化工作；当然，有的时候编解码不仅仅是对象序列化的工作，还有一些通信相关的字节流的额外解析部分。序列化工具有：hessian，protobuf，avro,thrift，json系，xml系等等。在RMI中这块是直接使用JDK自身的序列化组件。
> 4. transport：传输层，主要是服务端和客户端网络通信相关的功能。这里和下面的IO层区分开，主要是因为传输层处理server/client的网络通信交互，而不涉及具体底层处理连接请求和响应相关的逻辑。
> 5. I/O：这个模块主要是为了提高性能可能采用不同的IO模型和线程模型，当然，一般我们可能和上面的transport层联系的比较紧密，统一称为remote模块。

还有业务代码自己去实现的client和server层。client当需要远程调用服务时，会首先初始化一个API接口代理对象，然后调用某个代理方法。server在对外暴露服务时，需要首先实现对应API接口内部的方法，当请求过来时，通过反射找到对应的实例对象，执行对应的业务代码。 

#### 简单RPC组件实现

**protocol模块代码**：示例中的代码使用JSON来序列化/反序列化工作。由于JSON序列化组件比较弱，所以这边需要将执行调用方法相关的请求数据进行编码成`ProtocolModel`对象。 

**remote模块代码**：在server端的remote中，启动服务之前是需要绑定对外提供的服务的，也就是服务server启动，其内部需要指定序列化、服务处理器等逻辑。 

通用RPC的通信层，是非常复杂的，其需要考虑各种网络环境导致的数据半包，分包和粘包情况，需要考虑高性能NIO组件，多线程处理超时，连接复用等等。 

**processor模块代码**：PROCESSOR_INSTANCE_MAP publish这个逻辑，在**Spring环境**中，一般通过xml配置自动注入进来，然后从context中获取对应的实例。但是，不管怎样，底层其实都是一个map来维护映射关系 

**serviceClient模块代码**：（serviceProxyClient ），`ProxyClient`就是对客户端调用API时透明化底层序列化和网络操作相关细节。所以，在proxyClient内部，我们可以看到它封装代理了这块调用逻辑，业务代码直接使用`getInstance`方法就可以获取对象实例，然后按照正常使用api方法来执行调用逻辑，获取结果。

> 如果使用spring框架的话，可以进一步封装成一个bean，然后客户端业务代码只需要在xml中配置一下，就可以通过注解annotation等方式注入进来。

**server业务接口实现代码** 

### RPC技术深入

#### 序列化介绍

目前基于Java的序列化工具，主要有：

- JDK Serializable工具：很多时候，我们并不使用原生的JDK序列化工具进行序列化，主要原因是因为其序列化后的二进制流太大，并且序列化耗时也比较长。但是，其最大的优点就是原生支持，快速使用，引入成本低，此外，其支持java所有类型，所以在有些RPC组件中，其作为默认序列化工具。
- Hessian工具：hessian2的性能相对JDK来说，提高了很多，而且序列化完了之后的流也小了很多。由于hessian已经生产实践了很长时间，所以其还是很值得使用的。 
- Kryo工具：Kryo是一个快速高效的Java对象序列化框架，其在java的序列化上的性能指标甚至优于google著名的序列化框架protobuf，已经在Twitter、Groupon、Yahoo以及多个著名开源项目（如Hive、Storm）中广泛的使用。总之，Kryo性能非常霸道。 
- JSON工具：其性能上跟hessian差不多，并且反序列化兼容会很，但是其有个比较大的缺点，就是很多类型，可能JSON工具无法支持，并且其是基于String然后再转成二进制流的，所以流的大小，可能并没有想象的那么好。 

####  RPC协议编解码

> 除了序列化，在编码的上/下游还需要对二进制流或者对象做一些额外的处理，而这些处理本身和二进制流化没有太大关系。

比如dubbo给出的处理流程，可以清晰的看出序列化和编码之间的区别(个人觉得广义的编码应该包括序列化那部分)如下 

![](../../../md_pictures/序列化编解码区别.png)

每个RPC组件，基本上都是直接基于Socket来开发通信层功能，但是在网络传输的数据由于网络链路和协议的问题，会出现半包，分包和粘包情况。这样就需要设计编解码协议头来解码网络流，如上dubbo视图。 

**dubbo的协议编码格式**(具体参考：[远程通讯细节](http://dubbo.io/Developer+Guide-zh.htm#DeveloperGuide-zh-%E8%BF%9C%E7%A8%8B%E9%80%9A%E8%AE%AF%E7%BB%86%E8%8A%82))： 

![](../../../md_pictures/dubbo协议编码格式.png)

协议头固定长度`16`个字节，也就是128位。这样，当我们解码流的时候，会首先提取前16byte来解析。 

> `SerializationID`表示序列化类型ID，Dubbo支持多种序列化工具，比如hessian，jdk，fastjson等，所以需要在协议头里面指定序列化方式，这样在解码完了之后才能知道内容使用哪种工具反序列化。
>
> `event`表示事件，比如这个请求是`heartbeat`。`two way`表示请求是否是需要交互返回数据的请求。`req/res`表示该数据是请求还是响应。`status`表示状态位，当响应数据的时候，根据该字段判断是否成功。
>
> `id`表示请求id。这个ID真的真的很重要！！！这个id是请求客户端生成的唯一id，保证在服务运行期内id不会重复。此外，在阿里内部的RPC组件HSF最开始是将id放在data数据内，这样只有在反序列化的时候，才能拿到ReqId，但是有些时候ReqId对应的RPC请求可能由于超时或者已经被处理，导致客户端对于这种case直接丢弃就可以。因此，将id放在head里面，则直接解码的时候就可以拿到ReqId去check，而不需要额外反序列化工作。

#### RPC路由和负载均衡

> 路由策略，是完成单个机器对于服务方调用链路的选择策略,然后把客户端的服务请求传输到具体的某台服务端的机器上。负载均衡是完成路由的一种实现方式，其将前端请求根据一定算法策略来分发到不同机器上，使得集群中机器资源得到充分均衡的利用，此外还可以将不可用机器剔出请求列表。但是，显然路由除了负载均衡之外，还有其他方式。
>
> 我们知道，现在的服务后台都是多台机器部署的服务集群，在这些集群在请求的入口，一般会有负责负载均衡的机器部署，来完成请求的合理分发。RPC的结构也是客户端和服务端模式，但是其结构中我们发现是没有中间代理server层的，所以对于客户端在集群中的远程服务调用，就需要客户端自己来完成负载均衡的逻辑了。
>
> 除负载均衡之外，我们还会存在其他路由加强方式。比如，我们有多个机房都部署服务的时候，我需要优先选择同机房内的服务调用。

> 一般RPC组件中，会实现两个通用的负载均衡策略。随机和轮询。具体实现可以参考：<https://github.com/ketao1989/ourea/tree/master/ourea-core/src/main/java/com/taocoder/ourea/core/loadbalance> 

关于心跳请求的定时任务，可以参考使用Netty提供的`HashedWheelTimer`[netty源码解读之时间轮算法实现-HashedWheelTimer](https://zacard.net/2016/12/02/netty-hashedwheeltimer/)，其提供了在不要求高精度触发定时任务的场景下，性能非常高。 

对于一些完善的RPC框架，可能还会支持动态可配置路由规则。比如，我们可以按照机器ip来配置，某些客户端调用只能路由到某些服务端机器上。对于线上测试问题跟踪而言，可以很好的根据服务调用链路，来查看日志解决问题。 

#### RPC超时管理

增加超时设置是多么重要(当然，连接使用现在最大连接数的连接池也非常重要) 

RPC的调用实现，一般会有一个IO线程池来处理RPC调用，也就是我们的业务线程会将调用请求交给RPC线程来处理，返回一个future对象。远程调用处理完成之后，RPC线程会将结果填充到`futrue`对象内部，然后告知调用方调用完成，可以使用`futrue.get`来获取返回数据。如下所示： 

![](../../../md_pictures/rpc调用实现.png)

#### RPC 服务发现

在对外http服务里，我们有一个配套的支撑基础组件叫做DNS，其根据域名找到某几个外网ip地址。然后，请求打到网站内部，一般首先到nginx群，nginx也会根据url规则找到配置好的一组ip地址，此外，nginx根据healthcheck来检查http服务是否可用。但是，使用nginx时，我们通常需要把ip地址离线配置到nginx上。 

服务发现主要包括2部分：

1. 服务地址存储；
2. 服务状态感知。

服务地址存储

服务地址存储，首先需要一个组件来存放服务机器列表等RPC服务数据，提供存储服务的组件有很多，比如：zookeeper，redis，mysql等等。然后，在服务端正常启动可以提供服务之后，需要将自己的服务地址，比如ip，port，以及服务信息，比如接口，版本号等信息，提交到存储服务机器上。然后，客户端在启动的时候，从存储服务的机器上，根据接口，版本等服务信息来拿到提供对应服务的RPC地址列表，客户端根据这个列表就可以开始调用远程服务了。 对于客户端，为了服务治理，比如我们需要知道哪些客户端调用了我们对外提供的服务，就需要客户端在启动的时候，把自己的地址数据和调用的服务信息提交到存储服务上去。 

对于提供比较完善的服务治理功能，还可以提供后台操作界面，让某些服务端机器手动操作上/下线，这样让通过RPC调用的客户端不将流量打到下线的服务器上。 

![](../../../md_pictures/rpc方法和存储组件交互.png)

服务状态感知

这里的服务感知，包括客户端感知服务端状态，以及存储服务感知RPC参与方的状态。 

正常情况下，我们从存储组件那里拿到服务端地址后，自己来处理路由策略，然后选择一个服务端建立连接，执行远程调用。在执行的过程中，如果有服务不可用，我们可以从我们的服务列表中，将它剔除。但是，如果服务增加机器或者服务机器迁移了呢？这就需要我们及时了解服务端集群的整体机器状态。两种方式：

1. 客户端其一个定时调度任务，周期去存储组件处拉取最新的服务集群地址列表，但是这个**周期粒度**比较难控制。
2. 客户端和存储组件建立一个长连接，当存储组件发现有服务集群状态发生变更，推送给客户端。但是，这又要求存储组件具有**推送功能**。

目前有这个功能的存储组件，主要有zookeeper和redis，此外，也可以自己实现一个简单可靠的服务发现中间件，对外提供推送存储服务。 

我们在服务启动的时候，会告知存储组件我们对外提供服务的地址信息和客户端的地址信息；在服务已知操作的服务下线的时候，会将存储组件中存储的服务相关信息清除掉。但是，显然，在服务下线或者客户端下线的时候，都存在没有清除存储信息就宕机的情况，这个时候就需要存储组件需要有感知各个参与方的状态了。

一般，我们会让RPC两方都和存储组件保持连接，然后通过心跳等方式来探测对方是否下线。

目前提供这个功能的存储组件，主要有zookeeper和redis。当然，你也可以实现一个，可以和所有注册服务和查找服务的server保持长连接。由于，可能有大量的机器建立长连接，所以服务器性能一定要高。

> 基于zookeeper实现服务发现功能的代码，可以参考：<https://github.com/ketao1989/ourea> 

#### RPC 多线程IO模型

### 深入浅出RPC总结

文从单机到集群，从本地调用到远程调用的渐进过度。然后再从一个满足RPC结构图的简单示例开始，代码介绍每个模块，进而深入成熟RPC框架所需要考虑和优化的各个技术点。

本文的目的，旨在对RPC整体结构和各个模块进行介绍和深入，然后根据这些点，可以去分析开源的RPC框架或者自己写一个RPC组件。

*API* 这个词通常用于指代任何通过 REST (HTTP/JSON) 或 Web 服务（SOAP/HTTP） 公开的接口 

所有微服务组件和所涉及的 SaaS 应用程序都将使用常见的协议（比如 HTTP/JSON API）进行通信。 





## 另一文章介绍rpc

最早接触分布式计算时使用的 [CORBAR](https://en.wikipedia.org/wiki/Common_Object_Request_Broker_Architecture) 实现结构基本与此类似。 CORBAR 为了解决异构平台的 RPC，使用了 IDL（Interface Definition Language）来定义远程接口，并将其映射到特定的平台语言中。后来大部分的跨语言平台 RPC 基本都采用了此类方式，比如我们熟悉的 Web Service（SOAP），近年开源的 Thrift 等。他们大部分都通过 IDL 定义，并提供工具来映射生成不同语言平台的 user-stub 和 server-stub，并通过框架库来提供 RPCRuntime 的支持。 不过貌似每个不同的 RPC 框架都定义了各自不同的 IDL 格式，导致程序员的学习成本进一步上升（苦逼啊），Web Service 尝试建立业界标准，无赖标准规范复杂而效率偏低，否则 Thrift 等更高效的 RPC 框架就没必要出现了。 

rpc调用分为同步和异步

> -- 消息头 --
>
> magic      : 协议魔数，为解码设计
>
> header size: 协议头长度，为扩展设计
>
> version    : 协议版本，为兼容设计
>
> st         : 消息体序列化类型
>
> hb         : 心跳消息标记，为长连接传输层心跳设计
>
> ow         : 单向消息标记，
>
> rp         : 响应消息标记，不置位默认是请求消息
>
> status code: 响应消息状态码
>
> reserved   : 为字节对齐保留
>
> message id : 消息 id
>
> body size  : 消息体长度
>
> -- 消息体 --
>
> 采用序列化编码，常见有以下格式
>
> xml   : 如 webservie soap
>
> json  : 如 JSON-RPC
>
> binary: 如 thrift; hession; kryo 等

## RPC框架实例

https://my.oschina.net/huangyong/blog/361751

## gRpc

http://jcf94.com/2016/07/20/2016-07-20-grpc/

http://www.10tiao.com/html/46/201803/2651001766/1.html

https://kuaibao.qq.com/s/20180314A082OB00?refer=spider



## webservice

**一、概念** 

传统上，我们把计算机后台程序（Daemon）提供的功能，称为"服务"（service）。比如，让一个杀毒软件在后台运行，它会自动监控系统，那么这种自动监控就是一个"服务"。**通俗地说，"服务"就是计算机可以提供的某一种功能。** 可分为本地服务和网络服务，网络服务就是通过网络调用其他网站的资源。

```
除了WebService外,常见的远程调用技术还有RMI（Remote method invoke）和CORBA,由于WebService的跨平台和跨编程语言特点,因此比其他两种技术应用更为广泛,但性能略低.


什么是Web Services？
    Web Services 是应用程序组件
    Web Services 使用开放协议进行通信
    Web Services 是独立的（self-contained）并可自我描述
    Web Services 可通过使用UDDI来发现
    Web Services 可被其他应用程序使用
    XML 是 Web Services 的基础

Web services 平台的元素：
    SOAP (简易对象访问协议)
    UDDI (通用描述、发现及整合)
    WSDL (Web services 描述语言)


```



**二、webservice架构和云** 

如果一个软件的主要部分采用了"网络服务"，即它把存储或计算环节"外包"给其他网站了，那么我们就说这个软件属于Web Service架构。**Web Service架构的基本思想，就是尽量把非核心功能交给其他人去做，自己全力开发核心功能。**比如，如果你要开发一个相册软件，完全可以使用Flickr的网络服务，把相片都储存到它上面，你只要全力做好相册本身就可以了。总体上看，凡是不属于你核心竞争力的功能，都应该把它"外包"出去。

**最近很红的"云计算"（cloud computing）或者"云服务"（cloud services），实际上就是Web Service的同义词** ，**它们不说你把事情交给其他计算机去做，而说你把事情交给"云"去做。** 

**三、本地服务的缺陷** 

"网络服务"是未来软件开发和使用的趋势，本地服务将用得越来越少，主要因为以下三个原因：

> *　本地资源不足。很多数据和资料，本地得不到，只有向其他网站要。
>
> *　成本因素。本地提供服务，往往是不经济的，使用专业网站的服务更便宜。这里面涉及硬件和人员两部分，即使你买得起硬件，专门找一个人管理系统，也是很麻烦的事。
>
> *　可移植性差。如果你想把本机的服务，移植到其他机器上，往往很困难，尤其是在跨平台的情况下。

**四、Web Service的优势** 

除了本地服务的缺点以外，Web Service还有以下的优越性：

> *　平台无关。不管你使用什么平台，都可以使用Web service。
>
> *　编程语言无关。只要遵守相关协议，就可以使用任意编程语言，向其他网站要求Web service。这大大增加了web service的适用性，降低了对程序员的要求。
>
> *　对于Web service提供者来说，部署、升级和维护Web service都非常单纯，不需要考虑客户端兼容问题，而且一次性就能完成。
>
> *　对于Web service使用者来说，可以轻易实现多种数据、多种服务的聚合（mashup），因此能够做出一些以前根本无法想像的事情。

**五、Web service的发展趋势** 

根据我的观察，目前Web service有这样几种发展趋势。

> *　在使用方式上，RPC和soap的使用在减少，Restful架构占到了主导地位。
>
> *　在数据格式上，XML格式的使用在减少，json等轻量级格式的使用在增多。
>
> *　在设计架构上，越来越多的第三方软件让用户在客户端（即浏览器），直接与云端对话，不再使用第三方的服务器进行中转或处理数据。

感觉云计算，SaaS都是WebService的具体表现或延伸



不同平台，开发语言之间通过SOAP来交互的web service都能够较好的互通

由于REST只是一种基于Http协议实现资源操作的思想，因此各个网站的REST实现都自有一套，在性能和可用性上会大大高于SOAP发布的web service，但统一通用方面远远不及SOAP。由于这些大网站的SP往往专注于此网站的API开发，因此通用性要求不高。SOAP在成熟度上优于REST。REST还有一个很吸引开发者的就是能够很好的融合当前Web2.0的很多前端技术来提高开发效率。例如很多大型网站开放的REST风格的API都会有多种返回形式，除了传统的xml作为数据承载，还有（JSON,RSS,ATOM）等形式，这对很多网站前端开发人员来说就能够很好的mashup各种资源信息。

**xml+xsd，soap和wsdl是构成webservice平台的三大技术。** 

xml+xsd：webservice采用http协议传输数据，采用xml格式封装数据。

soap：webservice通过http协议发送请求和接收数据时，发送请求的内容和结果都采用xml格式封装，并增加了一些特定的http消息头，以说明http消息的内容格式，这些特定的http消息头和xml内容格式就是soap协议。soap提供了标准的RPC方法来调用webservice。soap协议 = http协议+xml数据格式

wsdl：webservice客户端要调用一个webservice服务，首先要知道这个服务器的地址在哪，以及这个服务器中有哪些方法可以调用，所以webservice服务器首先要通过一个wsdl文件来说明自己家中有哪些方法可以被外界调用，方法包括 参数、返回值等。

wsdl主要描述webservice参数及其返回值。他是webservice客户端和服务器都能理解的标准格式。wsdl文件保存在web服务器上，通过一个url地址就能够访问到它。所以，客户端要调用一个webservice服务之间，要知道该服务的wsdl文件的地址。**webservice服务商通过两种方式发布他的wsdl文件地址**：1>注册到UDDI服务器，以便被人查找；2>直接告诉给客户端调用者。

借助一些webservice框架很容易将自己的业务发布到webservice服务器上，java典型的框架包括：axif，xfire，cxf等。

所谓远程调用，就是一台计算机a上 的一个程序可以调用到另外一台计算机b上的一个对象的方法，譬如，银联提供给商场的pos刷卡系统，商场的POS机转账调用的转账方法的代码其实是跑在银 行服务器上。再比如amazon，天气预报系统，淘宝网，校内网，百度等把自己的系统服务以webservice服务的形式暴露出来，让第三方网站和程 序可以调用这些服务功能，这样扩展了自己系统的市场占有率，往大的概念上吹，就是所谓的SOA应用。



### SOAP

SOAP：简单对象访问协议，简单对象访问协议（SOAP）是一种轻量的、简单的、基于 XML 的协议，它被设计成在 WEB 上交换结构化的和固化的信息。 SOAP 可以和现存的许多因特网协议和格式结合使用，包括超文本传输协议（ HTTP），简单邮件传输协议（SMTP），多用途网际邮件扩充协议（MIME）。它还支持从消息系统到远程过程调用（RPC）等大量的应用程序。



在 RPC  上使用 SOAP 并不仅限于 HTTP 协议绑定。SOAP也可以绑定到TCP和UDP协议上。

SOAP 编码规则：它定义了一种序列化的机制，用于交换应用程序所定义的数据类型的实例。

SOAP RPC 表示：它定义了用于表示远程过程调用和应答的协定。

SOAP 绑定：定义了一种使用底层传输协议来完成在节点间交换SOAP封装的约定。

###  **Restful** 

 REST(Representational State Transfer)一种**轻量级的Web Service架构，可以完全通过HTTP协议实现**。其实现和操作比SOAP和XML-RPC更为简洁，还可以利用缓存Cache来提高响应速度，性能、效率和易用性上都优于SOAP协议。
REST架构对资源的操作包括获取、创建、修改和删除资源的操作正好对应HTTP协议提供的GET、POST、PUT和DELETE方法(Verb) 



**Restful与SOAP的区别** 

 安全性：SOAP会好于restful 

 效率和易用性(REST更胜一筹)

 成熟度(总的来说SOAP在成熟度上优于REST) 



Java中共有三种WebService规范，分别是  

​       JAX-WS(Java API For XML-WebService)  

​       JAXM&SAAJ(Java API For XML Message)  

​       JAX-RS，是Java针对REST(Representation State Transfer)制定的服务规范  

​	JAX-WS和JAXM&SAAJ都是基于SOAP的web服务，相比之下JAXM&SAAJ暴露了SOAP更多的底层细节，如果想控制SOAP消息的更多细节，可以使用；而JAX-WS更加抽象，隐藏了更多的细节，更加面向对象，实现时基本不需要关系SOAP的任何细节。   

​	JAX-RS    REST是一种软件架构模式，不是像SOAP那样本身承载着一种消息协议。REST中重要的两个概念就是资源定位和资源操作，而HTTP协议恰好完整的提供了这两个要点，HTTP协议中的URI可以完成资源定位，GET、POST、OPTION等方法可以完成资源操作，因此REST完全依赖HTTP协议就可以完成Web服务，而不像SOAP协议那样只利用HTTP的传输特性，定位与操作由SOAP协议自身完成。   



webservice的demo

JavaWeb,MVC,Spring IOC

1、知道什么是javaweb的MVC分层，2、知道Spring配置IOC是什么。[新手初识webservice-实践篇](https://my.oschina.net/achi/blog/53227) ，https://cxf.apache.org/ 



**怎样调用Web Service**  

可能这样说不太准确，应该是“怎样调用Web Service中定义的操作 ”      每个Web Service 都有一个描述文件（WSDL ），      它描述一个Web Service 的如下方面：      （1）服务的端口（接收SOAP消息的端口）      （2）服务提供的操作     （3）操作的输入输出格式的定义（通过XMLSchema 定义输入输出格式）       有了Web Service 的描述文件（WSDL ），我们就知道怎样调用这个Web Service 中定义的操作了。      （1）通过服务提供的操作找到你想调用的操作      （2）找到这个操作的输入格式的定义（XMLSchema ），按照这种输入格式构造一个SOAP消息      （3）将这个SOAP消息发送到服务的指定端口      （4）准备接收一个从Web Service服务器返回的 SOAP 响应吧 !



**Web Service服务器端和客户端** 

​    一个Web Service服务器，本质上和一个Web服务器是相同的。 
​    它主要做下面这些事：  
 --> 监听网络端口（监听服务端口）  
 --> 接收客户端请求（接收SOAP请求）  
 --> 解析客户端请求（解析SOAP消息，将SOAP消息转换为数据对象） 
 --> 调用业务逻辑 （调用Web Service实现类的特定操作，参数是由SOAP消息转换而来的数据对象） 
 --> 生成响应 （将返回值转换为SOAP消息） 
 --> 返回响应 （返回SOAP响应）

 

WebService开发可以分为服务器端开发和客户端开发两个方面：

**服务端开发**：把公司内部系统的业务方法发布成WebService服务，供远程合作单位和个人调用。

(借助一些WebService框架可以很轻松地把自己的业务对象发布成WebService服务，

Java方面的典型WebService框架包括：axis，xfire，cxf 等，java ee服务器通常也支持发布WebService服务，例如JBoss。)


**客户端开发：**调用别人发布的WebService服务，大多数人从事的开发都属于这个方面，例如，调用天气预报WebService服务。

（使用厂商的WSDL2Java之类的工具生成静态调用的代理类代码；

使用厂商提供的客户端编程API类；

使用SUN公司早期标准的jax-rpc开发包；

使用 SUN公司最新标准的jax-ws开发包。)

**WebService 的工作调用原理：** 

​	对客户端而言，我们给这各类WebService客户端API传递wsdl文件的url地址，这些API就会创建出底层的代理类，我调用这些代理，就可以访问到webservice服务。代理类把客户端的方法调用变成soap格式的请求数据再通过HTTP协议发出去，并把接收到的soap数据变成返回值返回。

​	 对服务端而言，各类WebService框架的本质就是一个大大的Servlet，当远程调用客户端给它通过http协议发送过来   soap格式的请求数据时，它分析这个数据，就知道要调用哪个java类的哪个方法，于是去查找或创建这个对象，并调用其方法，再把方法返回的结果包装成  soap格式的数据，通过http响应消息回给客户端。



![](/run/media/dhasijkd/learning/md_pictures/网络协议TCP.png)



引用：http://www.ruanyifeng.com/blog/2009/08/what_is_web_service.html





可以在未来物联网，依赖高网速实现服务器处理

RPC就是从一台机器（客户端）上通过参数传递的方式调用另一台机器（服务器）上的一个函数或方法（可以统称为服务）并得到返回的结果。

RPC 会隐藏底层的通讯细节（不需要直接处理Socket通讯或Http通讯） RPC 是一个请求响应模型。

客户端发起请求，服务器返回响应（类似于Http的工作方式） RPC 在使用形式上像调用本地函数（或方法）一样去调用远程的函数（或方法）。



接口调用通常包含两个部分，序列化和通信协议。
　　常见的序列化协议包括json、xml、hession、protobuf、thrift、text、bytes等；
　　通信比较流行的是http、soap、websockect，RPC通常基于TCP实现，常用框架例如dubbo，netty、mina、thrift

> 首先解释下两种接口调用：
> 　　Rest：严格意义上说接口很规范，操作对象即为资源，对资源的四种操作（post、get、put、delete），并且参数都放在URL上,但是不严格的说Http+json、Http+xml，常见的http api都可以称为Rest接口。
> 　　Rpc:我们常说的远程方法调用，就是像调用本地方法一样调用远程方法，通信协议大多采用二进制方式

在微服务中，使用什么协议来构建服务体系，一直是个热门话题。 争论的焦点集中在两个候选技术： (binary) RPC or Restful。
　　业内对微服务的实现，基本是确定一个组织边界，在该边界内，使用RPC; 边界外，使用Restful。这个边界，可以是业务、部门，甚至是全公司。

以Apache Thrift为代表的二进制RPC，支持多种语言（但不是所有语言），四层通讯协议，性能高，节省带宽。相对Restful协议，使用Thrifpt RPC，在同等硬件条件下，带宽使用率仅为前者的20%，性能却提升一个数量级。但是这种协议最大的问题: **无法穿透防火墙**。

在微服务架构中，各个服务之间可能千差万别，**rest接口更加灵活，如果使用RPC则会有很多约束**

- RPC：所谓的远程过程调用 (面向方法)
- SOA：所谓的面向服务的架构(面向消息)
- REST：所谓的 Representational state transfer (面向资源)



RPC阵营如下

\1. Web Service采用HTTP协议做传输层协议，采用SOAP做应用层协议

\2. XML-RPC，采用HTTP协议做传输层协议，使用自定义XML做应用层协议

\3. JMI, Thrift, Protobuf等都使用TCP做传输协议，使用自定义应用层协议

RPC阵营如下

\1. Web Service和XML-RPC采用基于文本的XML进行序列化

\2. RMI基于Java序列化协议

\3. Thrfit, Protobuf等采用了基于二进制流的序列化协议，比如就是采用消息头消息体的方式传输，通过消息头来定义长度，从而保证能够正确读写数据

REST近年来有取代Web Service的趋势，主要是Web Service的优点它都有，而且更轻量级，采用JSON来序列化数据性能也还可以，编程模型更加简单，适合组织内向组织外发布服务。

\1. RPC是面向动作的（方法调用）

\2. REST是面向资源的（URL表示资源，HTTP动词表示动作）

从变现形式来看，RPC的编程模型较重量级，REST的编程模型更轻量级

**Rest**：严格意义上说接口很规范，操作对象即为资源，对资源的四种操作（post、get、put、delete），并且参数都放在URL上,

​         但是不严格的说Http+json、Http+xml，常见的http api都可以称为Rest接口。

**Rpc**:   我们常说的远程方法调用，就是像调用本地方法一样调用远程方法，通信协议大多采用二进制方式

现在开源中间件，基本最先支持的几个协议都包含RESTful

**实际上内部都是采用性能更高的RPC方式。而对外开放的才是RESTful** 

Rest 调用及测试都很方便，Rpc就显得有点麻烦，但是Rpc的效率是毋庸置疑的，所以建议在多系统之间采用Rpc，对外提供服务，Rest是很适合的
  duboo在生产者和消费者两个微服务之间的通信采用的就是Rpc,无疑在服务之间的调用Rpc更变现的优秀

在微服务架构中，各个服务之间可能千差万别，rest接口更加灵活，如果使用RPC则会有很多约束

“描述性状态传递（**re**presentational **s**tate **t**ransfer）,” 





https://www.oschina.net/translate/understanding-rest-and-rpc-for-http-apis?print

https://segmentfault.com/q/1010000000484829

https://www.jianshu.com/p/b0343bfd216e

http://www.ruanyifeng.com/blog/2011/09/restful.html

https://books.studygolang.com/go-rpc-programming-guide/

http://windpoplar.iteye.com/blog/2354931

http://www.voidcn.com/article/p-nkbgkgbg-bd.html

https://www.zhihu.com/question/28557115

https://www.zhihu.com/question/41609070

https://www.zhihu.com/question/28570307

https://blog.csdn.net/sanluo11/article/details/52604584

rpc远程过程调用实例

tr069

protobuf api





# 参考文献

1. [linux内存管理浅析](https://yq.aliyun.com/articles/8931)
2. [微服务、SOA 和 API：是敌是友？](http://www.ibm.com/developerworks/cn/websphere/library/techarticles/1601_clark-trs/1601_clark.html)
3. [序列化和反序列化](http://tech.meituan.com/serialization_vs_deserialization.html)

（1）https://www.jianshu.com/p/dbfac2b876b1

（2）https://segmentfault.com/a/1190000012870365

[分布式RPC框架性能大比拼](http://colobu.com/2016/09/05/benchmarks-of-popular-rpc-frameworks/)