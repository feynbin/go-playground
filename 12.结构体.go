package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	Name string
}
type Student struct {
	Class
	Name string
}

func (s Student) Study() {
	fmt.Printf("%s\n", s.Name)
}
func (s Student) Info() {
	fmt.Printf("名字: %s 班级：%s\n", s.Name, s.Class.Name)
}
func (s *Student) SetName(name string) {
	s.Name = name
	fmt.Printf("method in : %p\n", s)

}

type UserInfo struct {
	Name string `json:"name"`
	Age  int
}

func (this *UserInfo) SetName(name string) {
	this.Name = name

	fmt.Printf("%p\n", this)
}

type User struct {
	Name     string `json:"zhangsan"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"-"`
}

func main() {
	// 结构体，又叫对象

	c1 := Class{Name: "三年一班"}
	s1 := Student{Class: c1, Name: "feynbin"}
	s1.Study()
	s1.Info()
	s2 := Student{Class: c1, Name: "zhangsan"}
	s2.Study()
	s2.Info()
	// 继承，又叫组合，就是把一个结构体提供给另一个结构体使用
	// 结构体指针
	fmt.Printf("method out : %p\n", &s2)
	s2.SetName("李四")
	s2.Info()

	// 结构体tag
	user := UserInfo{
		Name: "feynbin",
	}
	fmt.Printf("%p\n", &user)
	user.SetName("zhangsan")
	fmt.Println(user.Name)
	//
	user1 := User{Name: "feynbin", Age: 1, Password: "123456"}
	byteData, _ := json.Marshal(user1)
	fmt.Printf(string(byteData))
}
