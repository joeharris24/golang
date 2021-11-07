package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, argError{arg: arg, prob: "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	val, err := f1(42)
	fmt.Println(val)
	// > -1
	fmt.Println(err)
	// > can't work with 42
	val2, err2 := f2(42)
	fmt.Println(val2)
	// > -1
	fmt.Println(err2)
	// > 42 - can't work with it
}
