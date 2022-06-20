# 介绍

管道，既可在程序中使用，也可从shell中使用。管道的问题在于它们只能在具有共同祖先（指父子进程关系）的进程间使用，不过该问题已随有名管道（named pipe）即FIFO的引入而解决了。



System V消息队列（System V message queue）是在20世纪80年代早期加到System V内核中的。它们可用在同一主机上有亲缘关系或无亲缘关系的进程之间。当今多数版本的Unix却不论自己是否源自System V都支持它们

考虑到IPC，父进程可以在调用fork前建立某种形式的IPC（例如管道或消息队列），因为它知道随后派生的两个子进程将穿越fork继承该IPC对象。

Posix消息队列是由Posix实时标准加入的，它们可用在同一主机上有亲缘关系和无亲缘关系的进程之间。

互斥锁（mutex）和条件变量（condition variable）是由Posix线程标准（1003.1c-1995）定义的两种同步形式。尽管往往用于线程间的同步，它们也能提供不同进程间的同步。



Unix进程间共享信息的三种方式

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E8%BF%9B%E7%A8%8B%E5%85%B1%E4%BA%AB%E6%96%B9%E5%BC%8F.png)

(1)左边的两个进程共享存留于文件系统中某个文件上的某些信息。为访问这些信息，每个进程都得穿越内核（例如read、write、lseek等）。当一个文件有待更新时，某种形式的同步是必要的，这样既可保护多个写入者，防止相互串扰，也可保护一个或多个读出者，防止写入者的干扰。

(2)中间的两个进程共享驻留于内核中的某些信息。管道是这种共享类型的一个例子， System V消息队列和System V信号量也是。现在访问共享信息的每次操作涉及对内核的一次系统调用。

(3)右边的两个进程有一个双方都能访问的共享内存区。每个进程一旦设置好该共享内存区，就能根本不涉及内核而访问其中的数据。共享该内存区的进程需要某种形式的同步。

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/IPC%E5%AF%B9%E8%B1%A1%E6%8C%81%E7%BB%AD%E6%80%A7.png)

(1)随进程持续的（process-persistent）IPC对象一直存在到打开着该对象的最后一个进程关闭该对象为止。例如管道和FIFO就是这种对象。

(2)随内核持续的（kernel-persistent）IPC对象一直存在到内核重新自举或显式删除该对象为止。例如System V的消息队列、信号量和共享内存区就是此类对象。Posix的消息队列、信号量和共享内存区必须至少是随内核持续的，但也可以是随文件系统持续的，具体取决于实现。

(3)随文件系统持续的（filesystem-persistent）IPC对象一直存在到显式删除该对象为止。即使内核重新自举了，该对象还是保持其值。Posix消息队列、信号量和共享内存区如果是使用映射文件实现的（不是必需条件），那么它们就是随文件系统持续的。

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E5%90%84%E7%B1%BB%E5%9E%8BIPC%E6%8C%81%E7%BB%AD%E6%80%A7.png)

对于一种给定的IPC类型，其可能的名字的集合称为它的名字空间（name space）。名字空间非常重要，因为对于除普通管道以外的所有形式的IPC来说，名字是客户与服务器彼此连接以交换消息的手段。

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E5%90%84IPC%E5%90%8D%E5%AD%97%E7%A9%BA%E9%97%B4.png)

调用fork、exec和_exit对于IPC的影响

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/fork-exec%E5%AF%B9IPC%E5%BD%B1%E5%93%8D.png)



使用Posix IPC函数是大势所趋，因为它们比System V中的相应部分更具优势。

使用Posix线程环境（称为“Pthreads”）

## Posix IPC

Posix消息队列、信号量、共享内存

### 所有Posix IPC函数

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/PosixIPC%E5%87%BD%E6%95%B0%E6%B1%87%E6%80%BB.png)

IPC名字，前提是/tmp目录存在，而且我们在该目录中有写权限，对于多数Unix系统来说，这是正常情况，在Solaris下则失败。

