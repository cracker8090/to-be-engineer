[toc] 

# mysql必知必会

## 数据库基本知识

### 什么是数据库

数据库是一个以某种有组织的方式存储数据集 合的软件。 数据库：保存有组织的数据的容器，通 常是一个文件或一组文件。    数据库软件应称为DBMS（数据库管理系统） 。在很大程度上说，数据库究竟是 文件还是别的什么东西并不重要，因为你并不直接访问数据 库；你使用的是DBMS，它替你访问数据库。    

### 表

在文件柜 中创建文件夹，然后将相关的资料放进特定的文件中。 在数据库领域中，这种文件夹称为 表。表是一种结构化的文件，可用来存储某种特定类型的数据。表可以保存顾客清单，产品 目录，或者其它信息清单。

	 表：某种特定类型数据的机构化清单。    

这里关键的一点在于，存储在表中的数据是一种类型的数据或一个清单。绝不应该将顾客的 清单与订单的清单存储在同一个数据库表中。这样做将使以后的检索和访问很困难。 数据库中的每个表都有一个名字，用来标识自己。此名字时唯 一的，这表示数据库中没有其它表具有相同的名字。    

	表名：表名的唯一性取决于多个因素，如数据库名和表名等的结合。虽然在同 一个数据库不能使用相同的表名，但在不同的数据库中却可以使用相同的表名    

表具有一些特性，这些特性定义了数据在表中如何存储，如可以**存储什么样的数据，数据如何分解，各部分信息如何命名**，等等。描述表的这组信息就是所谓的模式，模式可以用来描 述数据库中特定的表以及整个数据库（和其中表的关系） 。

	 模式：关于数据库和表的布局以及特性的信息    

### 列和数据类型 

表由列组成。列中存储着表中某部分的信息。    列：表中的一个字段。所有表都是由一个或多个列组成的。 理解列的最好办法是将表想象为 一个网格。网格中每一列存储着一条特定的信息。例如，在顾客表中，一个列存储着顾客编 号，另一个列存储着顾客名，而地址，城市，州以及邮政编号全都存储在各自的列中。 数据 库中每个列都有相应的数据类型。数据类型定义列可以存储的数据种类。    

	数据类型：所允许存储的数据的类型。每个列都有相应的数据数据类型，它限制该列中 存储的数据。    

数据类型限制可以存储在列中的数据种类。数据类型还帮助正确地排序数据，并在优化磁盘 使用方面起重要作用。因此，在创建表时必须对数据类型给予特别的关注。    

### 行

表中的数据是按行存储的，所保存的每个记录存储在自己的行内。如果将表想象为网格，网 格中垂直的列为表列，水平行为表行。 例如，顾客表可以每行存储一个顾客。表中的行数为 记录的总数。    

### 主键

表中每一行都应该有可以唯一标识自己的列。一个顾客表可以使用顾客编号列，而订单表可 以使用订单ID，雇员表可以使用雇员ID或雇员社会保险号。    唯一标识表中每行的这个列称为主键。主键用来表示一个特定的行。    表中的任何列都可以作为 主键，只要它满足以下条件：    

（1）任意两行都不具有相同的主键值    

（2）每个行都必须有一个主键值（不允许为null值）。主键通常定义的表的一列上，但这并不是必 须的，也可以一起使用多个列作为主键。在使用多列作为主键时，上述条件必须应用到 构成主键的所有列，所有列值的组合必须是唯一的（单个列的值可以不唯一） 。 还有一 种非常重要的键，称为外键，我们将在第15章中介绍。    

	主键：其值能够唯一标识表中每行数据    

第15章介绍外键。

SQL（Structured Query Language）是一种专门用来与数据库通信的语言。

SQL有如下的优点。

  SQL不是某个特定数据库供应商专有的语言。几乎所有重要的 DBMS都支持SQL，所以，学习此语言使你几乎能与所有数据库 打交道。

  SQL简单易学。它的语句全都是由描述性很强的英语单词组成， 而且这些单词的数目不多。

  SQL尽管看上去很简单，但它实际上是一种强有力的语言，灵活 使用其语言元素，可以进行非常复杂和高级的数据库操作。    

## 安装使用

事实上，多数网络的建立使用户不具有对数据的访问权，甚至不 具有对存储数据的驱动器的访问权。    

<http://dev.mysql.com/downloads/windows/installer/> 

