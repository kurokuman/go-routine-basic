package main

import (
	"fmt"
)

func goroutine1(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		ch <- sum
	}
	close(ch)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	ch := make(chan int)

	go goroutine1(s, ch)

	for c := range ch {
		fmt.Println(c)
	}
}