为避免这些移植性问题，我们应该把Posix IPC名字的#define行放在一个便于修改的头文件中，这样应用程序转移到另一个系统上时，只需修改这个头文件。

### 创建与打开IPC通道

打开或创建Posix IPC对象所用的各种oflag常值

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/PosixIPC-oflag%E5%80%BC.png)

前3行指定怎样打开对象：只读、只写或读写。消息队列能以其中任何一种模式打开，信号量的打开不指定任何模式（任意信号量操作，都需要读写访问权），共享内存区对象则不能以只写模式打开。余下4行标志是可选的

创建一个新的消息队列、信号量或共享内存区对象时，至少需要另外一个称为mode的参数。该参数指定权限位，常值定义在<sys/stat.h>头文件中。

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/IPC-mode.png)

### 打开或创建一个IPC对象的逻辑

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E6%89%93%E5%BC%80IPC%E9%80%BB%E8%BE%91.png)

三种类型的Posix IPC——消息队列、信号量、共享内存区——都是用路径名标识的。但是这些路径名既可以是文件系统中的实际路径名，也可以不是，而这点不一致性会导致一个移植性问题。全书采用的解决办法是使用我们自己的px_ipc_name函数。

## System V IPC

System V消息队列、信号量、共享内存

### 所有System V IPC函数

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/systemVIPC%E5%87%BD%E6%95%B0.png)

key_t键和ftok函数

客户和服务器一旦在pathname和id上达成一致，双方就都能调用ftok函数把pathname和id转换成同一个IPC键。

ftok的典型实现调用stat函数，然后组合以下三个值。

(1)pathname所在的文件系统的信息（stat结构的st_dev成员）。

(2)该文件在本文件系统内的索引节点号（stat结构的st_ino成员）。

(3)id的低序8位（不能为0）。

### 创建与打开IPC通道

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/IPC%E9%94%AE%E7%94%9F%E6%88%90IPC%E6%A0%87%E8%AF%86.png)

### 创建或打开一个IPC对象的逻辑

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/SystemIPC%E5%88%9B%E5%BB%BA%E9%80%BB%E8%BE%91.png)

### ipcs和ipcrm

由于System V IPC的三种类型不是以文件系统中的路径名标识的，因此使用标准的ls和rm程序无法看到它们，也无法删除它们。不过实现了这些类型IPC的任何系统都提供两个特殊的程序：ipcs和ipcrm。ipcs输出有关System V IPC特性的各种信息，ipcrm则删除一个System V消息队列、信号量集或共享内存区。

### 内核限制

Solaris 2.6有20个这些限制。它们的当前值可使用sysdef命令输出，不过如果相应的内核模块尚未加载（也就是说尚未使用IPC机制），那么所输出的值为0。它们的值可通过在/etc/system文件中加入如下语句来修改，而/etc/system是自举内核时读入的。

至于Digital Unix 4.0B，sysconfig程序可用于查询或修改许多内核参数和限制。下面是使用-q选项时该程序的输出，它就ipc子系统查询内核以输出当前限制值。我们已省略掉了与System V IPC机制无关的一些行。

/sbin/sysconfig -q ipc

这些参数的默认值可通过在/etc/sysconfigtab文件中指定不同的值来修改，不过该文件应使用sysconfigdb程序维护。该文件是在系统自举时读入的。

> 使用System V IPC的最大问题在于多数实现在这些对象的大小上施加了人为的内核限制，这些限制可追溯到它们历史上的最初实现。这就是说，较多使用System V IPC的多数应用需要系统管理员修改这些内核限制

# 消息传递

管道 FIFO 消息队列

## 管道和FIFO

管道和FIFO都是使用通常的read和write函数访问的。

\#include <unistd.h>

int pipe(int fd[2]);

返回：若成功则为0，若出错则为-1

\#include <stdio.h>

FILE *popen(const char *command,const char *type);

