package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameOver struct {
	window fyne.Window
}

func NewGameOver(gameOver fyne.Window) *GameOver {
	gameOverScene := &GameOver{window: gameOver}
	gameOverScene.Render()
	return gameOverScene
}

func (s *GameOver) Render() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/Backgrounds/gameOver.png"))
	backgroundImage.Resize(fyne.NewSize(800, 600))
	backgroundImage.Move(fyne.NewPos(0, 0))

	btnGoMenu := widget.NewButton("Principal menu", s.backMenu)
	btnGoMenu.Resize(fyne.NewSize(150, 30))
	btnGoMenu.Move(fyne.NewPos(325, 500))

	btnRestart := widget.NewButton("Restart Game", s.restart)
	btnRestart.Resize(fyne.NewSize(150, 30))
	btnRestart.Move(fyne.NewPos(325, 460))

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, btnGoMenu, btnRestart))
}

func (s *GameOver) backMenu() {
	NewMenu(s.window)
}
func (s *GameOver) restart() {
	NewMainScene(s.window)
}
