package algo

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"enkya.org/playground/internal/practice/io"
)

type BinarySearch struct {
	description string
	examples    []io.IO
	testData    []io.IO
	versions    []func([]int, int) int
}

func (bs *BinarySearch) binarySearchV1(a []int, t int) int {
	if len(a) == 0 {
		return -1
	}

	if len(a) == 1 {
		return 0
	}

	l, r := 0, len(a)-1

	for l <= r {
		m := l + (r-l)/2

		switch {
		case t == a[m]:
			return m
		case t > a[m]:
			l = m + 1
		default:
			r = m - 1
		}
	}

	return -1
}

func (bs *BinarySearch) RunAlgo() {
	if err := bs.Test(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (bs *BinarySearch) Test() error {
	for i, v := range bs.versions {
		fmt.Println("Testing version", i+1)

		if err := bs.testFunction(v); err != nil {
			return err
		}
	}

	return nil
}

func (bs *BinarySearch) testFunction(f func([]int, int) int) error {
	functionNameParts := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	functionName := functionNameParts[len(functionNameParts)-1]
	fmt.Println("Function name:", functionName)

	for _, e := range bs.testData {
		nums, _ := e.Input.([]any)[0].([]int)
		target, _ := e.Input.([]any)[1].(int)
		expected, _ := e.Output.(int)
		result := f(nums, target)

		if result != expected {
			return fmt.Errorf("in %s for input %v: \n\texpected %v, got %v", functionName, e.Input, expected, result)
		}
	}

	return nil
}

func (bs *BinarySearch) Describe() {
	fmt.Printf("\nDescription: %s\n", bs.description)
	fmt.Println("Examples:")

	for _, e := range bs.examples {
		fmt.Printf("Input: %v\nOutput: %v\n", e.Input, e.Output)
	}
}

func (bs *BinarySearch) LoadTestData() {
	bs.testData = []io.IO{
		{
			Input:  []any{[]int{}, 3},
			Output: -1,
		},
		{
			Input:  []any{[]int{1, 3, 4, 7, 9}, 3},
			Output: 1,
		},
		{
			Input:  []any{[]int{1, 3, 4, 7, 9}, 7},
			Output: 3,
		},
	}
}

func NewBinarySearch() *BinarySearch {
	bs := &BinarySearch{
		description: "find target number position given sorted list of integers",
		examples: []io.IO{
			{
				Input:  []any{[]int{1, 3, 4, 7, 9}, 3},
				Output: 1,
			},
			{
				Input:  []any{[]int{1, 3, 4, 7, 9}, 8},
				Output: -1,
			},
		},
		versions: []func([]int, int) int{},
	}
	bs.versions = append(bs.versions, bs.binarySearchV1)

	return bs
}
