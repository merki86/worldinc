package window

import (
	"fmt"
	"strings"
)

type Window struct {
	Title           string
	ContentCallback func()
	Buttons         []Button
}

type Button struct {
	Label    string
	Callback func()
}

func (w Window) New() {
	drawHeader(w.Title, w.Buttons) // Print header of the window
	w.ContentCallback()
	handleAction(w.Buttons)
}

func drawHeader(title string, buttons []Button) {
	gap := 5
	content := title + strings.Repeat(" ", gap)

	var btnContent []string
	for i, v := range buttons {
		btnContent = append(btnContent, fmt.Sprintf("[%v] %v", i, v.Label))
	}
	content = "║ " + strings.Join([]string{content, strings.Join(btnContent, " ")}, "") + " ║"

	upperBorder := "╔" + strings.Repeat("═", len(content)-6) + "╗"
	downBorder := "╚" + strings.Repeat("═", len(content)-6) + "╝"
	header := strings.Join([]string{upperBorder, content, downBorder}, "\n") + "\n"

	fmt.Printf("%v", header)
}

func handleAction(buttons []Button) {
	var action string
	fmt.Print("> ")

	fmt.Scanln(&action)

	for i, v := range buttons {
		if action == fmt.Sprint(i) || strings.EqualFold(action, v.Label) {
			v.Callback()
		}
	}

	fmt.Println("Invalid action")
	handleAction(buttons)
}
