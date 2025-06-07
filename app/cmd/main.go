package main

import (
	"os"
	"sync"
	"worldinc/app/internal/model"

	"github.com/gdamore/tcell/v2"
)

type GameState struct {
	World        model.World
	CurrentScene model.Scene
	Mutex        sync.Mutex
	// Input string
	// Logs  []string
}

type gameScene struct{}

func (g gameScene) Draw() {
	os.Exit(0)
}

var game = GameState{
	World: model.World{
		Population: 5000,
		DaysPassed: 1,
		Regions: []model.Region{
			{
				Name:       "Europe",
				Population: 200,
			},
			{
				Name:       "Asia",
				Population: 2020,
			},
		},
	},
}

func main() {
	screen, _ := tcell.NewScreen()
	screen.Init()
	defer screen.Fini()

	go inputLoop(screen)
	// go logicLoop()
	// renderLoop(screen)
	for {

	}
}

func inputLoop(s tcell.Screen) {
	for {
		event := s.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			game.Mutex.Lock()
			switch event.Key() {
			case tcell.KeyEscape:
				// game.CurrentScene = gameScene.Next()
				gameScene{}.Draw()
			}
			game.Mutex.Unlock()
		}
	}
}
