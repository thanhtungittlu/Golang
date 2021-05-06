package main

import "fmt"



type Phanso struct{
	tuso int
	mauso int
}

func Tong(a,b Phanso) (sum Phanso)  {
	sum.tuso = a.tuso*b.mauso + a.mauso*b.tuso
	sum.mauso = a.mauso * b.mauso
	return RutGon(sum)
}

func Hieu(a,b Phanso) (minus Phanso)  {
	minus.tuso = a.tuso*b.mauso - a.mauso*b.tuso
	minus.mauso = a.mauso * b.mauso
	return RutGon(minus)
}

func Tich(a,b Phanso) (t Phanso)  {
	t.tuso = a.tuso*b.tuso
	t.mauso = a.mauso*b.mauso
	return RutGon(t)
}

func Thuong(a,b Phanso) (t Phanso )   {
	t.tuso = a.tuso*b.mauso
	t.mauso = a.mauso*b.tuso
	return RutGon(t)
}

func ucln(a int, b int) int {
	if a%b == 0 {
		return b
	}
	return ucln(b, a%b)
}
func RutGon(a Phanso) (b Phanso)  {
	b.tuso = a.tuso/ucln(a.tuso,a.mauso)
	b.mauso = a.mauso/ucln(a.tuso,a.mauso)
	return b
}
func output(a Phanso){
	if (a.mauso == 1) {
		fmt.Println(a.tuso)
	}else{
		fmt.Println(a.tuso,"/",a.mauso)
	}
	
}

func main()  {
	a := Phanso{1,2}
	b := Phanso{1,4}
	fmt.Print("Tổng của 2 phân số là: ")
	output(Tong(a,b))
	fmt.Print("Hieu của 2 phân số là: ")
	output(Hieu(a,b))
	fmt.Print("Tich của 2 phân số là: ")
	output(Tich(a,b))
	fmt.Print("Thương của 2 phân số là: ")
	output(Thuong(a,b))
}