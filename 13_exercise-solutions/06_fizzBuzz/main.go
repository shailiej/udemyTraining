package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, " -- FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, " -- FIZZ")
		} else if i%5 == 0 {
			fmt.Println(i, " -- BUZZ")
		} else {
			fmt.Println(i)
		}
	}
}

/* package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%15 != 0 {

			fmt.Println("Fizz")

		} else if i%5 == 0 && i%15 != 0 {
			fmt.Println("Buzz")

		} else if i%15 == 0 {

			fmt.Println("FizzBuzz")
		} else {

			fmt.Println(i)

		}

	}
}
*/
