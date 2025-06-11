package scene

import (
	"fmt"
	"os"
	"time"
	"worldinc/app/internal/model"
	"worldinc/app/pkg/print"

	"github.com/gdamore/tcell/v2"
)

type resultScene struct {
	game  *model.GameState
	isWin bool
}

func NewResultScene(game *model.GameState, isWin bool) *resultScene {
	return &resultScene{
		game:  game,
		isWin: isWin,
	}
}

func (s *resultScene) Update(t *time.Ticker) {
	if s.game.World.Speed != time.Second {
		s.game.World.Speed = time.Second
		t.Reset(s.game.World.Speed)
	}
}

func (s *resultScene) Draw(sc tcell.Screen) {
	s.game.Mutex.Lock()
	w := &s.game.World

	if s.isWin {
		print.Print(sc, 0, 1, "YOU WON!")
		print.Print(sc, 0, 2, fmt.Sprintf("It took you: %v to perish every single person on the Earth. Good job!", w.DaysPassed))
	} else {
		print.Print(sc, 0, 1, "GAME OVER!")
		print.Print(sc, 0, 2, fmt.Sprintf("The disease was destroyed. It had existed for %v days. %v people left.", w.DaysPassed, w.Healthy))
	}
	print.Print(sc, 0, 3, "Thank you for playing the game! Press any key to quit.")
	s.game.Mutex.Unlock()
}

func (s *resultScene) HandleEvent(ev tcell.Event) {
	if _, ok := ev.(*tcell.EventKey); ok {
		os.Exit(0)
	}
}
