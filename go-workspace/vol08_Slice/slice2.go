package main

import "fmt"

func main()  {
	a := []int{1,2,3,4,5,6}
	copy := a[2:5] // lấy giá trị index từ 2 đến 4
	fmt.Println("after", a)
	for i := range copy { // duyệt mình index i
		copy[i]+=10 //cho giá trị + thêm 10
	}
	fmt.Println("before", a)
}
// Slice là 1 con trỏ, chứ không tạo ra bản copy
// CÒn mảng thì ngược lại, nó tạo ra 1 bản copy