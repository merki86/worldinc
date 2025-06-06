package window

import (
	"fmt"
	"strings"
)

type Window struct {
	Title   string
	Buttons []Button
}

type Button struct {
	Label    string
	Literal  string
	Callback func()
}

func (w Window) New() {
	gap := 5
	content := w.Title + strings.Repeat(" ", gap)

	var btnLines []string
	for _, v := range w.Buttons {
		btnLines = append(btnLines, fmt.Sprintf("[%v] %v", v.Literal, v.Label))
	}
	content = "║ " + strings.Join([]string{content, strings.Join(btnLines, " ")}, "") + " ║"

	upperBorder := "╔" + strings.Repeat("═", len(content)-6) + "╗"
	downBorder := "╚" + strings.Repeat("═", len(content)-6) + "╝"
	header := strings.Join([]string{upperBorder, content, downBorder}, "\n")

	fmt.Printf("%v", header)
}
