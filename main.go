package main

import (
	"fmt"

	"github.com/rafaelpapastamatiou/go-lang-advanced/goroutines/channels"
	"github.com/rafaelpapastamatiou/go-lang-advanced/goroutines/concurrency"
	raceconditions "github.com/rafaelpapastamatiou/go-lang-advanced/goroutines/race"
)

func main() {
	concurrency.Describe()
	separator()
	raceconditions.Describe()
	separator()
	channels.Describe()
	separator()
}

func separator() {
	fmt.Printf("\n\n#================================================================#\n\n")
}
