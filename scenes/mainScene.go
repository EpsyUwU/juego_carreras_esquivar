package scenes

import (
	"TurboRace/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type MainScene struct {
	window   fyne.Window
	car      *models.Car
	obstacle *models.Obstacle
	road     *models.SceneRoad
}

func NewMainScene(window fyne.Window) *MainScene {
	mainScene := &MainScene{window: window}
	mainScene.Render()
	mainScene.StarGame()
	window.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		switch e.Name {
		case fyne.KeyLeft:
			mainScene.MoveCarLeft()
		case fyne.KeyRight:
			mainScene.MoveCarRight()
		case fyne.KeyUp:
			mainScene.MoveCarUp()
		case fyne.KeyDown:
			mainScene.MoveCarDown()
		}
	})
	return mainScene
}

func (s *MainScene) Render() {
	sceneRoad1 := canvas.NewImageFromURI(storage.NewFileURI("./assets/Road/road1.png"))
	sceneRoad1.Resize(fyne.NewSize(300, 1200))
	sceneRoad1.Move(fyne.NewPos(260, 0))

	sceneRoad2 := canvas.NewImageFromURI(storage.NewFileURI("./assets/Road/road2.png"))
	sceneRoad2.Resize(fyne.NewSize(300, 1200))
	sceneRoad2.Move(fyne.NewPos(260, 0))

	sceneRoad3 := canvas.NewImageFromURI(storage.NewFileURI("./assets/Road/road3.png"))
	sceneRoad3.Resize(fyne.NewSize(300, 1200))
	sceneRoad3.Move(fyne.NewPos(260, 0))

	s.road = models.NewSceneRoad([3]*canvas.Image{sceneRoad1, sceneRoad2, sceneRoad3}, s.window)

	car := canvas.NewImageFromURI(storage.NewFileURI("./assets/Car/car0.png"))
	car.Resize(fyne.NewSize(110, 110))
	car.Move(fyne.NewPos(353, 420))

	s.car = models.NewCar(353, 420, 100, 100, car)

	obstacleImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/Obstacle/pikachu.png"))
	obstacleImage.Resize(fyne.NewSize(100, 100))
	obstacleImage.Move(fyne.NewPos(305, 10))

	s.obstacle = models.NewObstacle(305, 10, 100, 100, obstacleImage)

	roadContainer := container.NewWithoutLayout()
	s.window.SetContent(roadContainer)
	s.road.Run(roadContainer)

	s.window.SetContent(container.NewWithoutLayout(sceneRoad1, sceneRoad2, sceneRoad3, car, obstacleImage))
}

func (s *MainScene) StarGame() {
	go s.car.Run()
	go s.obstacle.Run(s.car, s.GameOver)
}

func (s *MainScene) MoveCarLeft() {
	s.car.MoveLeft(s.obstacle, s.GameOver)
}

func (s *MainScene) MoveCarRight() {
	s.car.MoveRight(s.obstacle, s.GameOver)
}

func (s *MainScene) MoveCarUp() {
	s.car.MoveUp(s.obstacle, s.GameOver)
}

func (s *MainScene) MoveCarDown() {
	s.car.MoveDown(s.obstacle, s.GameOver)
}

func (s *MainScene) GameOver() {
	gameOverScene := NewGameOver(s.window)
	gameOverScene.Render()
}
