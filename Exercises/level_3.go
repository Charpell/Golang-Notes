package main

import "fmt"


// 1 Write a program that prints a number in decimal, binary, and Hex

// func main() {
// 	a := 42
// 	fmt.Printf("%d\t%b\t%#x\n" , a, a, a)
// }


// 2 Using the following operators, write expressions and assign their values to variables
// func main() {
// 	a := (42 == 42)
// 	b := (42 <= 43)
// 	c := (42 >= 43)
// 	d := (42 != 43)
// 	e := (42 < 43)
// 	f := (42 > 43)
	
// 	fmt.Println(a, b, c, d, e, f)
// }


// 3 For Loops
// func main() {
// 	// for i := 0; i <= 10; i++ {
// 	// 	fmt.Println(i)
// 	// }

// 	// loop inside loop
// 	for i := 0; i <= 5; i++ {
// 		fmt.Printf("The outer loop is: %d\n", i)
// 		for j := 0; j < 3; j++ {
// 			fmt.Printf("\t\t The inner loop: %d\n", j)
// 		}
// 	}
// }


// Exercise 2

// func main() {
// 	for i := 65; i <= 90; i++ {
// 		fmt.Println(i)
// 		for j := 0; j < 3; j++ {
// 			fmt.Printf("\t%#U\n", j)
// 		}
// 	}
// }


// Exercise 3

// for condition

// func main() {
// 	bd := 1989
// 	for bd <= 2019 {
// 		fmt.Println(bd)
// 		bd++
// 	}
// }


// Exercise 4

// if condition
// func main() {
// 	bd := 1989
// 	for {
// 		if bd > 2019 {
// 			break
// 		}
// 		fmt.Println(bd)
// 		bd++
// 	}
// }



// Exercise 5
// func main() {
// 	for i := 10; i <= 100; i++ {
// 		fmt.Printf("The reminder of %d is %d\n", i, i % 4)
// 	}
// }


// Exercise 6

// if statement
// func main() {
// 	x := "James Bond"

// 	if x == "James Bond" {
// 		fmt.Println(x)
// 	}
// }


// Exercise 7
// func main() {
// 	x := "Moneypenny"

// 	if x == "Moneypenny" {
// 		fmt.Println(x)
// 	} else if x == "James Bond" {
// 		fmt.Println("BONDDONBONDONBOND", x)
// 	} else {
// 		fmt.Println("neither")
// 	}
// }



// Exercise 8
// func main() {
// 	switch {
// 	case false:
// 		fmt.Println("should not print")
// 	case true:
// 		fmt.Println("should print")
// 	}
// }


// Exetrcise 9
func main() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}