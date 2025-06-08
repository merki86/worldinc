package main

import (
	"time"
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

	game.CurrentScene = scene.NewGameScene(&game)

	go handle(&game, screen)
	go logic(&game)
	render(&game, screen)
}

func handle(game *model.GameState, screen tcell.Screen) {
	for {
		ev := screen.PollEvent()

		game.Mutex.Lock()
		scene := game.CurrentScene
		game.Mutex.Unlock()

		if scene != nil {
			scene.HandleEvent(ev)
		}
	}
}

func logic(game *model.GameState) {
	tick := time.NewTicker(500 * time.Millisecond)
	defer tick.Stop()

	for range tick.C {
		game.Mutex.Lock()
		scene := game.CurrentScene
		game.Mutex.Unlock()

		if scene != nil {
			scene.Update(500 * time.Millisecond)
		}
	}
}

func render(game *model.GameState, screen tcell.Screen) {
	tick := time.NewTicker(1 * time.Second / 30) // 30 FPS
	defer tick.Stop()

	for range tick.C {
		screen.Clear()

		game.Mutex.Lock()
		scene := game.CurrentScene
		game.Mutex.Unlock()

		if scene != nil {
			scene.Draw(screen)
		}

		screen.Show()
	}
}
