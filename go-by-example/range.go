package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	// _, num in range = index, value (so blanking index)
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	// > 9

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
			// index: 1
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%v --> %v\n", k, v)
	}
	// > a --> apple \n b --> banana

	for k := range kvs {
		fmt.Println("key:", k)
	}
	// key: b \n key: a

	for i, c := range "go" {
		fmt.Println(i, c)
	}
	// 0 103 \n 1 111

}
