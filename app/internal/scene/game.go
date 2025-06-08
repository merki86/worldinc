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
	next model.Scene
}

var textPos int = 0

func NewGameScene(game *model.GameState) *gameScene {
	// Here we reset all vars of the scene
	textPos = 0

	return &gameScene{
		game: game,
	}
}

func (g *gameScene) Update(done chan struct{}) {
	select {
	case <-done:
		return
	default:
		g.game.Mutex.Lock()
		textPos++
		g.game.Mutex.Unlock()

		time.Sleep(100 * time.Millisecond)
	}
}

func (g *gameScene) Draw(s tcell.Screen, done chan struct{}) {
	select {
	case <-done:
		return
	default:
		g.game.Mutex.Lock()
		s.Clear()
		print.Print(s, textPos, 1, "Hewoo")
		s.Show()
		g.game.Mutex.Unlock()
	}
}

func (g *gameScene) HandleEvent(s tcell.Screen) {
	event := s.PollEvent()
	switch event := event.(type) {
	case *tcell.EventKey:
		g.game.Mutex.Lock()
		switch event.Key() {
		case tcell.KeyEscape:
			os.Exit(0)
		case tcell.KeyEnter:
			g.next = NewSampleScene(g.game)
		}
		g.game.Mutex.Unlock()
	}
}

func (g *gameScene) Next() model.Scene {
	return g.next
}
