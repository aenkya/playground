package logstream

import (
	"fmt"
	"strings"

	pio "enkya.org/playground/practice/io"
)

//nolint:govet // ignore struct field order
type LogStream struct {
	description       string
	currentQueryIndex int
	responseBuilder   *ResponseBuilder
	examples          []pio.IO
	testData          []pio.IO
	store             map[string][]*Query
	results           [][]string
	versions          []func([]string) []string
}

func (ls *LogStream) logStreamV1(lsInput []string) []string {
	output := []string{}

	for _, input := range lsInput {
		if ls.isQuery(input) {
			output = append(output, ls.processQuery(input))
		} else {
			if r, err := ls.processLog(input); err == nil {
				fmt.Println("r: ", r)
				output = append(output, r)
			}
		}
	}

	return output
}

func (ls *LogStream) PrintStreamResponses(done chan bool) {
	for {
		select {
		case r := <-ls.responseBuilder.outputChannel:
			fmt.Println(r.String())
		case <-done:
			close(ls.responseBuilder.outputChannel)
			return
		}
	}
}

func (ls *LogStream) ReceiveInputFromChannel(instr chan string, done chan bool) {
	go func() {
		for {
			select {
			case <-instr:
				// process input
			case <-done:
				close(instr)
				return
			}
		}
	}()
}

func (ls *LogStream) isQuery(s string) bool {
	return strings.HasPrefix(s, "Q")
}

func (ls *LogStream) ParseQuery(qs string) *Query {
	qt := strings.Split(qs, " ")

	query := &Query{
		ID:    ls.currentQueryIndex + 1,
		terms: qt[1:],
	}

	ls.currentQueryIndex++

	return query
}

func (ls *LogStream) ParseLog(l string) *Log {
	lt := strings.Split(l, " ")

	log := &Log{
		terms:          lt[1:],
		matchedQueries: []*Query{},
	}

	return log
}

func (ls *LogStream) processLog(l string) (string, error) {
	log := ls.ParseLog(l)
	ls.matchQueriesToLog(log)

	if len(log.matchedQueries) > 0 {
		r := ls.createResponse(MATCH, log)
		// TODO: add response to output channel

		return r.String(), nil
	}

	return "", fmt.Errorf("no queries matched log: %s", l)
}

func (ls *LogStream) processQuery(q string) string {
	query := ls.ParseQuery(q)
	ls.addQueryToStore(query)
	r := ls.createResponse(ACK, query)
	// TODO: add response to output channel

	return r.String()
}

func (ls *LogStream) addQueryToStore(q *Query) {
	terms := []string{}

	for _, term := range q.terms {
		terms = append(terms, strings.ToLower(term))
	}

	s := strings.Join(terms, " ")

	if _, ok := ls.store[s]; !ok {
		ls.store[s] = []*Query{q}
	} else {
		ls.store[s] = append(ls.store[s], q)
	}
}

func (ls *LogStream) matchQueriesToLog(log *Log) {
	for queryTerms, queries := range ls.store {
		if ls.matchQueryToLog(log, queryTerms) {
			log.matchedQueries = append(log.matchedQueries, queries...)
		}
	}
}

func (ls *LogStream) matchQueryToLog(log *Log, terms string) bool {
	qterms := strings.Split(terms, " ")

	for _, term := range qterms {
		if !log.containsTerm(term) {
			return false
		}
	}

	return true
}

func (l *Log) containsTerm(term string) bool {
	for _, t := range l.terms {
		if strings.Contains(strings.ToLower(t), term) {
			return true
		}
	}

	return false
}

func (ls *LogStream) createResponse(msgType MessageType, so Stream) *Response {
	rb := ls.responseBuilder
	defer rb.Clear()

	rb.SetMessageType(msgType).
		SetMessage(strings.Join(so.GetTerms(), " ")).
		SetTargets(so.GetTargets())

	return rb.Build()
}

func (ls *LogStream) RunAlgo() {
	for _, version := range ls.versions {
		for _, data := range ls.testData {
			input, _ := data.Input.([]string)
			ls.results = append(ls.results, version(input))
		}
	}
}

func (ls *LogStream) LoadTestData() {
	ls.testData = ls.examples
}

func (ls *LogStream) Describe() {
	fmt.Printf("\nDescription: %s\n", ls.description)
	fmt.Println("Examples:")

	for _, e := range ls.examples {
		fmt.Printf("Input: %v\nOutput: %t\n", e.Input, e.Output)
	}
}

func NewLogStream() *LogStream {
	ls := &LogStream{
		description: "LogStream",
		examples: []pio.IO{
			{
				Input: []string{
					"Q: create database",
					"Q: close database",
					"Q: create model",
					"L: Database created",
					"L: model created",
					"Q: close connection",
					"L: connection closed",
					"Q: Database closed",
				},
				Output: []string{
					"ACK: create database, 1",
					"ACK: close database, 2",
					"ACK: create model, 3",
					"M: Database created, 1",
					"M: model created, 3",
					"ACK: close connection, 4",
					"M: connection closed, 4",
					"ACK: Database closed, 5",
				},
			},
		},
		responseBuilder: NewResponseBuilder(),
		store:           make(map[string][]*Query),
	}

	ls.versions = []func([]string) []string{
		ls.logStreamV1,
	}

	return ls
}
