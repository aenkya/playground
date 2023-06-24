package io

type IO struct {
	Input  interface{}
	Output interface{}
}

type Algo interface {
	LoadTestData()
	RunAlgo()
	Describe()
}
