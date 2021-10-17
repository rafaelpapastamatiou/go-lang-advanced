package channels

import (
	"fmt"
	"sync"
)

func waitGroup() {
	fmt.Println("\n\nChannels - Wait Group")
	fmt.Println()
	channel := make(chan int)

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	for number := range channel {
		fmt.Println(number)
	}
}
