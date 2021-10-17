package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func Describe() {
	fmt.Printf("Concurrency\n\n")

	waitGroup.Add(2)

	go runProcess("p1", 10)
	go runProcess("p2", 10)

	waitGroup.Wait()
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, " => ", i)

		t := time.Duration(rand.Intn(255))

		time.Sleep(time.Millisecond * t)
	}

	waitGroup.Done()
}
