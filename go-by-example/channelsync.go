package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("dont")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done // Block until we receive notification from worker on this channel that it's done
}
