package main

import (
	"github.com/nsf/termbox-go"
)

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func tbprintRunes(x, y int, fg, bg termbox.Attribute, msg []rune) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func drawScreen() {
	_, height := termbox.Size()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	tbprint(0, 0, termbox.ColorWhite, termbox.ColorDefault, "Press 'q' to quit")

	tbprint(0, height-1, termbox.ColorWhite|termbox.AttrBold, termbox.ColorDefault, ":")
	tbprintRunes(2, height-1, termbox.ColorWhite, termbox.ColorDefault, buffer)

	termbox.Flush()
}

var buffer []rune

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	done := false
	for !done {
		drawScreen()
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch ev.Ch {
			case 'q':
				done = true

			case 0:
				switch ev.Key {
				case termbox.KeyBackspace, termbox.KeyBackspace2:
					if len(buffer) > 0 {
						buffer = buffer[:len(buffer)-1]
					}
				}

			default:
				buffer = append(buffer, ev.Ch)
			}
		}
	}
}
