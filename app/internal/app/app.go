package app

import (
	"fmt"
	"worldinc/app/pkg/window"
)

func Run() {
	fmt.Println("World Inc.")
	window := window.Window{
		Title: "World Inc.",
		Buttons: []window.Button{
			{
				Label:    "Restart",
				Literal:  "R",
				Callback: Run,
			},
			{
				Label:    "Quit",
				Literal:  "Q",
				Callback: Run,
			},
		},
	}
	window.New()
}
