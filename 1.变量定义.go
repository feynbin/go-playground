package main

import "fmt"

// 常量赋值后就不能修改了
// 跨包使用，首字母大写
const Version string = "2.0.1"

func main() {
	// 先声明
	var name string
	name = "feynbin"
	fmt.Println(name)

	// 声明并复制
	var name1 string = "feynbin1"
	// 省略类型
	var name2 = "feynbin2"

	// 声明并赋值 短声明符号
	name3 := "feynbin3"
	fmt.Println(name1, name2, name3)

}
