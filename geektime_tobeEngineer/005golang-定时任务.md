

定时任务分类



# time包

golang中提供了2种定时器timer和ticker，分别是一次性定时器和重复任务定时器。

- ticker定时器表示每隔一段时间就执行一次，一般可执行多次。
- timer定时器表示在一段时间后执行，默认情况下只执行一次，如果想再次执行的话，每次都需要调用 time.Reset()方法，此时效果类似ticker定时器。同时也可以调用stop()方法取消定时器
- timer定时器比ticker定时器多一个Reset()方法，两者都有Stop()方法，表示停止定时器,底层都调用了stopTimer()函数。





## time.NewTicker

```go
go func() {        
   tt := time.NewTimer(time.Duration(5)*time.Minute)
   for {
      select {
         case <-tt.C :
               nowTime := time.Now().Format("2006-01-02 15:04:05")
               log.Println("当前时间是：", nowTime, "执行定时任务;")                 
               action.ServerBalance()
               tt.Reset(time.Duration(5)*time.Minute)
      }
   }
   }()
```

启动的时候执行一次，以后每天晚上12点执行

```go
func startTimer(f func()) {
    go func() {
        for {
            f()
            now := time.Now()
            // 计算下一个零点
            next := now.Add(time.Hour * 24)
            next = time.Date(next.Year(), next.Month(), next.Day(), 0,0,0,0,next.Location())
            t := time.NewTimer(next.Sub(now))
            <-t.C
        }
    }()
}
```

## ticker

使用定时器每隔12小时从MySQL复制用户信息到Redis数据库

```go
func CopyUserInfo() {
   for {
      rows, err := MysqlClient.Query("SELECT name,mail,department,title FROM UsersInfo")
      if err != nil {
         log4go.Info("query mysqlDB fail")
         return
      }

      userInfos := make(map[int]models.UserInfo)
      userInfo := models.UserInfo{}
      i := 0
      for rows.Next() {
         rows.Scan(&userInfo.Name, &userInfo.Mail, &userInfo.Department, &userInfo.Title)
         userInfos[i] = userInfo
         i++
      }

      SetUserNameMail(userInfos)  //save userInfo into Redis
      SetUserDisplaynameMail(userInfos) //save userInfo into Redis
      fmt.Println("userinfo copy to redis successfully")
      ticker := time.NewTicker(time.Hour * 12)
      <-ticker.C
   }

}
```

## 可关闭定时器

```go
func UserTicker() chan bool {
	ticker := time.NewTicker(5 * time.Second)
	
	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		
		for {
			select {
				case <-ticker.C:
					fmt.Println("Ticker2....")
				case stop := <-stopChan:
					if stop {
						fmt.Println("Ticker2 Stop")
						return
					}
			}
		}
	}(ticker)
	
	return stopChan
通过select读取两个通道，当stop通道读到true的时候，函数返回，由于使用了defer，
会调用Ticker的Stop方法，之后从协程返回
func main() {
    ch := UserTicker()
    time.Sleep(20 * time.Second)
    ch <- true
    close(ch)
}
```

## 打包成函数

可传入函数且带参数，带参数需要定义。

```go
func startTimer(f func(now time.Time)) chan bool {
    done := make(chan bool, 1)
    go func() {
        t := time.NewTimer(time.Second * 3)
        defer t.Stop()
        select {
        case now := <-t.C:
            f(now)
        case <-done:
            fmt.Println("done")
            return
        }
    }()
    return done
}

done := startTimer(func(now time.Time) {
    fmt.Println(now)
})
time.Sleep(1 * time.Second)
close(done)



package ggtimer

import "time"

type TimeCallbackFunc func(time.Time)
type GGTimer chan bool
type GGTicker chan bool

func (t GGTimer) Close() {
    close(t)
}

func (t GGTicker) Close() {
    close(t)
}

func NewTicker(d time.Duration, f TimeCallbackFunc) GGTicker {
    done := make(chan bool, 1)
    go func() {
        t := time.NewTicker(d)
        defer t.Stop()

        for {
            select {
            case now := <-t.C:
                f(now)
            case <-done:
                return
            }
        }
    }()
    return done
}

func NewTimer(d time.Duration, f TimeCallbackFunc) GGTimer {
    done := make(chan bool, 1)
    go func() {
        t := time.NewTimer(d)
        defer t.Stop()
        select {
        case now := <-t.C:
            f(now)
        case <-done:
            return
        }
    }()
    return done
}
```



