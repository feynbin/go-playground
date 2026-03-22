package main

import "fmt"

func main() {
	// 判断语句就是根据不同的条件去执行不同的逻辑
	var age int
	fmt.Printf("请输入你的年龄: ")
	//fmt.Scan(&age)
	// 中断式 卫语句
	//if age <= 0 {
	//	fmt.Println("未出生")
	//	return
	//}
	//if age <= 18 {
	//	fmt.Printf("未成年")
	//	return
	//}
	//if age <=35 {
	//	fmt.Println("青年")
	//	return
	//}
	//fmt.Println("中年")

	// 非常推荐中断式

	// 嵌套式，不推荐
	//if age <= 18 {
	//	if age <= 0 {
	//		fmt.Println("未出生")
	//	} else {
	//		fmt.Println("未成年")
	//	}
	//} else {
	//}
	//if age <= 35 {
	//	fmt.Println("青年")
	//} else {
	//	fmt.Println("中年")
	//}

	// 多条件式
	if age <= 0 {
		fmt.Println("未出生")
	}
	if age <= 18 && age > 0 {
		fmt.Println("未成年")
	}
	if age <= 35 && age > 18 {
		fmt.Println("青年")
	}
	if age > 35 {
		fmt.Println("中年")
	}

	// && 与，两个条件同时满足才得到true
	// || 或，两个条件满足之一得到true
	// ! 取反
	fmt.Println(1 == 1 && 1 == 1, 1 == 1 && 1 == 0, 1 == 0 && 1 == 0)
	fmt.Println(1 == 1 || 1 == 1, 1 == 1 || 1 == 0, 1 == 0 || 1 == 0)
	fmt.Println(!(1 == 1), !(1 == 0))
	// 逻辑短路
	// && 第一个条件如果是false，后面的条件就不会走了
	// || 第一个条件如果是true，后面的条件就不会走了
	
}
