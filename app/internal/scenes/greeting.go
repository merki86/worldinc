package scenes

import (
	"fmt"
	"os"
	"worldinc/app/pkg/window"
)

type Greeting struct{}

func (s Greeting) Run() {
	fmt.Println("Greeting screen initialized")
	window := window.Window{
		Title: "World Inc.",
		Buttons: []window.Button{
			{
				Label:    "Start",
				Callback: Menu{}.Run,
			},
			{
				Label:    "Quit",
				Callback: func() { os.Exit(0) },
			},
		},
		ContentCallback: callback,
	}
	window.New()
}

func callback() {
	fmt.Println(
		"Welcome to World Inc! This is a remastered demo of my original game back in 2020.\n",
		"It was originally written in pygame, but was soon abandoned due to lack of my skills.\n",
		"You have a world map and a list of various symptoms. You should buy them to increase mortality of the disease.\n",
		"Your goal is to kill every single person on Earth before the disease will be discovered and destroyed.",
	)
}
