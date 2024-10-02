package main

import (
	"embed"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/common-nighthawk/go-figure"

	"github.com/SahilRathod17/go-wordle/game"
	"github.com/SahilRathod17/go-wordle/words"
)

//go:embed words.txt
var content embed.FS

// starting point with, handling the keyboard intrrupt
func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	handleInterrupt()
}

func main() {
	myFigure := figure.NewColorFigure("Welcome to wordle!!", "", "green", true)
	myFigure.Print()
	fmt.Println()
	wordList, err := words.LoadWords(content, "words.txt")
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
