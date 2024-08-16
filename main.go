package main

import (
	"fmt"
	"github.com/IDOMATH/CheetahEngine/Game"
	"time"
)

func main() {
	fmt.Println("Hello world!")
	game := Game.NewGame(20 * time.Millisecond)
	game.Run()

}
