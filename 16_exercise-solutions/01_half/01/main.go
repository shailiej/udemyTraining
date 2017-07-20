package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
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

func boolie(z int) (int, bool) {
	q := z / 2
	if z%2 == 0 {
		return q, true
	}
	return q, false

}
*/
