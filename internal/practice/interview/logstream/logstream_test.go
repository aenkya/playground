package logstream

import (
	"fmt"
	"testing"
)

func TestLogStream(t *testing.T) {
	ls := NewLogStream()

	fmt.Println("LogStreamV1")

	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "LogStreamV1",
			input: []string{
				"Q: create database",
				"Q: close database",
				"Q: create model",
				"L: Database created",
				"L: model created",
				"Q: close connection",
				"L: connection closed",
				"Q: Database closed",
			},
			expected: []string{
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := ls.logStreamV1(test.input)

			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", len(test.expected), len(actual))
			}

			for i := 0; i < len(actual); i++ {
				if actual[i] != test.expected[i] {
					t.Errorf("Expected %s, got %s", test.expected[i], actual[i])
				}
			}
		})
	}
}
