package main

import (
	"fmt"
	"time"
)

func main() {
	var b int = 0
	fmt.Println("This my channle examples")
	/*
		var channel_1 = make(chan string)
		go pong(channel_1)
		// go ping(channel_1)
		// go pung(channel_1)
		go print(channel_1)
	*/

	/*
		this below code we cannot see true result (b=200) because our code work synchronous first goroutune get value of b
		then until set new value that old value of b can be changed
	*/

	/*
		var channel_1 = make(chan int)
		var channel_2 = make(chan int)

		go func() {
			for i := 0; i < 100; i++ {
				t := b
				channel_1 <- t + 1
				// time.Sleep(time.Millisecond*5)
			}
		}()

		go func() {

			for i := 0; i < 100; i++ {
				t := b
				channel_2 <- t + 1
				// time.Sleep(time.Millisecond*4)
			}
		}()

		go func() {
			for {
				select {
				case msg1 := <-channel_1:
					{
						b = msg1
					}
				case msg2 := <-channel_2:
					{
						b = msg2
					}

				case <-time.After(time.Second):
					fmt.Println("time out b=", b)
					// default:
					// 	fmt.Println("nathing redy")
				}
			}
		}()
	*/

	/*
		i will try with a channel but our true(b=200) result cannot got
	*/
	var channel_1 = make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			t := b
			channel_1 <- t + 1
			time.Sleep(time.Millisecond * 5)
		}
	}()

	go func() {

		for i := 0; i < 100; i++ {
			t := b
			channel_1 <- t + 1
			time.Sleep(time.Millisecond * 10)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-channel_1:
				{
					if b == msg1 {
						fmt.Println(b)
					}
					b = msg1
				}

			case <-time.After(time.Second):
				fmt.Println("time out b=", b)
				// default:
				// 	fmt.Println("nathing redy")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)

}

/*

// this function can send and recive channel message
func pong(c chan string) {
	for i := 0; ; i++ {
		fmt.Println("pong", i)
		c <- "messasge pong " + fmt.Sprint(i)
	}
}

// this function can only send channel message
func ping(c chan<- string) {
	for i := 0; ; i++ {
		fmt.Println("ping", i)
		c <- "message ping " + fmt.Sprint(i)
	}
}

// this function can only send channel message
func pung(c chan<- string) {
	for i := 0; ; i++ {
		fmt.Println("pung", i)
		c <- "message pung " + fmt.Sprint(i)
	}
}

// this function can only recive channel message
func print(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 2)
	}
}
*/
