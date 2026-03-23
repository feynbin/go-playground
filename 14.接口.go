package main

import "fmt"

type SingInterface interface {
	Sing()
}
type Chicken struct {
	Name string
}
type Cat struct {
	Name string
}

func (c Chicken) Sing() {
	fmt.Println(c.Name, "在唱歌")
}
func (c Cat) Sing() {
	fmt.Println(c.Name, "在唱歌")
}
func sing(c SingInterface) {
	c.Sing()
	ch, ok := c.(Chicken) // 返回一个类型
	fmt.Println(ch, ok)

	switch server := c.(type) {
	case Chicken:
		fmt.Println(server)
	case Cat:
		fmt.Println(server)
	default:
		fmt.Println("其他")
	}
}

func print(val any) {
	fmt.Println(val)

}
func main() {
	// 接口是一组仅包含方法名、参数、返回值的为未具体实现的方法的集合。
	ch := Chicken{Name: "ik"}
	ca := Chicken{Name: "咪咪"}
	sing(ch)
	sing(ca)
	// 接口本身不能绑定方法，接口是值传递。
	// 类型断言，指定一种类型断言，或者枚举所有类型

	// 空接口，空接口可以接受任何类型，换一种说法，任何类型都实现了空接口
	print("int")
	print(123)
}
