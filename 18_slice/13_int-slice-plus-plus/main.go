package main

import "fmt"

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++
	fmt.Println(mySlice[0])
	//mySlice[0] = mySlice[0] + 1
	//fmt.Println(mySlice[0])
	//mySlice[0] += 1
	//fmt.Println(mySlice[0])
	// alternate ways but not ideal
}
