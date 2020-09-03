go语言中的reader和writer接口也类似。我们只需简单的读写字节，不必知道reader的数据来自哪里，也不必知道writer将数据发送到哪里。



# 基本操作
## 创建空文件

```go
// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666) 存在则会把原来删除
newFile ,err := os.Create("xbp.txt")
```

## Truncate文件

如果文件本来就少于100个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。

// 如果文件本来超过100个字节，则超过的字节会被抛弃。
    // 这样我们总是得到精确的100个字节的文件。 fileinfo.size()
    // 传入0则会清空文件。

```go
err := os.Truncate("xbp.txt",10)
```



## 文件信息

```go
fileInfo,err := os.Stat("xbp.txt")
if err != nil {
   log.Println(err)
}
fmt.Println("File name:", fileInfo.Name())
fmt.Println("Size in bytes:", fileInfo.Size())
fmt.Println("Permissions:", fileInfo.Mode())
fmt.Println("Last modified:", fileInfo.ModTime())
fmt.Println("Is Directory: ", fileInfo.IsDir())
fmt.Printf("System interface type: %T\n", fileInfo.Sys())
fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
```

## 重命名和移动

```go
err := os.Rename("xbp.txt","xbp1.txt")
```

## 删除文件

```go
err = os.Remove("xbp_copy.txt")
```

## 打开和关闭文件

```go
// 简单地以只读的方式打开。下面的例子会介绍读写的例子。
    file, err := os.Open("test.txt")
// OpenFile提供更多的选项
file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
// os.O_CREATE|os.O_TRUNC|os.O_WRONLY
// os.O_RDONLY // 只读
// os.O_WRONLY // 只写
// os.O_RDWR // 读写
// os.O_APPEND // 往文件中添建（Append）
// os.O_CREATE // 如果文件不存在则先创建
// os.O_TRUNC // 文件打开时裁剪文件
// os.O_EXCL // 和O_CREATE一起使用，文件不能存在
// os.O_SYNC // 以同步I/O的方式打开
```

## 文件是否存在

```go
fileInfo,err := os.Stat("xbp.txt")
if err != nil {
   if os.IsNotExist(err) {
      log.Fatal("File does not exist.")
   }
}
```

## 检测读写权限

```go
file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666) // O_RDONLY
if err != nil {
    if os.IsPermission(err) {
        log.Println("Error: Write permission denied.")
    }
}
//如果没有写权限则返回error。
// 注意文件不存在也会返回error，需要检查error的信息来获取到底是哪个错误导致。
```

## 改变权限、拥有者、时间戳

```go
err := os.Chmod("test.txt", 0777)
if err != nil {
    log.Println(err)
}

err = os.Chown("test.txt", os.Getuid(), os.Getgid())
if err != nil {
    log.Println(err)
}

// 改变时间戳
twoDaysFromNow := time.Now().Add(48 * time.Hour)
lastAccessTime := twoDaysFromNow
lastModifyTime := twoDaysFromNow
err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
if err != nil {
    log.Println(err)
}

```

## 硬链接和软链接

一个普通的文件是一个指向硬盘的inode的地方。
硬链接创建一个新的指针指向同一个地方。只有所有的链接被删除后文件才会被删除。硬链接只在相同的文件系统中才工作。你可以认为一个硬链接是一个正常的链接。

symbolic link，又叫软连接，和硬链接有点不一样，它不直接指向硬盘中的相同的地方，而是通过名字引用其它文件。他们可以指向不同的文件系统中的不同文件。并不是所有的操作系统都支持软链接。

```go
// 创建一个硬链接。
// 创建后同一个文件内容会有两个文件名，改变一个文件的内容会影响另一个。
// 删除和重命名不会影响另一个。
err := os.Link("original.txt", "original_also.txt")
if err != nil {
    log.Fatal(err)
}

fmt.Println("creating sym")
// Create a symlink
err = os.Symlink("original.txt", "original_sym.txt")
if err != nil {
    log.Fatal(err)
}

// Lstat返回一个文件的信息，但是当文件是一个软链接时，它返回软链接的信息，而不是引用的文件的信息。
// Symlink在Windows中不工作。
fileInfo, err := os.Lstat("original_sym.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Link info: %+v", fileInfo)

//改变软链接的拥有者不会影响原始文件。
err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())
if err != nil {
    log.Fatal(err)
}
```













