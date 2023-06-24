//go:build ignore
// +build ignore

package main

import "strings"

// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.

// Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.

// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

// Input: n = 4
// Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above

func NQueens(n int) [][]string {
	// Brute force: O(N^N)

	// optimized: O(n!)

	board := createBoard(n)
	results := make([][]string, n)
	cols, antidiags, diags := make(map[int]bool), make(map[int]bool), make(map[int]bool)

	return recursiveCall(results, board, 0, cols, diags, antidiags)
}

func recursiveCall(results, board [][]string, row int, cols, diags, antidiags map[int]bool) [][]string {
	n := len(board)
	if row == n-1 {
		results = append(results, constructResponse(board))
		return results
	}

	for col := 0; col < n; col++ {
		if ok := isValidPlacement(row, col, cols, diags, antidiags); ok {
			cols[col] = true
			diags[row-col] = true
			antidiags[row+col] = true
			board[row][col] = "Q"

			results = recursiveCall(results, board, row+1, cols, diags, antidiags)

			delete(cols, col)
			delete(diags, row-col)
			delete(antidiags, row+col)
			board[row][col] = "."

		}
	}

	return results
}

func isValidPlacement(row, col int, cols, diags, antidiags map[int]bool) bool {
	currDiag := row - col
	currAntiDiag := row + col
	if contains(cols, col) || contains(antidiags, currAntiDiag) || contains(diags, currDiag) {
		return false
	}

	return true
}

func contains(set map[int]bool, val int) bool {
	if _, ok := set[val]; ok {
		return true
	}

	return false
}

func createBoard(n int) [][]string {
	board := make([][]string, n)

	for i := 0; i < n; i++ {
		board[i] = make([]string, n)

		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	return board
}

func constructResponse(board [][]string) []string {
	result := make([]string, len(board))
	for row := range board {
		result = append(result, strings.Join(board[row], ""))
	}

	return result
}
