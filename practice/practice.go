//go:build !codeanalysis
// +build !codeanalysis

package practice

import "enkya.org/playground/practice/io"

type Type struct {
	algorithms    map[string]io.Algo
	challengeName string
}

func StartPractice() {
	ap := AlgoPractice()

	a, err := ap.GetAlgo("spiralMatrix")
	if err != nil {
		panic(err)
	}

	a.RunAlgo()
}