# 读写

## 复制文件

```go
srcFile,err := os.Open("xbp1.txt")
if err != nil {
   log.Println(err)
}
defer srcFile.Close()
dstFile,err := os.Create("xbp_copy.txt")
if err != nil {
   log.Println(err)
}
defer dstFile.Close()
bytes,err := io.Copy(dstFile,srcFile)
if err != nil {
   log.Println(err)
}
log.Println(bytes)

err = dstFile.Sync()
if err != nil {
   log.Println(err)
}
```

## 跳到文件指定位置

```go
	file, _ := os.Open("test.txt")
    defer file.Close()

    // 偏离位置，可以是正数也可以是负数
    var offset int64 = 5

    // 用来计算offset的初始位置
    // 0 = 文件开始位置
    // 1 = 当前位置
    // 2 = 文件结尾处
    var whence int = 0
    newPosition, err := file.Seek(offset, whence)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Just moved to 5:", newPosition)

    // 从当前位置回退两个字节
    newPosition, err = file.Seek(-2, 1)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Just moved back two:", newPosition)

    // 使用下面的技巧得到当前的位置
    currentPosition, err := file.Seek(0, 1)
    fmt.Println("Current position:", currentPosition)

    // 转到文件开始处
    newPosition, err = file.Seek(0, 0)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Position after seeking 0,0:", newPosition)
```

## 写文件

因为Go可执行包是静态链接的可执行文件，你import的每一个包都会增加你的可执行文件的大小。其它的包如`io`、｀ioutil｀、｀bufio｀提供了一些方法，但是它们不是必须的。

```go
file,err:=os.OpenFile("xbp1.txt",os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0666)
if err != nil {
   log.Fatal(err)
}
defer file.Close()
byteSlice := []byte("Bytes!\n")
bytesWritten, err :=file.Write(byteSlice)
if err != nil {
   log.Fatal(err)
}
log.Printf("Wrote %d bytes.\n", bytesWritten)
```

## 快写文件

`ioutil`包有一个非常有用的方法`WriteFile()`可以处理创建／打开文件、写字节slice和关闭文件一系列的操作。如果你需要简洁快速地写字节slice到文件中，你可以使用它。

```go
err :=ioutil.WriteFile("xbp1.txt",[]byte("HI!\n"),0666)
if err != nil {
   log.Fatal(err)
}
```

## 缓存写

`bufio`包提供了带缓存功能的writer，所以你可以在写字节到硬盘前使用内存缓存。当你处理很多的数据很有用，因为它可以节省操作硬盘I/O的时间。在其它一些情况下它也很有用，比如你每次写一个字节，把它们攒在内存缓存中，然后一次写入到硬盘中，减少硬盘的磨损以及提升性能。

```go
    file,err:=os.OpenFile("xbp1.txt",os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
    // buffered writer
    bufferedWriter := bufio.NewWriter(file)

    // write to buffer
    bytesWritten, err := bufferedWriter.Write(
        []byte{65, 66, 67},
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Bytes written: %d\n", bytesWritten)

    // also WriteRune() 和 WriteByte()   
    bytesWritten, err = bufferedWriter.WriteString(
        "Buffered string\n",
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Bytes written: %d\n", bytesWritten)

    // check bytes
    unflushedBufferSize := bufferedWriter.Buffered()
    log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

    // Available bytes
    bytesAvailable := bufferedWriter.Available()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Available buffer: %d\n", bytesAvailable)

    // buffer write to hard
    bufferedWriter.Flush()

    // reset flush buffer，clean error and new writer
    bufferedWriter.Reset(bufferedWriter)

    bytesAvailable = bufferedWriter.Available()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Available buffer: %d\n", bytesAvailable)

    // writersize default，4096。
    // resize
    bufferedWriter = bufio.NewWriterSize(bufferedWriter,8000)
```

## 最多读取N字节

```go
	file, err := os.Open("test.txt")
	// 从文件中读取len(b)字节的文件。
    // 返回0字节意味着读取到文件尾了
    // 读取到文件会返回io.EOF的error
    byteSlice := make([]byte, 16)
    bytesRead, err := file.Read(byteSlice)
```

