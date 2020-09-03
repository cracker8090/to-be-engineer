package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
)

func main(){
	// OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666) 存在则会把原来删除
	//newFile ,err := os.Create("xbp.txt")
	//if err != nil{
	//	log.Println(err)
	//	return
	//}
	//newFile.WriteString("xbp is ok36")
	//log.Println(newFile)
	//newFile.Close()

	//err := os.Truncate("xbp.txt",15)
	//if err != nil {
	//	log.Println(err)
	//}

	//fileInfo,err := os.Stat("xbp.txt")
	//_,err := os.Stat("xbp.txt")
	//if err != nil {
	//	log.Println("os stat error:",err)
	//}
	//fmt.Println("File name:", fileInfo.Name())
	//fmt.Println("Size in bytes:", fileInfo.Size())
	//fmt.Println("Permissions:", fileInfo.Mode())
	//fmt.Println("Last modified:", fileInfo.ModTime())
	//fmt.Println("Is Directory: ", fileInfo.IsDir())
	//fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	//fmt.Printf("System info: %+v\n\n", fileInfo.Sys())

	//err := os.Rename("xbp.txt","xbp1.txt")
	//if err != nil {
	//	log.Println(err)
	//}
	//srcFile,err := os.Open("xbp1.txt")
	//if err != nil {
	//	log.Println(err)
	//}
	//defer srcFile.Close()
	//dstFile,err := os.Create("xbp_copy.txt")
	//if err != nil {
	//	log.Println(err)
	//}
	//defer dstFile.Close()
	//bytes,err := io.Copy(dstFile,srcFile)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(bytes)
	//
	//err = dstFile.Sync()
	//if err != nil {
	//	log.Println(err)
	//}
	//dstFile.Close()
	//err = os.Remove("xbp_copy.txt")
	//if err != nil {
	//	log.Println(err)
	//}

	//fileInfo,err := os.Stat("xbp.txt")
	//if err != nil {
	//	if os.IsNotExist(err) {
	//		log.Fatal("File does not exist.")
	//	}
	//}
	//log.Println(fileInfo)

	//os.Chown("xbp1.txt",os.Geteuid(),os.Getegid())

	//file,_:=os.Open("xbp1.txt")
	//defer file.Close()
	//// 偏离位置，可以是正数也可以是负数
	//var offset int64 = 5
	//// 用来计算offset的初始位置
	//// 0 = 文件开始位置
	//// 1 = 当前位置
	//// 2 = 文件结尾处
	//var whence int =0
	//position,err :=file.Seek(offset,whence)
	//if err != nil{
	//	log.Println(err)
	//}
	//log.Println(position)
	//position, err = file.Seek(-2, 1)
	//
	//currentPosition, err := file.Seek(0, 1)
	//fmt.Println("Current position:", currentPosition)
	//
	//position, err = file.Seek(0, 0)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//file,err:=os.OpenFile("xbp1.txt",os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//byteSlice := []byte("Bytes!\n")
	//bytesWritten, err :=file.Write(byteSlice)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("Wrote %d bytes.\n", bytesWritten)

	//err :=ioutil.WriteFile("xbp1.txt",[]byte("HI!\n"),0666)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//file,err:=os.OpenFile("xbp1.txt",os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//bufW := bufio.NewWriter(file)
	//bytes,err:=bufW.Write([]byte{65,66,67})
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(bytes)
	//bytes,err = bufW.WriteString("Buffered string\n")
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(bytes)
	//
	//unflushedBufferSize := bufW.Buffered()
	//log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	//
	//bytesAvailable := bufW.Available()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("Available buffer: %d\n", bytesAvailable)
	//
	//bufW.Flush()


	//file,err:=os.Open("xbp1.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//byteslice := make([]byte,16)
	//n,err:=file.Read(byteslice)
	//dealError(err)
	//log.Println("byteslice:",byteslice)
	//log.Println("n slice:",n)

	//file,err:=os.Open("xbp1.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//byteslice := make([]byte,4)
	//n,err:=io.ReadFull(file,byteslice)
	//dealError(err)
	//log.Println("byteslice:",byteslice)
	//log.Println("n slice:",n)

	//file,err:=os.Open("xbp1.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//byteslice := make([]byte,4)
	//minByte := 3
	//n,err:=io.ReadAtLeast(file,byteslice,minByte)
	//dealError(err)
	//log.Println("byteslice:",byteslice)
	//log.Println("n slice:",n)

	//file,err:=os.Open("xbp1.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//bufR:=bufio.NewReader(file)
	//byteslice := make([]byte,5)
	//byteslice,err = bufR.Peek(5)
	//dealError(err)
	//fmt.Println(byteslice)

	//n,err:=bufR.Read(byteslice)
	//dealError(err)
	//fmt.Println(byteslice,n)
	//myBytes,err := bufR.ReadByte()
	//dealError(err)
	//fmt.Printf("%c\n",myBytes)

	//dataBytes,err:=bufR.ReadBytes('\n')
	//dealError(err)
	//fmt.Println(string(dataBytes),len(dataBytes))

	//dataString,err:=bufR.ReadString('\n')
	//dealError(err)
	//fmt.Println(dataString,len(dataString))

	//file,err:=os.Open("xbp1.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//
	//bufscan := bufio.NewScanner(file)
	//
	//bufscan.Split(bufio.ScanWords)
	//
	//success := bufscan.Scan()
	//if success == false {
	//	err = bufscan.Err()
	//	if err == nil {
	//		log.Println("scan complete and EOF")
	//	} else {
	//		log.Fatal(err)
	//	}
	//}
	//fmt.Println(bufscan.Text())

	//file,err:=os.Create("xbp1.zip")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//zipW:=zip.NewWriter(file)
	//var filesToArchive = []struct {
	//	Name, Body string
	//} {
	//	{"test.txt", "String contents of file"},
	//	{"test2.txt", "\x61\x62\x63\n"},
	//}
	//for _, file := range filesToArchive {
	//	fileWriter, err := zipW.Create(file.Name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	_, err = fileWriter.Write([]byte(file.Body))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//err = zipW.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//newFile,err:=os.Create("devdungeon.html")
	//dealError(err)
	//defer newFile.Close()
	//
	//url := "http://www.devdungeon.com/archive"
	//resp,err:=http.Get(url)
	//dealError(err)
	//defer resp.Body.Close()
	//n,err:=io.Copy(newFile,resp.Body)
	//dealError(err)
	//log.Println(n)

	data, err := ioutil.ReadFile("xbp1.txt")
	dealError(err)

	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))

}

func dealError(err error){
	if err != nil {
		log.Println(err)
	}
}


