//go:build !codeanalysis
// +build !codeanalysis

package practice

import (
	"enkya.org/playground/practice/algo"
)

func Practice() {
	ap := algo.Practice()

	a, err := ap.GetAlgo("threesum")
	if err != nil {
		panic(err)
	}

	a.RunAlgo()
}
