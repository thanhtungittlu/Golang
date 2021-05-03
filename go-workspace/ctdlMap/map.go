package main

import "fmt"

func main() {
	/*	studentNameAgeMap := map[string]int{ // trong ngoặc vuông là kiểu dữ liệu của key, ngoài ngoặc vuông là kiểu dữ liệu của value
			"Tùng":   22,
			"Thương": 23,
			"Nam":    22,
		}
		fmt.Println(studentNameAgeMap)
		fmt.Println() */
	studentNameAgeMap := make(map[string]int) // chưa biết giá trị key, value
	studentNameAgeMap = map[string]int{
		"Tùng":   22,
		"Thương": 23,
		"Nam":    22,
	}
	studentNameAgeMap["Hùng"] = 25 // Thêm mới or cập nhật
	fmt.Println(studentNameAgeMap)
	fmt.Println()
	fmt.Println(studentNameAgeMap["Hùng"])
	delete(studentNameAgeMap, "Nam") // Xóa
	fmt.Println(studentNameAgeMap)
	_, isExist := studentNameAgeMap["Nam"] // kiểm tra xem còn Key trong Map k
	fmt.Println(isExist)
	fmt.Println(len(studentNameAgeMap)) // Chiều dài, số cặp (key: Value) của Map
	copyMap := studentNameAgeMap        // Copy Map, nhưng cùng vị trí truy xuất nên, khi thay đổi 1 trong 2 thì cả 2 đều thay đổi
	fmt.Println(copyMap)
}
