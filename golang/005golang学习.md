# VSCODE和GO环境配置

GOPATH=D:\GoPath 工作目录

- src存放源代码的目录，新建项目都在该目录下。
- pkg编译过后生成的包文件存放目录。
- bin编译后生产的可执行文件和go相关的工具，都在此目录。把此目录加入Path环境变量下方便go的使用。

VScode中go插件安装

在`%GOPATH%\src\golang.org\x`目录下打开git bash，执行`git clone http://github.com/golang/tools`。 

在%GOPATH%\src\ 目录下，建立golang.org 文件夹，并再新建x文件夹。  目录为 "%GOPATH\src\golang.org\x\" 

git clone <https://github.com/golang/tools.git> tools 

go install github.com/ramya-rao-a/go-outline 等（https://blog.csdn.net/Yo_oYgo/article/details/79065966）

## FAQ

为什么在别处新建.go文件后，自动提示功能失效了？如在src下新建文件夹test，在test中新建test.go之后，没有自动提示功能了。
**可以把上面下载的工具，通过go get xxx下载的exe工具，拷贝到GOROOT下的bin目录中，再重启vscode试一下。**https://www.jianshu.com/p/b8810bbbd068



# 工具配置

## settings.json

{

​    "files.autoSave": "onFocusChange",

​    "go.buildOnSave": true,

​    "go.lintOnSave": true,

​    "go.vetOnSave": true,

​    "go.buildFlags": [],

​    "go.lintFlags": [],

​    "go.vetFlags": [],

​    "go.useCodeSnippetsOnFunctionSuggest": false,

​    "go.formatOnSave": false,

​    "go.formatTool": "goreturns",

​    "go.goroot": "D:/Program Files/Go",

​    "go.gopath": "${workspaceRoot}"

}其中go.goroot和go.gopath根据自己实际目录设置。 





## launch.json

调试工具配置

go get：获取远程包（需 提前安装 git或hg）

go run：直接运行程序

go build：测试编译，检查是否有编译错误

go fmt：格式化源码（部分IDE在保存时自动调用

go install：编译包文件并编译整个程序

go test：运行测试文件

go doc：查看文档

















# 学习参考资料

## 1.https://tour.golang.org    

是一个系列的Go语言入门教程， 它包含了诸多基本概念和结构相关的并可在线运行的互动小程序。从web页面运行Go语言的程序    https://tour.go-zh.org/basics/1

## 2.https://golang.org    

官方网站，完善的参考文档，包 括编程语言规范和标准库等诸多权威的帮助信息。

## 3.https://blog.golang.org     

Go语言 的官方博客



# 从Java到Golang快速入门

http://www.flysnow.org/2016/12/28/from-java-to-golang.html

http://www.flysnow.org/categories/Golang/



# Go前景

比如docker，k8s等 

