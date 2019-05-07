package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type person1 struct {
	First string
	Last string
}


func main() {
	p1 := person1 {
		First: "James",
		Last: "Bond",
	}

	p2 := person1 {
		First: "Klint",
		Last: "Drone",
	}

	people := []person1{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	xi := []int{4,7,3,42,99}
	xs := []string{"Ebuka", "Joy", "Chris"}

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi, xs)
	
}