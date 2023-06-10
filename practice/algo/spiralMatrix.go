package algo

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"enkya.org/playground/practice/io"
	"enkya.org/playground/utils"
)

type SpiralMatrix struct {
	description string
	examples    []io.IO
	testData    []io.IO
	versions    []func(matrix [][]int) []int
}

// uses recursion
func (sm *SpiralMatrix) spiralMatrixV1(m [][]int) []int {
	return recursiveSMG(m, 0, len(m)-1, 0, len(m[0])-1, []int{})
}

func recursiveSMG(m [][]int, starti, endi, startj, endj int, output []int) []int {
	if endi < starti {
		return output
	}

	if endj < startj {
		return output
	}

	if starti == endi {
		for j := startj; j < endj; j++ {
			output = append(output, m[starti][j])
		}

		return output
	}

	if startj == endj {
		for i := starti; i < endi; i++ {
			output = append(output, m[i][startj])
		}

		return output
	}

	for j := startj; j <= endj; j++ {
		output = append(output, m[starti][j])
	}

	for i := starti + 1; i <= endi; i++ {
		output = append(output, m[i][endj])
	}

	for j := endj - 1; j >= startj; j-- {
		output = append(output, m[endi][j])
	}

	for i := endi - 1; i > starti; i-- {
		output = append(output, m[i][startj])
	}

	return recursiveSMG(m, starti+1, endi-1, startj+1, endj-1, output)
}

func (sm *SpiralMatrix) spiralMatrixV2(m [][]int) []int {
	res := make([]int, 0, len(m)*len(m[0]))
	starti, endi, startj, endj := 0, len(m)-1, 0, len(m[0])-1

	if len(m) == 0 {
		return res
	}

	if endi < starti {
		return res
	}

	if endj < startj {
		return res
	}

	for starti <= endi {
		for j := startj; j <= endj; j++ {
			res = append(res, m[starti][j])
		}

		for i := starti + 1; i <= endi; i++ {
			res = append(res, m[i][endj])
		}

		if starti != endi {
			for j := endj - 1; j >= startj; j-- {
				res = append(res, m[endi][j])
			}
		}

		if startj != endj {
			for i := endi - 1; i > starti; i-- {
				res = append(res, m[i][startj])
			}
		}

		starti++
		endi--
		startj++
		endj--
	}

	return res
}

func (sm *SpiralMatrix) Describe() {
	fmt.Printf("\nDescription: %s\n", sm.description)
	fmt.Println("Examples:")

	for _, e := range sm.examples {
		fmt.Printf("Input: %v\nOutput: %v\n", e.Input, e.Output)
	}
}

func (sm *SpiralMatrix) RunAlgo() {
	fmt.Println("Running spiralMatrix algo....")
	defer fmt.Println("Finished running spiralMatrix algo....")

	if err := sm.Test(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (sm *SpiralMatrix) Test() error {
	for _, v := range sm.versions {
		if err := sm.testFunction(v); err != nil {
			return err
		}
	}

	return nil
}

func (sm *SpiralMatrix) testFunction(f func(m [][]int) []int) error {
	defer fmt.Println("..finished testing function")

	functionNameParts := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	functionName := functionNameParts[len(functionNameParts)-1]
	fmt.Println("Testing Function:", functionName)

	for _, e := range sm.testData {
		matrix, _ := e.Input.([][]int)
		expected, _ := e.Output.([]int)
		r := f(matrix)

		if !utils.CompareSlice(r, expected) {
			return fmt.Errorf("in %s for input %v: \n\texpected %v, got %v", functionName, e.Input, expected, r)
		}

		fmt.Printf(".")
	}

	return nil
}

func NewSpiralMatrix() *SpiralMatrix {
	m := &SpiralMatrix{
		description: `Given an m x n matrix, return all elements of the matrix in spiral order.`,
		examples: []io.IO{
			{
				Input: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25},
				},
				Output: []int{1, 2, 3, 4, 5, 10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13},
			},
			{
				Input:  [][]int{{1, 2}, {3, 4}},
				Output: []int{1, 2, 4, 3},
			},
		},
		testData: []io.IO{
			{
				Input: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25},
				},
				Output: []int{1, 2, 3, 4, 5, 10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13},
			},
		},
		versions: []func([][]int) []int{},
	}

	m.versions = append(m.versions, m.spiralMatrixV2, m.spiralMatrixV1)

	return m
}
