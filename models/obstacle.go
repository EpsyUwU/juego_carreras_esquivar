package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"math/rand"
	"time"
)

type Obstacle struct {
	posX, posY          float32
	widthObs, heightObs float32
	status              bool
	obstacle            *canvas.Image
	spriteAnimationTime time.Duration
}

func NewObstacle(posx float32, posy float32, width float32, height float32, img *canvas.Image) *Obstacle {
	return &Obstacle{
		posX:                posx,
		posY:                posy,
		widthObs:            width,
		heightObs:           height,
		status:              true,
		obstacle:            img,
		spriteAnimationTime: 100 * time.Millisecond,
	}
}

func (o *Obstacle) SetStatus(status bool) {
	o.status = status
}

func (o *Obstacle) GetStatus() bool {
	return o.status
}

func (o *Obstacle) Run(c *Car, onCollision func()) {
	o.status = true
	for o.status {
		o.posY += 10
		o.obstacle.Move(fyne.NewPos(o.posX, o.posY))

		if o.posY <= c.posY+c.heightCar && o.posY+o.heightObs >= c.posY &&
			o.posX <= c.posX+c.widthCar && o.posX+o.widthObs >= c.posX {
			o.SetStatus(false)
			c.SetStatus(false)
			println("me chocaron")
			onCollision()
		}

		if o.posY >= 400 {
			if rand.Intn(2) == 0 {
				o.posX = 305
			} else {
				o.posX = 420
			}
			o.posY = 10
		}
		time.Sleep(o.spriteAnimationTime)
	}
}
