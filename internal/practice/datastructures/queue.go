package datastructures

type Queue []interface{}

func (q *Queue) Enqueue(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() interface{} {
	if len(*q) == 0 {
		return nil
	}

	v := (*q)[0]
	*q = (*q)[1:]

	return v
}

func (q *Queue) Peek() interface{} {
	if len(*q) == 0 {
		return nil
	}

	return (*q)[0]
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func NewQueue() *Queue {
	return &Queue{}
}
