package tui

import (
	"github.com/gdamore/tcell/v2"
)

func (ui *UI) setEventHandlers() {
	ui.table.SetSelectionChangedFunc(func(row, column int) {
		if row >= 0 && row < len(*ui.features) {
			ui.infoBar.SetText((*ui.features)[row].Description)
		} else if row == len(*ui.features) {
			ui.infoBar.SetText("Press Enter or 'q' to exit the application.")
		}
	})

	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			ui.isNormalMode = !ui.isNormalMode
			ui.updateLayout()
			return nil
		}

		if ui.isNormalMode {
			r := event.Rune()
			featureIndex := -1
			switch r {
			case '1', '+':
				featureIndex = 0
			case '2', 'ě':
				featureIndex = 1
			case '3', 'š':
				featureIndex = 2
			case '4', 'č':
				featureIndex = 3
			case '5', 'ř':
				featureIndex = 4
			case '6', 'ž':
				featureIndex = 5
			case '7', 'ý':
				featureIndex = 6
			case '8', 'á':
				featureIndex = 7
			case '9', 'í':
				featureIndex = 8
			}

			if featureIndex != -1 && featureIndex < len(*ui.features) {
				(*ui.features)[featureIndex].Toggle()
				ui.populateTable()
				return nil
			}

			if r == 'q' {
				ui.app.Stop()
				return nil
			}
		}
		return event
	})

	ui.table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Toggle feature with Enter or Space
		if event.Key() == tcell.KeyEnter || event.Rune() == ' ' {
			row, _ := ui.table.GetSelection()
			if row >= 0 && row < len(*ui.features) {
				(*ui.features)[row].Toggle()
				ui.populateTable()
			} else if row == len(*ui.features) {
				ui.app.Stop()
			}
			return nil
		}

		row, col := ui.table.GetSelection()
		// WASD navigation
		switch event.Rune() {
		case 'w':
			row--
		case 's':
			row++
		case 'a':
			col--
		case 'd':
			col++
		default:
			return event // Pass other keys to the default handler
		}

		// Boundary checks and selection
		rowCount := ui.table.GetRowCount()
		if row < 0 {
			row = 0
		} else if row >= rowCount {
			row = rowCount - 1
		}
		columnCount := ui.table.GetColumnCount()
		if col < 0 {
			col = 0
		} else if col >= columnCount {
			col = columnCount - 1
		}
		ui.table.Select(row, col)
		return nil
	})
}