## 获取前后运行时差

```go
	t1 := time.Now()
	a := "3p4e5ople"
	fmt.Println(verify(&a))

	tt := time.Now()
	fmt.Println(tt.Sub(t1))
```

## 注意

```go
// Ticker 包含一个通道字段C，每隔时间段 d 就向该通道发送当时系统时间。
    // 它会调整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。
    // 如果d <= 0会触发panic。关闭该 Ticker 可以释放相关资源。

	ticker1 := time.NewTicker(5 * time.Second)
	// 一定要调用Stop()，回收资源
	defer ticker1.Stop()
	go func(t *time.Ticker) {
		for {
			// 每5秒中从chan t.C 中读取一次
			<-t.C
			fmt.Println("Ticker:", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker1)
```

如果在调用 time.Reset() 或time.Stop() 的时候，timer已经过期或者停止了，则会返回false。

使用time.NewTicker() 定时器时，需要使用Stop()方法进行资源释放，否则会产生内存泄漏，(Stop the ticker to release associated resources.)



# cron包

$ go get -v -u github.com/robfig/cron

### Cron 的定时任务表达式

```
 1 ┌─────────────second 范围 (0 - 60)
 2 │ ┌───────────── min (0 - 59)
 3 │ │ ┌────────────── hour (0 - 23)
 4 │ │ │ ┌─────────────── day of month (1 - 31)
 5 │ │ │ │ ┌──────────────── month (1 - 12)
 6 │ │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
 7 │ │ │ │ │ │                  Saturday)
 8 │ │ │ │ │ │
 9 │ │ │ │ │ │
10 * * * * * *
```



cron 表达式代表一个时间的集合，使用 6 个空格分隔的字段表示：

| 字段名            | 是否必须 | 允许的值        | 允许的特定字符 |
| ----------------- | -------- | --------------- | -------------- |
| 秒(Seconds)       | 是       | 0-59            | * / , -        |
| 分(Minute)        | 是       | 0-59            | * / , -        |
| 时(Hours)         | 是       | 0-23            | * / , -        |
| 日(Day of month)  | 是       | 1-31            | * / , - ?      |
| 月(Month)         | 是       | 1-12 或 JAN-DEC | * / , -        |
| 星期(Day of week) | 否       | 0-6 或 SUM-SAT  | * / , - ?      |



### 时间表达式例子

`"0 0 0 1 * *"` // 表示每个月1号的00:00:00

 `"0 1 1 1 * *"` // 表示每个月1号的01:01:00

每隔 5 秒执行一次：*/5 * * * * ?

每隔 1 分钟执行一次：0 */1 * * * ?

每天 23 点执行一次：0 0 23 * * ?

每天凌晨 1 点执行一次：0 0 1 * * ?

每月 1 号凌晨 1 点执行一次：0 0 1 1 * ?

在 26 分、29 分、33 分执行一次：0 26,29,33 * * * ?

每天的 0 点、13 点、18 点、21 点都执行一次：0 0 0,13,18,21 * * ?

```go
	c := cron.New()
    c.AddFunc("*/3 * * * * *", func() {
        fmt.Println("every 3 seconds executing")
    })
 
    go c.Start()
    defer c.Stop()
 
 
    select {
    case <-time.After(time.Second * 10):
        return
    }
//select{}
```

- 月(Month)和星期(Day of week)字段的值不区分大小写，如：SUN、Sun 和 sun 是一样的。

- 星期(Day of week)字段如果没提供，相当于是 *

  

- cron.New创建一个定时器管理器
- c.AddFunc添加一个定时任务，第一个参数是cron时间表达式，第二个参数是要触发执行的函数
- go c.Start()新启一个协程，运行定时任务
- c.Stop是等待停止信号结束任务



## 文件目录

```go
constantdelay.go      #一个最简单的秒级别定时系统。与cron无关
constantdelay_test.go #测试
cron.go               #Cron系统。管理一系列的cron定时任务（Schedule Job）
cron_test.go          #测试
doc.go                #说明文档
LICENSE               #授权书 
parser.go             #解析器，解析cron格式字符串城一个具体的定时器（Schedule）
parser_test.go        #测试
README.md             #README
spec.go               #单个定时器（Schedule）结构体。如何计算自己的下一次触发时间
spec_test.go          #测试
```

