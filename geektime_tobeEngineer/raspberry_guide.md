树莓派使用





# 上网设置

1.wifi上网

把WIFI网卡插入树莓派的USB口中，用 lsusb 命令查看USB设备列表

 例：可以见到USB设备列表中有 "RTL8188CUS 802.11N WLAN Adapter"（不同设备名字不同） 的字样 说明该USB设备已被系统识别，芯片是RTL8188。用 ifconfig 命令可以看到 wlan0 设备，但没有IP地址(未连接)

```undefined
sudo raspi-config
可以设置声音输出是3.5mm还是hdmi
amixer cset numid=3 1 	//设置为耳机输出
amixer cset numid=3 2	//设置为hdmi输出
aplay /usr/share/sounds/alsa/Front_Center.wav
```

树莓派VNC.pdf——主要介绍如何是用VNC接入树莓派

比较全的树莓派入门介绍.pdf——主要介绍sd卡安装几基本介绍

树莓派pi_初级教程.pdf——99%是电源问题，灯闪烁提示故障







# 外挂硬盘

JDownloader无广告bt下载

树莓派启动的时候固定是从 SD 卡的第一个分区（/boot）读取 `config.txt`、`cmdline.txt` 和 `kernel.img` 这三个文件。我配的是5V3A的电源，外接一个有自动休眠功能的硬盘，做NAS妥妥的。请注意，硬盘一定要有闲时休眠功能，否则用一段就挂了。

