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
	var num float64 // current number

	num = 0

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

			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				num = num*10 + float64(ev.Ch-'0')
				buffer = append(buffer, ev.Ch)

			case '+':
				push(pop() + pop())

			case '-':
				t := pop()
				push(pop() - t)

			case '*':
				push(pop() * pop())

			case '/':
				t := pop()
				push(pop() / t)

			case 0:
				switch ev.Key {
				case termbox.KeyBackspace, termbox.KeyBackspace2:
					if len(buffer) > 0 {
						buffer = buffer[:len(buffer)-1]
					}

				case termbox.KeyEnter:
					// if buffer is empty, push the last value again
					if len(buffer) == 0 {
						num = pop()
						push(num)
					}
					push(num)
					num = 0
					buffer = nil
				}

			default:
				buffer = append(buffer, ev.Ch)
			}
		}
	}
}
