//go:build !codeanalysis
// +build !codeanalysis

package main

import (
	"fmt"
	"sync"
	"time"

	"enkya.org/playground/practice"
)

func main() {
	practice.StartPractice()
	goByExample()
}

func goByExample() {
	x := "hello"
	switch x {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default")
	}

	var a [5]int
	a[2] = 7
	fmt.Println(a[2], a, len(a))

	testGoroutines()
}

func testGoroutines() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func(going string) {
		fmt.Printf("hello %s\n", going)
		wg.Done()
	}("goroutine")

	wg.Wait()
	testChannels()
}

func testChannels() {
	done := make(chan bool)

	go func(done chan bool) {
		time.Sleep(1 * time.Second)
		done <- true
	}(done)

	<-done
	fmt.Println("done")
}
