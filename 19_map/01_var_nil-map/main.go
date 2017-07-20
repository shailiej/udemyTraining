package main

import "fmt"

func main() {

	var myGreeting map[string]string
	//  var myGreeting = make(map[string]string) // or
	//	var myGreeting = map[string]string{} // this map won't be nil but empty & you can add in stuff to it
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)

	// add these lines:
	/*
		myGreeting["Tim"] = "Good morning."
		myGreeting["Jenny"] = "Bonjour."
		fmt.Println(myGreeting)
	*/
	// and you will get this:
	// panic: assignment to entry in nil map
}
