package main

import "fmt"

func rectProps(dai, rong float64) (float64, float64){
	var chuvi= (dai + rong)*2
	var dientich = dai*rong
	return chuvi,dientich
}

func main()  {
	var a,b float64
	a = 5.2
	b = 4.8
	fmt.Println(rectProps(a,b))
}