package main

import "fmt"

func ping(pings chan<- string, msg string) { // pings chan<- is a channel only for sending values
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // one channel for receives (pings) the other for sends (pongs)
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
