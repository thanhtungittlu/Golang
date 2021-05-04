package main

import "fmt"

func main()  {
	emp1 := Employee {
		name : "Pham Van hung",
		age: 22,
	}
	fmt.Printf("Employee name before change: %s\n", emp1.name)
	e := emp1.changeName("Le Thanh Tung")
	fmt.Println("alo: ", e.name)
	fmt.Printf("\nEmployee name after change: %s\n", emp1.name)


}

type Employee struct {
	name string
	age int
}

func (e Employee) changeName(newName string) Employee{ 
	e.name = newName
	fmt.Println("e: ", e.name)
	return e
}
