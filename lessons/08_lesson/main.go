package main

import "fmt"

func main() {
	/*
		var result, t, s = f3()

		fmt.Println(result)
		fmt.Println(t)
		fmt.Println(s)
	*/

	/*
		var s = add(1, 5, 8, 3)
		fmt.Println(s)

		var xs = []int{1, 2, 3, 4, 5}
		s := add(xs...)
		fmt.Println(s)
	*/

	/*
		//function inside function
		var add = func(a int, b int) int {
			return a + b
		}
		result := add(4, 5)
		fmt.Println(result)

		count := 1
		increment := func() int {
			count++
			return count
		}
		fmt.Println(increment())
		fmt.Println(increment())
	*/

	/*
		// style 5
		var nextEven = makeGenerator()
		fmt.Println(nextEven())
		fmt.Println(nextEven())
		fmt.Println(nextEven())
	*/

	/*
		// style 6
		fact := factorial(5)
		fmt.Println(fact)
	*/

	/*
		//style 7
		defer second()
		first()
		third()
		fmt.Println("finish")
	*/

	/*
		defer func() {
			panickArg := recover()
			fmt.Println("after recover")
			fmt.Println(panickArg, " recover")
		}()
		panic("PANICK")
		fmt.Println("after Panick")
	*/

	var fib = fib(8)
	fmt.Println(fib)
}

// // style 1
// func f() (r int, t int) {
// 	r = 5
// 	t = 8
// 	return
// }

// // style 2
// func f2() (r int, t int) {
// 	return 6, 7
// }

// // style 3
// func f3() (r int, t int, s int) {
// 	s = 8
// 	return 6, 7, s
// }

// // style 4
// func add(args ...int) int {
// 	var total int
// 	for _, v := range args {
// 		total += v
// 	}
// 	return total
// }

// // style 5
// func makeGenerator() func() uint {
// 	i := uint(0)
// 	return func() (ret uint) {
// 		ret = i
// 		i += 2
// 		return
// 	}
// }

// // style 6
// func factorial(index uint) uint {
// 	if index == 0 {
// 		return 1
// 	}
// 	return index * factorial(index-1)
// }

// //style 7
// func first() {
// 	fmt.Println("1st")
// }
// func second() {
// 	fmt.Println("2nd")
// }
// func third() {
// 	fmt.Println("3rd")
// }

// examples
func sum(slice []int) int {
	var total int
	for _, v := range slice {
		total += v
	}
	return total
}

func half(number uint) (uint, bool) {
	if number%2 == 0 {
		return number / 2, true
	} else {
		return number / 2, false
	}
}

func max(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	var max int = slice[0]
	for _, v := range slice {
		if max < v {
			max = v
		}
	}
	return max
}

func makeEvenGenerator() func() (number int) {
	var index = 0
	return func() (number int) {
		number = index
		index += 2
		return
	}
}
func makeOddGenerator() func() (number int) {
	var index = 1
	return func() (number int) {
		number = index
		index += 2
		return
	}
}

func fib(fibNumber uint) uint {
	if fibNumber == 0 {
		return 0
	}
	if fibNumber == 1 {
		return 1
	}
	return fib(uint(fibNumber-1)) + fib(uint(fibNumber-2))
}
