package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.height = 100
	return r.width * r.height
}

func (r rect) areacopy() int {
	r.height = 100
	return r.width * r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println(r.areacopy())
	fmt.Println(r)

}