返回：若成功则为文件指针，若出错则为NULL

int pclose(FILE *stream);

返回：若成功则为shell的终止状态，若出错则为-1

\#include <sys/types.h>

\#include <sys/stat.h>

int mkfifo(const char *pathname,mode_t mode);

返回：若成功则为0，若出错则为-1

其中pathname是一个普通的Unix路径名，它是该FIFO的名字。

mode参数指定文件权限位，类似于open的第二个参数，定义在<sys/stat.h>头文件中的6个常值

### O_NONBLOCK标志对管道和FIFO的影响

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/NOBLOCK%E5%AF%B9FIFO%E7%AE%A1%E9%81%93%E5%BD%B1%E5%93%8D.png)

拒绝服务型攻击

为避免这种攻击，在编写任何服务器程序的迭代部分时必须小心，要留意服务器可能在哪儿阻塞以及可能阻塞多久。处理这种问题的方法之一是在特定操作上设置一个超时时钟，但是把服务器程序编写成并发服务器而不是迭代服务器通常更为简单，这么一来，上述类型的拒绝服务型攻击只影响一个子进程，而不会影响主服务器。即使采用并发服务器，拒绝服务型攻击仍可能发生：一个恶意的客户可能发送大量的独立请求，导致服务器达到它的子进程数限制，从而使得后续的fork失败。

### 管道和FIFO限制

系统加于管道和FIFO的唯一限制为：

OPEN_MAX 一个进程在任意时刻打开的最大描述符数（Posix要求至少为16）；

PIPE_BUF（<limits.h>） 可原子地写往一个管道或FIFO的最大数据量（Posix要求至少为512）。

OPEN_MAX的值可通过调用sysconf函数查询。它通常可通过执行ulimit命令（Bourne shell或KornShell，我们马上会看到）或limit命令（C shell）从shell中修改。它也可通过调用setrlimit函数从一个进程中修改。

ulimit -nS　　　　　　　　　　输出最大描述符数，软限制64

ulimit -nH　　　　　　　　　　输出最大描述符数，硬限制1024

ulimit –ns 512　　　　　　　 设置软限制为512



管道和FIFO是许多应用程序的基本构建模块。管道普遍用于shell中，不过也可以从程序中使用，往往是用于从子进程向父进程回传信息。使用管道时涉及的某些代码（pipe、fork、close、exec和waitpid）可通过使用popen和pclose来避免，由它们处理具体细节并激活一个shell。

FIFO与管道类似，但它们是用mkfifo创建的，之后需用open打开。打开管道时必须小心，因为有许多规则（图4-21）制约着open的阻塞与否。

## Posix消息队列

### Posix和System V消息队列主要的差别

- 对Posix消息队列的读总是返回最高优先级的最早消息，对System V消息队列的读则可以返回任意指定优先级的消息。（System V msgrcv可以就返回哪一个消息指定三种不同的情形：所指定队列中最早的消息、具有某个特定类型的最早消息、其类型小于或等于某个值的最早消息）

- 当往一个空队列放置一个消息时，Posix消息队列允许产生一个信号或启动一个线程， System V消息队列则不提供类似机制。

Posix消息队列和System V消息队列都不具备的一个特性是：

- 向接收者准确地标识每条消息的发送者。这个信息在许多应用中可能有用。不幸的是，大多数IPC消息机制并不标识发送者。门15.5

队列中的每个消息具有如下属性：

- 一个无符号整数优先级（Posix）或一个长整数类型（System V）；

- 消息的数据部分长度（可以为0）；

- 数据本身（如果长度大于0）。

注意这些特征不同于管道和FIFO。后两者是字节流模型，没有消息边界，也没有与每个消息关联的类型。

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/Posix%E6%B6%88%E6%81%AF%E9%98%9F%E5%88%97%E5%B8%83%E5%B1%80.png)

### mq_open mq_close mq_unlink

