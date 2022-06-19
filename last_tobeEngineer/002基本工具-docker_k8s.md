

# docker&K8S学习

[TOC]



## 1.关系及区别



![dockerEE](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/dockerEE.png)



## 2.docker学习

**作用**：

1.环境变量。用户必须保证两件事：操作系统的设置，各种库和组件的安装。只有它们都正确，软件才能运行。举例来说，安装一个 Python 应用，计算机必须有 Python 引擎，还必须有各种依赖，可能还要配置环境变量。

2.虚拟机。虚拟机（virtual machine）就是带环境安装的一种解决方案。对于底层系统来说，虚拟机就是一个普通文件，不需要了就删掉，对其他部分毫无影响。缺点如下：

**（1）资源占用多**

虚拟机会独占一部分内存和硬盘空间。它运行的时候，其他程序就不能使用这些资源了。哪怕虚拟机里面的应用程序，真正使用的内存只有 1MB，虚拟机依然需要几百 MB 的内存才能运行。

**（2）冗余步骤多**

虚拟机是完整的操作系统，一些系统级别的操作步骤，往往无法跳过，比如用户登录。

**（3）启动慢**

启动操作系统需要多久，启动虚拟机就需要多久。可能要等几分钟，应用程序才能真正运行。

3.linux容器。Linux 容器（Linux Containers，缩写为 LXC）。**Linux 容器不是模拟一个完整的操作系统，而是对进程进行隔离。**或者说，在正常进程的外面套了一个[保护层](https://opensource.com/article/18/1/history-low-level-container-runtimes)。对于容器里面的进程来说，它接触到的各种资源都是虚拟的，从而实现与底层系统的隔离。

由于容器是进程级别的，相比虚拟机有很多优势。

**（1）启动快**

容器里面的应用，直接就是底层系统的一个进程，而不是虚拟机内部的进程。所以，启动容器相当于启动本机的一个进程，而不是启动一个操作系统，速度就快很多。

**（2）资源占用少**

容器只占用需要的资源，不占用那些没有用到的资源；虚拟机由于是完整的操作系统，不可避免要占用所有资源。另外，多个容器可以共享资源，虚拟机都是独享资源。

**（3）体积小**

	容器只要包含用到的组件即可，而虚拟机是整个操作系统的打包，所以容器文件比虚拟机文件要小很多。

总之，容器有点像轻量级的虚拟机，能够提供虚拟化的环境，但是成本开销小得多。

4.docker。**Docker 属于 Linux 容器的一种封装，提供简单易用的容器使用接口。**它是目前最流行的 Linux 容器解决方案。Docker 将应用程序与该程序的依赖，打包在一个文件里面。运行这个文件，就会生成一个虚拟容器。程序在这个虚拟容器里运行，就好像在真实的物理机上运行一样。

	总体来说，Docker 的接口相当简单，用户可以方便地创建和使用容器，把自己的应用放入容器。容器还可以进行版本管理、复制、分享、修改，就像管理普通的代码一样。

5.docker用途

Docker 的主要用途，目前有三大类。

**（1）提供一次性的环境。**比如，本地测试他人的软件、持续集成的时候提供单元测试和构建的环境。

**（2）提供弹性的云服务。**因为 Docker 容器可以随开随关，很适合动态扩容和缩容。

**（3）组建微服务架构。**通过多个容器，一台机器可以跑多个服务，因此在本机就可以模拟出微服务架构。

6.docker的安装

Docker 是一个开源的商业产品，有两个版本：社区版（Community Edition，缩写为 CE）和企业版（Enterprise Edition，缩写为 EE）。企业版包含了一些收费服务，个人开发者一般用不到。下面的介绍都针对社区版。

Docker CE 的安装请参考官方文档。

> - [Mac](https://docs.docker.com/docker-for-mac/install/)
> - [Windows](https://docs.docker.com/docker-for-windows/install/)
> - [Ubuntu](https://docs.docker.com/install/linux/docker-ce/ubuntu/)
> - [Debian](https://docs.docker.com/install/linux/docker-ce/debian/)
> - [CentOS](https://docs.docker.com/install/linux/docker-ce/centos/)
> - [Fedora](https://docs.docker.com/install/linux/docker-ce/fedora/)
> - [其他 Linux 发行版](https://docs.docker.com/install/linux/docker-ce/binaries/)

```bash
 sudo groupadd docker
 sudo usermod -aG docker $USER
```

## 3.CentOS安装docker

```
uname -r #Docker 要求 CentOS 系统的内核版本高于 3.10
yum update
yum remove docker  docker-common docker-selinux docker-engine #卸载旧版本
yum install -y yum-utils device-mapper-persistent-data lvm2 #安装需要的软件包， yum-util 提供yum-config-manager功能，另外两个是devicemapper驱动依赖的
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo 设置yum源
yum list docker-ce --showduplicates | sort -r #可以查看所有仓库中所有docker版本，并选择特定版本安装
yum install --nobest docker-ce -y 
yum install <FQPN>  # 例如：sudo yum install docker-ce-17.12.0.ce
启动并开机自启：
systemctl start docker
systemctl enable docker
docker version
rm -rf /var/lib/docker   #删除以前已有的镜像和容器,非必要
rm -rf /var/run/docker 
docker run hello-word 校验
docker pull ubuntu
docker run -it ubuntu

```





## Configure Docker to start on boot

```
sudo systemctl enable docker
sudo systemctl disable docker

sudo systemctl start docker
```



下载：https://download.docker.com/linux/static/stable/

tar xzvf /path/to/<FILE>.tar.gz

sudo cp docker/* /usr/bin/

sudo dockerd &



```shell
sudo pacman -S docker
sudo systemctl start docker	启用
sudo systemctl enable docker	开机启动
sudo pacman -R docker			删除
systemctl restart docker		重启
rm -rf /var/lib/docker			删除Docker运行过程中产生的镜像、容器等文件。用户生成的配置文件需要手工删除。
yaourt
```







1、docker pull [OPTIONS] NAME [:TAG]  ：此命令的作用是从docker远程的仓库拉取镜像到本地 （命令中的NAME项是必填的代表我们需要拉取的镜像名称；  [:TAG]是可选的，是代表镜像的版本； [OPTIONS]是代表拉取的镜像参数）                      

例如：拉取hello-world镜像命令：docker pull hello-world

2、docker images [OPTIONS][REPOSITORY[:TAG]]  ：此命令是来查看我们本机都有哪些镜像，也可以验证我们的pull是否执行成功（命令中[OPTIONS]是镜像的参数；



### 3-1. 国内加速站点

https://www.cnblogs.com/wushuaishuai/p/9984228.html

<https://registry.docker-cn.com>

<http://hub-mirror.c.163.com>

<https://3laho3y3.mirror.aliyuncs.com>

<http://f1361db2.m.daocloud.io>

<https://mirror.ccs.tencentyun.com>

    https://hub.daocloud.io



### 3-2. 使用命令来配置加速站点

```shell
mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["<http://hub-mirror.c.163.com>"]
}
```

通过修改启动脚本配置加速站点

/usr/lib/systemd/system/docker.service

```shell
vim /usr/lib/systemd/system/docker.service 
# 在dockerd后面加参数
ExecStart=/usr/bin/dockerd --registry-mirror=http://hub-mirror.c.163.com

sudo systemctl daemon-reload
sudo systemctl restart docker
```

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/docker%E7%BB%84%E6%88%90.png)





docker网络类型

bridge,host,none

端口映射

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/docker%E7%BD%91%E7%BB%9C.png)



制作镜像

dockerfile

docker build

http://jpress.io/









## 3.k8s学习











# 引用

## [1.Docker 入门教程](https://www.ruanyifeng.com/blog/2018/02/docker-tutorial.html) 



应用linux一个文件系统，unionFS，一个分层的文件系统，每层都是只读

容器的本质就是一个进程。

hub.docker.com

c.163.com



windows安装

win10外：https://www.docker.com/products/docker-toolbox

win10：https://www.docker.com/products/docker#windows



docker下安装启动jpress

#### 在 Linux 上安装 JPress :

```
wget https://gitee.com/fuhai/jpress/raw/master/docker/docker-compose.yml
docker-compose up -d
```

#### 停止 JPress

```
docker-compose stop
```

#### 启动 JPress

```
docker-compose start
```

#### 重启 JPress

```
docker-compose restart
```

#### 卸载 JPress

```
docker-compose down
```



Compose文件是一个定义服务，网络和卷的YAML文件。 Compose文件的默认文件名为docker-compose.yml





 





#  技巧









cat json | python -mjson.tool

















