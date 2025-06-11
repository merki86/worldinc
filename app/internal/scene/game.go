package scene

import (
	"fmt"
	"os"
	"strings"
	"time"
	"worldinc/app/internal/logic"
	"worldinc/app/internal/model"
	"worldinc/app/pkg/print"

	"github.com/gdamore/tcell/v2"
)

type gameScene struct {
	game *model.GameState
	t    *time.Ticker
}

func NewGameScene(game *model.GameState) *gameScene {
	return &gameScene{
		game: game,
	}
}

func (s *gameScene) Update(t *time.Ticker) {
	s.game.Mutex.Lock()
	w := &s.game.World
	logic.DoWorldTick(w)
	s.t = t

	if w.Infected <= 0 {
		if w.Healthy <= 0 {
			s.game.CurrentScene = NewResultScene(s.game, true)
		} else {
			s.game.CurrentScene = NewResultScene(s.game, false)
		}
	}
	s.game.Mutex.Unlock()
}

func (s *gameScene) Draw(sc tcell.Screen) {
	s.game.Mutex.Lock()
	w := &s.game.World

	print.Print(sc, 0, 1, fmt.Sprintf("DAY: %v | Speed: %v | Credit: %v", w.DaysPassed, w.Speed, w.Credit))
	print.Print(sc, 0, 2, "=== World ===")
	print.Print(sc, 0, 3, fmt.Sprintf("Healthy: %v", w.Healthy))
	print.Print(sc, 0, 4, fmt.Sprintf("Infected: %v +%v / Dead: %v +%v", w.Infected, w.NewInfected, w.Dead, w.NewDead))

	print.Print(sc, 0, 5, "=== Disease ===")
	print.Print(sc, 0, 6, fmt.Sprintf("Name: %v", w.Disease.Name))
	print.Print(sc, 0, 7, fmt.Sprintf("Mortality: %v", w.Disease.Mortality))
	print.Print(sc, 0, 8, fmt.Sprintf("Transmission: %v", w.Disease.Transmission))

	var unlocked = []model.Symptom{}
	for _, v := range s.game.Symptoms {
		if v.Unlocked {
			unlocked = append(unlocked, v)
		}
	}

	row := 9
	print.Print(sc, 0, row, "=== Symptoms ===")
	print.Print(sc, 0, row+1, "No symptoms")
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
			case "1":
				s.game.World.Speed = time.Second
				s.t.Reset(s.game.World.Speed)
			case "2":
				s.game.World.Speed = time.Second / 2
				s.t.Reset(s.game.World.Speed)
			case "3":
				s.game.World.Speed = time.Second / 4
				s.t.Reset(s.game.World.Speed)
			case "4":
				s.game.World.Speed = time.Second / 6
				s.t.Reset(s.game.World.Speed)
			}
		}
		s.game.Mutex.Unlock()
	}
}
