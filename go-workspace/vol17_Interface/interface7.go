package main

import "fmt"

type Describe interface {
	Describe()
}

type Person struct {
	name string
	age int
}

func (p Person) Describe(){
	fmt.Printf("%s is %d year old.", p.name, p.age)
}

func findType (i interface{}){
	switch v := i.(type) {
	case Describe:
		v.Describe()
	default:
		fmt.Printf("Unknown type \n")	
	}
}

func main()  {
	findType("Tung")
	p := Person {
		name: "Tung R",
		age: 25,
	}
	findType(p)
}
