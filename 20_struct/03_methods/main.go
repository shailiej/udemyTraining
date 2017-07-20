package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string { //declaring person as receiver, makes fullName() one of it's members
	return p.first + p.last
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	fmt.Println(p1.first, p1.last, p1.fullName(), p1.age) //fullname() has become one of the members of struct person
}

/*
func fullName(p person) string {     //o/p is same but here fullName is just a func that passes struct person
	return p.first + p.last
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(fullName(p1))
	fmt.Println(fullName(p2))
}
*/
