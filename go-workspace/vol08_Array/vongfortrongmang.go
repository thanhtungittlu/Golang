package main

import "fmt"

func main()  {
	a := [...]int{2,4,6,8,10}
	sum := 0
	for i, v := range a{
		fmt.Printf("%d the element of a is %d\n",i,v)
		sum += v
	}	
	fmt.Println(sum)
}
// i là index, v là giá trị