package main

import "testing"

type testData struct {
	user string
	computer string
	expectedWinner string
}

var data = []testData {
	testData{"stone", "stone", "draw"},
	testData{"paper", "scissors", "computer"},
	testData{"stone", "scissors", "user"},
	testData{"scissors", "paper", "user"},
	testData{"stone", "paper", "computer"},
	testData{"paper", "stone", "user"},

}

func TestDetermineWinner(t *testing.T) {
	for _, d := range data {
		actualWinner := determineWinner(Move{d.user}, Move{d.computer})
		if actualWinner != d.expectedWinner {
			t.Error("For", d.user, "and", d.computer, "expected", d.expectedWinner,"got", actualWinner)
		}
	}
}
