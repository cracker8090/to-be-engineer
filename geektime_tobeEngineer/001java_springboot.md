

	需求变化快和用户群体庞大，在这种情况下，如何从系统架构的角度出发，构建灵活、易扩展的系统，快速应对需求的变化；同时，随着用户的增加，如何保证系统的可伸缩性、高可用性，成为系统架构面临的挑战。分而治之的思想被提了出来，于是我们从单独架构发展到分布式架构，又从分布式架构发展到 SOA 架构，服务不断的被拆分和分解，粒度也越来越小，直到微服务架构的诞生。

	微服务架构是 SOA 架构的传承，但一个最本质的区别就在于微服务是真正的分布式的、去中心化的。把所有的“思考”逻辑包括路由、消息解析等放在服务内部，去掉一个大一统的 ESB，服务间轻通信，是比 SOA 更彻底的拆分。微服务架构强调的重点是业务系统需要彻底的组件化和服务化，原有的单个业务系统会拆分为多个可以独立开发，设计，运行和运维的小应用，这些小应用之间通过服务完成交互和集成。

	每个服务运行在其独立的进程中，服务和服务间采用轻量级的通信机制互相沟通（通常是基于 HTTP 的 RESTful API）。每个服务都围绕着具体业务进行构建，并且能够被独立地部署到生产环境、类生产环境等。另外，应尽量避免统一的、集中式的服务管理机制，对具体的一个服务而言，应根据业务上下文，选择合适的语言、工具对其进行构建。

**数据处理** 

在微服务架构中我们强调彻底的组件化和服务化，每个微服务都可以独立的部署和投产，其实也就意味着很多的微服务有自己独立的数据库。

	整个业务数据被分散在各个子服务之后会带来两个最明显的问题：1、业务管理系统对数据完整的查询，比如分页查询、多条件查询等，数据被割裂后如何来整合？2、如何对数据进一步的分析挖掘？这些需求可能需要分析全量的数据，并且在分析时不能影响到当前业务。
	
	从技术方案来讲，我们一般有两种选择来处理这些问题，第一种是在线处理数据，第二种是离线处理数据。

在线这种方案有两个弊端：1）一方面微服务数据方需要提供数据接口，一方面数据的使用者需要去写调用方法，并且调用者需要编写大量的代码进行数据处理；2）在对各个微服务进行调取数据时会影响微服务的正常业务处理性能。

离线处理数据方案，就是将业务数据准实时的同步到另外一个数据库中，在同步的过程中进行数据整合处理，以满足业务方对数据的需求，数据同步过来后，再提供另外一个服务接口专业负责对外输出数据信息。这种方案有两个特点：1）数据同步方案是关键，技术选型有很多，如何选择切合公司业务的技术方案；2）离线数据处理对微服务正常业务处理没有影响。

第二种方案

MongDB 和数据分析

MongoDB 称之为对开发人员最友好的数据库，不再强调传统关系数据库中的行和列，整个表可以看作一个 Json 文档，MongoDB 也被认为在 Nosql 中最像关系数据库的 Nosql 数据库，保留了类似关系数据库的数据库（DataBase）、集合（Collection）、文档对象（Document）。

MongoDB 最大的特点是支持的查询语言非常强大，其语法有点类似于面向对象的查询语言，几乎可以实现类似关系数据库单表查询的绝大部分功能，而且还支持对数据建立索引。MongoDB 在高可用和读写负载均衡上的实现非常简洁和友好，MongoDB 自带了副本集的概念，通过设计恰当的副本集和驱动程序，可以非常便地实现高可用、读写负载均衡。

MongoDB 的这些特性非常方便对数据进行高性能查询，MongoDB 支持 Aggregate 和 Mapreduce 利用分而治之的理念来处理大规模数据分析。Spring Boot 对 MongoDB 的支持非常友好，使用 Spring Boot 非常便利的处理对 MongoDB 查询和操作，Spring Boot 也提供了组件包来支持对 MongoDB的使用。

MongoDB 4.0 宣布将正式支持 ACID 事务，未来 MongoDB 的想象空间更加巨大！因此 MongDB + Spring Boot 是微服务架构中数据分析的理想选择之一。

# springboot

## Spring Boot 特性

使用 Spring 项目引导页面可以在几秒构建一个项目

方便对外输出各种形式的服务，如 REST API、WebSocket、Web、Streaming、Tasks

非常简洁的安全策略集成

支持关系数据库和非关系数据库

