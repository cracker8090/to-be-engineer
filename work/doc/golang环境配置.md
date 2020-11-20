





# github网络配置

```text
获取Github相关网站的ip
访问https://www.ipaddress.com，拉下来，找到页面中下方的“IP Address Tools – Quick Links”
分别输入github.global.ssl.fastly.net和github.com，查询ip地址
下面是我的配置
140.82.114.4	github.com
199.232.5.194	github.global.ssl.fastly.net
```



```
$ git config --global user.name "John Doe"
$ git config --global user.email johndoe@example.com
```



ssh-keygen -t rsa -C ["xxxxx@xxxxx.com](mailto:"xxxxx@xxxxx.com)"





# git设置代理



git config --global http.proxy 'socks5://127.0.0.1:1080'
git config --global https.proxy 'socks5://127.0.0.1:1080'

```
git config --global http.proxy 'socks5://127.0.0.1:1080'

git config --global https.proxy 'socks5://127.0.0.1:1080'
```



```
全部采用
# socks5协议，1080端口修改成自己的本地代理端口
git config --global http.proxy socks5://127.0.0.1:1080
git config --global https.proxy socks5://127.0.0.1:1080

# http协议，1081端口修改成自己的本地代理端口
git config --global http.proxy http://127.0.0.1:1081
git config --global https.proxy https://127.0.0.1:1081
```



仅仅针对github进行配置，让github走本地代理，其他的保持不变；

```
# socks5协议，1080端口修改成自己的本地代理端口
git config --global http.https://github.com.proxy socks5://127.0.0.1:1080
git config --global https.https://github.com.proxy socks5://127.0.0.1:1080

# http协议，1081端口修改成自己的本地代理端口
git config --global http.https://github.com.proxy https://127.0.0.1:1081
git config --global https.https://github.com.proxy https://127.0.0.1:1081
```

```
# 查看所有配置
git config -l
# reset 代理设置
git config --global --unset http.proxy
git config --global --unset https.proxy
```



Windows上

git config --global core.autocrlf true

git config --global --add core.filemode false

git config --global core.filemode false



git config --global core.filemode false
git config core.filemode false



```
git rm -r --cached .
git add .
git commit -m 'update .gitignore
```

git config ---global core.editor vim



下面简单的方法可以让git diff的时候忽略换行符的差异：

git config --global core.whitespace cr-at-eol

更好的方法是每个项目都有一个.gitattributes文件，里面配好了换行符的设置，参考

[https://help.github.com/articles/dealing-with-line-endings](https://link.jianshu.com?t=https%3A%2F%2Fhelp.github.com%2Farticles%2Fdealing-with-line-endings)



IDM下载地址

https://www.macxin.com/archives/9641.html



# golang代理



```
export GO111MODULE=on
export GOPROXY=https://goproxy.io
```

也可以通过置空这个环境变量来关闭，export GOPROXY=

对于 Windows 用户，可以在 PowerShell 中设置：

```
$env:GOPROXY = "https://goproxy.io"
```

```
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPATH=D:\gavin\go_code
go env -w GO111MODULE=on
```



gopath设置成你工作的目录



go intsall tools

安装完的工具会保留在我们的GOPATH/bin/下面



# vscode配置

插件

http://www.meicx.com/?p=22223





## 1. 一次搜索所有文件的文本

Windows: Ctrl + Shift + F Mac: Command + Shift + F



## 3.进程资源管理器

你是否发现你的VsCode 编辑器有时有点慢

Windows在VsCode 中按`Ctrl + Alt + Delete`



## 集成终端

通过 **Ctrl + `**可以打开或关闭终端



## 查看正在运行插件

命令面板(`Ctrl + Shift + P`)并输入`Show running extensions`来查看所有你安装的正在运行的插件。



## 重新打开 关闭的编辑页面

Windows: Ctrl + Shift + T Mac: command + Shift + T

当你处理一个文件很多的大型项目时，如果不小心关闭了一个页面，并且不得不在侧菜单中再次搜索它，这可能会有点令人沮丧。

现在，可以按 `Ctrl + Shift + T` 重新打开一个关闭的页面



## 通过匹配文本打开文件

Windows: Ctrl + T Mac: command + T

说到搜索文件，你可以动态地搜索和打开文件。这是我最喜欢的特性之一，因为不需要手动单击目录来重新打开一个不再打开的文



## 重新加载

个人认为这是 VsCode 最酷的特性之一。它允许你在重新加载编辑器时将窗口放在前面，同时具有与关闭和重新打开窗口相同的效果。

Windows: Ctrl + Alt + R Mac: Control + Option + R

## 删除上一个单词

要删除前一个单词，可以按`Ctrl + Backspace` (Mac: `option + delete`)。这在你打错字的时候非常有用。

## 启动性能

打开命令面板(`Ctrl + Shift + P`)，搜索S`tartup Performance`。

## 逐个选择文本

可以通过快捷键`Ctrl + Shift +右箭头`(Mac: `option + Shift +右箭头`)和`Ctrl + Shift +左箭头`(Mac: option + Shift +左箭头)逐个选择文本。

## 重复的行

一个非常强大和已知的功能是复制行。只需按 `Shift + Alt + 向下箭头` (Mac: `command + Shift + 向下箭头`

## 批量替换当前文件中所有匹配的文本

可以选择任何一组文本，如果该选中文本出现多个，可以通过按`Ctrl + F2` (Mac: `command + F2`)一次改所有出现的文本。

## 向上/向下移动一行

按`Alt + 向上箭头`(Mac: `option+ 向上箭头`)当前行向上移动，按`Alt + 向下箭头`(Mac: `option+ 向下箭头`))当前行向下移动。

## 删除一行

使用`Ctrl + X`剪切命令(`Mac：command + X`)来删除一行。

或者使用 `Ctrl + Shift + K` (Mac: `command + Shift + K`)命令。

## 将编辑器向左或向右移动

你可能会有一种无法控制的欲望，想要在一个组中重新排列选项卡，其中选项卡相互关联，左边的选项卡是比较重要文件，而右边的选项卡是相对不重要的文件。 通过 Ctrl+Shift+PgUp/PgDown(command + +Shift+PgUp/PgDown)向左/向右移动编辑器。

##  复制光标向上或者向上批量添加内容

![](https://user-gold-cdn.xitu.io/2019/7/22/16c170549b1d2320?imageslim) 



# golang相关

## golang中json的tag omitempty

**如果没有omitempty，该字段是会显示的。** 



## gomod

开启模块支持后，并不能与GOPATH共存,所以把项目从GOPATH中移出即可

```bash
$GOPATH/go.mod exists but should not
```

问题描述：$GOPATH/go.mod exists but should not

产生原因：开启模块支持后，并不能与$GOPATH共存,所以把项目从$GOPATH中移出即可

| 命令     | 说明                                                         |
| -------- | ------------------------------------------------------------ |
| download | download modules to local cache(下载依赖包 重要)             |
| edit     | edit go.mod from tools or scripts（编辑go.mod                |
| graph    | print module requirement graph (打印模块依赖图)              |
| init     | initialize new module in current directory（在当前目录初始化mod 重要） |
| tidy     | add missing and remove unused modules(拉取缺少的模块，移除不用的模块 重要) |
| vendor   | make vendored copy of dependencies(将依赖复制到vendor下)     |
| verify   | verify dependencies have expected content (验证依赖是否正确） |
| why      | explain why packages or modules are needed(解释为什么需要依赖) |





