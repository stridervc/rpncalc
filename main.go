package main

import (
	"github.com/nsf/termbox-go"
	"strconv"
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

	// print help at top
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	tbprint(0, 0, termbox.ColorWhite, termbox.ColorDefault, "Press 'q' to quit")

	// print stack
	y := height - 2 - len(stack)
	for _, num := range stack {
		tbprint(0, y, termbox.ColorWhite, termbox.ColorDefault, strconv.FormatFloat(num, 'g', -1, 64))
		y++
	}

	// print cursor followed by current buffer
	tbprint(0, height-1, termbox.ColorWhite|termbox.AttrBold, termbox.ColorDefault, ":")
	tbprintRunes(2, height-1, termbox.ColorWhite, termbox.ColorDefault, buffer)

	termbox.Flush()
}

func push(num float64) {
	stack = append(stack, num)
}

func pop() float64 {
	if len(stack) == 0 {
		return 0
	}

	num := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	return num
}

var buffer []rune
var stack []float64

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

			case '+':
				if len(buffer) == 0 {
					push(pop() + pop())
				} else {
					num, err := strconv.ParseFloat(string(buffer), 64)
					if err == nil {
						push(pop() + num)
						buffer = nil
					}
				}

			case '-':
				if len(buffer) == 0 {
					t := pop()
					push(pop() - t)
				} else {
					num, err := strconv.ParseFloat(string(buffer), 64)
					if err == nil {
						push(pop() - num)
						buffer = nil
					}
				}

			case '*':
				if len(buffer) == 0 {
					push(pop() * pop())
				} else {
					num, err := strconv.ParseFloat(string(buffer), 64)
					if err == nil {
						push(pop() * num)
						buffer = nil
					}
				}

			case '/':
				if len(buffer) == 0 {
					t := pop()
					push(pop() / t)
				} else {
					num, err := strconv.ParseFloat(string(buffer), 64)
					if err == nil {
						push(pop() / num)
						buffer = nil
					}
				}

			case 0:
				switch ev.Key {
				case termbox.KeyBackspace, termbox.KeyBackspace2:
					if len(buffer) > 0 {
						buffer = buffer[:len(buffer)-1]
					}

				case termbox.KeyEnter:
					// if buffer is empty, push last number again
					if len(buffer) == 0 {
						if len(stack) > 0 {
							t := pop()
							push(t)
							push(t)
						}
					} else {
						num, err := strconv.ParseFloat(string(buffer), 64)
						if err == nil {
							push(num)
							buffer = nil
						}
					}
				}

			default:
				buffer = append(buffer, ev.Ch)
			}
		}
	}
}