```go
type Cron struct {
    entries  []*Entry
    stop     chan struct{}   // 控制 Cron 实例暂停
    add      chan *Entry     // 当 Cron 已经运行了，增加新的 Entity 是通过 add 这个 channel 实现的
    snapshot chan []*Entry   // 获取当前所有 entity 的快照
    running  bool            // 当已经运行时为true；否则为false
}
//注意：有一个成员 stop，类型是 struct{}，即空结构体。
```

```go
 1// 将 job 加入 Cron 中
 2// 如上所述，该方法只是简单的通过 FuncJob 类型强制转换 cmd，然后调用 AddJob 方法
 3func (c *Cron) AddFunc(spec string, cmd func()) error
 4
 5// 将 job 加入 Cron 中
 6// 通过 Parse 函数解析 cron 表达式 spec 的到调度器实例(Schedule)，之后调用 c.Schedule 方法
 7func (c *Cron) AddJob(spec string, cmd Job) error
 8
 9// 获取当前 Cron 总所有 Entities 的快照
10func (c *Cron) Entries() []*Entry
11
12// 通过两个参数实例化一个 Entity，然后加入当前 Cron 中
13// 注意：如果当前 Cron 未运行，则直接将该 entity 加入 Cron 中；
14// 否则，通过 add 这个成员 channel 将 entity 加入正在运行的 Cron 中
15func (c *Cron) Schedule(schedule Schedule, cmd Job)
16
17// 新启动一个 goroutine 运行当前 Cron
18func (c *Cron) Start()
19
20// 通过给 stop 成员发送一个 struct{}{} 来停止当前 Cron，同时将 running 置为 false
21// 从这里知道，stop 只是通知 Cron 停止，因此往 channel 发一个值即可，而不关心值是多少
22// 所以，成员 stop 定义为空 struct
23func (c *Cron) Stop()
```

实例

```go
 3import (
 4    "log"
 5
 6    "github.com/robfig/cron"
 7)
 8
 9type Hello struct {
10    Str string
11}
12
13func(h Hello) Run() {
14    log.Println(h.Str)
15}
16
17func main() {
18    log.Println("Starting...")
19
20    c := cron.New()
21    h := Hello{"I Love You!"}
22    // 添加定时任务
23    c.AddJob("*/2 * * * * * ", h)
24    // 添加定时任务
25    c.AddFunc("*/5 * * * * * ", func() {
26        log.Println("hello word")
27    })
28
29    s, err := cron.Parse("*/3 * * * * *")
30    if err != nil {
31        log.Println("Parse error")
32    }
33    h2 := Hello{"I Hate You!"}
34    c.Schedule(s, h2)
35    // 其中任务
36    c.Start()
37    // 关闭任务
38    defer c.Stop()
39    select {
40    }
41}
```



## 停止定时器

由于goroutine没有线程id，所以无法从外部停止指定的定时任务。cron自带Stop()方法，支持在方法体里使用，来停止当前定时。

何时调用Stop()方法，需要根据自身业务触发，例如运行完成100次后停止，数据库当前状态为注销时停止，等等。

```go
func main() {
  c := cron.New(cron.WithSeconds()) //精确到秒
  
  //定时任务
  spec := "*/1 * * * * ?" //cron表达式，每秒一次
  c.AddFunc(spec, func() {
    status := getStatus() //获取定时任务的状态
    if status == true {
      fmt.Println("11111")
    } else { 
      c.Stop() //当前定时任务状态已注销
    }
  })
  
  c.Start()
  select {}  //阻塞主线程停止
}
```

一个c对象可以加载多个定时任务，此时在其中一个方法里调用Stop()方法，所有定时任务都会停止。



```
// 也可以在运行的Cron中添加任务
c.AddFunc("@daily", func() { fmt.Println("Every day") })
..
// 检查cron任务条目的下一个和上一个运行时间
inspect(c.Entries())
```



# 定时器问题



## ticker的值拷贝问题

