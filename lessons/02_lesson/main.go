package main

import (
	"fmt"
	"time"
)

func main() {
	//examples with time

	currentTime := time.Now()
	fmt.Println("\n=================\n")
	fmt.Println(currentTime)
	fmt.Println("\n=================\n")
	fmt.Println("UTC")
	fmt.Println(currentTime.UTC())

	/*
		with any format
		Format string date always will be 2006 2 january 15:04:05 Monday

		if you write like bottom line Format string does not working
		currentTime.Format("04.07.2008")
	*/

	fmt.Println("With Format")
	fmt.Println(currentTime.Format("01.02.2006"))

	createdDate := time.Date(1998, 12, 14, 5, 25, 37, 123, time.Local)

	fmt.Println("\n=================\n")

	fmt.Println("My birthday in Tashkent")
	fmt.Println(createdDate)

	/*
		UTC another Cordinated Unversal Time London is first time +0 Tashkent +5
	*/
	fmt.Println("\n=================\n")
	location, _ := time.LoadLocation("UTC")
	if location != nil {
		fmt.Println("My Birthday in London \n", createdDate.In(location).Format("2006_01_02 Mon 15:04:05"))
	}
	fmt.Println("\n=================\n")
	location2, _ := time.LoadLocation("Asia/Istanbul")
	if location != nil {
		fmt.Println("My Birthday in Istanbul \n", createdDate.In(location2))
	}

}
