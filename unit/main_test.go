package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("测试前")
}

func teardown() {
	fmt.Println("测试后")
}
func TestAdd2(t *testing.T) {
	fmt.Println("测试中")
}

// 测试的时候的主函数
func TestMain(m *testing.M) {
	fmt.Println("Testmain")
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
