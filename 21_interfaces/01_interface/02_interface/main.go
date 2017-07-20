package main

import "fmt"

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	fmt.Printf("%T\n", s)
	info(s)
}

/* "type implements an interface"
object struct (square) is implementing interface (shape)
because interface has defined same method (area())
which is attached to struct ie, area() is a member of struct */
