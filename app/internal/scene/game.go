package scene

import (
	"os"
	"time"
	"worldinc/app/internal/model"
	"worldinc/app/pkg/print"

	"github.com/gdamore/tcell/v2"
)

type gameScene struct {
	game *model.GameState
}

var textPos int = 0

func NewGameScene(game *model.GameState) *gameScene {
	// Here we reset all vars of the scene
	textPos = 0

	return &gameScene{
		game: game,
	}
}

func (s *gameScene) Update(dt time.Duration) {
	s.game.Mutex.Lock()
	textPos++
	s.game.Mutex.Unlock()

	time.Sleep(100 * time.Millisecond)
}

func (s *gameScene) Draw(sc tcell.Screen) {
	s.game.Mutex.Lock()
	sc.Clear()
	print.Print(sc, textPos, 1, "Hewoo")
	sc.Show()
	s.game.Mutex.Unlock()
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