mq_open函数创建一个新的消息队列或打开一个已存在的消息队列

```c
#include <mqueue.h>
mqd_t mq_open(const char *name,int oflag,... // 返回：若成功则为消息队列描述符，若出错则为-1
/* mode_t mode,struct mq_attr *attr */ );
// oflag参数是O_RDONLY、O_WRONLY或O_RDWR之一，可能按位或上O_CREAT、O_EXCL或O_NONBLOCK

int mq_close(mqd_t mqdes); //返回：若成功则为0，若出错则为-1
int mq_unlink(const char *name); // 返回：若成功则为0，若出错则为-1 引用计数达到0以删除该队列为止
```

调用进程可以不再使用该描述符，但其消息队列并不从系统中删除。一个进程终止时，它的所有打开着的消息队列都关闭，就像调用了mq_close一样。

要从系统中删除用作mq_open第一个参数的某个name，必须调用mq_unlink

### mq_getattr mq_setattr

每个消息队列有四个属性，mq_getattr返回所有这些属性，mq_setattr则设置其中某个属性。

```c
#include <mqueue.h>
int mq_getattr(mqd_t mqdes,struct mq_attr *attr);
int mq_setattr(mqd_t mqdes,const struct mq_attr *attr,struct mq_attr *oattr);
// 均返回：若成功则为0，若出错则为-1

struct mq_attr {
long mq_flags;　 /* message queue flag: 0,O_NONBLOCK */
long mq_maxmsg;　/* max number of messages allowed on queue */
long mq_msgsize; /* max size of a message (in bytes)*/
long mq_curmsgs; /* number of messages currently on queue */
};
```

### mq_send mq_receive mq_notify

```c
#include <mqueue.h>
int mq_send(mqd_t mqdes,const char *ptr,size_t len,unsigned int prio);
// 返回：若成功则为0，若出错则为−1
ssize_t mq_receive(mqd_t mqdes,char *ptr,size_t len,unsigned int *priop);
// 返回：若成功则为消息中字节数，若出错则为−1
// 这两个函数的前三个参数分别与write和read的前三个参数类似。

int mq_notify(mqd_t mqdes,const struct sigevent *notification);
// 返回：若成功则为0，若出错则为−1
```

这两个函数分别用于往一个队列中放置一个消息和从一个队列中取走一个消息。每个消息有一个优先级，它是一个小于MQ_PRIO_MAX的无符号整数。

mq_receive总是返回所指定队列中最高优先级的最早消息，而且该优先级能随该消息的内容及其长度一同返回。

mq_receive的操作不同于Systme V的msgrcv（6.4节）。System V消息有一个类似于优先级的类型字段，但使用msgrcv时，我们可以就返回哪一个消息指定三种不同的情形：所指定队列中最早的消息、具有某个特定类型的最早消息、其类型小于或等于某个值的最早消息。

mq_notify：为指定队列建立或删除异步事件通知

异步信号安全的函数

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E5%BC%82%E6%AD%A5%E4%BF%A1%E5%8F%B7%E5%AE%89%E5%85%A8%E7%9A%84%E5%87%BD%E6%95%B0.png)

使用内存映射I/O以及Posix互斥锁和条件变量完成的Posix消息队列的实现

创建出的消息队列最多容纳4个消息，每个消息7个字节。

使用内存映射文件实现Posix消息队列的各种数据结构的布局

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E5%86%85%E5%AD%98%E6%98%A0%E5%B0%84%E6%96%87%E4%BB%B6%E5%AE%9E%E7%8E%B0Posix%E6%B6%88%E6%81%AF%E9%98%9F%E5%88%97%E5%B8%83%E5%B1%80.png)

### 消息队列限制

mq_mqxmsg　　 队列中的最大消息数；

mq_msgsize　　给定消息的最大字节数。

MQ_OPEN_MAX 一个进程能够同时拥有的打开着消息队列的最大数目（Posix要求它至少为8）；

