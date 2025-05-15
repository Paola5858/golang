package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "image"
    "log"
	"runtime"
)

const (
	screenWidth  = 640
	screenHeight = 480
	playerSpeed  = 3
	playerWidth  = 64
	playerHeight = 64
	asteroidWidth  = 64
	asteroidHeight = 64
	asteroidSpeed = 2
	)
type Game struct {
	playerX, playerY float64
	player *ebiten.Image
	asteroid *ebiten.Image
	asteroidX, asteroidY float64
	asteroidSpeed float64
}
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.playerX > 0 {
		g.playerX -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.playerX < screenWidth-playerWidth {
		g.playerX += playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.playerY > 0 {
		g.playerY -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.playerY < screenHeight-playerHeight {
		g.playerY += playerSpeed
	}
	return nil
}

func Draw(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerX, g.playerY)
	screen.DrawImage(g.player, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.asteroidX, g.asteroidY)
	screen.DrawImage(g.asteroid, op)
}
