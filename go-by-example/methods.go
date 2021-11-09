package main

import "fmt"

type rect struct {
	width, height int
}

// Using pointer allows mutation of receiving structure
func (r *rect) area() int {
	r.height = 100
	return r.width * r.height
}

// Using value for reciever means calculations can be done but no mutation
func (r rect) areacopy() int {
	r.height = 100
	return r.width * r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println(r.areacopy())
	fmt.Println(r)

}
