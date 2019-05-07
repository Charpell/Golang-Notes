package main

import "fmt";

type person struct {
	first string
	last string
	age int
}

// Embedded Struct
type secretAgent struct {
	person
	first string
	ltk bool
}

func main() {
	p1 := person {
		first: "James",
		last: "Bond",
		age: 23,
	}

	p2 := person {
		first: "Miss",
		last: "New",
		age: 44,
	}


	sa1 := secretAgent {
		person: person{
			first: "Han",
			last: "Bond",
			age: 32,
		},
		first: "Something call",
		ltk: true,
	}

// Anonymous
	an1 := struct {
		first string
		last string
		age int
	} {
		first: "James",
		last: "Bond",
		age: 32,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(an1)

	fmt.Println(p1.first, p1.last)
	fmt.Println(sa1, sa1.person, sa1.person.first)


	// Exercise 1 

	type human struct {
		first string
		last string 
		favColor []string
	}

	h1 := human {
		first: "Ebuka",
		last: "Umeh",
		favColor: []string{ "Memo", "Stick", "Chocloate"},
	}

	fmt.Println(h1.first, h1.last)
	
	for i, v := range h1.favColor {
		fmt.Println(i, v)
	}

	// Creating map
	// m := map[string]person{}


	// Exercise 2
	m := map[string]person {
			p1.last: p1,
			p2.last: p2,
	}


	for k, v := range m {
		fmt.Println(k, v.first)
	}


	// Exercise 4
	sd := struct {
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Money": 555,
			"Q": 777,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	fmt.Println(sd.first)
	fmt.Println(sd.friends)
	fmt.Println(sd.favDrinks)

	for k, v := range sd.friends {
		fmt.Println(k, v)
	}

	for k, v := range sd.favDrinks {
		fmt.Println(k, v)
	}

}
