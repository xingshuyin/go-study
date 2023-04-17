package main

import (
	"fmt"
	"time"
)

// 需求：
// 要求统计1-200000的数字，哪些是素数？ 这个问题在本章开篇就提出了，现在我们有goroutine和channel的知识后，就可以完成了【测试数据：80000】

// 分析思路：
// 传统的方法，就是使用一个循环，循环的判断各个数是不是素数【ok】
// 使用并发/并行的方式，将统计素数的任务分配给多个（x个）goroutine去完成，完成任务时间短

// 1. 向intChan放入1-8000的数字
func putNum(intChan chan int) {
	for i := 1; i <= 20000; i++ {
		intChan <- i
	}
	// 关闭intChan
	close(intChan)
}

// 从intChan中取出数据，并判断是否为素数，如果是，就放入到primeChan
func primeNum1(intChan chan int, primeChan chan int, exitChan chan bool) {
	// 使用for循环
	var flag bool
	for {
		time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok { //intChan取不到，且关闭了管道
			break
		}
		flag = true // 假设是素数
		// 判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			// 将这个数放到primeChan
			primeChan <- num
		}
	}
	fmt.Println("有一个prieNum 协程因为取不到数据，退出")
	// 这里我们还不能关闭primeChan

	// 向exitChan写入true
	exitChan <- true
}

func main() {
	var intChan chan int = make(chan int, 1000)
	var primeChan chan int = make(chan int, 2000)
	// 标识退出的管道
	exitChan := make(chan bool, 4)

	// 开启一个协程，向intChan放入 1-8000个数
	go putNum(intChan)

	// 开启4个协程，从 intChan中取出数据，并判断是否为素数，如果是就放入到primeChan
	for i := 0; i < 10000; i++ {
		go primeNum1(intChan, primeChan, exitChan)
	}

	// 这里我们主线程，进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		// 当我们从exitChan中取出了4个结果，就可以放心的关闭primeChan
		close(primeChan)
	}()

	// 遍历我们的primeChan，把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		// 将结果取出
		fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main线程退出")
}
