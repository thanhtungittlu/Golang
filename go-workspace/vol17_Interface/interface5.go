package main

import "fmt"

func assert (i interface{}){
	s, ok := i.(int)  // Ép kiểu interfcae là int
	fmt.Println(s, ok)
}
func main(){
	var s interface{} =  40
	assert(s)
	var h interface{} = "Tung"
	assert(h)
}

/*
 s, ok := i.(T)
 Câu lệnh trên được hiểu là, nếu kiểu cụ thể của i là T thì s sẽ nhận giá trị cơ bản của i và ok là true.
 Nếu kiểu cụ thể của i không là T thì ok sẽ là false và s sẽ nhận zero value của kiểu T và chương
 */