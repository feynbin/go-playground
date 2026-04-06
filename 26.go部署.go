package main

// go项目的部署特别简单，编写完成后，只需要执行go build 即可打包为可执行文件

// go build xxx.go

// go build -o main.exe xxx.go

// 交叉编译
//set CGO_ENABLED=0
//set GOOS=linux
//set GOARCH=amd64
//go build -o main main.go

//去掉符号表和调试信息，去掉DWARF调试信息
//go build -ldflags="-s -w" -o myapp main.go