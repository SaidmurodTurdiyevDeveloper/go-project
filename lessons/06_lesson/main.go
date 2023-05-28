package main

import "fmt"

func main() {
	// fmt.Println("You have choose any number inside 1 to 10")
	// var a int
	// fmt.Scan(&a)
	// switch a {
	// case 1:
	// 	fmt.Print("you choose one")
	// case 2:
	// 	fmt.Print("you choose two")
	// case 3:
	// 	fmt.Print("you choose three")
	// case 4:
	// 	fmt.Print("you choose four")
	// case 5:
	// 	fmt.Print("you choose five")
	// case 6:
	// 	fmt.Print("you choose six")
	// case 7:
	// 	fmt.Print("you choose seven")
	// case 8:
	// 	fmt.Print("you choose eight")
	// case 9:
	// 	fmt.Print("you choose nine")
	// 	fallthrough
	// case 10:
	// 	{
	// 		if a == 9 {
	// 			fmt.Println("\nwith fallthrough")
	// 		}
	// 		fmt.Print("you choose ten")
	// 	}

	// default:
	// 	fmt.Print("you choose wrong number")
	// }

	var days = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	for index, day := range days {
		/*
			if index == 2 {
				//this code work when day will be wednesday and index == 2 then did not work buttom code this time inside of for loops barracks
				continue
			}
		*/

		/*
			if index == 5 {
				//this code work when day will be saturday and index == 5 then did not work after for loops
				break
			}
		*/
		/*
			if index == 1 {
				//this code work when day will be Tuesday and index == 1 then work my label code and did not work this for loops
				goto myLabel
			}
		*/
		fmt.Printf("valuea of %v\n", day)
	}
	/*
	   myLabel:
	   	fmt.Println("My label for go to")

	*/
}
