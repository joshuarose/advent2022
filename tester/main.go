package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/msoap/tcg"
)

func main() {
	tg, err := tcg.New(tcg.Mode2x3) // each terminal symbol contains a 2x3 pixels grid, also you can use 1x1, 1x2, and 2x2 modes
	if err != nil {
		log.Fatalf("create tg: %s", err)
	}

	i := 0
	for {
		pixColor := tg.Buf.At(10, 10)       // get color of pixel
		tg.Buf.Set(11, 11, pixColor)        // draw one pixel with color from 10,10
		tg.Buf.Line(0, 0, 50, i, tcg.Black) // draw a diagonal line
		tg.Show()                           // synchronize buffer with screen

		if ev, ok := tg.TCellScreen.PollEvent().(*tcell.EventKey); ok && ev.Rune() == 'q' {
			break // exit by 'q' key
		}
		i++
	}

	tg.Finish() // finish application and restore screen
}