MQ_PRIO_MAX 任意消息的最大优先级值加1（Posix要求它至少为32）。

这两个常值往往定义在<unistd.h>头文件中，也可以在运行时通过调用sysconf函数获取

### 总结

mq_open创建一个新队列或打开一个已存在的队列，mq_close关闭队列，mq_unlink则删除队列名。往一个队列中放置消息使用mq_send，从一个队列中读出消息使用mq_receive。队列属性的查询与设置使用mq_getattr和mq_setattr，函数mq_notify则允许我们注册一个信号或线程，它们在有一个消息被放置到某个空队列上时发送（信号）或激活（线程）。队列中的每个消息被赋予一个小整数优先级，mq_receive每次被调用时总是返回最高优先级的最早消息。

rnq_notify的使用给我们引入了Posix实时信号，它们在SIGRTMIN和SIGRTMAX之间。当设置SA_SIGINFO标志来安装这些信号的处理程序时，（1）这些信号是排队的，（2）排了队的信号是以FIFO顺序递交的，（3）给信号处理程序传递两个额外的参数。

## System V消息队列

跟Posix消息队列一样，在某个进程往一个队列中写入一个消息之前，不求另外某个进程正在等待该队列上一个消息的到达。

```c
struct msqid_ds {
struct ipc_perm　 msg_perm;　　/* read_write perms: Section 3.3 */
struct msg　　　　*msg_first;　 /* ptr to first message on queue */
struct msg　　　　*msg_last;　　/* ptr to last message on queue */

msglen_t　　　　　　msg_cbytes;　/* current # bytes on queue */
msgqnum_t　　　　　msg_qnum;　　/* current # of messages on queue */
msglen_t　　　　　　msg_qbytes;　/* max # of bytes allowed on queue */

pid_t　　　　　　　 msg_lspid;　 /* pid of last msgsnd()*/
pid_t　　　　　　　 msg_lrpid;　 /* pid of last msgrcv()*/
    
time_t　　　　　　　msg_stime;　 /* time of last msgsnd()*/
time_t　　　　　　　msg_rtime;　 /* time of last msgrcv()*/
time_t　　　　　　　msg_ctime;　 /* time of last msgctl()
(that changed the above)*/
};
```

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E5%86%85%E6%A0%B8%E4%B8%AD%E7%9A%84System%20V%E6%B6%88%E6%81%AF%E9%98%9F%E5%88%97%E7%BB%93%E6%9E%84.png)

msgget函数用于创建一个新的消息队列或访问一个已存在的消息队列

使用msgget打开一个消息队列后，我们使用msgsnd往其上放置一个消息

使用msgrcv函数从某个消息队列中读出一个消息

msgctl函数提供在一个消息队列上的各种控制操作

System V消息队列的问题之一是它们由各自的标识符而不是描述符标识。这意味着我们不能在消息队列上直接使用select或poll

### 消息队列限制

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/System%20V%E6%B6%88%E6%81%AF%E9%98%9F%E5%88%97%E7%9A%84%E5%85%B8%E5%9E%8B%E7%B3%BB%E7%BB%9F%E9%99%90%E5%88%B6.png)

System V消息队列与Posix消息队列类似。新的应用程序应考虑使用Posix消息队列，不过大量的现有代码使用System V消息队列。然而把一个应用程序从使用System V消息队列重新编写成使用Posix消息队列不是件难事。Posix消息队列缺失的主要特性是从队列中读出指定优先级的消息的能力。这两种消息队列都不使用真正的描述符，从而造成在消息队列上使用select或poll的困难。

# 同步-(线程同步)

互斥量 条件变量 读写锁 文件和记录锁 信号量

《线程同步》

使用fcntl记录上锁。然而你有可能碰到使用这些老式上锁技巧的代码，它们通常存在于fcntl上锁尚未广泛得以实现之前编写的程序中。

