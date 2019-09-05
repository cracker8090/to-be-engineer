

## [介绍](https://zh.wikipedia.org/wiki/Clojure) 

**Clojure**（[/ˈkloʊʒər/](https://zh.wikipedia.org/wiki/Help:%E8%8B%B1%E8%AA%9E%E5%9C%8B%E9%9A%9B%E9%9F%B3%E6%A8%99)）[[15\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-15)是[Lisp](https://zh.wikipedia.org/wiki/Lisp)[编程语言](https://zh.wikipedia.org/wiki/%E7%BC%96%E7%A8%8B%E8%AF%AD%E8%A8%80)在[Java](https://zh.wikipedia.org/wiki/Java%E5%B9%B3%E8%87%BA)平台上的现代、[动态](https://zh.wikipedia.org/wiki/%E5%8A%A8%E6%80%81%E8%AF%AD%E8%A8%80)及[函数式](https://zh.wikipedia.org/wiki/%E5%87%BD%E6%95%B8%E7%A8%8B%E5%BC%8F%E8%AA%9E%E8%A8%80)方言。[[16\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-platform/android-16)[[17\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-clojure.org-17) 与其他Lisp一样，Clojure视[代码为数据](https://zh.wikipedia.org/wiki/%E5%90%8C%E5%83%8F%E6%80%A7)且拥有一套Lisp[宏](https://zh.wikipedia.org/wiki/%E5%B7%A8%E9%9B%86)系统。[[18\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-reference/lisps-18) 其开发过程当前由[社区](https://zh.wikipedia.org/wiki/%E5%AF%A6%E8%B8%90%E7%A4%BE%E7%BE%A4)驱动，[[19\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-dev/dev-19) 其作者里奇·希基（Rich Hickey）则以[终身仁慈独裁者](https://zh.wikipedia.org/wiki/%E7%BB%88%E8%BA%AB%E4%BB%81%E6%85%88%E7%8B%AC%E8%A3%81%E8%80%85)（BDFL）的身份进行监督。

里奇·希基开发Clojure的原因是因为他想要一款适合[函数式编程](https://zh.wikipedia.org/wiki/%E5%87%BD%E6%95%B8%E7%A8%8B%E5%BC%8F%E8%AA%9E%E8%A8%80)的现代[Lisp](https://zh.wikipedia.org/wiki/Lisp)。该语言既需要与已创建的[Java平台](https://zh.wikipedia.org/wiki/Java%E5%B9%B3%E8%87%BA)共生又需要有适合[并发性](https://zh.wikipedia.org/wiki/%E5%B9%B6%E5%8F%91%E6%80%A7)的设计。[[22\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-about/rationale-22)[[23\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-msdn/inside-clojure-23)[[35\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-35)

Clojure对待变化（change）的方式以[标识](https://zh.wikipedia.org/wiki/%E6%A0%87%E8%AF%86)（identity）的概念为特征。[[21\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-about/state-21) 标识是指随着时间的推移而产生的一系列状态（state）。而状态则是指标识在某一特定时间点上的值（value）。需要强调的是，这里的值是不可变的（immutable）。由此引申，由于状态是不可变的值，任意数量的工作单位（workers）都可以在其上以并行（parallel）的方式实施操作。因此，并发性（concurrency）就成为一道管理状态间变化的问题［注意，这里的“变化”是指从一个状态到另外一个状态的跃迁（transition）而不是状态本身的变化（mutation）。］为此，Clojure提供了几个可变的（mutable）引用类型（reference type）。每个引用类型都有其明确定义的语义用于控制状态之间的跃迁。



| 版本       | 发布日期                                                     | 主要功能/改进                                                |
| ---------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
|            | 2007年10月16日[[36\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-36) | 首次公开发布                                                 |
| 1.0        | 2009年5月4日[[37\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-37) | 首个稳定版                                                   |
| 1.1        | 2009年12月31日[[38\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-38) | 将来（future）                                               |
| 1.2        | 2010年8月19日[[39\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-39) | 协议（protocol）                                             |
| 1.3        | 2011年9月23日[[40\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-40) | 增强对原始类型（primitive type）的支持                       |
| 1.4        | 2012年4月15日[[41\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-41) | 读取器字面量（reader literal）                               |
| 1.5        | 2013年3月1日[[42\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-42) | 归纳器（reducer）                                            |
| 1.5.1      | 2013年3月10日[[43\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-43) | 修复内存泄漏                                                 |
| 1.6        | 2014年3月25[[44\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-44) | Java API、经过改进的哈希算法                                 |
| 1.7        | 2015年6月30日[[45\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-release/clojure-17-45) | 变换归纳器（transducer）、读取器条件表达式（reader conditionals） |
| 1.8        | 2016年1月19日[[46\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-46) | 附加的字符串函数、直接连接（direct linking）、套接字服务器（socket server） |
| 1.9        | 2017年12月8日[[47\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-47) | 集成spec、命令行工具                                         |
| 1.10       | 2018年12月17日[[48\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-48) | 经过改进的错误报告、Java兼容性                               |
| **1.10.1** | 2019年6月6日[[8\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-release/clojure1-10-1-8) | 解决Java性能回归问题并改进clojure.main的错误报告             |

Clojure执行于[Java平台](https://zh.wikipedia.org/wiki/Java%E5%B9%B3%E8%87%BA)之上，因此，与Java紧密集成并完全支持从Clojure调用[Java](https://zh.wikipedia.org/wiki/Java)代码。[[49\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-about/jvm_hosted-49) 与此同时，也可以从Java调用Clojure代码。[[50\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-50) Leiningen是社区中普遍使用的项目自动化工具。Leiningen为[Maven](https://zh.wikipedia.org/wiki/Apache_Maven)集成提供支持，处理项目包管理和依赖项。Leiningen的配置使用的则是Clojure语法。[[51\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-51)

与其他大多数[Lisp](https://zh.wikipedia.org/wiki/Lisp)一样，Clojure的语法创建在[S-表达式](https://zh.wikipedia.org/wiki/S-%E8%A1%A8%E8%BE%BE%E5%BC%8F)之上。S-表达式在被编译之前先由读取器（reader）解析为数据结构。[[52\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-reference/reader-52) 除了列表（list）之外，Clojure的读取器还支持[映射](https://zh.wikipedia.org/wiki/%E5%93%88%E5%B8%8C%E8%A1%A8)（map）、集合（set）及[向量](https://zh.wikipedia.org/wiki/%E6%95%B0%E7%BB%84)（vector）等的字面量（literal）语法。这些字面量随后会被直接编译成上述数据结构。[[52\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-reference/reader-52) Clojure是[Lisp-1](https://zh.wikipedia.org/wiki/Common_Lisp#%E5%87%BD%E6%95%B0%E5%90%8D%E5%AD%97%E7%A9%BA%E9%97%B4)且有一套与其它Lisp不兼容的数据结构，因此，Clojure不支持与Lisp的其它方言之间的代码级兼容性。[[18\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-reference/lisps-18)

作为一门Lisp方言，[函数](https://zh.wikipedia.org/wiki/%E5%AD%90%E7%A8%8B%E5%BA%8F#%E5%87%BD%E6%95%B8)在Clojure中是[一等公民](https://zh.wikipedia.org/wiki/%E7%AC%AC%E4%B8%80%E9%A1%9E%E7%89%A9%E4%BB%B6)。此外，Clojure还支持[读取﹣求值﹣输出循环](https://zh.wikipedia.org/wiki/%E8%AF%BB%E5%8F%96%EF%B9%A3%E6%B1%82%E5%80%BC%EF%B9%A3%E8%BE%93%E5%87%BA%E5%BE%AA%E7%8E%AF)（REPL）以及一套宏系统。[[6\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-reference/macros-6) Clojure的[Lisp宏](https://zh.wikipedia.org/wiki/%E5%B7%A8%E9%9B%86)系统与[Common Lisp](https://zh.wikipedia.org/wiki/Common_Lisp)的系统极为相似。唯一不同的是，Clojure的[重音符](https://zh.wikipedia.org/wiki/%E9%87%8D%E9%9F%B3%E7%AC%A6#%E7%B7%A8%E7%A8%8B%E7%94%A8%E9%80%94)［称为“语法引用”（syntax quote）］用[名字空间](https://zh.wikipedia.org/wiki/%E5%91%BD%E5%90%8D%E7%A9%BA%E9%97%B4)（namespace）来限定符号（symbol）。这有助于防止意外的名字捕获（unintended name capture），因为Clojure禁止绑定（binding）到用名字空间限定的名字（namespace-qualified name）上。如果需要强制捕获宏扩展（capturing macro expansion，）那么就需要显示地完成该过程。Clojure不支持用户定义的读取器宏（reader macro，）但Clojure的读取器支持更具约束力的语法扩展（syntactic expansion）形式。[[53\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-53) Clojure支持[多重方法](https://zh.wikipedia.org/wiki/%E5%A4%9A%E5%88%86%E6%B4%BE)（multimethods。）[[54\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-54) 对于类似[接口](https://zh.wikipedia.org/wiki/%E4%BB%8B%E9%9D%A2_(%E8%B3%87%E8%A8%8A%E7%A7%91%E6%8A%80))的抽象，Clojure提供基于协议（protocol）[[55\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-55)的多态性（polymorphism）以及基于[记录](https://zh.wikipedia.org/wiki/%E8%AE%B0%E5%BD%95)（record）[[56\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-56)的数据类型系统。 Clojure通过这些设计来提供高性能且动态的多态性以避免所谓的“表达式问题”（"expression problem"。）

Clojure支持[惰性序列](https://zh.wikipedia.org/wiki/%E6%83%B0%E6%80%A7%E6%B1%82%E5%80%BC)，并鼓励[不可变性](https://zh.wikipedia.org/wiki/%E4%B8%8D%E5%8F%AF%E8%AE%8A%E7%89%A9%E4%BB%B6)（immutability）与持久数据结构（persistent data structure。）Clojure作为一门[函数式编程语言](https://zh.wikipedia.org/wiki/%E5%87%BD%E6%95%B8%E7%A8%8B%E5%BC%8F%E8%AA%9E%E8%A8%80)将重点放在[递归](https://zh.wikipedia.org/wiki/%E9%80%92%E5%BD%92)与[高阶函数](https://zh.wikipedia.org/wiki/%E9%AB%98%E9%98%B6%E5%87%BD%E6%95%B0)上而不是基于[副作用](https://zh.wikipedia.org/wiki/%E5%87%BD%E6%95%B0%E5%89%AF%E4%BD%9C%E7%94%A8)的[循环](https://zh.wikipedia.org/wiki/%E8%BF%B4%E5%9C%88)流程上。Clojure不支持自动[尾调用](https://zh.wikipedia.org/wiki/%E5%B0%BE%E8%B0%83%E7%94%A8)优化，因为JVM还不支持该项优化，[[57\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-57)[[58\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-58)[[59\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-59)但是，可以用`recur`关键字显式地进行该项优化。[[60\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-60) 对于[并行](https://zh.wikipedia.org/wiki/%E5%B9%B6%E8%A1%8C%E8%AE%A1%E7%AE%97)与[并发](https://zh.wikipedia.org/wiki/%E5%B9%B6%E5%8F%91%E8%AE%A1%E7%AE%97)计算，Clojure提供[软件事务内存](https://zh.wikipedia.org/wiki/%E8%BD%AF%E4%BB%B6%E4%BA%8B%E5%8A%A1%E5%86%85%E5%AD%98)、[[61\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-61) [响应式](https://zh.wikipedia.org/wiki/%E5%93%8D%E5%BA%94%E5%BC%8F%E7%BC%96%E7%A8%8B)[代理系统](https://zh.wikipedia.org/wiki/%E4%B8%AA%E4%BD%93%E4%B8%BA%E6%9C%AC%E6%A8%A1%E5%9E%8B)[[1\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-reference/agents-1)及基于[通道](https://zh.wikipedia.org/wiki/%E4%BA%A4%E8%AB%87%E5%BE%AA%E5%BA%8F%E7%A8%8B%E5%BC%8F)（channel）的并发编程。[[62\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-62)

Clojure 1.7引入了读取器条件表达式（reader conditional）从而允许在同一名字空间中嵌入Clojure与ClojureScript代码。[[45\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-release/clojure-17-45)[[52\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-reference/reader-52) 变换归纳器（transducer）的加入则提供了另一种组合变换的方法。变换归纳器可以使高阶函数（如，`map`和`fold`）更加抽象从而使之独立于其输入数据源。传统地说，这些函数一般被应用于[序列](https://zh.wikipedia.org/wiki/%E4%B8%B2%E5%88%97_(%E6%8A%BD%E8%B1%A1%E8%B3%87%E6%96%99%E5%9E%8B%E5%88%A5))（sequence）上，而变换归纳器允许这些函数被应用于通道（channel）上并让用户定义她们自己的变换归纳（transduction）模型。

#### 平台

Clojure的主要平台是[Java](https://zh.wikipedia.org/wiki/Java%E5%B9%B3%E8%87%BA)[[17\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-clojure.org-17)[[49\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-about/jvm_hosted-49)但也存在其他目标平台上的实现。其中，最值得关注的是ClojureScript[[66\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-66)（可被编译成[ECMAScript](https://zh.wikipedia.org/wiki/ECMAScript) 3[[67\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-67)）和ClojureCLR[[68\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-68)（[.NET](https://zh.wikipedia.org/wiki/%E9%80%9A%E7%94%A8%E8%AF%AD%E8%A8%80%E6%9E%B6%E6%9E%84)平台上的完整移植版，可与其生态系统互操作。）2013年对1,060名受访者进行的Clojure社区调查[[69\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-69)发现，47％的受访者在使用Clojure的同时也使用ClojureScript。2014年，这一数字增长到了55％，[[70\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-70)而到了2015年，则达到了66％（根据2,445名受访者）。[[71\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-71) 人气较高的ClojureScript项目包括[React](https://zh.wikipedia.org/wiki/React)实现，如Reagent[[72\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-72)、re-frame[[73\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-73)、Rum[[74\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-74)及Om



| 集成开发环境/编辑器                                          | Clojure插件                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [Atom](https://zh.wikipedia.org/wiki/Atom_(%E6%96%87%E5%AD%97%E7%B7%A8%E8%BC%AF%E5%99%A8)) | Chlorine[[105\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-105) |
| [Emacs](https://zh.wikipedia.org/wiki/Emacs)                 | CIDER[[106\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-106) |
| [IntelliJ IDEA](https://zh.wikipedia.org/wiki/IntelliJ_IDEA) | Clojure-Kit[[107\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-107)或Cursive[[108\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-108)（提供免费的非商业许可证） |
| Light Table[[109\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-109) | （不适用）                                                   |
| [Vim](https://zh.wikipedia.org/wiki/Vim)                     | fireplace.vim[[110\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-110)[[111\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-111) |
| [Visual Studio Code](https://zh.wikipedia.org/wiki/Visual_Studio_Code) | Calva[[112\]](https://zh.wikipedia.org/wiki/Clojure#cite_note-112) |

 

### Clojure 开发经验总结

<https://liujiacai.net/blog/2019/04/21/experience-in-clojure/> 



## 实例

网站可以看：  [www.joudou.com.cn](https://link.zhihu.com/?target=http%3A//www.joudou.com.cn/)   (网站的架子是php的，但是数据接口都是Clojure的），微信服务号： 九斗数据  ， 里面各个股票，公告的信息处理，全部使用的Clojure。

ps：我们在招人，公司已经获得IDG等机构PreA轮投资，公司名 九斗数据。坐标：北京三元桥。可接受一半时间远程办公。

Clojure Cookbook



整个后端基本都是clojure实现的，API、推送等等。

已经改名为 http://leancloud.cn

控制台、存储等 API、实时通信都是基于 Clojure 实现的，当然还有一些内部管理系统。

我们这个 code base 在国内应该算得上 Clojure 数一数二的应用。



## 设计原则

Clojure 的设计原则可以概括成 5 个词汇：简单、专注、实用、一致和清晰。这不是我概括的，而是《The joy of clojure》概括的。

- **简单**： 鼓励纯函数，极简的语法（少数 special form），个人也认为 clojure 不能算是多范式的语言（有部分 OO 特性），为了支持多范式引入的复杂度，我们在 C++ 和 Scala 身上都看到了。 
- **专注**：前缀运算符不需要去考虑优先级，也没有什么菱形继承的问题，动态类型系统（有利有弊），REPL 提供的探索式编程方法（告别修改 / 编译 / 运行的死循环，所见即所得）。
- **实用**：前面提到，构建在 JVM 之上，跟 Java 语言的互操作非常容易。直接调用 Java 方法，不去发明一套新的调用语法，努力规避 Java 语言中繁琐的地方 (doto, 箭头宏等等）。
- 清晰：纯函数（前面提到），immutable var，immutable 数据结构，STM 避免锁问题。不可变减少了心智的负担，降低了多线程编程的难度，纯函数也更利于测试和调试。
- **一致**：语法的一致性：例如 doseq 和 for 宏类似，都支持 destructring, 支持相同的 guard 语句（when,while）。数据结构的一致性：sequence 抽象之上的各种高阶函数。
- 具体到 **STM**，我个人认为这个特性在日常编程中，其实你用到的机会不多。在 web 编程里，你的并发模型 Web Container 已经帮你处理（tomcat,jetty），事务也是数据库帮你处理，几乎找不到场合去使用 STM。这个特性在做一些中间件或者底层 framework 的时候才可能用到。这个特性的设计上面已经提到，跟 clojure 的设计目标是紧密相关的，跟 immutable 数据结构也是密不可分，同时它也不是没有代价，事务历史记录和慢事务频繁回滚的代价，有时候你还是需要退回去使用 Java 那套锁机制，庆幸的是 Clojure 不阻止你去使用，并且提供了类似 locking 这样的宏来方便你使用。

## 缺陷

Clojure 的设计缺陷不能说是缺陷，这是由于它设计的目标决定的，有得必有失。

- 首先还是 JVM，基于 JVM 有种种好处，但是 JVM 的启动速度实在悲剧，因此用 Clojure 写一些小的 script 处理日常事务，显得还是不够得心应手，这样的工作我还是用 Ruby，Python 的脚本语言来搞定更便捷。不过目前 Clojure 有一些其他语言之上的实现，比如 [clojure-py](https://github.com/halgari/clojure-py)、[joxa](http://joxa.org/)、[clojureclr](http://clojure.org/clojureclr) 这些实现应该会比 JVM 的启动快很多（抱歉，我没测试过）。
- 不仅如此，因为 Clojure 跟 JVM 平台的绑定如此之深，并且为了真正发挥 Clojure 的威力，你还需要去熟悉 Java 平台的东西，熟悉 Java 语言、类库、内存模型、GC 优化、多线程和网络编程、开源类库等等。可以这样认为：想成为一个好的 Clojure 程序员，首先需要是一名好的 Java 程序员。这也一定程度上阻碍了 Clojure 的推广，提高了学习成本。
- 其次，Clojure 的 API 设计上，有时候不符合你的直觉，而是符合 Clojure 的哲学，比如 `contains?`函数对 vector 等数组型集合的调用上。关于这一点，Rich 的回答是**Elegance and familiarity are orthogonal.**，也就是优雅和熟悉是正交关系的。保持 API 内在的一致性，比直觉的**熟悉**更重要，这是更深入思考、理性的直觉。
- 第三，弱类型的好处足够多，灵活，减少声明代码，适合探索式编程；同样，坏处也不是没有，没有类型保障，错误可能要等到运行时才能发现，静态代码检查工具也没有办法帮你发现，这就需要你一定程度的测试代码来保证运行时行为。
- 第四，性能上，虽然 clojure 生成的字节码已经很高效，也有 type hint 这样的技术来帮助提升性能，但是会有不少的转型 (checkcast)、装箱和拆箱（boxing and unboxing）以及类型判断分支跳转的多余指令，这在一些性能敏感的应用里可能会暴露出来。尽管我认为大多数网站型的应用瓶颈都会落在 IO 上。



[《Clojure in Action》第二版](https://www.manning.com/books/clojure-in-action-second-edition) 

## [Clojure 开发那些事](https://liujiacai.net/blog/2016/12/31/dev-in-clojure/) 



### 高并发

java的锁和CAS都是开销惊人的(锁会调用本地系统进行线程阻塞和唤醒开销巨大).

Actor的消息队列也会有自己的严重问题(使用Actor又会使得编程复杂度大大提高).

所以分布式开发其实是按照业务特性,妥协出最优方案. 



在写少读多的并发场景下：clojure的STM,通过多版本控制以空间换时间,提供了类似MVCC的内存事务回滚, ,简单高效到了极致.

大量写的情况下go/rust在CAS上抽象出的task机制也可以让CAS开发变得相当简单


如果你是 Clojure 新手，并有兴趣深入学习，并且工作前景还过的去，那么我建议你从[ Light Table](http://lighttable.com/)的 [instaREPL](http://www.baidu.com/link?url=3lgmrJCAyAkp4VS3QgTq6V1gy5qPXkvgwNAcDrR_EF6H_gyseU0Ni27nsSEyTPA4) 和一些东西像 [Clojure for the Brave and True](http://www.braveclojure.com/) 开始。



### [Clojure编程语言 – Loretta](https://podtail.com/no/podcast/%E4%BB%A3%E7%A0%81%E6%97%B6%E9%97%B4/clojure%E7%BC%96%E7%A8%8B%E8%AF%AD%E8%A8%80-loretta/ ) 



<https://trends.google.com/trends/explore?date=all&q=scala,Clojure,Closure,%2Fm%2F09gbxjr> 





东南亚

grab go jek lazada agoda

<https://36kr.com/p/5220162> 

[面试line](https://igene.tw/line-tokyo-interview) （[line](https://medium.com/@FreddieWang/line-fukuoka%E7%9A%84%E5%B7%A5%E4%BD%9C%E5%BF%83%E5%BE%97-e47e61468ed3)） 

[名企Google、Line的面试流程大揭秘](http://zone.hirede.com/1723-%e5%90%8d%e4%bc%81google%e3%80%81line%e7%9a%84%e9%9d%a2%e8%af%95%e6%b5%81%e7%a8%8b%e5%a4%a7%e6%8f%ad%e7%a7%98.html) 

Monster.com是Monster Worldwide，Inc。拥有和运营的全球雇佣网站。 

hashmap原理

















