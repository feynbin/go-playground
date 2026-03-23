package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int) // 声明并初始化一个长度为0的信道

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)

	moneyChan <- money
	wait.Done()
}
func main() {
	// Goroutine 是Go运行时管理的轻量化线程
	// 在Go中，开启一个协程是非常简单的

	// 协程，异步函数或者叫耗时函数
	// 主线程结束，协程函数就跟着结束

	var wait sync.WaitGroup
	wait.Add(3)
	startTime := time.Now()

	go pay("zhangsan", 2, &wait)
	go pay("wangwu", 3, &wait)
	go pay("lisi", 5, &wait)
	go func() {
		wait.Wait()
		close(moneyChan)
	}()
	var moneyList []int
	for money := range moneyChan {
		moneyList = append(moneyList, money)
	}
	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println(money, ok)
	//	if !ok {
	//		break
	//	}
	//}

	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println("moneyList", moneyList)

	//
}
