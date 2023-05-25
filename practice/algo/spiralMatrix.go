package algo

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"enkya.org/playground/utils"
)

type SpiralMatrix struct {
	description string
	examples    []IO
	testData    []IO
	versions    []func(matrix [][]int) []int
}

func (sm *SpiralMatrix) Describe() {
	fmt.Printf("\nDescription: %s\n", sm.description)
	fmt.Println("Examples:")

	for _, e := range sm.examples {
		fmt.Printf("Input: %v\nOutput: %v\n", e.input, e.output)
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
	for i, v := range sm.versions {
		fmt.Println("Testing version", i+1)

		if err := sm.testFunction(v); err != nil {
			return err
		}
	}

	return nil
}

func (sm *SpiralMatrix) testFunction(f func(m [][]int) []int) error {
	functionNameParts := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	functionName := functionNameParts[len(functionNameParts)-1]
	fmt.Println("Function name:", functionName)

	for _, e := range sm.testData {
		matrix, _ := e.input.([][]int)
		expected, _ := e.output.([]int)
		r := f(matrix)

		if !utils.CompareIntSlice(r, expected) {
			return fmt.Errorf("in %s for input %v: \n\texpected %v, got %v", functionName, e.input, expected, r)
		}
	}

	return nil
}

func NewSpiralMatrix() *SpiralMatrix {
	m := &SpiralMatrix{
		description: `Given an m x n matrix, return all elements of the matrix in spiral order.`,
		examples: []IO{
			{
				[][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25},
				},
				[]int{1, 2, 3, 4, 5, 10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13},
			},
			{
				[][]int{{1, 2}, {3, 4}},
				[]int{1, 2, 4, 3},
			},
		},
		testData: []IO{
			{
				[][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25},
				},
				[]int{1, 2, 3, 4, 5, 10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13},
			},
		},
		versions: []func([][]int) []int{},
	}

	m.versions = append(m.versions, m.spiralMatrixV1)

	return m
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
