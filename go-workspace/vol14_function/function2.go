package main

import "fmt"

func main()  {
	res, err := calDivide(4,2)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(res)
	}
	
}

func calDivide(a,b int) (int, error)  {
	result := 0
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	result = a / b 
	return result, nil
}