//go:build !codeanalysis
// +build !codeanalysis

package practice

import (
	"enkya.org/playground/practice/algo"
)

func Practice() {
	ap := algo.Practice()
	ap.SetChallengeName("ReverseArray")
	ap.RunAlgo()
}
