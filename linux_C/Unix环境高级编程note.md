|         进程          |         线程          |                          进程间通信                          |
| :-------------------: | :-------------------: | :----------------------------------------------------------: |
| [进程介绍](#进程介绍) | [线程介绍](#线程介绍) |         [管道](#管道) [FIFO有名管道](#FIFO有名管道)          |
| [守护进程](#守护进程) | [线程同步](#线程同步) | [XSI-IPC](#XSI-IPC)  ：[消息队列](#消息队列) [信号量](#信号量) [共享存储](#共享存储) |
|                       | [线程控制](#线程控制) |        [POSIX信号量](#POSIX信号量) [套接字](#套接字)         |

- [进程](#进程) 
  - [进程介绍](#进程介绍)
  - [守护进程](#守护进程)

- [线程](#线程)
  - [线程介绍](#线程介绍)
    - [pthread_create](#pthread_create)
    - [pthread_exit](#pthread_exit)
    - [pthread_join](#pthread_join)
    - [pthread_cancel](#pthread_cancel)
    - [线程清理](#线程清理)
    - [pthread_detach](#pthread_detach)
  - [线程同步](#线程同步)
    - [互斥量mutex互斥锁](#互斥量mutex互斥锁) 
    - [读写锁rwlock](#读写锁rwlock)
    - [条件变量cond](#条件变量cond)
    - [自旋锁spin](#自旋锁spin)
    - [barrier](#barrier)
  - [线程控制](#线程控制)
    - [线程属性](#线程属性)
    - [同步属性](#同步属性)
      - [互斥锁属性](#互斥锁属性)
      - [读写锁属性](#读写锁属性)
      - [条件变量属性](#条件变量属性)
      - [barrier属性](#barrier属性)
    - [线程特定数据](#线程特定数据)
    - [线程和信号](#线程和信号)
    - [线程和fork](#线程和fork)
- [进程间通信IPC](#进程间通信IPC)
  - [管道](#管道)
    - [创建管道](#创建管道)
    - [读写管道规则](#读写管道规则)
    - [popen/pclose](#popen/pclose)
  - [FIFO有名管道](#FIFO有名管道)
    - [创建FIFO](#创建FIFO)
    - [打开FIFO](#打开FIFO)
    - [读写FIFO](#读写FIFO)
  - [XSI-IPC](#XSI-IPC)
    - [XSI-IPC介绍](#XSI-IPC介绍) 
      - [标识符和键](#标识符和键)
      - [权限和结构](#权限和结构)
      - [优缺点](#优缺点)
    - [消息队列](#消息队列)
      - [创建/打开消息队列](#创建/打开消息队列)
      - [添加到队列](#添加到队列)
      - [获取消息](#获取消息)
      - [操作消息队列](#操作消息队列)
    - [信号量](#信号量)
    - [共享存储](#共享存储)
  - [POSIX信号量](#POSIX信号量)
    - [创建/获取信号量](#创建/获取信号量)
    - [关闭信号量](#关闭信号量)
    - [销毁信号量](#销毁信号量)
    - [信号量操作](#信号量操作)
    - [未命名信号量](#未命名信号量)
  - [套接字](#套接字) 

网页在线阅读：[unix环境高级编程note](https://hellolinux.xyz/posts/e39ad6a7.html) 

# C语言

- void * malloc(size_t n); 未初始化的，memset初始化
- void free(void * p); 只是释放指针指向的内容 野指针，null
- void *calloc(size_t n, size_t size); 初始化为0，适合为数组申请空间
- void * realloc(void * p, size_t n); 将指针 p指向的内存块的大小改变为n字节，可能新地址

# 高级I/O



# 进程

## 进程介绍

fork函数创建子进程，通常情况用来执行另一个程序的，如果是执行同一个进程那么多数采用线程。

vfork 和fork 之间的另一个区别是：vfork 保证子进程先运行，在她调用exec 或exit 之后父进程才可能被调度运行。如果在调用这两个函数之前子进程依赖于父进程的进一步动作，则会导致死锁。

当子进程退出时，系统不会立即删除该子进程的进程描述符，并将进程状态设置为Z（zombie），然后等待父进程处理子进程的退出。如果父进程不处理，那么子进程就会一直处于僵尸状态。

父进程处理子进程使用wait函数族。当使用wait函数处理完子进程的终止状态后，子进程才彻底的消失了，在次调用wait函数处理该子进程将会出错。

在使用wait函数时，一般等待三种进程的退出

- 等待任意一子进程退出
- 等待具体的某个子进程退出
- 等待某个进程组的任意一子进程退出

一般情况下wait函数的目的就是取出子进程的终止状态，但是通过设置一些选项，我们可以控制wait的一些行为，比如非阻塞式的，只查看是否有子进程退出，而不取出该子进程的终止状态等。

wait阻塞、waitpid不阻塞



popen、system函数（调用了fork、exec、waitpid，三种返回值）

system()来执行一个shell命令，popen()也是执行shell命令并且通过管道和shell命令进行通信。

system在执行期间调用进程会一直等待shell命令执行完成(waitpid等待子进程结束)才返回，但是popen无须等待shell命令执行完成就返回了。我们可以理解system为串行执行，在执行期间调用进程放弃了”控制权”，popen为并行执行。如果你没有在调用popen后调用pclose那么这个子进程就可能变成”僵尸”。



## 守护进程

- 一般在系统启动时装入，仅在系统关闭时终止。大多数守护进程以超级用户特权运行。
- 用户层守护进程的父进程是 init 进程。
- 所有的守护进程都没有控制终端，其终端名设置为问号。
  - 内核守护进程以无控制终端方式启动。
  - 用户层守护进程可以通过调用 setsid 实现。



守护进程是没有控制终端的，无法将自己的消息输出到标准输出或标准错误上，需要一个集中的守护进程记录设施，即 syslog。

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/deamon_syslog.png) 

主要有 3 中产生日志消息的方式：

1.内核例程调用 log 函数
2.大多数用户进程调用 syslog 函数
3.将日志消息发送到 UDP 的 514 端口

syslogd 守护进程接收这些日志消息，在其启动前会读取配置文件（/etc/syslog.conf），以决定各类消息的处理方式。

```c
#include <syslog.h>
void openlog(const char *ident, int option, int facility);
void syslog(int priority, const char *format, ...);
void closelog(void);
// Returns: previous log priority mask value
int setlogmask(int maskpri);
```

注：在没有调用 openlog 的情况下，先调用了 syslog，会自动调用 openlog。

ident 参数指向的字符串会被加到日志消息中去，因此一般指定为程序名称。

option 参数指定各种选项的位屏蔽

facility 参数可选值见图

priority 参数包含 facility 和 level 的组合，如果参数中没有指定 facility，则会使用 openlog 中指定的 facility，如果没有调用 openlog，那么会使用默认值 LOG_USER。

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/openlog_option.png) 

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/openlog_facility.png) 

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/syslog_level.png)  

守护进程编程规则

1.调用 umask 函数将文件模式创建屏蔽字设置为指定值（通常为 0）。守护进程可能需要创建一些文件，如果使用继承的屏蔽字，可能文件的权限会不符合预期。

2.调用 fork 后，使父进程 exit。这样可以保证子进程不是进程组的组长进程。

3.调用 setsid 创建新会话。这可以保证当前进程没有控制终端，且成为新会话的首进程和新进程组的组长进程。

4.将当前工作目录改为根目录或某个指定位置。

5.关闭不再需要的文件描述符。可以使用 getrlimit 函数获取最高文件描述符值，并关闭直到该值的所有描述符。

6.某些守护进程将文件描述符 0、1 和 2 指向 /dev/null，这样任何需要输入输出的库例程都不会产生影响。

单实例守护进程

某些守护进程在同一时刻只能运行一个实例程序，这时候可以使用文件和记录锁（简称文件锁）来实现这个功能。

守护进程只要创建一个固定名字的文件（一般在 /var/run 目录中），并在该文件整体上加一把写锁，那么此后其他的守护进程如果想要给该文件加锁就会失败，也就不应该继续运行。在守护进程终止时，锁会被自动删除，简化了复原过程。



# 线程

## 线程介绍

线程包含了表示进程内执行环境的必须信息，其中包含进程中表示线程的线程ID，一组寄存器值，栈，调度优先级和策略。信号屏蔽字，errno值以及线程私有数据。

**进程的所有信息对该进程的所有线程都是共享的，包括可执行的程序文本，程序的全局内存和堆内存，栈以及文件描述符。** 

- 线程标识

就像每一个进程都有一个进程ID一样，每一个线程也有一个线程ID，进程ID在整个系统中是唯一的，但线程ID不同，线程ID只在它所属的进程环境中有效。线程ID用pthread_t数据类型来表示，（Linux使用无符号长整数表示pthread_t结构）。实现的时候可以用一个结构来代表pthread_t数据类型，所以可移植的系统不能把它当做整数来处理，因此必须使用函数来对两个线程ID进行比较。

**线程可以通过pthread_self函数获取自身的线程ID** 

```c
#include<pthread.h>
int pthread_equal(pthread_t tid1,pthread_t tid2); //返回0表示相等
pthread_t pthread_self(void) //返回线程的线程ID
```

### pthread_create

线程创建

```c
pthread_t pthread_create(pthread_t  *tidp,const pthread_attr_t *attr,void *(*start_rtn)(void),void * arg)//若成功返回0，否则返回错误编号
```

新创建的线程的线程 ID 被设置成 tidp 指向的内存单元；attr 参数定制线程的不同属性；start_rtn 函数是线程开始时执行的函数，其参数可以通过 arg 进行传递。

注意：新线程最好不要通过 tidp 指向的内存空间获取自己的线程 ID，因为如果新线程在主线程调用 pthread_create 返回前就运行了，那么它看到的就是未经初始化的内容，很可能并不是正确的线程 ID。可以使用 pthread_self 函数获取自己的线程 ID。

### pthread_exit

- 线程退出

如果进程中某一个线程调用了exit，_exit或者_Exit，那么整个进程就会终止。类似的，如果信号的默认动作是终止进程，那么把该信号发送到某个线程，整个进程都会终止。

单个线程可以有三种方式退出：
1.直接从启动实例中返回，返回值是线程退出码

2.被同一进程的其他线程取消

3.调用 pthread_exit

```c
#include<pthread.h>
void pthread_exit(void *rval_ptr)
```

### pthread_join

```c
#include<pthread.h>
int pthread_join(pthread_t thread,void ** rval_ptr) //成功返回0，否则返回错误码
```

rval_ptr是一个无类型指针，进程中的其他线程可以通过调用pthread_join函数访问到这个指针

调用 pthread_join 的线程会一直阻塞，直到指定的线程终止。如果指定的线程直接返回或者是调用 pthread_exit 终止，则可以通过 rval_ptr 查看其返回值；如果线程是被取消的，则 rval_ptr 被设置为 PTHRERAD_CANCELED。

**如果线程已经处于分离状态，那么，pthread_join调用就会失败。**

**pthread_create和pthread_exit函数的无类型指针参数能传递的数值不止一个，该指针可以传递更复杂信息的结构地址，但是注意这个结构所使用的内存，在调用者完成调用以后必须仍然是有效的，否则就会出现无效或者非法内存访问。例如在调用线程的栈上分配了该结构，那么其他线程在使用这个结构时内存可能就已经改变了。（可以使用全局栈结构malloc调用分配结构）。** 

### pthread_cancel

- 取消线程

```c
// Returns: 0 if OK, error number on failure
int pthread_cancel(pthread_t tid);
```

***线程可以安排退出时它调用的函数，这样的函数成为线程清理处理程序。处理程序记录在栈中，也会就是说他们的执行顺序与注册时相反。*** 

### 线程清理

- 线程清理处理程序

```c
void pthread_cleanup_push(void (*rtn)(void *), void *arg);
void pthread_cleanup_pop(int execute);
```

清理函数 rtn 只有在以下情况会执行：

1.调用 pthread_exit
2.响应取消请求
3.用非零execute 参数调用 pthread_cleanup_pop（为 0 时，清理函数不会被调用）

两个函数需要成对使用。

### pthread_detach

- 线程分离

```c
// Returns: 0 if OK, error number on failure
int pthread_detach(pthread_t tid);
```

默认情况下，线程的终止状态会保留，直到调用 pthread_join。如果线程被分离，则资源会在线程终止后被立即收回。


## 线程同步

### 互斥量mutex互斥锁

```c
// All return: 0 if OK, error number on failure
int pthread_mutex_init(pthread_mutex_t *restrict mutex,const pthread_mutexattr_t *restrict attr);
int pthread_mutex_destroy(pthread_mutex_t *mutex);

int pthread_mutex_lock(pthread_mutex_t *mutex);
int pthread_mutex_trylock(pthread_mutex_t *mutex);
int pthread_mutex_unlock(pthread_mutex_t *mutex);
```

互斥变量为 pthread_mutex_t 类型，如果使用静态分配方式，可以直接使用 PTHREAD_MUTEX_INITIALIZER 进行初始化。对于动态分配的互斥量，在释放内存前需要调用 pthread_mutex_destroy。

- 带有超时的互斥锁

```c
#include <time.h>
// Returns: 0 if OK, error number on failure
int pthread_mutex_timedlock(pthread_mutex_t *restrict mutex,const struct timespec *restrict tsptr);
```

如果不希望线程在访问加锁的互斥量时无限等待，可以通过 pthread_mutex_timedlock 指定等待的绝对时间。

```c
#include <pthread.h>
#include <time.h>

#include "apue.h"

int main()
{
  int err;
  struct timespec tout;
  struct tm *tmp;
  char buf[64];
  pthread_mutex_t lock = PTHREAD_MUTEX_INITIALIZER;

// 加锁
  pthread_mutex_lock(&lock);
  printf("mutex is locked.\n");
  clock_gettime(CLOCK_REALTIME, &tout);
  tmp = localtime(&tout.tv_sec);
  strftime(buf, sizeof(buf), "%r", tmp);
  printf("current time is %s\n", buf);

// 设置超时
  tout.tv_sec += 10;
  err = pthread_mutex_timedlock(&lock, &tout);
  clock_gettime(CLOCK_REALTIME, &tout);
  tmp = localtime(&tout.tv_sec);
  strftime(buf, sizeof(buf), "%r", tmp);
  printf("the time is now %s\n", buf);

  if(err == 0) {
    printf("mutex locked.\n");
  } else {
    printf("can't lock mutex:%s\n",strerror(err));
  }

  return 0;
}
```

### 读写锁rwlock

读写锁有 3 中状态：不加锁、读模式加锁和写模式加锁。一次只有一个线程可以占有写模式的读写锁，但是多个线程可以同时占有读模式的读写锁。

读写锁适合对数据结构读的次数远大于写的情况。

```c
int pthread_rwlock_init(pthread_rwlock_t *restrict rwlock,
                      const pthread_rwlockattr_t *restrict attr);
int pthread_rwlock_destroy(pthread_rwlock_t *rwlock);

// All return: 0 if OK, error number on failure
int pthread_rwlock_rdlock(pthread_rwlock_t *rwlock); // 读模式锁定
int pthread_rwlock_wrlock(pthread_rwlock_t *rwlock); // 写模式锁定
int pthread_rwlock_unlock(pthread_rwlock_t *rwlock);

int pthread_rwlock_tryrdlock(pthread_rwlock_t *rwlock);
int pthread_rwlock_trywrlock(pthread_rwlock_t *rwlock);

// 带有超时读写锁
// Both return: 0 if OK, error number on failure
int pthread_rwlock_timedrdlock(pthread_rwlock_t *restrict rwlock,const struct timespec *restrict tsptr);
int pthread_rwlock_timedwrlock(pthread_rwlock_t *restrict rwlock,const struct timespec *restrict tsptr);
```

### 条件变量cond

当线程等待的条件变量被满足后，该线程就会被唤醒。条件变量需要和互斥量配合使用，条件本身是由互斥量保护的。

在使用条件变量之前，必须对其进行初始化（有静态和动态 2 种方式）。

```c
// All return: 0 if OK, error number on failure
int pthread_cond_init(pthread_cond_t *restrict cond,const pthread_condattr_t *restrict attr);
int pthread_cond_destroy(pthread_cond_t *cond);

int pthread_cond_wait(pthread_cond_t *restrict cond,pthread_mutex_t *restrict mutex);
int pthread_cond_timedwait(pthread_cond_t *restrict cond,pthread_mutex_t *restrict mutex,
                         const struct timespec *restrict tsptr);

int pthread_cond_signal(pthread_cond_t *cond);    // 至少唤醒一个
int pthread_cond_broadcast(pthread_cond_t *cond); // 全部唤醒
```

pthread_cond_wait 操作主要执行如下操作步骤

> 1. 解锁互斥量 mutex
> 2. 阻塞调用线程，直至另一线程就条件变量 cond 发出信号
> 3. 重新锁定 mutex

在使用 pthread_cond_wait 函数之前，应该已经取得 mutex 锁。对 pthread_cond_wait 的调用应该放在 while 循环中，因为从 wait 函数返回时，并不能确定条件已经得到满足（其他线程先醒来、虚假唤醒等），需要重新对条件进行判断。

```c
// 消费者进程
void *process_msg(void *arg)
{
  for (;;) {
    pthread_mutex_lock(&qlock);
    while (count <= 0) {
      printf("%s wait msg\n", tag);
      pthread_cond_wait(&qready, &qlock);
    }
    count--;
    pthread_mutex_unlock(&qlock);
    /* 处理消息 */
    // 放弃cpu，让另一个处理进场有机会得到数据
    sleep(1);
  }
  return NULL;
}
// 生产者进程
int main(void)
{
  for (;;) {
    pthread_mutex_lock(&qlock);
    count += 4;
    pthread_mutex_unlock(&qlock);
    // 测试两种唤醒方式
#if 1
    pthread_cond_broadcast(&qready);
#else
    pthread_cond_signal(&qready);
#endif
    // 保证两个消费者进程都可以有时间处理数据
    sleep(3);
  }
  return 0;
}
```

### 自旋锁spin

自旋锁与互斥量大体类似，主要的不同之处在于自旋锁在获取锁之前会一直忙等。因此，使用自旋锁应该保证持有锁的时间很短。

自旋锁和互斥量的接口类似

```c
// All return: 0 if OK, error number on failure
int pthread_spin_init(pthread_spinlock_t *lock, int pshared);
int pthread_spin_destroy(pthread_spinlock_t *lock);

int pthread_spin_lock(pthread_spinlock_t *lock);
int pthread_spin_trylock(pthread_spinlock_t *lock);
int pthread_spin_unlock(pthread_spinlock_t *lock);
```

pshared 表示进程共享（process-shared）属性，表明自旋锁的获取方式。它仅在支持线程进程共享同步（Thread Process-Shared Synchronization）的平台上有效，当设置为 PTHREAD_PROCESS_SHARED，则只要线程可以访问锁底层内存，即使是不同进程的线程都可以获得锁；而设置为 PTHREAD_PROCESS_PRIVATE 后，只有初始化该锁的进程内部的线程可以访问它。

### barrier

屏障允许多个线程等待，直到所有合作线程满足某个点后，从该点继续执行。主线程可以将某个任务分解多个小任务交给不同的线程，等到所有线程工作完成后，主线程在此基础上继续执行。

使用 8 个线程分解 800 万个数的排序工作，每个线程对其中的 100 万个数排序，最后由主线程将这些结果进行合并。

```c
// Both return: 0 if OK, error number on failure
int pthread_barrier_init(pthread_barrier_t *restrict barrier,
                       const pthread_barrierattr_t *restrict attr,unsigned int count);
// 初始化函数中的 count 参数用于指定所有线程继续运行前，必须到达屏障的线程数。
int pthread_barrier_destroy(pthread_barrier_t *barrier);

// Returns: 0 or PTHREAD_BARRIER_SERIAL_THREAD if OK, error number on failure
int pthread_barrier_wait(pthread_barrier_t *barrier);
```

wait 函数表明当前线程已完成工作，准备等待其他线程。当线程调用该函数后满足屏障计数，那么函数的返回值为 PTHREAD_BARRIER_SERIAL_THREAD，其余线程该函数返回值为 0。这一特点使得可以很容易的将一个线程作为主线程，它可以工作在其他所有线程已完成的工作结果上。

```c
#include <pthread.h>
#include "apue.h"

pthread_barrier_t pb;
pthread_t t1, t2;

void *th1(void *a)
{
  printf("start t1\n");
  sleep(1);
  // 最后一个完成的线程，返回值应该为-1
  int r = pthread_barrier_wait(&pb);
  printf("th1  r:%d\n", r);
  return NULL;
}

void *th2(void *a)
{
  printf("start t2\n");
  int r = pthread_barrier_wait(&pb);
  printf("th2  r:%d\n", r);
  return NULL;
}

int main()
{
  int r;
  pthread_barrier_init(&pb, NULL, 3);

  pthread_create(&t1, NULL, th1, NULL);
  pthread_create(&t2, NULL, th2, NULL);

  r = pthread_barrier_wait(&pb);
  printf("main r:%d\n", r);

  // 等待子进程结束
  pthread_join(t1, NULL);
  pthread_join(t2, NULL);
  return 0;
}
```

## 线程控制

线程属性和同步原语属性，基于进程的系统调用如何与线程进行交互。

管理这些属性的函数大概有以下几类：

1.初始化函数，负责给属性设置为默认值
2.销毁函数，负责释放初始化函数分配的资源
3.获取属性值的函数
4.设置属性值的函数

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/thread_limit.png)

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/thread_limit2.png) 

### 线程属性

```c
// Both return: 0 if OK, error number on failure
// 初始化和销毁
int pthread_attr_init(pthread_attr_t *attr);
int pthread_attr_destroy(pthread_attr_t *attr);
// destroy 函数除了释放资源外，还会用无效的值初始化属性对象，这样当线程创建函数误用该对象时，会返回错误信息

// Both return: 0 if OK, error number on failure
// 分离状态属性
int pthread_attr_getdetachstate(const pthread_attr_t *restrict attr,int *detachstate);
int pthread_attr_setdetachstate(pthread_attr_t *attr, int detachstate);
// 该状态可以设置成 PTHREAD_CREATE_DETACHED 或 PTHREAD_CREATE_JOINABLE，分别表示以分离状态或正常方式启动线程。

// Both return: 0 if OK, error number on failure
// 线程栈的相关属性 stackaddr 参数指定的是栈的最低内存地址
int pthread_attr_getstack(const pthread_attr_t *restrict attr,
                  void **restrict stackaddr,size_t *restrict stacksize);
int pthread_attr_setstack(pthread_attr_t *attr,void *stackaddr, size_t stacksize);

// Both return: 0 if OK, error number on failure
// 如果不想手动设定栈地址，可以通过下面的函数来仅指定栈大小。
int pthread_attr_getstacksize(const pthread_attr_t *restrict attr,
                         size_t *restrict stacksize);
int pthread_attr_setstacksize(pthread_attr_t *attr, size_t stacksize);

// Both return: 0 if OK, error number on failure
// guardsize 控制线程栈末尾之后用以避免栈溢出的扩展内存的大小。当此值设置为 0 或者修改了线程属性 stackaddr 后，系统不会提供警戒缓冲区。
int pthread_attr_getguardsize(const pthread_attr_t *restrict attr,
                         size_t *restrict guardsize);
int pthread_attr_setguardsize(pthread_attr_t *attr, size_t guardsize);
```

### 同步属性

#### 互斥锁属性

```c
// Both return: 0 if OK, error number on failure
// mutex属性
int pthread_mutexattr_init(pthread_mutexattr_t *attr);
int pthread_mutexattr_destroy(pthread_mutexattr_t *attr);

// Both return: 0 if OK, error number on failure
// 进程共享属性（process-shared）
int pthread_mutexattr_getpshared(const pthread_mutexattr_t *
                           restrict attr, int *restrict pshared);
int pthread_mutexattr_setpshared(pthread_mutexattr_t *attr,int pshared);
// 默认情况下，仅相同进程的线程可以访问同一个同步对象（PTHREAD_PROCESS_PRIVATE），但是在某些情况下，需要多个进程访问同一个同步对象，这时候可以将属性设置为 THREAD_PROCESS_SHARED

// All return: 0 if OK, error number on failure
// 健壮属性（robust）
int pthread_mutexattr_getrobust(const pthread_mutexattr_t *restrict attr,
                              int *restrict robust);
int pthread_mutexattr_setrobust(pthread_mutexattr_t *attr, int robust);
int pthread_mutex_consistent(pthread_mutex_t * mutex);
// 当某个线程在终止时没有释放持有的锁，那么当其他线程尝试获取该锁时，会发生问题。如果使用默认的设置（PTHREAD_MUTEX_STALLED），则请求的线程会一直阻塞。可以通过设置为 PTHREAD_MUTEX_ROBUST 解决这个问题，此时 lock 函数的返回值为 EOWNERDEAD
// 如果线程加锁时发现返回值为 EOWNERDEAD，那么在解锁前需要调用 consistent 函数，声明互斥量的一致性（与该互斥量相关的状态在互斥量解锁之前是一致的）。如果没有调用 consistent 函数就解锁，那么互斥量将不再可用，其他线程调用 lock 函数会返回 ENOTRECOVERABLE。

// Both return: 0 if OK, error number on failure
// 类型属性（type） 控制互斥量的锁定特性
int pthread_mutexattr_gettype(const pthread_mutexattr_t *restrict attr,
                            int *restrict type);
int pthread_mutexattr_settype(pthread_mutexattr_t *attr, int type);
```

PTHREAD_MUTEX_NORMAL ：标准互斥量，不进行错误检查或死锁检测。
PTHREAD_MUTEX_ERRORCHECK ：提供错误检查
PTHREAD_MUTEX_RECURSIVE ：允许同一线程在解锁前多次加锁。
PTHREAD_MUTEX_DEFAULT ：提供默认的特性和行为，操作系统可以将其映射为其他类型。

Mutex type	Relock without unlock?	Unlock when not owned?	Unlock when unlocked?
PTHREAD_MUTEX_NORMAL	deadlock	undefined	undefined
PTHREAD_MUTEX_ERRORCHECK	returns error	returns error	returns error
PTHREAD_MUTEX_RECURSIVE	allowed	returns error	returns error
PTHREAD_MUTEX_DEFAULT	undefined	undefined	undefined

#### 读写锁属性

读写锁非常适合于对数据结构读的次数远大于写的情况

```c
// All return: 0 if OK, error number on failure
// 读写锁仅支持进程共享属性
int pthread_rwlockattr_init(pthread_rwlockattr_t *attr);
int pthread_rwlockattr_destroy(pthread_rwlockattr_t *attr);

int pthread_rwlockattr_getpshared(const pthread_rwlockattr_t *restrict attr,
                                int *restrict pshared);
int pthread_rwlockattr_setpshared(pthread_rwlockattr_t *attr,int pshared);

```

#### 条件变量属性

```c
// All return: 0 if OK, error number on failure
// 支持进程共享属性和时钟属性 时钟属性用于控制 pthread_cond_timedwait 函数使用哪个系统时钟
int pthread_condattr_init(pthread_condattr_t *attr);
int pthread_condattr_destroy(pthread_condattr_t *attr);

int pthread_condattr_getpshared(const pthread_condattr_t *restrict attr,
                              int *restrict pshared);
int pthread_condattr_setpshared(pthread_condattr_t *attr, int pshared);

int pthread_condattr_getclock(const pthread_condattr_t *restrict attr,
                            clockid_t *restrict clock_id);
int pthread_condattr_setclock(pthread_condattr_t *attr,clockid_t clock_id);
```

#### barrier属性

```c
// All return: 0 if OK, error number on failure
// 只有进程共享属性
int pthread_barrierattr_init(pthread_barrierattr_t *attr);
int pthread_barrierattr_destroy(pthread_barrierattr_t *attr);

int pthread_barrierattr_getpshared(const pthread_barrierattr_t *restrict attr,
                                 int *restrict pshared);
int pthread_barrierattr_setpshared(pthread_barrierattr_t *attr, int pshared);
```

### 线程特定数据

线程模型促进了进程中数据和属性的共享，但是在部分场景下，我们又希望线程的部分数据可以是私有的。

一个进程中的所有线程都可以访问进程的整个地址空间，因此线程没有办法阻止另一个线程访问它的数据（除非使用寄存器），即使是接下来介绍的线程特定数据（thread-specific data）机制，也不能做到这一点。但是通过这种机制，可以提高线程间的独立性，使得线程不太容易访问到其他线程的线程特定数据。

每个线程通过 **键（key）** 来访问线程特定数据，键在进程中被所有线程使用，每个线程把自己的线程特定数据和键关联起来。这样，通过同一个键，每个线程可以管理与自己关联的数据。

```c
// Both return: 0 if OK, error number on failure
int pthread_key_create(pthread_key_t *keyp, void (*destructor)(void *));
int pthread_key_delete(pthread_key_t key);
```

创建新键时，每个线程的数据地址为空。同时，在创建的时候可以指定一个析构函数，当线程退出时，如果数据地址不为空，则会调用这个析构函数（参数是数据地址）。

所有的线程都可以调用删除函数来取消键与数据之间的关联，但是这不会触发析构函数。

```c
// Returns: thread-specific data value or NULL if no value has been associated with the key
void *pthread_getspecific(pthread_key_t key);
// Returns: 0 if OK, error number on failure
int pthread_setspecific(pthread_key_t key, const void *value);
// 可以通过 get 函数的返回值来确定是否需要调用 set 函数。
```

- 取消选项

有 2 个额外的线程属性并没有包含在上述的 pthread_attr_t 中，它们分别是可取消状态和可取消类型。

```c
// Returns: 0 if OK, error number on failure
// 可以设置成 PTHREAD_CANCEL_ENABLE 或 PTHREAD_CANCEL_DISABLE
int pthread_setcancelstate(int state, int *oldstate);
// set 函数把当前的可取消状态设置为 state，同时将原来的状态通过 oldstate 返回

// 可以使用 pthread_testcancel 函数手动添加取消点
void pthread_testcancel(void);
```

如果将状态设置为 PTHREAD_CANCEL_DISABLE，那么调用 pthread_cancle 函数并不会杀死线程，取消请求会一直处于挂起状态，直到状态被设置为 ENABLE。同理，此时调用 pthread_testcancel 没有任何效果。

```c
// Returns: 0 if OK, error number on failure
// 可以设置成 PTHREAD_CANCEL_DEFERRED 或 PTHREAD_CANCEL_ASYNCHRONOUS
int pthread_setcanceltype(int type, int *oldtype);
```

默认设置为 PTHREAD_CANCEL_DEFERRED，即推迟取消，线程到达取消点之前不会被真正取消。如果设置为 PTHREAD_CANCEL_ASYNCHRONOUS，即异步取消，那么线程可以在任意时间撤销，而不必等待到达取消点

### 线程和信号

每个线程有自己的信号屏蔽字，通过 pthread_sigmask 函数进行设置，参数与 sigprocmask 类似

```c
#include <signal.h>
// Returns: 0 if OK, error number on failure
// 需要注意的是，如果在主线程中屏蔽了一些信号，那么被创建的线程会继承当前的信号屏蔽字
int pthread_sigmask(int how, const sigset_t *restrict set,sigset_t *restrict oset);

// Returns: 0 if OK, error number on failure
int sigwait(const sigset_t *restrict set, int *restrict signop);
// 线程可以通过 sigwait 函数等待一个或多个信号出现。如果多个线程通过该函数等待信号，则在传递信号的时候，只有一个线程可以从该函数返回。

// Returns: 0 if OK, error number on failure
// 可以调用 pthread_kill 函数将信号发送给指定的线程（需属于同一进程）
int pthread_kill(pthread_t thread, int signo);
// 如果传递给 signo 的值是 0，则可以用来检测线程是否存在。如果接收信号的线程没有对应的处理函数，则该信号会发送给主线程
```

```c
int main()
{
  int err;
  sigset_t mask, old;
  pthread_t pt1, pt2;

  sigemptyset(&mask);
  sigaddset(&mask, SIGQUIT); /* 如果不屏蔽QUIT信号，则主线程会收到该信号 */
  sigaddset(&mask, SIGINT);
  err = pthread_sigmask(SIG_BLOCK, &mask, &old);
  assert(err == 0);

  signal(SIGQUIT, main_q); /* QUIT信号处理函数 */

  err = pthread_create(&pt1, NULL, th1, NULL);
  assert(err == 0);

  sleep(1);
  printf("main:send QUIT signal.\n");
  // 线程1未屏蔽QUIT信号，但没有处理程序，会返回给主线程
  pthread_kill(pt1, SIGQUIT);

  sleep(10);

  return 0;
}
// 线程1
void* th1(void* a)
{
  int err, signo;
  sigset_t mask;

  sigemptyset(&mask);
  sigaddset(&mask, SIGINT);
  pthread_sigmask(SIG_BLOCK, &mask, NULL);

  while (1) {
    err = sigwait(&mask, &signo);
    assert(err == 0);
    switch (signo) {
      case SIGINT:
        printf("\nth1:INT.\n");
        break;
      default:
        printf("\nth1:unexcepted signal %d.\n", signo);
        break;
    }
  }
}
```

在多线程中，一般安排专用线程处理信号，通过互斥量的保护，信号处理线程可以安全地改动数据。

### 线程和fork

线程调用 fork 时，为子进程创建了整个进程地址空间的副本，同时还继承了互斥量、读写锁和条件变量的状态。为此，子进程返回后，如果不是马上调用 exec，则需要清理锁的状态。因为子进程中只含有调用 fork 的那个线程的副本，父进程中其他占有锁的线程在子进程中不存在。

要清除锁的状态，可以使用 pthread_atfork 函数建立 fork 处理程序。

```c
// Returns: 0 if OK, error number on failure
int pthread_atfork(void (*prepare)(void), void (*parent)(void),void (*child)(void));
```

prepare 由父进程在 fork 创建子进程前调用。任务是获取父进程定义的所有锁。
parent 在 fork 创建子进程后、返回之前在父进程上下文中调用。任务是对获取的所有锁进行解锁。
child 在 fork 返回前在子进程上下文中调用。任务是释放所有的锁。

可以多次调用该函数以设置多套 fork 处理程序。对于不需要的某个处理程序，可以传入空指针。多次调用时，parent 和 child 以注册时的顺序执行，而 prepare 的执行顺序与注册时相反。

# 进程间通信IPC

管道pipe-半双工（高级管道popen、有名管道FIFO-半双工）、消息队列、信号量（sem，用来控制多个进程对共享资源的访问，常作为一种锁机制，同步手段）、信号（signal）、共享存储（结合信号量）、套接字socket

FIFO命名管道

## 管道

一般来说，管道是半双工的（即数据只能在一个方向上流动），并且只能在具有公共祖先的两个进程之间使用。通常，父进程创建管道后会接着调用 fork，从而利用管道在父子进程之间通信。

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E5%8D%8A%E5%8F%8C%E5%B7%A5%E7%AE%A1%E9%81%93.png) 

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/fork-%E5%8D%8A%E5%8F%8C%E5%B7%A5%E7%AE%A1%E9%81%93.png)

父子进程可以分别关闭管道的读 / 写端，以利用管道在父子进程中传递信息。例如，如果想要创建从父进程到子进程的管道，则可以关闭父进程的读端和子进程的写端。由于管道半双工的特性，想要在父子进程间双向传递信息需要建立 2 个管道。

### 创建管道

```c
#include <unistd.h>
// Returns: 0 if OK, −1 on error
int pipe(int fd[2]);
```

fd 参数返回两个文件描述符，fd[0] 为读而打开，fd[1] 为写而打开。fd[1] 的输出是 fd[0] 的输入。

### 读写管道规则

- 当读一个写端被关闭的管道，在所有数据被读取后，read 返回 0

- 当写一个读端被关闭的管道，会产生 SIGPIPE 信号。如果忽略该信号或从信号处理程序返回，则 write 返回 - 1，且设置 errno 为 EPIPE

- 写入不超过 PIPE_ BUF 字节的操作是原子的，如果写入数据的大小超过该值，在多个进程同时写一个管道时，所写的数据可能交叉

### popen/pclose

管道的通常用法是创建一个连接到另一个进程的管道，然后读取其输出或者向其输入端发送数据。可以使用 popen 和 pclose 实现这一功能。

这两个函数实现的操作是：创建一个管道，fork 一个子进程，关闭未使用的管道，执行 shell 运行命令，然后等待命令终止。

```c
#include <stdio.h>
// Returns: file pointer if OK, NULL on error
FILE *popen(const char *cmdstring, const char *type);
// Returns: termination status of cmdstring, or −1 on error
int pclose(FILE *fp);
```

popen 先执行 fork，然后调用 exec 执行 cmdstring，并且返回一个标准 I/O 文件指针。

如果 type 是 "r"，则文件指针连接到 cmdstring 的标准输出，如果是 "w" 则连接到标准输入。cmdstring 以 sh -c cmdstring 的方式执行。pclose 函数关闭标准 I/O 流，等待命令终止，然后返回 shell 的终止状态。（注意不要使用 fclose 函数，它不会等待子进程结束）。

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/popen.png)

## FIFO命名管道

管道只能用于两个进程之间线性连接，FIFO有名字，可用于非线性连接。使得不相关的进程间也能交换数据，也是一种文件类型，创建 FIFO 与创建文件类似，需要指定其路径。

> `FIFO`是一种文件类型，通过`stat`结构的`st_mode`成员的编码可以知道文件是否是`FIFO`类型。可以用`S_ISFIFO`宏对此进行测试

FIFO有以下2个**用途**：

1. shell命令使用FIFO将数据从一条管道传送到另一条时，无需创建中间临时文件
2. 客户进程-服务器进程应用程序中，FIFO用作汇聚点（多个客户进程向服务器同一个众所周知的FIFO写），在客户进程和服务器进程二者之间传递数据

### 创建FIFO

```c
#include <sys/stat.h>
// Both return: 0 if OK, −1 on error
int mkfifo(const char *path, mode_t mode);
int mkfifoat(int dirfd, const char *path, mode_t mode);
```

- mkfifo：mode 参数指明 FIFO 的文件权限，与 open 函数中的 mode 相同。

- mkfifoat：path 参数有如下几种情况：
  - 如果指定为绝对路径，则会忽略 dirfd 参数，行为与 mkfifo 类似
  - 如果指定为相对路径，则该路径与 dirfd 打开的目录有关
  - 如果指定为相对路径，且 dirfd 有参数 AT_FDCWD，那么路径以当前目录开始

> 应用程序可以用`mknod`和`mknodat`函数创建FIFO。因为POSIX.1原先并没有包括`mknod`函数，所以`mkfifo`是专门为POSIX.1设计的。`mknod`和`mknodat`函数现在已包括在POSIX.1的XSI扩展中

### 打开FIFO

可以使用 open 打开 FIFO

- 一般情况下，在打开时如果没有设置非阻塞标志 O_NONBLOCK，只读（O_RDONLY）`open`要阻塞到某个其它进程为写而打开这个FIFO为止。类似地，只写（O_ WRONL ）`open`要阻塞到某个其它进程为读而打开它为止

但是，不应该使用 O_RDWR 的方式来绕过这种阻塞行为，而应该使用非阻塞标志。使用读写方式打开 FIFO，会导致读取数据时永远看不到文件结束，因为至少会有一个写描述符是打开着的。

### 读写FIFO

- 若`write`一个尚无进程为读而打开的FIFO，则产生信号`SIGPIPE` 
- 若某个FIFO的最后一个写进程关闭了该FIFO，则将为该FIFO的读进程产生一个文件结束标志

一个给定的FIFO有多个写进程是常见的。这意味着，如果不希望多个进程所写的数据交叉，则必须考虑原子写操作。和管道一样，常量`PIPE_BUF`说明了可被原子地写到FIFO的最大数据量。



![image-20220617161038339](C:\Users\xubei\AppData\Roaming\Typora\typora-user-images\image-20220617161038339.png) 

使用 FIFO 进行客户进程与服务器进程之间的通信。每个客户进程可以将自己的请求写到一个公共的 FIFO 文件中（请求长度需要小于 PIPE_BUF 以避免客户进程之间的数据交叉），服务器进程针对每个客户进程创建 FIFO，用于向客户进程发送数据。客户进程的 FIFO 的路径名可以使用客户进程的 PID 号作为基础，如 /tmp/servv1.PID，这样客户进程就知道该从哪个 FIFO 读取服务器进程返回的数据了。

## XSI-IPC

每个 IPC 对象与键（key）相关联，以使得多个进程可以通过它进行联系。在创建 IPC 结构时，必须指定一个键。而在系统内部，则使用标识符引用 IPC 结构。——**消息队列、信号量、共享内存。** 

### XSI-IPC介绍

#### 标识符和键

标识符

- 每个内核中的IPC结构（消息队列、信号量或共享内存）都用一个非负整数的**标识符**加以引用
- 与文件描述符不同，**IPC标识符**不是小的整数。当一个IPC结构被创建，然后又被删除时，与这种结构相关的标识符连续加1，直至到达一个整形数的最大正值，然后又回转到0

键

- **标识符是IPC对象的内部名**。为使多个合作进程能够在同一IPC对象上汇聚，需要提供一个外部命名方案。为此，**每个IPC对象都与一个键相关联，将这个键作为该对象的外部名**（创建IPC结构时，应指定一个键）。**键的类型是基本系统数据类型`key_t`**，通常在`<sys/types.h>`中被定义为长整形。这个键由内核变换成标识符。

#### 权限和结构

每个IPC结构关联了一个`ipc_perm`结构（`<sys/ipc.h>`），规定了权限和所有者，至少包括以下成员

```c
struct ipc_perm{
    uid_t uid;      /* 拥有者的有效用户ID */
    gid_t gid;      /* 拥有者的有效组ID */
    uid_t cuid;     /* 创建者的有效用户ID */
    gid_t cgid;     /* 创建者的有效组ID */
    mode_t mode;    /* 访问模式 */
    ...
};
```

- 创建IPC结构时，对所有字段都赋初值
- IPC结构的创建者或超级用户可以调用`msgctl`、`semctl`或`shmctl`修改`uid`、`gid`和`mode`字段。修改这些字段类似于对文件调用chown和chmod

**对于任何IPC结构都不存在执行权限**，消息队列和共享内存使用术语”读“和”写“，信号量则用”读“和”更改“。

键的创建方式，主要有如下几种：

- 指定为 IPC_PRIVATE，这会创建一个新的 IPC 结构，可以将返回的标识符存入文件供其他进程使用，也可直接给 fork 后的子进程使用

- 在公共头文件中定义一个键，然后由一个进程（通常是服务器进程）根据这个键来创建新的 IPC 结构。但是这种方式可能会与已经存在的键冲突，需要进程删除原有的 IPC 结构再重新创建。

- 使用 ftok 函数，将路径名和某个数字（0-255）变换为一个键。

> key_t ftok(const char *path, int id);  // Returns: key if OK, (key_t)−1 on error

path 参数必须引用的是现有的文件，id 参数只使用其低 8 位。

在创建 IPC 结构时还需要指定其权限，与文件权限类似，但是不存在执行权限。

注意：

- IPC_PRIVATE 只能用于创建新的 IPC 结构，而不能用来引用一个现有的 IPC 结构。
- 如果希望确保新创建的 IPC 结构没有引用具有同一标识符的现有 IPC 结构，则可以在 flag 中同时指定 IPC_CREAT 和 IPC_EXCL。这样，如果已经存在则会返回 EEXIST。

#### 优缺点

- 在系统范围内起作用，没有引用计数
  - 如果创建一个消息队列，放入消息后终止，消息队列和内容不会删除，直到调用msgrcv或msgctl读取或删除消息队列，或者ipcrm删除消息队列或自举系统删除消息队列，

- 在文件系统重没有名字
  - 不能用ls查看IPC对象
  - 不能对它们使用多路转接I/O函数（select-poll） 

### 消息队列

消息队列是消息的链接表，存储在内核中，由消息队列标识符标识。以下简称队列。相关的数据结构很少用到，后面的信号量和共享存储同理。

结构

每个队列都有一个`msqid_ds`结构与其关联，这个结构定义了队列的当前状态

```c
struct msqid_ds{
    struct ipc_perm    msg_perm;
    msgqnum_t          msg_qnum;    /* 队列中的消息数 */
    msglen_t           msg_qbytes;  /* 队列中消息的字节 */
    pid_t              msg_lspid;   /* 最后调用msgsnd()的进程ID */
    pid_t              msg_lrpid;   /* 最后调用msgrcv()的进程ID */
    time_t             msg_stime;   /* 最后调用msgsnd()的时间 */ 
    time_t             msg_rtime;   /* 最后调用msgrcv()的时间 */
    time_t             msg_ctime;   /* 最后一次修改队列的时间 */
    ...
};
```

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/msg%E7%B3%BB%E7%BB%9F%E9%99%90%E5%88%B6.png)

最大消息队列数，消息队列最大容量，一个消息最大长度8192。

[linux 消息队列的限制](https://www.programminghunter.com/article/8679861259/) ipcs -q

#### 创建/打开消息队列

msgget 用于创建或打开一个队列

```c
#include <sys/msg.h>
// Returns: message queue ID if OK, −1 on error
int msgget(key_t key, int flag);
// 
```

key 参数可以是通过 ftok 函数生成的，也可以是 IPC_PRIVATE。flag 用于设定读写权限，如果是新建该 IPC 结构则可以添加 IPC_CREAT。

#### 添加到队列

将新消息添加到队列尾端

```c
//  Returns: 0 if OK, −1 on error
int msgsnd(int msqid, const void *ptr, size_t nbytes, int flag);
// msqid 是 get 函数返回的队列 ID，nbytes 是消息数据的长度
// ptr 指向一个结构，其包含一个正的消息类型，和消息数据（nbytes 为 0 则无消息数据）,可以定义其结构如下
struct msgbuf {
    long mtype;       /* message type, must be > 0 */
    char mtext[1];    /* message data, of length nbytes */
};
```

flag 可以指定为 IPC_NOWAIT，当消息队列满时（或达到系统限制），会立即出错返回 EAGAIN。

否则，进程会一直阻塞直到：有空间容纳消息；队列被删除（返回 EIDRM)；或捕捉到信号并从处理程序返回（返回 EINTR）。

#### 获取消息

用于从队列中取出消息，可以指定获取某些类型的数据，而不是必须按照先进先出的次序。

```c
ssize_t msgrcv(int msqid, void *ptr, size_t nbytes, long type, int flag);
// ptr 指向的结构与 snd 函数一样，而 nbytes 则指定了消息长度，如果返回的消息长度 > nbytes，而 flag 中设置了 MSG_NOERROR，则消息被截断。如果没有设置则出错返回 E2BIG，而消息仍然留在队列中。
```

type==0：返回队列中的第一个消息

type>0：返回消息类型为 type 的第一个消息

type<0：返回消息类型≤type 绝对值的消息，如果有若干个满足则取类型最小的。

flag 参数同样可以指定为非阻塞

#### 操作消息队列

```c
int msgctl(int msqid, int cmd, struct msqid_ds *buf );
```

对队列执行多种操作，msqid 队列ID（标识符），`msgget`的返回值

cmd 参数指定队列需要执行的操作：

- IPC_STAT：获取队列的 msqid_ds 结构信息，存放于 buf 指向的结构中
- IPC_SET：将 msg_perm.uid，msg_perm.gid，msg_perm.mode 和 msg_qbytes 通过 buf 复制到队列的 msqid_ds 结构中。该命令只能由超级用户或者有效用户 ID 等于 msg_perm.cuid 或 msg_perm.uid 的用户执行。
- IPC_RMID：删除队列及其中的数据。也只能由上述的两类用户执行。

这 3 条命令也适用与信号量（semctl）和共享存储（shmctl）。

### 信号量

信号量是一个计数器，用于为多个进程提供对共享数据对象的访问。创建的时候需要指明信号量的个数，在使用的时候也要指明用的是哪个信号量。

```c
#include <sys/sem.h>
// Returns: semaphore ID if OK, −1 on error// 
// 创建或打开一个信号量合集 id
int semget(key_t key, int nsems, int flag);
// nsems 用于指定该集合中的信号量数，如果是创建新集合，则需要指定数量；如果是引用现有的集合，则将其设置为 0

// 包含多种信号量操作
int semctl(int semid, int semnum, int cmd, ... /* union semun arg */ );
// 第 4 个参数 arg 由 cmd 的实际值来决定是否使用，注意该参数并不是指针。如果需要使用该参数，其类型需要自己定义
union semun {
    int              val;   /* for SETVAL */
    struct semid_ds *buf;   /* for IPC_STAT and IPC_SET */
    unsigned short  *array; /* for GETALL and SETALL */
};
// 参数 semnum 用于指定信号量集合中的某个成员，该值在 0 ~ nsmes-1 之间
// cmd 由如下 10 个可选项: 1.IPC_STAT，IPC_SET，IPC_RMID：与队列类似；2.GETVAL，SETVAL：返回 / 设置（通过 arg.val）semnum 指定的成员的信号量值（semval）;3.GETPID，GETNCNT，GETZCNT：返回指定成员的 sempid，semncnt，semzcnt；4.GETALL，SETALL：取 / 设置所有的信号量值（通过 arg.array）
// 除 GETALL 以外所有的 GET 命令都由函数的返回值返回，其他命令则是成功返回 0，失败返回 - 1 并设置 errno

// 自动执行信号量集合上的操作数组
int semop(int semid, struct sembuf semoparray[], size_t nops);
// nops 是数组 semoparray 的元素个数，semoparray 是一个信号量操作数组，其中存放每个信号量的操作，其结构如下：
struct sembuf {
  unsigned short sem_num; /* member # in set (0, 1, ..., nsems-1) */
  short          sem_op;  /* operation (negative, 0, or positive) */
  short          sem_flg; /* IPC_NOWAIT, SEM_UNDO */
};
```

sem_flg 的 SEM_UNDO 标志标识当进程终止时，该操作修改的信号量值会被恢复，即重新设置为调用该操作之前的数值。

sem_op 可以指定如下 3 种值：

- 正值，表示进程释放的占用的资源数，sem_op 值会加到对应的信号量的值上。

- 0，表示进程希望等待该信号量值变为 0。IPC_NOWAIT 标志可以控制进程是否阻塞，相关的出错返回信息可以查阅手册，此处省略。

- 负值，表示进程想要获取的资源数。如果信号量值≥sem_op 的绝对值（满足需求），则会从当前的信号量值上减去对应的值，否则由 IPC_NOWAIT 标志决定进程是否阻塞。
  semop 函数具有原子性，即要么执行数组中所有的操作，要么什么也不做。

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E4%BF%A1%E5%8F%B7%E9%87%8F%E7%B3%BB%E7%BB%9F%E9%99%90%E5%88%B6.png) 

### 共享存储

共享存储允许两个或多个进程共享一个给定的存储区。但是，需要注意存储区访问的同步问题，当进程在写入数据时其他进程不应该去读取这些数据。一般使用信号量来解决这一同步问题。

相比与通过文件映射的方式来共享存储区的方式，XSI 共享存储没有相关的文件，它共享的是内存的匿名段。

> mmap就是共享存储的一种形式，但是XSI共享存储与其区别在于，XSI共享存储没有相关文件。XSI共享存储段是内存的匿名段

```c
#include <sys/shm.h>
// Returns: shared memory ID if OK, −1 on error
// 用于创建或引用一个共享存储段
int shmget(key_t key, size_t size, int flag);
// 实现一般将大小向上取整为系统页长的整数倍，若指定的 size 不是整数倍，则余下的空间是不可使用的

//  IPC_STAT，IPC_SET 和 IPC_RMID，相关解释可以参考消息队列部分
int shmctl(int shmid, int cmd, struct shmid_ds *buf );

// 将共享存储段连接到进程的地址空间。具体连接到地址空间的什么位置由 2、3 两个参数决定
void *shmat(int shmid, const void *addr, int flag);
// flag 还可以指定 SHM_RDONLY 以只读方式连接共享段

// 分离共享存储段
int shmdt(const void *addr);
// 这一操作不会删除系统中共享存储段的标识符及其数据结构。想要删除对应的数据结构，需要调用 shmctl 的 IPC_RMID 命令。
```

shmat 用于将共享存储段连接到进程的地址空间。具体连接到地址空间的什么位置由 2、3 两个参数决定。

1.addr=0，则连接到内核选择的第一个可用地址上。（推荐）
2.addr≠0，且 flag 没有指定 SHM_RND，那么连接到 addr 指定的地址。
3.addr≠0，且指定了 SHM_RND，那么系统会按照公式 (addr-(addr % SHMLBA)) 决定连接地址。该公式作用是将地址向下取最近的 SHMLBA 的倍数，而常数 SHMLBA 表示 “低边界地址倍数”。
flag 还可以指定 SHM_RDONLY 以只读方式连接共享段。

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/%E5%85%B1%E4%BA%AB%E5%AD%98%E5%82%A8%E7%B3%BB%E7%BB%9F%E9%99%90%E5%88%B6.png) 

## POSIX信号量

POSIX 信号量与 XSI 信号量最大的不同就是没有信号量集的概念，一次只能操作一个信号量。还有就是在删除信号量时，正在使用 XSI 信号量的操作会失败；而 POSIX 信号量的操作会正常执行，直到该信号量的最有一个引用被释放。

POSIX 信号量有两种形式：命名的和未命名的。两者的差异在于创建和销毁的形式上，使用的方式是一样的。未命名的信号量只存在于内存中，因此想要使用这些信号量的进程需要有对应的访问权限，如同一进程中的线程，或者是不同进程中映射相同的内存内容到自己的地址空间的线程。而命名信号量可以被任何直到它们名字的进程访问。

命名信号量：名字的第一个字符应该是 /。因为一般 POSIX 信号量的实现要使用文件系统；名字不应该包含其他斜杠；名字长度是实现定义的，不应长于_POSIX_NAME_MAX。

#### 创建/获取信号量

用于创建一个新的信号量或使用一个现有的信号量 oflag可为0

```c
#include <semaphore.h>
// Returns: Pointer to semaphore if OK, SEM_FAILED on error
sem_t *sem_open(const char *name, int oflag, ... /* mode_t mode,unsigned int value */ );
```

当 oflag 包含 O_CREAT 标志时，如果信号量不存在则会创建新的，如果存在则会被使用，但不会重新初始化。指定此标志时，还需要提供后面的 2 个参数。mode 指定访问权限，这与打开文件的权限相同；value 指定信号量的初值。
如果 oflag 同时指定了 O_EXCL 标志，则在创建信号量时，如果信号量已经存在就会出错

- 使用现有的命名信号量时，仅指定2个参数：
  - `name`：信号量的名字
  - `oflag`：设为0
- 创建新的命名信号量
  - `name`：信号量的名字
  - `oflag`：指定了`O_CREAT`标志。当该参数置为`O_CREAT|O_EXCL`并且信号量存在时，函数会失败
  - `mode`：谁可以访问信号量，值与[open函数](https://github.com/arkingc/note/blob/master/操作系统/UNIX环境高级编程.md#21-打开文件)的权限位相同
  - `value`：信号量的初始值（`0~SEM_VALUE_MAX`）

**为了移植性，信号量的命名应该遵循下列规则**：

- 名字的第一个字符应该为斜杠(`/`)
- 名字不应该包含其他斜杠以此避免实现定义的行为
- 信号量名字的最大长度是实现定义的，不应该鲳鱼`_POSIX_NAME_MAX`个字符长度。因为这是文件系统的实现能允许的最大名字长度的限制

#### 关闭信号量

```c
int sem_close(sem_t *sem);
```

sem_close关闭一个信号量，释放相关资源, 进程退出时如果没有调用该函数，系统也会自动关闭打开的信号量。POSIX 信号量没有 UNDO 机制，所以信号量的值不会受到影响

#### 销毁信号量

```c
int sem_unlink(const char *name);
// sem_unlink，
```

删除信号量的名字,如果没有打开的信号量引用，信号量会被立即销毁，否则会延迟到最后一个打开的引用关闭

#### 信号量操作

与 XSI 信号量不同，POSIX 信号量一次操作只能 + 1 或者 - 1

```c
#include <time.h>
// All return: 0 if OK, −1 on error
int sem_trywait(sem_t *sem);
int sem_wait(sem_t *sem);
int sem_timedwait(sem_t *restrict sem,const struct timespec *restrict tsptr);
// 这 3 个函数实现信号量的 - 1 操作
// 当信号量计数为 0 时，使用 sem_wait 函数会阻塞，直到成功使信号量 - 1 或者被信号中断；而 sem_trywait 会返回 - 1 且设置 errno 为 EAGAIN。
// 使用 sem_timedwait 可以设定等待时间，超时后会返回 - 1 且设置 errno 为 ETIMEOUT

int sem_post(sem_t *sem);
// 使信号量计数 + 1。如果有进程被改信号量阻塞，那么进程会被唤醒。

int sem_getvalue(sem_t *restrict sem, int *restrict valp);
// 获取信号量值，该数值存储在 valp 指向的地址处。注意函数返回的数值有可能是过时的
```

如果在多个进程间共享一个资源，则可使用3种技术中的一种来协调访问，可以使用映射到两个进程地址空间中的信号量、记录锁或者（共享存储中的）互斥量，（共享存储中的）互斥量更快，但是记录锁简单也快

#### 未命名信号量

主要用于单个进程

```c
// 创建一个未命名信号量，value 指定其初值，pshared 值为 0 时，信号量仅在进程的线程之间共享；不为 0 则表明会在进程之间共享
int sem_init(sem_t *sem, int pshared, unsigned int value);
int sem_destroy(sem_t *sem);
// 销毁未命名信号量。销毁之后不能使用任何带有 sem 的信号量函数，除非通过 sem_init 重新初始化它
```

## 套接字

```c
#include <sys/socket.h>
// 成功时返回一个套接字描述符，失败则返回-1
int socket(int domain, int type, int protocol);
```

int domain : 用于确定网络类型

int type : 用于确定协议类型

int protocol : 用于消歧义，决定网络协议，通常是0

参考：[Unix高级编程笔记](https://github.com/arkingc/note/blob/master/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/UNIX%E7%8E%AF%E5%A2%83%E9%AB%98%E7%BA%A7%E7%BC%96%E7%A8%8B.md) 

书籍：[《UNIX环境高级编程(第三版)》](https://awesome-programming-books.github.io/linux/UNIX环境高级编程(第三版).pdf) 