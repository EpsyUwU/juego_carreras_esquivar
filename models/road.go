package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type SceneRoad struct {
	posY                float32
	status              bool
	background          [3]*canvas.Image
	index               int
	window              fyne.Window
	spriteAnimationTime time.Duration
}

func NewSceneRoad(imgs [3]*canvas.Image, window fyne.Window) *SceneRoad {
	return &SceneRoad{
		posY:                0,
		status:              true,
		background:          imgs,
		index:               0,
		window:              window,
		spriteAnimationTime: 120 * time.Millisecond,
	}
}

func (e *SceneRoad) Run(sceneContainer *fyne.Container) {
	e.status = true
	for _, img := range e.background {
		img.Hide()
		sceneContainer.Add(img)
	}
	go func() {
		for e.status {
			for _, img := range e.background {
				img.Hide()
			}
			e.background[e.index].Show()
			sceneContainer.Refresh()
			time.Sleep(e.spriteAnimationTime)
			e.index = (e.index + 1) % len(e.background)
		}
	}()
}

func (e *SceneRoad) SetStatus(status bool) {
	e.status = status
}
