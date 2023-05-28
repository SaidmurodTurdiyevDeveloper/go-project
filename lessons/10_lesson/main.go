package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print("i=", i, ",")
		go f(i)
	}
	fmt.Println("Tugadi")
	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	fmt.Println("\n---", n, "---")
	for i := 0; i < 10; i++ {

		amp := time.Duration(rand.Intn(250))
		fmt.Println(n, ":", i, "amp", amp)
		time.Sleep(time.Millisecond * amp)

	}
}
