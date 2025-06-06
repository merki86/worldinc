package scenes

import (
	"fmt"
	"os"
	"worldinc/app/pkg/window"
)

type Menu struct{}

func (s Menu) Run() {
	fmt.Println("Menu initialized")
	window := window.Window{
		Title:   "World Inc.",
		Message: "Welcome to World Inc! This is a remastered demo of my original game back in 2020.",
		Buttons: []window.Button{
			{
				Label:    "Start",
				Callback: s.Run,
			},
			{
				Label:    "Restart",
				Callback: s.Run,
			},
			{
				Label:    "Quit",
				Callback: func() { os.Exit(0) },
			},
		},
	}
	window.New()
}
