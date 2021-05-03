package main

import "fmt"

func main() {
	m := map[string]int{
		"Tùng": 19,
		"Tom":  20,
	}
	if age, isExist := m["Tùng"]; isExist == true {
		fmt.Println(age)
		fmt.Println()
	} else {
		fmt.Print("Key không có trong map")
	}
	fmt.Println()
	a := 90
	switch a {
	case 1:
		fmt.Println("No")
	case 2:
		fmt.Println("No")
	case 90:
		fmt.Println("Yes")
	default:
		fmt.Println("Not others")
	}
	/*
		number :=3
		switch {
			case number <=3:
				fmt.Println("Nhỏ hơn hoặc bằng 3")
				fallthrough // Vẫn xét tiếp các trường hợp khác
			case number <=5:
				fmt.Println("Nhỏ hơn hoặc bằng 5")
			default:
				fmt.Println(" Không thuộc trường hợp nào")

		}
	*/

}

// age, isExist := m["Tung"]  : Kiếm tra xem Key "Tung" có trong Map không,
// isExist sẽ trả về true or false, nếu true thì value sẽ gán vào biến age
// || hoặc,  && và
