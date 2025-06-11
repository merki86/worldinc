package scene

import (
	"fmt"
	"os"
	"time"
	"worldinc/app/internal/model"
	"worldinc/app/pkg/print"

	"github.com/gdamore/tcell/v2"
)

type sampleScene struct {
	game *model.GameState
}

var textPos int = 0

func NewSampleScene(game *model.GameState) *sampleScene {
	// Here we reset all vars of the scene
	textPos = 0

	return &sampleScene{
		game: game,
	}
}

func (s *sampleScene) Update(t *time.Ticker) {
	s.game.Mutex.Lock()
	textPos++
	s.game.Mutex.Unlock()
}

func (s *sampleScene) Draw(sc tcell.Screen) {
	s.game.Mutex.Lock()
	for i := 1; i < 5; i++ {
		print.Print(sc, textPos-i, i, fmt.Sprintf("%v.sample.%v", textPos, textPos))
	}
	for i := 1; i < 5; i++ {
		print.Print(sc, textPos+i-5, i+4, fmt.Sprintf("%v.sample.%v", textPos, textPos))
	}
	s.game.Mutex.Unlock()
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
