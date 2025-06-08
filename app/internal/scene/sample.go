package scene

import (
	"os"
	"worldinc/app/internal/model"

	"github.com/gdamore/tcell/v2"
)

type sampleScene struct {
	game *model.GameState
	next model.Scene
}

func NewSampleScene(game *model.GameState) *sampleScene {
	return &sampleScene{
		game: game,
	}
}

func (g *sampleScene) Update(done chan struct{}) {
	// Update
}

func (g *sampleScene) Draw(s tcell.Screen, done chan struct{}) {

	s.Fill('s', tcell.StyleDefault)
	s.Show()

}

func (g *sampleScene) Next() model.Scene {
	return g.next
}

func (g *sampleScene) HandleEvent(s tcell.Screen) {

	event := s.PollEvent()
	switch event := event.(type) {
	case *tcell.EventKey:
		g.game.Mutex.Lock()
		switch event.Key() {
		case tcell.KeyEscape:
			os.Exit(0)
		case tcell.KeyEnter:
			g.next = NewGameScene(g.game)

		}
		g.game.Mutex.Unlock()
	}

}
