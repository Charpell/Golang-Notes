package main

import "fmt"

// Array
// func main() {
// 	var x [5]int
// 	fmt.Println(x)

// 	x[4] = 42
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// }


// Composite literals
// func main() {
// 	// x := type{values}  // Composite Literals
	
// 	x  := []int{4, 5, 7, 8, 42}
// 	fmt.Println(x)

// 	// a SLICE allows you to group together VALUES of the same TYPE
// }



// Slice Range
// func main() {
// 	x := []int{4, 5, 7, 7, 42}

// 	for i, v := range x {
// 		fmt.Println(i, v)
// 	}
// }




// Slicing a Slice
// func main() {
// 	x := []int{4, 5, 7, 7, 42}	

// 	fmt.Println(x[1:])
// 	fmt.Println(x[1:3])


// 	for i := 0; i <= 4; i++ {
// 		fmt.Println(x[i])
// 	}
// }



// Append and Delete

// func main() {
// 	x := []int{4, 5, 7, 7, 42}
// 	fmt.Println(x)

// 	x = append(x, 77, 78, 99, 1014)
// 	fmt.Println(x)


// 	y := []int{43, 15, 27, 37, 42}

// 	x = append(x, y...)
// 	fmt.Println(x)

// 	x = append(x[:2], x[4:]...)
// 	fmt.Println(x)
// }





// Make
// func main() {
// 	x := make([]int, 10, 12)
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	x[0] = 42
// 	x[9] = 999
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	x = append(x, 3423)
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	x = append(x, 3423)
// 	x = append(x, 3423)
// 	x = append(x, 3423)
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))
// }



// Multidimensional Slice
// func main() {
// 	jb := []string{"James", "Bond", "Chocolate", "martini"}
// 	fmt.Println(jb)

// 	mp := []string{"Miss", "Money", "Straw", "Haz"}
// 	fmt.Println(mp)

// 	xp := [][]string{jb, mp}
// 	fmt.Println(xp) 
// }



// Maps
// func main() {
// 	m := map[string]int{
// 		"James": 32,
// 		"Miss Moneypenny": 27,
// 	}
// 	fmt.Println(m)

// 	fmt.Println(m["James"])
// 	fmt.Println(m["Barnabas"])  // value will be 0 because it does not exist

// 	// add 
// 	m["todd"] = 33

// 	// Check if value exist
// 	if v, ok := m["Miss Moneypenny"]; ok {
// 		fmt.Println("THIS IS THE IF PRINT", v)
// 	}

// 	// loop
// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}


// 	// delete
// 	delete(m, "James")
// 	fmt.Println(m)

// 	// Check if value exist
// 	if v, ok := m["Miss Moneypenny"]; ok {
// 		fmt.Println("value", v)
// 		delete(m, "Miss Moneypenny")
// 	}

// 	fmt.Println(m)
// }




// Exercice 1
// func main() {
// 	x := [5]int{42, 43, 44, 45, 46}

// 	for i, v := range x {
// 		fmt.Println(i, v)
// 	}
// 	fmt.Printf("%T\n", x)
// }


// Exercise 2
// func main() {
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 	for i, v := range x {
// 		fmt.Println(i, v)
// 	}
// 	fmt.Printf("%T\n", x)
// }


// [42 43 44 45 46]
// [47 48 49 50 51]
// [44 45 46 47 48]
// [43 44 45 46 47]


// Exercise 3
// func main() {
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 	fmt.Println(x[:5])
// 	fmt.Println(x[5:])
// 	fmt.Println(x[2:7])
// 	fmt.Println(x[1:6])
// 	fmt.Println(x)
// }


// Exercise 4
// func main() {
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// 	x = append(x, 52)
// 	fmt.Println(x)
// 	x = append(x, 53, 54, 55)
// 	fmt.Println(x)

// 	y := []int{56, 57, 58, 59, 60}
// 	x = append(x, y...)
// 	fmt.Println(x)
// }


// Exercise 5
// func main() {
// 	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	

// 	x = append(x[:3], x[6:]...)
// 	fmt.Println(x)
// }


// Exercise 6
// func main() {
// 	x := make([]string, 50, 50)

// 	x = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	for i := 0; i < len(x); i++ {
// 		fmt.Println(i, x[i])
// 	}
// }



// Exercise 7
// func main() {
// 	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
// 	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
// 	fmt.Println(xs1)
// 	fmt.Println(xs2)

// 	xxs := [][]string{xs1, xs2}
// 	fmt.Println(xxs)

// 	for i, xs := range xxs {
// 		fmt.Println("record: ", i)
// 		for j, val := range xs {
// 			fmt.Printf("\t index position: %v \t value: \t %v", j, val)
// 		}
// 	}
// }




// Exercise 8
// func main() {
// 	m := map[string][]string{
// 		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
// 		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
// 		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
// 	}

// 	// fmt.Println(m)

// 	for k, v := range m {
// 		fmt.Println("This is the record for", k)
// 		for i, v2 := range v {
// 			fmt.Println("\t", i, v2)
// 		}
// 	}
// }


// Exercise 9
// func main() {
// 	m := map[string][]string{
// 		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
// 		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
// 		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
// 	}

// 	// fmt.Println(m)

// 	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

// 	for k, v := range m {
// 		fmt.Println("This is the record for", k)
// 		for i, v2 := range v {
// 			fmt.Println("\t", i, v2)
// 		}
// 	}


	// Exercise 10
	func main() {
		m := map[string][]string{
			`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
			`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
			`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
		}
	
		// fmt.Println(m)
	
		m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	
		delete(m, `no_dr`)
	
		for k, v := range m {
			fmt.Println("This is the record for", k)
			for i, v2 := range v {
				fmt.Println("\t", i, v2)
			}
		}
	}
	