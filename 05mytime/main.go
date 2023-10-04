package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello, welcome to the time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("Monday"))

	// creating time or date
	createdDate := time.Date(2001, time.October, 27, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
