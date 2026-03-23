package main

import "fmt"

type Code int

func (c Code) GetMsg() string {
	switch c {
	case SuccessCode:
		return "success"
	case ServiceErrCode:
		return "service error"
	case NetworkErrCode:
		return "network error"
	}
	return ""
}

const (
	SuccessCode    Code = 0
	ServiceErrCode Code = 1001 //服务错误
	NetworkErrCode Code = 1002 // 网络错误

)

func getCodeMsg(code Code) (msg string) {
	switch code {
	case SuccessCode:
		return "success"
	case ServiceErrCode:
		return "服务错误"
	case NetworkErrCode:
		return "网络错误"
	}
	return
}

func (c Code) Ok() (code Code, msg string) {
	return c, c.GetMsg()
}
func webServer(name string) (code Code, msg string) {
	if name == "1" {
		return ServiceErrCode.Ok()

	}
	if name == "2" {
		return NetworkErrCode.Ok()

	}
	return

}

type MyCode int
type MyAliasCode = int

func (m MyCode) Code() {

}

const myCode MyCode = 1
const myAliasCode MyAliasCode = 2

func main() {
	// 在go语言中，自定义类型指的是使用type关键字定义的新类型，
	// 它可以是基本类型的别名，也可以是结构体、函数等组合而成的新类型。
	// 自定义类型可以帮助我们更好的抽象和封装数据，让代码更加易读，易懂，易维护。

	// 类型别名，和自定义类型很像，但是有一些地方和自定义类型有很大差异
	// 1.不能绑定方法，2.打印类型还是原始类型，3.和原始类型比较，类型别名不用转换。
	fmt.Printf("%v,%T\n", myCode, myCode)
	fmt.Printf("%v,%T\n", myAliasCode, myAliasCode)
	var age = 1
	if myCode == MyCode(age) {

	}
	if myAliasCode == age {
	}
}
