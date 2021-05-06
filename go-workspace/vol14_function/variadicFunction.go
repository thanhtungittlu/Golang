// Hàm bất định
package main

import "fmt"

func find(num int, nums ...int)  {
	fmt.Printf("Type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "Found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
	fmt.Println()
}

func main()  {
	find(89,89,90,05)
	find(45,56,67,45,90,100)
	find(45,56,67,90,100)
	find(87)
}