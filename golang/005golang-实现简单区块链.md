[toc] 



# 目标

golang实现简单区块链，编写和运行一个本地的区块链，并且可以在浏览器中查看它。



通过本文，你将可以做到：

- 创建自己的区块链
- 理解 hash 函数是如何保持区块链的完整性
- 如何创造并添加新的块
- 多个节点如何竞争生成块
- 通过浏览器来查看整个链
- 所有其他关于区块链的基础知识

但是，对于比如**工作量证明算法（PoW）** 以及 **权益证明算法（PoS）**这类的共识算法文章中将不会涉及。同时为了让你更清楚的查看区块链以及块的添加，我们将网络交互的过程简化了，关于 P2P 网络比如“全网广播”这个过程等内容将在下一篇文章中补上。

# 1.设置

github.com/davecgh/go-spew/spew

`spew` 可以帮助我们在 `console` 中直接查看 `struct` 和 `slice` 这两种数据结构。

github.com/gorilla/mux

`Gorilla` 的 mux 包非常流行， 我们用它来写 web handler。

github.com/joho/godotenv

godotenv 可以帮助我们读取项目根目录中的 `.env` 配置文件，这样我们就不用将 http port 之类的配置硬编码进代码中

# 2.数据模型

```go
type Block struct {
    Index     int
    Timestamp string
    BPM       int
    Hash      string
    PrevHash  string
}
```

- `Index` 是这个块在整个链中的位置
- `Timestamp` 显而易见就是块生成时的时间戳
- `Hash` 是这个块通过 SHA256 算法生成的散列值
- `PrevHash` 代表前一个块的 SHA256 散列值
- `BPM` 每分钟心跳数，也就是心率。实际的数据BPM

再定义一个结构表示整个链，最简单的表示形式就是一个 `Block` 的 slice：

var Blockchain []Block

使用散列算法（SHA256）来确定和维护链中块和块正确的顺序，确保每一个块的 PrevHash 值等于前一个块中的 Hash 值，这样就以正确的块顺序构建出链：

![](https://raw.githubusercontent.com/cracker8090/imgbed/master/blogImg/20200902190159.png) 

# 3.散列和生成块

需要散列的原因：

- 在节省空间的前提下去唯一标识数据。散列是用整个块的数据计算得出，在我们的例子中，将整个块的数据通过 `SHA256` 计算成一个定长不可伪造的字符串。
- 维持链的完整性。通过存储前一个块的散列值，我们就能够确保每个块在链中的正确顺序。任何对数据的篡改都将改变散列值，同时也就破坏了链。以我们从事的医疗健康领域为例，比如有一个恶意的第三方为了调整“人寿险”的价格，而修改了一个或若干个块中的代表不健康的 BPM 值，那么整个链都变得不可信了。

```go
// 计算区块的SHA256 散列值，calculateHash 函数接受一个块，通过块中的 Index，Timestamp，BPM，以及 PrevHash 值来计算出 SHA256 散列值。
func calculateHash(block Block) string {
    record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}
```

```go
// 生成块的函数,Index 是从给定的前一块的 Index 递增得出，时间戳是直接通过 time.Now() 函数来获得的，Hash 值通过前面的 calculateHash 函数计算得出，PrevHash 则是给定的前一个块的 Hash 值。
func generateBlock(oldBlock Block, BPM int) (Block, error) {
    var newBlock Block

    t := time.Now()
    newBlock.Index = oldBlock.Index + 1
    newBlock.Timestamp = t.String()
    newBlock.BPM = BPM
    newBlock.PrevHash = oldBlock.Hash
    newBlock.Hash = calculateHash(newBlock)

    return newBlock, nil
}
```

# 4.校验块

```go
// 检查 PrevHash 与前一个块的 Hash 是否一致,检查 Index 来看这个块是否正确得递增,通过 calculateHash 检查当前块的 Hash 值是否正确
func isBlockValid(newBlock, oldBlock Block) bool {
    if oldBlock.Index+1 != newBlock.Index {
        return false
    }
    if oldBlock.Hash != newBlock.PrevHash {
        return false
    }
    if calculateHash(newBlock) != newBlock.Hash {
        return false
    }
    return true
}
```

除了校验块以外，我们还会遇到一个问题：两个节点都生成块并添加到各自的链上，那我们应该以谁为准？

通常来说，更长的链表示它的数据（状态）是更新的，所以我们需要一个函数

```go
// 将本地的过期的链切换成最新的链
func replaceChain(newBlocks []Block) {
    if len(newBlocks) > len(Blockchain) {
        Blockchain = newBlocks
    }
}
```

# 5.web服务

```go
// 借助 Gorilla/mux 包,端口号是通过前面提到的 .env 来获得
func run() error {
    mux := makeMuxRouter()
    httpAddr := os.Getenv("ADDR")
    log.Println("Listening on ", os.Getenv("ADDR"))
    s := &http.Server{
        Addr:           ":" + httpAddr,
        Handler:        mux,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    if err := s.ListenAndServe(); err != nil {
        return err
    }

    return nil
}

func makeMuxRouter() http.Handler {
    muxRouter := mux.NewRouter()
    muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
    muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
    return muxRouter
}
// 为了简化，我们直接以 JSON 格式返回整个链
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
    bytes, err := json.MarshalIndent(Blockchain, "", "  ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    io.WriteString(w, string(bytes))
}

type Message struct {
    BPM int
}
// POST handler 接受请求后就能获得请求体中的 BPM 值，接着借助生成块的函数以及校验块的函数就能生成一个新的块
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
    var m Message
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&m); err != nil {
        respondWithJSON(w, r, http.StatusBadRequest, r.Body)
        return
    }
    defer r.Body.Close()
    newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], m.BPM)
    if err != nil {
        respondWithJSON(w, r, http.StatusInternalServerError, m)
        return
    }
    if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
        newBlockchain := append(Blockchain, newBlock)
        replaceChain(newBlockchain)
        spew.Dump(Blockchain)
    }
    respondWithJSON(w, r, http.StatusCreated, newBlock)
}
// 使用spew.Dump 这个函数可以以非常美观和方便阅读的方式将 struct、slice 等数据打印在控制台里，方便我们调试。
// POST 请求处理完之后，无论创建块成功与否，我们需要返回客户端一个响应
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
    response, err := json.MarshalIndent(payload, "", "  ")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("HTTP 500: Internal Server Error"))
        return
    }
    w.WriteHeader(code)
    w.Write(response)
}
```

# 6.main函数

```go
// genesisBlock （创世块）是 main 函数中最重要的部分，通过它来初始化区块链，毕竟第一个块的 PrevHash 是空的。
func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal(err)
    }

    go func() {
        t := time.Now()
        genesisBlock := Block{}
		genesisBlock = Block{0, t.String(), 0, calculateHash(genesisBlock), ""}
		spew.Dump(genesisBlock)
        Blockchain = append(Blockchain, genesisBlock)
    }()
    log.Fatal(run())

}
```



# 总结

它具备块生成、散列计算、块校验等基本能力。还需要学习 工作量证明、权益证明这样的共识算法，或者是智能合约、Dapp、侧链等等。这个实现中不包括任何 P2P 网络的内容。





# 参考

[golang创建一个区块链](http://blog.hubwiz.com/2019/01/17/blockchain-build-go)

[golang实现简单区块链](https://colobu.com/2018/02/05/code-your-own-blockchain-in-less-than-200-lines-of-go/) 

[golang实现简单区块链(英文)](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc) 

[github repo](https://github.com/mycoralhealth/blockchain-tutorial) 

