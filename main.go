package main

import (
	"TurboRace/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TurboRace")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(800, 600))

	menu := scenes.NewMenu(myWindow)
	menu.Render()

	myWindow.ShowAndRun()
}
