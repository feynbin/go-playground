package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan1 = make(chan int)
var nameChan1 = make(chan string)
var doneChan = make(chan struct{})
var done = make(chan struct{})

func send(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)
	moneyChan1 <- money
	nameChan1 <- name

	wait.Done()
}

func main() {
	//
	var wait sync.WaitGroup
	wait.Add(3)
	startTime := time.Now()

	go send("zhangsan", 2, &wait)
	go send("wangwu", 3, &wait)
	go send("lisi", 5, &wait)
	go func() {
		defer close(moneyChan1)
		defer close(nameChan1)
		defer close(doneChan)
		wait.Wait()

	}()
	var moneyList []int
	var nameList []string

	var event = func() {
		for {
			select {
			case money := <-moneyChan1:
				moneyList = append(moneyList, money)
			case name := <-nameChan1:
				nameList = append(nameList, name)
			case <-doneChan:

				return
			}

		}
	}
	event()
	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println("moneyList", moneyList)
	fmt.Println("nameList", nameList)
	//fmt.Println("购买完成", time.Since(startTime))
	//fmt.Println("moneyList", moneyList)
	//fmt.Println("nameList", nameList)
	go chaoshi()
	select {
	case <-done:
		fmt.Println("协程执行完毕")
	case <-time.After(3 * time.Second):
		fmt.Println("超时")
		return
	}

}

func chaoshi() {
	fmt.Println("协程执行开始")
	time.Sleep(2 * time.Second)
	fmt.Println("协程执行结束")
	close(done)
}
