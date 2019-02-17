package views

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".tmpl"
)

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
	for i, f := range files {
		files[i] = f + TemplateExt
	}

	layoutFiles, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}

	files = append(files, layoutFiles...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// Render builds a template using data
func (v *View) Render(data interface{}) {
	err := v.Template.ExecuteTemplate(os.Stdout, v.Layout, data)
	if err != nil {
		fmt.Println(err)
	}
}
