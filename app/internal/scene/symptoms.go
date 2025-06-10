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

type symptomsScene struct {
	game *model.GameState
}

func NewSymptomsScene(game *model.GameState) *symptomsScene {
	return &symptomsScene{
		game: game,
	}
}

func (s *symptomsScene) Update() {
	world := &s.game.World

	logic.DoWorldTick(world)
}

func (s *symptomsScene) Draw(sc tcell.Screen) {
	// world := &s.game.World
	symptomsList := &s.game.Symptoms

	print.Print(sc, 0, 1, "=== Symptoms store ===")
	row := 2
	for i, v := range *symptomsList {
		print.Print(sc, 0, row, fmt.Sprintf("%v. %v = $%v [%v]", i+1, v.Name, v.Cost, v.Unlocked))
		print.Print(sc, 0, row+1, fmt.Sprintf("   MT / TR bonus: %v / %v", v.MortalityBonus, v.TransmissionBonus))
		row += 2
	}
}

func (s *symptomsScene) HandleEvent(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		s.game.Mutex.Lock()
		switch ev.Key() {
		case tcell.KeyEscape:
			os.Exit(0)
		case tcell.KeyRune:
			switch strings.ToLower(string(ev.Rune())) {
			case "d":
				s.game.CurrentScene = NewGameScene(s.game)
			}
		}
		s.game.Mutex.Unlock()
	}
}
