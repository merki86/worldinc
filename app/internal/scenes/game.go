package scenes

import (
	"fmt"
	"worldinc/app/pkg/window"
)

type Game struct{}

func (s Game) Run() {
	fmt.Println("Game initialized")
	window := window.Window{
		Title: "General stats",
		Buttons: []window.Button{
			{
				Label:    "Symptoms",
				Callback: s.Run,
			},
			{
				Label:    "Detailed stats",
				Callback: Greeting{}.Run,
			},
			{
				Label:    "Back",
				Callback: Menu{}.Run,
			},
		},
		ContentCallback: gameCallback,
	}
	window.New()
}

func gameCallback() {
	fmt.Print(
		" ▄▄██▀██▄ ▀   ▄▄█▄  \n",
		"▀█████ ▀  ▄██████▀▀▄\n",
		"  ▀█▀   ▄█▄████▀ ▀  \n",
		"    ▀██▄ ▀██   ▀▄▄  \n",
		"     █▀   ▀    ▀▀█ ▄\n",
	)
}
