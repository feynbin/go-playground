package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	res = a / b
	return
}

func server() (res int, err error) {
	res, err = div(2, 0)
	if err != nil {
		// 把错误向上抛
		return
	}
	// 执行其他的逻辑
	return
}

//func init() {
//	_, err := os.ReadFile("12344")
//	if err != nil {
//		//log.Fatalln("错误")
//		panic("错误了")
//	}
//
//}

func read() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
		}
	}()
	var list []int = []int{1, 2}
	fmt.Println(list[2])
}
func main1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
		}
	}()

	read()
}

func main() {
	// go的异常处理是go唯一的一个诟病
	// 由于go语言没有捕获异常的机制，导致每调一个函数都要接一下这个函数的error
	// 网上有一个梗，叫做error是go的一等共鸣

	// 常用的异常处理
	// 1.向上抛，一般是用于框架层
	// 2.中断程序，一般是用于初始化
	// 3.恢复程序，

	// 1.向上抛
	//res, err := server()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(res)

	// 2.中断程序
	// 除了初始化一般不要使用中断程序，可以使用logs。Fatalln来打印错误，也可以直接提供一个panic

	//3.恢复程序，也就是异常捕获
	main1()
	fmt.Println("正常逻辑")
}
