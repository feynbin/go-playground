package main

import (
	"encoding/json"
	"fmt"
)

func plus[T Number](n1, n2 T) T {
	return n1 + n2
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

//	func plusUint(n1, n2 uint) uint {
//		return n1 + n2
//	}

func myPrint[T int, K string | int](n1 T, n2 K) {

}

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func main() {
	// 泛型的支持，也就是类型参数
	//plus(1, 2)
	//var u1, u2 = uint(2), uint(3)
	//plusUint(u1, u2)
	//plus(int(u1), int(u2))

	//plus(u1, u2)
	// 泛型没有出来之前，一个是加上对应类型的函数，一个是对参数进行类型转换

	// 有了泛型,可以通过T占位符，然后通过标明可用类型使用
	type User struct {
		Name string `json:"name"`
	}
	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	// 泛型结构体
	//user := Response{
	//	Code: 0,
	//	Msg:  "ok",
	//	Data: User{
	//		Name: "feynbin",
	//	},
	//}
	//userInfo := Response{
	//	Code: 0,
	//	Msg:  "ok",
	//	Data: UserInfo{
	//		Name: "feynbin",
	//		Age:  24,
	//	},
	//}
	//byteData, _ := json.Marshal(user)
	//fmt.Println(string(byteData))
	//byteData, _ = json.Marshal(userInfo)
	//fmt.Println(string(byteData))

	var userResponse Response[User]
	json.Unmarshal([]byte(`{"code":0,"msg":"ok","data":{"name":"feynbin"}}`), &userResponse)
	fmt.Println(userResponse.Data.Name)
	var userInfoResponse Response[UserInfo]
	json.Unmarshal([]byte(`{"code":0,"msg":"ok","data":{"name":"feynbin","age":24}}`), &userInfoResponse)
	fmt.Println(userInfoResponse.Data.Name, userInfoResponse.Data.Age)

	// 泛型切片
	//type MySlice[T int| string] []T
	//var mySlice = MySlice[int]{1,2,3}
	//fmt.Println(mySlice[0]+1)

	// 泛型map,key只能是简单类型
	type MyMap[T string, K any] map[T]K
	var myMap = MyMap[string, int]{
		"name": 12,
	}
	fmt.Println(myMap)

}
