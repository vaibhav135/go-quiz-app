package gui

import (
	"github.com/rivo/tview"
	"github.com/vaibhav135/go-quiz-app/pkg/quiz"
)

// Ref: https://github.com/rivo/tview/wiki/Grid
func QuizGUI(duration int, quizContent []quiz.QuizContent) {
	app = tview.NewApplication()

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	timerPrimitive := func() tview.Primitive {
		return timer(duration)
	}

	header := timerPrimitive()
	main := newPrimitive("")
	footer := newPrimitive("")

	grid := tview.NewGrid().
		SetRows(6, 0, 5).
		SetBorders(true)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(header, 0, 0, 1, 3, 0, 0, false).
		AddItem(main, 1, 0, 1, 3, 0, 0, false).
		AddItem(footer, 2, 0, 1, 3, 0, 0, false)

	app.SetRoot(grid, true).SetFocus(grid)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