## 正好N字节

```go
	// file.Read()可以读取一个小文件到大的byte slice中，
    // 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误
    byteSlice := make([]byte, 2)
    numBytesRead, err := io.ReadFull(file, byteSlice)
```

## 至少N字节

```go
	byteSlice := make([]byte, 512)
    minBytes := 8
    // io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的文件保留
    numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
```

## 全部读取

```go
	// os.File.Read(), io.ReadFull() 和
    // io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
    // 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
    data, err := ioutil.ReadAll(file)
```

## 快读到内存

```go
// 读取文件到byte slice中
    data, err := ioutil.ReadFile("test.txt")
```

## 缓存读

缓存reader会把一些内容缓存在内存中。它会提供比`os.File`和`io.Reader`更多的函数,缺省的缓存大小是4096，最小缓存是16。

```go
	bufferedReader := bufio.NewReader(file)
    // 得到字节，当前指针不变
    byteSlice := make([]byte, 5)
    byteSlice, err = bufferedReader.Peek(5) //不够则会返回error

    // 读取，指针同时移动
    numBytesRead, err := bufferedReader.Read(byteSlice)

    // 读取一个字节, 如果读取不成功会返回Error
    myByte, err := bufferedReader.ReadByte()

    // 读取到分隔符，包含分隔符，返回byte slice
    dataBytes, err := bufferedReader.ReadBytes('\n')
      
    // 读取到分隔符，包含分隔符，返回字符串
    dataString, err := bufferedReader.ReadString('\n')
```

## sanner使用

`Scanner`是`bufio`包下的类型,在处理文件中以分隔符分隔的文本时很有用。

通常使用换行符作为分隔符将文件内容分成多行。在CSV文件中，逗号一般作为分隔符。

`os.File`文件可以被包装成`bufio.Scanner`，它就像一个缓存reader。

我们会调用`Scan()`方法去读取下一个分隔符，使用`Text()`或者`Bytes()`获取读取的数据。

分隔符可以不是一个简单的字节或者字符，有一个特殊的方法可以实现分隔符的功能，以及将指针移动多少，返回什么数据。
如果没有定制的`SplitFunc`提供，缺省的`ScanLines`会使用`newline`字符作为分隔符，其它的分隔函数还包括`ScanRunes`和`ScanWords`,皆在`bufio`包中。

```go
	scanner := bufio.NewScanner(file)

    // 缺省的分隔函数是bufio.ScanLines,这里使用ScanWords。
    // 也可以定制一个SplitFunc类型的分隔函数
    scanner.Split(bufio.ScanWords)

    // scan下一个token.
    success := scanner.Scan()
    if success == false {
        // 出现错误或者EOF是返回Error
        err = scanner.Err()
        if err == nil {
            log.Println("Scan completed and reached EOF")
        } else {
            log.Fatal(err)
        }
    }

    // 得到数据，Bytes() 或者 Text()
    fmt.Println("First word found:", scanner.Text())

    // 再次调用scanner.Scan()发现下一个token
```




# 压缩



## 打包(zip) 文件

```go
    outFile, err := os.Create("test.zip")

    zipWriter := zip.NewWriter(outFile)

    // 往打包文件中写文件。
    // 这里我们使用硬编码的内容，你可以遍历一个文件夹，把文件夹下的文件以及它们的内容写入到这个打包文件中。
    var filesToArchive = []struct {
        Name, Body string
    } {
        {"test.txt", "String contents of file"},
        {"test2.txt", "\x61\x62\x63\n"},
    }

    // 下面将要打包的内容写入到打包文件中，依次写入。
    for _, file := range filesToArchive {
            fileWriter, err := zipWriter.Create(file.Name)
            if err != nil {
                    log.Fatal(err)
            }
            _, err = fileWriter.Write([]byte(file.Body))
            if err != nil {
                    log.Fatal(err)
            }
    }

    // 清理
    err = zipWriter.Close()
```

## 解压unzip

