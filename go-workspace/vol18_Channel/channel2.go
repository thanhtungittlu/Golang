package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg = sync.WaitGroup{}
	ch := make(chan string)
	wg.Add(2)
	go func ()  {
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}()
	go func ()  {
		ch <- "Le Thanh Tung"
		wg.Done()
	}()
	wg.Wait()
}
