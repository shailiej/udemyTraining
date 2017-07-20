package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1) //go right to left to understand
	// xx := json.NewEncoder(os.Stdout) // xx is pointer to Encoder (*Encoder)
	// xx.Encode(p1) //p1 getting passed to Encode which writes it to the *Encoder, which then prints it to Stdout
}
