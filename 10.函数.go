package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hello")

}

func param1(id, userName string) {

	fmt.Println(id)

}

func add(numberList ...int) {
	var sum int
	for _, i := range numberList {
		sum += i
	}
	fmt.Println(sum)
}

func r1() {
	// 没有返回值
	return
}

func r2() bool {
	return false
}

func r3() (string, bool) {
	return "", true
}

// 命名式函数返回
func r4() (val string, ok bool) {
	if 1 > 2 {
		val = "12"
		return val, ok
	}
	return
}

func Copy(name string) {
	fmt.Printf("infun %p\n", &name)
	name = "feynbin不知道"
}

func Set(name *string) {
	fmt.Printf("infun %p\n", name)
	*name = "feynbin不知道"
	// 通过内存地址去找值
}
func main() {
	// 函数是一段封装了特定功能的可重用代码块，用于执行特定的任务或计算
	// 函数接受输入参数并产生输出，也就是返回值。

	// 定义，使用func关键字来定义一个函数

	//sayHello()
	//param1("123", "feynbin")
	//add(1, 2)
	//add(1, 2, 3)
	//add(1, 2, 3, 4)

	//返回值

	// 匿名函数, 不支持在函数中再写一个函数, 需要在函数体中创建函数的时候使用匿名函数

	//var getName = func() string { //相当于创建了一个函数
	//	return "feynbin"
	//}
	//var setName = func(name string) {
	//	fmt.Println(name)
	//	return
	//}
	//fmt.Println(getName())
	//setName("zhangsan")
	//
	//// 高阶函数,就是把函数当作参数传到另一个函数中.
	//fmt.Println("请输入你要执行的操作: ")
	//fmt.Println("1 登录: ")
	//fmt.Println("2 注册: ")
	//fmt.Println("3 个人中心: ")
	//var index int
	//fmt.Scanln(&index)
	//
	//var funMap = map[int]func(){
	//	1: login,
	//	2: register,
	//	3: userCenter,
	//}
	//fun, ok := funMap[index]
	//if ok {
	//	fun()
	//}
	//switch index {
	//case 1:
	//
	//	login()
	//case 2:
	//	registrer()
	//case 3:
	//	userCenter()
	//}

	// 闭包
	// 设计一个函数，先传一个参数表示延时，后面再次传参数就是将参数求和
	//
	//t1 := time.Now()
	//
	//sum := awaitAdd(2)(1, 2, 3)
	//subTime := time.Since(t1)
	//fmt.Println(sum)
	//fmt.Println(sum, subTime)
	// 以此类推，函数可以再嵌套嵌套

	// 值传递和引用传递
	var name = "feynbin"
	fmt.Printf("%p\n", &name)
	Set(&name)
	fmt.Println(name)
}

func login() {
	fmt.Println("登录")
}
func register() {
	fmt.Println("注册")
}
func userCenter() {
	fmt.Println("个人中心")
}

func awaitAdd(awitSecond int) func(...int) int {
	time.Sleep(time.Duration(awitSecond) * time.Second)
	return func(numberList ...int) (sum int) {

		for _, i2 := range numberList {
			sum += i2
		}
		return sum
	}

}
