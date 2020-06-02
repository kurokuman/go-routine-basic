package main

import (
	"fmt"
	"time"
)

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(3)
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 10
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)

	go goroutine1(s, c)
	go goroutine2(s, c)
	// goroutine1,2のうち、先にc<-　をした方の値がxに渡される.
	x := <-c
	fmt.Printf("x : %d \n", x)
	y := <-c
	fmt.Printf("y : %d \n", y)
}
