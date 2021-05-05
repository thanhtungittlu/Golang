package main

import "fmt"

type SalaryCalculator interface{
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface{
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary()  {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
	
}

func (e Employee) CalculateLeavesLeft() int  {
	return e.totalLeaves - e.leavesTaken
}

func main()  {
	e:= Employee{"Tung","Le Thanh", 1000,200,30,5}
	var e1 EmployeeOperations = e
	e1.DisplaySalary()
	fmt.Println("\n Ket qua = ", e1.CalculateLeavesLeft())
}