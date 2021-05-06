package main	
import "fmt"

func subtactOne(number []int)  {
	for i := range number {
		number[i] -= 2
	}
}

func main()  {
	nos := []int{8, 7, 6}
	fmt.Println("Slice before function call", nos)
	subtactOne(nos)
	fmt.Println("Slice before function call", nos)
}