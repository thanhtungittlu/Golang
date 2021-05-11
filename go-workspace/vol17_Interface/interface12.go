package main

import "fmt"

type Animal interface{
	speak()
	eat()
}

type Movement interface{
	move()
}

type Dog struct{}
type Cat struct{}

func (d Dog) speak()  { //implement tới interface Animal
	fmt.Println("Go Go")
}
func (d Dog) eat() {
	fmt.Println("Eat Dog")
}
func (d Dog) move()  { //implement tới interface Movement. 1 đối tượng có thể implement đến nhiều interface
	fmt.Println("Go Go")
}

func (d Dog) foot()  { //Không sự dụng interface
	fmt.Println("Four")
}
func (c Cat) speak()  {
	fmt.Println("Meow meow")
}
func (c Cat) eat()  {
	fmt.Println("Eat Cat")
}

func main()  {
	
	dog := Dog{}

	var m Movement = dog
	var a Animal = dog
	var d Dog // Vì không dùng interface nên phải khai báo thêm biên d là kiểu Dog
	var cat Animal
	
	cat = Cat{}
	
	d.foot()
	a.speak() ; a.eat() ; m.move()
	cat.speak() ; cat.eat()
}