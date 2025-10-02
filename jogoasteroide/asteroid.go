package main

import "github.com/hajimehoshi/ebiten/v2"

type Asteroid struct {
	position Vector
	velocity Vector
	size     float64
	angle    float64
	rotSpeed float64
}

func (a *Asteroid) Update() {
	a.position.Add(a.velocity)
	a.angle += a.rotSpeed
	if a.position.X < -a.size {
		a.position.X = ScreenWidth + a.size
	}
	if a.position.X > ScreenWidth+a.size {
		a.position.X = -a.size
	}
	if a.position.Y < -a.size {
		a.position.Y = ScreenHeight + a.size
	}
	if a.position.Y > ScreenHeight+a.size {
		a.position.Y = -a.size
	}
}

func (a *Asteroid) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	s := a.size / float64(ImgAsteroid.Bounds().Dx())
	op.GeoM.Translate(-float64(ImgAsteroid.Bounds().Dx())/2, -float64(ImgAsteroid.Bounds().Dy())/2)
	op.GeoM.Rotate(a.angle)
	op.GeoM.Scale(s, s)
	op.GeoM.Translate(a.position.X, a.position.Y)
	screen.DrawImage(ImgAsteroid, op)
}
