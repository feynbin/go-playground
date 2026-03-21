package main

import "fmt"

func main() {
	// 学号-名称
	var userMap map[int]string = map[int]string{
		1: "feynbin",
		2: "张三",
		4: "",
	}
	fmt.Println(userMap)
	// 取值
	fmt.Println(userMap[1])
	fmt.Printf("%#v\n", userMap[3])
	value, ok := userMap[3]
	// 两个值，第一个值是变量，第二个值是布尔值.
	fmt.Println(value, ok)

	// 设置值
	userMap[1] = "哈夫克"
	fmt.Println(userMap)

	// 删除值
	delete(userMap, 1)
	fmt.Println(userMap)

	// map创建的时候一定要赋值
	//var aMap map[string]string
	//var aMap = make(map[string]string)
	var aMap = map[string]string{}
	//aMap["name"] = "阿萨拉"
	fmt.Println(aMap)
}
