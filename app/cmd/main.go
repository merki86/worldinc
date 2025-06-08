package main

import (
	"fmt"
	"time"
	"worldinc/app/internal/model"
	"worldinc/app/internal/scene"
	"worldinc/app/pkg/print"

	"github.com/gdamore/tcell/v2"
)

var game = model.GameState{
	World: model.World{
		Population: 5000,
		Infected:   0,
		Dead:       0,
		Disease: model.Disease{
			Name:         "Bacteria",
			Mortality:    0,
			Transmission: 0,
			Discovered:   false,
			Symptoms: []model.Symptom{
				{
					Name:           "Nausea",
					MortalityBonus: 0,
					Cost:           5,
					Unlocked:       true,
				},
			},
		},
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
		DaysPassed: 0,
	},
}

var (
	frameCount int
	lastTime   = time.Now()
	fps        int
)

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
	tick := time.NewTicker(200 * time.Millisecond)
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
	tick := time.NewTicker(1 * time.Second / 30)
	defer tick.Stop()

	for range tick.C {
		frameCount++
		if time.Since(lastTime) >= time.Second {
			fps = frameCount
			frameCount = 0
			lastTime = time.Now()
		}

		screen.Clear()

		game.Mutex.Lock()
		scene := game.CurrentScene
		game.Mutex.Unlock()

		if scene != nil {
			scene.Draw(screen)
		}

		print.Print(screen, 0, 0, fmt.Sprintf("FPS: %v", fps))

		screen.Show()
	}
}
