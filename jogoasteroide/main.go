package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

)

// Constants defining game parameters and screen size
const (
	screenWidth  = 1280
	screenHeight = 720

	maxAsteroids = 12

	playerMaxSpeed = 6.5
	playerAccel    = 0.35
	playerFriction = 0.06

	bulletSpeed  = 14.0
	bulletMaxAge = 90
	fireCooldown = 10

	explosionFrames = 15

	scale = 1.0
)

// Colors for UI and game elements
var (
	bgColor        = color.White
	textColor      = color.RGBA{107, 114, 128, 255} // neutral gray (#6b7280)
	bulletColor    = color.RGBA{0, 0, 0, 255}        // Black bullets
	explosionColor = color.RGBA{255, 69, 0, 160}     // translucent burning orange
)

// Different game states
type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StateGameOver
)

// Vector struct represents 2D position or velocity
type Vector struct {
	X, Y float64
}

func (v *Vector) Add(v2 Vector) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v Vector) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Normalize() {
	l := v.Len()
	if l == 0 {
		return
	}
	v.X /= l
	v.Y /= l
}

func (v Vector) Scaled(s float64) Vector {
	return Vector{v.X * s, v.Y * s}
}

// Player struct holds the ship's position, velocity, angle, image and cooldown info
type Player struct {
	position Vector
	velocity Vector
	angle    float64

	img *ebiten.Image

	width, height float64

	fireCooldown int
}

// Bullet struct holds position, velocity and age to expire the bullet
type Bullet struct {
	position Vector
	velocity Vector
	age      int
}

// Asteroid struct holds position, velocity, size, rotation angle and rotation speed
type Asteroid struct {
	position Vector
	velocity Vector
	size     float64
	angle    float64
	rotSpeed float64
}

// Explosion struct for visual effect; frames count to animate fading
type Explosion struct {
	position Vector
	frame    int
	maxFrame int
}

// Game struct holds entire game state and assets
type Game struct {
	player     Player
	bullets    []Bullet
	asteroids  []Asteroid
	explosions []Explosion
	score      int
	state      GameState
	frames     int
	fontFace   font.Face

	imgPlayer   *ebiten.Image
	imgAsteroid *ebiten.Image
	playerW     float64
	playerH     float64
	asteroidW   float64
	asteroidH   float64
}

// loadFont returns a basic, readable font face
func loadFont() font.Face {
	return basicfont.Face7x13
}

// NewGame loads images, initializes fonts and state
func NewGame() *Game {
	fontFace := loadFont()

	g := &Game{
		fontFace: fontFace,
		state:    StateMenu,
	}

	// Load images from external files
	playerImg, _, err := ebitenutil.NewImageFromFile("nave.png")
if err != nil {
	log.Fatalf("failed to load nave.png: %v", err)
}

asteroidImg, _, err := ebitenutil.NewImageFromFile("asteroide.png")
if err != nil {
	log.Fatalf("failed to load asteroide.png: %v", err)
}


	// Store images and their natural sizes
	g.imgPlayer = playerImg
	bounds := playerImg.Bounds()
	g.playerW = float64(bounds.Dx())
	g.playerH = float64(bounds.Dy())

	g.imgAsteroid = asteroidImg
	boundsAst := asteroidImg.Bounds()
	g.asteroidW = float64(boundsAst.Dx())
	g.asteroidH = float64(boundsAst.Dy())

	// Initialize player with scaled size to about 64x64 for neat UI

	g.playerW = 64
	g.playerH = 64

	// g.player.Scale = playerScale  // se tiver esse campo, por exemplo

	g.player = Player{
		position: Vector{screenWidth / 2, screenHeight / 2},
		width:    g.playerW,
		height:   g.playerH,
		img:      playerImg,
	}

	return g
}

// Reset initializes game state for a new play session
func (g *Game) Reset() {
	g.player = Player{
		position:    Vector{screenWidth / 2, screenHeight / 2},
		width:       g.playerW,
		height:      g.playerH,
		img:         g.imgPlayer,
		fireCooldown: 0,
		velocity:    Vector{0, 0},
		angle:      0,
	}

	g.bullets = nil
	g.asteroids = nil
	g.explosions = nil
	g.score = 0
	g.state = StatePlaying
	g.frames = 0

	// Spawn initial asteroids with various sizes
	for i := 0; i < maxAsteroids; i++ {
		g.spawnAsteroid()
	}
}

// spawnAsteroid randomly creates a new asteroid at the top part of the screen
func (g *Game) spawnAsteroid() {
	minSize := 40.0
	maxSize := 96.0
	size := minSize + rand.Float64()*(maxSize-minSize)

	pos := Vector{
		X: rand.Float64() * float64(screenWidth),
		Y: rand.Float64()*float64(screenHeight)/4 - size,
	}
	vel := Vector{
		X: (rand.Float64()*2 - 1) * 1.5,
		Y: rand.Float64()*2 + 1,
	}
	rotSpeed := (rand.Float64()*2 - 1) * 0.04 // smoother rotation

	g.asteroids = append(g.asteroids, Asteroid{
		position: pos,
		velocity: vel,
		size:     size,
		angle:    0,
		rotSpeed: rotSpeed,
	})
}

