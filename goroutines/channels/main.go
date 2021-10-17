package channels

import (
	"fmt"
)

func Describe() {
	fmt.Printf("Channels\n\n")

	waitGroup()
	semaphore()
	pipelinePattern()
	funIn()
}
