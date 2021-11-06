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
