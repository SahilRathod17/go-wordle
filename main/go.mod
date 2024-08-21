module github.com/SahilRathod17/go-wordle/main

go 1.20

require (
	github.com/SahilRathod17/go-wordle/game v0.0.0
	github.com/SahilRathod17/go-wordle/words v0.0.0
)

require github.com/SahilRathod17/go-wordle/verifier v0.0.0 // indirect

replace (
	github.com/SahilRathod17/go-wordle/game => ../game
	github.com/SahilRathod17/go-wordle/verifier => ../verifier
	github.com/SahilRathod17/go-wordle/words => ../words
)
