package verifier

import (
	"testing"
)

func TestCompareAndPrint(t *testing.T) {
	tests := []struct {
		correctWord string
		guess       string
		expected    string
		isCorrect   bool
	}{
		{
			correctWord: "apple",
			guess:       "apple",
			expected:    "\033[32mapple\033[0m", // All Green
			isCorrect:   true,
		},
		{
			correctWord: "apple",
			guess:       "apric",
			expected:    "\033[32map\033[31mr\033[31mi\033[31mc\033[0m",
			isCorrect:   false,
		},
		{
			correctWord: "blimp",
			guess:       "plumb",
			expected:    "\033[33mp\033[32ml\033[31mu\033[32mm\033[33mb\033[0m",
			isCorrect:   false,
		},
		{
			correctWord: "sword",
			guess:       "words",
			expected:    "\033[33mw\033[33mo\033[33mr\033[33md\033[33ms\033[0m",
			isCorrect:   false,
		},
		{
			correctWord: "happy",
			guess:       "worry",
			expected:    "\033[31mw\033[31mo\033[31mr\033[31mr\033[32my\033[0m",
			isCorrect:   false,
		},
	}

	for _, test := range tests {
		isCorrect := CompareAndPrint(test.correctWord, test.guess)
		result := FormatResult(test.correctWord, test.guess)

		if isCorrect != test.isCorrect {
			t.Errorf("For correctWord '%s' and guess '%s', expected isCorrect %v, got %v", test.correctWord, test.guess, test.isCorrect, isCorrect)

		}
		expectedASCII := toASCIIString(test.expected)
		resultASCII := toASCIIString(result)
		if result != test.expected {
			t.Errorf("For correctWord '%s' and guess '%s', expected output %s (%s), got %s (%s)", test.correctWord, test.guess, test.expected, expectedASCII, result, resultASCII)
		}
	}
}
