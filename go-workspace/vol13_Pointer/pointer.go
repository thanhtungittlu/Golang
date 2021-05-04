package main

import "fmt"

func main() {
	var a int = 12
	var b *int = &a // Con trỏ b, trỏ về địa chỉ của biến a
	fmt.Println(a ,b )
	// Xuất ra giá trị tại địa chỉ của con trỏ, dùng *
	fmt.Println (*b)
	a = 24;
	fmt.Println(a, *b) // Khi a thay đổi thì cái giá trị của con trỏ b cũng thay đổi theo
	*b = 32
	fmt.Println(a, *b)

	fmt.Println("----Thao tác trên mảng-----")

	c:= [3]int{1, 2, 3}
	d:= c
	fmt.Println(c,d)
	c[1] = 9
	fmt.Println(c,d)
	//=>> Trên mảng (array) khi copy sẽ cấp phát vùng nhớ mới
	fmt.Println("----Thao tác trên slice-----")

	g:= []int{1, 2, 3}
	h:= g
	fmt.Println(g,h)
	g[1] = 9
	fmt.Println(g,h)
	fmt.Println(&g[1],&h[1])
	// =>> Thao tác trên slice khi copy vẫn là cùng 1 vùng nhớ
	fmt.Println("----Thao tác trên struct-----")
	var i *myStruct // Con trỏ của kiểu tự định nghĩa (struct)
	fmt.Println(i)
	i = new(myStruct) // Cấp phát 1 vùng nhớ ban đầu
	fmt.Println(i)
	(*i).number = 20 // hoặc có thể viết gọn hơn: i.number = 20 vì Golang tự hiểu
	fmt.Println((*i).number,i)
	fmt.Println("----Thao tác trên map-----")
	var k = map[string]string{"Tung":"Thuong","Nam":"Hang"}
	t := k
	fmt.Println(k,t)
	k["Nam"] = "Van"
	fmt.Println(k,t)
	// Bản chất k cũng là 1 con trỏ, trỏ đến các địa chỉ key và value, nên khi thao trên map cũng giống thao tác trên con trỏ
}
type myStruct struct{
	number int
}
