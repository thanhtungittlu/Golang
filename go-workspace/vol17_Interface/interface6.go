package main

import "fmt"

func findType (i interface{}){
	switch i.(type) {
	case string:
		fmt.Printf("I am a %T and my value is %s \n",i, i.(string))
	case int:
		fmt.Printf("i am an %T and my value is %d\n",i, i.(int))
	default:
		fmt.Printf("I don't know \n")
		
	}
}

func main()  {
	findType("Tung")
	findType(50)
	findType(50.2)
}