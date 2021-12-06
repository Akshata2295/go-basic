package main

import "fmt"

func main() {
	fmt.Println("1 dimensional array")
	var intArr [5]int
	fmt.Println(intArr)

	fmt.Println("2 dimensional array")
	var doubleArray [3][3]int
	doubleArray[1][1] = 8
	doubleArray[2][2] = 5
	fmt.Println(doubleArray)
}
