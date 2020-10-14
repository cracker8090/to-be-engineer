



准备

保证 Windows 10 至少为 Windows 10 Fall Creators Update 及之后的版本。

保证为 Windows 10 开启了「Windows Subsystem for Linux」的可选功能。

Windows 用户名不能有空格。（有空格会对 WSL 环境造成影响。）

Windows 用户名不能为中文。（两个系统下的编码格式不一样，会对 VSCode 的调试功能造成影响。）



# windows下samba连接阿里云

\# yum -y install samba

\# rpm -qa | grep samba

检查samba服务包的安装情况，会显示类似如下两个包：
samba-common-3.0.33-3.7.el5_3.1  //服务器和客户端均需要的文件
samba-3.0.33-3.7.el5_3.1        //服务器端文件

\# whereis samba

samba: /usr/lib64/samba /etc/samba /usr/share/man/man7/samba.7.gz



useradd  -M  smb/gavin

-M参数表示不创建家目录

使其成为samba用户，会提示输入密码：

smbpasswd  -a  smb/gavin

使能该用户：

smbpasswd  -e  smb/gavin



service  smb  restart









