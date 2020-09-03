[java架构学习导图](https://www.processon.com/view/link/5cb6c8a4e4b059e209fbf369#map) 

47.116.8.77

whatsapp,bridgefy,signal app

[TOC]

redis原理

数据类型

为什么快





# [ehcache、memcache、redis比较](https://www.cnblogs.com/qlqwjy/p/7788912.html)  

## **Ehcache** 

在[Java](http://lib.csdn.net/base/javase)项目广泛的使用。它是一个开源的、设计于提高在数据从RDBMS中取出来的高花费、高延迟采取的一种缓存方案。正因为Ehcache具有健壮性（基于java开发）、被认证（具有apache 2.0  license）、充满特色（稍后会详细介绍），所以被用于大型复杂分布式web application的各个节点中。 

### 1.够快

Ehcache的发行有一段时长了，经过几年的努力和不计其数的性能[测试](http://lib.csdn.net/base/softwaretest)，Ehcache终被设计于large, high concurrency systems.

### 2.够简单

开发者提供的接口非常简单明了，从Ehcache的搭建到运用运行仅仅需要的是你宝贵的几分钟。其实很多开发者都不知道自己用在用Ehcache，Ehcache被广泛的运用于其他的开源项目

比如：[hibernate](http://lib.csdn.net/base/javaee)

### 3.够袖珍

关于这点的特性，官方给了一个很可爱的名字small foot print ，一般Ehcache的发布版本不会到2M，V 2.2.3  才 668KB。

### 4.够轻量

核心程序仅仅依赖slf4j这一个包，没有之一！

### 5.好扩展

Ehcache提供了对[大数据](http://lib.csdn.net/base/hadoop)的内存和硬盘的存储，最近版本允许多实例、保存对象高灵活性、提供LRU、LFU、FIFO淘汰[算法](http://lib.csdn.net/base/datastructure)，基础属性支持热配置、支持的插件多

### 6.监听器

缓存管理器监听器 （CacheManagerListener）和 缓存监听器（CacheEvenListener）,做一些统计或数据一致性广播挺好用的



```
CacheManager manager = CacheManager.newInstance("src/config/ehcache.xml");
Ehcache cache = new Cache("testCache", 5000, false, false, 5, 2);
cacheManager.addCache(cache);
```

## **memcache** 

memcache 是一种高性能、分布式对象缓存系统，最初设计于缓解动态网站[数据库](http://lib.csdn.net/base/mysql)加载数据的延迟性，你可以把它想象成一个大的内存HashTable，就是一个key-value键值缓存。Danga Interactive为了LiveJournal所发展的，以BSD license释放的一套开放源代码软件。

### 1.依赖

memcache [C语言](http://lib.csdn.net/base/c)所编写，依赖于最近版本的GCC和libevent。GCC是它的编译器，同事基于libevent做socket io。在安装memcache时保证你的系统同事具备有这两个环境。

### 2.多线程支持

memcache支持多个cpu同时工作，在memcache安装文件下有个叫threads.txt中特别说明，By default, memcached is compiled as a single-threaded application.默认是单线程编译安装，如果你需要多线程则需要修改./configure --enable-threads，为了支持多核系统，前提是你的系统必须具有多线程工作模式。开启多线程工作的线程数默认是4，如果线程数超过cpu数 容易发生操作死锁的概率。结合自己业务模式选择才能做到物尽其用。

### 3.高性能

通过libevent完成socket 的通讯，理论上性能的瓶颈落在网卡上。

### 4.简单安装

```
1.分别把memcached和libevent下载回来，放到 /tmp 目录下：

\# cd /tmp

\# wget http://www.danga.com/memcached/dist/memcached-1.2.0.tar.gz

\# wget http://www.monkey.org/~provos/libevent-1.2.tar.gz

![img](http://img.my.csdn.net/uploads/201301/17/1358386757_5640.jpg)

2.先安装libevent：

\# tar zxvf libevent-1.2.tar.gz

\# cd libevent-1.2

\# ./configure -prefix=/usr

\# make （如果遇到提示gcc 没有安装则先安装gcc)

\# make install

3.测试libevent是否安装成功：

\# ls -al /usr/lib | grep libevent

lrwxrwxrwx 1 root root 21 11?? 12 17:38 libevent-1.2.so.1 -> libevent-1.2.so.1.0.3

-rwxr-xr-x 1 root root 263546 11?? 12 17:38 libevent-1.2.so.1.0.3

-rw-r-r- 1 root root 454156 11?? 12 17:38 libevent.a

-rwxr-xr-x 1 root root 811 11?? 12 17:38 libevent.la

lrwxrwxrwx 1 root root 21 11?? 12 17:38 libevent.so -> libevent-1.2.so.1.0.3

还不错，都安装上了。

4.安装memcached，同时需要安装中指定libevent的安装位置：

\# cd /tmp

\# tar zxvf memcached-1.2.0.tar.gz

\# cd memcached-1.2.0

\# ./configure -with-libevent=/usr

\# make

\# make install

如果中间出现报错，请仔细检查错误信息，按照错误信息来配置或者增加相应的库或者路径。

安装完成后会把memcached放到 /usr/local/bin/memcached ，

5.测试是否成功安装memcached：

\# ls -al /usr/local/bin/mem*

-rwxr-xr-x 1 root root 137986 11?? 12 17:39 /usr/local/bin/memcached

-rwxr-xr-x 1 root root 140179 11?? 12 17:39 /usr/local/bin/memcached-debug
```



### 5.启动Memcache的服务器端：

```
\# /usr/local/bin/memcached -d -m 8096 -u root -l 192.168.77.105 -p 12000 -c 256 -P /tmp/memcached.pid

-d选项是启动一个守护进程，

-m是分配给Memcache使用的内存数量，单位是MB，我这里是8096MB，

-u是运行Memcache的用户，我这里是root，

-l是监听的服务器IP地址，如果有多个地址的话，我这里指定了服务器的IP地址192.168.77.105，

-p是设置Memcache监听的端口，我这里设置了12000，最好是1024以上的端口，

-c选项是最大运行的并发连接数，默认是1024，我这里设置了256，按照你服务器的负载量来设定，

-P是设置保存Memcache的pid文件，我这里是保存在 /tmp/memcached.pid，


2.如果要结束Memcache进程，执行：

\# cat /tmp/memcached.pid 或者 ps -aux | grep memcache   （找到对应的进程id号）

\# kill 进程id号

也可以启动多个守护进程，不过端口不能重复。

 memcache 的连接

telnet  ip   port 

注意连接之前需要再memcache服务端把memcache的防火墙规则加上

-A RH-Firewall-1-INPUT -m state --state NEW -m tcp -p tcp --dport 3306 -j ACCEPT 

重新加载防火墙规则

service iptables restart

OK ,现在应该就可以连上memcache了

在客户端输入stats 查看memcache的状态信息

![img](http://img.my.csdn.net/uploads/201301/17/1358386780_1884.jpg)



pid              memcache服务器的进程ID

uptime      服务器已经运行的秒数

time           服务器当前的unix时间戳

version     memcache版本

pointer_size         当前[操作系统](http://lib.csdn.net/base/operatingsystem)的指针大小（32位系统一般是32bit）

rusage_user          进程的累计用户时间

rusage_system    进程的累计系统时间

curr_items            服务器当前存储的items数量

total_items           从服务器启动以后存储的items总数量

bytes                       当前服务器存储items占用的字节数

curr_connections        当前打开着的连接数

total_connections        从服务器启动以后曾经打开过的连接数

connection_structures          服务器分配的连接构造数

cmd_get get命令          （获取）总请求次数

cmd_set set命令          （保存）总请求次数

get_hits          总命中次数

get_misses        总未命中次数

evictions     为获取空闲内存而删除的items数（分配给memcache的空间用满后需要删除旧的items来得到空间分配给新的items）

bytes_read    读取字节数（请求字节数）

bytes_written     总发送字节数（结果字节数）

limit_maxbytes     分配给memcache的内存大小（字节）

threads         当前线程数
```

## redis

 [Redis](http://lib.csdn.net/base/redis)是 在memcache之后编写的，大家经常把这两者做比较，如果说它是个key-value store 的话但是它具有丰富的数据类型，我想暂时把它叫做缓存数据流中心，就像现在物流中心那样，order、package、store、 classification、distribute、end。现在还很流行的LAMP [PHP](http://lib.csdn.net/base/php)[架构](http://lib.csdn.net/base/architecture) 不知道和 redis+[MySQL](http://lib.csdn.net/base/mysql) 或者 redis + [MongoDB](http://lib.csdn.net/base/mongodb)的性能比较（听群里的人说mongodb分片不稳定）。













### 1.支持持久化

     redis的本地持久化支持两种方式：RDB和AOF。RDB 在redis.conf配置文件里配置持久化触发器，AOF指的是redis没增加一条记录都会保存到持久化文件中（保存的是这条记录的生成命令），如果 不是用redis做DB用的话还会不要开AOF ，数据太庞大了，重启恢复的时候是一个巨大的工程！

### 2.丰富的数据类型

    redis 支持 String 、Lists、sets、sorted sets、hashes 多种数据类型,新浪微博会使用redis做nosql主要也是它具有这些类型，时间排序、职能排序、我的微博、发给我的这些功能List 和 sorted set的强大操作功能息息相关

### 3.高性能

   这点跟memcache很想象，内存操作的级别是毫秒级的比硬盘操作秒级操作自然高效不少，较少了磁头寻道、数据读取、页面交换这些高开销的操作！这也是NOSQL冒出来的原因吧，应该是高性能

  是基于RDBMS的衍生产品，虽然RDBMS也具有缓存结构，但是始终在app层面不是我们想要的那么操控的。

### 4.replication

    redis提供主从复制方案，跟mysql一样增量复制而且复制的实现都很相似，这个复制跟AOF有点类似复制的是新增记录命令，主库新增记录将新增脚本 发送给从库，从库根据脚本生成记录，这个过程非常快，就看网络了，一般主从都是在同一个局域网，所以可以说redis的主从近似及时同步，同事它还支持一 主多从，动态添加从库，从库数量没有限制。 主从库搭建，我觉得还是采用网状模式，如果使用链式（master-slave-slave-slave-slave·····）如果第一个slave出 现宕机重启，首先从master 接收 数据恢复脚本，这个是阻塞的，如果主库数据几TB的情况恢复过程得花上一段时间，在这个过程中其他的slave就无法和主库同步了。

### 5.更新快

   这点好像从我接触到redis到目前为止 已经发了大版本就4个，小版本没算过。redis作者是个非常积极的人，无论是邮件提问还是论坛发帖，他都能及时耐心的为你解答，维护度很高。有人维护的 话，让我们用的也省心和放心。目前作者对redis 的主导开发方向是redis的集群方向。

### 6.改进

为什么redis不支持多线程多核心处理呢？作者也发表了一下自己的看法，首先是多线程不变于bug的修复，其实是不 易软件的扩展，还有数据一致性问题因为redis所有的操作都是原子操作，作者用到一个词nightmare 噩梦，呵呵！  当然不支持多线程操作，肯定也有他的弊端的比如性能想必必然差，作者从2.2版本后专注redis cluster的方向开发来缓解其性能上的弊端，说白了就是纵向不行，横向提高。

## 应用场景

```
ehcache直接在jvm虚拟机中缓存，速度快，效率高；但是缓存共享麻烦，集群分布式应用不方便。

redis是通过socket访问到缓存服务，效率比ecache低，比数据库要快很多，处理集群和分布式缓存方便，有成熟的方案。

如果是单个应用或者对缓存访问要求很高的应用，用ehcache。

如果是大型系统，存在缓存共享、分布式部署、缓存内容很大的，建议用redis。

补充下：ehcache也有缓存共享方案，不过是通过RMI或者Jgroup多播方式进行广播缓存通知更新，缓存共享复杂，维护不方便；简单的共享可以，但是涉及到缓存恢复，大数据缓存，则不合适

redis和memcached相比的独特之处：
1、redis可以用来做存储(storage),而memcached是用来做缓存(cache)
这个特点主要因为其有持久化功能

2、redis中存储的数据有多种结构，而memcached存储的数据只有一种类型“字符串”
```

## 实际代码操作

```
第二种理解:
第一：两者之间的介绍

Redis：属于独立的运行程序，需要单独安装后，使用JAVA中的Jedis来操纵。因为它是独立，所以如果你写个单元测试程序，放一些数据在Redis中，然后又写一个程序去拿数据，那么是可以拿到这个数据的。，

ehcache：与Redis明显不同，它与java程序是绑在一起的，java程序活着，它就活着。譬如，写一个独立程序放数据，再写一个独立程序拿数据，那么是拿不到数据的。只能在独立程序中才能拿到数据。

第二：使用及各种配置：

两者都可以集群：

1.Redis可以做主从来集群，例如，在A电脑上装个Redis，作为主库；在其他电脑上装Redis，作为从库；这样主库拥有读和写的功能，而从库只拥有读的功能。每次主库的数据都会同步到从库中。

1.默认方式启动

Linux下使用Redis 
安装：从官网上下载tar.gz格式的包，然后使用tar zxvf redis-2.8.24.tar.gz命令解压，然后进入Redis文件夹目录下的src目录，使用make编译一下 
   
1.开启：进入/usr/local/redis-3.2.1/src 
然后./redis-server 

2.如果我们想修改端口，设置密码：那么得修改配置文件的redis.conf
port 6379                                         //端口修改
requirepass redis123                  //设置密码  redis123为密码

配置主从：主库的配置文件不用修改，从库的配置文件需要修改，因为从库需要绑定主库，以便可以获取主库的数据

slaveof 192.168.1.100 6379                              //主库的IP地址和端口号
masterauth redis123                                      //主库设定的密码

3.要让配置文件的属性生效，那么启动的redis的时候，要将配置文件加上去

进入/usr/local/redis-3.2.1/src
然后 ./redis-server  redis.conf

那么将成功的启动redis，如果没有加入配置的话，按照普通方式启动的话，端口仍然还是6379.

4.客户端连接远程的Redis
第一步：在远程端处设置密码：config set requirepass 123       //123为密码
第二步：可以在客户端登录  redis-cli.exe -h 114.215.125.42 -p 6379 
第三步：认证：auth 123                                       //123为密码
本地端设置密码后，要使用密码登录；如果Redis重启的话，密码需要重新设置

5.主从配置后，为保证主库写的能力，一般不在主库做持久化，而是在从库做持久化：

主库配置：
将save注释，不使用rdb
# save 900 1
# save 300 10
# save 60 10000

appendonly no        不使用aof

从库配置：
save 900 1
save 300 10
save 60 10000

appendonly yes

这样做的优缺点：
优点：保证了主库写的能力。

缺点：主库挂掉后，重启主库，然后进行第一次写的动作后，主库会先生成rdb文件，然后传输给从库，从而覆盖掉从库原先的rdb文件，造成数据丢失。但是第二次写的时候，主库会以快照方式直接传数据给从库，不会重新生成rdb文件。

解决方案：先复制从库中的数据到主库后，再启动主库。

使用：
引入jedis包
    <dependency>  
        <groupId>redis.clients</groupId>  
        <artifactId>jedis</artifactId>  
        <version>2.7.3</version>  
    </dependency>  
简单的写个类玩玩吧
    public class RedisMain {  
        public static void main(String [] str)  
        {  
              
             Jedis jedis = new Jedis("114.215.125.42",6379);  
             jedis.auth("123");     //密码认证  
              System.out.println("Connection to server sucessfully");  
              //查看服务是否运行  
              jedis.set("user","namess");  
             // System.out.println("Server is running: "+jedis.ping());  
              System.out.println(jedis.get("user").toString());  
              jedis.set("user","name");  
              System.out.println(jedis.get("user"));  
        }  
    }  
Redis完毕

下面说说Ehcache:

Ehcache的使用：
1.首先引入包
    <dependency>  
               <groupId>net.sf.ehcache</groupId>  
               <artifactId>ehcache-core</artifactId>  
               <version>2.6.6</version>  
           </dependency>  
       
           <dependency>  
               <groupId>org.slf4j</groupId>  
               <artifactId>slf4j-log4j12</artifactId>  
               <version>1.6.6</version>  
           </dependency>  

2.创建一个ehcache.xml文件，里面配置cache的信息，这个配置是包含了集群的配置：与192.168.93.129：40001的 机器集群了：Ip为192.168.93.129机子的配置要将rmiUrls对应的数据改为这个配置文件的机子的IP地址，和对应的缓存名字

    <ehcache xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"  
       xsi:noNamespaceSchemaLocation="ehcache.xsd">  
        <cacheManagerPeerProviderFactory   
       class="net.sf.ehcache.distribution.RMICacheManagerPeerProviderFactory"  
    properties="peerDiscovery=manual,rmiUrls=//192.168.93.129:40001/demoCache"/>           <!--另一台机子的ip缓存信息-->  
    <cacheManagerPeerListenerFactory class="net.sf.ehcache.distribution.RMICacheManagerPeerListenerFactory"  
    properties="hostName=localhost,port=40001,socketTimeoutMillis=2000" />                  <!--hostName代表本机子的ip-->     
      <diskStore path="java.io.tmpdir"/>  
      <defaultCache  
        maxElementsInMemory="10000"  
        maxElementsOnDisk="0"  
        eternal="true"  
        overflowToDisk="true"  
        diskPersistent="false"  
        timeToIdleSeconds="0"  
        timeToLiveSeconds="0"  
        diskSpoolBufferSizeMB="50"  
        diskExpiryThreadIntervalSeconds="120"  
        memoryStoreEvictionPolicy="LFU"  
        >  
         <cacheEventListenerFactory    
                    class="net.sf.ehcache.distribution.RMICacheReplicatorFactory"/>                
        </defaultCache>  
      <cache name="demoCache"  
        maxElementsInMemory="100"  
        maxElementsOnDisk="0"  
        eternal="false"  
        overflowToDisk="false"  
        diskPersistent="false"  
        timeToIdleSeconds="119"  
        timeToLiveSeconds="119"  
        diskSpoolBufferSizeMB="50"  
        diskExpiryThreadIntervalSeconds="120"  
        memoryStoreEvictionPolicy="FIFO"  
        >  
        <cacheEventListenerFactory class="net.sf.ehcache.distribution.RMICacheReplicatorFactory"/>          <!--监听这个cache-->  
        </cache>  
    </ehcache>  

配置完后写代码：

放数据：
    @RequestMapping("/testehcache.do")  
        public void testehcache(HttpServletResponse response) throws IOException  
        {  
            URL url = getClass().getResource("ehcache.xml");   
             CacheManager singletonmanager = CacheManager.create(url);  
            Cache cache = singletonmanager.getCache("demoCache");  
            //使用缓存  
            Element element = new Element("key1", "value1");  
            cache.put(element);  
            cache.put(new Element("key2", "value2"));  
             
            response.getWriter().println("我存放了数据");  
        }  

拿数据：
@RequestMapping("/getcache.do")  
    public void getcache(HttpServletResponse response) throws IOException  
    {  
        CacheManager singletonmanager = CacheManager.create();   
        Cache cache = singletonmanager.getCache("demoCache");  
        String one=cache.get("key1").getObjectValue().toString();  
        String two=cache.get("key2").getObjectValue().toString();  
        response.getWriter().println(one+two);  
    } 
配置集群后，A机器放数据，在B机器上能拿到数据，B机器放数据，A机器也可以拿到数据
```



# Redis学习

## 1.书籍推荐

[redis实战](https://book.douban.com/subject/26612779/)  《Redis实战》是《Redis in Action》

在线阅读地址：<http://redisinaction.com/index.html>

redis设计与实现  [带有详细注释的 Redis 2.6 源码](https://github.com/huangz1990/annotated_redis_source/) 

在线阅读地址：<http://redisbook.readthedocs.io/en/latest/index.html>

[redis在线学习网站](http://try.redis.io/) 

[the-little-redis-book](https://github.com/JasonLai256/the-little-redis-book/blob/master/cn/redis.md)  

[redis命令参考](http://doc.redisfans.com/) 

<https://redis.readthedocs.io/en/latest/> 

## 2.[非关系型数据库可以学习的课程](https://www.jianshu.com/p/b768b7adae12) 

**课程：面向文档数据库——【[mongoDB基础教程](https://www.shiyanlou.com/courses/12)】**

Mongo最大的特点是他支持的查询语言非常强大，其语法有点类似于面向对象的查询语言，几乎可以实现类似关系数据库单表查询的绝大部分功能，而且还支持对数据建立索引。

通过课程了解monggoDB的基本操作、数据查询、文档操作、以及一些高级语法；

课程：键值存储数据库——【[Redis基础教](https://www.shiyanlou.com/courses/106)程】

Redis 是一个高性能的key-value数据库。Redis支持主从同步，可执行单层树复制。

课程介绍Redis系统的基本配置和使用方法。

课程：列存储数据库——【[HBASE基础教程](https://www.shiyanlou.com/courses/37)】

HBASE是Hadoop项目的一部分，运行于HDFS文件系统之上，为 Hadoop 提供类似于BigTable 规模的服务。

通过课程了解HBASE的基础配置以及使用方法。

## 3.Redis实战

该书深入浅出地介绍了 Redis 的字符串、列表、散列、集合、有序集合等五种结构， 并通过文章聚合网站、cookie、购物车、网页缓存、日志、计数器、IP 所属地址查询程序、自动补全、分布式锁、计数信号量、任务队列、消息队列、搜索程序、广告定向程序、社交网站等一系列实用示例展示了 Redis 的用法。

除此之外， 《Redis实战》还介绍了使用短结构、分片、事务、流水线、复制、Lua 脚本等手段来扩展和优化 Redis 的方法， 这些技术可以大幅地扩展系统的性能， 并尽可能地降低程序所需的内存数量。

综上所述， 《Redis实战》将是一本对于学习和使用 Redis 来说不可多得的参考书籍， 无论是 Redis 新手还是有一定经验的 Redis 使用者， 应该都能从本书中获益。







### redis配置

　　daemonize：如需要在后台运行，把该项的值改为yes

　　pdifile：把pid文件放在/var/run/redis.pid，可以配置到其他地址

　　bind：指定redis只接收来自该IP的请求，如果不设置，那么将处理所有请求，在生产环节中最好设置该项

　　port：监听端口，默认为6379

　　timeout：设置客户端连接时的超时时间，单位为秒

　　loglevel：等级分为4级，debug，revbose，notice和warning。生产环境下一般开启notice

　　logfile：配置log文件地址，默认使用标准输出，即打印在命令行终端的端口上

　　database：设置数据库的个数，默认使用的数据库是0

　　save：设置redis进行数据库镜像的频率

　　rdbcompression：在进行镜像备份时，是否进行压缩

　　dbfilename：镜像备份文件的文件名

　　dir：数据库镜像备份的文件放置的路径

　　slaveof：设置该数据库为其他数据库的从数据库

　　masterauth：当主数据库连接需要密码验证时，在这里设定

　　requirepass：设置客户端连接后进行任何其他指定前需要使用的密码

　　maxclients：限制同时连接的客户端数量

　　maxmemory：设置redis能够使用的最大内存

　　appendonly：开启appendonly模式后，redis会把每一次所接收到的写操作都追加到appendonly.aof文件中，当redis重新启动时，会从该文件恢复出之前的状态

　　appendfsync：设置appendonly.aof文件进行同步的频率

　　vm_enabled：是否开启虚拟内存支持

　　vm_swap_file：设置虚拟内存的交换文件的路径

　　vm_max_momery：设置开启虚拟内存后，redis将使用的最大物理内存的大小，默认为0

　　vm_page_size：设置虚拟内存页的大小

　　vm_pages：设置交换文件的总的page数量

　　vm_max_thrrads：设置vm IO同时使用的线程数量







## 4.Netty、Redis、Zookeeper高并发实战pdf







## 5.redis and docker

### 1.阿里云docker部署redis

```
docker search redis
docker pull redis
docker images

docker run -p 6379:6379 -v $PWD/data:/data --name redis_1 -d redis redis-server --appendonly yes
docker run -p 6379:6379 -v /docker/redis/redis.conf:/etc/redis/redis.conf -v /docker/redis/data:/data --name redis1 -d redis redis-server --appendonly yes

    docker run -p 6379:6379 -v /docker/redis/redis.conf:/etc/redis/redis.conf -v /docker/redis/data:/data --name myredis -d redis redis-server --appendonly yes


//-p  指定端口  
//-v  指定目录下的data目录挂载
//-name 指定容器名称
//-d 指定镜像名称后台运行
//redis-server --appendonly yes  在容器执行redis-server启动命令，并打开redis持久化配置


下载和编译
$ wget http://download.redis.io/releases/redis-5.0.7.tar.gz
$ tar xzf redis-5.0.7.tar.gz
$ cd redis-5.0.7
$ make

阿里云docker配置：
redis.conf文件中有个daemonize yes的初始默认值，默认redis以后台形式运行会和docker的启动选项冲突，必须注释掉或者改为no
redis.conf的bind 127.0.0.1注释掉
requirepass 123456
docker run -d -p 6379:6379 -v /docker/redis/redis.conf:/etc/redis/redis.conf -v /docker/redis/data:/data --name redis_1 redis redis-server /etc/redis/redis.conf
-p：指定端口映射，前面的代表容器的pid，后面的代表服务器的pid
-v：配置文件和持久化存储的挂载，前面的是服务器的文件路径，后面的是容器的路径。
最后指定使用容器中的/etc/redis/redis.conf配置文件启动redis。

#bind 127.0.0.1 //允许远程连接
protected-mode no
appendonly yes //持久化
requirepass 123456 //密码 

netstat -lntp | grep 6379

config set requirepass "yourpassword" // 设置当前密码，服务重新启动后又会置为默认，即无密码；
config get requirepass // 获取当前密码
添加requirepass yourpassword（此处注意，行前不能有空格），保存之后，重启Redis服务。

config set requirepass 123456
config get requirepass
service redis restart

master配置了密码，slave如何配置
若master配置了密码则slave也要配置相应的密码参数否则无法进行正常复制的。
slave中配置文件内找到如下行，移除注释，修改密码即可
#masterauth  mstpassword
```

### 2.阿里云部署redis集群

[使用Docker Compose部署基于Sentinel的高可用Redis集群](https://yq.aliyun.com/articles/57953?spm=a2c6h.13066369.0.0.ab14346bYQ0Fkr&do=login) 

<https://www.imooc.com/article/277939> 容器配置内网ip 桥接



### 3.docker部署redis与其他容器

使用Docker的时候，经常可能需要连接到其他的容器，比如：web服务需要连接数据库。按照往常的做法，需要先启动数据库的容器，映射出端口来，然后配置好客户端的容器，再去访问。其实针对这种场景，Docker提供了--link 参数来满足。
--link=container_name or id:name

```
$ docker run --name some-app --link some-redis:redis -d application-that-uses-redis
或者
$ docker run -it --link some-redis:redis --rm redis redis-cli -h redis -p 6379

docker run --name=mysql_server -d -P kongxx/mysql_server
docker run --name=mysql_client1 --link=mysql_server:db -t -i kongxx/mysql_client /usr/bin/mysql -h db -u root -pletmein
“–link=mysql_server:db”，这个参数就是告诉Docker容器需要使用“mysql_server”容器，并将其别名命名为db，这样在这两个容器里就可以使用“db”来作为提供mysql数据库服务的机器名。所以在最后启动参数里我们使用的是“/usr/bin/mysql -h db -u root -pletmein”来连接mysql数据库的。
jdbc:mysql://db:3306/test?useUnicode=true&characterEncoding=utf8&useSSL=false

```



*systemctl statusfirewalld*

*systemctl stop firewalld* 

docker network ls

默认情况下创建的所有容器都会在bridge网段

docker network inspect bridge 查看bridge网段详情

docker run -it -d --name test1 busybox
docker run -it -d --name test2 busybox

*sudo yum install -y bridge-utils这样执行 sudo brctl show*
可以清晰简单的看到连接到各网段的容器

docker run -d -it --link test2 --name test3 busybox

ping -c 4 test2



docker network create --driver bridge my-bridge

其中–driver是表示基于后面参数bridge建立的网段my-bridge
docker run -it -d --net=my-bridge --name test4 busybox 创建一个在my-bridge网段的容器

test4 和test1.2.3都不在一个网段

docker network connect bridge test4 把test4也加入bridge网段（test4会有两个ip）

进入test4网段随意ping一下bridge网段的容器





#### docker4种网络模式

1：bridge模式，--net=bridge(默认)。
这是dokcer网络的默认设置。安装完docker，系统会自动添加一个供docker使用的网桥docker0，我们创建一个新的容器时，容器通过DHCP获取一个与docker0同网段的IP地址。并默认连接到docker0网桥，以此实现容器与宿主机的网络互通。

2：host模式，--net=host。

  这个模式下创建出来的容器，将不拥有自己独立的Network Namespace，即没有独立的网络环境。它使用宿主机的ip和端口。

3：container模式，--net=container:NAME_or_ID。

这个模式就是指定一个已有的容器，共享该容器的IP和端口。除了网络方面两个容器共享，其他的如文件系统，进程等还是隔离开的。

4：none模式，--net=none。

这个模式下，dokcer不为容器进行任何网络配置。需要我们自己为容器添加网卡，配置IP。

因此，若想使用pipework配置docker容器的ip地址，必须要在none模式下才可以



docker通信工具：Pipework和Open vSwitch

pipework的工具。号称是容器网络的SDN解决方案，可以在复杂的场景下将容器连接起来。它既支持普通的LXC容器，也支持Docker容器。



1.同一主机容器之间



2.不同主机容器之间

```ruby
brctl addbr br0
[root@node1 ~]# ip link set dev br0 up
[root@node1 ~]# ip addr add 192.168.114.1/24 dev br0//这个ip相当于br0网桥的网关ip，可以随意设定。
docker run -ti -d --net=none --name=my-test2 docker.io/nginx /bin/bash
[root@node1 ~]# pipework br0 -i eth0 my-test12 192.168.114.200/24@192.168.114.1
[root@node1 ~]# pipework br0 -i eth0 my-test2 192.168.114.200/24@192.168.114.1
```

上面使用的pipework工具还，还可以使用虚拟交换机(Open vSwitch)进行docker容器间的网络通信

iproute2 正在逐步取代旧的 net-tools（ifconfig）



### 4.linux安装redis

```
cd redis-2.8.17
make 
cd src
make install PREFIX=/usr/local/redis

/usr/local/redis/bin目录下的几个关键文件

redis-benchmark：redis性能测试工具
redis-check-aof：检查aof日志的工具
redis-check-dump：检查rdb日志的工具
redis-cli：连接用的客户端
redis-server：redis服务进程

配置介绍
daemonize：如需要在后台运行，把该项的值改为yes
pdifile：把pid文件放在/var/run/redis.pid，可以配置到其他地址
bind：指定redis只接收来自该IP的请求，如果不设置，那么将处理所有请求，在生产环节中最好设置该项
port：监听端口，默认为6379
timeout：设置客户端连接时的超时时间，单位为秒
loglevel：等级分为4级，debug，revbose，notice和warning。生产环境下一般开启notice
logfile：配置log文件地址，默认使用标准输出，即打印在命令行终端的端口上
database：设置数据库的个数，默认使用的数据库是0
save：设置redis进行数据库镜像的频率
rdbcompression：在进行镜像备份时，是否进行压缩
dbfilename：镜像备份文件的文件名
dir：数据库镜像备份的文件放置的路径
slaveof：设置该数据库为其他数据库的从数据库
masterauth：当主数据库连接需要密码验证时，在这里设定
requirepass：设置客户端连接后进行任何其他指定前需要使用的密码
maxclients：限制同时连接的客户端数量
maxmemory：设置redis能够使用的最大内存
appendonly：开启appendonly模式后，redis会把每一次所接收到的写操作都追加到appendonly.aof文件中，当redis重新启动时，会从该文件恢复出之前的状态

appendfsync：设置appendonly.aof文件进行同步的频率
vm_enabled：是否开启虚拟内存支持
vm_swap_file：设置虚拟内存的交换文件的路径
vm_max_momery：设置开启虚拟内存后，redis将使用的最大物理内存的大小，默认为0
vm_page_size：设置虚拟内存页的大小
vm_pages：设置交换文件的总的page数量
vm_max_thrrads：设置vm IO同时使用的线程数量
```



pip install docker-compose



1. Connection is idle for 30seconds and Redis kills the same.
2. Before "Redis connection factory" in application detects the broken connection, it gets allocation for read or write request
3. Code tries to use this connection but as it is broken, it is unable to send command for read/write. Thus we get "JedisConnectionException: Unexpected end of stream" exception

**Solution** 

1. set Redis timeout to Zero
2. Using custom JedisPoolConfig set the minEvictableIdleTimeMillis to desired value. This will ensure idle connections are released from Jedis connection pool



### 6.Jedispool操作redis

Jedis连接就是资源，JedisPool管理的就是Jedis连接。

1、资源设置和使用

maxTotal：资源池中最大连接数；默认值：8	设置建议见下节
maxIdle：资源池允许最大空闲的连接数；默认值：8；使用建议：设置建议见下节
minIdle：资源池确保最少空闲的连接数；默认值：0；使用建议：设置建议见下节
blockWhenExhausted：当资源池用尽后，调用者是否要等待。只有当为true时，下面的maxWaitMillis才会生效；默认值：true；使用建议：建议使用默认值
maxWaitMillis：当资源池连接用尽后，调用者的最大等待时间(单位为毫秒)	-1：表示永不超时；使用建议：不建议使用默认值
testOnBorrow：向资源池借用连接时是否做连接有效性检测(ping)，无效连接会被移除；默认值：false；使用建议：业务量很大时候建议设置为false(多一次ping的开销)。
testOnReturn：向资源池归还连接时是否做连接有效性检测(ping)，无效连接会被移除；默认值：false；使用建议：业务量很大时候建议设置为false(多一次ping的开销)。
jmxEnabled：是否开启jmx监控，可用于监控；默认值：true；使用建议：建议开启，但应用本身也要开启

2、空闲资源监测

空闲Jedis对象检测，下面四个参数组合来完成，testWhileIdle是该功能的开关。

testWhileIdle：是否开启空闲资源监测；默认值：false；使用建议：true
timeBetweenEvictionRunsMillis：空闲资源的检测周期(单位为毫秒)；默认值：-1：不检测；使用建议：建议设置，周期自行选择，也可以默认也可以使用下面JedisPoolConfig中的配置
minEvictableIdleTimeMillis：资源池中资源最小空闲时间(单位为毫秒)，达到此值后空闲资源将被移除；默认值：1000 60 30 = 30分钟；使用建议：可根据自身业务决定，大部分默认值即可，也可以考虑使用下面JeidsPoolConfig中的配置
numTestsPerEvictionRun：做空闲资源检测时，每次的采样数；默认值：3；使用建议：可根据自身应用连接数进行微调,如果设置为-1，就是对所有连接做空闲监测

```
# REDIS（RedisProperties）
# （普通集群，不使用则不用开启）在群集中执行命令时要遵循的最大重定向数目。
# spring.redis.cluster.max-redirects=
# （普通集群，不使用则不用开启）以逗号分隔的“主机：端口”对列表进行引导。
# spring.redis.cluster.nodes=
# 连接工厂使用的数据库索引。
spring.redis.database=0
# 连接URL，将覆盖主机，端口和密码（用户将被忽略），例如：redis://user:password@example.com:6379
#spring.redis.url=
# Redis服务器主机。
#spring.redis.host=localhost
spring.redis.host=47.116.8.77
# 登录redis服务器的密码。
spring.redis.password=123456
# 启用SSL支持。
spring.redis.ssl=false
# 池在给定时间可以分配的最大连接数。使用负值无限制。
#spring.redis.pool.max-active=8
spring.redis.jedis.pool.max-active=8
# 池中“空闲”连接的最大数量。使用负值表示无限数量的空闲连接。
#spring.redis.pool.max-idle=8
spring.redis.jedis.pool.max-idle=8
# 连接分配在池被耗尽时抛出异常之前应该阻塞的最长时间量（以毫秒为单位）。使用负值可以无限期地阻止。
#spring.redis.pool.max-wait=-1
spring.redis.jedis.pool.max-wait=-1
# 目标为保持在池中的最小空闲连接数。这个设置只有在正面的情况下才有效果。
#spring.redis.pool.min-idle=0
spring.redis.jedis.pool.min-idle=0

# Redis服务器端口。
spring.redis.port=6379
# （哨兵模式，不使用则不用开启）Redis服务器的名称。
# spring.redis.sentinel.master=
# （哨兵模式，不使用则不用开启）主机：端口对的逗号分隔列表。 
# spring.redis.sentinel.nodes=
# 以毫秒为单位的连接超时。
spring.redis.timeout=0

# （普通集群，不使用则不用开启）在群集中执行命令时要遵循的最大重定向数目。
#spring.redis.cluster.max-redirects=
# （普通集群，不使用则不用开启）以逗号分隔的“主机：端口”对列表进行引导。
#spring.redis.cluster.nodes=127.0.0.1:1001,127.0.0.1:1002
```

<http://127.0.0.1:8080/redis/get/xbp3> 

<https://blog.lqdev.cn/2018/07/23/springboot/chapter-eleven/> 

Redis command timed out; nested exception is io.lettuce.core.RedisCommandTimeoutException: Command timed out after no timeout



### 7.springboot操作redis

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-redis</artifactId>
    <version>1.3.8.RELEASE</version>
</dependency>
```

<http://127.0.0.1:8080/redis/set/name/xubeiping> 

`Spring Cache`是`Spring`框架提供的对缓存使用的抽象类，支持多种缓存，比如`Redis`、`EHCache`等，集成很方便。同时提供了多种注解来简化缓存的使用，可对方法进行缓存。

修改`RedisConfig`配置类，加入注解`@EnableCaching`，同时设置`CacheManager`缓存管理类，这里使用`RedisCacheManager`，其他的管理类还有：`SimpleCacheManager`、`ConcurrentMapCacheManager`等，默认提供的在类`org.springframework.cache.support`下，可自行查阅。















































