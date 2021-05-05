package main

import "fmt"

func describe( i interface{}) {
	fmt.Printf("Type = %T, Value = %v \n", i,i)
}


func main()  {
	s:= "Hello My Friend"
	describe(s)
	i:= 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Tung",
	}
	describe(strt)
}

// describe (i interface{}) lấy một interface rỗng làm đối số do đó nó có thể truyền vào bất kỳ kiểu dữ liệu nào.