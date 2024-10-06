//go:build !codeanalysis
// +build !codeanalysis

package main

import (
	"enkya.org/playground/playground"
)

func main() {
	// practice.StartPractice()
	playground.NewGame().Start()
}
