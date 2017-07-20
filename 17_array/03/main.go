package main

import "fmt"

func main() {
	var x [256]int

	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = i
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
		//	fmt.Printf("%v - %T - %b\n", v, v, v) will print from 0 to 50
	}
}
