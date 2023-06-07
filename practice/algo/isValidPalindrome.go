package algo

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unicode"

	"enkya.org/playground/practice/io"
)

type IsValidPalindrome struct {
	description string
	examples    []io.IO
	testData    []io.IO
	versions    []func(s string) bool
}

func (vp *IsValidPalindrome) isValidPalindromeV1(s string) bool {
	s = strings.ToLower(s)
	s = removeNonAlphanumeric(s)

	if len(s) < 2 {
		return true
	}

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func removeNonAlphanumeric(s string) string {
	var b strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}

	return b.String()
}

func (vp *IsValidPalindrome) RunAlgo() {
	fmt.Println("Running IsValidPalindrome algo....")
	defer fmt.Println("Finished running IsValidPalindrome algo....")

	if err := vp.Test(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (vp *IsValidPalindrome) Describe() {
	fmt.Printf("\nDescription: %s\n", vp.description)
	fmt.Println("Examples:")

	for _, e := range vp.examples {
		fmt.Printf("Input: %v\nOutput: %t\n", e.Input, e.Output)
	}
}

func (vp *IsValidPalindrome) Test() error {
	for i, v := range vp.versions {
		fmt.Println("Testing version", i+1)

		if err := vp.testFunction(v); err != nil {
			return err
		}
	}

	return nil
}

func (vp *IsValidPalindrome) testFunction(f func(s string) bool) error {
	functionNameParts := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	functionName := functionNameParts[len(functionNameParts)-1]
	fmt.Println("Function name:", functionName)

	for _, e := range vp.testData {
		nums, _ := e.Input.(string)
		expected, _ := e.Output.(bool)
		result := f(nums)

		if result != expected {
			return fmt.Errorf("in %s for input %v: \n\texpected %v, got %v", functionName, e.Input, expected, result)
		}
	}

	return nil
}

func NewIsValidPalindrome() *IsValidPalindrome {
	p := &IsValidPalindrome{
		description: "Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.",
		examples: []io.IO{
			{Input: "radar", Output: true},
			{Input: "level", Output: true},
			{Input: "not", Output: false},
		},
		testData: []io.IO{
			{Input: "race a car", Output: false},
			{Input: "A man, a plan, a canal: Panama", Output: true},
			{Input: "0P", Output: false},
		},
		versions: []func(s string) bool{},
	}
	p.versions = append(p.versions, p.isValidPalindromeV1)

	return p
}
