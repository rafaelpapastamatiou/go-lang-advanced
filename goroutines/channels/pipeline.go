package channels

import (
	"fmt"
)

func pipelinePattern() {
	fmt.Println("\n\nChannels - Pipeline Pattern")
	fmt.Println()

	numbers := pipelineGenerate(2, 4, 6, 8, 10)
	result := pipelineDivide(numbers)

	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
}

func pipelineGenerate(numbers ...int) chan int {
	channel := make(chan int)

	go func() {
		for _, number := range numbers {
			channel <- number
		}
	}()

	return channel
}

func pipelineDivide(input chan int) chan int {
	channel := make(chan int)

	go func() {
		for number := range input {
			channel <- number / 2
		}
	}()

	return channel
}
