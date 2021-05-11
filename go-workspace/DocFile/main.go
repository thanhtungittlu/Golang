package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"strings"
)

func countFirstFile(result chan int, filePath string, keyword string)  {
	// Tính số lần xuất hiện từ khóa trong file
	var numberOfOcc int
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		result <- 0
		return 
	}
	numberOfOcc = strings.Count(string(fileContent), keyword )
	result <- numberOfOcc
	defer close(result)


}
func countSecondFile(result chan int, filePath string, keyword string)  {

	var numberOfOcc int
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		result <- 0
		return 
	}
	numberOfOcc = strings.Count( string(fileContent), keyword )
	result <- numberOfOcc
	
	defer close(result)
}

func main()  {
	countFirstChan := make (chan int)
	countSecondChan := make (chan int)

	
	go countFirstFile(countFirstChan,"1.txt", "bạn")
	go countSecondFile(countSecondChan,"2.txt", "bạn")

	log.Println("Kết quả bài toán: ", <-countFirstChan + <-countSecondChan)	
	fmt.Println("Done")
}

