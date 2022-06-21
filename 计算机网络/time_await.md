# 查看time_wait连接

netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'

# 原因

一般是客户端发起的。某些服务如pop/smtp、ftp却是服务端收到客户端的QUIT命令后主动关闭连接，这就造成这类服务器上容易出现大量的TIME_WAIT状态的连接，而且并发量越大处于此种状态的连接越多。
time_wait会占用系统资源，socket是有限的65535
会出现在负载均衡服务器上，会与服务进行处理

# 处理方式

1.长连接
2.如果是服务器端发起的，可以解决ip，启用Nginx代理等
3.linux系统层修改tcp参数，可重用或者2MSL时间修改
4.由客户端来主动断开连接

如果主动关闭的一方不维护这样一个TIME_WAIT状态，那么当被动关闭一方重发的FIN到达时，主动关闭一方的TCP传输层会用RST包响应对方，这会被对方认为是有错误发生，然而这事实上这只是正常的关闭连接过程，并非异常。

确保被动关闭方收到ACK，连接正常关闭，且不因被动关闭方重传 FIN 影响下一个新连接。

```bash
net.ipv4.tcp_tw_reuse = 1
表示开启重用。允许将TIME-WAIT sockets重新用于新的TCP连接，默认为0，表示关闭；

net.ipv4.tcp_tw_recycle = 1 
表示开启TCP连接中TIME-WAIT sockets的快速回收，默认为0，表示关闭。

net.ipv4.tcp_syncookies = 1 
表示开启SYN Cookies。当出现SYN等待队列溢出时，启用cookies来处理，可防范少量SYN攻击，默认为0，表示关闭；

tcp_max_tw_buckets
即表示系统允许同时存在的处于TIME_WAIT状态的socket数量。该配置项可以用来防范简单的Dos攻击，在某些情况下可以适当调大，但绝对不应当调小，否则后果自负

ip_local_port_range
该项说明了local port的分配范围，从上面可以看到默认的可用端口数不到3W。

tcp_max_syn_backlog
incomplete connection queue的最大长度

tcp_syn_retries
三次握手时SYN的最大重试次数

tcp_fin_timeout
对于本端断开的socket连接，TCP保持在FIN_WAIT_2状态的时间。对方可能会断开连接或一直不结束连接或不可预料的进程死亡。默认值为60秒。过去在2.2版本的内核中是180秒。您可以设置该值，但是需要注意，如果你的机器为负载很重的Web服务器，你可能要冒内存被大量无效数据报填满的风险。FIN_WAIT_2 sockets的危险性低于FIN_WAIT_1，因为它们最多只吃1.5K的内存，但是它们存在时间更长。

tcp_timestamps
为1表示开启TCP时间戳，用来计算往返时间RTT（Round-Trip Time）和防止序列号回绕
```

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/timeawait.png)

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/握手SYN.png) 