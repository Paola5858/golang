package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	screenWidth  = 1280
	screenHeight = 720

	maxAsteroids   = 12
	minAsteroidSize = 20.0

	playerMaxSpeed = 6.5
	playerAccel    = 0.35
	playerMaxAccel = 0.5
	playerFriction = 0.06

	bulletSpeed  = 14.0
	bulletMaxAge = 90
	maxBullets   = 10
	fireCooldown = 10

	explosionFrames = 15

	scale = 1.0
)

var (
	bgColor        = color.White
	textColor      = color.RGBA{107, 114, 128, 255}
	bulletColor    = color.RGBA{0, 0, 0, 255}
	explosionColor = color.RGBA{255, 69, 0, 160}
	imgPlayer      *ebiten.Image
	imgAsteroid    *ebiten.Image
	imgBullet      *ebiten.Image
	imgExplosion   *ebiten.Image
	imgHealthBg    *ebiten.Image
	imgCooldownBg  *ebiten.Image
)

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StatePaused
	StateGameOver
)

type Game struct {
	player       Player
	bullets      []Bullet
	asteroids    []Asteroid
	explosions   []Explosion
	score        int
	highScore    int
	state        GameState
	frames       int
	fontFace     font.Face
	playerW      float64
	playerH      float64
	asteroidW    float64
	asteroidH    float64
	scale        float64
	message      string
	messageTimer int
}

func NewGame() *Game {
	fontFace := loadFont()
	g := &Game{
		fontFace: fontFace,
		state:    StateMenu,
	}

	imgPlayer = loadImage("nave.png")
	boundsPlayer := imgPlayer.Bounds()
	g.playerW = float64(boundsPlayer.Dx())
	g.playerW = 64
	g.playerH = 64

	imgAsteroid = loadImage("asteroide.png")
	boundsAst := imgAsteroid.Bounds()
	g.asteroidW = float64(boundsAst.Dx())
	g.asteroidH = float64(boundsAst.Dy())

	imgBullet = generateCircleImage(12, bulletColor)
	imgExplosion = generateCircleImage(40, color.RGBA{255, 69, 0, 255})
	imgHealthBg = ebiten.NewImage(200, 20)
	imgHealthBg.Fill(color.RGBA{255, 0, 0, 255})
	imgCooldownBg = ebiten.NewImage(200, 20)
	imgCooldownBg.Fill(color.RGBA{0, 0, 255, 255})

	g.player = Player{
		position: Vector{screenWidth / 2, screenHeight / 2},
		width:    g.playerW,
		height:   g.playerH,
		img:      imgPlayer,
		health:   3,
	}

	return g
}

func (g *Game) Reset() {
	g.player = Player{
		position:     Vector{screenWidth / 2, screenHeight / 2},
		width:        g.playerW,
		height:       g.playerH,
		img:          imgPlayer,
		velocity:     Vector{0, 0},
		angle:        0,
		fireCooldown: 0,
		health:       3,
	}
	g.bullets = nil
	g.asteroids = nil
	g.explosions = nil
	g.score = 0
	g.state = StatePlaying
	g.frames = 0
	for i := 0; i < maxAsteroids; i++ {
		g.spawnAsteroid()
	}
}

func (g *Game) spawnAsteroid() {
	minSize := 40.0
	maxSize := 96.0
	size := minSize + rand.Float64()*(maxSize-minSize)
	pos := Vector{X: rand.Float64() * float64(screenWidth), Y: rand.Float64()*float64(screenHeight)/4 - size}
	vel := Vector{X: (rand.Float64()*2 - 1) * 1.5, Y: rand.Float64()*2 + 1}
	rotSpeed := (rand.Float64()*2 - 1) * 0.04
	g.asteroids = append(g.asteroids, Asteroid{position: pos, velocity: vel, size: size, rotSpeed: rotSpeed})
}

func (g *Game) Update() error {
	g.frames++
	switch g.state {
	case StateMenu:
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.Reset()
		}
		if ebiten.IsKeyPressed(ebiten.KeyP) {
			g.state = StatePaused
		}
	case StatePlaying:
		if ebiten.IsKeyPressed(ebiten.KeyP) {
			g.state = StatePaused
		} else {
			g.updatePlaying()
		}
	case StatePaused:
		if ebiten.IsKeyPressed(ebiten.KeyP) {
			g.state = StatePlaying
		}
	case StateGameOver:
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			g.Reset()
		}
	}
	return nil
}

func (g *Game) updatePlaying() {
	g.player.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.player.fireCooldown <= 0 && len(g.bullets) < maxBullets {
		g.fireBullet()
		g.player.fireCooldown = fireCooldown
	}
	g.updateBullets()
	g.updateAsteroids()
	g.updateExplosions()
	for _, a := range g.asteroids {
		if circleCollision(g.player.position.X, g.player.position.Y, g.player.width/2, a.position.X, a.position.Y, a.size/2) {
			g.player.health--
			if g.player.health <= 0 {
				g.state = StateGameOver
				if g.score > g.highScore {
					g.highScore = g.score
				}
			} else {
				g.message = "Você foi atingido!"
				g.messageTimer = 120
			}
			break
		}
	}
	if len(g.asteroids) < maxAsteroids && g.frames%60 == 0 {
		g.spawnAsteroid()
	}
}

func (g *Game) fireBullet() {
	offsetX := math.Sin(g.player.angle) * g.player.height / 2
	offsetY := -math.Cos(g.player.angle) * g.player.height / 2
	bulletVel := Vector{X: math.Sin(g.player.angle) * bulletSpeed, Y: -math.Cos(g.player.angle) * bulletSpeed}
	bulletPos := Vector{X: g.player.position.X + offsetX, Y: g.player.position.Y + offsetY}
	g.bullets = append(g.bullets, Bullet{position: bulletPos, velocity: bulletVel})
}

