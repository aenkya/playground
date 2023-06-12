package algo

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

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
func (nq *NQueens) nQueensV1(n int) [][]string {
	// pseudo code
	// create a board of size n x n
	// place a queen in the first row
	// check if the queen is safe
	// if the queen is not safe, move it to the next column
	// if the queen is safe, place the next queen in the next row
	// continue this until the queen is safe
	// if the queen is safe, place the next queen in the next row
	// repeat this until all the queens are placed
	// if all the queens are placed, return the board
	return [][]string{}
}

func (nq *NQueens) nQueensV2(n int) [][]string {
	emptyBoard := make([][]string, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			emptyBoard[i] = append(emptyBoard[i], ".")
		}
	}

	nq.backtrack(0, make(map[int]bool, 0), make(map[int]bool, 0), make(map[int]bool), emptyBoard)

	return nq.results
}

func (nq *NQueens) backtrack(row int, cols, diags, antidiags map[int]bool, state [][]string) {
	n := len(state)

	if row == n {
		results := nq.results
		nq.results = append(results, nq.createBoard(state))
	}

	for col := 0; col < n; col++ {
		currDiag := row - col
		currAntiDiag := row + col
		if utils.Contains(cols, col) || utils.Contains(diags, currDiag) || utils.Contains(antidiags, currAntiDiag) {
			continue
		}

		cols[col] = true
		diags[currDiag] = true
		antidiags[currDiag] = true
		state[row][col] = "Q"

		nq.backtrack(row+1, cols, diags, antidiags, state)

		delete(cols, col)
		delete(diags, currDiag)
		delete(antidiags, currAntiDiag)
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
