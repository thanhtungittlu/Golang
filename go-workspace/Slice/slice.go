package main

import "fmt"

func main() {
	slice1 := []int{2, 5, 10, 12, 45, 50} // Mảng động,

	// 1 2 3 4 5 6 7 8 9  <-
	//         ^     ^         len = 4 , cap = 5
	slice2 := slice1[:]
	slice3 := slice1[3:]  // 3,4,5
	slice4 := slice1[:5]  // 0,1,2,3,4
	slice5 := slice1[3:5] // 3,4
	slice5[1] = 100       // Tất cả đều truy xuất đên 1 địa chỉ, nên khi thay đổi giá trị thì tất cả đều thay đổi
	fmt.Printf("%v,%v,%v\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("%v,%v,%v\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("%v,%v,%v\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("%v,%v,%v\n", slice4, len(slice4), cap(slice4))
	fmt.Printf("%v,%v,%v\n", slice5, len(slice5), cap(slice5))
	// Len: chiều dài hiện tại của mảng, Cap: Chiều dài tối đa có thể duyệt
	fmt.Println()
	a := make([]int, 5, 20)
	fmt.Printf("a %v,%v,%v\n", a, len(a), cap(a))
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3, 4, 5)
	fmt.Printf("a %v,%v,%v\n", a, len(a), cap(a))
	b := []int{6, 7, 8}
	a = append(a, b...) // Khi thêm append 1 mảng mới vào thì phải thêm dấu ...
	fmt.Printf("a %v,%v,%v\n", a, len(a), cap(a))
	fmt.Println()
	fmt.Println("Ngăn xếp")
	c := []int{1, 2, 3, 4, 5, 6}
	//d := append(c[:2], c[3:]...) // Thao tác với các số giữa
	//	fmt.Println(d)
	//f := c[:5]     // Thao tác với phần tử cuối
	//fmt.Println(f)
	e := c[1:] // Thao tác với phần tử đầu
	fmt.Println(e)
}
