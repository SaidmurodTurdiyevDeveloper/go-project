package main

import "fmt"

//this viewo about array
func main() {

	fmt.Println("Array exaple\n")
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[3] = "Banana"
	fruits[4] = "Orange"
	fmt.Println("Your fruit list hes this fruet", fruits)
	fmt.Println("Your first fruit", fruits[0])
	fmt.Println("Fruitd list size", len(fruits))

	fmt.Println("\n====================\n")

	var vegitable = [6]string{"mashroom", "tometo", "carrot", "cucumber"}
	fmt.Println("You vegitable list has this vegitable")

	for i := 0; i < len(vegitable); i++ {
		fmt.Println(i, vegitable[i])
	}

}
