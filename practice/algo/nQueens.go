package algo

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	ds "enkya.org/playground/practice/datastructures"
	"enkya.org/playground/practice/io"
	"enkya.org/playground/utils"
)

type NQueens struct {
	description string
	examples    []io.IO
	testData    []io.IO
	results     [][]string
	versions    []func(n int) [][]string
}

// Runtime complexity: O(n^2N)
func (nq *NQueens) nQueensV1(_ int) [][]string {
	// TODO: Implement brute force solution
	return [][]string{}
}

func (nq *NQueens) nQueensV2(n int) [][]string {
	emptyBoard := make([][]string, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			emptyBoard[i] = append(emptyBoard[i], ".")
		}
	}

	cols, diags, antidiags := ds.NewOrderedSet(), ds.NewOrderedSet(), ds.NewOrderedSet()

	nq.backtrack(0, cols, diags, antidiags, emptyBoard)

	return nq.results
}

func (nq *NQueens) backtrack(row int, cols, diags, antidiags *ds.OrderedSet, state [][]string) {
	n := len(state)

	if row == n {
		nq.results = append(nq.results, nq.createBoard(state))
		return
	}

	for col := 0; col < n; col++ {
		currDiag := row - col
		currAntiDiag := row + col

		if cols.Contains(col) ||
			diags.Contains(currDiag) ||
			antidiags.Contains(currAntiDiag) {
			continue
		}

		cols.Add(col)
		diags.Add(currDiag)
		antidiags.Add(currAntiDiag)

		state[row][col] = "Q"

		nq.backtrack(row+1, cols, diags, antidiags, state)

		cols.Remove(col)
		diags.Remove(currDiag)
		antidiags.Remove(currAntiDiag)

		state[row][col] = "."
	}
}

func (nq *NQueens) createBoard(state [][]string) []string {
	board := make([]string, 0, len(state))

	for _, v := range state {
		s := strings.Join(v, "")
		board = append(board, s)
	}

	return board
}

func (nq *NQueens) RunAlgo() {
	if err := nq.Test(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (nq *NQueens) Test() error {
	for _, v := range nq.versions {
		if err := nq.testFunction(v); err != nil {
			return err
		}
	}

	return nil
}

func (nq *NQueens) testFunction(f func(n int) [][]string) error {
	functionNameParts := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	functionName := functionNameParts[len(functionNameParts)-1]
	fmt.Println("Function name:", functionName)

	for _, e := range nq.testData {
		s, _ := e.Input.(int)
		expected, _ := e.Output.([][]string)
		result := f(s)
		nq.results = [][]string{}

		for i, r := range result {
			if !utils.CompareSlice(r, expected[i]) {
				return fmt.Errorf("in %s for input %v: \nexpected: \n\t%v, \ngot: \n\t%v", functionName, e.Input, expected, result)
			}
		}
	}

	return nil
}

func (nq *NQueens) Describe() {
	fmt.Printf("\nDescription: %s\n", nq.description)
	fmt.Println("Examples:")

	for _, e := range nq.examples {
		fmt.Printf("\tInput: %v\n\tOutput: %v\n", e.Input, e.Output)
	}
}

func NewNQueens() *NQueens {
	nq := &NQueens{
		description: "Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.",
		examples: []io.IO{
			{
				Input: 4,
				Output: [][]string{
					{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."},
				},
			},
			{Input: 1, Output: [][]string{{"Q"}}},
		},
		results: [][]string{},
		testData: []io.IO{
			{Input: 4, Output: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
			{Input: 1, Output: [][]string{{"Q"}}},
		},
	}

	nq.versions = []func(n int) [][]string{
		nq.nQueensV1,
		nq.nQueensV2,
	}

	return nq
}
