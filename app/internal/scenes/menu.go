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
		Title: "World Inc.",
		Buttons: []window.Button{
			{
				Label:    "Start",
				Callback: Game{}.Run,
			},
			{
				Label:    "Back",
				Callback: Greeting{}.Run,
			},
			{
				Label:    "Quit",
				Callback: func() { os.Exit(0) },
			},
		},
		ContentCallback: func() {},
	}
	window.New()
}
