package main

import "fmt"

func printSth(msg string) chan string {
	result := make(chan string)
	go func ()  {
		for i := 0; i <= 5; i++ {
			result <- fmt.Sprintf("%s %d",msg,i)
		}
	}()
	return result

}

func main()  {
	a := printSth("Tung")
	for i := 0; i <= 5; i++ {
		fmt.Println(<- a)
	}
}