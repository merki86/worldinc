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
	game       *model.GameState
	hovered    int
	message    string
	isSelected bool
}

func NewSymptomsScene(game *model.GameState) *symptomsScene {
	return &symptomsScene{
		game:       game,
		hovered:    1,     // cause the hovered item should be the very first one
		isSelected: false, // TODO: a []int to select multiple
	}
}

func (s *symptomsScene) Update() {
	s.game.Mutex.Lock()
	w := &s.game.World

	logic.DoWorldTick(w)
	if s.isSelected {
		i, err := logic.BuySymptom(s.hovered, s.game)
		if err != nil {
			s.message = err.Error()
		} else {
			logic.ApplySymptom(i, w)
		}
		s.isSelected = false
	}
	s.game.Mutex.Unlock()
}

func (s *symptomsScene) Draw(sc tcell.Screen) {
	s.game.Mutex.Lock()
	w := &s.game.World
	symptomsList := &s.game.Symptoms

	print.Print(sc, 0, 1, "=== Symptoms store ===")
	row := 2
	for i, v := range *symptomsList {
		if v.ID == s.hovered {
			print.Print(sc, 0, row, fmt.Sprintf("[x] %v. %v - $%v [ID: %v] status: %v", i+1, v.Name, v.Cost, v.ID, v.Unlocked))
			if s.isSelected {
				print.Print(sc, 0, row, fmt.Sprintf("[v] %v. %v - $%v [ID: %v] status: %v", i+1, v.Name, v.Cost, v.ID, v.Unlocked))
			}
		} else {
			print.Print(sc, 0, row, fmt.Sprintf("[ ] %v. %v - $%v [ID: %v] status: %v", i+1, v.Name, v.Cost, v.ID, v.Unlocked))
		}
		print.Print(sc, 0, row+1, fmt.Sprintf("    MT / TR bonus: %v / %v", v.MortalityBonus, v.TransmissionBonus))
		row += 2
	}

	print.Print(sc, 0, row, fmt.Sprintf("Credit: %v", w.Credit))
	if s.isSelected {
		print.Print(sc, 0, row, fmt.Sprintf("Buying...: %v. %v", s.hovered, s.game.Symptoms[s.hovered-1].Name))
	}
	if s.message != "" {
		print.Print(sc, 0, row, fmt.Sprintf("Error: %v", s.message))
	}
	s.game.Mutex.Unlock()
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
			case "w":
				if s.hovered > 1 {
					s.hovered -= 1
					s.isSelected = false
					s.message = ""
				}
			case "s":
				if s.hovered < len(s.game.Symptoms) {
					s.hovered += 1
					s.isSelected = false
					s.message = ""
				}
			case "e":
				if s.hovered > 0 && s.hovered <= len((*s).game.Symptoms) {
					if (*s).game.Symptoms[s.hovered-1].Unlocked {
						s.isSelected = false
						s.message = "The item is already bought"
					} else {
						s.isSelected = !s.isSelected
					}
				} else {
					s.isSelected = !s.isSelected
				}
			}
		}
		s.game.Mutex.Unlock()
	}
}
