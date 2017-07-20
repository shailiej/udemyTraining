package main

import "fmt"

func half(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}

func main() {
	h, even := half(5)
	fmt.Println(h, even)
}

/* package main

import "fmt"

func main() {
	//ret, xs := boolie(5)
	//fmt.Println(ret, xs)
	fmt.Println(boolie(5))
}

func boolie(z int) (float64, bool) {
	q := float64(z)/ 2
	if z%2 == 0 {
		return q, true
	}
	return q, false

}
*/