[GitHub上优秀的Go开源项目](http://www.flysnow.org/2016/12/27/golang-hot-project-in-github.html) 

微服务

对go语言发展，很利好的消息就是微服务。微服务的发展让我们把一些模块独立成服务，这样子的话，我们就可以为我们特定的服务来选择最适合的语言，这样子的话，Go就在中间件、网络服务、高并发等应用场景中有很大的优势，就会被优先考虑和选择。 

区块链

另外一个更具有竞争力的，促进Go语言发展的就是区块链着这个技术。尤其是以太坊的出现，完全是用go来写的，对go语言的促进非常非常大。另外一个就是区块链联盟，由ibm所发起的超级账本，它主要针对企业级的联盟链儿 

大家有愿意一起共事的，可以跟我联系。 

[Go语言 | 哪些大公司在用go语言？](http://www.flysnow.org/2017/09/13/go-for-company.html)

[go开源项目整理-新手篇](https://blog.csdn.net/haolipengzhanshen/article/details/78205942)

[Google Go 语言从入门到应用所需要的开源项目](https://coderknock.com/blog/2016/12/12/Go.html)

[有哪些值得学习的 Go 语言开源项目？](https://www.zhihu.com/question/20801814)

[go-wiki](https://code.google.com/p/go-wiki/)、[Go Walker](http://gowalker.org/)、[Go Language Resources](http://go-lang.cat-v.org/library-bindings)

[go相关qq群](https://docs.google.com/spreadsheet/lv?key=0AqIvOG5Y0CJ6dFFJV0JwSm1kbEtEdmg5Nk1uZndzakE)







# GO设计意义

为什么会有go？它的意义？

## 1.什么是Go？

Go是一门 **并发支持** 、**垃圾回收** 的 **编译型** **系统编程语言**，旨在创造一门具有在静态编译语言的 **高性能** 和动态语言的 **高效开发** 之间拥有良好平衡点的一门编程语言。



## 2.Go的主要特点有哪些？

——类型安全 和 内存安全以

——非常直观和极低代价的方案实现 高并发

——高效的垃圾回收机制

——快速编译（同时解决C语言中头文件太多的问题）

——为多核计算机提供性能提升的方案

——UTF-8编码支持

## 3.[Go在谷歌：以软件工程为目的的语言设计](https://www.oschina.net/translate/go-at-google-language-design-in-the-service-of-software-engineering)

Go语言设计时考虑的因素，除了大家较为了解的内置并发和内存垃圾自动回收这些方面之外，还包括严格的依赖管理、对随系统增大而在体系结构方面发生变化的适应性、跨组件边界的健壮性（robustness）。 

Go语言运行效率高，具有较强的可伸缩性(scalable)，而且使用它进行工作时的效率也很高。 虽然因此而设计出的语言不会是一门在研究领域里具有突破性进展的语言，但它却是大型软件项目中软件工程方面的一个非常棒的工具。 

Go语言考虑更多的是软件工程而不是编程语言方面的科研。或者，换句话说，它是为软件工程服务而进行的语言设计。 

这些问题包括：

- Build速度缓慢
- 失控的依赖关系
- 每个程序员使用同一门语言的不同子集
- 程序难以理解（代码难以阅读，文档不全面等待）
- 很多重复性的劳动
- 更新的代价大
- 版本偏斜（version skew）
- 难以编写自动化工具
- 语言交叉Build（cross-language build）产生的问题

一门语言每个单个的特性都解决不了这些问题。这需要从软件工程的大局观，而在Go语言的设计中我们试图致力于解决*所有这些*问题。

在单个的计算机上build出Google的服务器二进制程序就变得不太实际了，因此我们创建了一个[**大型分布式编译系统**](http://google-engtools.blogspot.com/2011/06/build-in-cloud-accessing-source-code.html)。 该系统非常复杂（这个Build系统本身也是个大型程序）还使用了大量机器以及大量缓存，藉此在Google进行Build才算行得通了，尽管还是有些困难。 

尽管现在的讨论更专注于依赖关系，这里依然还有很多其他需要关注的问题。这一门成功语言的主要因素是：

- 它必须适应于大规模开发，如拥有大量依赖的大型程序，且又一个很大的程序员团队为之工作。

- 它必须是熟悉的，大致为 C 风格的。谷歌的程序员在职业生涯的早期，对函数式语言，特别是 C 家族更加熟稔。要想程序员用一门新语言快速开发，新语言的语法不能过于激进。

- 它必须是现代的。C、C++以及Java的某些方面，已经过于老旧，设计于多核计算机、网络和网络应用出现之前。新方法能够满足现代世界的特性，例如内置的并发。

Go的依赖图还有另外一个特性，就是它不包含循环。 

尽管如此，Go语言为了提高程序的健壮性，还是对C语言的语意做出了很多小改动。它们包括：

- 不能对指针进行算术运算
- 没有隐式的数值转换
- 数组的边界总是会被检查
- 没有类型别名（进行type X int的声明后，X和int是两种不同的类型而不是别名）
- ++和--是语句而不是表达式
- 赋值不是一种表达式
- 获取栈变量的地址是合法的（甚至是被鼓励的）
- 其他

还有一些很大的改变，同传统的C 、C++ 、甚至是JAVA 的模型十分不同。它包含了对以下功能的支持：

- 并发
- 垃圾回收
- 接口类型
- 反射
- 类型转换

下面的章节从软件工程的角度对 Go 语言这几个主题中的两个的讨论：**并发和垃圾回收**。 

### **并发**

运行于多核机器之上并拥有众多客户端的web服务器程序，可称为Google里最典型程序。在这样的现代计算环境中，并发很重要。这种软件用C++或Java做都不是特别好，因为它们缺在与语言级对并发支持的都不够好。 Go采用了一流的**channel，体现为CSP的一个变种**。之所以选择CSP，部分原因是因为大家对它的熟悉程度（我们中有一位同事曾使用过构建于CSP中的概念之上的前任语言），另外还因为CSP具有一种在无须对其模型做任何深入的改变就能轻易添加到过程性编程模型中的特性。也即，对于类C语言，CSP可以一种最长正交化（orthogonal）的方式添加到这种语言中，为该语言提供额外的表达能力而且还不会对该语言的其它用它施加任何约束。简言之，就是该语言的其它部分仍可保持“通常的样子”。 

假设Web服务器必须验证它的每个客户端的安全证书；在Go语言中可以很容易的使用CSP来构建这样的软件，将客户端以独立执行的过程来管理，而且还具有编译型语言的执行效率，足够应付昂贵的加密计算。 总的来说，CSP对于Go和Google来说非常实用。在编写Web服务器这种Go语言的典型程序时，这个模型简直是天作之合。 

有一条警告很重要：因为有并发，所以Go不能成为纯的内存安全（memory safe）的语言。共享内存是允许的，通过channel来传递指针也是一种习惯用法（而且效率很高）。 

有些并发和函数式编程专家很失望，因为Go没有在并发计算的上下文中采用只写一次的方式作为值语义，比如这一点上Go和Erlang就太象。 Go让使得简单而安全的并发编程*成为可能*，但它并不*阻止*糟糕的编程方式。这个问题我们通过惯例来折中，训练程序员将消息传递看做拥有权限控制的一个版本。有句格言道：“**不要通过共享内存来通信，要通过通信来共享内存。**” 

### **垃圾回收**

Go语言没有显式的内存释放操作，那些被分配的内存只能通过垃圾回收器这一唯一途径来返回内存池。 此外，拥有自动化的内存管理机制对于一门并发的面向对象的编程语言来说很关键，因为一个内存块可能会在不同的并发执行单元间被来回传递，要管理这样一块内存的所有权对于程序员来说将会是一项挑战。将行为与资源的管理分离是很重要的。 

当然，垃圾回收机制会带来很大的成本：资源的消耗、回收的延迟以及复杂的实现等。 但给编程者带来了更好的体验。

有个关键点在于，Go为程序员提供了通过控制数据结构的格式来限制内存分配的手段 

```go
type X struct {
    a, b, c int
    buf [256]byte
}
```

在Java中，buffer字段需要再次进行内存分配，因为需要另一层的间接访问形式。然而在Go中，该缓冲区同包含它的struct一起分配到了一块单独的内存块中，无需间接形式。对于系统编程，这种设计可以得到更好的性能并减少回收器（collector）需要了解的项目数。要是在大规模的程序中，这么做导致的差别会非常巨大。

有个更加直接一点的例子，**在Go中，可以非常容易和高效地提供二阶内存分配器（second-order allocator）**，例如，为一个由大量struct组成的大型数组分配内存，并用一个自由列表（a free list）将它们链接起来的arena分配器（an arena allocator）。在重复使用大量小型数据结构的库中，可以通过少量的提前安排，就能不产生任何垃圾还能兼顾高效和高响应度。

虽然Go是一种支持内存垃圾回收的编程语言，但是资深程序员能够限制施加给回收器的压力从而提高程序的运行效率（Go的安装包中还提供了一些非常好的工具，用这些工具可以**研究程序运行过程中动态内存的性能**。） 

要给程序员这样的灵活性，Go必需支持指向分配在堆中对象的指针，我们将这种指针称为***内部指针***。上文的例子中X.buff字段保存于struct之中，但也可以保留这个内部字段的地址。比如，可以将这个地址传递给I/O子程序。在Java以及许多类似的支持垃圾回收的语音中，不可能构造象这样的内部指针，但在Go中这么做很自然。这样设计的指针会影响可以使用的回收算法，并可能会让算法变得更难写，但经过慎重考虑，我们决定允许内部指针是必要的，因为这对程序员有好处，让大家具有降低对（可能实现起来更困难）回收器的压力的能力。到现在为止，我们的将大致相同的Go和Java程序进行对比的经验表明，**使用内部指针能够大大影响arena总计大型、延迟和回收次数**。 

总的说来，Go是一门支持垃圾回收的语言，但它同时也提供给程序员一些手段，可以对回收开销进行控制。 

垃圾回收器目前仍在积极地开发中。当前的设计方案是**并行的边标示边扫描（mark-and-sweep）的回收器**，未来还有机会提高其性能甚至其设计方案。（**Go语言规范中并没有限定必需使用哪种特定的回收器实现方案**）。 

**要组合不要继承（面向对象编程）**





### **总结**

造成这种效果的因素有：

- 清晰的依赖关系
- 清晰的语法
- 清晰的语义
- 偏向组合而不是继承
- 编程模型（垃圾回收、并发）所代理的简单性
- 易于为它编写工具（Easy tooling ）(gotool、gofmt、godoc、gofix)



# sublime GO配置



## Go常用命令简介

go get：获取远程包（需 提前安装 git或hg）

go run：直接运行程序

go build：测试编译，检查是否有编译错误

go fmt：格式化源码（部分IDE在保存时自动调用）

go install：编译包文件并编译整个程序

go test：运行测试文件

go doc：查看文档（CHM手册）

## Go内置关键字（25个均为小写）

break        default           func        interface        select

case          defer              go           map               struct

chan          else                goto       package        switch

const         fallthrough    if             range             type

continue   for                  import    return             var 

- var和const ：变量和常量的声明
- var varName type  或者 varName : = value
- package and import: 导入，package可用别名{ to "fmt"}
- func： 用于定义函数和方法
- return ：用于从函数返回
- defer someCode ：在函数退出之前执行
- go : 用于并行
- select 用于选择不同类型的通讯
- interface 用于定义接口
- struct 用于定义抽象数据类型
- break、case、continue、for、fallthrough、else、if、switch、goto、default 流程控制
- chan用于channel通讯
- type用于声明自定义类型，进行结构(struct)或接口(interface)的声明
- map用于声明map类型数据
- range用于读取slice、map、channel数据



使用 大小写 来决定该 常量、变量、类型、接口、结构或函数 是否可以被外部包所调用

## 数据类型

### 整型类型

| 类型名称 | 有无符号 | bit数         |
| -------- | -------- | ------------- |
| int8     | Yes      | 8             |
| int16    | Yes      | 16            |
| int32    | Yes      | 32            |
| int64    | Yes      | 64            |
| uint8    | No       | 8             |
| uint16   | No       | 16            |
| uint32   | No       | 32            |
| uint64   | No       | 64            |
| int      | Yes      | 等于cpu位数   |
| uint     | No       | 等于cpu位数   |
| rune     | Yes      | 与 int32 等价 |
| byte     | No       | 与 uint8 等价 |
| uintptr  | No       | -             |

 **rune** 类型是 Unicode 字符类型，和 int32 类型等价，通常用于表示一个 Unicode 码点。rune 和 int32 可以互换使用。

**byte** 是uint8类型的等价类型，byte类型一般用于强调数值是一个原始的数据而不是 一个小的整数。

**uintptr** 是一种无符号的整数类型，没有指定具体的bit大小但是足以容纳指针。 uintptr类型只有在底层编程是才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方。 

不管它们的具体大小，int、uint和uintptr是不同类型的兄弟类型。其中int和int32也是 不同的类型， 即使int的大小也是32bit，在需要将int当作int32类型的地方需要一个显式 的类型转换操作，反之亦然。 

### 浮点数

Go语言提供了两种精度的浮点数，float32和float64。它们的算术规范由IEEE754浮点数国际标准定义， 该浮点数规范被所有现代的CPU支持。

这些浮点数类型的取值范围可以从很微小到很巨大。浮点数的范围极限值可以在math包找到。常量 math.MaxFloat32表示float32能表示的最大数值，大约是 3.4e38；对应的math.MaxFloat64常量大约是 1.8e308。它们分别能表示的最小值近似为1.4e-45和4.9e-324。

一个float32类型的浮点数可以提供大约6个十进制数的精度，而float64则可以提供约15个十进制数的精度。

函数math.IsNaN用于测试一个数是否是非数NaN，math.NaN则返回非数对应的值。虽然可以用math.NaN来 表示一个非法的结果，但是测试一个结果是否是非数NaN则是充满风险的，因为NaN和任何数都是不相等的。 如果一个函数返回的浮点数结果可能失败，最好的做法是用单独的标志报告失败。

### 复数

Go语言提供了两种精度的复数类型：complex64和complex128，分别对应float32和float64两种浮点数精度。内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部。

```go
z := x + yi
x = real(z)
y = imag(z)
```

布尔型和字符串































































