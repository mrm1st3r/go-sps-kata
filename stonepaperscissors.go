package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"encoding/json"
	"fmt"
	"math/rand"
)

type Move struct {
	Move string
}

var validMoves = []string{"stone", "paper", "scissors"}

type Result struct {
	PlayerMove string
	ComputerMove string
	Winner string
}

func main() {
	router := fasthttprouter.New()
	router.POST("/play", Play)
	fasthttp.ListenAndServe(":12345", router.Handler)
}

func Play(ctx *fasthttp.RequestCtx) {
	var m Move
	err := json.Unmarshal(ctx.PostBody(), &m)
	if err != nil {
		fmt.Fprintf(ctx, "Unable to parse request data: %s", err)
		ctx.SetStatusCode(400)
		return
	}

	if !isValidMove(m) {
		fmt.Fprintf(ctx, "Unknown move: %s", m.Move)
		ctx.SetStatusCode(400)
		return
	}

	computerMove := computerMove()
	winner := determineWinner(m, computerMove)
	result := Result{m.Move, computerMove.Move, winner}
	res, err := json.Marshal(result)
	fmt.Fprint(ctx, string(res))
}

func determineWinner(user Move, computer Move) string {
	if user.Move == computer.Move {
		return "draw"
	}
	if user.Move == "stone" && computer.Move == "scissors" ||
		user.Move == "scissors" && computer.Move == "paper" ||
		user.Move == "paper" && computer.Move == "stone" {
		return "user"
	}
	return "computer"
}

func isValidMove(move Move) bool {
	for _, m := range validMoves {
		if m == move.Move {
			return true
		}
	}
	return false
}

func computerMove() Move {
	n := rand.Int31n(int32(len(validMoves)))
	return Move{validMoves[n]}
}
