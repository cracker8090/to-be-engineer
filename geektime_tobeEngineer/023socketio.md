

## 介绍

​	socket.io是一个基于WebSocket的CS的实时通信库，它底层基于engine.io。engine.io使用WebSocket和xhr-polling(或jsonp)封装了一套自己的协议，在不支持WebSocket的低版本浏览器中(支持websocket的浏览器版本见[这里](https://link.jianshu.com/?t=https%3A%2F%2Fcaniuse.com%2F%23feat%3Dwebsockets))使用了长轮询(long polling)来代替。socket.io在engine.io的基础上增加了namespace，room，自动重连等特性。

Socket.IO是node.js的一个模块，它用于浏览器与服务端之间实时通信。它提供了服务器和客户端的组件，只需一个模块就可以给应用程序对webSocket的支持。Socket.IO解决了各个浏览器支持的问题。



**Socket.IO支持如下方式进行通信，会根据浏览器的支持程度，自动选择使用哪种技术进行通信：** 

WebSocket
Flash Socket
AJAX long-polling
AJAX multiple streaming
Forever IFrame
JSONP polling



**Socket.IO的API** 

// 监听客户端连接，回调函数会传递本次连接的socket
io.on('connection', function(socket){}); 

// 给所有客户端广播消息
io.sockets.emit('event_name', data); 

// 给指定的客户端发送消息
io.sockets.socket(socketid).emit('event_name', data);

// 监听发送的消息
socket.on('event_name', function(data) {});

// 给该socket的客户端发送消息
socket.emit('event_name', data);

// 给除了自己以外的客户端广播消息
socket.broadcast.emit("event_name", data);

[使用实例](https://www.cnblogs.com/tugenhua0707/p/8699309.html) [使用实例2](https://juejin.im/entry/59b126c46fb9a02487556964) [实例3](https://blog.csdn.net/zsj777/article/details/83189558) 

[socketio官网](https://socket.io/docs/) 

[Socket.io创建连接的参数](https://blog.csdn.net/taoerchun/article/details/48253013) 

[介绍](https://www.infoq.cn/article/2015/01/socket-io-websocket) [介绍2](https://www.jianshu.com/p/a3e06ec1a3a0) 





当我们io.connect()建立一个socket连接的时候，返回的是namespace实例，namespace实例中有个socket实例，当新建一个连接，或者发送一条消息的时候，namespace->socket->transport->websocket(xhrpolling...)，其实发送一条消息真正的发送者还是底层的websocket或是xhrpolling或其他的几种，而socket.io只是一个组织者，当我们需要建立连接的时候，它自己会在其内部挑选一种连接方式，然后实现连接。具体的socket.io内部实现还需要进一步学习源码。
	默认情况下，无论我们创建了多少次socket连接，指向都是同一个socket实例，除非我们在connect()的第二个参数中指定"force new socket"强制建立一个新的连接。socket实例监听了其connect，connect_failed等等事件，不要小看这些事件，其实都很有用，弄明白他们，socket.io基本无大问题，我们一一来分析一下。刚刚创建一个socket的时候，其实这只是一个空的socket，而且并未分配一个id，只有当与服务器连接成功后，才会生成一个id，因此我们在connect中才能得到获取socket的id，有些时候，socket的id是很有用的，有时会用于存储对应socket的用户信息。socket连接成功后，服务端与客户端会定时发送心跳包，若某一方长时间未回应，默认连接断开，当我们的网络不稳定或者有其他影响连接的因素存在的时候，则会与服务器断开连接，断开时会触发disconnect监听，若我们没有声明reconnect的时候，断开就不再与服务器重新连接，若我们在创建的时候如上声明了reconnect：true，那么断开后客户端还是会尝试新的连接，就会触发reconnecting监听，若连接未成功，会继续尝试多次连接，若连接成功，则会触发reconnect监听，此时，原有的socket连接已经不再，内部会创建一个新的连接，并且socket得到一个新的id，最后会再一次触发connect监听。
	用原来的对象重连发起连接就可以了
	socket = socket.connect.

[socketit客户端开发](https://www.jianshu.com/p/11d45bfd03ed) 

能够监听的事件有：

| 事件名称         | 事件说明           | 事件回调参数说明             |
| ---------------- | ------------------ | ---------------------------- |
| connect          | 连接成功的事件     | 无参数                       |
| connect_failed   | 连接失败的事件     | 参数Object，里面包含错误信息 |
| disconnect       | 断开连接的事件     | 无参数                       |
| connecting       | 正在链接的事件     | 无参数                       |
| error            | 连接错误事件处理   |                              |
| connect_timeout  | 连接超时的事件     | 无参数                       |
| reconnect        | 成功重新连接的事件 | 参数Number，重连的次数       |
| reconnecting     | 正在重新连接       | 参数Number，重连的次数       |
| reconnect_failed | 重新连接失败       | 无参数                       |







### 心跳

在使用websocket的过程中，有时候会遇到网络断开的情况，但是在网络断开的时候服务器端并没有触发onclose的事件。这样会有：服务器会继续向客户端发送多余的链接，并且这些数据还会丢失。所以就需要一种机制来检测客户端和服务端是否处于正常的链接状态。因此就有了websocket的心跳了。还有心跳，说明还活着，没有心跳说明已经挂掉了。

**1. 为什么叫心跳包呢？**
它就像心跳一样每隔固定的时间发一次，来告诉服务器，我还活着。

**2. 心跳机制是？**
心跳机制是每隔一段时间会向服务器发送一个数据包，告诉服务器自己还活着，同时客户端会确认服务器端是否还活着，如果还活着的话，就会回传一个数据包给客户端来确定服务器端也还活着，否则的话，有可能是网络断开连接了。需要重连~





### WebSocket介绍

​	在HTTP 协议开发的时候，并不是为了双向通信程序准备的，起初的 web 应用程序只需要 “请求-响应” 就够了。由于历史原因，在创建拥有双向通信机制的 web 应用程序时，就只能利用 HTTP 轮询的方式，由此产生了 “短轮询” 和 “长轮询”(注意区分短连接和长连接)。

​	短轮询通过客户端定期轮询来询问服务端是否有新的信息产生，缺点也是显而易见，轮询间隔大了则信息不够实时，轮询间隔过小又会消耗过多的流量，增加服务器的负担。长轮询是对短轮询的优化，需要服务端做相应的修改来支持。客户端向服务端发送请求时，如果此时服务端没有新的信息产生，并不立刻返回，而是Hang住一段时间等有新的信息或者超时再返回，客户端收到服务器的应答后继续轮询。可以看到长轮询比短轮询可以减少大量无用的请求，并且客户端接收取新消息也会实时不少。

​	虽然长轮询比短轮询优化了不少，但是每次请求还是都要带上HTTP请求头部，而且在长轮询的连接结束之后，服务器端积累的新消息要等到下次客户端连接时才能传递。更好的方式是只用一个TCP连接来实现客户端和服务端的双向通信，WebSocket协议正是为此而生。WebSocket是基于TCP的一个独立的协议，它与HTTP协议的唯一关系就是它的握手请求可以作为一个`Upgrade request`经由HTTP服务器解析，且与HTTP使用一样的端口。WebSocket默认对普通请求使用80端口，协议为`ws://`，对TLS加密请求使用443端口，协议为`wss://`。

​	握手是通过一个`HTTP Upgrade request`开始的，一个请求和响应头部示例如下(去掉了无关的头部)。WebSocket握手请求头部与HTTP请求头部是兼容的（见RFC2616）。

```csharp
## Request Headers ##
Connection: Upgrade
Host: socket.io.demo.com
Origin: http://socket.io.demo.com
Sec-WebSocket-Extensions: permessage-deflate; client_max_window_bits
Sec-WebSocket-Key: mupA9l2rXciZKoMNQ9LphA==
Sec-WebSocket-Version: 13
Upgrade: websocket

## Response Headers ##
101 Web Socket Protocol Handshake
Connection: upgrade
Sec-WebSocket-Accept: s4VAqh7eedG0a11ziQlwTzJUY3s=
Sec-WebSocket-Origin: http://socket.io.demo.com
Server: nginx/1.6.2
Upgrade: WebSocket
```

- Upgrade 是HTTP/1.1中规定的用于转换当前连接的应用层协议的头部，表示客户端希望用现有的连接转换到新的应用层协议WebSocket协议。
- Origin 用于防止跨站攻击，浏览器一般会使用这个来标识原始域，对于非浏览器的客户端应用可以根据需要使用。
- 请求头中的 Sec-WebSocket-Version 是WebSocket版本号，Sec-WebSocket-Key 是用于握手的密钥。Sec-WebSocket-Extensions 和 Sec-WebSocket-Protocol 是可选项，暂不讨论。
- 响应头中的 Sec-WebSocket-Accept 是将请求头中的 Sec-WebSocket-Key 的值加上一个固定魔数`258EAFA5-E914-47DA-95CA-C5AB0DC85B11`经SHA1+base64编码后得到。计算过程的python代码示例（uwsgi中的实现见 core/websockets.c的 `uwsgi_websocket_handshake`函数）

```java
magic_number = '258EAFA5-E914-47DA-95CA-C5AB0DC85B11'
key = 'mupA9l2rXciZKoMNQ9LphA=='
accept = base64.b64encode(hashlib.sha1(key + magic_number).digest())
assert(accept == 's4VAqh7eedG0a11ziQlwTzJUY3s=')
```

- 客户端会检查响应头中的status code 和 Sec-WebSocket-Accept 值是否是期待的值，如果发现Accept的值不正确或者状态码不是101，则不会建立WebSocket连接，也不会发送WebSocket数据帧。

ebSocket协议使用帧（Frame）收发数据，帧格式如下。基于[安全考量](https://link.jianshu.com/?t=https%3A%2F%2Ftools.ietf.org%2Fhtml%2Frfc6455%23page-51)，**客户端发送给服务端的帧必须通过4字节的掩码（Masking-key）加密**，服务端收到消息后，用掩码对数据帧的Payload Data进行异或运算解码得到数据（详见uwsgi的 core/websockets.c 中的uwsgi_websockets_parse函数），如果服务端收到未经掩码加密的数据帧，则应该马上关闭该WebSocket。而**服务端发给客户端的数据则不需要掩码加密**，客户端如果收到了服务端的掩码加密的数据，则也必须关闭它。

```ruby
 0                   1                   2                   3
      0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
     +-+-+-+-+-------+-+-------------+-------------------------------+
     |F|R|R|R| opcode|M| Payload len |    Extended payload length    |
     |I|S|S|S|  (4)  |A|     (7)     |             (16/64)           |
     |N|V|V|V|       |S|             |   (if payload len==126/127)   |
     | |1|2|3|       |K|             |                               |
     +-+-+-+-+-------+-+-------------+ - - - - - - - - - - - - - - - +
     |     Extended payload length continued, if payload len == 127  |
     + - - - - - - - - - - - - - - - +-------------------------------+
     |                               |Masking-key, if MASK set to 1  |
     +-------------------------------+-------------------------------+
     | Masking-key (continued)       |          Payload Data         |
     +-------------------------------- - - - - - - - - - - - - - - - +
     :                     Payload Data continued ...                :
     +---------------------------------------------------------------+
```

帧分为控制帧和数据帧，控制帧不能分片，数据帧可以分片。主要字段说明如下：

- FIN: 没有分片的帧的FIN为1，分片帧的第一个分片的FIN为0，最后一个分片FIN为1。
- opcode: 帧类型编号，其中控制帧：0x8 (Close), 0x9 (Ping), and 0xA (Pong)，数据帧主要有：0x1 (Text), 0x2 (Binary)。
- MASK：客户端发给服务端的帧MASK为1，Masking-key为加密掩码。服务端发往客户端的MASK为0，Masking-key为空。
- Payload len和Payload Data分别是帧的数据长度和数据内容。



### engine.io和socket.io

前面提到socket.io是基于engine.io的封装，engine.io（协议版本3）有一套自己的协议，任何engine.io服务器都必须支持polling(包括jsonp和xhr)和websocket两种传输方式。**engine.io使用websocket时有一套自己的ping/pong机制，使用的是opcode为0x1(Text)类型的数据帧，不是websocket协议规定的ping/pong类型的帧，标准的 ping/pong 帧被uwsgi使用**。

engine.io的数据编码分为Packet和Payload，其中 Packet是数据包，有6种类型：

- 0 open：从服务端发出，标识一个新的传输方式已经打开。
- 1 close：请求关闭这条传输连接，但是它本身并不关闭这个连接。
- 2 ping：客户端周期性发送ping，服务端响应pong。注意这个与uwsgi自带的ping/pong不一样，uwsgi里面发送ping，而浏览器返回pong。
- 3 pong：服务端发送。
- 4 message：实际发送的消息。
- 5 upgrade：在转换transport前，engine.io会发送探测包测试新的transport（如websocket）是否可用，如果OK，则客户端会发送一个upgrade消息给服务端，服务端关闭老的transport然后切换到新的transport。
- 6 noop：空操作数据包，客户端收到noop消息会将之前等待暂停的轮询暂停，用于在接收到一个新的websocket强制一个新的轮询周期。

而Payload是指一系列绑定到一起的编码后的Packet，它只用在poll中，websocket里面使用websocket帧里面的Payload字段来传输数据。如果客户端不支持XHR2，则payload格式如下，其中length是数据包Packet的长度，而packet则是编码后的数据包内容(测试发现客户端发送给服务端的poll请求中的payload用的这种字符编码)。

```ruby
<length1>:<packet1>[<length2>:<packet2>[...]]
```

若支持XHR2，则payload中内容全部以字节编码，其中第1位0表示字符串，1表示二进制数据，而后面接着的数字则是表示packet长度，然后以\xff结尾。如果一个长度为109的字符类型的数据包，则前面长度编码是 `\x00\x01\x00\x09\xff`，然后后面接packet内容。(测试发现服务端返回给客户端的payload为这种字节编码)

```tsx
<0 for string data, 1 for binary data><Any number of numbers between 0 and 9><The number 255><packet1 (first type,
then data)>[...]
```

engine.io服务器维护了一个socket的字典结构用于管理连接到该机的客户端，而客户端的标识就是sid。**如果有多个worker，则需要保证同一个客户端的请求落在同一台worker上(可以配置nginx根据sid分发)**。因为每个worker只维护了一部分客户端连接，如果要支持广播，room等特性，则后端需要使用 redis 或者 RabbitMQ 消息队列，使用redis的话则是通过redis的订阅发布机制实现多机多worker之间的消息推送。

**socket.io是engine.io的封装，在其基础上增加了自动重连，多路复用，namespace，room等特性。**socket.io本身也有一套协议，它Packet类型分为`(CONNECT 0, DISCONNECT 1, EVENT 2, ACK 3, ERROR 4, BINARY_EVENT 5, BINARY_ACK 6)`。注意与engine.io的Packet类型有所不同，但是socket.io的packet实际是借助的engine.io的Message类型发送的，在后面实例中可以看到Packet的编码方式。当连接出错的时候，socket.io会通过自动重连机制重新连接。













































