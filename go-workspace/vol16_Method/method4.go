package main

import "fmt"

type rectangle struct {
	length int
	width int
}
func area(r rectangle){
	fmt.Printf("Area Function result: %d \n", (r.length*r.width))
}

func (r rectangle) area(){
	fmt.Printf("Area Method result: %d\n",(r.length*r.width))
}

func main()  {
	r := rectangle{
		length: 10,
		width:5,
	}
	area(r)
	r.area()
	
	p:= &r // Khởi tạo con trỏ tới r

	// area(p) : biên dịch sẽ báo lỗi vì khi này p là con trỏ

	p.area() // nhưng phương thức này thì sẽ k bị báo lỗi vì Go sẽ ngầm hiểu là (*p).area()
	
}