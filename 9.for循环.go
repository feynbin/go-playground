package main

import "fmt"

func main() {
	// 传统循环
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 死循环，条件永远成立
	// for i := 1; true; i++ {...|
	// for {} 或者 for true {}

	// while 模式循环
	//var s int
	//var i int = 1
	//for i <= 100 {
	//	sum += i
	//	i++
	//}
	//fmt.Println(s)

	// do-while 模式 不管是否满足条件，都进行一次循环体

	//var su int
	//var id int = 1
	//for {
	//	su += id
	//	id++
	//	if id == 101 {
	//		break
	//	}
	//}
	//fmt.Println(su)

	// 遍历slice
	var list = []string{"feynbin", "feynin1"}
	for i := 0; i < len(list); i++ {
		fmt.Println(i, list[i])

	}
	for i, s := range list {
		fmt.Println(i, s)
	}

	// 遍历map
	var userMap = map[string]string{"name": "feynbin"}
	for key, value := range userMap {
		fmt.Println(key, value)
	}

	// break ， continue
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()

	}
}
