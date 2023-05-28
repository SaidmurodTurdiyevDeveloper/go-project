package main

import "fmt"

//this is about poiter
func main() {
	// mvvalue := 23
	// var ptr = &mvvalue
	// fmt.Println("value = ", mvvalue)
	// fmt.Println("value of pointer is ", ptr)
	// fmt.Println("pointer of value is ", *ptr)

	/*
		x := 5
		changaToZeroWithoutPointer(x)
		fmt.Println("x= ", x) // x is still

		changeToZeroWithPointer(&x)
		fmt.Println("x= ", x) // x is 0
	*/

	/*
		//create ponter like this
		t := new(int)
		changeToZeroWithPointer(t)
		fmt.Println("t= ", *t, " adress ", t)
	*/

	/*
		//example
		var y float64 = 12
		square(&y)
		fmt.Println(y)
	*/
	var (
		x = 1
		y = 2
	)
	fmt.Println("Before", x, y)
	swap(&x, &y)
	fmt.Println("After", x, y)

}

// func changeToZeroWithPointer(xPtr *int) {
// 	*xPtr = 0
// }
// func changaToZeroWithoutPointer(x int) {
// 	x = 0
// }
// func square(x *float64) {
// 	*x = *x * *x
// }

func swap(x *int, y *int) {
	*x = 2
	*y = 1
}

/**
How do you get the memory address of a variable?
We can get adress of variables with & syimbol like this &x in here x is variable.
*/
