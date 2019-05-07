package main;

import (
	"fmt"
	"math"
)

func main() {
	// s := sum(2,3,4,5,6,7)
	xi := []int{2,3,4,5,6,7}
	s := sum(xi...)
	fmt.Println(s)

	st := woo("Moneypenny")
	fmt.Println(st)

	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)

	defer foo()
	bar()

	sa1 := stsoul {
		soul: soul{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa1.speak()

	so1 := soul {
		first: "Dr",
		last: "Yes",
	}
	
	praInter(sa1)
	praInter(so1)


	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)


	f1 := func(x int){
		fmt.Println("my first func expression", x)
	}

	f1(1984)


	re1 := returnFunc()
	fmt.Println(re1)

	fmt.Println(insideFunc()())

	inF := insideFunc()
	fmt.Printf("%T\n", inF)

	inFr := inF()
	fmt.Println(inFr)

	
	xxi := []int{1,2,3,4,5,6,7,8,9,}
	sum1 := newSum(xxi...)
	fmt.Println("all numbers", sum1)

	evenSum1 := even(sum, xxi...)
	fmt.Println("even numbers", evenSum1)
	
	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	n := factorial(4)
	fmt.Println(n)

	l := loopFactorial(4)
	fmt.Println(l)

	circ := circle{
		radius: 12.345,
	}

	squa := square {
		length: 15,
	}

	info(circ)
	info(squa)

}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
} 

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ",ln, `, says "Hello"`)
	b := false
	return a, b
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is,", sum)
	return sum
}


func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

type soul struct {
	first string
	last string 
}

type stsoul struct {
	soul
	ltk bool
}

func (s soul) speak() {
	fmt.Println("I am", s.first, s.last, " - soul speak")
}

func (s stsoul) speak() {
	fmt.Println("I am", s.first, s.last, " - the person speak")
}


// Interface
type human interface {
	speak()
}

func praInter(h human) {
	switch h.(type) {
	case soul:
		fmt.Println("I was passed into this function", h.(soul).first)
	case stsoul:
		fmt.Println("I was passed into this function", h.(stsoul).first)
	}
}


func returnFunc() string {
	s := "Hello world"
	return s
}

func insideFunc() func() int {
	return func() int{
		return 451
	}
}




func newSum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
		var yi []int
		for _, v := range vi {
			if v % 2 == 0 {
				yi = append(yi, v)
			}
		}
		return f(yi...)
}


func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}


func factorial(n int) int {
	if n == 0 {
		return 1
	}
	
	return n * factorial(n-1)
}

func loopFactorial(n int) int {
	total := 1
	for ; n > 0; n -- {
		total *= n
	}
	return total
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}