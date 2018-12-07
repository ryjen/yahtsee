package internal

import (
	"log"
	"runtime"
)

type Game struct {
	ui *Ui
	config *Config
}

func NewGame() *Game {
	return &Game{
		NewUi(),
		NewConfig(),
	}
}

func (game *Game) Start() error {

	err := game.config.Load()

	if err != nil {
		log.Print("Warning: ", err)
	}

	runtime.LockOSThread()

	return game.ui.Init(game.config)
}

func (game *Game) IsOver() bool {
	return game.config.Over || !game.ui.IsOpen()
}

func (game *Game) Update() {
	game.ui.Update()
}

func (game *Game) Render() {
	game.ui.Render()
}

func (game *Game) Finish() {
	game.ui.Close()
}