// Update is the main game loop called every frame (~60 FPS)
func (g *Game) Update() error {
	g.frames++

	switch g.state {
	case StateMenu:
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.Reset()
		}
	case StatePlaying:
		g.updatePlaying()
	case StateGameOver:
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			g.Reset()
		}
	}

	return nil
}

// updatePlaying contains core game updates and input processing
func (g *Game) updatePlaying() {
	g.updatePlayer()

	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.player.fireCooldown <= 0 {
		g.fireBullet()
		g.player.fireCooldown = fireCooldown
	} else if g.player.fireCooldown > 0 {
		g.player.fireCooldown--
	}

	g.updateBullets()
	g.updateAsteroids()
	g.updateExplosions()

	// Check collision player-asteroid for game over
	for _, a := range g.asteroids {
		if circleCollision(g.player.position.X, g.player.position.Y, g.player.width/2, a.position.X, a.position.Y, a.size/2) {
			g.state = StateGameOver
			break
		}
	}

	// Spawn asteroids to maintain challenge
	if len(g.asteroids) < maxAsteroids && g.frames%60 == 0 {
		g.spawnAsteroid()
	}
}

// updatePlayer processes movement and rotation input
func (g *Game) updatePlayer() {
	// Rotation input
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.angle -= 0.09
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.angle += 0.09
	}

	// Accelerate forward
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		accel := Vector{
			X: math.Sin(g.player.angle) * playerAccel,
			Y: -math.Cos(g.player.angle) * playerAccel,
		}
		g.player.velocity.Add(accel)
	} else {
		g.player.velocity.X *= (1 - playerFriction)
		g.player.velocity.Y *= (1 - playerFriction)
	}

	// Cap velocity
	speed := g.player.velocity.Len()
	if speed > playerMaxSpeed {
		g.player.velocity.Normalize()
		g.player.velocity = g.player.velocity.Scaled(playerMaxSpeed)
	}

	// Update position
	g.player.position.Add(g.player.velocity)

	// Screen wraparound for smooth gameplay
	if g.player.position.X < 0 {
		g.player.position.X = screenWidth
	}
	if g.player.position.X > screenWidth {
		g.player.position.X = 0
	}
	if g.player.position.Y < 0 {
		g.player.position.Y = screenHeight
	}
	if g.player.position.Y > screenHeight {
		g.player.position.Y = 0
	}
}

// fireBullet creates a bullet traveling from player's tip
func (g *Game) fireBullet() {
	offsetX := math.Sin(g.player.angle) * g.player.height / 2
	offsetY := -math.Cos(g.player.angle) * g.player.height / 2

	bulletVel := Vector{
		X: math.Sin(g.player.angle) * bulletSpeed,
		Y: -math.Cos(g.player.angle) * bulletSpeed,
	}
	bulletPos := Vector{
		X: g.player.position.X + offsetX,
		Y: g.player.position.Y + offsetY,
	}

	g.bullets = append(g.bullets, Bullet{
		position: bulletPos,
		velocity: bulletVel,
		age:      0,
	})
}

// updateBullets moves bullets and handles collisions and lifetimes
func (g *Game) updateBullets() {
	activeBullets := g.bullets[:0]

	for i := range g.bullets {
		b := &g.bullets[i]
		b.position.Add(b.velocity)
		b.age++

		if b.age > bulletMaxAge {
			continue
		}

		hit := false
		for j, a := range g.asteroids {
			if circleCollision(b.position.X, b.position.Y, 5, a.position.X, a.position.Y, a.size/2) {
				// Create explosion effect
				g.explosions = append(g.explosions, Explosion{
					position: a.position,
					frame:    0,
					maxFrame: explosionFrames,
				})
				// Increase score proportionally to asteroid size
				g.score += int(a.size) * 10

				// Remove asteroid cleanly
				g.asteroids = append(g.asteroids[:j], g.asteroids[j+1:]...)
				hit = true
				break
			}
		}
		if !hit {
			activeBullets = append(activeBullets, *b)
		}
	}
	g.bullets = activeBullets
}

