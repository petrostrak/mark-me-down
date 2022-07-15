package main

import (
	"io"

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
			// keep track of what the current file is
			c.CurrentFile = w.URI()

			defer w.Close()

			win.SetTitle(win.Title() + " + " + w.URI().Name())
			c.SaveMenuItem.Disabled = false
		}, win)

		saveDialog.Show()
	}
}

func (c *config) openFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if r == nil {
				// user cancelled
				return
			}

			defer r.Close()

			data, err := io.ReadAll(r)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			c.EditWidget.SetText(string(data))

			// keep track of what the current file is
			c.CurrentFile = r.URI()

			// update window title
			win.SetTitle(win.Title() + " + " + r.URI().Name())
			c.SaveMenuItem.Disabled = false
		}, win)

		openDialog.Show()
	}
}