```go
	zipReader, err := zip.OpenReader("test.zip")
    defer zipReader.Close()

    // 遍历打包文件中的每一文件/文件夹
    for _, file := range zipReader.Reader.File {
        // 打包文件中的文件就像普通的一个文件对象一样
        zippedFile, err := file.Open()
        defer zippedFile.Close()

        // 指定抽取的文件名。
        // 你可以指定全路径名或者一个前缀，这样可以把它们放在不同的文件夹中。
        // 我们这个例子使用打包文件中相同的文件名。
        targetDir := "./"
        extractedFilePath := filepath.Join(
            targetDir,
            file.Name,
        )

        // 抽取项目或者创建文件夹
        if file.FileInfo().IsDir() {
            // 创建文件夹并设置同样的权限
            log.Println("Creating directory:", extractedFilePath)
            os.MkdirAll(extractedFilePath, file.Mode())
        } else {
            //抽取正常的文件
            log.Println("Extracting file:", file.Name)
            outputFile, err := os.OpenFile(
                extractedFilePath,
                os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
                file.Mode(),
            )
            defer outputFile.Close()

            // 通过io.Copy简洁地复制文件内容
            _, err = io.Copy(outputFile, zippedFile)
        }
    }
```

## 压缩文件

```go
	outputFile, err := os.Create("test.txt.gz")

    gzipWriter := gzip.NewWriter(outputFile)
    defer gzipWriter.Close()

    // 当我们写如到gizp writer数据时，它会依次压缩数据并写入到底层的文件中。
    // 我们不必关心它是如何压缩的，还是像普通的writer一样操作即可。
    _, err = gzipWriter.Write([]byte("Gophers rule!\n"))
```

## 解压缩文件

```go
	// 打开一个gzip文件。
    // 文件是一个reader,但是我们可以使用各种数据源，比如web服务器返回的gzipped内容，
    // 它的内容不是一个文件，而是一个内存流
    gzipFile, err := os.Open("test.txt.gz")

    gzipReader, err := gzip.NewReader(gzipFile)
    defer gzipReader.Close()

    // 解压缩到一个writer,它是一个file writer
    outfileWriter, err := os.Create("unzipped.txt")
    defer outfileWriter.Close()

    // 复制内容
    _, err = io.Copy(outfileWriter, gzipReader)
```



# 其他

## 临时文件和目录

ioutil`提供了两个函数: `TempDir()` 和 `TempFile()，使用完毕后，调用者负责删除这些临时文件和文件夹。

有一点好处就是当你传递一个空字符串作为文件夹名的时候，它会在操作系统的临时文件夹中创建这些项目（/tmp on Linux）。`os.TempDir()`返回当前操作系统的临时文件夹。

```go
	// 在系统临时文件夹中创建一个临时文件夹
     tempDirPath, err := ioutil.TempDir("", "myTempDir")

     tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")

     // ... do something ...
     err = tempFile.Close()
     err = os.Remove(tempFile.Name())
     err = os.Remove(tempDirPath)
```

## http下载

```go
	newFile, err := os.Create("devdungeon.html")
     defer newFile.Close()

     url := "http://www.devdungeon.com/archive"
     response, err := http.Get(url)
     defer response.Body.Close()

     // Body满足reader接口，因此我们可以使用ioutil.Copy
     numBytesWritten, err := io.Copy(newFile, response.Body)
```

## 哈希和摘要

```go
    data, err := ioutil.ReadFile("test.txt")

    fmt.Printf("Md5: %x\n\n", md5.Sum(data))
    fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
    fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
    fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
```

另一个方式是创建一个hash writer, 使用`Write`、`WriteString`、`Copy`将数据传给它。
下面的例子使用 md5 hash,但你可以使用其它的Writer。

```go
	file, err := os.Open("test.txt")
    defer file.Close()

    //new hasher,writer interface
    hasher := md5.New()
    _, err = io.Copy(hasher, file)

    // 传递 nil 作为参数，因为我们不通参数传递数据，而是通过writer接口。
    sum := hasher.Sum(nil)
```




# 参考
1. 原文: [Working with Files in Go](http://www.devdungeon.com/content/working-files-go#write_bytes), 作者: [NanoDano](http://www.devdungeon.com/blogs/nanodano)  

2. [https://colobu.com/2016/10/12/go-file-operations/](https://colobu.com/2016/10/12/go-file-operations/) 