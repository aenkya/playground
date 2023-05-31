package io

type IO struct {
	Input  interface{}
	Output interface{}
}

type Algo interface {
	RunAlgo()
	Describe()
	Test() error
}
