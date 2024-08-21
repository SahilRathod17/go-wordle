package game

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

type MockInputProvider struct {
	Inputs []string
	index  int
}

func (mip *MockInputProvider) guessFromUser() string {
	if mip.index >= len(mip.Inputs) {
		return ""
	}
	input := mip.Inputs[mip.index]
	mip.index++
	return input
}

func TestPlayGame(t *testing.T) {
	tests := []struct {
		correctWord    string
		guessedWords   []string
		expectedOutput string
		attempt        int
	}{
		{
			correctWord:    "apple",
			guessedWords:   []string{"apple"},
			expectedOutput: "Congratulation!! You guessed the word!",
			attempt:        1,
		},
		{
			correctWord:    "plane",
			guessedWords:   []string{"benten", "wrodle", "airbnb"},
			expectedOutput: "Too many worng attempts, please start again.",
			attempt:        3,
		},
		{
			correctWord:    "happy",
			guessedWords:   []string{"sad", "happy"},
			expectedOutput: "Congratulation!! You guessed the word!",
			attempt:        2,
		},
		{
			correctWord:    "flame",
			guessedWords:   []string{"sad", "bull", "apple", "waves", "total", "begun", "bangl", "black"},
			expectedOutput: "Sorry, you've used all your attempts, The word was: flame",
			attempt:        5,
		},
	}

	for _, test := range tests {

		origin := os.Stdout

		r, w, _ := os.Pipe()
		os.Stdout = w

		mockInput := &MockInputProvider{Inputs: test.guessedWords}

		PlayGame(test.correctWord, mockInput)
		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)

		os.Stdout = origin

		output := removeColorCoding(buf.String())
		lines := strings.Split(output, "\n")
		if len(lines) < 1 {
			t.Fatalf("No output received!!")
		}

		lastLine := lines[len(lines)-2]
		fmt.Print(output)

		if lastLine != test.expectedOutput {
			t.Errorf("Expected %q but got %q", test.expectedOutput, output)
		}
	}
}

func removeColorCoding(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(input, "")
}
