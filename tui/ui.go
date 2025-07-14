package tui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// UI represents the user interface of the TUI.
type UI struct {
	app          *DebugTUI
	table        *tview.Table
	infoBar      *tview.TextView
	rootFlex     *tview.Flex
	features     *FeatureList
	isNormalMode bool
}

// InitUI creates a new UI.
func InitUI(app *DebugTUI) *UI {
	table := tview.NewTable().SetSelectable(true, false)
	table.SetBackgroundColor(tcell.ColorDefault)

	infoBar := tview.NewTextView().SetDynamicColors(true)
	infoBar.SetBackgroundColor(tcell.ColorDefault)

	rootFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	rootFlex.SetBackgroundColor(tcell.ColorDefault)

	return &UI{
		app:      app,
		table:    table,
		infoBar:  infoBar,
		rootFlex: rootFlex,
		features: GetFeatureList(),
	}
}

// Run initializes and runs the UI.
func (ui *UI) Run() error {
	ui.table.SetBorder(true).SetTitle("Features")
	ui.infoBar.SetBorder(true).SetTitle("Info")

	// Check and start features with StartOnStartup
	for i := range *ui.features {
		feature := &(*ui.features)[i]
		if feature.StartOnStartup {
			if feature.OnStart != nil {
				feature.OnStart(feature)
			}
		}
	}

	ui.populateTable()
	ui.updateLayout()
	ui.setEventHandlers()

	fmt.Fprintln(ui.app.GetLogWriter(), "[yellow]Press 'Esc' to toggle command mode.[-]")
	ui.app.SetRoot(ui.rootFlex, true).EnableMouse(true)

	return ui.app.Application.Run()
}

func (ui *UI) updateLayout() {
	ui.rootFlex.Clear()
	if ui.isNormalMode {
		mainContent := tview.NewFlex().
			AddItem(ui.table, 0, 1, true).
			AddItem(ui.app.LogView, 0, 2, false)
		ui.rootFlex.AddItem(mainContent, 0, 1, false).
			AddItem(ui.infoBar, 3, 0, false)
		ui.app.SetFocus(ui.table)
	} else {
		ui.rootFlex.AddItem(ui.app.LogView, 0, 1, true)
		ui.app.SetFocus(ui.app.LogView)
	}
}

func (ui *UI) populateTable() {
	current_row, _ := ui.table.GetSelection()
	ui.table.Clear()
	for i, feature := range *ui.features {
		ui.table.SetCell(i, 0,
			tview.NewTableCell(fmt.Sprintf("[grey](%d)[-]", i+1)).
				SetAlign(tview.AlignCenter))
		checkbox := "[ ]"
		if feature.Enabled {
			checkbox = "[[green]X[-]]"
		}
		ui.table.SetCell(i, 1,
			tview.NewTableCell(fmt.Sprintf(" %s ", checkbox)).
				SetAlign(tview.AlignCenter))
		ui.table.SetCell(i, 2,
			tview.NewTableCell(feature.Name).
				SetExpansion(1))
	}
	ui.table.SetCell(len(*ui.features), 0, tview.NewTableCell("[grey](q)[-]").
		SetAlign(tview.AlignCenter))
	ui.table.SetCell(len(*ui.features), 1, tview.NewTableCell(""))
	ui.table.SetCell(len(*ui.features), 2, tview.NewTableCell("Quit"))
	ui.table.Select(current_row, 0)
}
