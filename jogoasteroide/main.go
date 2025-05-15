// jogo/main.go
package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth   = 640
	screenHeight  = 480
	playerSpeed   = 4
	playerWidth   = 64
	playerHeight  = 64
	asteroidWidth = 64
	asteroidHeight = 64
	asteroidSpeed = 2
	numAsteroids  = 5
	scale         = 0.1
)

type Asteroid struct {
	x, y float64
	img  *ebiten.Image
}

type Game struct {
	playerX, playerY float64
	player           *ebiten.Image
	asteroids        []Asteroid
	asteroidImage    *ebiten.Image
	gameOver         bool
	startTime        time.Time
	score int
}

func (g *Game) Update() error {
	if g.gameOver {
		return nil
	}

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

	for i := range g.asteroids {
		g.asteroids[i].y += asteroidSpeed
		if g.asteroids[i].y > screenHeight {
			g.asteroids[i].y = -asteroidHeight
			g.asteroids[i].x = float64(rand.Intn(screenWidth - asteroidWidth))
			g.score++ // soma pontos ao evitar o asteroide
		}
		if checkCollision(g.playerX, g.playerY, g.asteroids[i].x, g.asteroids[i].y) {
			g.gameOver = true
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	if g.gameOver {
		msg := fmt.Sprintf("Game Over! Score: %d", g.score)
		ebitenutil.DebugPrint(screen, msg)
		return
	}

	opPlayer := &ebiten.DrawImageOptions{}
	opPlayer.GeoM.Scale(0.1, 0.1)
	opPlayer.GeoM.Translate(g.playerX, g.playerY)
	screen.DrawImage(g.player, opPlayer)

	for _, a := range g.asteroids {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.1, 0.1)
		op.GeoM.Translate(a.x, a.y)
		screen.DrawImage(a.img, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.score))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func checkCollision(px, py, ax, ay float64) bool {
	return px < ax+asteroidWidth*scale &&
		px+playerWidth*scale > ax &&
		py < ay+asteroidHeight*scale &&
		py+playerHeight*scale > ay
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func main() {
	rand.Seed(time.Now().UnixNano())

	playerImg := loadImage("nave.png")
	asteroidImg := loadImage("asteroide.png")

	asteroids := make([]Asteroid, numAsteroids)
	for i := range asteroids {
		asteroids[i] = Asteroid{
			x: float64(rand.Intn(screenWidth - asteroidWidth)),
			y: float64(rand.Intn(screenHeight - asteroidHeight)),
			img: asteroidImg,
		}
	}

	game := &Game{
		playerX:        300,
		playerY:        400,
		player:         playerImg,
		asteroids:      asteroids,
		asteroidImage:  asteroidImg,
		startTime:      time.Now(),
		score:          0,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Asteroides da Paola 💫")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
