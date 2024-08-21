package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/SahilRathod17/go-wordle/verifier"
)

const (
	maxWordLength = 5
	maxAttempts   = 6
)

type InputProvider interface {
	guessFromUser() string
}

type RealInputProvider struct{}

func GetWord(wordList []string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return wordList[rand.Intn(len(wordList))]
}

func PlayGame(correctWord string, input InputProvider) {
	attempts := 0
	wrongAttempts := 0

	if input == nil {
		input = &RealInputProvider{}
	}

	for attempts < maxAttempts {
		guess := input.guessFromUser()

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
			return
		}

		attempts++
		fmt.Printf("You have %d attempts left.\n", maxAttempts-attempts)
	}
	fmt.Printf("Sorry, you've used all your attempts, The word was: %s\n", correctWord)
}

func (rip *RealInputProvider) guessFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What's your guess: ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(input))
}
