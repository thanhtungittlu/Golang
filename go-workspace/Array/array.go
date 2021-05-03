package main

import "fmt"

func main() {
	points := [...]int{7, 10, 5} //Cach 1
	// var points [3]int          //Cach2
	// points[0] = 7
	// points[1] = 10
	// points[2] = 5
	fmt.Printf("%v,%T\n", points, points)
	fmt.Printf("%v,%T\n", points[1], points[1])
	fmt.Println(len(points)) // Chiều dài của mảng hiện tại
	points1 := &points       // copy mảng ban đầu, khi thay đổi thì thay đổi cả 2
	// points1 := points // Copy mảng ban đầu, không liên quan gì tới mảng cũ
	points1[1] = 30
	fmt.Println(points1)
	fmt.Println(points)
}
