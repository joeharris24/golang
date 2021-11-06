package main

import "fmt"

// Can support anonymous functions (which can form closures)
// Anonymous functions are good for defining functions inline without naming

// intSeq returns another function, which we define anonymously in the body of intSeq
// returned function closes over the variable i to form a closure

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	// 1
	fmt.Println(nextInt())
	// 2
}

// we call intSeq, assigning the result (a function) to nextInt
// This function value captures its own i value, which will be updated each time we call nextInt
