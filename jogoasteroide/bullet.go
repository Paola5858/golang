package main

import "github.com/hajimehoshi/ebiten/v2"

type Bullet struct {
	position Vector
	velocity Vector
	age      int
}

func (b *Bullet) Update() {
	b.position.Add(b.velocity)
	b.age++
}

func (b *Bullet) IsOffScreen() bool {
	return b.position.X < -10 || b.position.X > screenWidth+10 || b.position.Y < -10 || b.position.Y > screenHeight+10
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	r := 6.0
	op.GeoM.Translate(b.position.X-r, b.position.Y-r)
	screen.DrawImage(imgBullet, op)
}
