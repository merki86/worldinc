package scene

import (
	"fmt"
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
	for y := 1; y < 3; y += 10 {
		for i := 1; i < 5; i++ {
			print.Print(sc, textPos-i-y, i, fmt.Sprintf("%v.upper.%v", textPos, textPos))
		}
		for i := 1; i < 5; i++ {
			print.Print(sc, textPos+i-5-y, i+4, fmt.Sprintf("%v.bottom.%v", textPos, textPos))
		}
	}
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
