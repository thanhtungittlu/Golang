package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId int
	basicpay int
	pf int	
}

type  Contract struct {
	empID int
	basicpay int
}

// Phương thức tính lương của nhân viên Permanent
func (p Permanent) CalculateSalary() int  {
	return p.basicpay + p.pf
}

//Phương thức tính lương của nhân viên contract chỉ là basic pay

func (c Contract) CalculateSalary() int  {
	return c.basicpay
}

// Tổng chi phí tính bằng cách tính tổng mức lương của nhân viên, và duyệt qua slice []SalaryCalculator

func totalExpense(s []SalaryCalculator)  {
	expense :=0 
	for _, v:= range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Println("Total Expense Per Month: ", expense)
}

func main()  {
	p1 := Permanent{1,5000,20}
	p2 := Permanent{2,5000,20}
	p3 := Permanent{3,5000,20}
	p4 := Contract{4,2000}
	p5 := Contract{5,3000}
	employees := []SalaryCalculator{p1,p2,p3,p4,p5}
	totalExpense(employees)
}
