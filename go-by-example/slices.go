package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	// [   ]

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	// [a b c]

	s = append(s, "d", "e")
	fmt.Println(s)
	// [a b c d e]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	// [a b c d e]

	l := s[2:5]
	fmt.Println(l)
	// [c d e]

	l = s[:3]
	fmt.Println(l)
	// [a b c]

	t := []string{"m", "n", "o"}
	fmt.Println(t)
	// [m n o]

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

}

/*
	Go blog: slices usage and internals

	Slice type is an abstraction built atop Go's array type
	Array type definition specifies length and element type
		var a [4] int
		a[0] = 1
		i := a[01]
		// i == 1
	Can have compiler count array elements for you
		b := [...]string{"P", "T"}
		// b type is [2]string


	Slice type has no specified length
	Slice literal is declared just like array literal, except you leave out the element count
		letters := []string{"a", "b", "c"}
	A slice can be created with the built-in function make
		func make([]T, len, cap) []T

		var s []byte
		s = make([]byte, 5, 5)
		// s = []byte{0, 0, 0, 0, 0}
	When the capacity argument is omitteed, it defaults to specified length.
		s := make([]byte, 5)
		// len(s) == 5, cap(s) == 5

	A slice can also be formed by slicing an existing slice or array
		b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
		// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b

		x := [3]string{"G", "d", "a"}
		s := x[:] // a slice referencing the storage of x

	A slice is a descriptor of an array segment, consisting of pointer to the array, length of the segment, and its capacity
	Slicing doesn't copy the slice's data, it creates a new slice value that points to the original array
	This makes slice operations as efficient as manipulating array indicies
	Modifying the elements (not the slice itself) of a reslice modifies the elements of the original slice
		d := []byte{'r', 'o', 'a', 'd'}
		e := d[2:]
		// e == []byte{'a', 'd'}
		e[1] = 'm'
		// e == []byte{'a', 'm'}
		// d == []byte{'r', 'o', 'a', 'm'}

	A slice can be regrown through
	 	s = s[:cap(s)]
	However this can't go to previous elements in the underlying array/slice again

	To copy slices
		func copy(dst, src []T) int

	To append to a slice
	func append(s []T, x ...T) []T
		a := make([]int, 1)
		// a == []int{0}
		a = append(a, 1, 2, 3)
		// a == []int{0, 1, 2, 3}

	To append one slice to another use ...
		a := []string{"J", "P"}
		b := []string{"G", "R"}
		a = append(a, b...) // equivalent to append(a, b[0], b[1])
		// a == []string{"J", "P", "G", "R"}

	To stop continued holding of full array in memory (rather than the slice we're looking at within)
	We need to copy the data we're interested in into a new slice, allowing garbage collector to release the underlying array from memory

*/