// updateAsteroids moves asteroids, applies rotation and screen wrap
func (g *Game) updateAsteroids() {
	for i := range g.asteroids {
		a := &g.asteroids[i]

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
}

// updateExplosions progresses and removes finished explosion animations
func (g *Game) updateExplosions() {
	active := g.explosions[:0]
	for i := range g.explosions {
		e := &g.explosions[i]
		e.frame++
		if e.frame < e.maxFrame {
			active = append(active, *e)
		}
	}
	g.explosions = active
}

// circleCollision detects overlap between two circles (for collision detection)
func circleCollision(x1, y1, r1, x2, y2, r2 float64) bool {
	dx := x1 - x2
	dy := y1 - y2
	return dx*dx+dy*dy < (r1+r2)*(r1+r2)
}

// Draw renders game content depending on current game state
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	switch g.state {
	case StateMenu:
		g.drawMenu(screen)
	case StatePlaying:
		g.drawPlaying(screen)
	case StateGameOver:
		g.drawGameOver(screen)
	}
}

// drawMenu shows game title and instructions with elegant typography
func (g *Game) drawMenu(screen *ebiten.Image) {
	title := "ASTEROIDES PROFISSIONAL"
	instr := "Setas ← → para girar, ↑ para acelerar\nBarra de espaço para atirar\n\nPressione ENTER para começar"
	yStart := screenHeight/2 - 80

	text.Draw(screen, title, g.fontFace, screenWidth/2-len(title)*7, yStart, textColor)
	text.Draw(screen, instr, g.fontFace, screenWidth/2-170, yStart+40, textColor)
}

// drawPlaying renders the player, asteroids, bullets, explosions and score
func (g *Game) drawPlaying(screen *ebiten.Image) {
	// Draw player ship with rotation and scaled nicely to 64x64
	op := &ebiten.DrawImageOptions{}
	scaleFactor := 64.0 / float64(g.imgPlayer.Bounds().Dx())
	op.GeoM.Translate(-float64(g.imgPlayer.Bounds().Dx())/2, -float64(g.imgPlayer.Bounds().Dy())/2)
	op.GeoM.Rotate(g.player.angle)
	op.GeoM.Scale(scaleFactor, scaleFactor)
	op.GeoM.Translate(g.player.position.X, g.player.position.Y)
	screen.DrawImage(g.imgPlayer, op)

	// Draw asteroids scaled to their size, with rotation
	for _, a := range g.asteroids {
		op := &ebiten.DrawImageOptions{}
		asteroidBaseSize := float64(g.imgAsteroid.Bounds().Dx())
		scaleA := a.size / asteroidBaseSize
		op.GeoM.Translate(-asteroidBaseSize/2, -asteroidBaseSize/2)
		op.GeoM.Rotate(a.angle)
		op.GeoM.Scale(scaleA, scaleA)
		op.GeoM.Translate(a.position.X, a.position.Y)
		screen.DrawImage(g.imgAsteroid, op)
	}

	// Draw bullets as crisp black filled circles
	for _, b := range g.bullets {
		op := &ebiten.DrawImageOptions{}
		r := 6.0
		op.GeoM.Translate(b.position.X-r, b.position.Y-r)
		bulletImg := generateCircleImage(int(r*2), bulletColor)
		screen.DrawImage(bulletImg, op)
	}

	// Draw explosions as pulsating orange circles fading out
	for _, e := range g.explosions {
		op := &ebiten.DrawImageOptions{}
		alpha := uint8(180 * (1 - float64(e.frame)/float64(e.maxFrame)))
		color := color.RGBA{255, 69, 0, alpha}
		explImg := generateCircleImage(40, color)
		op.GeoM.Translate(e.position.X-20, e.position.Y-20)
		screen.DrawImage(explImg, op)
	}

	// Draw score top-left with clean margin and font size
	scoreText := fmt.Sprintf("Pontos: %d", g.score)
	text.Draw(screen, scoreText, g.fontFace, 24, 40, textColor)
}

// drawGameOver presents final score and restart instruction elegantly centered
func (g *Game) drawGameOver(screen *ebiten.Image) {
	lines := []string{
		"🌟 FIM DE JOGO 🌟",
		fmt.Sprintf("Pontos finais: %d", g.score),
		"Pressione R para tentar novamente",
	}

	yStart := screenHeight/2 - (len(lines) * 20 / 2)
	for i, line := range lines {
		bounds := text.BoundString(g.fontFace, line)
		x := screenWidth/2 - bounds.Dx()/2
		y := yStart + i*32
		text.Draw(screen, line, g.fontFace, x, y, color.RGBA{255, 69, 0, 255})
	}
}

// generateCircleImage makes a filled circle image of specified diameter and color
func generateCircleImage(d int, clr color.Color) *ebiten.Image {
	img := ebiten.NewImage(d, d)
	c := float64(d) / 2
	for y := 0; y < d; y++ {
		for x := 0; x < d; x++ {
			dx := float64(x) - c
			dy := float64(y) - c
			if dx*dx+dy*dy <= c*c {
				img.Set(x, y, clr)
			}
		}
	}
	return img
}

// Layout returns fixed game window size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Asteroides Profissional - Clean & Elegant")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
