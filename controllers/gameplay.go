package controllers

import (
	"github.com/tkivisik/playfulgo/views"
)

func NewGameplays() *Gameplays {
	return &Gameplays{
		NewGameplay: views.NewView("cli", "gameplay"),
	}
}

type Gameplays struct {
	NewGameplay *views.View
}
