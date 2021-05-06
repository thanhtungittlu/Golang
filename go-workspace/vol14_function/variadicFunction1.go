package main

import "fmt"

func change(s ...string)  {
	s[0] = "Go"
	s = append(s, "play")
	fmt.Println(s)
}

func main()  {
	welcome := []string{"hello","world"}
	change(welcome...)
	fmt.Printf("%T\n",welcome)
	fmt.Println(welcome)
}