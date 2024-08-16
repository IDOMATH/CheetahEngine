package Game

import (
	"fmt"
	"os"
	"time"
)

type Game struct {
	refresh       time.Duration
	buttonPressed chan string
	counter       int
}

func NewGame(refresh time.Duration) *Game {
	return &Game{
		refresh:       refresh,
		buttonPressed: make(chan string),
		counter:       0,
	}
}

func (g *Game) Run() {
	fmt.Println("Game running")
	var b []byte = make([]byte, 1)

	for {
		time.Sleep(g.refresh)
		fmt.Println("Frame: ", g.counter)
		g.counter++

		os.Stdin.Read(b)
		fmt.Println("key: ", string(b))
		if string(b) == "x" {
			return
		}
	}
}
