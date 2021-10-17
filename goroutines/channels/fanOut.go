package channels

import (
	"fmt"
)

func fanOut() {
	fmt.Println("\n\nChannels - Fan Out")
	fmt.Println()

	//c := fanOutGenerate(5, 10, 15, 20)

	//d1 := fanOutDivide(c)
	//d2 := fanOutDivide(c)
}

func fanOutGenerate(numbers ...int) chan int {
	channel := make(chan int)

	go func() {
		for _, n := range numbers {
			channel <- n
		}

		close(channel)
	}()

	return channel
}

func fanOutDivide(input chan int) chan int {
	channel := make(chan int)

	go func() {
		for number := range input {
			channel <- number / 2
		}
	}()

	return channel
}
