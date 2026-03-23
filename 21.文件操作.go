package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 文件读取
	//byteData, err := os.ReadFile("hello.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(byteData) )

	// 分片读取
	//file, err := os.Open("hello.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//for {
	//	var byteData = make([]byte, 1024)
	//	n, err := file.Read(byteData)
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(byteData[:n]))
	//
	//}
	// 带缓冲读
	//file, err := os.Open("hello.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	// 按行读
	//buf := bufio.NewReader(file)
	//
	//for {
	//	line, _, err := buf.ReadLine()
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(line))
	//}

	// 按分隔符读
	//buf := bufio.NewScanner(file)
	//buf.Split(bufio.ScanWords)
	//var index int
	//for buf.Scan() {
	//	index++
	//	fmt.Println(index, buf.Text())
	//}

	//文件写入
	//file, err := os.OpenFile("w.txt", os.O_CREATE|os.O_RDONLY, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	////file.Write([]byte("你好"))
	//byteData, err := io.ReadAll(file)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(byteData))

	//文件复制

	rFile, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rFile.Close()
	wFile, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wFile.Close()
	io.Copy(wFile, rFile)

	// 目录操作
	dir, err := os.ReadDir("../go-playground")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, entry := range dir {
		info, _ := entry.Info()

		fmt.Println(entry.IsDir(), entry.Name(), info.Size())
	}
}
