package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 person
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`) //NewReader reads string and puts it into a Reader (rdr here)
	json.NewDecoder(rdr).Decode(&p1)                                       //go left to right to understand
	// NewDecoder takes in a Reader and not string so we converted our json string into reader in previous step

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
