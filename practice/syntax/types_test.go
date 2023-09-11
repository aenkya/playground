package syntax

import "testing"

func TestTypes(t *testing.T) {
	t.Run("TestTypes", func(t *testing.T) {
		types()
	})

	t.Run("TestCheckTypes", func(t *testing.T) {
		if checkTypes(1) != "int" {
			t.Error("Expected int")
		}
		if checkTypes("1") != "string" {
			t.Error("Expected string")
		}
		if checkTypes(1.0) != "unknown" {
			t.Error("Expected unknown")
		}
	})
}
