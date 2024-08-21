module github.com/SahilRathod17/go-wordle/main

go 1.20

require (
	github.com/SahilRathod17/go-wordle/game v0.0.1
	github.com/SahilRathod17/go-wordle/words v0.0.1
	github.com/SahilRathod17/go-wrodle/verifier v0.0.1
)

replace github.com/SahilRathod17/go-wordle/game => ../game

replace github.com/SahilRathod17/go-wordle/words => ../words

replace github.com/SahilRathod17/go-wordle/verifier => ../verifier
