package main

import "fmt"

func main()  {
	emp1 := Employee {
		name : "Pham Van hung",
		age: 22,
	}
	fmt.Printf("Employee name before change: %s\n", emp1.name)
	emp1.changeName("Le Thanh Tung")// Tạo 1 bản sao, nên tên vẫn không thay đổi
	fmt.Printf("\nEmployee name after change: %s\n", emp1.name) 
	emp1copy := emp1.changeName("Le Thanh Tung") // Bản sao đấy là e, nên muốn xuất ra tên sau thì phải e.name
	fmt.Println("alo: ", emp1copy.name)
	fmt.Println("--------Age---------")
	fmt.Printf("Employee age before change: %d\n", emp1.age)
	emp1.changeAge(23)
	fmt.Printf("\nEmployee age after change: %d\n", emp1.age) 


}

type Employee struct {
	name string
	age int
}

func (e Employee) changeName(newName string) Employee {  //Phương thức trỏ tới giá trị
	e.name = newName
	return e
}

func (e *Employee) changeAge(newAge int)  {  //Phương thức trỏ tới địa chỉ
	e.age = newAge
}


