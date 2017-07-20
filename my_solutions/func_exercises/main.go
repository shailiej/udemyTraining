/*package main

import "fmt"

func main() {
	//ret, xs := boolie(5)
	//fmt.Println(ret, xs)

	funcExp := func(z int) (int, bool) {
		q := z / 2
		if z%2 == 0 {
			return q, true
		}
		return q, false
	}

	//	z, y := funcExp(5)
	//	fmt.Println(z, y)
	fmt.Println(funcExp(5))

}


package main

import "fmt"

func main() {
	//ret, xs := boolie(5)
	//fmt.Println(ret, xs)
	fmt.Println(boolie(5))
}

func boolie(z int) (float64, bool) {
	q := float64(z) / 2
	if z%2 == 0 {
		return q, true
	}
	return q, false

}
*/

package main

import "fmt"

func variadic(num ...int) int {

	var largest int

	for i, v := range num {
		if v > largest || i == 0 {
			largest = v
		}
	}
	return largest

}

func main() {

	greatest := variadic(-56, 54, 23, 67)
	fmt.Println(greatest)
}
