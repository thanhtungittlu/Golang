package main

import (
	"fmt"
	"sync"
)

func printNumber(wg *sync.WaitGroup)  {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d ",i)	
	}
	wg.Done()
}

func printChar(wg *sync.WaitGroup)  {
	for i := 'A'; i < 'A' + 26 ; i++ {
		fmt.Printf("%c ",i)	
	}
	wg.Done()
}

func main()  {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go printNumber(&wg)
	go printChar(&wg)
	wg.Wait()
	fmt.Println("Done")
}

//B1: tạo 1 biến có kiểu sync.WaitGroup: wg
//B2: wg.Add(Số lượng gorotine)
//B3: wg.Done() trước khi return, 
//B4: Gọi method wg.Wait()