package handler

import "testing"

func TestBase62Encode(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{61, "Z"},
		{62, "10"},
		{125, "21"},
		{3844, "100"}, // 62^2
	}

	for _, test := range tests {
		result := Base62Encode(test.input)
		if result != test.expected {
			t.Errorf("Base62Encode(%d) = %s; expected %s", test.input, result, test.expected)
		}
	}
}