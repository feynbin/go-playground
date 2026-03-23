package main

import (
	"fmt"
	"sync"
)

// var sum int
// var wait2 sync.WaitGroup
//
//	func addd() {
//		for i := 0; i < 100000; i++ {
//			sum++
//		}
//		wait2.Done()
//	}
//
//	func sub() {
//		for i := 0; i < 100000; i++ {
//			sum--
//		}
//		wait2.Done()
//	}
var sum int
var lock sync.Mutex
var wait1 sync.WaitGroup

func addd() {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		sum++
	}
	lock.Unlock()
	wait1.Done()
}
func sub() {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		sum--
	}
	lock.Unlock()
	wait1.Done()
}
func main() {
	//wait2.Add(2)
	//go addd()
	//go sub()
	//wait2.Wait()
	//fmt.Println(sum)

	// 线程安全，也就是并发问题，
	wait1.Add(2)
	go addd()
	go sub()
	wait1.Wait()
	fmt.Println(sum)
	
}
