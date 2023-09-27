package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Car struct {
	posX, posY          float32
	status              bool
	widthCar, heightCar float32
	car                 *canvas.Image
}

func NewCar(posx float32, posy float32, width float32, height float32, img *canvas.Image) *Car {
	car := &Car{
		posX:      posx,
		posY:      posy,
		widthCar:  width,
		heightCar: height,
		status:    true,
		car:       img,
	}

	return car
}

func (a *Car) Run() {
	a.status = true
}

func (a *Car) SetStatus(status bool) {
	a.status = status
}

func (a *Car) GetStatus() bool {
	return a.status
}

func (a *Car) MoveLeft(o *Obstacle, onCollision func()) {
	if a.posX > 290 {
		a.posX -= 10
		a.car.Move(fyne.NewPos(a.posX, a.posY))

		if a.CheckCollision(o) {
			a.SetStatus(false)
			o.SetStatus(false)
			onCollision()
		}
	}
}

func (a *Car) MoveRight(o *Obstacle, onCollision func()) {
	if a.posX < 430 {
		a.posX += 10
		a.car.Move(fyne.NewPos(a.posX, a.posY))

		if a.CheckCollision(o) {
			a.SetStatus(false)
			o.SetStatus(false)
			onCollision()
		}
	}
}

func (a *Car) MoveDown(o *Obstacle, onCollision func()) {
	if a.posY < 500 {
		a.posY += 10
		a.car.Move(fyne.NewPos(a.posX, a.posY))

		if a.CheckCollision(o) {
			a.SetStatus(false)
			o.SetStatus(false)
			onCollision()
		}
	}
}

func (a *Car) MoveUp(o *Obstacle, onCollision func()) {
	if a.posY > 10 {
		a.posY -= 10
		a.car.Move(fyne.NewPos(a.posX, a.posY))

		if a.CheckCollision(o) {
			a.SetStatus(false)
			o.SetStatus(false)
			onCollision()
		}
	}
}

func (a *Car) CheckCollision(o *Obstacle) bool {
	if a.posY >= o.posY && a.posY <= o.posY+o.heightObs &&
		a.posX >= o.posX && a.posX <= o.posX+o.widthObs {
		println("se choco")
		return true
	}
	return false
}
