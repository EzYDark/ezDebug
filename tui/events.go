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
			}

			if featureIndex != -1 {
				(*ui.features)[featureIndex].Enabled = !(*ui.features)[featureIndex].Enabled
				if (*ui.features)[featureIndex].Action != nil {
					(*ui.features)[featureIndex].Action((*ui.features)[featureIndex].Enabled)
				}
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
		if event.Key() == tcell.KeyEnter {
			row, _ := ui.table.GetSelection()
			if row >= 0 && row < len(*ui.features) {
				featureIndex := row
				(*ui.features)[featureIndex].Enabled = !(*ui.features)[featureIndex].Enabled
				if (*ui.features)[featureIndex].Action != nil {
					(*ui.features)[featureIndex].Action((*ui.features)[featureIndex].Enabled)
				}
				ui.populateTable()
			} else if row == len(*ui.features) {
				ui.app.Stop()
			}
			return nil
		}
		return event
	})
}
