package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	out1 := make(chan string)
	out2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 500)
			out1 <- "out1 output"
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 100)
			out2 <- "out2 output"
		}
		wg.Done()
	}()

	for {
		select {
		case msg, open := <-out1:
			fmt.Println(msg, "-", open)
		case msg, open := <-out2:
			fmt.Println(msg, "-", open)
		}
	}

}
