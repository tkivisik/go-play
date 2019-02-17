package controllers

import (
	"strings"

	"github.com/tkivisik/playfulgo/games"
	"github.com/tkivisik/playfulgo/views"
)

func NewGameplays() *Gameplays {
	return &Gameplays{
		NewGameplay: views.NewView("general", "layouts/general"),
		Boards:      []games.Board{*games.NewBoard(), *games.NewBoard()},
	}
}

type Gameplays struct {
	NewGameplay *views.View
	Boards      []games.Board
}

func (g Gameplays) InitBoards() {
	g.Boards[0].Init(false)
	g.Boards[1].Init(true)
}

func (g Gameplays) Play() {
	//g.NewGameplay.Template
	data := PlayData{Legend: games.Legend}
	data.BoardsSlice[0] = strings.Split(g.Boards[0].String(false), "\n")
	data.BoardsSlice[1] = strings.Split(g.Boards[1].String(true), "\n")
	g.NewGameplay.Render(data)

	for g.Boards[0].HitCount < 2 && g.Boards[1].HitCount < 2 {
		g.Boards[0].ShootThisBoard(true)
		g.Boards[1].ShootThisBoard(false)

		data.BoardsSlice[0] = strings.Split(g.Boards[0].String(false), "\n")
		data.BoardsSlice[1] = strings.Split(g.Boards[1].String(true), "\n")
		g.NewGameplay.Render(data)
	}

}

type PlayData struct {
	Legend      games.LegendStruct
	BoardsSlice [2][]string
}
