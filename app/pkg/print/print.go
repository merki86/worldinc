package print

import "github.com/gdamore/tcell/v2"

func Print(s tcell.Screen, x, y int, text string) {
	for i, v := range text {
		s.SetContent(x+i, y, v, nil, tcell.StyleDefault)
	}
}
