package sort

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"enkya.org/playground/practice/io"
	"enkya.org/playground/utils"
)

type MergeSort struct {
	description string
	examples    []io.IO
	testData    []io.IO
	versions    []func([]int) []int
}

func (bs *MergeSort) MergeSortV1(a []int) []int {
	if len(a) == 0 || len(a) == 1 {
		return a
	}

	return recursiveMergeSort(a)
}

func recursiveMergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	mid := len(a) / 2
	left := recursiveMergeSort(a[:mid])
	right := recursiveMergeSort(a[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}

		if len(right) == 0 {
			return append(result, left...)
		}

		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func (bs *MergeSort) RunAlgo() {
	if err := bs.Test(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (bs *MergeSort) Test() error {
	for i, v := range bs.versions {
		fmt.Println("Testing version", i+1)

		if err := bs.testFunction(v); err != nil {
			return err
		}
	}

	return nil
}

func (bs *MergeSort) testFunction(f func([]int) []int) error {
	defer println("...testing complete")

	functionNameParts := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	functionName := functionNameParts[len(functionNameParts)-1]
	fmt.Printf("Function name:%s...\n", functionName)

	for _, e := range bs.testData {
		nums, _ := e.Input.([]int)
		expected, _ := e.Output.([]int)
		actual := f(nums)

		if !utils.CompareSlice(expected, actual) {
			return fmt.Errorf("expected %v, got %v", expected, actual)
		}
	}

	return nil
}

func (bs *MergeSort) Describe() {
	fmt.Printf("\nDescription: %s\n", bs.description)
	fmt.Println("Examples:")

	for _, e := range bs.examples {
		fmt.Printf("Input: %v\nOutput: %v\n", e.Input, e.Output)
	}
}

func NewMergeSort() *MergeSort {
	bs := &MergeSort{
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
		testData: []io.IO{
			{
				Input:  []int{},
				Output: []int{},
			},
			{
				Input:  []int{1, 8, 3, 4, 7, 9},
				Output: []int{1, 3, 4, 7, 8, 9},
			},
			{
				Input:  []int{4, 1, 0, 8, 3, 4, 7, 9},
				Output: []int{0, 1, 3, 4, 4, 7, 8, 9},
			},
		},
		versions: []func([]int) []int{},
	}
	bs.versions = append(bs.versions, bs.MergeSortV1)

	return bs
}
