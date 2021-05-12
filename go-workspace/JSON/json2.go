package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var payload = `{
		"id":   1,
		"name": "My",
		"age":  18
	}`
	fmt.Println(payload)

	var data = map[string]interface{}{}

	json.Unmarshal([]byte(payload), &data) // Chuyển về Map
	fmt.Println(data)

}