```go
func tickerTest1() {
	ticker := *time.NewTicker(time.Second)
	count := 0
	go func() {
		time.Sleep(3*time.Second)
		ticker.Stop()
	}()
	for range ticker.C {
		count ++
		fmt.Println("tickerTest1：", count)
	}
}

func tickerTest2() {
	ticker := time.NewTicker(time.Second)
	count := 0
	go func() {
		time.Sleep(3*time.Second)
		ticker.Stop()
	}()
	for range ticker.C {
		count ++
		fmt.Println("tickerTest2：", count)
	}
}

	go tickerTest1()
	tickerTest2()
```

tickerTest2： 1
tickerTest1： 1
tickerTest2： 2
tickerTest1： 2
tickerTest2： 3
tickerTest1： 3
tickerTest1： 4
tickerTest1： 5```

故`tmp.C`和`ticker.C`的值是一样的，故`for range ticker.C`是能够接受到值，即本质是`tmp.C`发送的值，因此可以持续输出；而`tmp.r`与`ticker.r`的地址是不一样的。而在在`runtime`包中，发现`stop`的时候有这么一个判断：

```
if i < 0 || i > last || tb.t[i] != t { //tb.t[i]就是ticker.r的地址，t就是tmp.r的地址
    return false
}
复制代码
```

关键在于`tb.t[i] != t`，则这个条件肯定为真，即不从最小堆四叉树中移除，相当于没有 `stop`。所以计时器没有停止。



## ticker.Stop()并不会关闭Channel

```go
func UseTickerWrong() *time.Ticker {
	ticker := time.NewTicker(5 * time.Second)
	go func(ticker *time.Ticker) {
		for range ticker.C {
			fmt.Println("Ticker1....")
		}
		
		fmt.Println("Ticker1 Stop")
	}(ticker)
	
	return ticker
}
	ticker1 := UseTickerWrong()
	time.Sleep(20 * time.Second)
	ticker1.Stop()
