package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg = sync.WaitGroup{}
	ch := make (chan int, 50)
	wg.Add(2)
	go func (ch <-chan int)  { // Chỉ nhận thôitừ channel
		for i:= range ch{
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // Chỉ gửi vào channel
		ch <- 17
		ch <- 42
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}