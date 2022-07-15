package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func (c *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(w fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if w == nil {
				// user cancelled
				return
			}

			// save file
			w.Write([]byte(c.EditWidget.Text))
			c.CurrentFile = w.URI()

			defer w.Close()

			win.SetTitle(win.Title() + " + " + w.URI().Name())
			c.SaveMenuItem.Disabled = false
		}, win)

		saveDialog.Show()
	}
}
