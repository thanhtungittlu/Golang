package main

import "fmt"

func printarray( a [3][2]int)  {
	for _,v1 := range a{
		for _,v2 := range v1{
			fmt.Printf("%d ",v2)
		}
		fmt.Println()
	}
}

func main()  {
	a:= [3][2]int{
		{1,2},
		{4,5},
		{9,10},
	}
	printarray(a)
}