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
		a.position.X = screenWidth + a.size
	}
	if a.position.X > screenWidth+a.size {
		a.position.X = -a.size
	}
	if a.position.Y < -a.size {
		a.position.Y = screenHeight + a.size
	}
	if a.position.Y > screenHeight+a.size {
		a.position.Y = -a.size
	}
}

func (a *Asteroid) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	s := a.size / float64(imgAsteroid.Bounds().Dx())
	op.GeoM.Translate(-float64(imgAsteroid.Bounds().Dx())/2, -float64(imgAsteroid.Bounds().Dy())/2)
	op.GeoM.Rotate(a.angle)
	op.GeoM.Scale(s, s)
	op.GeoM.Translate(a.position.X, a.position.Y)
	screen.DrawImage(imgAsteroid, op)
}
