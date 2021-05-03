package main

import (
	"fmt"
)

const (
	red = iota // int, giá trị ban đầu là 0, rồi tăng dần, _ = iota, để bắt đầu từ 1
	yellow
	blue
	black
)

func main() {
	const i int16 = 42
	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T\n", red, red)
	fmt.Printf("%v,%T\n", yellow, yellow)
	fmt.Printf("%v,%T\n", blue, blue)
	fmt.Printf("%v,%T\n", black, black)
}
