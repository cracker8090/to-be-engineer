

# linux系统推荐

Which Linux distros are suitable for developers

https://fossbytes.com/best-linux-distros-for-programming-developers/

https://fossbytes.com/secure-linux-distros-privacy-anonymity/



Tails

Tails做到的其实仅是隐藏你的真实IP，却不能保证你的软件和系统的安全。

如果你希望更安全的隐私保护和0Day保险，请使用Qubes。这款操作系统并不帮你隐藏身份，它的设计理念就是默认你所有的程序都具有潜在威胁。每个程序都会在一个虚拟机中运行，类似于程序沙盘，所以就算一个程序遭到攻击也不会泄漏你所有的信息。只要配置一下Tor做个全局代理就能变成比Tails更加牢固的防御系统了。

和Whonix一样。。。估计不连代理是不能连tor的。我朝网监已经破解了这个系统建议不要用



- [Tor](https://zh.wikipedia.org/wiki/Tor) 包括： Stream isolation，regular，obfs2，obfs3，obfs4和ScrambleSuit bridges support，[Vidalia](https://zh.wikipedia.org/wiki/Vidalia)图形前端（它的入门教程在“[这里](https://program-think.blogspot.com/2009/09/break-through-gfw-with-tor.html)”，FAQ在“[这里](https://program-think.blogspot.com/2013/11/tor-faq.html)”）
- [NetworkManager](https://zh.wikipedia.org/wiki/NetworkManager) 用来进行简单的网络配置
- [Tor Browser](https://zh.wikipedia.org/wiki/Tor) 是一个基于[火狐](https://zh.wikipedia.org/wiki/%E7%81%AB%E7%8B%90)进行改造以保护匿名性的Web浏览器，其含有的Torbutton用来保持匿名和对抗[JavaScript](https://zh.wikipedia.org/wiki/JavaScript)，默认使所有的cookies被欺骗为临时cookies；[HTTPS Everywhere](https://zh.wikipedia.org/wiki/HTTPS_Everywhere) 用来启用针对于大部分网站的[Transport Layer Security](https://zh.wikipedia.org/wiki/Transport_Layer_Security)，并通过[NoScript](https://zh.wikipedia.org/wiki/NoScript)来限制JavaScript，[Adblock Plus](https://zh.wikipedia.org/wiki/Adblock_Plus)用来移除广告
- [Pidgin](https://zh.wikipedia.org/wiki/Pidgin) 使用[OTR协议](https://zh.wikipedia.org/wiki/Off-the-Record_Messaging)进行端对端加密[即时通讯](https://zh.wikipedia.org/wiki/%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF)
- [Claws Mail](https://zh.wikipedia.org/wiki/Claws_Mail) 带有[GnuPG](https://zh.wikipedia.org/wiki/GnuPG)支持的电子邮件客户端
- Icedove ([Thunderbird](https://zh.wikipedia.org/wiki/Mozilla_Thunderbird)) 带有基于[OpenPGP](https://zh.wikipedia.org/wiki/OpenPGP)支持的[Enigmail](https://zh.wikipedia.org/wiki/Enigmail)的电子邮件客户端
- [Liferea](https://zh.wikipedia.org/wiki/Liferea) [聚合器](https://zh.wikipedia.org/wiki/%E8%81%9A%E5%90%88%E5%99%A8)
- [Gobby](https://zh.wikipedia.org/w/index.php?title=Gobby&action=edit&redlink=1) 用来进行文本[协同写作](https://zh.wikipedia.org/wiki/%E5%8D%94%E5%90%8C%E5%AF%AB%E4%BD%9C)
- [Aircrack-NG](https://zh.wikipedia.org/wiki/Aircrack-NG) 用来对[Wi-Fi](https://zh.wikipedia.org/wiki/Wi-Fi)网络进行[信息安全审计](https://zh.wikipedia.org/w/index.php?title=%E4%BF%A1%E6%81%AF%E5%AE%89%E5%85%A8%E5%AE%A1%E8%AE%A1&action=edit&redlink=1)
- [I2P](https://zh.wikipedia.org/wiki/I2P) 匿名网络（它的入门教程在“[这里](https://program-think.blogspot.com/2012/06/gfw-i2p.html)”）
- Electrum 易于使用的[bitcoin](https://zh.wikipedia.org/wiki/Bitcoin)客户端

加密和隐私

- [LUKS](https://zh.wikipedia.org/w/index.php?title=LUKS&action=edit&redlink=1)和[GNOME Disks](https://zh.wikipedia.org/w/index.php?title=GNOME_Disks&action=edit&redlink=1) 用来安装和使用加密存储设备，比如[优盘](https://zh.wikipedia.org/wiki/%E5%84%AA%E7%9B%A4)等
- [GnuPG](https://zh.wikipedia.org/wiki/GnuPG) 用来对电子邮电和数据进行加密和签名的[OpenPGP](https://zh.wikipedia.org/wiki/OpenPGP)的GNU分发版
- Monkeysign 用来对OpenPGP密钥进行签名和交换
- PWGen 一个强[随机密码产生器](https://zh.wikipedia.org/w/index.php?title=%E9%9A%8F%E6%9C%BA%E5%AF%86%E7%A0%81%E4%BA%A7%E7%94%9F%E5%99%A8&action=edit&redlink=1)
- [Shamir's Secret Sharing](https://zh.wikipedia.org/w/index.php?title=Shamir%27s_Secret_Sharing&action=edit&redlink=1) 使用gfshare和ssss
- Florence[虚拟键盘](https://zh.wikipedia.org/wiki/%E8%99%9A%E6%8B%9F%E9%94%AE%E7%9B%98) 用来针对[键盘监听](https://zh.wikipedia.org/wiki/%E9%94%AE%E7%9B%98%E7%9B%91%E5%90%AC)的一种对策
- MAT 用来匿名化文件的元数据
- [KeePassX](https://zh.wikipedia.org/w/index.php?title=KeePassX&action=edit&redlink=1) 密码管理器
- GtkHash 用来计算校验和
- Keyringer 用来通过[Git](https://zh.wikipedia.org/wiki/Git)进行秘密共享加密的命令行工具
- Paperkey 用来将OpenPGP的安全密钥记录在纸上的命令行工具
- TrueCrypt——跨平台的磁盘加密工具（俺写的系列教程在“[这里](https://program-think.blogspot.com/2011/05/recommend-truecrypt.html)”，共5篇）





## kalilinux









## linux Mint

`Linux Mint` 默认桌面环境是**Cinnamon**，一款由**GNOME**衍生出的并专为 `Linux Mint` 打造的桌面环境，拥有的清晰、简洁的外观设计，提供桌面部件功能，程序加载与显示速度令人满意，功能配置图形化并具有相当灵活性，整体表现相当稳定。许多用户甚至因为**Cinnamon**而选择 `Linux Mint`。当然，如果你不喜欢**Cinnamon**, `Linux Mint` 还提供基于**MATE**、**Xfce**、**KDE**桌面环境的发行版本，如果你对这些没有一个清晰的概念或者感受的话，我建议还是使用 `Linux Mint` 推荐的**Cinnamon**，毕竟它是专为 `Linux Mint` 定制开发的，其稳定性肯定更有保障。



## Ubuntu

得益于 Canonical 和开源社区的支持,Debian 衍生物在云和服务器应用程序中被大量使用。它也有多种风格来满足人们的不同需求。LTS 版本提供了良好的体验，可以快速解决问题。 Ubuntu 也支持流行的 .deb 包管理系统。







## debian

	`Linux` 世界赫赫有名的**APT**与**dpkg**也是它提供的。`Debian` 的发行及其软件源有五个分支：旧稳定分支（oldstable）、稳定分支（stable）、测试分支（testing）、不稳定分支（unstable）、实验分支（experimental）。目前稳定分支的代号**Jessie** 
	
	Debian 含有大量的软件包，提供良好的稳定性和大量的教程，帮助开发人员解决问题。Debian 测试分支，它有所有最新的软件，并且非常稳定。适合高级程序员和系统管理员。Debian 有很多开源库，另外，它的 .deb 软件包管理也是值得推荐的一点。

## Fedora CentOS 

	Fedora 的赞助商是红帽公司，以提供 Linux 桌面世界最尖端的功能而闻名。其最新版本经常激发其他 Linux 发行版采用新功能并进行更改。其智能的自动配置和更新的软件包使其成为开发人员完美的编程操作系统。Fedora 只包含开源组件，适合开源爱好者。即使貌似 Linus Torvalds 就很喜欢 Fedora。
	
	由于它是从RHEL源代码编译的，所以为 RHEL 构建的大多数商业软件都可以在 CentOS 上运行。它的安装和设置过程几乎就像 Fedora 一样。 它大量的红帽软件集合和 CentOS 存储库能满足不同的软件需求。 它允许使用 Xen 虚拟化来开发应用程序。 它使用 YUM 进行包管理。

## Gentoo 

它可以完成源代码的所有工作，还可以为特定的 CPU 架构重建整个系统。它对技术要求较高，但是时间越久，钻得越深，懂得越多。Gentoo 安装时候有些折腾，但你得到的是一个超级稳定的系统。



## archlinux Manjaro

### archlinux

	Arch Linux 是高度可定制的。Arch Linux 是硬核 Linux 爱好者最喜爱的 Linux 发行版，它随附有 Linux 内核和软件包管理器。如果需要做一些渗透测试工作，可以将 Arch Linux 安装转换成 BlackArch 安装。
	
	它最大特点就是滚动升级，完全靠网络分发，不像其他常见的发行版或基于 `Debian`，或基于 `Red Hat`，它完全是一个独立的发行版本。`Arch Linux` 从安装就是一个定制构建的过程，所有的安装配置基本都是基于命令，对使用者的要求很高。它以**pacman**进行包管理，所有软件安装，包括内核，都提供最新的，软件源也可以根据自己的需要灵活配置，可以这么说，它是一个为“极客而生”的系统。**AUR**软件仓库可以称得上是世界上最全的 `Linux` 软件源，有丰富的**Wiki**资料与活跃的社区，只要你是真正热爱它，并真喜欢折腾，那么肯定能找到不少志同道合者。[以官方Wiki的方式安装ArchLinux](https://www.viseator.com/2017/05/17/arch_install/) ，[ArchLinux的安装步骤](https://wiki.archlinux.org/index.php/Installation_guide_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87)) 。

```
# ls /sys/firmware/efi/efivars
# loadkeys layout
# ping -c 3 archlinux.org
# timedatectl set-ntp true
# fdisk -l
# mkfs.ext4 /dev/sda1
# mount /dev/sda1 /mnt
# nano /etc/pacman.d/mirrorlist
# pacstrap /mnt
# genfstab -U /mnt >> /mnt/etc/fstab
# arch-chroot /mnt /bin/bash
# ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# hwclock --systohc --utc
# nano /etc/locale.gen
# locale-gen
# echo LANG=en_US.UTF-8 > /etc/locale.conf
# echo myhostname > /etc/hostname
# nano /etc/hosts
# mkinitcpio -p linux
# passwd
# grub-install --target=i386-pc /dev/sdx
# grub-mkconfig -o /boot/grub/grub.cfg
# umount -R /mnt
```

### Manjaro

	[manjaro介绍](https://www.manjaro.cn/) 如果您知道如何设置 Arch Linux 系统，则不需要 Manjaro Linux。但是，如果你想要一个基于 Arch 的系统，而不需要考虑更新，你可以试试 Manjaro。安装非常简单，就像 Ubuntu 或 Linux Mint 一样。

[DistroWatch排名及信息更新](https://distrowatch.com/table.php?distribution=manjaro) 

**《[一张列表展示ArchLinux系软件有多丰富——看哭百万Debian、RedHat系同学](https://www.lulinux.com/archives/2787)》** ,[欢迎使用manjaro](https://www.manjaro.cn/) 

1. **安装系统**：有图形化界面，一路点击下一步即可（**节约8到无限个工时**【以archlinux或gentoo为负面参照物】）
2. **硬件驱动**：安装之后，不会出现硬件没响应的问题，比如说连网都上不了、X桌面无法进入**（节约1到无限个工时）**
3. **基础使用**：安装之后马上就可以干最基本的工作，如打开ntfs分区、听mp3、看flash视频等**（节约1到无限个工时）**
4. **高级使用**：想要的软件都能通过软件中心找到，QQ有，网易云音乐有，teamviewer有，skype有，wps有，sublime有，phpstorm有……**（节约无限个工时）**
5. **界面体验**：为懒人塑造了人性化体验，如关闭窗口只要鼠标点击就行、有“开始”菜单，可以自动平铺窗口等等**（节约无限个工时）**
6. **稳定性**：开启和操作软件都很快，不会出现资源高耗、卡死、崩溃的bug**（节约无限个工时）**

Manjaro 在桌面环境上有着许多选择，比如 Xfce、KDE、Deepin、BspWM、Budgie、i3、LXDE、Cinnamon、Enlightenment、Netbook、Fluxbox、Gnome、JWM、LXQT、MATE、Openbox 和 PekWM。所有这些桌面环境在 Manjaro 中都十分漂亮。Manjaro 官方的桌面环境是 Xfce 和 KDE，而其他桌面环境则是社区支持的。

Manjaro 则使用 Octopi 来使其变得更加用户友好，Octopi 是一个用 Qt 编写的 pacman 的图形前端。Manjaro 很好的维护了他自己的库，这也是它的一个优势。



A 64 bit installation of Manjaro running KDE uses about 455MB of memory.

###### Install a basic KDE Plasma environment

```
sudo pacman -S plasma kio-extras
```

systemd-analyze命令用于诊断系统启动时间

systemd-analyze blame	#按时间排序，查看服务启动耗费时间

systemd-analyze critical-chain updatedb.service	#查看关联性服务启动耗费时间

sudo systemctl list-unit-files --state=enabled	#查看已经启用的服务

sudo systemctl disable xxxx.service	#关闭服务自启动

sudo systemctl mask xxxx.service	#将服务启动文件重定向到/dev/null，一般用于static类型的服务，该条命令谨慎使用，除非你确认其相关的服务真的都不需要启动



更换桌面系统

1、安装dde

​	sudo pacman -S deepin deepin-extra

2、修改 /etc/lightdm/lightdm.conf

​	sudo cp /etc/lightdm/lightdm.conf /etc/lightdm/lightdm.conf.bak

​	sudo sed -i 's/greeter-session=lightdm-.*/greeter-session=lightdm-deepin-greeter/g' /etc/lightdm/lightdm.conf

​	 sudo sed -i 's/user-session=xfce/user-session=deepin/g'  /etc/lightdm/lightdm.conf

3、选择桌面、重启

​	锁定——选择deepin桌面图标（一般在右下角）——重启系统就ok了。

DockBarX XFCE

plank，docky，cairo-dock，tint2，AWN，lattedock，glx-dock

# 选择

## 桌面

默认桌面选择xfce4、lxde、mate甚至仿制windows界面都是不错的。

## 易安装性

拿archlinux为反面典型，虽然其性能高可以节省工作时间，但是如果安装它都要从头开始学习ABC，那价值就大打折扣。就安装系统的便捷性来说，archlinux、gentoo、lfs这样的系统真没必要尝试。

## 软件包

centos、slackware这些软件奇缺或者需要非常复杂的途径才能找到安装源的系统就没必要尝试了。ubuntu系软件比较丰富，但QQ/TIM这样的基本软件还需要折腾一番，还未必能稳定使用。[archlinux系软件包异常丰富](https://www.lulinux.com/archives/2787)，无情碾压deb和rpm系诸多发行版，例如manjaro下可以一条命令安装好无比稳定、功能全面的deepinwine-tim或deepinwine-qq。





# distrowatch操作系统评论网站

https://distrowatch.com/



https://git.kernel.org/pub/scm/linux/kernel/git/





# manjarolinux使用

## 1.安装

安装过程（Windows10 与 Linux 共存）

1. 将系统镜像写入U盘，以 DD 模式写入，或者直接用 ImageWrite 写入，千万不要 syslinux。
2. 由于 Windows10 本身的引导分区就是 ESP 分区，所以到分区的界面，只需要将 /boot/efi 挂载到 ESP 分区即可。
3. 以上操作之后，系统应该正常安装结束，重启之后会直接进入 Xfce 桌面。

```
1. 将源排序，用速度最快的
sudo pacman-mirrors -g或者sudo pacman-mirrors -i -c China -m rank
2. 添加 Arch 源（需要编辑 /etc/pacman.conf ）
[archlinuxcn]
SigLevel = Optional TrustedOnly
Server = https://mirrors.ustc.edu.cn/archlinuxcn/$arch
3.执行
sudo pacman -Syy && sudo pacman -S archlinuxcn-keyring
4. 安装搜狗输入法
sudo pacman -S fcitx-sogoupinyin
sudo pacman -S fcitx-im 
sudo pacman -S fcitx-configtool
~/.profile 这个文件或者在/etc/environment
export GTK_IM_MODULE=fcitx
export QT_IM_MODULE=fcitx
export XMODIFIERS="@im=fcitx"
5. 输入 fcitx 命令，重启计算机
6.美化 Xfce
使用了 OSX 风格的主题 OSX-Arc-White下载好主题之后，直接解压到  /usr/share/themes 下，就可以在 xfce 自带的窗口管理配置了
7.使用yaourt从AUR安装软件
pacman -S base-devel yaourt
yaourt -S vim
8.pacman基本用法
# pacman -S package_name1 package_name2 ...     # 安装软件
# pacman -R package_name                        # 删除软件
# pacman -Syu                                   # 更新软件和系统
$ pacman -Ss string1 string2 ...                # 搜索
sudo pacman -Syyu //更新系统软件包
9.安装软件
$ sudo pacman -S vim                    # 安装vim编辑器
$ sudo pacman -S netease-cloud-music    # 安装网易云音乐，需archlinuxcn源
$ sudo pacman -S google-chrome          # 安装Chrome，需archlinuxcn源
$ sudo pacman -S visual-studio-code-bin # 安装VSCode，需archlinuxcn源
10.git
sudo pacman -S git
git config --global user.name "github昵称" 
git config --global user.email "注册邮箱" 
11.代理
shadowsocks
浏览器switchy

```



使用全局菜单。包括 GTK，QT 的程序以及像 Firefox 这个菜单非基于 GTK/QT 的，都已经成功实现了全局菜单了

基于 Wine 较完美运行微信，QQ 以及迅雷（迅雷是为了解决 ED2K 的下载）

系统功能齐全，工作正常，包括蓝牙，MTP 等

Shadowsocks / Redsocks / Iptables / PcapDNSProxy

```bash
1.Shadowsocks / Redsocks / Iptables / PcapDNSProxy
yaourt -S shadowsocks redsocks-git pcap-dnsproxy-git
实现网络层（TCP）的代理，国内流量（中国 IP 地址白名单）不走代理直接请求，其余流量通过 Iptables 转发给 redsocks 并进一步转发给 shadowsocks。Pcap-DNSProxy 则是用来解决 DNS 污染的。
后面两个带 -git 的其实是 AUR 的包，如果下载缓慢或者由于网络问题经常出错，可以先安装 proxychains-ng 用代理去下载安装。

2.dropbox
yaourt -S dropbox
KDE 可以用 kio-gdrive，但是我 KDE 的 Google 帐号一直没法登上去。

3.zsh / oh-my-zsh   终端就这一套了
yaourt -S zsh oh-my-zsh-git

4.wqy-microhei / wqy-zenhei
yaourt -S wqy-microhei wqy-zenhei
Manjaro 字体已经处理的不错了，但是还是装多这两个字体吧

5.studio-3t / dbeaver / medis
yaourt -S studio-3t dbeaver medis
分别是 MongoDB，MySQL 等关系型数据库，Redis 的管理 GUI。Medis 是基于 Electron 构建的，我之前一直用的是 fastonsql，但是 Manjaro 搜不到该包好像。

6.typora
yaourt -S typora

7.wireshark-qt / telegram-desktop / netease-cloud-music / firefox / chromium / vscode / virtualbox / wiznote
yaourt -S wireshark-qt telegram-desktop netease-cloud-music firefox chromium visual-stuido-code-bin virtualbox wiznote

8.enpass
yaourt -S enpass
密码管理器。为什么不是 lastpass，onepass 等，因为他们都没有 Linux 客户端呀

9.eepin-wechat / deepin-wine-tim / deepin-wine-thunderspeed
yaourt -S deepin-wechat deepin-wine-tim deepin-wine-thunderspeed
这三个就屌了，在 Linux 运行 Windows 版本的微信，TIM 以及迅雷。都比较完美了

10.docker / docker-compose / kitematic
yaourt -S docker docker-compose kitematic
本地使用 MySQL，Redis 之类的我一般都不直接装，而是通过 Docker。Kitematic 是 Docker 应用管理，界面不错，也挺好用的。

11.mailspring / nylas
yaourt -S mailspring
邮件管理器。界面很美观，功能也强大。

12.gitkraken
yaourt -S gitkraken
管理 Git 的，类似 SourceTree。基于 Electron。

13.aria2-fast
yaourt -S aria2-fast
突破线程限制的 Aria2。配合 Firefox 的 Aria2 插件就很完美了

14.deluge / uget
yaourt -S deluge uget
前者是下 BT 的，后者常规下载（支持 Aria2）。我其实现在都没装，我都用 Aria2，Firefox 装一个 Aria2 的插件来下载。然后 ED2K 就用 WIne 迅雷。

15.kio-gdrive
yaourt -S kio-gdrive
让 dolphin 支持 Google Drive，经测试 manjaro stable 版本添加帐号报错了，testing 反倒没问题。

16.eepin-screenshot
yaourt -S deepin-screenshot
Deepin 截图炒鸡好用。可以在设置里面添加下快捷键配置。这个截图和 QQ 截图类似都可以检测应用边框。不过可惜的一点是没有上传到网络的功能，GNOME 自带的截图我记得是可以自动上传图床的。

17.latte-dock
yaourt -S latte-dock
如果想要 Dock 栏的，非安装这个软件不可了。

18.Redshift
yaourt -S redshift
对自己眼睛好一点，我安装了一个 KDE 插件叫 Redshift-Control，他依赖这个，没装的话运行不了。作用就是调节色温了，晚上屏幕暖一点，不那么刺眼。

19.ssh
pacman -Syy openssh
systemctl start sshd
systemctl enable sshd.service

20.ifconfig，ip
pacman -S net-tools
pacman -S net-tools dnsutils inetutils iproute2
ifconfig、route、arp和netstat等命令行工具（统称为net-tools）
net-tools通过procfs(/proc)和ioctl系统调用去访问和改变内核网络配置，而iproute2则通过netlink套接字接口与内核通讯。抛开性能而言，iproute2的用户接口比net-tools显得更加直观。比如，各种网络资源（如link、IP地址、路由和隧道等）均使用合适的对象抽象去定义，使得用户可使用一致的语法去管理不同的对象。更重要的是，到目前为止，iproute2仍处在持续开发中。

21.终端分屏工具tmux
```





## 2.配置

[FSearch——适用于Linux的快速独立搜索工具](https://www.sysgeek.cn/fsearch/) 



[TCP Optimizer](https://www.speedguide.net/downloads.php) 是一款绿色小工具，它可以帮助新手或高级用户调整 Windows 系统中相关的 TCP/IP 参数，从而可以轻松地将系统调整为适用所使用的 Internet 连接类型。



# 双系统安装



## 1.关闭快速启动和安全启动 

### **关闭Windows的快速启动** 

“电源选项”，“选择电源按钮的功能”，“更改当前不可用的设置”，“快速启动”

	快速启动是 Windows 8 (及更新的版本) 中的一项功能,通过休眠来提高启动速度.但是如果你休眠 Windows 然后进入另一个系统修改文件,可能会造成数据丢失.即使你不打算在双系统中共享文件,这也容易损坏 EFI 系统分区.因此你应该在安装 Linux 前禁用快速启动

### **关闭Security Boot** 

进入BIOS（本人联想台式F1）Security --- Security Boot ---- Enabled改为Disabled，F10保存退出。

要求人们继续支持反Secure Boot垄断，这个呼吁很重要。如果我们不支持，未来就无法自由地使用硬件、安装自己想要的软件。

自从个人电脑诞生后，就一直如此。过去30年我们都在使用类似上图的画面，设置硬件参数。不用说，BIOS已经变得日益不适用了。

1998年，Intel牵头，联合AMD、AMI、Apple、Dell、HP、IBM、Lenovo、Microsoft和Phoenix等业界主要厂商，开始制定新一代BIOS。这个项目叫做"统一的可扩展固定接口"（Unified Extensible Firmware Interface），简称[UEFI](http://en.wikipedia.org/wiki/Unified_Extensible_Firmware_Interface)。2005年推出1.1版。电脑运行的将不是BIOS，而是UEFI BIOS。等它运行结束，再载入操作系统。微软感兴趣的不是整个UEFI，而是UEFI的一个子规格Secure Boot。它要强行部署Secure Boot。

	Secure Boot的目的，是防止恶意软件侵入。它的做法就是采用密钥。UEFI规定，主板出厂的时候，可以内置一些可靠的公钥。然后，任何想要在这块主板上加载的操作系统或者硬件驱动程序，都必须通过这些公钥的认证。也就是说，这些软件必须用对应的私钥签署过，否则主板拒绝加载。由于恶意软件不可能通过认证，因此就没有办法感染Boot。

这个设想是好的。但是，UEFI没规定哪些公钥是可靠的，也没规定谁负责颁发这些公钥，都留给硬件厂商自己决定。现在，微软就是要求，主板厂商内置Windows 8的公钥。费者购买一台预装Windows 8的台式机或笔记本，想要在上面再安装其他操作系统（包括以前版本的Windows）是不可能的，除非关闭Secure Boot，或者其他操作系统能够通过Windows 8公钥的认证。如果选择关闭Secure Root，那么预装的Windows 8将无法使用，需要重新安装。

Secure Boot对移动设备的影响，比PC还要严重。微软的平板电脑Surface RT的Secure Boot是打开的，没法关闭。使用Windows Phone 8操作系统的智能手机也将采用这种做法。那么，用户也就不可能在Windows Phone上安装其他操作系统了。除了微软公司，苹果公司也有这种倾向。在新一代的iPhone和iPad上面安装其他操作系统，似乎是不可能的。

### rootkit 

不同类型的 rootkit 会在 PC 启动过程的不同阶段加载：

- **Firmware rootkit** 会覆盖 PC 的基本输入/输出系统或其他硬件的固件，以便 rootkit 能够在 Windows 之前启动。
- **Bootkit** 可以取代操作系统的引导加载程序，以便 PC 在操作系统之前加载bootkit。
- **Kernel rootkit** 可以替换操作系统内核的一部分，以便 rootkit 可以在操作系统加载时自动启动。
- **Driver rootkit** 会伪装是 Windows 用来与 PC 硬件进行通信的值得信赖的驱动程序之一。



### *UEFI模式下，硬盘应该是GPT分区形式*  

“msinfo32”，如果显示的是“传统”，即为BIOS启动方式；如果显示的是“UEFI”，则为UEFI启动方式。

如果你知晓GPT与UEFI启动之间的关系，那么你可能知道，Windows想要从GPT硬盘引导，就必须以UEFI方式启动(反之则不成立)。对于 UEFI-GPT 模式,这里的值为 "UEFI",对于 BIOS-MBR 模式,这里的值为 "Legacy"。

| Windows 版本  | BIOS-MBR            | UEFI-GPT       |        |        |
| ------------- | ------------------- | -------------- | ------ | ------ |
| x86 (IA32)    | x86_64              |                |        |        |
| Windows XP    | x86                 | 支持（大多数） | 不支持 | 不支持 |
| x64           |                     |                |        |        |
| Windows Vista | x86                 |                |        |        |
| x64           | 带有 SP1 的版本支持 |                |        |        |
| Windows 7     | x86                 | 不支持         |        |        |
| x64           | 支持                |                |        |        |
| Windows 8/8.1 | x86                 | 部分支持       | 不支持 |        |
| x64           | 不支持              | 支持           |        |        |
| Windows 10    | x86                 | 部分支持       | 不支持 |        |
| x64           | 不支持              | 支持           |        |        |

### Windows 中的文件名限制

	Windows 中的文件名长度最长为260字符。这是 Windows 而非 NTFS 的限制,因此其它系统依然可以在 NTFS 文件系统中创建文件名中带有这些符号的文件,但 Windows 无法识别这些文件. 运行 `chkdsk` 时也可能删除它们,这是一个引发数据丢失的风险.

## 2.安装

推荐先安装 Windows 再安装 Linux , 安装 Windows 时只使用硬盘的部分空间建立需要的分区.在 Windows 安装完毕后,进入 Linux 安装环境中后你可以对硬盘未分配的空间进行分区而不改动 Windows 分区. 对于 UEFI 系统,可以使用 Windows 安装时建立的 EFI 系统分区.

### （1）bios系统

#### 使用 Linux 启动管理器

	你可以使用grub

#### 使用 Windows 启动管理器

	需要设置 Windows 启动管理器使它加载另一个启动管理器 (例如 GRUB ),然后启动 Arch Linux.





### （2）准备

rufus 2.18 或者usrwriter

manjaro 17.1.7.iso

按下开始请选择dd模式，manjaro必须要dd模式才可以安装

写好的U盘，多半是同时支持UEFI和BIOS启动的，所以我们要手动选择以那种方式方式启动，这里我们选择UEFI启动下面的U盘。然后就可以进入manjaro的安装了。



https://wiki.archlinux.org/index.php/Dual_boot_with_Windows_



root(20GB)、home(64GB)、swap(16GB)



掛載 /boot/efi ，找出一開始要記的 EFI 空間 (剛剛是 260MB)，我因為是不同顆硬碟，所以選擇另外一顆

找出剛剛的大小然後點選，然後點 "編輯"。內容選 "keep (保持)"，掛載點處選擇或輸入 "/boot/efi"，旗標要勾選 "boot (開機)"、"esp"



### （3）替换引导文件

先用 DG 打开 EFI 分区，你会看到多了一个文件夹，名称取决于你安装的是哪一个发行版。我安装的是 Manjaro Linux，名称就是 Manjaro，打开之后会发现里面有一个名为 grubx64.efi 的文件，这就是启动 Linux 的引导文件。和 Windows 10 的 bootmgfw.efi 类似，我们想要用 grubx64.efi 引导代替掉 bootmgfw.efi，这样就可以用 GRUB 引导了。步骤：

1. 进入管理员命令行。方法：win + x，再按 a
2. 输入 `bcdedit /set {bootmgr} path \EFI\Manjaro\grubx64.efi`。提示操作成功的话，就完成了。

> 注：经人提醒，如果输入以上命令提示「参数错误」的话，将 {bootmgr} 改为 ‘{bootmgr}’，原因是 PowerShell 和 CMD 语法的差别。

至此，如果你安装的是除 Arch 之外绝大多数发行版，那么接下来就和你没有啥关系了，你已经成功了，好好享受吧！

开机之后会发现进入 GRUB 的引导了，通常会包含至少三个选项（以 Manjaro 举例）：Manjaro、Manjaro 高级选项和 Windows Manager。这就代表你已经完美的解决了 Windows 和 Linux 双系统引导的问题。

### （4）修复 Windows 引导

这一点是我安装 Arch Llinux 的时候发现的，Arch Linux 安装过程是手动安装的，在编写 GRUB 的时候会扫描不到 Windows Manager 所在的分区（当然可能不是所有人都会遇到），所以在 GRUB 界面可能会看不到 Windows Manager 选项，导致进不去 Windows 10，这里就需要手动编辑 GRUB 信息，我们打开 `/boot/grub/grub.cfg` 文件，发现里面确实没有 Windows 10 的启动信息，在后面加上：

```
menuentry "Microsoft Windows 10" {
  insmod part_get
  insmod fat
  insmod search_fs_uuid
  insmod chain
  search --fs-uuid --set=root $hints_string $fs_uuid
  chainloader /EFI/Microsoft/Boot/bootmgfw.efi
}
```

**注意**：

这里的 `$hints_string`，代表的是终端执行命令：

```
sudo grub-probe --target=hints_string /boot/efi/EFI/Microsoft/Boot/bootmgfw.efi
```

后的输出；

而 `$fs_uuid` 代表的是：

```
sudo grub-probe --target=fs_uuid /boot/efi/EFI/Microsoft/Boot/bootmgfw.efi
```

的输出。

然后保存。在终端执行命令：`sudo grub-mkconfig -o /boot/grub/grub.cfg`，就 OK 了。

到此，Arch Linux 和 Windows 10 双系统也配置完毕了。

### （5）附加问题

在使用这一年多的时间，遇到了以下的几个问题：

1. 在 Windows 10 进行了一个大更新后，会发现 GRUB 引导界面没有了，还是直接进入了 Windows 10，这时只需要按照 `替换引导文件` 的方法重新输入一遍命令就行。
2. 使用 Linux 某个发行版一段时间之后，难免会想尝试一下另一个发行版。这时请务必将之前的发型版的引导文件删除，否则可能会出现无论怎么设置都无法进入 GRUB 的情况。例如：我之前用的是 Ubuntu，我现在换成了 Manjaro，我就需要用 DG 删除 EFI 分区的 Ubuntu 文件夹。
3. 在我使用 Manjaro 更新了一次 Linux 的内核后，进不去 Windows 10 了，这个时候千万不要直接修复 Windows 10 引导，这会格式化 EFI 分区，只需要按上面 [修复 Windows 引导](https://itswincer.com/posts/ad42f575/#%E4%BF%AE%E5%A4%8D-Windows-%E5%BC%95%E5%AF%BC) 的方法编辑一下 GRUB 就可以了。

选择我们建立的efi分区，然后在下面选择保留，标记选择boot和esp，挂载点选择/boot/efi 

分给linuxswap分区8G空间，就是传说中的虚拟存储啦；分给/home 和 / 主分区每个一半的剩余空间。我建议这几个分区先压缩出来，在进行分区安装。

其中linuxswap是在文件系统选的（可能记错？）；/home 和 / 都是挂载点。

### （6）解决方法

1.磁盘a已安装win10，磁盘b分三个区和另一个未分配的空间（166G），这个未分配的用来安装manjaro系统。

2.制作USB启动盘（manjaro），进去之后按正常安装，进到磁盘分区选择时，选择磁盘b，选择166G未分配的块，选择新建分区，ext4，然后挂载不选，确定；再在这个ext4盘上选择新建分区，后面就直接到用户名密码设置了，后面就简单，确认分区后就会安装，重启选择对于HDD盘启动即可。

（）文件互相读取

linux可以读取NTFS文件系统

windows不能读取ext4文件，可推荐Ext2Fsd，Ext2Read，Ext2Explore，DiskInternals Linux Reader





