fcntl记录上锁提供了对一个文件的劝告性或强制性上锁功能，而我们是通过该文件打开着的描述符来访问它的。这些锁用于不同进程间的上锁，而不是同一进程内不同线程间的上锁。术语“记录”是个不确切的名字，因为Unix内核没有文件内记录的概念。更好的称谓是“范围上锁（range locking）”。

NFS就是网络文件系统，它在TCPvl第29章中讨论。作为对NFS的一种扩展，NFS的大多数实现支持fcntl记录上锁。Unix系统通常以两个额外的守护进程支持NFS记录上锁，它们是lockd和statd。当某个进程调用fcntl以获取一个锁，而且内核检测出其描述符引用通过NFS安装的某个文件系统上的一个文件时，本地的lockd就向服务器的lockd发送这个请求。statd守护进程跟踪着持有锁的各个客户，它与lockd交互以提供NFS上锁的崩溃恢复功能。

## 信号量

Posix信号量

Posix定义了两个信号量限制：

SEM_NSEMS_MAX 一个进程可同时打开着的最大信号量数（Posix要求至少为256）；

SEM_VALUE_MAX 一个信号量的最大值（Posix要求至少为32767）。

这两个常值通常定义在<unistd.h>头文件中，也可在运行时通过调用sysconf函数获取

System V信号量限制

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/System%20V%E4%BF%A1%E5%8F%B7%E9%87%8F%E7%9A%84%E5%85%B8%E5%9E%8B%E9%99%90%E5%88%B6%E5%80%BC.png) 

# 共享内存

匿名 具名

## mmap munmap msync

mmap、munmap（删除映射关系）、msync（MS_ASYNC、MS_SYNC、MS_INVALIDATE）

mmap函数把一个文件或一个Posix共享内存区对象映射到调用进程的地址空间。MS_INVALIDATE使用该函数有三个目的：

(1)使用普通文件以提供内存映射I/O

(2)使用特殊文件以提供匿名内存映射

(3)使用shm_open以提供无亲缘关系进程间的Posix共享内存区

内存映射文件

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E5%86%85%E5%AD%98%E6%98%A0%E5%B0%84%E6%96%87%E4%BB%B6.png)

mmap大小等于文件大小时的内存映射

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/mmap%E5%A4%A7%E5%B0%8F%E7%AD%89%E4%BA%8E%E6%96%87%E4%BB%B6%E5%A4%A7%E5%B0%8F%E6%97%B6%E7%9A%84%E5%86%85%E5%AD%98%E6%98%A0%E5%B0%84.png) 

mmap大小超过文件大小时的内存映射，SIGBUS信号（Bus Error）、SIGSEGV信号（Segmentation Fault）

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/mmap%E5%A4%A7%E5%B0%8F%E8%B6%85%E8%BF%87%E6%96%87%E4%BB%B6%E5%A4%A7%E5%B0%8F%E6%97%B6%E7%9A%84%E5%86%85%E5%AD%98%E6%98%A0%E5%B0%84.png)

## Poxis共享内存

Posix共享内存区调用shm_open后调用mmap的是，System V共享内存先调用shmget，再调用shmat。

```c
#include <sys/mman.h>
int shm_open(const char *name,int oflag,mode_t mode);
int shm_unlink(const char *name);
// 返回：若成功则为非负描述符，若出错则为−1
// 返回：若成功则为0，若出错则为−1

#include <unistd.h>
int ftruncate(int fd,off_t length);
// 返回：若成功则为0，若出错则为−1

#include <sys/types.h>
#include <sys/stat.h>
int fstat(int fd,struct stat *buf);
// 返回：若成功则为0，若出错则为−1
```

# 远程过程调用

Solaris门 Sun RPC

几乎所有Unix系统都可使用Sun RPC，门则是Solaris特有的特性（到目前为止是这样）。

## 门

三种不同类型的过程调用

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E4%B8%89%E7%A7%8D%E4%B8%8D%E5%90%8C%E7%B1%BB%E5%9E%8B%E7%9A%84%E8%BF%87%E7%A8%8B%E8%B0%83%E7%94%A8.png)

