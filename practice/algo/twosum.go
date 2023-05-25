package algo

import (
	"fmt"
)

type TwoSum struct {
	description string
	examples    []IO
	testData    []IO
	versions    []func(nums []int, target int) []int
}

func (ts *TwoSum) RunAlgo() {
	fmt.Println("Running TwoSum algo....")
	defer fmt.Println("Finished running TwoSum algo....")

	if err := ts.Test(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (ts *TwoSum) twoSumV1(nums []int, target int) []int {
	// Brute force
	for i, e := range nums {
		for j, f := range nums {
			if i != j && e+f == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// twoSumV2 uses a map to store the difference between the target and the current element.
func (ts *TwoSum) twoSumV2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, e := range nums {
		if j, ok := m[target-e]; ok {
			return []int{j, i}
		}
		m[e] = i
	}
	return []int{}
}

func (ts *TwoSum) Test() error {
	for _, v := range ts.versions {
		if err := ts.TestFunction(v); err != nil {
			return err
		}
	}

	return nil
}

func (ts *TwoSum) TestFunction(f func(nums []int, target int) []int) error {
	for _, e := range ts.testData {
		nums, _ := e.input.([]any)[0].([]int)
		target, _ := e.input.([]any)[1].(int)
		expected, _ := e.output.([]int)
		actual := f(nums, target)

		if !compareIntSlice(expected, actual) {
			return fmt.Errorf("expected %v, got %v", expected, actual)
		}
	}

	return nil
}

func compareIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, e := range a {
		if e != b[i] {
			return false
		}
	}

	return true
}

func (ts *TwoSum) Describe() {
	fmt.Printf("TwoSum: %s\n", ts.description)

	for i, e := range ts.examples {
		nums, _ := e.input.([]any)[0].([]int)
		target, _ := e.input.([]any)[1].(int)
		fmt.Printf("Example %d:\n\tInput: \tnums = %v, target = %d\n\tOutput:  %v\n", i, nums, target, e.output)
	}
}

func (ts *TwoSum) CastInput() error {
	for _, e := range ts.testData {
		nums, _ := e.input.([]any)[0].([]int)
		target, _ := e.input.([]any)[1].(int)
		e.input = []any{nums, target}
	}

	return nil
}

func NewTwoSum() *TwoSum {
	t := &TwoSum{
		description: `Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

		You may assume that each input would have exactly one solution, and you may not use the same element twice.
		
		You can return the answer in any order.`,
		examples: []IO{
			{
				input:  []any{[]int{2, 7, 11, 15}, 9},
				output: []int{0, 1},
			},
			{
				input:  []any{[]int{3, 2, 4}, 6},
				output: []int{1, 2},
			},
			{
				input:  []any{[]int{3, 3}, 6},
				output: []int{0, 1},
			},
		},
		testData: []IO{
			{
				input:  []any{[]int{2, 7, 11, 15}, 9},
				output: []int{0, 1},
			},
			{
				input:  []any{[]int{3, 2, 4}, 6},
				output: []int{1, 2},
			},
			{
				input:  []any{[]int{3, 3}, 6},
				output: []int{0, 1},
			},
		},
		versions: []func(nums []int, target int) []int{},
	}
	t.versions = append(t.versions, t.twoSumV1, t.twoSumV2)
	return t
}
