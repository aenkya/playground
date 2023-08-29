package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:govet // ignore struct optimisation
func TestCreation(t *testing.T) {
	tests := []struct {
		name     string
		original Person
		fields   []Pair[string, string]
		want     Person
	}{
		{
			name: "creates a new person with new name",
			original: Person{
				Name: "John",
				Address: &Address{
					StreetAddress: "123 London Rd",
					City:          "London",
					Country:       "UK",
				},
			},
			fields: []Pair[string, string]{
				{"name", "Jane"},
			},
			want: Person{
				Name: "Jane",
				Address: &Address{
					StreetAddress: "123 London Rd",
					City:          "London",
					Country:       "UK",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			assert.Equal(t, tt.want, Creational(tt.original, tt.fields...))
		})
	}
}
