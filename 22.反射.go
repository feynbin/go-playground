package main

import (
	"fmt"
	"reflect"
	"strings"
)

func getType(obj any, value any) {
	v1 := reflect.ValueOf(obj)
	v2 := reflect.ValueOf(value)
	if v1.Elem().Kind() != v2.Kind() {
		return
	}
	switch v1.Elem().Kind() {
	case reflect.String:
		//fmt.Println("String", v1.String())
		v1.Elem().SetString(value.(string))
	case reflect.Int:
		//fmt.Println("Int", v1.Int())
		v1.Elem().SetInt(v2.Int())
	case reflect.Struct:
		fmt.Println("Struct")

	}

}

type Student1 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	IsMan bool
}

func ParseJson(obj any) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		jsonTag := tf.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = tf.Name
		}
		fmt.Printf("%s,%#v\n", tf.Name, jsonTag)
		fmt.Println(v.Field(i))

	}

}

type User1 struct {
	Name1 string `big:"-"`
	Name2 string
}

func SetStruct(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		big := t.Field(i).Tag.Get("big")
		if big == "" {
			continue
		}
		value.SetString(strings.ToUpper(value.String()))
	}
}

type Class1 struct {
	Name string `feng-orm:"name"`
	Id   int    `feng-orm:"id"`
}

//	func Find(obj any, query ...any) (sql string, err error) {
//		// obj 必须得是结构体
//		t := reflect.TypeOf(obj)
//		if t.Kind() != reflect.Struct {
//			err = errors.New("obj必须得是结构体")
//			return
//		}
//		var where string
//		// 验证条件
//		if len(query) > 0 {
//			q := query[0]
//			// 验证第一个参数的类型，必须是string
//			qs, ok := q.(string)
//			if !ok {
//				err = errors.New("查询的第一个参数必须是字符串")
//				return
//			}
//			// 算？的个数
//			if strings.Count(qs, "?")+1 != len(query) {
//				err = errors.New("查询参数个数不匹配")
//				return
//			}
//			// 拼接where
//			for i, a := range query[1:] {
//				switch s := a.(type) {
//				case string:
//					strings.Replace(qs, "?", fmt.Sprintf("%s?", s), 1)
//				case int:
//					qs = strings.Replace(qs, "?", fmt.Sprintf("%d?", s), 1)
//				}
//				strings.Replace(qs, "?")
//
//			}
//		}
//		return
//	}
func main() {
	// 类型判断，判断一个变量是否是结构体，切片，map
	//getType("23")
	//getType(123)
	//getType(struct {
	//	Name string
	//}{Name: "John"})
	// 获取值，valueof

	// 修改值，
	//var name = "feynbin"
	//var age = 24
	//getType(&name, "zhangsan")
	//fmt.Println(name)
	//getType(&age, 25)
	//fmt.Println(age)

	// 结构体反射,获取结构体的值，修改结构体的值
	//s := Student1{Name: "feynbin", Age: 24, IsMan: true}
	//ParseJson(s)
	//s := User1{Name1: "name1", Name2: "name2"}

	//SetStruct(&s)
	//fmt.Println(s)

	// 调用结构体方法。

	//Find(Class1{}, "name = ?", "三年一班")
	// select name, id from class_models where name = '三年一班';

}
