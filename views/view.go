package views

import (
	"text/template"
)

func NewView(layout string, files ...string) *View {
	return &View{Template: template.New("legend"),
		Layout: ""}
}

type View struct {
	Template *template.Template
	Layout   string
}
