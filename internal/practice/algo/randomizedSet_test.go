package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		actions []string
		inputs  []any
		expects []any
	}{
		{
			[]string{"insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"},
			[]any{1, 2, 2, nil, 1, 2, nil},
			[]any{true, false, true, nil, true, false, 2},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			rs := Constructor()

			for i, action := range tt.actions {
				switch action {
				case "insert":
					got := rs.Insert(tt.inputs[i].(int))
					assert.Equalf(t, tt.expects[i], got, "%s not equal to %s", tt.expects[i], got)
				case "remove":
					got := rs.Remove(tt.inputs[i].(int))
					assert.Equalf(t, tt.expects[i], got, "%s not equal to %s", tt.expects[i], got)
				}
			}
		})
	}
}
