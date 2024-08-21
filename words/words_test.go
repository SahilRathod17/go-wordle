package words

import (
	"embed"
	"testing"
)

//go:embed test_words.txt
var testContent embed.FS

func TestLoadWords(t *testing.T) {

	filePath := "test_words.txt"
	words, err := LoadWords(testContent, filePath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(words) != 1 {
		t.Errorf("Expected 1 word, got %d", len(words))
	}

	expectedWords := map[string]bool{
		"apple": true,
	}

	for _, word := range words {
		if _, ok := expectedWords[word]; !ok {
			t.Errorf("Unexpected word: %s", word)
		}
	}

}
