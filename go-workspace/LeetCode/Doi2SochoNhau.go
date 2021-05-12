package main

import "fmt"

func swapTwoNumber(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b

}

func main() {
	a, b := 5, 6

	fmt.Println("After: ", a, " ", b)
	swapTwoNumber(&a, &b)
	fmt.Println("Before: ", a, " ", b)
}
