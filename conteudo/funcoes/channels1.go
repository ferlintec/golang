package main

import (
	"fmt"
	"time"
)

// Channel permite que a thread principal finalize somente quando
// a goroutine finalizar também.
func say(s string, done chan string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	done <- "Terminei"

}

func main() {
	done := make(chan string)
	go say("# thread 1", done)
	fmt.Println(<-done)

	fmt.Println("thread principal executando...")

	go say("## thread 2", done)
	fmt.Println(<-done)

	fmt.Println("FIM thread principal")
}

// world
// world
// world
// world
// world
// Terminei
// FIM thread principal
