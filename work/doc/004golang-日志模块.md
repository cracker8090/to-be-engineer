# ssh工具

TabNine是他们用过的最好的代码补全工具，这是属于程序员的**杀手级应用**。 

 1、cmder   

​    网站：https://cmder.net/

​    说到这个软件，大概就是我找xshell代替的，在百度上最多的回答之一了，看过之后也确实很有感觉，直接输入ssh命令就可以直接访问远程主机，还能自定义保存我们的主机配置，方便连接多个主机使用，官方的文档和教程也科普了很多的快捷键和配置说明，使用非常简单，还能随时跳转Windows自带的cmd和powershell，非常值得使用。

  2、electerm 

​    网站：https://electerm.html5beta.com/

​    这个软件，是笔者在捣鼓深度的deepinOS偶然间发现的，在深度操作系统上使用感觉还颇为不错，后来百度了一下，发现也是开源的并且制霸全平台的一个基于UI的软件，于是乎笔者在mac，Windows上都下载了一下，发现确实也都有很不错的体验，UI里面自带的标签功能也真的是很实用，只不过这个软件在macos上的表现还有待改进，因为他自带的UI 应该是以windows为标准写的，所以对于窗口的操作是在右上角，和mac系统统一的使用习惯（左上角）有些不同，所以可能会让有些人用的不爽。





Bitvise SSH Client

finalshell

mRemoteNG 

wsl

Termius

MobaXterm

[Best Putty Alternatives for SSH Clients for Windows (FREE!)](https://www.pcwdld.com/best-alternatives-to-putty) 



gpg验证签名完整性



MyDB Studio

制造商：**H2LSoft, Inc.**
网站：[http://www.mydb-studio.com/](http://www.mydb-studio.com/?spm=a2c6h.12873639.0.0.49ff175fM5R6QR)
价格：免费
授权：免费软件
支持平台：Windows 2000，XP，Vista，7





# 日志输出



command  | tee stdout.log

**这里有一个需要注意的坑点，上面的命令只是把标准输出，也就是 STDOUT 写进文件，并没有处理错误输出，也就是 STDERR** 

command > >(tee -a stdout.log) 2> >(tee -a stderr.log >&2)

command  2>&1 | tee -a log

Golang自带的日志并没有提供分级功能（Trace、Debug、Info、Warn、Error、Fatal）





## 输出日志到文件和控制台

```go
func main() {
    fmt.Println("---------------")
    log.Println("------ log printl ----")
    func_log2file()
    func_log2fileAndStdout()
}
func func_log2file() {
    //创建日志文件
    f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatal(err)
    }
    //完成后，延迟关闭
    defer f.Close()
    // 设置日志输出到文件
    log.SetOutput(f)
    // 写入日志内容
    log.Println("check to make sure it works")
}
func func_log2fileAndStdout() {
    //创建日志文件
    f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatal(err)
    }
    //完成后，延迟关闭
    defer f.Close()
    // 设置日志输出到文件
    // 定义多个写入器
    writers := []io.Writer{
        f,
        os.Stdout}
    fileAndStdoutWriter := io.MultiWriter(writers...)
    // 创建新的log对象
    logger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
    // 使用新的log对象，写入日志内容
    logger.Println("--> logger :  check to make sure it works")
}
```







日志可以用于查看和分析应用程序的运行状态。日志一般可以支持输出多种形式，如命令行、文件、网络等。

本例将搭建一个支持多种写入器的日志系统，可以自由扩展多种日志写入设备













https://github.com/sirupsen/logrus

https://github.com/sirupsen/logrus?utm_source=gold_browser_extension





















