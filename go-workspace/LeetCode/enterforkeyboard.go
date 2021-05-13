package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your city: ")

	//fmt.Println(reader)
	city, _ := reader.ReadString('\n')
	fmt.Printf("%T  ", city)

}
