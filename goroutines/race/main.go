package raceconditions

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var result int
var waitGroup sync.WaitGroup
var m sync.Mutex

func Describe() {
	fmt.Printf("Race conditions\n\n")

	waitGroup.Add(2)

	go runProcess("p1", 10)
	go runProcess("p2", 10)

	waitGroup.Wait()

	fmt.Println("\nFinal result: ", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		m.Lock()
		result++
		fmt.Println(name, " => ", i, " | Partial result: ", result)
		m.Unlock()
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}

	waitGroup.Done()
}
