package main

import (
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
	alpha := float64(180) * (1 - float64(e.frame)/float64(e.maxFrame)) / 255
	op.ColorM.Scale(1, 1, 1, alpha)
	op.GeoM.Translate(e.position.X-20, e.position.Y-20)
	screen.DrawImage(ImgExplosion, op)
}