支持运行期内嵌容器，如 Tomcat、Jetty

强大的开发包，支持热启动

自动管理依赖

自带应用监控

支持各种 IED，如 IntelliJ IDEA 、NetBeans

### 优点

总结一下，使用 Spring Boot 至少可以给我们带来以下几方面的改进：

Spring Boot 使编码变简单，Spring Boot 提供了丰富的解决方案，快速集成各种解决方案提升开发效率。

Spring Boot 使配置变简单，Spring Boot 提供了丰富的 Starters，集成主流开源产品往往只需要简单的配置即可。

Spring Boot 使部署变简单，Spring Boot 本身内嵌启动容器，仅仅需要一个命令即可启动项目，结合 Jenkins 、Docker 自动化运维非常容易实现。

Spring Boot 使监控变简单，Spring Boot 自带监控组件，使用 Actuator 轻松监控服务各项状态。

总结，Spring Boot 是 Java 领域最优秀的微服务架构落地技术，没有之一。

	微服务架构下，数据被分隔到 N 个独立的微服务中，如何应对市场、业务对大量数据的查询、分析就变的非常急迫，利用 Spring Boot 和 MongoDB 可以轻松的解决这个问题，通过技术手段将分裂到 N 个微服务的数据同步到 MongoDB 集群中，在同步的过程中进行数据清洗，来满足公司的各项业务需求。Spring Boot 对 MongoDB 的支持非常友好，一方面 Spring Data 技术预生成很多常用方法便于使用，另一方面 Spring Boot 封装了分布式计算的相关函数，可以让我们以较简洁的方式来实现统计查询。

Spring Boot 是 Java 领域微服务架构最优落地技术，Spring Boot+MongoDB 方案是在微服务架构下数据治理的最佳方案之一。



## springboot开源学习



### pom文件解析

#### 什么是pom

pom代表项目对象模型，它是Maven中工作的基本组成单位。它是一个XML文件，始终保存在项目的基本目录中的pom.xml文件中。pom包含的对象是使用maven来构建的，pom.xml文件包含了项目的各种配置信息。 创建一个POM之前，应该要先决定项目组(groupId)，项目名(artifactId)和版本（version），因为这些属性在项目仓库是唯一标识的。需要特别注意，每个项目都只有一个pom.xml文件。





















# 笔记

## KindleNote

<https://github.com/BadTudou/KindleNote-Rails>，https://ruby-china.org/topics/34814

**KindleNote**可以导出您**Kindle**中的**标注**与**笔记**，并支持将它们转换为**MarkDown**文件。

您可以选择将导出的**标记**与**笔记**存储于**Evernote**、**有道云笔记**等云笔记平台，或者**KindleNote**的服务器中。

## 功能

- [x] **笔记导出为Markdown** 
- [x] **笔记保存到Evernote** 
- [x] **批量导出为Markdown / 导出到第三方云笔记 / 删除** 
- [x] **通过豆瓣图书自动获取笔记对应的图书信息** 
- [x] **重复笔记自动合并** 
- [x] **第三方登录：QQ** 
- [x] **第三方登录：Evernote** 
- [ ] **分享到QQ空间、微博等社交网站** 
- [ ] **笔记保存到有道云笔记** 
- [ ] **搜索笔记** 



## 笔记管理工具Clippings

 Kindle Mate

Klib: <https://toolinbox.net/Klib/>



对比一下我用过的几款 PDF 标记软件，主要是 [Marginnote](https://itunes.apple.com/cn/app/marginnote-pro/id723205553?mt=8&at=10lJSw)，[PDF Expert](https://itunes.apple.com/us/app/pdf-expert-by-readdle/id743974925?mt=8)，[iAnnotate 4](https://itunes.apple.com/cn/app/iannotate-4-pdfs-more/id1093924230?mt=8&at=10lJSw)，以及 [Evernote](https://itunes.apple.com/cn/app/evernote/id281796108?mt=8&at=10lJSw) （你没看错，就是Evernote），希望思路对大家有帮助。



# JAVA学习

## 跨平台

需要安装JVM，java虚拟机



## JRE和JDK

java runtime environment，java运行环境

包括JVM和java程序需要的核心类库

如果要运行一个java程序则需要安装jre即可

java development kit，JDKjava开发工具包

JDK提供给开发人员使用，其中保护开发工具，也包括了JRE。其中开发工具：编译工具javac.exe，打包工具jar.exe

JDK开发完成的程序交给jre去运行。