[树莓派添加USB外接硬盘](https://blog.csdn.net/huayucong/article/details/48812573) ，[给树莓派挂载移动硬盘或U盘](http://shumeipai.nxez.com/2013/09/08/raspberry-pi-to-mount-the-removable-hard-disk.html) ，[树莓派 RASPBERRY PI SD卡系统备份与还原](https://zvv.me/xiankan/RASPBERRY-PI-SD-backup.html) 





# [树莓派外网访问](https://github.com/ma6174/blog/issues/7) 



在树莓派上面写一个脚本（任意语言）将采集到的传感器上传到VPS上面的网站程序。又或者使用现在流行的物联云，例如著名的yeelink来实现，这样就连vps也不用了。就能随处查看树莓派上面的传感器信息。



## Vpn

> 顾名思义就是让Raspberry与你当前的Client处于同一个网段，然后通过内网ssh.
>
> 所以要在`vps`上搭建`VPN`, 然后树莓派和控制端都连上`VPN`

网上找到[这样的方法](http://lukin.cn/201312/Ubuntu_command_How_to_connect_VPN/)连vpn：`sudo pptpsetup --create vpnname --server ip --username test --password test --encrypt --start`，但是树莓派一执行这个命令就断网，原因未知，只能重启。

## SSH内网穿透

> 原理是这样的假设`vps`地址是`10.10.10.10`，树莓派通过`ssh`连接到`vps`，同时将`vps`上某个端口比如`8888`映射到树莓派的`ssh`端口比如`22`，这样在`vps`上访问8888端口就相当于访问树莓派的`22`端口。
>
> ```
> ssh参数解释：
> -N：不执行任何指令
> -f：背景执行
> -R：主要建立reverse tunnel的参数
> 
> $ ssh -f -N -R 8888:localhost:22 username@10.10.10.10
> 
> ssh -f -N -R 57612:127.0.0.1:223 root@65.49.228.80
> ssh -f -N -R 57612:127.0.0.1:223 root@65.49.228.80
> 
> （1）ssh -p 26435 -f -N -R 30000:localhost:22 root@65.49.228.80
> autossh -p 26435 -M 30001 -f -N -R 30000:localhost:22 root@65.49.228.80
> autossh -p 26435 -M 30001 -f -N -R '*:30000:localhost:22' root@65.49.228.80
> 
> lsof -i :30000
> （2）局域网手机（ssh -N -f -L 2222:127.0.0.1:2222 root@vps的IP；ssh -p 2222 username@localhost）
> 
> ssh   -CfNg -R 57612:127.0.0.1:223  65.49.228.80
> 
> vps修改防火墙：
> iptables-L-n可以查看当前iptables的开放端口情况
> vim /etc/sysconfig/iptables
> 
> -A INPUT -p tcp -m multiport --dport 20,21  -m state --state NEW -j ACCEPT   --开启20,21端口
> -A INPUT -p tcp -m state --state NEW -m tcp --dport 21 -j ACCEPT            --开启21主动端口
> -A INPUT -p tcp --dport 30000:31000 -j ACCEPT             --开启被动端口
> 
> 开放某端口：iptables -I INPUT -p tcp --dport 80 -j ACCEPT
> 关闭某端口 : iptables -D INPUT -p tcp --dport 80 -j ACCEPT
> 屏蔽某个IP请求 : iptables -I INPUT -s 192.168.0.1 -j DROP (屏蔽单个IP192.168.0.1)
> 屏蔽IP某段请求 : iptables -I INPUT  -s 192.168.0.0/16  -j DROP(屏蔽单个IP192.168.0.0-192.168.255.255)
> 屏蔽整个IP段请求 ：iptables -I INPUT -s 192.168.0.0/16   -j DROP(屏蔽单个IP192.0.0.0-192.255.255.255)
> 添加iptables配置项：service iptables save  
> 重新启动服务：service iptables restart
> 
> 目标A：内网服务器（开发服务器），服务器端口8080，映射端口8088
> 目标B：外网VPS，服务器端口8089，映射本地端口8088，公网IP：192.168.1.111(举例IP，替换成相应的地址即可)
> 访问者C : 可以访问VPS，不能访问A
> 1）目标A，启动服务(原理SSH 反向代理)
> 	ssh -fCNR 8088:localhost:8080 root@192.168.1.111
> 2）目标B，启动服务(原理SSH 正向代理)
> 	ssh -fCNL *:8089:localhost:8088 localhost
> 3）访问者C， 访问 http://192.168.1.111:8089, 完成访问
> 
> ssh 连接不稳定，一段时间可能断开的问题，使用autossh解决
> #-M 6677为 autossh 维持访问的端口, 其他参数同ssh
> ➜  ~ autossh -M 6677 -fCNR 8088:localhost:8088 root@XXX
> 
> sudo pacman -S autossh
> autossh -p 22 -M 6777 -NR 6766:localhost:22 usera@a.site
> 打洞修改autossh -p 22 -M 6777 -NR '*:6766:localhost:22' usera@a.site
> 
> B $ sudo useradd -m autossh
> B $ sudo passwd autossh
> 免密ssh-copy-id username@1.1.1.1
> autossh -M 5678 -fNR 2222:localhost:22 username@1.1.1.1
> ```

[autossh](http://arondight.me/2016/02/17/%E4%BD%BF%E7%94%A8SSH%E5%8F%8D%E5%90%91%E9%9A%A7%E9%81%93%E8%BF%9B%E8%A1%8C%E5%86%85%E7%BD%91%E7%A9%BF%E9%80%8F/) 

常用ss命令：

ss -l 显示本地打开的所有端口

ss -pl 显示每个进程具体打开的socket

ss -t -a 显示所有tcp socket

ss -u -a 显示所有的UDP Socekt

ss -o state established '( dport = :smtp or sport = :smtp )' 显示所有已建立的SMTP连接

ss -o state established '( dport = :http or sport = :http )' 显示所有已建立的HTTP连接

ss -x src /tmp/.X11-unix/* 找出所有连接X服务器的进程

ss -s 列出当前socket详细信息





## Dynamic DNS

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



## [Ngrok](https://ngrok.com/)

> `ngrok`目前是非常流行的反向代理服务，可以进行内网穿透，支持80端口以及自定义`tcp`端口转发。
> 这样即使你的树莓派没有公网IP也可以使用SSH远程登陆。
>
> Offical Website: <https://ngrok.com/>
> Open Source: <https://github.com/inconshreveable/ngrok> 



## [frp建立服务器注册](http://www.10tiao.com/html/357/201801/2247485670/1.html) 





# 连接摄像头

**sudo apt-get update**

**sudo apt-get upgrade** 

sudo raspi-config，选择接口->摄像头，重启

内存卡修复工具SDFormatter.exe，ipscan22.exe（https://pan.baidu.com/s/1skNBZsT#list/path=%2F，3oar）





# 摄像头 VLC远程监控

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



# 实现视频通话







# python,scratch,c,ruby

ruby on rails







# 媒体播放中心

XBMC开源项目已经移植到raspberry上了

1、什么是XBMC（Kodi）
XBMC是一个优秀的自由和开源的（GPL）媒体中心软件。XBMC最初为Xbox而开发，可以运行在Linux、OSX、Windows、Android系统。让用户播放本地或网络中的大多数视频、音乐、播客及各种常见文件。XBMC 的可定制性很高，有许多[皮肤](http://zh.wikipedia.org/w/index.php?title=%E7%9A%AE%E8%82%A4(%E7%94%B5%E8%84%91)&action=edit&redlink=1)可以更改软件外观及各种可以访问网络内容的[插件](http://zh.wikipedia.org/wiki/%E6%8F%92%E4%BB%B6)，附带录制直播节目的数字视频录像机图形界面前端，同时支持电子节目指南和高清视频录制。
从XBMC14.0开始，XBMC正式更名为Kodi 。

可以看到XBMC支持的连接方式还是很多的；设置 该目录包含  为 电影，选择刮削器 为 The Movie Database，接下来 点击设置，注意将首选语言改为 zh。接下来一路确定就可以了。如果你不喜欢刻个刮削器的话，还可以安装豆瓣的刮削器，可以看到豆瓣里对电影的介绍。安装方法往下看。

**插件**：丰富的插件资源是XBMC的另一大特色，可以安装射手的插件在线搜索字幕，安装bt的插件可以下载等等等等，还有一些特殊的影视插件，你懂的。安装插件的方法有很多，从zip安装插件或者插件包，从获取插件下直接安装等等。以安装射手字幕插件为例，选择字幕-Shooter，点击安装即可。

这一部分的最后，提供两个中文插件包以及一个号称是全球最大的插件包，下载地址：http://pan.baidu.com/s/1xqEDK

皮肤：其实皮肤也是一种插件，安装方法可上面一样，但为什么还要说呢，因为大部分皮肤都是英文的皮肤，不支持中文，所以你要是想找能用的皮肤的时候最好先换回英文语言。
接下来介绍一个我最喜欢的，在大神汉化的基础上而来的，也是在XBMC里面非常出名的一个主题——Aeon MQ 5 
进入 设置-皮肤-皮肤，选择获取更多，安装Aeon MQ 5

**外挂播放器** ：如果用着XBMC自带的播放器感觉不顺手的话，可以试试看外挂播放器，毕竟像PotPlayer，TMT之类的播放效果可能更好一些。

**中文支持**：接下来关闭XBMC，进入XBMC的安装目录

【SKIN】，然后选右边的【FONTS】为【Arial based】，右边的【Language】为【Chinese (Simple)】、【Character set】为【Chinese Simplified（GBK)】。[KODI插件库下载汇总](http://www.kodiplayer.cn/) 

Windows的一般在C:\Users\用户名\AppData\Roaming\XBMC /addons下
Android的一般在/sdcard/Android//org.xbmc.xbmc/files/.xbmc/addons下
进入skin.aeonmq5目录，添加替换以下文件/
相关文件下载：http://pan.baidu.com/s/1sjnYEgX
1)替换720P文件夹中的Font.XML
2)在字体文件夹fonts中添加.ttf 和msyh.ttf 两个字体
3)添加Chinese (Simple)到语言文件夹language中
重启XBMC后，将fonts设置为，重新选择语言为Chinese（Simple），就可以看到漂亮的中文了

这个皮肤相对于原生而言，可自定义程度很高，你可以随意关闭或者更改相关项目。
比如：在下图位置按向上键可以调出视图设置

**自定义菜单**：进入设置-皮肤设置-主菜单-自定义主菜单

里面对于菜单的修改便很齐全了，包括显示，替换等等，可以按需修改。特别注意的是，如果发现重命名不支持中文输入的话，一种方法是安装中文输入插件，另一种方法是直接进入皮肤文件进行修改。
具体位置在安装目录XBMC\user\下，修改guisettings.XML，搜索并替换你要修改的名字就可以了



我在Android上装的版本是SPMC，定制版，详见[官网](http://spmc.semperpax.com/)
为了防止皮肤自动更新替换掉字体，可以关闭自动更新插件功能。
可以打开开机自动扫描媒体库功能，方便省事。
如果喜欢特效字幕的话，可以去下面的这个站点去下载，应该是国内最棒的特效字幕下载站了。
[触摸春天音轨字幕网](http://cmct.cc/?fromuid=314583) 



2.[搭建私人家庭影院](https://www.jianshu.com/p/339a90f1df7c) https://www.zhihu.com/question/29004138



kodi，这个播放器与mpc-hc基本可以看做一家人,目前也就这个分支在继续开发,独立安装的播放器Kodi

https://github.com/taxigps/xbmc-addons-chinese/blob/master/plugin.video.asia-tv/Users_Guide.txt

[追美剧](http://allenlow.com/blog/2017/04/02/kodi%E9%85%8D%E7%BD%AE%E4%B8%AD%E6%96%87%E5%B9%B6%E5%AE%89%E8%A3%85%E6%8F%92%E4%BB%B6%E8%A7%82%E7%9C%8B%E7%9B%B4%E6%92%ADtv%E5%92%8C%E8%BF%BD%E7%BE%8E%E5%89%A7%E6%95%99%E7%A8%8B/) 



Xbian，变成盒子



smb配置

```shell
sudo pacman -S samba gvfs-smb thunar-shares-plugin
sudo pacman -S manjaro-settings-samba会自动做一些配置，/etc/samba/smb.conf
[kodi_share]
comment = Kodishare Directories
path = /run/media/dhasijkd/39B544B735E40466/classic_video
browseable = yes
read only = no
create mask = 0700
directory mask = 0700 
添加分享用户并设置密码
gpasswd sambashare -a username
smbpasswd -a username
启用smaba 服务
systemctl enable smb nmb
systemctl start smb nmb
```







