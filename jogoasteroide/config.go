package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game configuration constants
const (
	ScreenWidth  = 1280
	ScreenHeight = 720
	Scale        = 1.0
)

// Player configuration
const (
	PlayerMaxSpeed = 6.5
	PlayerAccel    = 0.35
	PlayerFriction = 0.06
	PlayerWidth    = 64
	PlayerHeight   = 64
)

// Bullet configuration
const (
	BulletSpeed  = 14.0
	BulletMaxAge = 90
	MaxBullets   = 10
	FireCooldown = 10
)

// Asteroid configuration
const (
	MaxAsteroids   = 12
	MinAsteroidSize = 20.0
)

// Explosion configuration
const (
	ExplosionFrames = 15
)

// Power-up configuration
const (
	PowerUpMaxAge = 600 // 10 seconds at 60fps
)

// Colors
var (
	BgColor        = color.White
	TextColor      = color.RGBA{107, 114, 128, 255}
	BulletColor    = color.RGBA{0, 0, 0, 255}
	ExplosionColor = color.RGBA{255, 69, 0, 160}
)

// Images (to be loaded)
var (
	ImgPlayer   *ebiten.Image
	ImgAsteroid *ebiten.Image
	ImgBullet   *ebiten.Image
	ImgExplosion *ebiten.Image
	ImgHealthBg  *ebiten.Image
	ImgCooldownBg *ebiten.Image
)
