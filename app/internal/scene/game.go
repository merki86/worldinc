package scene

import (
	"fmt"
	"os"
	"time"
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

func (s *gameScene) Update(dt time.Duration) {

}

func (s *gameScene) Draw(sc tcell.Screen) {
	print.Print(sc, 0, 1, fmt.Sprintf("DAY: %v === World ===", s.game.World.DaysPassed))
	print.Print(sc, 0, 2, fmt.Sprintf("Population: %v", s.game.World.Population))
	print.Print(sc, 0, 3, fmt.Sprintf("Infected: %v / Dead: %v", s.game.World.Infected, s.game.World.Dead))

	print.Print(sc, 0, 4, "=== Disease ===")
	print.Print(sc, 0, 5, fmt.Sprintf("Name: %v", s.game.World.Disease.Name))
	print.Print(sc, 0, 6, fmt.Sprintf("Mortality: %v", s.game.World.Disease.Mortality))
	print.Print(sc, 0, 7, fmt.Sprintf("Transmission: %v", s.game.World.Disease.Transmission))
	print.Print(sc, 0, 8, fmt.Sprintf("Discovered: %v", s.game.World.Disease.Discovered))

	print.Print(sc, 0, 9, "=== Regions ===")
	row := 10
	for i, v := range s.game.World.Regions {
		print.Print(sc, 0, row, fmt.Sprintf("%v. %v", i+1, v.Name))
		print.Print(sc, 0, row+1, fmt.Sprintf("   Population: %v", v.Population))
		print.Print(sc, 0, row+2, fmt.Sprintf("   Infected: %v / Dead: %v", v.Infected, v.Dead))
		row += 3
	}

	print.Print(sc, 0, row, "=== Symptoms ===")
	for i, v := range s.game.World.Disease.Symptoms {
		print.Print(sc, 0, row+1, fmt.Sprintf("%v. %v = $%v [%v]", i+1, v.Name, v.Cost, v.Unlocked))
		print.Print(sc, 0, row+2, fmt.Sprintf("   MT / TR bonus: %v / %v", v.MortalityBonus, v.TransmissionBonus))
	}
}

func (s *gameScene) HandleEvent(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			os.Exit(0)
		case tcell.KeyEnter:
			s.game.Mutex.Lock()
			s.game.CurrentScene = NewSampleScene(s.game)
			s.game.Mutex.Unlock()
		}
	}
}
