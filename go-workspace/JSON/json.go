package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var soi = map[string]interface{}{
		"id":   1,
		"name": "My",
		"age":  18,
	}
	fmt.Println(soi)
	// Chuyển đổi dữ liệu sang string json
	data, _ := json.Marshal(soi)

	var myJsonString = string(data)
	fmt.Println(myJsonString)
	fmt.Printf("%T\n", myJsonString)
	fmt.Println(soi)

}
