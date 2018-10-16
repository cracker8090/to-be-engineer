# 树莓派使用





## 树莓派外网访问



在树莓派上面写一个脚本（任意语言）将采集到的传感器上传到VPS上面的网站程序。又或者使用现在流行的物联云，例如著名的yeelink来实现，这样就连vps也不用了。就能随处查看树莓派上面的传感器信息。



### Vpn

> 顾名思义就是让Raspberry与你当前的Client处于同一个网段，然后通过内网ssh.
>
> 所以要在`vps`上搭建`VPN`, 然后树莓派和控制端都连上`VPN`



### SSH内网穿透

> 原理是这样的假设`vps`地址是`10.10.10.10`，树莓派通过`ssh`连接到`vps`，同时将`vps`上某个端口比如`8888`映射到树莓派的`ssh`端口比如`22`，这样在`vps`上访问8888端口就相当于访问树莓派的`22`端口。
>
> ```
> $ ssh -f -N -R 8888:localhost:22 username@10.10.10.10
> ```



#### Dynamic DNS

> 常用的[花生壳](http://hsk.oray.com/)动态域名, 当然还有很多其他免费的`DNS`运营商。
>
> <http://www.noip.com/>
> <https://duckdns.org/>
> <http://www.dnsdynamic.org/>
> <http://www.dynu.com/>
> <http://www.changeip.com/dns.php>
> ……
>
> 注册一个免费的壳域名，在路由器端`DNS`解析输入申请的壳域名账号及密码，配置完成。所有的工作都交给了壳域名来操作。
>
> 如果还做了端口映射，请用壳域名+映射端口来`ssh`
>
> Resource:
> <http://hsk.oray.com/get/?icn=hsk_get&ici=hsk_home-grid#topology>
> <http://hsk.oray.com/news/4168.html>



#### [Ngrok](https://ngrok.com/)

> `ngrok`目前是非常流行的反向代理服务，可以进行内网穿透，支持80端口以及自定义`tcp`端口转发。
> 这样即使你的树莓派没有公网IP也可以使用SSH远程登陆。
>
> Offical Website: <https://ngrok.com/>
> Open Source: <https://github.com/inconshreveable/ngrok> 





## 摄像头 VLC远程监控

```bsh
sudo apt-get update
sudo apt-get install vlc
sudo raspivid -o - -t 0 -w 640 -h 360 -fps 25|cvlc -vvv stream:///dev/stdin --sout '#standard{access=http,mux=ts,dst=:8090}' :demux=h264
```

 第一行是更新软件数据库

 第二行是安装vlc

 第三行是使用PI官方的raspivid捕获视频工具把视频流输出到vlc，通过vlc转码成h264网络视频流通过http协议以ts的形式封装，然后输出到8090端口，用这个当监控只要网络稳定绝对不卡。

打开客户端--》媒体--》网络串流--》如下图输入  **http://PI的IP地址:8090** 

摄像头模块工作的时候那个红色的灯会一直亮，嘿嘿，要想禁用它的话 

 sudo nano /boot/config.txt 

 然后加入一行  disable_camera_led=1  



## 实现视频通话





































