package main

import (
	"fmt"
)

type user struct {
	FirstName, LastName string
}

type customer struct {
	id       int
	fullName string
}

func (u *user) name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (c *customer) name() string {
	return c.fullName
}

type namer interface {
	name() string
}

func greet(n namer) string {
	return fmt.Sprintf("Dear %s", n.name())
}

func main() {
	u := &user{"Matt", "Aimonetti"}
	fmt.Println(greet(u))
	c := &customer{42, "Francesc"}
	fmt.Println(greet(c))
}
