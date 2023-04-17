package main

import (
	"fmt"
	"sync"
	"time"
	// "time"
)

var group sync.WaitGroup
var judge_group sync.WaitGroup

func su_num(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func add_num(ch chan int) {
	for i := 1; i < 1000000; i++ {
		ch <- i
	}
	close(ch)
	fmt.Println("添加结束")
	group.Done()
}
func judge_num(ch chan int, ch_su chan int) {
	for v := range ch {
		// fmt.Println("judge_num", v)
		// time.Sleep(time.Second)
		// runtime.Gosched()
		if su_num(v) {
			ch_su <- v
		}
	}
	fmt.Println("判断结束")
	judge_group.Done()
}
func get_num(ch_su chan int) {
	for v := range ch_su {
		fmt.Println(v)
	}
	group.Done()
}
func main() {
	start := time.Now().Unix()
	group.Add(3)
	judge_group.Add(5)
	var ch = make(chan int, 100)
	var ch_su = make(chan int, 100)
	go add_num(ch)
	go judge_num(ch, ch_su)
	go judge_num(ch, ch_su)
	go judge_num(ch, ch_su)
	go judge_num(ch, ch_su)
	go judge_num(ch, ch_su)
	// go judge_num(ch, ch_su)
	// go judge_num(ch, ch_su)
	// go judge_num(ch, ch_su)
	// go judge_num(ch, ch_su)
	// go judge_num(ch, ch_su)
	go get_num(ch_su)
	judge_group.Wait()
	close(ch_su)
	group.Done()
	group.Wait()
	end := time.Now().Unix()
	fmt.Println("执行时间", end-start)

}
