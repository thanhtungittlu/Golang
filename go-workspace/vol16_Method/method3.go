package main

import "fmt"

type address struct{
	city string
	state string
}
func (a address) fullAddress(){
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct{
	firstName string
	lastName string
	address
}

func main()  {
	p:= person {
		firstName : "Tùng",
		lastName : "Lê",
		address: address{
			city : "Hà Nội",
			state : "Hoàng Mai",
		},
	}
	p.fullAddress() // đáng ra phải là p.address.fullAddress()
}