package controllers

import (
	"fmt"
	"os"
	"strings"

	"github.com/tkivisik/playfulgo/games"
	"github.com/tkivisik/playfulgo/views"
)

const (
	BOT   bool = true
	HUMAN bool = false
)

func NewGameplays() *Gameplays {
	return &Gameplays{
		GameplayView:  views.NewView("general", "layouts/general"),
		FriendlyBoard: games.NewBoard(),
		HostileBoard:  games.NewBoard(),
	}
}

type Gameplays struct {
	//ShipchoosingView *views.View
	GameplayView  *views.View
	FriendlyBoard *games.Board
	HostileBoard  *games.Board
}

func (g Gameplays) InitBoards() {
	//	data := PlayData{Legend: games.Legend}
	data := PlayData{Legend: games.Legend}
	g.updateAndRender(data)

	for nShipsPlaced := 0; nShipsPlaced < games.MaxShips; nShipsPlaced++ {
		g.FriendlyBoard.AddShipBy(HUMAN)
		g.HostileBoard.AddShipBy(BOT)

		g.updateAndRender(data)
	}
}

func (g Gameplays) Play() {
	//g.GameplayView.Template
	data := PlayData{Legend: games.Legend}
	g.updateAndRender(data)

	for g.FriendlyBoard.HitCount < 2 && g.HostileBoard.HitCount < 2 {
		g.FriendlyBoard.ShootThisBoard(BOT)
		g.HostileBoard.ShootThisBoard(HUMAN)

		g.updateAndRender(data)
	}

	myScore := int(g.HostileBoard.HitCount)
	enemyScore := int(g.FriendlyBoard.HitCount)

	if enemyScore >= games.MaxShips {
		if myScore >= games.MaxShips {
			fmt.Println("IT'S A DRAW, GAME OVER, WELL DONE")
		} else {
			fmt.Println("GAME OVER, YOU LOST")
		}
		os.Exit(0)
	} else {
		if myScore >= games.MaxShips {
			fmt.Println("GAME OVER. YOU WON !!!")
			os.Exit(0)
		}
	}

}

type PlayData struct {
	Legend              games.LegendStruct
	FriendlyBoardString []string
	HostileBoardString  []string
}

func (g Gameplays) updateAndRender(data PlayData) {
	data.FriendlyBoardString = strings.Split(g.FriendlyBoard.String(HUMAN), "\n")
	data.HostileBoardString = strings.Split(g.HostileBoard.String(BOT), "\n")

	g.GameplayView.Render(data)
}
