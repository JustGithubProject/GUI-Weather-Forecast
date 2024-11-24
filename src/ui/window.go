package ui

import (
	"github.com/therecipe/qt/widgets"
)

func RunWindow(lenOfArgs int, args []string) {
	
	// Initialize Qt application
	app := widgets.NewQApplication(lenOfArgs, args)

	// Create a main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle(windowTitle)
	window.SetMinimumSize2(windowWidth, windowHeight)

	window.Show()

	app.Exec()

}