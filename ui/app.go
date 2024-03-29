package ui

import (
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
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

func (a *App) loadStatusBar() *fyne.Container {
	a.statusBar = container.NewVBox(
		widget.NewSeparator(),
		container.NewHBox(
			layout.NewSpacer(),
		))
	return a.statusBar
}

func (a *App) loadFileList(filenames []string, scriptFiles map[string][][]byte) *widget.List {
	data := filenames

	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(data[id])
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(data[id])
		icon.SetResource(theme.DocumentIcon())

		if scriptFiles != nil {
			a.rawScriptData.SetText(convertRawScriptInstructionsToString(scriptFiles[filenames[id]]))
			a.convertedScriptCode.SetText(convertScriptInstructionsToCode(scriptFiles[filenames[id]]))
		}
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}
	// Select first item at the top
	list.Select(0)

	a.fileListBar = list
	return a.fileListBar
}

func (a *App) loadMainUI() fyne.CanvasObject {
	a.mainWin.SetMaster()
	// set main mod key to super on darwin hosts, else set it to ctrl
	if runtime.GOOS == "darwin" {
		a.mainModKey = desktop.SuperModifier
	} else {
		a.mainModKey = desktop.ControlModifier
	}

	// main menu
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Open", a.openFileDialog),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				dialog.ShowCustom("About", "Ok", container.NewVBox(
					widget.NewLabel("Original Resident Evil 2 / Biohazard 2 Script Viewer."),
				), a.mainWin)
			}),
		),
	)
	a.mainWin.SetMainMenu(mainMenu)

	a.loadKeyboardShortcuts()

	a.rawScriptData = widget.NewMultiLineEntry()
	a.rawScriptData.Wrapping = fyne.TextWrapWord
	a.rawScriptData.SetText("")

	a.convertedScriptCode = widget.NewMultiLineEntry()
	a.convertedScriptCode.Wrapping = fyne.TextWrapWord
	a.convertedScriptCode.SetText("")

	a.split = container.NewHSplit(
		a.rawScriptData,
		a.convertedScriptCode,
	)
	a.split.SetOffset(0.50)
	layout := container.NewBorder(nil, a.loadStatusBar(), a.loadFileList([]string{}, nil), nil, a.split)
	return layout
}

func RunApp() {
	curApp := app.NewWithID("bio2-scd-viewer")
	mainWindow := curApp.NewWindow("Biohazard 2 Script Viewer")
	userInterface := &App{app: curApp, mainWin: mainWindow}
	userInterface.init()
	mainWindow.SetContent(userInterface.loadMainUI())
	mainWindow.Resize(fyne.NewSize(1200, 750))
	mainWindow.ShowAndRun()
}
