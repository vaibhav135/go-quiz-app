package gui

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const refreshInterval = 500 * time.Millisecond

var (
	view *tview.Box
	app  *tview.Application
)

var duration int
var prevElapsed = 0
var start = time.Now()

func printRemainingTime(elapsed int) string {
	timeRemaining := duration - elapsed
	minutes := timeRemaining / 60
	seconds := timeRemaining - minutes*60
	return fmt.Sprintf("Time Remaining: %d:%d", minutes, seconds)
}

func drawTime(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	elapsed := int(time.Since(start).Seconds())
	if elapsed != prevElapsed {
		prevElapsed = elapsed
	}

	if elapsed >= duration {
		tview.Print(screen, "Times Up!!!", x, height/2, width, tview.AlignCenter, tcell.ColorLime)
	} else {
		tview.Print(screen, printRemainingTime(elapsed), x, height/2, width, tview.AlignCenter, tcell.ColorLime)
	}
	return 0, 0, 0, 0
}

func refresh() {
	tick := time.NewTicker(refreshInterval)
	for {
		<-tick.C
		app.Draw()
	}
}

func timer(durationArg int) tview.Primitive {
	duration = durationArg

	view = tview.NewBox().SetDrawFunc(drawTime)

	go refresh()

	return view
}
