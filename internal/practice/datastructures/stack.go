package datastructures

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}

	index := len(*s) - 1
	v := (*s)[index]
	*s = (*s)[:index]

	return v
}

func (s *Stack) Peek() interface{} {
	if len(*s) == 0 {
		return nil
	}

	return (*s)[len(*s)-1]
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
