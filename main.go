//go:build !codeanalysis
// +build !codeanalysis

package main

import (
	"fmt"
	"sync"
	"time"

	"enkya.org/playground/practice"
	_ "enkya.org/playground/practice"
)

func main() {
	practice.StartPractice()
	// goByExample()
	// proto()
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

func variadic(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}

func pointers() {
	x := 7
	y := &x
	z := x
	fmt.Println(x, *y)
	*y = 8
	fmt.Println(x, *y)
	fmt.Println(z)
}

func testDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func testGenerics() {
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	// fmt.Println(getKeysNotNum(m, 2))
	fmt.Println(getKeysNotNum[int, string](m, 2))
}

func getKeysNotNum[K comparable, V any](m map[K]V, n K) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		if k == n {
			continue
		}
		keys = append(keys, k)
	}

	return keys
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
