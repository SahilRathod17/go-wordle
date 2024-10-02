package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/SahilRathod17/go-wordle/verifier"
)

const (
	maxWordLength = 5
	maxAttempts   = 6
)

// Handling hints
var usedHint bool

// Provide hint
func giveHint(correctWord string) string {
	if usedHint {
		return "You can not use more than one hint..."
	}
	usedHint = true
	randomIndex := rand.Intn(len(correctWord))
	hint := string(correctWord[randomIndex])
	return fmt.Sprintf("Hint: The letter at position %d is '%s'\n", randomIndex+1, hint)
}

// Reset hint
func resetHint() {
	usedHint = false
}

// InputProvider is an interface for providing user input.
type InputProvider interface {
	guessFromUser() string
}

// RealInputProvider implements InputProvider and reads input from the real user.
type RealInputProvider struct{}

// GetWord returns a random word from the given list.
func GetWord(wordList []string) string {
	return wordList[rand.Intn(len(wordList))]
}

// PlayGame starts a game with the given correct word and input provider.
func PlayGame(correctWord string, input InputProvider) {
	attempts := 0
	wrongAttempts := 0

	if input == nil {
		input = &RealInputProvider{}
	}

	for attempts < maxAttempts {
		guess := input.guessFromUser()

		// Check if player asked for hint
		if strings.ToLower(guess) == "hint" {
			fmt.Println(giveHint(correctWord))
			attempts++
			fmt.Printf("You have %d attempts left.\n", maxAttempts-attempts)
			continue
		}

		if wrongAttempts > 2 {
			fmt.Println("Too many worng attempts, please start again.")
			return
		}

		if len(guess) != maxWordLength {
			wrongAttempts++
			fmt.Printf("Guessed word must be %d letters long.\n", maxWordLength)
			continue
		}

		isCorrect := verifier.CompareAndPrint(correctWord, guess)
		if isCorrect {
			fmt.Println("Congratulation!! You guessed the word!")
			resetHint()
			return
		}

		attempts++
		fmt.Printf("You have %d attempts left.\n", maxAttempts-attempts)
	}
	fmt.Printf("Sorry, you've used all your attempts, The word was: %s\n", correctWord)
}

func (rip *RealInputProvider) guessFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	if usedHint {
		fmt.Print("What's your guess: ")
	} else {
		fmt.Print("What's your guess (or type 'hint' for a hint): ")
	}
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(input))
}
