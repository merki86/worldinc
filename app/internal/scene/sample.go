package scene

import (
	"os"
	"time"
	"worldinc/app/internal/model"

	"github.com/gdamore/tcell/v2"
)

type sampleScene struct {
	game *model.GameState
}

func NewSampleScene(game *model.GameState) *sampleScene {
	return &sampleScene{
		game: game,
	}
}

func (s *sampleScene) Update(dt time.Duration) {

}

func (s *sampleScene) Draw(sc tcell.Screen) {
	sc.Fill('a', tcell.StyleDefault)
}

func (s *sampleScene) HandleEvent(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			os.Exit(0)
		case tcell.KeyEnter:
			s.game.Mutex.Lock()
			s.game.CurrentScene = NewGameScene(s.game)
			s.game.Mutex.Unlock()
		}
	}
}