func (g *Game) updateBullets() {
	active := g.bullets[:0]
	for i := range g.bullets {
		b := &g.bullets[i]
		b.Update()
		if b.age > bulletMaxAge || b.IsOffScreen() {
			continue
		}
		hit := false
		for j, a := range g.asteroids {
			if circleCollision(b.position.X, b.position.Y, 5, a.position.X, a.position.Y, a.size/2) {
				g.explosions = append(g.explosions, Explosion{position: a.position, maxFrame: explosionFrames})
				g.score += int(a.size) * 10
				if a.size > minAsteroidSize {
					// Split into 2 smaller asteroids
					newSize := a.size * 0.6
					for k := 0; k < 2; k++ {
						angle := float64(k) * math.Pi + rand.Float64()*math.Pi/2
						vel := Vector{X: math.Cos(angle) * 2, Y: math.Sin(angle) * 2}
						g.asteroids = append(g.asteroids, Asteroid{position: a.position, velocity: vel, size: newSize, rotSpeed: (rand.Float64()*2 - 1) * 0.04})
					}
				}
				g.asteroids = append(g.asteroids[:j], g.asteroids[j+1:]...)
				hit = true
				break
			}
		}
		if !hit {
			active = append(active, *b)
		}
	}
	g.bullets = active
}

func (g *Game) updateAsteroids() {
	for i := range g.asteroids {
		a := &g.asteroids[i]
		a.Update()
	}
}

func (g *Game) updateExplosions() {
	active := g.explosions[:0]
	for i := range g.explosions {
		e := &g.explosions[i]
		e.Update()
		if e.frame < e.maxFrame {
			active = append(active, *e)
		}
	}
	g.explosions = active
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)
	switch g.state {
	case StateMenu:
		g.drawMenu(screen)
	case StatePlaying:
		g.drawPlaying(screen)
	case StatePaused:
		g.drawPaused(screen)
	case StateGameOver:
		g.drawGameOver(screen)
	}
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	title := "ASTEROIDES PROFISSIONAL"
	instr := "Setas ← → para girar, ↑ para acelerar\nBarra de espaço para atirar\n\nPressione ENTER para começar"
	y := screenHeight / 2
	text.Draw(screen, title, g.fontFace, screenWidth/2-len(title)*7, y-80, textColor)
	text.Draw(screen, instr, g.fontFace, screenWidth/2-170, y-40, textColor)
	text.Draw(screen, fmt.Sprintf("Melhor pontuação: %d", g.highScore), g.fontFace, screenWidth/2-100, y-120, textColor)
}

func (g *Game) drawPlaying(screen *ebiten.Image) {
	g.player.Draw(screen)
	for _, a := range g.asteroids {
		a.Draw(screen)
	}
	for _, b := range g.bullets {
		b.Draw(screen)
	}
	for _, e := range g.explosions {
		e.Draw(screen)
	}
	text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), g.fontFace, 24, 40, textColor)
	text.Draw(screen, fmt.Sprintf("Melhor: %d", g.highScore), g.fontFace, 24, 70, textColor)

	// Draw health bar
	healthBarWidth := 200.0
	healthBarHeight := 20.0
	healthBarX := 24.0
	healthBarY := 100.0
	healthFg := ebiten.NewImage(int(healthBarWidth*float64(g.player.health)/3), int(healthBarHeight))
	healthFg.Fill(color.RGBA{0, 255, 0, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(healthBarX, healthBarY)
	screen.DrawImage(imgHealthBg, op)
	screen.DrawImage(healthFg, op)

	// Draw cooldown bar
	cooldownBarWidth := 200.0
	cooldownBarHeight := 20.0
	cooldownBarX := 24.0
	cooldownBarY := 130.0
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(cooldownBarX, cooldownBarY)
	screen.DrawImage(imgCooldownBg, op2)
	if g.player.fireCooldown > 0 {
		cooldownFg := ebiten.NewImage(int(cooldownBarWidth*float64(g.player.fireCooldown)/fireCooldown), int(cooldownBarHeight))
		cooldownFg.Fill(color.RGBA{255, 255, 0, 255})
		screen.DrawImage(cooldownFg, op2)
	}

	// Draw message if any
	if g.messageTimer > 0 {
		bounds := text.BoundString(g.fontFace, g.message)
		x := screenWidth/2 - bounds.Dx()/2
		y := screenHeight/2 - bounds.Dy()/2
		text.Draw(screen, g.message, g.fontFace, x, y, color.RGBA{255, 0, 0, 255})
		g.messageTimer--
	}
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	lines := []string{"🌟 FIM DE JOGO 🌟", fmt.Sprintf("Pontos finais: %d", g.score), fmt.Sprintf("Melhor pontuação: %d", g.highScore), "Pressione R para tentar novamente"}
	y := screenHeight / 2
	for i, line := range lines {
		bounds := text.BoundString(g.fontFace, line)
		x := screenWidth/2 - bounds.Dx()/2
		text.Draw(screen, line, g.fontFace, x, y+i*32, color.RGBA{255, 69, 0, 255})
	}
}

func (g *Game) drawPaused(screen *ebiten.Image) {
	text.Draw(screen, "PAUSADO - Pressione P para continuar", g.fontFace, screenWidth/2-150, screenHeight/2, textColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.scale = math.Min(float64(outsideWidth)/screenWidth, float64(outsideHeight)/screenHeight)
	return outsideWidth, outsideHeight
}
