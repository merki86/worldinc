package main

import (
	"worldinc/app/internal/model"
	"worldinc/app/internal/scene"

	"github.com/gdamore/tcell/v2"
)

var game = model.GameState{
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

	var current model.Scene = scene.NewGameScene(&game)
	startScene(current, screen)
}

func startScene(current model.Scene, screen tcell.Screen) {
	done := make(chan struct{})

	go func() {
		current.Draw(screen, done)
	}()
	go func() {
		current.Update(done)
	}()
	for {
		current.HandleEvent(screen)
		if next := current.Next(); next != nil {
			close(done)
			game.Mutex.Lock()
			current = next
			game.Mutex.Unlock()
			startScene(current, screen)
			return
		}
	}
}