//没有打印Ticker1 Stop
```

首先，创建`Ticker`的协程并不负责计时，只负责从`Ticker`的管道中获取事件；其次，系统协程只负责定时器计时，向管道中发送事件，并不关心上层协程如何处理事件。即我们在创建timer的时候会构建`runtimeTimer`对象，里面有`sendTime`回调方法及初始化的`channel`。`timerproc`是`golang runtime`的定时扫描器，当发现有任务到期后，进行相应的方法回调。但如果我们在stop里把channel给关闭了，那么`timerproc`有可能就`panic`了。

## 在定时器触发之前，gc不会回收 Timer

```go
go func() {
	for {
		timerC := time.After(2 * time.Second)
		select {
		case num := <-ch:
			fmt.Println("get num is ", num)
		case <-timerC:
			fmt.Println("time's up !!!")
		}
	}
}()
```

每循环一次，`timerC`都是重新创建的，即`timerC`只对一个`select`有效。下次处理时重新计算超时时间，每个操作的处理都是独立的。 若`ch`触发的频道远远高于`timerC`的触发频道，将会导致`timerC`不停地被创建。但由于在定时器触发之前，`gc`是不会对`timerC`进行回收。因此造成了大量的内存浪费（由于`timerC`还在四叉树里被引用着，故不能进行`gc`回收）。

```go
go func() {
	idleDuration := 5 * time.Second
	idleDelay:=time.NewTimer(idleDuration)
	defer  idleDelay.Stop()
	for {
		idleDelay.Reset(idleDuration) //重置定时器
		select {
		case num := <-ch:
			fmt.Println("get num is ", num)
		case <-idleDelay.C:
			fmt.Println("time's up !!!")
		}
	}
}()
```

## 创建的周期性定时器，若不进行主动`stop`，将会导致内存泄漏

使用time.NewTicker() 定时器时，需要使用Stop()方法进行资源释放，否则会产生内存泄漏。

如果调用 time.Stop() 时，timer已过期或已stop，则并不会关闭通道。

创建一个`ticker`，应该紧跟着使用`defer ticker.Stop()`语句，当函数退出时自动从最小堆四叉树中移除。

```go
func  run() {
	timeout := time.NewTicker(p.Timeout)
	defer timeout.Stop()  
	interval := time.NewTicker(p.Interval)
	defer interval.Stop() 
	for {
		select {
		case <-p.done:       
			wg.Wait()
			return
		case <-timeout.C:   
			close(p.done)
			wg.Wait()
			return
		case <-interval.C:
			if p.Count > 0 && p.PacketsSent >= p.Count {
				continue
			}
			err = p.sendICMP(conn)
			if err != nil {
				fmt.Println("FATAL: ", err.Error())
			}
		case r := <-recv:
			err := p.processPacket(r)
			if err != nil {
				fmt.Println("FATAL: ", err.Error())
			}
		}
		if p.Count > 0 && p.PacketsRecv >= p.Count { 
			close(p.done)
			wg.Wait()
			return
		}
	}
}
```

## timer引起的内存泄漏和CPU问题

在高性能场景下，不应该使用time.After，而应该使用New.Timer并在不再使用该Timer后（无论该Timer是否被触发）调用Timer的Stop方法来及时释放资源。不然内存资源可能被延时释放。

使用time.NewTicker时，在Ticker对象不再使用后（无论该Ticker是否被触发过），一定要调用Stop方法，否则会造成内存和cpu泄漏。



After函数等待参数duration所指定的时间到达后才返回，返回的chan中Time的值为当前时间。After函数等价于time.NewTimer(d).C。



## 高性能时间轮

golang在1.9.x版本后优化了定时器，以前为一个四叉堆，现在是多个四叉堆，具体你能用到多少个堆，要看你的gomaxprocs配置大小了。go timer init里初始化了64个timerBucket，那么time会如何选择timerBucket? 是通过g.m.p.id % 64取摸拿到的timerBucket。

```go
func (t *timer) assignBucket() *timersBucket {
    id := uint8(getg().m.p.ptr().id) % timersLen
    t.tb = &timers[id].timersBucket
    return t.tb
}
```

golang标准库里实现的定时器都是高精度的，高精度带来了频发的读写操作heap和锁的竞争压力。

时间轮的意义在于把精度放低，比如同一秒内的定时任务放凑在一起，数据结构改用map, 那么时间轮只需要o(1)的时间复杂度就可以把任务放进去，只需要一次加锁放锁就可以把同一个时间精度的任务都捞出来。

https://github.com/rfyiamcool/go-timewheel

```go
// 初始化时间轮, 精度为1s，槽位360个
tw, err := NewTimeWheel(1 * time.Second, 360)
if err != nil {
    panic(err)
}

// 启动时间轮
tw.Start()

// 关闭时间轮
tw.Stop()

// 添加任务
task := tw.Add(5 * time.Second, func(){})

// 删除任务
tw.Remove(task)

// 添加周期性任务
task := tw.AddCron(5 * time.Second, func(){
    ...
})

// 实现time.Sleep
tw.Sleep(5 * time.Second)
similar to time.After()

// 实现After
<- tw.After(5 * time.Second)
similar to time.NewTimer

// 实现NewTimer定时器
timer :=tw.NewTimer(5 * time.Second)
<- timer.C
timer.Reset(1 * time.Second)
timer.Stop()

// 实现NewTicker定时器
timer :=tw.NewTicker(5 * time.Second)
<- timer.C
timer.Stop()


// 实现go time的AfterFunc
runner :=tw.AfterFunc(5 * time.Second, func(){})
<- runner.C
runner.Stop()
```

如果性能还是达不到要求，可以封装多个时间轮到一个池子，类似go time标准库里timerprc的实现。

时间轮的效率确实很高，在一个高频的推送服务里会有各种定时器的使用，时间轮可以减少cpu消耗。





# 参考

https://ld246.com/article/1534497226108

https://github.com/ouqiang/gocron

https://juejin.im/post/6844903901418749960

https://blog.haohtml.com/archives/19859

https://juejin.im/post/6844904164577771533#heading-2

[定时器未释放](https://www.mdeditor.tw/pl/235s) 

[时间轮](http://xiaorui.cc/archives/6160) 

[golang 定时任务方面time.Sleep和time.Tick的优劣对比](https://aimuke.github.io/go/2019/12/13/go-timer/) 

[timer机制](https://aimuke.github.io/go/2019/12/12/go-timer-ticker/) 

[linux下cron使用](https://www.cnblogs.com/mq0036/p/12930998.html) 