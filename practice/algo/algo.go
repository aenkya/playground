package algo

import "fmt"

type PracticeType struct {
	challengeName string
	testData      []byte
}

func NewAlgoPractice() PracticeType {
	return PracticeType{}
}

func (ap *PracticeType) RunAlgo() {
	fmt.Println("Running algo practice")
}

func (ap *PracticeType) SetChallengeName(name string) {
	ap.challengeName = name
}

func (ap *PracticeType) SetTestData(data []byte) {
	ap.testData = data
}

func Practice() PracticeType {
	return NewAlgoPractice()
}
