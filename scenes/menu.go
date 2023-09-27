package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type Menu struct {
	window fyne.Window
}

func NewMenu(windowMenu fyne.Window) *Menu {
	scene := &Menu{window: windowMenu}
	scene.Render()
	return scene
}

func (s *Menu) Render() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/Backgrounds/background1.jpg"))
	backgroundImage.Resize(fyne.NewSize(800, 600))
	backgroundImage.Move(fyne.NewPos(0, 50))

	botonIniciar := widget.NewButton("Start Game", s.StartGame)
	botonIniciar.Resize(fyne.NewSize(180, 50))
	botonIniciar.Move(fyne.NewPos(310, 290))

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, botonIniciar))
}

func (s *Menu) StartGame() {
	NewMainScene(s.window)
}
