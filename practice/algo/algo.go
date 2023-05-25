package algo

import "fmt"

type Algo interface {
	RunAlgo()
	Describe()
}

type PracticeType struct {
	algorithms    map[string]Algo
	challengeName string
}

type IO struct {
	input  interface{}
	output interface{}
}

func NewAlgoPractice() PracticeType {
	return PracticeType{
		algorithms: make(map[string]Algo),
	}
}

func (ap *PracticeType) RunAlgo() {
	fmt.Println("Running algo practice")
}

func (ap *PracticeType) SetAlgo(name string, algo Algo) {
	ap.algorithms[name] = algo
}

func (ap *PracticeType) GetAlgo(name string) (Algo, error) {
	if a := ap.algorithms[name]; a != nil {
		a.Describe()
		return a, nil
	}

	return nil, fmt.Errorf("algo %s not found", name)
}

func (ap *PracticeType) SetChallengeName(name string) {
	ap.challengeName = name
}

func Practice() PracticeType {
	p := NewAlgoPractice()
	p.SetChallengeName("leetcode")

	ts := NewTwoSum()
	p.SetAlgo("twosum", ts)

	ts3 := NewThreeSum()
	p.SetAlgo("threesum", ts3)

	return p
}
