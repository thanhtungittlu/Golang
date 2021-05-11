package main

import (
	"fmt"
	"log"
)

func firstChan(nums ...int) chan int  {
	result := make(chan int)

	go func() {
		for i := 0; i < len(nums); i++ {
			result <- nums[i]
		}
		close(result)
	}()

	return result
}

func secondChan(firstInput chan int) chan int  {
	result := make(chan int)
	go func() {
		for item:= range firstInput {
			result <- item *item
		}
		close(result)
	}()
	return result
}

func main()  {
	firstChan := firstChan(1,2,3,4,5)
	secondChan:= secondChan(firstChan)

	for item:= range secondChan {
		log.Println(item)
	}

	fmt.Println("Done")
}