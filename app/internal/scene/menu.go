package scene

import (
	"os"
	"worldinc/app/internal/model"
	"worldinc/app/pkg/print"

	"github.com/gdamore/tcell/v2"
)

type menuScene struct {
	game *model.GameState
}

func NewMenuScene(game *model.GameState) *menuScene {
	return &menuScene{
		game: game,
	}
}

func (s *menuScene) Update() {
}

func (s *menuScene) Draw(sc tcell.Screen) {
	s.game.Mutex.Lock()

	print.Print(sc, 0, 1, "Welcome to World Inc! This is a remastered demo of my original game back in 2020.")
	print.Print(sc, 0, 2, "It was originally written in pygame, but was soon abandoned due to lack of my skills.")
	print.Print(sc, 0, 3, "You have a world map and a list of various symptoms. You should buy them to increase mortality of the disease.")
	print.Print(sc, 0, 4, "Your goal is to kill every single person on Earth before the disease will be discovered and destroyed.")

	print.Print(sc, 0, 6, "=== Controls ===")
	print.Print(sc, 0, 7, "A   - go to Symptoms store")
	print.Print(sc, 0, 8, "D   - go back to World statistics")
	print.Print(sc, 0, 9, "W   - go up on the Symptoms store")
	print.Print(sc, 0, 10, "S   - go down on the Symptoms store")
	print.Print(sc, 0, 11, "E   - buy hovered symptom on the Symptoms store")
	print.Print(sc, 0, 12, "1   - Tick speed 1s")
	print.Print(sc, 0, 13, "2   - Tick speed 500ms")
	print.Print(sc, 0, 14, "3   - Tick speed 250ms")
	print.Print(sc, 0, 15, "4   - Tick speed 166ms")
	print.Print(sc, 0, 16, "ESC - quit the game")

	print.Print(sc, 0, 18, "Press ENTER to start! ESC to quit.")

	s.game.Mutex.Unlock()
}

func (s *menuScene) HandleEvent(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		s.game.Mutex.Lock()
		switch ev.Key() {
		case tcell.KeyEscape:
			os.Exit(0)
		case tcell.KeyEnter:
			s.game.CurrentScene = NewGameScene(s.game)
		}
		s.game.Mutex.Unlock()
	}
}
