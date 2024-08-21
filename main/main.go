package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/SahilRathod17/go-wordle/game"
	"github.com/SahilRathod17/go-wordle/words"
)

//cat words.txt | tr '[:upper:]' '[:lower:]' | awk 'length($0) == 5' | sed 's/^[ \t]*//;s/[ \t]*$//' > valid_words.txt
//shuf valid_words.txt > random_valid_words.txt

func main() {
	fmt.Println("Welcome to wordle!!")

	wordList, err := words.LoadWords("words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	correctWord := game.GetWord(wordList)
	game.PlayGame(correctWord, nil)

}

func handleInterrupt() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("\nGame interrupted. Exiting...")
		os.Exit(0)

	}()
}

func init() {
	handleInterrupt()
}
