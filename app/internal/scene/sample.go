package scene

import (
	"os"
	"worldinc/app/internal/model"

	"github.com/gdamore/tcell/v2"
)

type sampleScene struct {
	next model.Scene
}

func NewSampleScene() *sampleScene {
	return &sampleScene{}
}

func (g *sampleScene) Update() {
	// Update
}

func (g *sampleScene) Draw(s tcell.Screen) {
	s.Fill('s', tcell.StyleDefault)
	s.Show()
}

func (g *sampleScene) Next() model.Scene {
	return g.next
}

func (g *sampleScene) HandleEvent(game *model.GameState, s tcell.Screen) {
	event := s.PollEvent()
	switch event := event.(type) {
	case *tcell.EventKey:
		game.Mutex.Lock()
		switch event.Key() {
		case tcell.KeyEscape:
			os.Exit(0)
		case tcell.KeyEnter:
			game.CurrentScene = g.next
		}
		game.Mutex.Unlock()
	}
}
