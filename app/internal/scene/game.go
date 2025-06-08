package scene

import (
	"os"
	"time"
	"worldinc/app/internal/model"

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
