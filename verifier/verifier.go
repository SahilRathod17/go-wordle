package verifier

import (
	"fmt"
	"strings"
)

const (
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Reset  = "\033[0m"
)

func CompareAndPrint(correctWord, guess string) bool {
	result := make([]byte, len(correctWord))
	isCorrect := true

	for i := 0; i < len(correctWord); i++ {
		if guess[i] == correctWord[i] {
			result[i] = 'G'
		} else if strings.ContainsRune(correctWord, rune(guess[i])) {
			result[i] = 'Y'
			isCorrect = false
		} else {
			result[i] = 'R'
			isCorrect = false
		}
	}
	printResult(result, guess)
	return isCorrect
}

func printResult(result []byte, guess string) {
	for i, r := range result {
		switch r {
		case 'G':
			fmt.Printf("%s%c%s", Green, guess[i], Reset)
		case 'Y':
			fmt.Printf("%s%c%s", Yellow, guess[i], Reset)
		default:
			fmt.Printf("%s%c%s", Red, guess[i], Reset)
		}
	}
	fmt.Println()
}

// For testing
func FormatResult(correctWord, guess string) string {
	result := make([]byte, len(correctWord))

	for i := 0; i < len(correctWord); i++ {
		if guess[i] == correctWord[i] {
			result[i] = 'G'
		} else if strings.ContainsRune(correctWord, rune(guess[i])) {
			result[i] = 'Y'
		} else {
			result[i] = 'R'
		}
	}
	return formatResult(result, guess)
}

func formatResult(result []byte, guess string) string {
	var sb strings.Builder
	for i, r := range result {
		// if r == 'G' || r == 'Y' {
		// 	sb.WriteByte(guess[i])
		// }
		switch r {
		case 'G':
			sb.WriteString(Green + guess[i:i+1] + Reset)
		case 'Y':
			sb.WriteString(Yellow + guess[i:i+1] + Reset)
		default:
			sb.WriteString(Red + guess[i:i+1] + Reset)
		}
	}
	return sb.String()
}

func toASCIIString(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		sb.WriteString(fmt.Sprintf("%d ", s[i]))
	}
	return sb.String()
}
