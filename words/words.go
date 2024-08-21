package words

import (
	"bufio"
	"embed"
	"fmt"
	"strings"
)

// LoadWords loads words from a file and returns them as a slice of strings.
func LoadWords(content embed.FS, filePath string) ([]string, error) {
	file, err := content.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening embedded file: %w", err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if len(word) == 5 {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return words, nil
}
