package main

import (
	"fmt"
	"time"
	"sync"
)

func main()  {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func ()  {
		count("sheep")
		wg.Done()
	}()
	go func ()  {
		count("fish")
		wg.Done()
	}()
	wg.Wait()
	
	fmt.Println("Done")
}

func count(name string)  {
	for i := 1; i <= 5; i++ {
		fmt.Println(name,i)
		time.Sleep(time.Second)
	}
}