package main

import "fmt"

func main() {
	var nameList [3]string = [3]string{"a", "b", "c"}
	fmt.Println(nameList)
	// a b c
	// 0 1 2
	// -3  -2  -1
	fmt.Println(nameList[0])
	fmt.Println(nameList[2])
	fmt.Println(nameList[len(nameList)-1])
	// nameList[-1] go不支持

}
