package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// App represents the whole application with all its windows, widgets and functions
type App struct {
	app     fyne.App
	mainWin fyne.Window

	mainModKey desktop.Modifier

	split               *container.Split
	rawScriptData       *widget.Entry
	convertedScriptCode *widget.Entry

	fileListBar *widget.List
	statusBar   *fyne.Container

	fullscreenWin fyne.Window
}

func (a *App) init() {
	// theme
	switch a.app.Preferences().StringWithFallback("Theme", "Dark") {
	case "Light":
		a.app.Settings().SetTheme(theme.LightTheme())
	case "Dark":
		a.app.Settings().SetTheme(theme.DarkTheme())
	}

	// show/hide statusbar
	if a.app.Preferences().BoolWithFallback("statusBarVisible", true) == false {
		a.statusBar.Hide()
	}
}

func main() {
	a := app.NewWithID("bio2-scd-viewer")
	w := a.NewWindow("Biohazard 2 Script Viewer")
	ui := &App{app: a, mainWin: w}
	ui.init()
	w.SetContent(ui.loadMainUI())
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("Error while opening the file: %v\n", err)
		}
		ui.open(file, true)
	}
	w.Resize(fyne.NewSize(1200, 750))
	w.ShowAndRun()
}
