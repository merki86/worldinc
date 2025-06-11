package scene

import (
	"fmt"
	"os"
	"strings"
	"worldinc/app/internal/logic"
	"worldinc/app/internal/model"
	"worldinc/app/pkg/print"

	"github.com/gdamore/tcell/v2"
)

type gameScene struct {
	game *model.GameState
}

func NewGameScene(game *model.GameState) *gameScene {
	return &gameScene{
		game: game,
	}
}

func (s *gameScene) Update() {
	s.game.Mutex.Lock()
	w := &s.game.World
	logic.DoWorldTick(w)
	s.game.Mutex.Unlock()
}

func (s *gameScene) Draw(sc tcell.Screen) {
	s.game.Mutex.Lock()
	w := &s.game.World

	print.Print(sc, 0, 1, fmt.Sprintf("DAY: %v | Credit: %v", w.DaysPassed, w.Credit))
	print.Print(sc, 0, 2, "=== World ===")
	print.Print(sc, 0, 3, fmt.Sprintf("Healthy: %v", w.Healthy))
	print.Print(sc, 0, 4, fmt.Sprintf("Infected: %v +%v / Dead: %v +%v", w.Infected, w.NewInfected, w.Dead, w.NewDead))

	print.Print(sc, 0, 5, "=== Disease ===")
	print.Print(sc, 0, 6, fmt.Sprintf("Name: %v", w.Disease.Name))
	print.Print(sc, 0, 7, fmt.Sprintf("Mortality: %v", w.Disease.Mortality))
	print.Print(sc, 0, 8, fmt.Sprintf("Transmission: %v", w.Disease.Transmission))
	print.Print(sc, 0, 9, fmt.Sprintf("Discovered: %v", w.Disease.Discovered))

	print.Print(sc, 0, 10, "=== Regions ===")
	row := 11
	for i, v := range w.Regions {
		print.Print(sc, 0, row, fmt.Sprintf("%v. %v", i+1, v.Name))
		print.Print(sc, 0, row+1, fmt.Sprintf("   Population: %v", v.Population))
		print.Print(sc, 0, row+2, fmt.Sprintf("   Infected: %v / Dead: %v", v.Infected, v.Dead))
		row += 3
	}

	var unlocked = []model.Symptom{}
	for _, v := range s.game.Symptoms {
		if v.Unlocked {
			unlocked = append(unlocked, v)
		}
	}

	print.Print(sc, 0, row, "=== Symptoms ===")
	for i, v := range unlocked {
		print.Print(sc, 0, row+1, fmt.Sprintf("%v. %v = $%v MT / TR bonus: %v / %v", i+1, v.Name, v.Cost, v.MortalityBonus, v.TransmissionBonus))
		row += 1
	}
	s.game.Mutex.Unlock()
}

func (s *gameScene) HandleEvent(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		s.game.Mutex.Lock()
		switch ev.Key() {
		case tcell.KeyEscape:
			os.Exit(0)
		case tcell.KeyTab:
			s.game.CurrentScene = NewSampleScene(s.game)
		case tcell.KeyRune:
			switch strings.ToLower(string(ev.Rune())) {
			case "a":
				s.game.CurrentScene = NewSymptomsScene(s.game)
			}
		}
		s.game.Mutex.Unlock()
	}
}
