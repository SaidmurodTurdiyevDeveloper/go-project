package main

import "fmt"

func main() {
	var myMap map[string]int = make(map[string]int)

	myMap["first"] = 10
	myMap["second"] = 8

	for _, v := range myMap {
		fmt.Println(v)
	}

}
