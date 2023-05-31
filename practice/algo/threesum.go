package algo

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"strings"

	"enkya.org/playground/practice/io"
	"enkya.org/playground/utils"
)

type ThreeSum struct {
	description string
	examples    []io.IO
	testData    []io.IO
	versions    []func(nums []int) [][]int
}

func (ts *ThreeSum) RunAlgo() {
	fmt.Println("Running ThreeSum algo....")
	defer fmt.Println("Finished running ThreeSum algo....")

	if err := ts.Test(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (ts *ThreeSum) threeSumV1(nums []int) [][]int {
	// Brute force O(n^3)
	sort.Ints(nums)
	res := make([][]int, 0, len(nums))

	for i, e := range nums {
		if i != 0 && e == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			f := nums[j]

			if j != i && f == nums[j-1] {
				continue
			}

			for k := j + 1; k < len(nums); k++ {
				g := nums[k]

				if k != j+1 && g == nums[k-1] {
					continue
				}

				if e+f+g == 0 {
					res = append(res, []int{e, f, g})
				}
			}
		}
	}

	return res
}

// threeSumV2 uses a map to store the difference between the target and the current element.
func (ts *ThreeSum) threeSumV2(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	resMap := make(map[string][]int)
	m := make(map[int]int)

	sort.Ints(nums)

	for i, e := range nums {
		if i != 0 && e == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			f := nums[j]
			fmt.Println("i: ", i, "e: ", e, "j: ", j, "f: ", f)

			if k, ok := m[-e-f]; ok {
				if z := resMap[fmt.Sprintf("%d%d%d", e, f, k)]; z != nil {
					fmt.Printf("Found duplicate: %v\n", z)
					continue
				}

				resMap[fmt.Sprintf("%d%d%d", e, f, k)] = []int{e, f, k}
				res = append(res, []int{e, f, k})

				for j+1 < len(nums) && nums[j] == nums[j+1] {
					continue
				}
			}

			m[f] = f
		}

		m[e] = i
	}

	return res
}

// threeSumV3 uses the 2 pointer approach.
func (ts *ThreeSum) threeSumV3(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, len(nums))

	for i, e := range nums {
		if i != 0 && e == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			f, g := nums[l], nums[r]
			sum := e + f + g

			//nolint:gocritic // no switches
			if sum == 0 {
				res = append(res, []int{e, f, g})

				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

func (ts *ThreeSum) Test() error {
	for i, v := range ts.versions {
		fmt.Println("Testing version", i+1)

		if err := ts.testFunction(v); err != nil {
			return err
		}
	}

	return nil
}

func (ts *ThreeSum) testFunction(f func(nums []int) [][]int) error {
	functionNameParts := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	functionName := functionNameParts[len(functionNameParts)-1]
	fmt.Println("Function name:", functionName)

	for _, e := range ts.testData {
		nums, _ := e.Input.([]int)
		expected, _ := e.Output.([][]int)
		result := f(nums)

		for i, r := range result {
			if !utils.CompareIntSlice(r, expected[i]) {
				return fmt.Errorf("in %s for input %v: \n\texpected %v, got %v", functionName, e.Input, expected, result)
			}
		}
	}

	return nil
}

func (ts *ThreeSum) Describe() {
	fmt.Printf("\nDescription: %s\n", ts.description)
	fmt.Println("Examples:")

	for _, e := range ts.examples {
		fmt.Printf("Input: %v\nOutput: %v\n", e.Input, e.Output)
	}
}

func NewThreeSum() *ThreeSum {
	t := &ThreeSum{
		description: `Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

		Notice that the solution set must not contain duplicate triplets.`,
		examples: []io.IO{
			{
				Input:  []any{[]int{-1, 0, 1, 2, -1, -4}},
				Output: [][]int{{-1, 0, 1}, {-1, -1, 2}},
			},
			{
				Input:  []any{[]int{0, 1, 1}},
				Output: [][]int{},
			},
			{
				Input:  []any{[]int{0, 0, 0}},
				Output: [][]int{{0, 0, 0}},
			},
		},
		testData: []io.IO{
			{
				Input:  []int{-1, 0, 1, 2, -1, -4},
				Output: [][]int{{-1, 0, 1}, {-1, -1, 2}},
			},
			{
				Input:  []int{0, 1, 1},
				Output: [][]int{},
			},
			{
				Input:  []int{0, 0, 0},
				Output: [][]int{{0, 0, 0}},
			},
		},
		versions: []func(nums []int) [][]int{},
	}
	t.versions = append(t.versions, t.threeSumV1, t.threeSumV2, t.threeSumV3)

	return t
}
