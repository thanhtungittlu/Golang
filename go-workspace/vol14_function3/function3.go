package main

import "fmt"

func main()  {
	func () {
		x:= 10
		fmt.Println("Hello", x)
	}()
	// Không truy xuất dc x ở hàm main
	fmt.Println("---------")
	g := greeter {
		greeting : "Hello",
		name: "Tung",
	}
	g.greet()
	fmt.Println(g.name)
}
type greeter struct {
	greeting string
	name string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name ="TYuh"
}
