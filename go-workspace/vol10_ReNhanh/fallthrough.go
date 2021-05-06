package main

import "fmt"

func main()  {
	num := 20
	switch  {
	case num <10:
		fmt.Println("<10")
		fallthrough
	case num <20:
		fmt.Println("<20")
		fallthrough
	case num <30:
		fmt.Println("<30")
		fallthrough
	case num <40:
		fmt.Println("<40")
	}
}
// fallthrough sử dụng khi muốn kiểm tra tất cả các case, vì mặc định khi có case đúng chương trình sẽ thoát ra