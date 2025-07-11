package tui

import (
	"io"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// DebugTUI is the main TUI application.
type DebugTUI struct {
	*tview.Application
	LogView *tview.TextView
}

// Init creates a new TUI application.
func Init() *DebugTUI {
	app := tview.NewApplication()
	logView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	logView.SetChangedFunc(func() {
		logView.ScrollToEnd()
		app.Draw()
	})

	logView.SetBorder(true).SetTitle("Logs")
	logView.SetBackgroundColor(tcell.ColorDefault)

	return &DebugTUI{
		Application: app,
		LogView:     logView,
	}
}

// GetLogWriter returns an io.Writer that streams log data into the log view.
func (a *DebugTUI) GetLogWriter() io.Writer {
	return tview.ANSIWriter(a.LogView)
}

// Start starts the TUI application. (Blocking call)
func (a *DebugTUI) Start() error {
	ui := InitUI(a)
	return ui.Run()
}
