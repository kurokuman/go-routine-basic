package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	//buffer  error
	// ch <- 300
	fmt.Println(len(ch))

	//success case
	x := <-ch
	fmt.Println(x)
	ch <- 300
	fmt.Println(len(ch))

	//for
	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}
