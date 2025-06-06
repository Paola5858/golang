package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// func loadFont loads a default font face (fallback if no TTF is embedded)

func loadFont(size float64) font.Face {
	
	face, err := opentype.NewFace(nil, &opentype.FaceOptions{
		Size:    size,
		DPI:     fontDPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return face
}


const (
	screenWidth  = 1280
	screenHeight = 720

	playerWidth  = 64
	playerHeight = 64

	asteroidMinSize = 30
	asteroidMaxSize = 90

	maxAsteroids = 10

	fontSize  = 24
	fontDPI   = 72

	playerMaxSpeed  = 6.0
	playerAccel     = 0.2
	playerFriction  = 0.05
	bulletSpeed    = 10.0
	bulletMaxAge   = 60 // frames bullet lasts
	fireCooldown   = 10 // frames between shots
	explosionFrames = 12 // frames of explosion animation

	scale = 1.0 // <<< ADICIONE ESTA LINHA
)

var (
	bgColor    = color.RGBA{5, 5, 12, 255}
	textColor  = color.RGBA{220, 220, 220, 255}
)

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StateGameOver
)

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

type Player struct {
	position Vector
	velocity Vector
	angle    float64
	img      *ebiten.Image

	width, height float64

	fireCooldown int
}

type Bullet struct {
	position Vector
	velocity Vector
	age      int
}

type Asteroid struct {
	position Vector
	velocity Vector
	size     float64
	angle    float64
	rotSpeed float64
}

type Explosion struct {
	position Vector
	frame    int
	maxFrame int
}

type Game struct {
	player      Player
	bullets     []Bullet
	asteroids   []Asteroid
	explosions  []Explosion
	score       int
	state       GameState
	frames      int
	fontFace    font.Face
	audioCtx    *audio.Context
	soundShoot  *audio.Player
	soundExpl   *audio.Player
}

func NewGame() *Game {
	audioCtx := audio.NewContext(44100)

	fontFace := loadFont(fontSize)

	g := &Game{
		fontFace: fontFace,
		audioCtx: audioCtx,
		state:    StateMenu,
	}
	g.player = Player{
		position: Vector{screenWidth / 2, screenHeight / 2},
		width:    playerWidth,
		height:   playerHeight,
		img:      generatePlayerImage(playerWidth, playerHeight),
	}

	// Load sounds
	g.soundShoot = loadSound(audioCtx, shootSoundMp3)
	g.soundExpl = loadSound(audioCtx, explosionSoundMp3)

	// Uncomment to enable looping music if you add music bytes
	// g.musicPlayer = loadSound(audioCtx, backgroundMusicMp3)
	// g.musicPlayer.SetVolume(0.3)
	// g.musicPlayer.Play()

	return g
}

func (g *Game) Reset() {
	g.player = Player{
		position: Vector{screenWidth / 2, screenHeight / 2},
		width:    playerWidth,
		height:   playerHeight,
		img:      generatePlayerImage(playerWidth, playerHeight),
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
	size := float64(rand.Intn(asteroidMaxSize-asteroidMinSize) + asteroidMinSize)
	pos := Vector{
		X: rand.Float64() * float64(screenWidth),
		Y: rand.Float64()*float64(screenHeight)/4 - size, // Spawn in top quarter, slightly out of screen
	}
	vel := Vector{
		X: (rand.Float64()*2 - 1) * 1.5,
		Y: rand.Float64()*2 + 1,
	}
	rotSpeed := (rand.Float64()*2 - 1) * 0.05 // rotation speed -0.05 to 0.05

	g.asteroids = append(g.asteroids, Asteroid{
		position: pos,
		velocity: vel,
		size:     size,
		angle:    0,
		rotSpeed: rotSpeed,
	})
}

func generatePlayerImage(w, h int) *ebiten.Image {
	img := ebiten.NewImage(w, h)
	// Draw a simple white triangle (ship)
	white := color.White
	// draw triangle pointing up
	points := []image.Point{{w / 2, 0}, {0, h}, {w, h}}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if pointInTriangle(image.Point{x, y}, points[0], points[1], points[2]) {
				img.Set(x, y, white)
			}
		}
	}
	return img
}

func pointInTriangle(pt, v1, v2, v3 image.Point) bool {
	d1 := sign(pt, v1, v2)
	d2 := sign(pt, v2, v3)
	d3 := sign(pt, v3, v1)
	hasNeg := (d1 < 0) || (d2 < 0) || (d3 < 0)
	hasPos := (d1 > 0) || (d2 > 0) || (d3 > 0)
	return !(hasNeg && hasPos)
}

