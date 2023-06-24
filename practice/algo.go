package practice

import (
	"fmt"

	"enkya.org/playground/practice/algo"
	"enkya.org/playground/practice/io"
	"enkya.org/playground/practice/sort"
)

func NewAlgoPractice() Type {
	return Type{
		algorithms: make(map[string]io.Algo),
	}
}

func (ap *Type) RunAlgo() {
	fmt.Println("Running algo practice")
}

func (ap *Type) SetAlgo(name string, a io.Algo) *Type {
	ap.algorithms[name] = a
	return ap
}

func (ap *Type) GetAlgo(name string) (io.Algo, error) {
	if a := ap.algorithms[name]; a != nil {
		a.Describe()
		return a, nil
	}

	return nil, fmt.Errorf("algo %s not found", name)
}

func (ap *Type) SetChallengeName(name string) {
	ap.challengeName = name
}

func AlgoPractice() Type {
	p := NewAlgoPractice()
	p.SetChallengeName("leetcode")
	p.SetAlgo("longestsubstring", algo.NewLongestSubstring()).
		SetAlgo("twosum", algo.NewTwoSum()).
		SetAlgo("threesum", algo.NewThreeSum()).
		SetAlgo("isValidPalindrome", algo.NewIsValidPalindrome()).
		SetAlgo("spiralMatrix", algo.NewSpiralMatrix()).
		SetAlgo("binarySearch", algo.NewBinarySearch()).
		SetAlgo("mergeSort", sort.NewMergeSort()).
		SetAlgo("repeatedWords", algo.NewRepeatedWords()).
		SetAlgo("nQueens", algo.NewNQueens()).
		SetAlgo("median2sortedarrays", algo.NewMedian2SortedArrays())

	return p
}
