package scene

import (
	"os"
	"worldinc/app/internal/model"

	"github.com/gdamore/tcell/v2"
)

type gameScene struct {
	next model.Scene
}

func NewGameScene() *gameScene {
	return &gameScene{}
}

func (g *gameScene) Update() {
	// Update
}

func (g *gameScene) Draw(s tcell.Screen) {
	s.SetContent(0, 0, '5', []rune{'a', 's', 'a', 'p', '!'}, tcell.StyleDefault)
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
