package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PowerUpType int

const (
	PowerUpShield PowerUpType = iota
	PowerUpRapidFire
	PowerUpMultiShot
	PowerUpExtraLife
)

type PowerUp struct {
	position Vector
	velocity Vector
	powerType PowerUpType
	size      float64
	age       int
	maxAge    int
}

func (p *PowerUp) Update() {
	p.position.Add(p.velocity)
	p.age++
	// Wrap around screen
	if p.position.X < 0 {
		p.position.X = ScreenWidth
	}
	if p.position.X > ScreenWidth {
		p.position.X = 0
	}
	if p.position.Y < 0 {
		p.position.Y = ScreenHeight
	}
	if p.position.Y > ScreenHeight {
		p.position.Y = 0
	}
}

func (p *PowerUp) Draw(screen *ebiten.Image) {
	var col color.RGBA
	switch p.powerType {
	case PowerUpShield:
		col = color.RGBA{0, 255, 255, 255} // Cyan
	case PowerUpRapidFire:
		col = color.RGBA{255, 255, 0, 255} // Yellow
	case PowerUpMultiShot:
		col = color.RGBA{255, 0, 255, 255} // Magenta
	case PowerUpExtraLife:
		col = color.RGBA{0, 255, 0, 255} // Green
	}
	img := generateCircleImage(int(p.size), col)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X-p.size/2, p.position.Y-p.size/2)
	screen.DrawImage(img, op)
}

func (p *PowerUp) IsExpired() bool {
	return p.age > p.maxAge
}

type PowerUpPool struct {
	pool []PowerUp
}

func (p *PowerUpPool) Get() *PowerUp {
	if len(p.pool) > 0 {
		pw := &p.pool[len(p.pool)-1]
		p.pool = p.pool[:len(p.pool)-1]
		return pw
	}
	return &PowerUp{}
}

func (p *PowerUpPool) Put(pw *PowerUp) {
	*pw = PowerUp{} // reset
	p.pool = append(p.pool, *pw)
}
