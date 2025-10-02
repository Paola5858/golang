package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Explosion struct {
	position Vector
	frame    int
	maxFrame int
}

func (e *Explosion) Update() {
	e.frame++
}

func (e *Explosion) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	alpha := uint8(180 * (1 - float64(e.frame)/float64(e.maxFrame)))
	img := generateCircleImage(40, color.RGBA{255, 69, 0, alpha})
	op.GeoM.Translate(e.position.X-20, e.position.Y-20)
	screen.DrawImage(img, op)
}
