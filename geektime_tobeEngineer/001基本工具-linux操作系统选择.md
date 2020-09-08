# linux系统推荐

[TOC]

Which Linux distros are suitable for developers

https://fossbytes.com/best-linux-distros-for-programming-developers/

https://fossbytes.com/secure-linux-distros-privacy-anonymity/

[https://en.m3uiptv.com/](https://www.youtube.com/redirect?q=https%3A%2F%2Fen.m3uiptv.com%2F&event=video_description&v=B_vMWJ0tCHQ&redir_token=Uqwizl1GOpmmuSnHXRS1OSxm3kx8MTU0Njc2MjUyMUAxNTQ2Njc2MTIx) (VLC直播源)

# 系统选择

distrowatch操作系统评论网站

https://distrowatch.com/

**1.桌面选择**：默认桌面选择xfce4、lxde、mate甚至仿制windows界面都是不错的。

**2.易安装性**：拿archlinux为反面典型，虽然其性能高可以节省工作时间，但是如果安装它都要从头开始学习ABC，那价值就大打折扣。就安装系统的便捷性来说，archlinux、gentoo、lfs这样的系统真没必要尝试。

**3.软件包**：centos、slackware这些软件奇缺或者需要非常复杂的途径才能找到安装源的系统就没必要尝试了。ubuntu系软件比较丰富，但QQ/TIM这样的基本软件还需要折腾一番，还未必能稳定使用。[archlinux系软件包异常丰富](https://www.lulinux.com/archives/2787)，无情碾压deb和rpm系诸多发行版，例如manjaro下可以一条命令安装好无比稳定、功能全面的deepinwine-tim或deepinwine-qq。

## Tails

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



## archlinux

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

## Manjaro

	如果你想要一个基于 Arch 的系统，而不需要考虑更新，你可以试试 Manjaro。安装非常简单，就像 Ubuntu 或 Linux Mint 一样。

[DistroWatch排名及信息更新](https://distrowatch.com/table.php?distribution=manjaro) 

**《[一张列表展示ArchLinux系软件有多丰富——看哭百万Debian、RedHat系同学](https://www.lulinux.com/archives/2787)》** ,[欢迎使用manjaro](https://www.manjaro.cn/) 











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
~/.xprofile 这个文件或者在/etc/environment
export GTK_IM_MODULE=fcitx
export QT_IM_MODULE=fcitx
export XMODIFIERS="@im=fcitx"
5. 输入 fcitx 命令，重启计算机
6.美化 Xfce
使用了 OSX 风格的主题 OSX-Arc-White下载好主题之后，直接解压到  /usr/share/themes 下，就可以在 xfce 自带的窗口管理配置了
7.使用yaourt从AUR安装软件
sudo pacman -S base-devel yaourt
sudo yaourt -S vim
8.pacman基本用法
# pacman -S package_name1 package_name2 ...     # 安装软件
# pacman -R package_name                        # 删除软件
# pacman -Syu                                   # 更新软件和系统
$ pacman -Ss string1 string2 ...                # 搜索
sudo pacman -Syyu //更新系统软件包

yaourt回退软件版本

pacman -S downgrade
[archlinuxfr]
SigLevel = Never
Server = http://repo.archlinux.fr/$arch

sudo pacman -Sy更新仓库
```

使用全局菜单。包括 GTK，QT 的程序以及像 Firefox 这个菜单非基于 GTK/QT 的，都已经成功实现了全局菜单了

基于 Wine 较完美运行微信，QQ 以及迅雷（迅雷是为了解决 ED2K 的下载）

系统功能齐全，工作正常，包括蓝牙，MTP 等

```bash
4.wqy-microhei / wqy-zenhei
yaourt -S wqy-microhei wqy-zenhei
Manjaro 字体已经处理的不错了，但是还是装多这两个字体吧

8.enpass
yaourt -S enpass
密码管理器。为什么不是 lastpass，onepass 等，因为他们都没有 Linux 客户端呀

9.eepin-wechat / deepin-wine-tim / deepin-wine-thunderspeed
yaourt -S deepin-wechat deepin-wine-tim deepin-wine-thunderspeed
这三个就屌了，在 Linux 运行 Windows 版本的微信，TIM 以及迅雷。都比较完美了

11.mailspring / nylas
yaourt -S mailspring
邮件管理器。界面很美观，功能也强大。

13.aria2-fast
yaourt -S aria2-fast
突破线程限制的 Aria2。配合 Firefox 的 Aria2 插件就很完美了

14.deluge / uget
yaourt -S deluge uget
前者是下 BT 的，后者常规下载（支持 Aria2）。我其实现在都没装，我都用 Aria2，Firefox 装一个 Aria2 的插件来下载。然后 ED2K 就用 WIne 迅雷。

16.eepin-screenshot
yaourt -S deepin-screenshot
Deepin 截图炒鸡好用。可以在设置里面添加下快捷键配置。这个截图和 QQ 截图类似都可以检测应用边框。不过可惜的一点是没有上传到网络的功能，GNOME 自带的截图我记得是可以自动上传图床的。

17.latte-dock
yaourt -S latte-dock
如果想要 Dock 栏的，非安装这个软件不可了。

18.Redshift
yaourt -S redshift
对自己眼睛好一点，我安装了一个 KDE 插件叫 Redshift-Control，他依赖这个，没装的话运行不了。作用就是调节色温了，晚上屏幕暖一点，不那么刺眼。

20.ifconfig，ip
pacman -S net-tools
pacman -S net-tools dnsutils inetutils iproute2
ifconfig、route、arp和netstat等命令行工具（统称为net-tools）
net-tools通过procfs(/proc)和ioctl系统调用去访问和改变内核网络配置，而iproute2则通过netlink套接字接口与内核通讯。抛开性能而言，iproute2的用户接口比net-tools显得更加直观。比如，各种网络资源（如link、IP地址、路由和隧道等）均使用合适的对象抽象去定义，使得用户可使用一致的语法去管理不同的对象。更重要的是，到目前为止，iproute2仍处在持续开发中。

21.终端分屏工具tmux

screenfetch展示系统信息
xmind

22.
uget 媲美迅雷的下载工具
pacman -S file-roller unrar unzip p7zip
openssh

sudo vim /usr/share/applications/netease-cloud-music.desktop
23.
arch下安装deb
    yaourt -S debtap
    sudo debtap -u
    debtap quadrapassel_3.22.0-1.1_arm64.deb
24.aurman，yay，yaourt，pacman
Yaourt 已死！在 Arch 上使用这些替代品https://zhuanlan.zhihu.com/p/42287487
aurman 是最好的 AUR 助手之一，也能勝任 Yaourt 替代品。他對所有 pacman 的操作有着一樣的語法。你可以搜索 AUR，解決包依賴，安裝前檢查 PKGBUILD 的內容等等。 
aurman 的特性
    aurman 支持所有 pacman 操作並且引入了可靠的包依賴解決，衝突判定和分包（split package）支持
    分線程的 sudo 循環會在後來運行所以你每次安裝只需要輸入一次管理員密碼
    提供開發者包支持並且可以區分顯性安裝和隱性安裝的包
    支持搜索AUR
    你可以檢視並編輯 PKGBUILD 的內容
    可以用作單獨的 包依賴解決
安裝 aurman
git clone https://aur.archlinux.org/yay.git
git clone https://aur.archlinux.org/aurman.git
cd aurman
makepkg -si

用名字搜索：aurman -Ss <package-name>
安裝：aurman -S <package-name>

 yay 是我們列表上下一個選項。它使用 Go 語言寫成，宗旨是提供 pacman 的界面並且讓用户輸入最少化，yay 自己幾乎沒有任何依賴軟件。
yay 的特性
    yay 提供 AUR 表格補全並且從 ABS 或 AUR 下載 PKGBUILD
    支持收窄搜索，並且不需要引用 PKGBUILD 源
    yay 的二進制文檔除了 pacman 以外別無依賴
    提供先進的包依賴解決以及在編譯安裝之後移除編譯時的依賴
    支持日色彩輸出，使用 /etc/pacman.conf 文檔配置
    yay 可被配置成只支持 AUR 或者 repo 裏的軟件包

```

源帮助

https://mirrors.ustc.edu.cn/help/manjaro.html



## 2.配置

[TCP Optimizer](https://www.speedguide.net/downloads.php) 是一款绿色小工具，它可以帮助新手或高级用户调整 Windows 系统中相关的 TCP/IP 参数，从而可以轻松地将系统调整为适用所使用的 Internet 连接类型。



## 3.linux下安装程序（单独）

Intellij Idea最新版安装包  ideaIU-2018.1.1-no-jdk.tar.gz

tar -zxvf ideaIU-2018.1.1-no-jdk.tar.gz  

在 /usr/share/applications目录下创建idea.desktop文件

使用 vim /usr/share/applications/idea.desktop 修改文件中内容

[Desktop Entry]
 Version=1.0
 Type=Application
 Name=程序开发
 Icon=/home/shuai/deve/toolbox/intellij/idea-1703/bin/idea.png
 Exec="/home/shuai/deve/toolbox/intellij/idea-1703/bin/idea.sh" %f
 Comment=The Drive to Develop
 Categories=Development;IDE;
 Terminal=false
 StartupWMClass=jetbrains-idea

保存退出然后在启动器中已经可以看到idea图标了





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

1. **安装系统**：有图形化界面，一路点击下一步即可（**节约8到无限个工时**【以archlinux或gentoo为负面参照物】）
2. **硬件驱动**：安装之后，不会出现硬件没响应的问题，比如说连网都上不了、X桌面无法进入**（节约1到无限个工时）**
3. **基础使用**：安装之后马上就可以干最基本的工作，如打开ntfs分区、听mp3、看flash视频等**（节约1到无限个工时）**
4. **高级使用**：想要的软件都能通过软件中心找到，QQ有，网易云音乐有，teamviewer有，skype有，wps有，sublime有，phpstorm有……**（节约无限个工时）**
5. **界面体验**：为懒人塑造了人性化体验，如关闭窗口只要鼠标点击就行、有“开始”菜单，可以自动平铺窗口等等**（节约无限个工时）**
6. **稳定性**：开启和操作软件都很快，不会出现资源高耗、卡死、崩溃的bug**（节约无限个工时）**

Manjaro 在桌面环境上有着许多选择，比如 Xfce、KDE、Deepin、BspWM、Budgie、i3、LXDE、Cinnamon、Enlightenment、Netbook、Fluxbox、Gnome、JWM、LXQT、MATE、Openbox 和 PekWM。所有这些桌面环境在 Manjaro 中都十分漂亮。Manjaro 官方的桌面环境是 Xfce 和 KDE，而其他桌面环境则是社区支持的。

Manjaro 则使用 Octopi 来使其变得更加用户友好，Octopi 是一个用 Qt 编写的 pacman 的图形前端。Manjaro 很好的维护了他自己的库，这也是它的一个优势。



A 64 bit installation of Manjaro running KDE uses about 455MB of memory.

Install a basic KDE Plasma environment

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

```
sudo pacman -S deepin deepin-extra
```

2、修改 /etc/lightdm/lightdm.conf

```
sudo cp /etc/lightdm/lightdm.conf /etc/lightdm/lightdm.conf.bak

sudo sed -i 's/greeter-session=lightdm-.*/greeter-session=lightdm-deepin-greeter/g' /etc/lightdm/lightdm.conf

 sudo sed -i 's/user-session=xfce/user-session=deepin/g'  /etc/lightdm/lightdm.conf
```

3、选择桌面、重启

```
锁定——选择deepin桌面图标（一般在右下角）——重启系统就ok了。
```

DockBarX XFCE

plank，docky，cairo-dock，tint2，AWN，lattedock，glx-dock

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

### （7）文件互相读取

linux可以读取NTFS文件系统

windows不能读取ext4文件，可推荐Ext2Fsd，Ext2Read，Ext2Explore，DiskInternals Linux Reader

### （8）git补全



### （9）zsh使用



### (10)降级软件

```
sudo pacman -S downgrade
sudo downgrade typora
```







last reboot可以看到Linux系统历史启动的时间





浏览器内存占用率高

great suspender 插件，google浏览器高级设置中退出即关闭程序。

firefox

如何关闭plasma隐藏的程序

**一.解决CPU占用率高，打开网页停顿的问题：**

 工具--选项--内容--启用Java这一项（去掉前面的勾）--确认，重启即可。（当使用firefox火狐浏览器浏览时，可能会造成某些需使用java项无法正常使用等，如有问题重新勾选即可）



Firefox浏览器，在地址栏中输入about:config，在过滤器中输入browser.cache.memory.enable





ps -aux` 和 `ps -lax，pstree 可以以进程树的形式列出当前进程.xkill 是 X 环境的杀进程命令。只要输入 xkill，鼠标轻轻一点，鼠标指向的窗口（对应的进程）就立马被杀死.

**信号** 

Unix 和 Linux 有信号的概念。信号可以控制进程的运行。

如果想手动发送信号，可通过 kill 命令（因此 kill 不完全是一个杀进程的命令）。例如：

- `kill -STOP pid`：发送 SIGSTOP 信号，停止进程但不消灭进程。
- `kill -CONT pid`：发送 SIGCONT 信号，重新开始已停止的进程。
- `kill -KILL pid`：发送 SIGKILL 信号，强迫进程停止。这个命令可以用于对付无法正常结束的进程。



Firefox 的 Fission Memshrink （ [Project Fission](https://careers.mozilla.org/position/gh/1234648) 的一部分）虽然目的也是减少内存使用，但它可能更偏向于是一个清理过程，最终的 Project Fission 将真正提高网页的响应能力。而 Chrome 的 Page Lifecycle 也同样不意味着能即时优化，为了使其发挥最佳效果，Web 开发者需要提供支持，以便浏览器可以更智能地判断页面元素的优先级。





qemu kvm，调试内核环境



c/c++开发环境



## linux下工具

https://linuxtools-rst.readthedocs.io/zh_CN/latest/tool/top.html

https://briteming.blogspot.com/2017/04/awesome-linux.html

netease-cloud-music / firefox / chromium/telegram-desktop/google-chrome

### 播放器

mpv，mplayer，vlc，SMPlayer。可用于二次开发比较好的播放器有MPV、VLC。但首推VLC，因为技术比较成熟，有许多程序员将其作为开发对象，所以[适合做为二次开发](https://juejin.im/post/5b7bd7fde51d4538a423ba8f)。

**mpv快捷键** 

> - 鼠标左键双击: 进入/退出全屏
>
>   鼠标右键单击: 暂停/继续播放
>
>   p: 暂停/继续播放
>
>   m: 静音
>
>   方向键: 快进/快退（具体的自己试）
>
>   <>: 上一个/下一个（播放列表中）
>
>   Enter: 下一个（播放列表中）
>
>   q/Q: 停止播出并退出/保存当前播放进度并退出
>
> - j 循环选择字幕
>
> - J 反向循环选择字幕
>
> - [#] 循环切换音轨
>
> - f 切换全屏状态
>
> - T 切换视频窗口是否总在最前
>
> - s 视频截图，包含字幕
>
> - S 视频截图，不带字幕
>
> - Alt+s 自动逐帧视频截图，再按一次停止截图
>
> - . 进到下一帧
>
> - , 退到前一帧
>
> - 9 音量减 2
>
> - 0 音量加 2
>
> - [ 0.9091 倍慢速播放
>
> - ] 1.1 倍快速播放
>
> - { 0.5 倍慢速播放
>
> - } 2 倍快速播放
>
> - Backspace 重置为正常播放速度

**打开多个文件** 

Linux 和 OS: 选中多个文件，右键选中MPV打开

Linux: 终端执行命令： ~/.config/mpv/（常用配置）

**创建一个mpv.conf文件，这是主配置文件**

```
#没有边框
no-border
#记住断点
save-position-on-quit
```

**创建一个input.conf文件，这是自定义快捷键文件**

```
#向前滚轮音量加5
AXIS_UP volume 5
#向后滚轮音量减5
AXIS_DOWN  volume -5
```

脚本是mpv配置的重头戏。mpv启动时，保存在配置目录的scripts文件夹里的Lua脚本自动加载且执行。官方wiki里有脚本的[List](https://github.com/mpv-player/mpv/wiki/User-Scripts)，可以自行查阅，寻找自己喜欢的功能脚本。

分享一下我用的几个脚本：

[autoload.lua](https://github.com/mpv-player/mpv/blob/master/TOOLS/lua/autoload.lua)：自动加载当前播放文件目录里的视频文件到播放列表，也就是自动连播，原生mpv不具有这个功能。[mpv_thumbnail_script.lua](https://github.com/TheAMM/mpv_thumbnail_script)：显示预览缩略图



### shell

echo $SHELL，cat /etc/shells ，echo` `$0

若想要更改系统创建用户时的默认shell，可以使用在 /etc/default/useradd 目录下更改SHELL一行 
`SHELL=/bin/bash //可以更改此行来达到目的`

zsh / oh-my-zsh   终端就这一套了
yaourt -S zsh oh-my-zsh-git

chsh修改当前用户shell，如果你要换成 bash, 请输入 /bin/bash 并回车确认。vi /etc/passwd中修改/bin/bash

切换shell的命令如下：

```
`chsh -s /bin/zsh`
安装oh-my-zsh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
使用了默认的主题（robbyrussell）
切换主题
vi ~/.zshrc
ZSH_THEME="agnoster"
source ~/.zshrc
```

vim NERDTree zsh(steeef theme) + tmux

https://github.com/robbyrussell/oh-my-zsh/wiki/Themes



### 代理

shadowsocks：sudo pacman -S shadowsocks-qt5
浏览器switchy

1.Shadowsocks / Redsocks / Iptables / PcapDNSProxy
yaourt -S shadowsocks redsocks-git pcap-dnsproxy-git
实现网络层（TCP）的代理，国内流量（中国 IP 地址白名单）不走代理直接请求，其余流量通过 Iptables 转发给 redsocks 并进一步转发给 shadowsocks。Pcap-DNSProxy 则是用来解决 DNS 污染的。
后面两个带 -git 的其实是 AUR 的包，如果下载缓慢或者由于网络问题经常出错，可以先安装 proxychains-ng 用代理去下载安装。

```
`vim /etc/proxychains.conf`
```

curl ip.gs ；proxychains4 curl ip.cn;curl ip.cn

archlinuxcn/besttrace 1.2-2
​    IPIP.net 开发的加强版 traceroute，附带链路可视化

proxychains



### 邮件

Thunderbird

1.  使用户能够拥有个性化的电子邮件地址 

2.  一个单击通讯录 

3.  附件提醒 

4.  多频道聊天 

5.  标签和搜索 

6.  启用搜索网络 

7.  快速过滤器工具栏 

8.  消息存档 

9.  活动经理 

10.  大文件管理 

11.  安全功能，如网络钓鱼保护，无跟踪 

12.  自动更新等等


### 云盘

yaourt -S dropbox

kfilebox

KDE 可以用 kio-gdrive，但是我 KDE 的 Google 帐号一直没法登上去。

yaourt -S kio-gdrive
让 dolphin 支持 Google Drive，经测试 manjaro stable 版本添加帐号报错了，testing 反倒没问题。

proxychains dropbox start -i &

**nutstore**（坚果云，速度很快，按流量限制，用来存储文件挺好的）。

**mega**（mega.nz，国外的网盘，免费容量50g，速度还可以）。

用 Wuala 吧，据说是客户端本地加密，而且客户端也比 SpiderOak （安全性更好）好用。

Resilio Sync（速度快） , Dropbox, google drive

自己撘 seafile ，全平台同步，很爽的

**BitTorrent Sync** ，Syncthing是一种开源服务，可以替代BitTorrent Sync

### 下载工具

aria2是一款下载工具，它支持http(s)/ftp/BitTorrent/Metalink五种协议，aria2有强大的分块下载能力，它可以通过多个来源和多种协议下载同一个文件，让你的带宽爆满。aria2甚至可以同一时间使用http(s)/ftp/BitTorrent四种协议下载同一个文件（变态），此时他会把http(s)/ftp下载部分使用bt上传。当然，aria2仍然为你提供了前所未有的强大的常规http(s)/ftp下载性能，它提供了metalink协议的验证纠错功能。

aria2还自动验证通过BT协议下载的数据的正确性。

	*支持cookie，有些地址需要cookie才能下载，例如linuxsir的附件你单把地址复制到命令行用mytget下载是不行的。aria2能功过参数制定cookie文件，而且能自动载入firefox的cookie
	  *支持断电续传，之所以强，是他支持除了aria2自身产生的断点续传文件外，还支持浏览器的和wget产生的断点文件
	  *http代理，以及通过http代理的ftp代理？
	  *支持多线程，但默认为1，需要参数'-s'设定线程数，但用BT/metalink协议自动多线程。

使用：

	下载：aria2c http://www.kernel.org/pub/linux/kernel/v2.6/linux-2.6.22.6.tar.bz2
	
	分段下载：aria2c -s 2 http://www.kernel.org/pub/linux/kernel/v2.6/linux-2.6.22.6.tar.bz2(s 后面的参数值介于 1~5 之间)
	
	断点续传：aria2c -c http://www.kernel.org/pub/linux/kernel/v2.6/linux-2.6.22.6.tar.bz2
	
	下载torrent：aria2c -o gutsy.torrent <http://cdimage.ubuntu.com/daily-live/current/gutsy-desktop-i386.iso.torrent>

其配置文件位于~/.aria2/aria2.conf

设置上传和下载：aria2c --max-download-limit=100K --max-upload-limit=10K

同时下载两个BT文件aria2c -j2 file1.torrent file2.torrent

查看BT文件内的内容aria2c －S file1.torrent
选择下载文件那几个文件aria2c --select-file=1-4,8 file.torrent
给下载的文件设定路径并重命名aria2c --dir=/tmp --index-out=1=mydir/base.iso --index-out=2=dir/driver.iso file.torrent

aria2c magnet:?xt=urn:btih:AQWTG32X5UZQEV2JA7JHJMWWCJ2VUQBQ  -d  test/ 
-d  test/是指明下载文件保存在test目录。



先去“<https://github.com/ziahamza/webui-aria2>”下载最新的WebUI压缩包

文件夹上传到服务器或者复制到本机的HTML网页根目录下（如/var/www/html）

进入到/var/www/html目录“mv  ./webui-aria2  ./webui”重命名方便访问

“chmod 755 /var/www/html/webui-aria2”设置权限目录权限，让其有权限下载文件到本地

安装完成后，“aria2c  --enable-rpc  --rpc-listen-all”启用监听RPC

“systemctl  enable  httpd”、“systemctl  start  httpd”（SystemV管理方式）

 “/etc/init.d/httpd”（INIT）



http://IP地址/webui，就可访问成功；如果报错，在“设置”--> “服务器设置”中“主机：”后后面填写自己的IP地址就可；

如果Aria2.conf配置文件中启用了RPC安全认证，需要在WEB客户端设置中填入RPC用户名和密码，否则客户端报错！如果是自己用，个人觉得没必要用RPC认证，打开/etc/aria2.conf删除里面的RPC用户名和密码



sudo apt-get install aria2
sudo mkdir /etc/aria2    #新建文件夹  
sudo touch /etc/aria2/aria2.session    #新建session文件
sudo chmod 777 /etc/aria2/aria2.session    #设置aria2.session可写 

sudo mousepad /etc/aria2/aria2.conf    #创建配置文件

<http://aria2c.com/>

**配置** 

` ＝＝＝＝＝＝＝＝＝文件保存目录自行修改
dir=/home/username/Desktop/Downloads #username根据实际情况修改
disable-ipv6=true

#打开rpc的目的是为了给web管理端用
enable-rpc=true
rpc-allow-origin-all=true
rpc-listen-all=true
#rpc-listen-port=6800
#断点续传
continue=true
input-file=/etc/aria2/aria2.session
save-session=/etc/aria2/aria2.session

#最大同时下载任务数
max-concurrent-downloads=20

save-session-interval=120

Http/FTP 相关

connect-timeout=120
#lowest-speed-limit=10K
#同服务器连接数
max-connection-per-server=10
#max-file-not-found=2
#最小文件分片大小, 下载线程数上限取决于能分出多少片, 对于小文件重要
min-split-size=10M
#单文件最大线程数, 路由建议值: 5
split=10
check-certificate=false

http-no-cache=true

` 

aria2c --conf-path=/root/aria2/aria2.conf

aria2c --conf-path=/etc/aria2/aria2.conf





推荐两个：Aria2 Web UI或者YAAW

aria2+ariang+nginx linux 离线下载部署

AriaNg 是一个简单易用的linux web前台，可以提供图形化的aria控制管理界面(https://www.jianshu.com/p/8124b5b6ef95)



### 服务器

安装nginx

/bin/nginx 启用

nginx -s stop 停止

nginx -s reload 重启



```
server {
        listen 80;  #监听端口
        server_name 127.0.0.1 192.168.1.116;    #主机ip
        
        location / {
            root /home/x/aria2/AriaNg; #站点目录
            }
        }

```











### 笔记及markdown

yaourt -S typora

remarkable

wiznote

**nixnote**（evernote的第三方客户端，且不会占用免费用户的同步设备数量）。wine evernote

**calibre**（电子书的制作编辑管理格式转换等额等，我一般配合kindle用）

leanote

KindleNote

emacs+orgmode

zim wiki + dropbox

geeknote

linux下安装第三方evernote客户端everpad，并与中国版印象笔记进行同步的方法



### 文件搜索

[FSearch——适用于Linux的快速独立搜索工具](https://www.sysgeek.cn/fsearch/) 

angrysearch

FSearch文件搜索，ANGRYsearch，CatFish

### 专业软件

wireshark-qt ，vim，visual-studio-code-bin ，git，virtualbox，filezilla，gitkraken（类似 SourceTree。基于 Electron）

git config --global user.name "github昵称" 
git config --global user.email "注册邮箱" 

**ssh** 
pacman -Syy openssh
systemctl start sshd
systemctl enable sshd.service


10.docker / docker-compose / kitematic
yaourt -S docker docker-compose kitematic
本地使用 MySQL，Redis 之类的我一般都不直接装，而是通过 Docker。Kitematic 是 Docker 应用管理，界面不错，也挺好用的。
5.studio-3t / dbeaver / medis
yaourt -S studio-3t dbeaver medis
分别是 MongoDB，MySQL 等关系型数据库，Redis 的管理 GUI。Medis 是基于 Electron 构建的，我之前一直用的是 fastonsql，但是 Manjaro 搜不到该包好像。

### anki

插件

Zooming images 2.0 & 2.1 sans limitation 146%

可以比较方便地控制卡片字体大小而不用调整代码。安装后**按住ctrl键，上下滚动鼠标滚轮**。

**True Retention***推荐指数★★★★★*
 在统计中更准确更多样地反映学习情况



### kindle阅读相关

calibre

```
sudo -v && wget -nv -O- https://download.calibre-ebook.com/linux-installer.sh | sudo sh /dev/stdin

https://calibre-ebook.com/download_linux
```



### PDF Studio Viewer



### git使用

通过将Git配置变量 core.quotepath 设置为false，就可以解决中文文件名称在这些Git命令输出中的显示问题，

	$ git config --global core.quotepath false



Visual Studio Code，Eclipse – Egit，TortoiseGit，SourceTree是老牌的Git GUI管理工具，GitHub for Desktop，gitkraken



### geany



### VNC

**Remmina** 

tigervnc

VNC Viewer



### nmap

用namp对局域网扫描一遍，然后查看arp缓存表就可以知道局域内ip对应的mac了。namp比较强大也可以直接扫描mac地址和端口。执行扫描之后就可以 cat /proc/net/arp查看arp缓存表了。

nmap -sP 192.168.1.0/24　//进行ping扫描，打印出对扫描做出响应的主机.探测C段存活主机，可以用 |grep up 过滤存活的主机

nmap -sL 192.168.1.0/24　//仅列出指定网络上的每台主机，不发送任何报文到目标主机

nmap -PS 192.168.1.234　探测目标主机开放的端口，可以指定一个以逗号分隔的端口列表(如-PS  22，23，25，80)

nmap -PU 192.168.1.0/24	//使用UDP ping探测主机

nmap -sS 192.168.1.0/24	//使用频率最高的扫描选项（SYN扫描,又称为半开放扫描），它不打开一个完全的TCP连接，执行得很快.SYN扫描，指定IP范围指定端口,nmap -sS 1.1.1.1-30 -p 80

nmap -sV 1.1.1.1 -p 1-65535	//探测端口的服务和版本

nmap -O 1.1.1.1 或 nmap -A 1.1.1.1		//探测操作系统类型和版本

 https://blog.csdn.net/keepSmi1e/article/details/9370049

https://wizardforcel.gitbooks.io/nmap-man-page/content/10.html

https://nmap.org/man/zh/index.html#man-description



### 媒体播放器

（[煲机音乐](https://jingyan.baidu.com/article/3065b3b6848caebecff8a4e3.html)）

vlc C++

Mplay C语音（SMPlayer），mpv，baba ；PotPlayer韩国人

kodi，这个播放器与mpc-hc基本可以看做一家人,目前也就这个分支在继续开发,独立安装的播放器Kodi

https://github.com/taxigps/xbmc-addons-chinese/blob/master/plugin.video.asia-tv/Users_Guide.txt



### 安装软件工具

aurman，yay，yaourt，pacman
Yaourt 已死！在 Arch 上使用这些替代品https://zhuanlan.zhihu.com/p/42287487
aurman 是最好的 AUR 助手之一，也能勝任 Yaourt 替代品。他對所有 pacman 的操作有着一樣的語法。你可以搜索 AUR，解決包依賴，安裝前檢查 PKGBUILD 的內容等等。 
aurman 的特性
​    aurman 支持所有 pacman 操作並且引入了可靠的包依賴解決，衝突判定和分包（split package）支持
​    分線程的 sudo 循環會在後來運行所以你每次安裝只需要輸入一次管理員密碼
​    提供開發者包支持並且可以區分顯性安裝和隱性安裝的包
​    支持搜索AUR
​    你可以檢視並編輯 PKGBUILD 的內容
​    可以用作單獨的 包依賴解決
安裝 aurman
git clone https://aur.archlinux.org/yay.git
git clone https://aur.archlinux.org/aurman.git
cd aurman
makepkg -si

用名字搜索：aurman -Ss <package-name>
安裝：aurman -S <package-name>

 yay 是我們列表上下一個選項。它使用 Go 語言寫成，宗旨是提供 pacman 的界面並且讓用户輸入最少化，yay 自己幾乎沒有任何依賴軟件。
yay 的特性
​    yay 提供 AUR 表格補全並且從 ABS 或 AUR 下載 PKGBUILD
​    支持收窄搜索，並且不需要引用 PKGBUILD 源
​    yay 的二進制文檔除了 pacman 以外別無依賴
​    提供先進的包依賴解決以及在編譯安裝之後移除編譯時的依賴
​    支持日色彩輸出，使用 /etc/pacman.conf 文檔配置
​    yay 可被配置成只支持 AUR 或者 repo 裏的軟件包

### 翻译软件

stardict，自己安装词库，很强大的。开源（https://github.com/farseerfc/ydcv-rs）

 GoldenDict  这个软件是个翻译聚合。非常好用。可以自行添加词库，比如牛津八，柯林斯等词典

我自己用 python 写了一个插件，可以完美解决 Linux 屏幕取词问题。对于屏幕取词有需要，不妨试试。https://github.com/easeflyer/gd_plugin 

加载词典包：

编辑-->词典--> 词典来源 --> 词典文件所在目录

popup-dict安装

```
sudo pip install popupdict
```

[popup-dict - Linux 下的划词翻译工具](https://bianjp.com/posts/2017/12/14/announcing-popup-dict)  (https://github.com/bianjp/popup-dict)



### 蓝牙

```
sudo pacman -S  bluez bluez-utils
$ sudo systemctl enable bluetooth.service
$ sudo systemctl start bluetooth.service
 sudo modprobe btusb
 bluetoothctl
 
 $ sudo bluetoothctl
[bluetooth]# power on
[bluetooth]# agent on
[bluetooth]# default-agent
[bluetooth]# pairable on
[bluetooth]# discoverable on
[bluetooth]# scan on
[bluetooth]# devices
Device XX:XX:XX:XX:XX:XX Some device
[bluetooth]# trust XX:XX:XX:XX:XX:XX
[bluetooth]# connect XX:XX:XX:XX:XX:XX
Attempting to connect to XX:XX:XX:XX:XX:XX
[CHG] Device XX:XX:XX:XX:XX:XX Connected: yes
Request confirmation
[agent] Confirm passkey 360914 (yes/no): yes
[bluetooth]# pair XX:XX:XX:XX:XX:XX
https://docs.khadas.com/zh-cn/vim2/HowToSetupBluetooth.html

dmesg | grep -i Bluetooth
systemctl status bluetooth
rfkill list all
rfkill block bluetooth 
rfkill unblock bluetooth
sudo systemctl restart bluetooth

sudo wget -O /lib/firmware/brcm/BCM.hcd https://github.com/winterheart/broadcom-bt-firmware/blob/master/brcm/BCM43142A0-105b-e065.hcd

yaourt hcitool
hcitool dev

```

terminator

### 源码阅读工具

主要的交叉索引工具有：ctags、cscope、global、lxr、KScope、sourcenav、calltree、CodeViz、ncc、gprof等。

Global Global是GNU出品的交叉索引工具，能生成交叉索引的web页，很适合用来做程序的文档。

LXR：可能我分析得不深入，它的代码不是用C写的，是网页脚本，程序也很短，没有看到有应用数据库。很多的东西都是你浏览的时候才生成的，所以拿过来用或改造的潜力有限，而且代码分析的不够Global细腻。 

KScope KScope是cscope的图形前端

#### sourcenav snavigator

下载源码sudo ./configure –prefix=/home/h8yung/opt/sourcenav

sudo make，以及sudo make install

./snavigator



Graphviz是专门绘制有向图和无向图的工具

Doxygen是源码文档化工具，也能绘制调用图，它似乎是自己分析源码获得函数调用关系的

vim + ctags + cscope



### github插件

Octotree

Sourcegraph for GitHub

Insight.io

GitHub Hovercard：鼠标悬停看人资料

Awesome Autocomplete for GitHub：智能搜索插件

ZenHub for GitHub：持续集成，敏捷开发





### soho

方便监管。工业革命的产物，目的是让老板可控——信任

\- 协作互助。过去我们不谈「虚」的，要协作就要面对面，现在可以谈点「虚」的，我给你开发个软件，可以全程不见你——协作平台Phabricator——分解任务，Wiki共享，看板，这些功能让各个项目、项目中的细节和状态都一目了然。

\- 随时沟通。你就在那，离我不到1米的工位上，想找你我就喂一声，你自然就回去看了看我，于是就聊了10块的。——「平时日常沟通Bearychat」、「远程视频会议桌面共享软件ZOOM」、「紧急和重要情况下的Phone」



做和程序猿当前技能有关的SOHO工作。
（1）可以去国内外的知名外包网站接单，建议都去国外网站接单，比如elance,freelancer,guru。国内的taskcity和csto可以看看，猪八戒对于程序猿不是个好去处。
物以稀为贵，现在大家都蜂拥而上做手游，做app，其实可以去看一下elance这些网站的需求，虽然app开发需求多，但去抢单的更多，反倒是体感游戏开发或者其它目前大家关注较少但是也有需求的技能很吃香。
（2）针对海外市场做独立游戏开发，这个需要强大的个人技能（开发，策划，美术，市场）。不一定自己一个人做，可以跟人合作，甚至可以和老外合作。老外很容易接受跨国跨地区的线上合作，反倒是国人不一定能接受。
产品不一定针对iOS,Android市场，当别人都盯着这里的时候，你反倒可以试试做PC的，体感的，可以投放到steam等平台。
（3）针对海外市场做独立软件开发或者建个网站做SAAS服务。
很多老外独立软件开发者都是这样做的，特别是做SAAS服务比较靠谱。但前提是海外市场。
（4）如果是回老家，有一些本地资源关系，还可以做一些to B的网站或软件开发，但难度比较大，一个人会有问题。但如果是做一些管理维护或咨询工作，就相对简单些。
总的来说，不管是接外包项目，还是做独立开发者（游戏，软件，SAAS），还是利用关系做东西或者咨询，都要有个人品牌意识（固定的个人网站，博客，相对固定的产品和服务类型），不要打一枪换个地方。而且在某个领域做的久了，哪天浪潮来了就可以顺势而起，利用自己的累计和沉淀做一番事业。
多去尝试下海外市场吧，在这片神奇的土地上做SOHO比创业还难。

oracle，redhat，canonical 都有开源soho岗位。
还有37signal，github 之类。

答题前，我需要说明我不是程序员，我是SOHO，负责推广的
在我认识的SOHO程序员中，较为成功的（就是挣钱比较多），他们都有一个共同点：
**对推广以及产品都相当了解**
如果你在之前的公司工作了很久，但找不到技术和产品以及推广的某些联系，我不建议你SOHO，因为这样你最多只能写个APP发布到苹果应用商店，成功机会很低很低。
**我的建议是：有相当把握（在公司业余时间操作才算有把握），掌握了推广和产品的窍门后再SOHO，这样才能将你的技术价值最大化。当然你也可以找会推广的人合作，但自身要对推广有一点认识，才不会被猪队友坑！**

https://segmentfault.com/a/1190000000269342

https://www.sohotask.com/zh-cn/

https://www.douban.com/group/topic/84746155/



### 系统启动速度优化

Arch Linux 的 systemd-analyze 是个很不错的工具，利用它你可以很直观地观察到系统启动的时间都花到哪儿去了

```shell
systemd-analyze
```

```sh
systemd-analyze blame
```



### 安装mysql

```
sudo pacman -S mariadb mariadb-clients #安装MariaDb和其客户端工具,MariaDb默认的引擎还是Innodb
sudo pacman -S mariadb libmariadbclient mariadb-clients

#初始化MariaDb的数据目录
sudo mysql_install_db --user=mysql --basedir=/usr --datadir=/var/lib/mysql

sudo systemctl start mysqld #先启动MariaDb

mysqladmin -u root password '12345678' #为root用户设置一个新密码

mysql -u root -p 12345678 #尝试登录MariaDb，如果登录成功，说明配置完成了

sudo systemctl enable mysqld #MariaDb开机自动启动
```





图形化数据库管理工具

navicat(lite)

linux下破解https://github.com/cniPatch/NavicatPremium/blob/master/patch.sh

第一次执行start_navicat时，会在用户主目录下生成一个名为.navicat的隐藏文件夹。

—-把此文件夹删除后(删除文件夹命令是rm -rf .navicat)，下次启动navicat 会重新生成此文件夹，30天试用期会按新的时间开始计算。

等到期时找到这个system.reg文件删除，navicat又会重新计算过期时间，又可以用了，这个有点麻烦。

如果字体还不正常可以用gedit或vim打开start_navicat 文件将代码做如下更改

```
export LANG="en_US.UTF-8" 改为 ———> export LANG="zh_CN.UTF-8"
```



cp navicat.desktop /home/single/Desktop





https://www.dyxmq.cn/windows/software/navicate-premium-12-cracked.html

https://www.52pojie.cn/thread-867986-1-1.html





[phpMyAdmin](https://www.phpmyadmin.net/) 





Workbench

这是一个Sun Systems/Oracle开发的免费工具。对于Microsoft Windows, Mac OS X和Linux平台来说，Workbench都十分有用。





```text
export https_proxy=http://127.0.0.1:12333
export https_proxy="socks5://127.0.0.1:1080"
export ALL_PROXY=socks5://127.0.0.1:1080
source ~/.bashrc
```

[破解地址](https://github.com/DoubleLabyrinth/navicat-keygen/blob/windows/README_FOR_LINUX.zh-CN.md) 

NAVO-FOV4-YXU6-AJLM



### wine





