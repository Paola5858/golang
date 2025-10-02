package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position       Vector
	velocity       Vector
	acceleration   Vector
	angle          float64
	img            *ebiten.Image
	width, height  float64
	fireCooldown   int
	isAccelerating bool
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.angle -= 0.09
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.angle += 0.09
	}
	p.isAccelerating = ebiten.IsKeyPressed(ebiten.KeyUp)
	if p.isAccelerating {
		p.acceleration = Vector{X: math.Sin(p.angle) * playerAccel, Y: -math.Cos(p.angle) * playerAccel}
	} else {
		p.acceleration = Vector{0, 0}
	}
	// Apply acceleration to velocity
	p.velocity.Add(p.acceleration)
	// Apply friction
	p.velocity.X *= 1 - playerFriction
	p.velocity.Y *= 1 - playerFriction
	// Clamp speed
	speed := p.velocity.Len()
	if speed > playerMaxSpeed {
		p.velocity.Normalize()
		p.velocity = p.velocity.Scaled(playerMaxSpeed)
	}
	// Update position
	p.position.Add(p.velocity)
	// Wrap around screen
	if p.position.X < 0 {
		p.position.X = screenWidth
	}
	if p.position.X > screenWidth {
		p.position.X = 0
	}
	if p.position.Y < 0 {
		p.position.Y = screenHeight
	}
	if p.position.Y > screenHeight {
		p.position.Y = 0
	}
	if p.fireCooldown > 0 {
		p.fireCooldown--
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	sf := 64.0 / float64(p.img.Bounds().Dx())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(p.img.Bounds().Dx())/2, -float64(p.img.Bounds().Dy())/2)
	op.GeoM.Rotate(p.angle)
	op.GeoM.Scale(sf, sf)
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.img, op)

	if p.isAccelerating {
		// Draw thruster flame
		op := &ebiten.DrawImageOptions{}
		r := 8.0
		op.GeoM.Translate(-r, -r)
		op.GeoM.Rotate(p.angle)
		op.GeoM.Translate(p.position.X-math.Sin(p.angle)*35, p.position.Y+math.Cos(p.angle)*35)
		img := generateCircleImage(int(r*2), color.RGBA{255, 165, 0, 200})
		screen.DrawImage(img, op)
	}
}
