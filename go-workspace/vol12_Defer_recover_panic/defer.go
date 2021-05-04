package main

import "fmt"

func main() {
	fmt.Println("Line 1")
	defer fmt.Println("Line 2") // Câu lệnh cuối cùng. chương trình chạy hết lệnh mới chạy lệnh sau defer
	fmt.Println("Line 3")
	fmt.Println("Line 4")
	defer fmt.Println("Line 5") //defer sẽ để vào 1 ngăn xếp, sau khi xuất ra sẽ xuất từ trên ra
	fmt.Println("ví dụ 2")
	a:=12
	defer fmt.Println(a) // Defer sẽ lưu giá trị a = 12 trước.
	a = 24
	fmt.Println("Ví dụ 3")
	c:= 10
	d:= 3
	//panic("Chia cho một số là 0")
	fmt.Println(c/d)
	fmt.Println()
	defer fmt.Println("----Các lệnh sau defer")
}