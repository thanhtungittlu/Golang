package main

import "fmt"

func main() {
	fmt.Println("-----Hàm-----")
	a:= 20
	b:= 50
	max := findMax(a,b) // Max và kết quả trả về sẽ cùng 1 bộ nhớ để tiết kiệm bộ nhớ
	fmt.Println(*max) // Max cũng là 1 con trỏ
	min := findMin(a,b) // Min dc cấp phát bộ nhớ mới
	fmt.Println(min)
	fmt.Println("----------")
	k := []int{1,2,3,4,5}
	sum := findSum(k)
	fmt.Println(sum)
	sumCach2 := findSumCach2("Ket qua: ", 1, 2, 3, 4, 5) //Tự add vào 1 slice
	fmt.Println(sumCach2)
}
func findMax( number1,number2 int)  *int { // Kết quả trả về là *int (con trỏ)
	if (number1 > number2){
		
		return &number1
	}else{
		return &number2
	}	
}
func findMin(number1,number2 int) (min int) { // Cách 2
	if (number1 < number2){
		min = number1
	}else{
		min = number2
	}
	return 
}
func findSum(a []int) (sum int)  {
	for _, i:= range a{ // range a lấy các giá trị của slice a
		sum+=i
	}
	return	
}
func findSumCach2(s string, a ...int) (sum int)  {
	for _, i:= range a{ // range a lấy các giá trị của slice a
		sum+=i
	}
	fmt.Print(s)
	return	
}