与所有其他形式的消息传递机制相比较，门不是更快也至少一样快。

远程过程调用可能是同步的，也可能是异步的，不过我们将看到门调用是同步的。

从一个进程到另一个进程的表象过程调用

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E4%BB%8E%E4%B8%80%E4%B8%AA%E8%BF%9B%E7%A8%8B%E5%88%B0%E5%8F%A6%E4%B8%80%E4%B8%AA%E8%BF%9B%E7%A8%8B%E7%9A%84%E8%A1%A8%E8%B1%A1%E8%BF%87%E7%A8%8B%E8%B0%83%E7%94%A8.png) 

从一个进程到另一个进程的过程调用的真正控制流

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E4%BB%8E%E4%B8%80%E4%B8%AA%E8%BF%9B%E7%A8%8B%E5%88%B0%E5%8F%A6%E4%B8%80%E4%B8%AA%E8%BF%9B%E7%A8%8B%E7%9A%84%E8%BF%87%E7%A8%8B%E8%B0%83%E7%94%A8%E7%9A%84%E7%9C%9F%E6%AD%A3%E6%8E%A7%E5%88%B6%E6%B5%81.png)

### door_call

door_call函数由客户调用，它会调用在服务器进程的地址空间中执行的一个服务器过程。

\#include <door.h>

int door_call(int fd,door_arg_t *argp);

返回：若成功则为0，若出错则为−1

### door_create

服务器进程通过调用door_create建立一个服务器过程。

\#include <door.h>

typedef void Door_server_proc(void *cookie,char *dataptr,size_t datasize,

door_desc_t *descptr,size_t ndesc);

int door_create(Door_server_proc *proc,void *cookie,u_int attr);

返回：若成功则为非负描述符，若出错则为−1

### door_return

服务器过程完成工作时通过调用door_return返回。这会使客户中相关联的door_call调用返回。

\#include <door.h>

int door_return(char *dataptr,size_t datasize,door_desc_t *descptr,size_t *ndesc);

返回：若成功则不返回到调用者，若出错则为−1

数据结果由dataptr和datasize指定，描述符结果由descptr和ndesc指定。

### door_cred

门有一个很不错的特性：服务器过程能够获取每个调用对应的客户凭证。这是由door_cred函数完成的。

\#include <door.h>

int door_cred(door_cred_t *cred);

返回：若成功则为0，若出错则为−1

其中由cred指向的door_cred_t结构将在返回时填入客户的凭证。

### door_info

客户也可通过调用door_info函数找出有关服务器的信息。

\#include <door.h>

int door_info(int fd,door_info_t *info);

返回：若成功则为0，若出错则为−1

其中fd指定一个已打开的门。由info指向的door_info_t结构将在返回时填入关于服务器的信息。

> 描述door_call函数时我们提到过，如果结果缓冲区太小而容纳不了服务器的结果，门函数库就会自动分配一个新的缓冲区。
>
> 通过调用mmap分配的这个新缓冲区可使用munmap返还给系统。客户也可以给后续的door_call调用一直使用该缓冲区。

### door_server_create

如果我们想要改变上述任何特性，或者希望亲自管理服务器线程池，那么开源调用door_server_create以指定我们自己的服务器创建过程（sever creation procedure）。

\#include <door.h>

typedef void Door_create_proc(door_info_t *);

Door_create_proc *door_server_create(Door_create_proc *proc);

返回：指向先前的服务器创建过程的指针

### door_bind door_unbind door_revoke

\#include <door.h>

int door_bind(int fd);

int door_unbind(void);

int door_revoke(int fd);

均返回：若成功则为0，若出错则为−1

door_bind函数。它把调用线程捆绑到与描述符为fd的门关联的私用服务器池中。如果调用线程已绑定在另外某个门上，那就执行一个隐式的松绑操作。