安装路径可以修改一下 ，安装过程中的设置的密码很重要，要牢记。安装好之后，重要的一步是打开服务：控制面板 -> 管理工具 -> 服务，在里面找到MySQL56选项，右击->属性，在启动类型中选择 自动，然后点击启动，最后确定即可。 

 点击 MySQL 5.6 Command Line Client 启动命令行模式，输入密码： 

mysql工具

mysql命令行实用程序    

MySQL Administrator（MySQL管理器）

MySQL Query Browser



port：3306

root:1a2b3c4d5e6f:;!

windows service name:MySQL80



使用mysql



客户端

MySQL Workbench

Navicat for MySQL 

phpMyAdmin



# mysql安装使用

在CentOS中默认安装有MariaDB，这个是MySQL的分支，但为了需要，还是要在系统中安装MySQL，而且安装完成之后可以直接覆盖掉MariaDB。（mysql -u root -p 104104aa）

## 1 下载并安装MySQL官方的 Yum Repository

```
[root@localhost ~]# wget -i -c http://dev.mysql.com/get/mysql57-community-release-el7-10.noarch.rpm
```

  使用上面的命令就直接下载了安装用的Yum Repository，大概25KB的样子，然后就可以直接yum安装了。

```
[root@localhost ~]# yum -y install mysql57-community-release-el7-10.noarch.rpm
```

**注意：必须进入到 /etc/yum.repos.d/目录后再执行以下脚本** 

yum install mysql-server

systemctl start mysqld 

service mysqld start 

## 2.MariaDB 

此外,你也可以使用 MariaDB 代替，MariaDB 数据库管理系统是 MySQL 的一个分支，主要由开源社区在维护，采用 GPL 授权许可。开发这个分支的原因之一是：甲骨文公司收购了 MySQL 后，有将 MySQL 闭源的潜在风险，因此社区采用分支的方式来避开这个风险。

MariaDB的目的是完全兼容MySQL，包括API和命令行，使之能轻松成为MySQL的代替品。

```
yum install mariadb-server mariadb
```

```
systemctl start mariadb  #启动MariaDB
systemctl stop mariadb  #停止MariaDB
systemctl restart mariadb  #重启MariaDB
systemctl enable mariadb  #设置开机启动

vi /etc/my.cnf #添加 [mysqld] character_set_server=utf8 init_connect='SET NAMES utf8'

其他默认配置文件路径： 

配置文件：/etc/my.cnf 日志文件：/var/log//var/log/mysqld.log 服务启动脚本：/usr/lib/systemd/system/mysqld.service socket文件：/var/run/mysqld/mysqld.pid
select version();查看版本
```

```
mysqladmin --version
```

## 3.命令

你可以在 MySQL Client(Mysql客户端) 使用 mysql 命令连接到 MySQL 服务器上，默认情况下 MySQL 服务器的登录密码为空，所以本实例不需要输入密码。

命令如下：

```
[root@host]# mysql

SHOW DATABASES;列出 MySQL 数据库管理系统的数据库列表。
Mysql安装成功后，默认的root用户密码为空，你可以使用以下命令来创建root用户的密码：
mysqladmin -u root password "new_password";
可以通过以下命令来连接到Mysql服务器：
mysql -u root -p
mysql -u ben -p -h myserver -P 9999
```

```
ps -ef | grep mysqld
启动mysql服务器
oot@host# cd /usr/bin
./mysqld_safe &
关闭mysql服务器
root@host# cd /usr/bin
./mysqladmin -u root -p shutdown
以下为添加用户的的实例，用户名为guest，密码为guest123，并授权用户可进行 SELECT, INSERT 和 UPDATE操作权限：
SHOW TABLES:列表
显示指定数据库的所有表，使用该命令前需要使用 use 命令来选择要操作的数据库
SHOW COLUMNS FROM 数据表:
显示数据表的属性，属性类型，主键信息 ，是否为 NULL，默认值等其他信息。
SHOW INDEX FROM 数据表:
显示数据表的详细索引信息，包括PRIMARY KEY（主键）。

SHOW TABLE STATUS LIKE [FROM db_name] [LIKE 'pattern'] \G: 
该命令将输出Mysql数据库管理系统的性能及统计信息。

mysql> SHOW TABLE STATUS  FROM RUNOOB;   # 显示数据库 RUNOOB 中所有表的信息
mysql> SHOW TABLE STATUS from RUNOOB LIKE 'runoob%';     # 表名以runoob开头的表的信息
mysql> SHOW TABLE STATUS from RUNOOB LIKE 'runoob%'\G;   # 加上 \G，查询结果按列打印

mysql> show engines;查看MySQL提供的所有存储引擎

SHOW STATUS，用于显示广泛的服务器状态信息；
SHOW CREATE DATABASE和SHOW CREATE TABLE，分别用来显示创建特定数据库或表的MySQL语句；
SHOW GRANTS，用来显示授予用户（所有用户或特定用户）的安全权限；
SHOW ERRORS和SHOW WARNINGS， 用来显示服务器错误或警告消息



```

