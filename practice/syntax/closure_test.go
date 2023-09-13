package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosure(t *testing.T) {
	t.Run("TestClosure", func(t *testing.T) {
		closure()
	})

	t.Run("TestArrayUpdating", func(t *testing.T) {
		arraySize()
		assert.Equal(t, true, false)
	})
}
