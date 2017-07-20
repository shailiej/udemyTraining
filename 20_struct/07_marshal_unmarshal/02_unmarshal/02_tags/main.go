package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p1 person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`)
	//if "Age" was passed instead of "wisdom score", Age would default to 0 because it'd be looking for winsdom scrore & not Age
	json.Unmarshal(bs, &p1)

	fmt.Println("--------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Println(p1)
	fmt.Printf("%T \n", p1)
}
