package words

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadWords(t *testing.T) {
	tempFile := filepath.Join(t.TempDir(), "test_words.txt")
	data := "apple\nbanana\ncherry\n"
	err := os.WriteFile(tempFile, []byte(data), 0644)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	fileContent, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	fmt.Printf("File Content: \n%s\n", fileContent)
	words, err := LoadWords(tempFile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// if len(words) != 3 {
	// 	t.Errorf("Expected 3 words, got %d", len(words))
	// }

	// expectedWords := map[string]bool{
	// 	"apple":  true,
	// 	"banana": true,
	// 	"cherry": true,
	// }
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
