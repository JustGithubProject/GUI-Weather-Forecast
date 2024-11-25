package ui

import (
	"github.com/therecipe/qt/widgets"
)

func updateLabelTextEvent(label *widgets.QLabel, text string) {
	label.SetText("Searching for: " + text)
}