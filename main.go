package main

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/SahilRathod17/go-wordle/game"
	"github.com/SahilRathod17/go-wordle/words"
)

//go:embed words.txt
var content embed.FS

func main() {
	file, err := content.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening embedded file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

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
