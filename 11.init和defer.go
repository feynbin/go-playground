package main

import "fmt"

var db int

func init() {
	db = 10
	fmt.Println("连接数据库成功")
	fmt.Println("init-1")

}
func init() {
	fmt.Println("init-2")

}
func init() {
	fmt.Println("init-3")
}
func main() {
	// init是一个特殊的函数
	// 1. 不能被其他函数调用，在main函数之前执行，自动调用，如果有多个init，那么按照顺序执行
	// 2. init 函数不能作为参数传入
	// 3. init 函数不能有传入参数和返回值

	// 一般用作初始化

	//fmt.Println("main", db)

	// defer函数，又叫延迟函数
	// 1. 关键字defer用于注册延迟调用
	// 2. defer直到return前才被执行，因此，可以用来做资源清理
	// 3. 多个defer语句，按先进后出的方式执行，谁离return近谁先执行
	// 4. defer语句中的变量，在defer声明时就决定了。也就是只能获得defer之前的变量
	defer func() {
		fmt.Println("defer-2")
	}()
	defer func() {

		fmt.Println("defer-1")
	}()
	//var name = "feynbin" // defer获取不到
	return
}
