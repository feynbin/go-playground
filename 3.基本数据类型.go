package main

import "fmt"

func main() {
	// var age = 12
	// var u8 uint8 = 255
	// uint8
	// 0 0 0 0 0 0 0 0
	// 1 1 1 1 1 1 1 1 = 255 = 2^8-1
	// int8
	// 0    0 0 0 0 0 0 0 = 0
	// 0    1 1 1 1 1 1 1 = 127
	// 1    0 0 0 0 0 0 1 = -1 (原码) 补码 (取反=反码+1)
	// 1    0 0 0 0 0 0 0 = -128
	var a byte = 'a'
	fmt.Printf("%c %d\n ", a, a)
	var a1 uint8 = 97
	fmt.Printf("%c %d\n ", a1, a1)
	var z rune = '中'
	fmt.Printf("%c %d\n ", z, z)
}
