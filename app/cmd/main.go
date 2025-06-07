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

type gameScene struct {
	next model.Scene
}

func NewGameScene() *gameScene {
	return &gameScene{}
}

func (g *gameScene) Update() {
	// Update
}

func (g *gameScene) Draw(s tcell.Screen) {
	// Draw
}

func (g *gameScene) Next() model.Scene {
	return g.next
}

func (g *gameScene) EventHandler(s tcell.Screen) {
	event := s.PollEvent()
	switch event := event.(type) {
	case *tcell.EventKey:
		game.Mutex.Lock()
		switch event.Key() {
		case tcell.KeyEscape:
			// game.CurrentScene = gameScene.Next()
			os.Exit(0)
		}
		game.Mutex.Unlock()
	}
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

	var current model.Scene = NewGameScene()

	for {
		current.Update()
		current.Draw(screen)
		current.EventHandler(screen)

		if next := current.Next(); next != nil {
			current = next
		}
	}
}