func sign(p1, p2, p3 image.Point) int {
	return (p1.X-p3.X)*(p2.Y-p3.Y) - (p2.X-p3.X)*(p1.Y-p3.Y)
}


func (g *Game) Update() error {
	switch g.state {
	case StateMenu:
		return g.updateMenu()
	case StatePlaying:
		return g.updatePlaying()
	case StateGameOver:
		return g.updateGameOver()
	}
	return nil
}

func (g *Game) updateMenu() error {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.Reset()
	}
	return nil
}

func (g *Game) updateGameOver() error {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.Reset()
	}
	return nil
}

func (g *Game) updatePlaying() error {
	g.frames++

	// Update player movement
	g.updatePlayer()

	// Fire bullets
	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.player.fireCooldown <= 0 {
		g.fireBullet()
		g.player.fireCooldown = fireCooldown
	} else if g.player.fireCooldown > 0 {
		g.player.fireCooldown--
	}

	// Update bullets
	g.updateBullets()

	// Update asteroids
	g.updateAsteroids()

	// Update explosions
	g.updateExplosions()

	// Collision detection player & asteroids
	for _, a := range g.asteroids {
		if circleCollision(g.player.position.X, g.player.position.Y, g.player.width*scale/2, a.position.X, a.position.Y, a.size/2) {
			// Player died
			g.state = StateGameOver
		}
	}

	// Spawn new asteroids if needed
	if len(g.asteroids) < maxAsteroids && g.frames%60 == 0 {
		g.spawnAsteroid()
	}

	return nil
}

func (g *Game) updatePlayer() {
	// Rotate
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.angle -= 0.08
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.angle += 0.08
	}

	// Accelerate if up pressed
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		accel := Vector{
			X: math.Sin(g.player.angle) * playerAccel,
			Y: -math.Cos(g.player.angle) * playerAccel,
		}
		g.player.velocity.Add(accel)
	} else {
		// Friction slow down velocity
		g.player.velocity.X *= 1 - playerFriction
		g.player.velocity.Y *= 1 - playerFriction
	}

	// Limit max speed
	speed := g.player.velocity.Len()
	if speed > playerMaxSpeed {
		g.player.velocity.Normalize()
		g.player.velocity = g.player.velocity.Scaled(playerMaxSpeed)
	}

	// Update player position
	g.player.position.Add(g.player.velocity)

	// Screen wrap
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

func (g *Game) fireBullet() {
	offsetX := math.Sin(g.player.angle) * g.player.height * scale / 2
	offsetY := -math.Cos(g.player.angle) * g.player.height * scale / 2

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

	go g.soundShoot.Rewind()
	go g.soundShoot.Play()
}

func (g *Game) updateBullets() {
	activeBullets := g.bullets[:0]
	for i := range g.bullets {
		b := &g.bullets[i]
		b.position.Add(b.velocity)
		b.age++
		if b.age > bulletMaxAge {
			continue
		}
		// Check collision with asteroids
		hit := false
		for j, a := range g.asteroids {
			if circleCollision(b.position.X, b.position.Y, 3, a.position.X, a.position.Y, a.size/2) {
				// Destroy asteroid
				g.explosions = append(g.explosions, Explosion{
					position: a.position,
					frame:    0,
					maxFrame: explosionFrames,
				})
				// Play explosion sound
				go g.soundExpl.Rewind()
				go g.soundExpl.Play()

				g.score += int(a.size)

				// Remove asteroid by swap and truncate
				g.asteroids = append(g.asteroids[:j], g.asteroids[j+1:]...)
				hit = true
				break
			}
		}
		if hit {
			continue
		}
		activeBullets = append(activeBullets, *b)
	}
	g.bullets = activeBullets
}

