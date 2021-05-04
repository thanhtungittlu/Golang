package main

import "fmt"

func main()  {
	emp1 := Employee {
		name: "Tung",
		salary : 1000,
		currency: "$",
	}
	emp1.displaySalary()
}

type Employee struct {
	name string
	salary int
	currency string
}

func (e Employee) displaySalary(){
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}
// Phương thức displaySalary() có vật nhật(kiểu dữ liệu đặc biệt)là Employee

