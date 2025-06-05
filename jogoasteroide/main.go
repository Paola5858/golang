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
	screenWidth    = 1280
	screenHeight   = 720
	playerSpeed    = 4
	playerWidth    = 64
	playerHeight   = 64
	asteroidWidth  = 64
	asteroidHeight = 64
	asteroidSpeed  = 2
	numAsteroids   = 5
	scale          = 0.1
	spawnInterval  = 60
)

type Asteroid struct {
	x, y, speedY float64
	img          *ebiten.Image
}

type Game struct {
	playerX, playerY float64
	player           *ebiten.Image
	asteroids        []Asteroid
	frameCount       int
	asteroidImage    *ebiten.Image
	gameOver         bool
	startTime        time.Time
	score int
}

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			g.restart()
		}
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.playerX > 0 {
		g.playerX -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.playerX < screenWidth-playerWidth*scale {
		g.playerX += playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.playerY > 0 {
		g.playerY -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.playerY < screenHeight-playerHeight*scale {
		g.playerY += playerSpeed
	}

	g.frameCount++
	if g.frameCount%spawnInterval == 0 {
		scaledAW := asteroidWidth * scale
		newAsteroid := Asteroid{
			x:      float64(rand.Intn(screenWidth - int(scaledAW))),
			y:      -asteroidHeight * scale,
			speedY: 2 + rand.Float64()*3,
			img:    g.asteroidImage,
		}
		g.asteroids = append(g.asteroids, newAsteroid)
	}

	activeAsteroids := g.asteroids[:0]
	for i := range g.asteroids {
		g.asteroids[i].y += g.asteroids[i].speedY
		if g.asteroids[i].y > screenHeight {
			continue
		}
		if checkCollision(g.playerX, g.playerY, g.asteroids[i].x, g.asteroids[i].y) {
			g.gameOver = true
		}
		activeAsteroids = append(activeAsteroids, g.asteroids[i])
	}
	g.asteroids = activeAsteroids

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	if g.gameOver {
		msg := fmt.Sprintf("Game Over! Score: %d | Pressione R para reiniciar", g.score)
		textX := (screenWidth - len(msg)*7) / 2
		textY := screenHeight / 2
		ebitenutil.DebugPrintAt(screen, msg, textX, textY)
		return
	}

	opPlayer := &ebiten.DrawImageOptions{}
	opPlayer.GeoM.Scale(scale, scale)
	opPlayer.GeoM.Translate(g.playerX, g.playerY)
	screen.DrawImage(g.player, opPlayer)

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
	playerW := playerWidth * scale
	playerH := playerHeight * scale
	asteroidW := asteroidWidth * scale
	asteroidH := asteroidHeight * scale

	return px < ax+asteroidW &&
		px+playerW > ax &&
		py < ay+asteroidH &&
		py+playerH > ay
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func (g *Game) restart() {
	g.playerX = 300
	g.playerY = 400
	g.score = 0
	g.frameCount = 0
	g.asteroids = nil
	g.gameOver = false
}

func main() {
	rand.Seed(time.Now().UnixNano())

	playerImg := loadImage("nave.png")
	asteroidImg := loadImage("asteroide.png")

	scaledWidth := asteroidWidth * scale
	scaledHeight := asteroidHeight * scale

	asteroids := make([]Asteroid, numAsteroids)
	for i := range asteroids {
		x := rand.Intn(screenWidth - int(scaledWidth))
		y := rand.Intn(screenHeight - int(scaledHeight))
		asteroids[i] = Asteroid{
			x:      float64(x),
			y:      float64(y),
			speedY: 2 + rand.Float64()*3,
			img:    asteroidImg,
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

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Asteroides da Paola 🌫️")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
