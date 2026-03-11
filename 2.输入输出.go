package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println("你好")
	// 格式化输出
	fmt.Printf("%s,哇，你好美\n", "feynbin")
	// %s字符串，%d整数，%f小数，%T类型，%v任意，%#v %c
	var f = fmt.Sprintf("%.2f", 3.1415926)
	// 将输出赋值
	fmt.Println(f)

	fmt.Print("请输入你的名字: ")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)
	fmt.Print("请输入你的年龄: ")
	var age int
	n, err := fmt.Scan(&age)
	fmt.Println(n, err, age)
}
