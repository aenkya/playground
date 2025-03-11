package syntax

import "testing"

func TestSyntax(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestSyntax"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MongoDriverPractice()
		})
	}
}
