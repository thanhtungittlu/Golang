// Cho 1 mảng các số tự nhiên, và 1 số. Đưa ra chỉ số khi tổng của 2 số đó cổng lại bằng số đã cho
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				fmt.Println(i, j)
				return
			}
		}
	}
}

func main() {
	nums := []int{2, 4, 6, 8, 5, 7, 8, 0, 4, 7, 9, 3, 11, 45, 67, 3, 7, 90}
	target := 52
	twoSum(nums, target)
}

// func main()  {
// 	fmt.Println("---Leetcode---")
// 	nums := []int{2,7,11,15,7,2}
// 	target := 9

// 	for i := 0; i < len(nums); i++ {
// 		for j := i+1; j < len(nums); j++ {
// 			if ( (nums[i] + nums[j]) == target) {
// 				fmt.Println(i,j)
// 				break
// 			}
// 		}
// 		break
// 	}
// }
