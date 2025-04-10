package algo

import "testing"

func TestDecodeVariations(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"12", 2},    // "AB" (1,2) and "L" (12)
		{"226", 3},   // "BZ" (2,26), "VF" (22,6), "BBF" (2,2,6)
		{"0", 0},     // No valid decoding
		{"06", 0},    // No valid decoding
		{"10", 1},    // "J" (10)
		{"27", 1},    // "BG" (2,7)
		{"11106", 2}, // "AAJF" (1,1,10,6) and "KJF" (11,10,6)
		{"", 0},      // Empty string
		{"1", 1},     // "A" (1)
		{"111", 3},   // "AAA" (1,1,1), "KA" (11,1), "AK" (1,11)
	}

	for _, test := range tests {
		result := decodeVariations(test.input)
		if result != test.expected {
			t.Errorf("decodeVariations(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}
