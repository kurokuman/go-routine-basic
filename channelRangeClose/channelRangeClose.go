package main

import (
	"fmt"
)

func goroutine1(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		//chに値が渡され、 mainのfor　range でその値が取り出される
		ch <- sum
	}
	//main のforで空のchから値を取らないように、closeを宣言
	close(ch)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(s))

	go goroutine1(s, ch)

	for c := range ch {
		fmt.Println(c)
	}
}