door_unbind显式地把调用线程从其已绑定的门上松绑。

door_revoke撤销对于由fd标识的门的访问。一个门描述符只能由创建它的进程撤销。调用该函数时已在进展中的任何门激活实例仍允许正常地完成。

## Sun RPC

构建一个RPC客户-服务器程序所需的步骤汇总

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E6%9E%84%E5%BB%BA%E4%B8%80%E4%B8%AARPC%E5%AE%A2%E6%88%B7-%E6%9C%8D%E5%8A%A1%E5%99%A8%E7%A8%8B%E5%BA%8F%E6%89%80%E9%9C%80%E7%9A%84%E6%AD%A5%E9%AA%A4%E6%B1%87%E6%80%BB.png)

### rpcinfo

通过执行rpcinfo程序，我们可以看到已向端口映射器注册了的所有RPC程序。我们可执行该程序来验证端口映射器本身使用端口号111。

rpcinfo -p

rpcinfo程序显示本系统上当前已注册的程序。一个给定系统上所支持RPC程序的相关信息的另一个来源是目录/usr/inelude/rpcsvc中的.x文件。

默认情况下，由rpcgen创建的服务器可由inetd超级服务器激活。

### 超时重传

查看Sun RPC使用的超时和重传策略。Sun RPC使用了两个超时值。

(1)总超时（total timeout）：一个客户等待其服务器的应答的总时间量。TCP和UDP都使用该值。

(2)重试超时（retry timeout）：只用于UDP，是一个客户在等待其服务器的应答期间每次重传请求的间隔时间。



验证UDP重试超时的唯一办法是使用tcpdump观察分组。这种观察表明，第一个数据报是在客户一启动后就发送的，下一个数据报的发送则在约15秒之后。

### RCP请求和响应

封装在一个TCP分节中的RPC请求

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E5%B0%81%E8%A3%85%E5%9C%A8%E4%B8%80%E4%B8%AATCP%E5%88%86%E8%8A%82%E4%B8%AD%E7%9A%84RPC%E8%AF%B7%E6%B1%82.png)

可能的RPC应答

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E5%8F%AF%E8%83%BD%E7%9A%84RPC%E5%BA%94%E7%AD%94.png)

封装为一个UDP数据报的成功的RPC应答

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E5%B0%81%E8%A3%85%E4%B8%BA%E4%B8%80%E4%B8%AAUDP%E6%95%B0%E6%8D%AE%E6%8A%A5%E7%9A%84%E6%88%90%E5%8A%9F%E7%9A%84RPC%E5%BA%94%E7%AD%94.png)

XDR运行时系统也是构建使用RPC的应用程序过程中的一个基本部件。XDR定义了在不同的系统间交换各种数据格式的一种标准方法，这些系统可能具有不同的整数大小、不同的字节序、不同的浮点数格式等。正如我们所示的那样，XDR可独立于RPC软件包单独使用，其目的纯粹是为了以一种标准的格式交换数据，而数据的交换可以使用任意形式的真正传送数据的通信手段（例如使用套接字或XTI编写的程序、软盘、CD-ROM等）。

Sun RPC提供了另外三种认证形式：Unix认证（提供客户的主机名、用户ID和组ID）、DES认证（基于私钥和公钥加密技术）和Kerberos认证。

# 总结

管道和FIFO是字节流，没有消息边界。Posix消息和System V消息则有从发送者向接收者维护的记录边界。（考虑到UNPv1中讲述的网际协议族，TCP是没有记录边界的字节流， UDP则提供具有记录边界的消息。）

在众多的同步技术——互斥锁、条件变量、读写锁、记录锁、Posix信号量和System V信号量——中，可从信号处理程序中调用的函数（图5-10）只有sem_post和fcntl。

在众多的消息传递技术——管道、FIFO、Posix消息队列和System V消息队列——中，可从一个信号处理程序中调用的函数只有read和write（适用于管道和FIFO）。