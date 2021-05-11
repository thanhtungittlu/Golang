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

func fanIn(a,b chan string) chan string  {
	c := make(chan string)
	go func (){
		for {
			select{
			case <- a:
				c <- <-a
			case <- b:
				c <- <-b
			}
		}
	}()
	return c	
}

func main()  {
	a := printSth("Tung")
	b := printSth("Thuong")
	c := fanIn(a,b)
	for i := 0; i <= 5; i++ {
		fmt.Println(<- c)
	}
}