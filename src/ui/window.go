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
	window.SetStyleSheet("QWidget { background-color: #2e2e2e; }")

	// Creating a vertical layout to hold the widgets
	layout := widgets.NewQVBoxLayout()
	
	searchField := widgets.NewQLineEdit(nil)
	searchField.SetPlaceholderText(searchFieldPlaceholder)

	// Create a label to show the search input
	label := widgets.NewQLabel(nil, 0)

	searchField.ConnectTextChanged(func(text string) {
		updateLabelTextEvent(label, text)
	})

	layout.AddWidget(searchField, 0, 0)
	layout.AddWidget(label, 0, 0)

	// Set the layout to the central widget of the window
	container := widgets.NewQWidget(nil, 0)
	container.SetLayout(layout)
	window.SetCentralWidget(container)

	
	window.Show()

	app.Exec()

}