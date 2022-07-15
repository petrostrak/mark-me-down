package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

func (c *config) makeUI() (edit *widget.Entry, preview *widget.RichText) {
	edit = widget.NewMultiLineEntry()
	preview = widget.NewRichTextFromMarkdown("")

	c.EditWidget = edit
	c.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return
}

func (c *config) createMenuItems(win fyne.Window) {
	openMenuItem := fyne.NewMenuItem("Open...", c.openFunc(win))

	saveMenuItem := fyne.NewMenuItem("Save", func() {})
	c.SaveMenuItem = saveMenuItem
	c.SaveMenuItem.Disabled = true

	saveAsMenuItem := fyne.NewMenuItem("Save as...", c.saveAsFunc(win))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}
