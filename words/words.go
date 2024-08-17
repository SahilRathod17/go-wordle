package words

import (
	"bufio"
	"os"
	"strings"
)

func LoadWords(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		// fmt.Printf("Processing Raw Word: '%s' | Length: %d\n", word, len(word))
		if len(word) == 5 {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
