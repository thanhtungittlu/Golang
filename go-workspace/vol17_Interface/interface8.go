package main

import "fmt"

type Describe interface {
	Describe()
}

type Person struct {
	name string
	age int		
}
func (p Person) Describe()  {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct{
	state string
	country string
}

func (a *Address) Describe()  {
	fmt.Printf("State %s, Country %s", a.state, a.country)
}

func main()  {
	var d1 Describe
	p1 := Person{"Tung", 22}
	d1 = p1
	d1.Describe()
	p2 := Person{"Thuong", 23}
	d1 = &p2
	d1.Describe()

	var d2 Describe

	a:= Address {"Ha Noi", "Viet Nam"}

	d2 = &a // K thể dùng d2 = a

	d2.Describe()
	
}
