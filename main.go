// Copyright © 2018 Taavi Kivisik
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// go-play is a package for playful exploration of Golang
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tkivisik/playfulgo/controllers"

	"github.com/tkivisik/playfulgo/games"
)

type gameplayData struct {
	Legend       games.LegendStruct
	BoardsString [2][]string
	Boards       []games.Board
}

func main() {
	//	games.Battleship()
	gameplayC := controllers.NewGameplays()

	// tmpl, err := template.New("legend").ParseFiles("views/layouts/legend.tmpl", "views/layouts/gameplay.tmpl", "views/layouts/boards.tmpl", "views/layouts/footer.tmpl")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	data := gameplayData{
		Legend:       games.Legend,
		BoardsString: [2][]string{},
		Boards:       [](games.Board){*games.NewBoard(), *games.NewBoard()},
	}
	data.BoardsString[0] = strings.Split(data.Boards[0].String(false), "\n")
	data.BoardsString[1] = strings.Split(data.Boards[1].String(true), "\n")
	err := gameplayC.NewGameplay.Template.Execute(os.Stdout, data)
	//err = tmpl.ExecuteTemplate(os.Stdout, "gameplay", data)
	if err != nil {
		fmt.Println(err)
	}
	/*	fmt.Println("b")
		fmt.Printf("%+v", tmpl)
		tmpl.Execute(os.Stdout, nil)*/
}
