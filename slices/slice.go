package main

import "fmt"

func main() {
	s := make([]int, 10)
	fmt.Println("s initialized as ", s)
	s[1] = 2
	s[2] = 5
	s[3] = 3
	s[4] = 6
	s[5] = 8
	s[6] = 1
	sLength := len(s)
	fmt.Println(s)
	fmt.Println("length of slice is: ", sLength)
}
