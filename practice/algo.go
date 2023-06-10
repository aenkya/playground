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

func (ap *Type) SetAlgo(name string, a io.Algo) {
	ap.algorithms[name] = a
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

	ts := algo.NewTwoSum()
	p.SetAlgo("twosum", ts)

	ts3 := algo.NewThreeSum()
	p.SetAlgo("threesum", ts3)

	isValidPalindrome := algo.NewIsValidPalindrome()
	p.SetAlgo("isValidPalindrome", isValidPalindrome)

	spiralMatrix := algo.NewSpiralMatrix()
	p.SetAlgo("spiralMatrix", spiralMatrix)

	bs := algo.NewBinarySearch()
	p.SetAlgo("binarySearch", bs)

	ms := sort.NewMergeSort()
	p.SetAlgo("mergeSort", ms)

	rw := algo.NewRepeatedWords()
	p.SetAlgo("repeatedWords", rw)

	nq := algo.NewNQueens()
	p.SetAlgo("nQueens", nq)

	return p
}
