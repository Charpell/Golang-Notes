// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists


package main;

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		fmt.Println("Hello from one")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from two")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from three")
		wg.Done()
	}()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}