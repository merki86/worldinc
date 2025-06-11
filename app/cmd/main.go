package main

import (
	"fmt"
	"time"
	"worldinc/app/internal/model"
	"worldinc/app/internal/scene"
	"worldinc/app/pkg/print"

	"github.com/gdamore/tcell/v2"
)

// TODO: move the shit outta here
var SymptomsList = []model.Symptom{
	{
		ID:                1,
		Name:              "Nausea",
		MortalityBonus:    0,
		TransmissionBonus: 0.001,
		Cost:              5,
		Unlocked:          false,
	},
	{
		ID:                2,
		Name:              "Insomnia",
		MortalityBonus:    0.0005,
		TransmissionBonus: 0.003,
		Cost:              30,
		Unlocked:          false,
	},
	{
		ID:                3,
		Name:              "Skin Rash",
		MortalityBonus:    0.001,
		TransmissionBonus: 0.016,
		Cost:              45,
		Unlocked:          false,
	},
	{
		ID:                4,
		Name:              "Cough",
		MortalityBonus:    0.02,
		TransmissionBonus: 0.2,
		Cost:              60,
		Unlocked:          false,
	},
	{
		ID:                5,
		Name:              "Fever",
		MortalityBonus:    0.05,
		TransmissionBonus: 0.1,
		Cost:              50,
		Unlocked:          false,
	},
}

// TODO: Execute function that generates the struct
var game = model.GameState{
	World: model.World{
		Healthy:  80000,
		Infected: 1,
		Dead:     0,
		Disease: model.Disease{
			Name:         "Bacteria",
			Mortality:    0,
			Transmission: 0,
		},
		DaysPassed: 0,
		Credit:     5,
		Speed:      time.Second,
	},
	Symptoms:   SymptomsList,
	Gameticker: time.NewTicker(time.Second),
}

// TODO: Into another file maybe?
var (
	frameCount int
	lastTime   = time.Now()
	fps        int
)

func main() {
	// for {
	// 	l.DoWorldTick(&game.World)
	// }
	screen, _ := tcell.NewScreen()
	screen.Init()
	defer screen.Fini()

	game.CurrentScene = scene.NewMenuScene(&game)

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
	defer game.Gameticker.Stop()

	for range game.Gameticker.C {
		game.Mutex.Lock()
		scene := game.CurrentScene
		game.Mutex.Unlock()

		if scene != nil {
			scene.Update()
		}
	}
}

func render(game *model.GameState, screen tcell.Screen) {
	tick := time.NewTicker(time.Second / 30)
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
