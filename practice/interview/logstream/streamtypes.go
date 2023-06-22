package logstream

type Stream interface {
	GetID() int
	GetTerms() []string
	GetTargets() []int
}

type Query struct {
	terms []string
	ID    int
}

func (q *Query) GetTerms() []string {
	return q.terms
}

func (q *Query) GetID() int {
	return q.ID
}

func (q *Query) GetTargets() []int {
	return []int{q.ID}
}

type Log struct {
	terms          []string
	matchedQueries []*Query
}

func (l *Log) GetID() int {
	return -1
}

func (l *Log) GetTerms() []string {
	return l.terms
}

func (l *Log) GetTargets() []int {
	targets := []int{}

	for _, q := range l.matchedQueries {
		targets = append(targets, q.ID)
	}

	return targets
}
