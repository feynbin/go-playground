package main

import (
	"fmt"
	"sort"
)

func main() {
	var nameList []string
	var nameList1 = []string{}
	nameList2 := []string{}
	var nameList3 = []string{}
	nameList3 = make([]string, 0)
	nameList = append(nameList, "feynbin")
	nameList = append(nameList, "feynbin1")
	// 只能同类型。
	fmt.Println(nameList[0])
	nameList1 = append(nameList1, "feynbin")
	fmt.Println(nameList1[0])
	nameList2 = append(nameList2, "feynbin")
	nameList3 = append(nameList3, "feynbin")

	// 除了基本数据类型，其他数据类型如果只定义不赋值，那么实际的值就是nil
	var list = make([]string, 0)
	fmt.Println(list, len(list), cap(list))
	fmt.Println(list == nil)
	ageList := make([]int, 3)
	fmt.Println(ageList)

	array := [3]int{1, 2, 3}
	slices := array[:]
	fmt.Println(slices)
	fmt.Println(array[0:2])

	var ints = []int{4, 8, 3, 2}
	sort.Ints(ints)
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints))) // 降序
	fmt.Println(ints)
}
