package main

import "fmt"

type Describe interface{
	Describe()
}

func main()  {
	var d1 Describe
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T, value %v \n", d1,d1)
	}
}