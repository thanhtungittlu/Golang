package main

import (
	"log"
	"time"
)


// google Search
func googleSearch(result chan string)  {
	time.Sleep(3* time.Second)
	result <- "found from Google"
}

// bing Seacrch

func bingSearch(result chan string)  {
	time.Sleep(5* time.Second)
	result <- "found from bing"
}

func main()  {
	chanGoogle := make(chan string)
	chanBing := make(chan string)

	go googleSearch(chanGoogle)
	go bingSearch(chanBing)

	select{
	case result:= <- chanGoogle :
		log.Println(result)
	case result:= <- chanBing :
		log.Println(result)
	}
	log.Println("done")
}