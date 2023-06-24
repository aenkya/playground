package algo

import (
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"enkya.org/playground/practice/io"
)

type RepeatedWords struct {
	description string
	examples    []io.IO
	testData    []io.IO
	versions    []func(string) int
}

func removeNonAlphanumericRegex(s string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9 ]+")
	return re.ReplaceAllString(s, "")
}

func (r *RepeatedWords) numOfRepeatedWordsv1(s string) int {
	s = strings.ToLower(s)
	s = removeNonAlphanumericRegex(s)
	words := strings.Split(s, " ")
	wordsMap := make(map[string]int)
	totalNumOfRepeatedWords := 0

	for _, w := range words {
		wordsMap[w]++
	}

	for _, v := range wordsMap {
		if v > 1 {
			totalNumOfRepeatedWords = totalNumOfRepeatedWords + v - 1
		}
	}

	return totalNumOfRepeatedWords
}

func (r *RepeatedWords) RunAlgo() {
	if err := r.Test(); err != nil {
		panic(err)
	}
}

func (r *RepeatedWords) Test() error {
	for _, v := range r.versions {
		if err := r.testFunction(v); err != nil {
			return err
		}
	}

	return nil
}

func (r *RepeatedWords) testFunction(f func(string) int) error {
	functionNameParts := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	functionName := functionNameParts[len(functionNameParts)-1]
	fmt.Println("Function name:", functionName)

	for _, e := range r.testData {
		s, _ := e.Input.(string)
		expected, _ := e.Output.(int)
		result := f(s)

		if result != expected {
			return fmt.Errorf("in %s for input %v: \n\texpected %v, got %v", functionName, e.Input, expected, result)
		}
	}

	return nil
}

func (r *RepeatedWords) Describe() {
	fmt.Printf("Repeated Sum: %s\n", r.description)

	for i, e := range r.examples {
		str, _ := e.Input.(string)
		fmt.Printf("Example %d:\n\tInput: \tstring = %v\n\tOutput:  %v\n", i, str, e.Output)
	}
}

func (r *RepeatedWords) LoadTestData() {
	r.testData = []io.IO{
		{
			Input:  "This is a test string. This is a test string.",
			Output: 5,
		},
		{
			Input:  "This is a test string. This is a test string. This is a test string.",
			Output: 10,
		},
		{
			Input:  "This is a test string. This is a test string. This is a test string. This is a test string.",
			Output: 14,
		},
	}
}

func NewRepeatedWords() *RepeatedWords {
	t := &RepeatedWords{
		description: "Given a string, find the number of repeated words in it.",
		examples: []io.IO{
			{Input: "This is a test string. This is a test string.", Output: 5},
			{Input: "This is a test string. This is a test string. This is a test string.", Output: 10},
			{Input: "This is a test string. This is a test string. This is a test string. This is a test string.", Output: 14},
		},
		versions: []func(string) int{},
	}
	t.versions = append(t.versions, t.numOfRepeatedWordsv1)

	return t
}
