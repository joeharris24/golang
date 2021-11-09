package main

import "fmt"

func zeroval(val int) {
	val = 0
}

func zeroptr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 1
	fmt.Println(i)
	// > 1

	zeroval(i)
	fmt.Println(i)
	// > 1

	zeroptr(&i)
	fmt.Println(i)
	// > 0

}
