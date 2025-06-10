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
	world := &s.game.World
	logic.DoWorldTick(world)
	s.game.Mutex.Unlock()
}

func (s *gameScene) Draw(sc tcell.Screen) {
	s.game.Mutex.Lock()
	world := &s.game.World

	print.Print(sc, 0, 1, fmt.Sprintf("DAY: %v === World ===", world.DaysPassed))
	print.Print(sc, 0, 2, fmt.Sprintf("Healthy: %v", world.Healthy))
	print.Print(sc, 0, 3, fmt.Sprintf("Infected: %v +%v / Dead: %v +%v", world.Infected, world.NewInfected, world.Dead, world.NewDead))

	print.Print(sc, 0, 4, "=== Disease ===")
	print.Print(sc, 0, 5, fmt.Sprintf("Name: %v", world.Disease.Name))
	print.Print(sc, 0, 6, fmt.Sprintf("Mortality: %v", world.Disease.Mortality))
	print.Print(sc, 0, 7, fmt.Sprintf("Transmission: %v", world.Disease.Transmission))
	print.Print(sc, 0, 8, fmt.Sprintf("Discovered: %v", world.Disease.Discovered))

	print.Print(sc, 0, 9, "=== Regions ===")
	row := 10
	for i, v := range world.Regions {
		print.Print(sc, 0, row, fmt.Sprintf("%v. %v", i+1, v.Name))
		print.Print(sc, 0, row+1, fmt.Sprintf("   Population: %v", v.Population))
		print.Print(sc, 0, row+2, fmt.Sprintf("   Infected: %v / Dead: %v", v.Infected, v.Dead))
		row += 3
	}

	print.Print(sc, 0, row, "=== Symptoms ===")
	for i, v := range world.Disease.Symptoms {
		print.Print(sc, 0, row+1, fmt.Sprintf("%v. %v = $%v [%v]", i+1, v.Name, v.Cost, v.Unlocked))
		print.Print(sc, 0, row+2, fmt.Sprintf("   MT / TR bonus: %v / %v", v.MortalityBonus, v.TransmissionBonus))
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
