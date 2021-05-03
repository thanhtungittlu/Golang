package main

import (
	"fmt"
)

type People struct {
	name   string
	isMale bool
}

type Student struct { // Student, S viet hoa để mọi packet đều dùng được
	number int `Maximum can't over 100`
	People
	subjects []string
}

// Có thể khai báo 1 struct trong 1 struct
// tag `Noi Dung`
func main() {
	/*	studentTung := Student{ // Có thể khao báo tuần tự, không cần tên thuộc tính
			number: 3,
			name:   "Tung",
			isMale: true,
			subjects: []string{
				"Math",
				"English",
			},
		}
		fmt.Println(studentTung)
		fmt.Println(studentTung.name) */
	studentThuong := Student{}
	studentThuong.number = 4
	studentThuong.name = "Thuong"
	studentThuong.isMale = false
	studentThuong.subjects = []string{"Math", "English", "Computer"}
	fmt.Println(studentThuong)
	fmt.Println()
	// Cách 2 khai báo Struct :
	student2 := struct{ name string }{name: "Tung"}
	copy := student2 // Copy nhưng k truy xuất đến cùng 1 giá trị, nên 2 giá trị sẽ k liên quan tới nhau
	fmt.Println(student2)
	copy.name = "Thuong"
	fmt.Println(copy)

}
