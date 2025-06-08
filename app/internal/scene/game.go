package scene

import (
	"os"
	"worldinc/app/internal/model"
	"worldinc/app/pkg/print"

	"github.com/gdamore/tcell/v2"
)

type gameScene struct {
	next model.Scene
}

func NewGameScene() *gameScene {
	return &gameScene{}
}

var textPos int = 0

func (g *gameScene) Update() {
	textPos++
}

func (g *gameScene) Draw(s tcell.Screen) {
	s.Clear()
	print.Print(s, textPos, 1, "Hewoo")
	s.Show()
}

func (g *gameScene) Next() model.Scene {
	return g.next
}

func (g *gameScene) HandleEvent(game *model.GameState, s tcell.Screen) {
	event := s.PollEvent()
	switch event := event.(type) {
	case *tcell.EventKey:
		// game.Mutex.Lock()
		switch event.Key() {
		case tcell.KeyEscape:
			os.Exit(0)
		case tcell.KeyEnter:
			g.next = NewSampleScene()

		}
		// game.Mutex.Unlock()
	}
}
