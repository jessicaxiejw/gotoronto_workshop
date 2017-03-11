package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	jsonData, err := ioutil.ReadFile("./people.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	people := make([]Person, 0)
	err = json.Unmarshal(jsonData, &people)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people)
}
