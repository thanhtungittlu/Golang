package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println()
	j := 1
	for {
		if j%2 == 1 {
			if j == 5 {
				j++
				continue
			}
			fmt.Println(j)
		}
		j++
		if j > 11 {
			break
		}
	}
	fmt.Println()
	fmt.Println()
	array := [3]int{2, 3, 4}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
	fmt.Println()
	for index, value := range array {
		fmt.Println(index, "  ", value)
	}
	fmt.Println()

	m := map[string]int{
		"Tung":   22,
		"Thuong": 23,
	}
	for key, value := range m {
		fmt.Println(key, "  ", value)
	}
	fmt.Println()
	s := "Hello World"
	for index, value := range s {
		fmt.Println(index, "  ", string(value))
	}
}
