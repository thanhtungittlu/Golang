package main

import "fmt"

func main()  {
	r:= rectangle{
		length: 10,
		width: 5,
	}
	p:= &r // địa chỉ của r
	perimeter(p) // Hợp lệ vì p là địa chỉ của r
	//perimeter(r) // Không hợp lệ vì biến truyền vào k phải địa chỉ
	p.perimeter()
	r.perimeter() // Dùng giá trị r gọi đến phương thức perimeter() có vật nhận là con trỏ
	// r.perimeter() = (&r).perimeter()

}

type rectangle struct {
	length int
	width int
}

func perimeter(r *rectangle)  {
	fmt.Println("Perimeter function output: ", 2*(r.length + r.width))
}

func (r *rectangle) perimeter()  {
	fmt.Println("Perimeter method output: ",2*(r.length+r.width) )
}


