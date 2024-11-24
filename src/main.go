package main

import (
	"os"
	"JustGithubProject/GUI-Weather-Forecast/src/ui"
)


func main() {
	ui.RunWindow(len(os.Args), os.Args)
}