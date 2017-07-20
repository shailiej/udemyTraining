package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}

	p3 := person{
		first: "shailie",
		last:  "j",
		age:   16,
	}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(p3)
	// fmt.Printf("%T %v \n", p1, p1)
	//fmt.Println(p1)
	//fmt.Println(p2)

}
