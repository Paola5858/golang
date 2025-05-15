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
	screenWidth    = 640
	screenHeight   = 480
	playerSpeed    = 4
	playerWidth    = 64
	playerHeight   = 64
	asteroidWidth  = 64
	asteroidHeight = 64
	initialSpeed   = 2
	numAsteroids   = 5
	scale          = 0.1
)

type Asteroid struct {
	x, y, speed float64
	img         *ebiten.Image
}

type Game struct {
	player1X, player1Y float64
	player2X, player2Y float64
	player             *ebiten.Image
	asteroids          []Asteroid
	asteroidImages     []*ebiten.Image
	gameOver           bool
	startTime          time.Time
	score       int
	difficultyT float64
}

func (g *Game) Update() error {
	if g.gameOver {
		return nil
	}

	// Movimento Player 1 (setas)
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.player1X > 0 {
		g.player1X -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.player1X < screenWidth-playerWidth {
		g.player1X += playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.player1Y > 0 {
		g.player1Y -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.player1Y < screenHeight-playerHeight {
		g.player1Y += playerSpeed
	}

	// Movimento Player 2 (WASD)
	if ebiten.IsKeyPressed(ebiten.KeyA) && g.player2X > 0 {
		g.player2X -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && g.player2X < screenWidth-playerWidth {
		g.player2X += playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) && g.player2Y > 0 {
		g.player2Y -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.player2Y < screenHeight-playerHeight {
		g.player2Y += playerSpeed
	}

	elapsed := time.Since(g.startTime).Seconds()
	g.difficultyT = 1 + elapsed/15

	for i := range g.asteroids {
		g.asteroids[i].y += g.asteroids[i].speed * g.difficultyT
		if g.asteroids[i].y > screenHeight {
			g.asteroids[i].y = -asteroidHeight
			g.asteroids[i].x = float64(rand.Intn(screenWidth - asteroidWidth))
			g.asteroids[i].img = g.asteroidImages[rand.Intn(len(g.asteroidImages))]
			g.asteroids[i].speed = initialSpeed + rand.Float64()*2
			g.score++
		}
		if checkCollision(g.player1X, g.player1Y, g.asteroids[i].x, g.asteroids[i].y) ||
			checkCollision(g.player2X, g.player2Y, g.asteroids[i].x, g.asteroids[i].y) {
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

	// Player 1
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Scale(scale, scale)
	op1.GeoM.Translate(g.player1X, g.player1Y)
	screen.DrawImage(g.player, op1)

	// Player 2
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(scale, scale)
	op2.GeoM.Translate(g.player2X, g.player2Y)
	screen.DrawImage(g.player, op2)

	for _, a := range g.asteroids {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
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
	asteroid1 := loadImage("asteroide.png")
	asteroid2 := loadImage("asteroide2.png")

	asteroids := make([]Asteroid, numAsteroids)
	for i := range asteroids {
		asteroids[i] = Asteroid{
			x:     float64(rand.Intn(screenWidth - asteroidWidth)),
			y:     float64(rand.Intn(screenHeight - asteroidHeight)),
			img:   []*ebiten.Image{asteroid1, asteroid2}[rand.Intn(2)],
			speed: initialSpeed + rand.Float64()*2,
		}
	}

	game := &Game{
		player1X:        200,
		player1Y:        400,
		player2X:        400,
		player2Y:        400,
		player:          playerImg,
		asteroids:       asteroids,
		asteroidImages:  []*ebiten.Image{asteroid1, asteroid2},
		startTime:       time.Now(),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Asteroides da Paola 💫 - 2 Players")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}