[UNIX网络编程卷1笔记](https://github.com/arkingc/note/blob/master/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/UNIX%E7%BD%91%E7%BB%9C%E7%BC%96%E7%A8%8B%E5%8D%B71.md) 

[源码1](https://github.com/arkingc/unpv13e) 

| I/O复用 | 高级I/O函数 | 非阻塞式I/O |
| :-----: | :---------: | :---------: |
|         |             |             |

- I/O复用

select

pselect

poll

epoll

- 高级I/O函数

Unix I/O函数，recv和send、readv和writev、recvmsg和sendmsg，5组I/O函数对比

标准I/O函数

- 非阻塞式I/O

非阻塞读写、非阻塞connect、非阻塞accept