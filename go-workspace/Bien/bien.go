package main

import "fmt"

var global int = 50 // Biến toàn cục, chỉ được khai báo theo cách 2
var (
	m int    = 10
	n int    = 20
	h string = "abc"
)

// lúc chiều em mới làm mình cái hello.go thôi cũng k dc

var Number int = 10 // Viết hoa chữ đầu của biến thì đều sử dụng được cho các package

func main() {
	var i int //Cach1: Khai báo biến
	i = 3000
	var j int = 10       // Cach2
	k := "Le Thanh Tung" // Cach3, go lang tự nhận biết kiểu dữ liệu
	fmt.Println(i)
	fmt.Println(j)
	fmt.Printf("%v, %T", k, k) // %v: giá trị của biên
	// %T: Kiểu giữ liệu của biến
	fmt.Println()
	fmt.Println(global, m, n, h)
	fmt.Println()
	var so int = 42
	var xau string = string(so)
	fmt.Println(xau)
}

/* Cách đặt tên biến: Camel
// 	 - Chữ cái đầu tiên của từ đầu tiên sẽ không biết hoa.
// 	 - Chữ cái đầu tiên của các từ tiếp theo sẽ viết hoa.
// 	 - Ví dụ: helloTung
// 	 - Đặt tên biến với ý nghĩa giống value
// 	 - Ví dụ: numberOfDayPerYear int = 365
// 	 - Hạn chế sự dụng các biến chỉ 1 ký tự.
// Cách chuyển kiểu dữ liệu

// var i int = 10
// var j float32 = float32(i)
// Kết quả j = 10*/
