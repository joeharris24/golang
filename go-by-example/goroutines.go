package main

import (
	"fmt"
	"time"
)

/*
func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	go f("goroutine")

	go func() {
		fmt.Println("msg")
	}()

	time.Sleep(time.Second)
	fmt.Println("done")
}

// Goroutines and channels
*/

func numbers(c chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100)
		c <- i
	}
}

func letters() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(time.Millisecond * 25)
		fmt.Printf("%c\n", i)
	}
}

func main() {

	//var wg sync.WaitGroup
	//wg.Add(2)

	c := make(chan int)

	go numbers(c)
	go letters()

	for range c {
		select {
			case c <
		}
		fmt.Println(c)
	}

	//fmt.Println("test1")

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(i)
	//	}
	//	wg.Done()
	//}()

	//fmt.Println("test2wdqdwadwdaddwaaddwefffffffffffffffwdaw")

	//wg.Wait()
	fmt.Scanln()

}
