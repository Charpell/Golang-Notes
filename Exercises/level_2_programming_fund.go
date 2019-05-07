
package main

import "fmt"


// var x int8 = 127 

// func main()  {
// 	fmt.Println(x)
// 	fmt.Printf("%T\n", x)
// }



// OS and Architecture
// func main()  {
// 	fmt.Println(runtime.GOOS)
// 	fmt.Println(runtime.GOARCH)
// }



// String Type

// func main() {
// 	s := "Hello, playground"
// 	fmt.Println(s)
// 	fmt.Printf("%T\n", s)

// 	bs := []byte(s)
// 	fmt.Println(bs)
// 	fmt.Printf("%T\n", bs)

// // UTF 8
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%#U ", s[i])
// 	}


// 	for i, v := range s {
// 		fmt.Printf("at index position %d we have hex %#x\n", i, v)
// 	}
// }



// Const

// const a = 42
// const b = 42.78

// const (
// 	a = 34
// 	b = "Hello World"
// )

// const (
// 	a int = 42
// )

// func main() {
// 	fmt.Println(a)
// }



// Iota

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

// const (
// 	d = iota
// 	e
// 	f
// )

// func main() {
// 		fmt.Println(d)
// 		fmt.Println(e)
// 		fmt.Println(f)
// 	}



// Bit Shifting

// func main() {
// 	x := 4
// 	fmt.Printf("%d\t\t%b\n", x, x)

// 	y := x << 1
// 	fmt.Printf("%d\t\t%b\n", y, y)
// }


// Using Iota to create bit shift

// const (
// 	_ = iota
// 	kb = 1 << (iota * 10)
// 	mb = 1 << (iota * 10)
// 	gb = 1 << (iota * 10)
// )

// func main() {
// 	// kb := 1024
// 	// mb := 1024 * kb
// 	// gb := 1024 * mb

// 	fmt.Printf("%d\t\t\t%b\n", kb, kb)
// 	fmt.Printf("%d\t\t\t%b\n", mb, mb)
// 	fmt.Printf("%d\t\t\t%b\n", gb, gb)
// }




// Exercise 1
// func main() {
// 	a := 42
// 	fmt.Printf("%d\t%b\t%#x", a, a, a)
// }


// Exercise 2
// func main() {
// 	a := (42 == 42)
	
// 	fmt.Println(a)
// }


// Exercise 3
// const (
// 	a = 42
// 	b int = 43
// )

// func main() {
// 	fmt.Println(a, b)
// }


// Exercis 4
// func main() {
// 	a := 42
// 	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
// 	b := a << 1
// 	fmt.Printf("%d\t%b\t%#x", b, b, b)
// }


// Exercice 5
// func main() {
// 	a := `here is something
// 	as 
// 	a 
// 	raw string
// 	literal
// 	"you see"
// 	another thing`
// 	fmt.Println(a)
// }



// Exercise 6
const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
	d = 2017 + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}