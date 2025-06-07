package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
	"worldinc/app/internal/model"

	"github.com/gdamore/tcell/v2"
)

type GameState struct {
	World        model.World
	CurrentScene Scene
	Input        string
	Logs         []string // Логи команд
	Mutex        sync.Mutex
}

type Scene int

const (
	SceneMain Scene = iota
	SceneSymptoms
)

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
	Logs: []string{"Welcome to World Inc!"},
}

func main() {
	screen, _ := tcell.NewScreen()
	screen.Init()
	defer screen.Fini()

	go inputLoop(screen)
	go logicLoop()
	renderLoop(screen)
}

func renderLoop(s tcell.Screen) {
	for {
		s.Clear()
		game.Mutex.Lock()
		DrawGame(s, &game)
		game.Mutex.Unlock()
		s.Show()
		time.Sleep(100 * time.Millisecond)
	}
}

func inputLoop(s tcell.Screen) {
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			game.Mutex.Lock()
			switch ev.Key() {
			case tcell.KeyEnter:
				// ProcessCommand(game.Input)
				game.Input = ""
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if len(game.Input) > 0 {
					game.Input = game.Input[:len(game.Input)-1]
				}
			default:
				game.Input += string(ev.Rune())
			}
			game.Mutex.Unlock()
		}
	}
}

func logicLoop() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		game.Mutex.Lock()

		// 🔽 Здесь твоя логика
		game.World.DaysPassed++
		for i := range game.World.Regions {
			r := &game.World.Regions[i]
			if r.Infected > 0 && r.Infected < r.Population {
				r.Infected += r.Infected / 10 // +10% в день
				if r.Infected > r.Population {
					r.Infected = r.Population
				}
			}
		}

		game.Mutex.Unlock()
	}
}

func DrawGame(s tcell.Screen, g *GameState) {
	printText(s, 1, 0, "Day: "+strconv.Itoa(g.World.DaysPassed))
	y := 2
	for _, r := range g.World.Regions {
		printText(s, 0, y, fmt.Sprintf("%s Pop:%d Infected:%d Dead:%d", r.Name, r.Population, r.Infected, r.Dead))
		y++
	}
	printText(s, 0, y+1, "> "+g.Input)
}

func printText(s tcell.Screen, x, y int, text string) {
	for i, r := range text {
		s.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
}
