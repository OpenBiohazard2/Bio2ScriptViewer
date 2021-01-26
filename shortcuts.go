package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (a *App) loadKeyboardShortcuts() {
	// keyboard shortcuts
	// ctrl+o to open file
	a.mainWin.Canvas().AddShortcut(&desktop.CustomShortcut{
		KeyName:  fyne.KeyO,
		Modifier: a.mainModKey,
	}, func(shortcut fyne.Shortcut) { a.openFileDialog() })

	// ctrl+q to quit application
	a.mainWin.Canvas().AddShortcut(&desktop.CustomShortcut{
		KeyName:  fyne.KeyQ,
		Modifier: a.mainModKey,
	}, func(shortcut fyne.Shortcut) { a.app.Quit() })

	a.mainWin.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		// close dialogs with esc key
		case fyne.KeyEscape:
			if len(a.mainWin.Canvas().Overlays().List()) > 0 {
				a.mainWin.Canvas().Overlays().Top().Hide()
			}
		}
	})
}