## MyISAM和InnoDB区别

MyISAM是MySQL的默认数据库引擎（5.5版之前）。虽然性能极佳，而且提供了大量的特性，包括全文索引、压缩、空间函数等，但MyISAM不支持事务和行级锁，而且最大的缺陷就是崩溃后无法安全恢复。不过，5.5版本之后，MySQL引入了InnoDB（事务性数据库引擎），MySQL 5.5版本后默认的存储引擎为InnoDB。

大多数时候我们使用的都是 InnoDB 存储引擎，但是在某些情况下使用 MyISAM 也是合适的比如读密集的情况下。（如果你不介意 MyISAM 崩溃恢复问题的话）。

**两者的对比：**

1. **是否支持行级锁** : MyISAM 只有表级锁(table-level locking)，而InnoDB 支持行级锁(row-level locking)和表级锁,默认为行级锁。
2. **是否支持事务和崩溃后的安全恢复： MyISAM** 强调的是性能，每次查询具有原子性,其执行速度比InnoDB类型更快，但是不提供事务支持。但是**InnoDB** 提供事务支持事务，外部键等高级数据库功能。 具有事务(commit)、回滚(rollback)和崩溃修复能力(crash recovery capabilities)的事务安全(transaction-safe (ACID compliant))型表。
3. **是否支持外键：** MyISAM不支持，而InnoDB支持。
4. **是否支持MVCC** ：仅 InnoDB 支持。应对高并发事务, MVCC比单纯的加锁更高效;MVCC只在 `READ COMMITTED` 和 `REPEATABLE READ` 两个隔离级别下工作;MVCC可以使用 乐观(optimistic)锁 和 悲观(pessimistic)锁来实现;各数据库中MVCC实现并不统一。推荐阅读：[MySQL-InnoDB-MVCC多版本并发控制](https://segmentfault.com/a/1190000012650596)
5. ......

《MySQL高性能》上面有一句话这样写到:

> 不要轻易相信“MyISAM比InnoDB快”之类的经验之谈，这个结论往往不是绝对的。在很多我们已知场景中，InnoDB的速度都可以让MyISAM望尘莫及，尤其是用到了聚簇索引，或者需要访问的数据都可以放入内存的应用。

一般情况下我们选择 InnoDB 都是没有问题的，但是某些情况下你并不在乎可扩展能力和并发能力，也不需要事务支持，也不在乎崩溃后的安全恢复问题的话，选择MyISAM也是一个不错的选择。但是一般情况下，我们都是需要考虑到这些问题的。 



# MySQL 基本架构

下图是 MySQL  的一个简要架构图，从下图你可以很清晰的看到用户的 SQL 语句在 MySQL 内部是如何执行的。

先简单介绍一下下图涉及的一些组件的基本作用帮助大家理解这幅图，在后面会详细介绍到这些组件的作用。

- **连接器：** 身份认证和权限相关(登录 MySQL 的时候)。
- **查询缓存:**  执行查询语句的时候，会先查询缓存（MySQL 8.0 版本后移除，因为这个功能不太实用）。
- **分析器:**  没有命中缓存的话，SQL 语句就会经过分析器，分析器说白了就是要先看你的 SQL 语句要干嘛，再检查你的 SQL 语句语法是否正确。
- **优化器：**  按照 MySQL 认为最优的方案去执行。
- **执行器:** 执行语句，然后从存储引擎返回数据。

![](https://user-gold-cdn.xitu.io/2019/3/23/169a8bc60a083849?w=950&h=1062&f=jpeg&s=38189)

简单来说 MySQL  主要分为 Server 层和存储引擎层：

- **Server 层**：主要包括连接器、查询缓存、分析器、优化器、执行器等，所有跨存储引擎的功能都在这一层实现，比如存储过程、触发器、视图，函数等，还有一个通用的日志模块 binglog 日志模块。
- **存储引擎**： 主要负责数据的存储和读取，采用可以替换的插件式架构，支持 InnoDB、MyISAM、Memory 等多个存储引擎，其中 InnoDB 引擎有自有的日志模块 redolog 模块。**现在最常用的存储引擎是 InnoDB，它从 MySQL 5.5.5 版本开始就被当做默认存储引擎了。**

## Server 层基本组件介绍

### 1) 连接器

连接器主要和身份认证和权限相关的功能相关，就好比一个级别很高的门卫一样。

主要负责用户登录数据库，进行用户的身份认证，包括校验账户密码，权限等操作，如果用户账户密码已通过，连接器会到权限表中查询该用户的所有权限，之后在这个连接里的权限逻辑判断都是会依赖此时读取到的权限数据，也就是说，后续只要这个连接不断开，即时管理员修改了该用户的权限，该用户也是不受影响的。

### 2) 查询缓存(MySQL 8.0 版本后移除)

查询缓存主要用来缓存我们所执行的 SELECT 语句以及该语句的结果集。

连接建立后，执行查询语句的时候，会先查询缓存，MySQL 会先校验这个 sql 是否执行过，以 Key-Value 的形式缓存在内存中，Key 是查询预计，Value 是结果集。如果缓存 key 被命中，就会直接返回给客户端，如果没有命中，就会执行后续的操作，完成后也会把结果缓存起来，方便下一次调用。当然在真正执行缓存查询的时候还是会校验用户的权限，是否有该表的查询条件。

MySQL 查询不建议使用缓存，因为查询缓存失效在实际业务场景中可能会非常频繁，假如你对一个表更新的话，这个表上的所有的查询缓存都会被清空。对于不经常更新的数据来说，使用缓存还是可以的。

所以，一般在大多数情况下我们都是不推荐去使用查询缓存的。

MySQL 8.0 版本后删除了缓存的功能，官方也是认为该功能在实际的应用场景比较少，所以干脆直接删掉了。

### 3) 分析器

MySQL 没有命中缓存，那么就会进入分析器，分析器主要是用来分析 SQL 语句是来干嘛的，分析器也会分为几步：

**第一步，词法分析**，一条 SQL 语句有多个字符串组成，首先要提取关键字，比如 select，提出查询的表，提出字段名，提出查询条件等等。做完这些操作后，就会进入第二步。

**第二步，语法分析**，主要就是判断你输入的 sql 是否正确，是否符合 MySQL 的语法。

完成这 2 步之后，MySQL 就准备开始执行了，但是如何执行，怎么执行是最好的结果呢？这个时候就需要优化器上场了。

### 4) 优化器 

优化器的作用就是它认为的最优的执行方案去执行（有时候可能也不是最优，这篇文章涉及对这部分知识的深入讲解），比如多个索引的时候该如何选择索引，多表查询的时候如何选择关联顺序等。

可以说，经过了优化器之后可以说这个语句具体该如何执行就已经定下来。

### 5) 执行器

当选择了执行方案后，MySQL 就准备开始执行了，首先执行前会校验该用户有没有权限，如果没有权限，就会返回错误信息，如果有权限，就会去调用引擎的接口，返回接口执行的结果。





## MySQL发展历史

应用程序都可以按照统一的方式直接操作数据，也就是应用程序和数据都具有了高度的独立性。

![img](https://images2015.cnblogs.com/blog/63651/201703/63651-20170308091123641-636480917.jpg)



## 常见数据库技术品牌、服务与架构

发展了这么多年市场上出现了许多的数据库系统，最强的个人认为是Oracle，当然还有许多如：DB2、Microsoft SQL Server、MySQL、SyBase等，下图列出常见数据库技术品牌、服务与架构。

![img](https://images2015.cnblogs.com/blog/63651/201703/63651-20170308091629219-1179178843.jpg)



## MySQL引擎比较

MySQL 的存储引擎接口定义良好。有兴趣的开发者可以通过阅读文档编写自己的存储引擎。

![img](https://images2015.cnblogs.com/blog/63651/201703/63651-20170308122741453-767743951.png)

## MySQL操作实例

[一个小时学会MySQL数据库](https://www.cnblogs.com/best/p/6517755.html) 