func (g *Game) updateAsteroids() {
	for i := range g.asteroids {
		a := &g.asteroids[i]

		a.position.Add(a.velocity)
		a.angle += a.rotSpeed

		// Screen wrap asteroid
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

func circleCollision(x1, y1, r1, x2, y2, r2 float64) bool {
	dx := x1 - x2
	dy := y1 - y2
	distance := math.Sqrt(dx*dx + dy*dy)
	return distance < r1+r2
}

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

func (g *Game) drawMenu(screen *ebiten.Image) {
	str := "ASTERÓIDES DA PAOLA 🌫️\n\n" +
		"Use setas esquerda/direita para girar\n" +
		"Seta para cima para acelerar\n" +
		"Espaço para atirar\n\n" +
		"Pressione ENTER para começar"
	drawTextCentered(screen, str, screenWidth/2, screenHeight/2, g.fontFace, textColor)
}

func (g *Game) drawPlaying(screen *ebiten.Image) {
	// Draw player
	opPlayer := &ebiten.DrawImageOptions{}
	opPlayer.GeoM.Translate(-g.player.width/2*scale, -g.player.height/2*scale)
	opPlayer.GeoM.Rotate(g.player.angle)
	opPlayer.GeoM.Translate(g.player.position.X, g.player.position.Y)
	opPlayer.GeoM.Scale(scale, scale)
	screen.DrawImage(g.player.img, opPlayer)

	// Draw asteroids
	for _, a := range g.asteroids {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-a.size/2, -a.size/2)
		op.GeoM.Rotate(a.angle)
		op.GeoM.Translate(a.position.X, a.position.Y)
		// Draw circle for asteroid (simple)
		img := generateCircleImage(int(a.size), color.RGBA{150, 150, 150, 255})
		screen.DrawImage(img, op)
	}

	// Draw bullets
	for _, b := range g.bullets {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(b.position.X-2, b.position.Y-2)
		bulletImg := generateCircleImage(4, color.RGBA{255, 255, 255, 255})
		screen.DrawImage(bulletImg, op)
	}

	// Draw explosions
	for _, e := range g.explosions {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.position.X-8, e.position.Y-8)
		expl := generateExplosionImage(e.frame, e.maxFrame)
		screen.DrawImage(expl, op)
	}

	scoreText := fmt.Sprintf("Pontos: %d", g.score)
	text.Draw(screen, scoreText, g.fontFace, 20, 40, textColor)
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	lines := []string{
		"🌟 GAME OVER 🌟",
		fmt.Sprintf("Pontos finais: %d", g.score),
		"Pressione R para jogar novamente",
	}

	for i, line := range lines {
		textBounds := text.BoundString(g.fontFace, line)
		x := screenWidth/2 - textBounds.Dx()/2
		y := screenHeight/2 - len(lines)*20 + i*40
		text.Draw(screen, line, g.fontFace, x, y, color.RGBA{255, 70, 70, 255})
	}
}

func drawTextCentered(screen *ebiten.Image, s string, x, y int, face font.Face, clr color.Color) {
	bounds := text.BoundString(face, s)
	text.Draw(screen, s, face, x-bounds.Dx()/2, y-bounds.Dy()/2, clr)
}

func generateCircleImage(d int, clr color.Color) *ebiten.Image {
	img := ebiten.NewImage(d, d)
	for y := 0; y < d; y++ {
		for x := 0; x < d; x++ {
			// Circle check
			dx := float64(x - d/2)
			dy := float64(y - d/2)
			if dx*dx+dy*dy <= float64(d*d)/4 {
				img.Set(x, y, clr)
			}
		}
	}
	return img
}

func generateExplosionImage(frame, maxFrame int) *ebiten.Image {
	d := 16
	img := ebiten.NewImage(d, d)
	alpha := uint8(255 * (1 - float64(frame)/float64(maxFrame)))
	for y := 0; y < d; y++ {
		for x := 0; x < d; x++ {
			img.Set(x, y, color.RGBA{255, 140, 30, alpha})
		}
	}
	return img
}

func loadSound(ctx *audio.Context, data []byte) *audio.Player {
	dec, err := mp3.Decode(ctx, bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	p, err := audio.NewPlayer(ctx, dec)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

// Layout implements ebiten.Game interface.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Asteroides da Paola - Profissional 🌫️")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

// These are base64-decoded mp3 sounds or can be loaded as raw []byte or external files.
// For now, they are placeholders. Replace them with actual sound data or files.
var shootSoundMp3 = []byte{
	// Put your shoot.mp3 bytes here or load from a file
}

var explosionSoundMp3 = []byte{
	// Put your explosion.mp3 bytes here or  load from a file
}

	// You can embed a TTF font file bytes here to have a nice font, e.g., OpenSans-Regular.ttf
	// or load via go:embed if you want.

// var fontTTF = []byte{} // Removed embedded font as no file is present
