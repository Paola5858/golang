// jogo/main.go
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

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth     = 800
	screenHeight    = 600
	playerSpeed     = 4
	asteroidBaseSpeed = 2
	numAsteroids    = 10
	playerScale     = 0.12
	asteroidScale   = 0.15
	explosionFrames = 18
	soundVolume     = 0.15
)

// Colors used in game
var (
	bgColor          = color.RGBA{20, 20, 40, 255}
	player1Color     = color.RGBA{50, 180, 255, 255}
	player2Color     = color.RGBA{255, 80, 80, 255}
	// Removed unused variable scoreTextColor to fix the compile error
	gameOverTextColor= color.RGBA{255, 40, 40, 255}
	menuTextColor    = color.RGBA{220, 220, 220, 255}
	buttonColor      = color.RGBA{70, 160, 255, 255}
	buttonHoverColor = color.RGBA{110, 210, 255, 255}
)

// GameState enumerates the state of the game flow
type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StateGameOver
)

// ControlKeys defines the keys for player movement
type ControlKeys struct {
	Left, Right, Up, Down ebiten.Key
}

// Layout defines the screen dimensions for the game.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Player represents a player's ship
type Player struct {
	X, Y        float64
	Img         *ebiten.Image
	Width       float64
	Height      float64
	Color       color.Color
	Score       int
	Keys        ControlKeys
	IsExploding bool
	Explosion   Explosion
}

// Asteroid represents an asteroid object
type Asteroid struct {
	X, Y        float64
	Speed       float64
	Img         *ebiten.Image
	Width       float64
	Height      float64
}

// Explosion represents an explosion effect at a position
type Explosion struct {
	X, Y       float64
	Frame      int
	IsPlaying  bool
	MaxFrames  int
}

// Button represents an interactive UI button
type Button struct {
	X, Y, W, H float64
	Label      string
	IsHovered  bool
}

// Game encapsulates the entire game state
type Game struct {
	Player1        Player
	Player2        Player
	Asteroids      []Asteroid
	PlayerImg      *ebiten.Image
	AsteroidImgs   []*ebiten.Image
	Explosions     []Explosion
	State          GameState
	StartButton    Button
	RestartButton  Button
	AudioContext   *audio.Context
	ExplosionSound *audio.Player
	StartSound     *audio.Player
	GameOverSound  *audio.Player
	StartTime      time.Time
	Duration       float64
	Difficulty     float64
	FPS            float64
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Asteroides da Paola 💫 - 2 Players")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

// NewGame creates a new game instance setting initial values and generating resources
func NewGame() *Game {
	audioCtx := audio.NewContext(44100)
	playerImg := generatePlayerImage()
	asteroid1Img := generateAsteroidImage(false)
	asteroid2Img := generateAsteroidImage(true)

	explosionSound := loadAudioPlayer(audioCtx, generateExplosionSound())
	startSound := loadAudioPlayer(audioCtx, generateStartSound())
	gameOverSound := loadAudioPlayer(audioCtx, generateGameOverSound())

	player1 := Player{
		X:      screenWidth * 0.3,
		Y:      screenHeight - 120,
		Img:    playerImg,
		Width:  float64(playerImg.Bounds().Dx()) * playerScale,
		Height: float64(playerImg.Bounds().Dy()) * playerScale,
		Color:  player1Color,
		Score:  0,
		Keys: ControlKeys{
			Left:  ebiten.KeyArrowLeft,
			Right: ebiten.KeyArrowRight,
			Up:    ebiten.KeyArrowUp,
			Down:  ebiten.KeyArrowDown,
		},
	}

	player2 := Player{
		X:      screenWidth * 0.7,
		Y:      screenHeight - 120,
		Img:    playerImg,
		Width:  float64(playerImg.Bounds().Dx()) * playerScale,
		Height: float64(playerImg.Bounds().Dy()) * playerScale,
		Color:  player2Color,
		Score:  0,
		Keys: ControlKeys{
			Left:  ebiten.KeyA,
			Right: ebiten.KeyD,
			Up:    ebiten.KeyW,
			Down:  ebiten.KeyS,
		},
	}

	asteroids := make([]Asteroid, numAsteroids)
	for i := range asteroids {
		img := asteroid1Img
		if i%2 == 0 {
			img = asteroid2Img
		}
		aw := float64(img.Bounds().Dx()) * asteroidScale
		ah := float64(img.Bounds().Dy()) * asteroidScale
		asteroids[i] = Asteroid{
			X:      float64(rand.Intn(screenWidth-int(aw))),
			Y:      float64(rand.Intn(screenHeight/2)) - ah*float64(rand.Intn(5)),
			Speed:  asteroidBaseSpeed + rand.Float64()*2,
			Img:    img,
			Width:  aw,
			Height: ah,
		}
	}

	return &Game{
		Player1:        player1,
		Player2:        player2,
		Asteroids:      asteroids,
		PlayerImg:      playerImg,
		AsteroidImgs:   []*ebiten.Image{asteroid1Img, asteroid2Img},
		Explosions:     []Explosion{},
		State:          StateMenu,
		StartButton:    Button{X: float64(screenWidth/2 - 100), Y: float64(screenHeight/2 - 25), W: 200, H: 60, Label: "START (Enter)"},
		RestartButton:  Button{X: float64(screenWidth/2 - 105), Y: float64(screenHeight/2 + 40), W: 210, H: 60, Label: "REINICIAR (R)"},
		AudioContext:   audioCtx,
		ExplosionSound: explosionSound,
		StartSound:     startSound,
		GameOverSound:  gameOverSound,
		Duration:       0,
		Difficulty:     1,
	}
}

// Update is the main game update loop called each frame
func (g *Game) Update() error {
	g.FPS = ebiten.CurrentFPS()

	switch g.State {
	case StateMenu:
		g.handleMenuInput()
	case StatePlaying:
		g.updatePlaying()
	case StateGameOver:
		g.handleGameOverInput()
	}
	return nil
}

func (g *Game) handleMenuInput() {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.startGame()
	}
}

func (g *Game) handleGameOverInput() {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.resetGame()
	}
}

func (g *Game) startGame() {
	g.State = StatePlaying
	g.StartSound.Rewind()
	g.StartSound.Play()
	g.StartTime = time.Now()
	g.Duration = 0
	g.Difficulty = 1
	g.Player1.Score = 0
	g.Player2.Score = 0
	g.Player1.X = screenWidth * 0.3
	g.Player1.Y = screenHeight - 120
	g.Player2.X = screenWidth * 0.7
	g.Player2.Y = screenHeight - 120
	for i := range g.Asteroids {
		ast := &g.Asteroids[i]
		ast.X = float64(rand.Intn(screenWidth - int(ast.Width)))
		ast.Y = float64(rand.Intn(screenHeight/2)) - ast.Height*float64(rand.Intn(5))
		ast.Speed = asteroidBaseSpeed + rand.Float64()*2
	}
	g.Explosions = []Explosion{}
}

func (g *Game) resetGame() {
	g.State = StateMenu
}

func (g *Game) updatePlaying() {
	g.Duration = time.Since(g.StartTime).Seconds()
	g.Difficulty = 1 + g.Duration/30 // difficulty scale over time

	g.Player1.update(g.Difficulty)
	g.Player2.update(g.Difficulty)

	g.updateAsteroids()

	g.updateExplosions()
}

func (g *Game) updateAsteroids() {
	for i := range g.Asteroids {
		ast := &g.Asteroids[i]
		ast.Y += ast.Speed * g.Difficulty
		if ast.Y > screenHeight {
			ast.Y = -ast.Height
			ast.X = float64(rand.Intn(screenWidth-int(ast.Width)))
			ast.Speed = asteroidBaseSpeed + rand.Float64()*2
			// Increment score alternately per asteroid index
			if i%2 == 0 {
				g.Player1.Score++
			} else {
				g.Player2.Score++
			}
		}

		if collide(g.Player1.X, g.Player1.Y, g.Player1.Width, g.Player1.Height,
			ast.X, ast.Y, ast.Width, ast.Height) && !g.Player1.IsExploding {
			g.triggerExplosion(g.Player1.X+g.Player1.Width/2, g.Player1.Y+g.Player1.Height/2)
			g.Player1.IsExploding = true
			g.GameOver()
		}

		if collide(g.Player2.X, g.Player2.Y, g.Player2.Width, g.Player2.Height,
			ast.X, ast.Y, ast.Width, ast.Height) && !g.Player2.IsExploding {
			g.triggerExplosion(g.Player2.X+g.Player2.Width/2, g.Player2.Y+g.Player2.Height/2)
			g.Player2.IsExploding = true
			g.GameOver()
		}
	}
}

func (g *Game) updateExplosions() {
	newExplosions := make([]Explosion, 0, len(g.Explosions))
	for i := range g.Explosions {
		exp := &g.Explosions[i]
		if exp.IsPlaying {
			exp.Frame++
			if exp.Frame >= exp.MaxFrames {
				exp.IsPlaying = false
			} else {
				newExplosions = append(newExplosions, *exp)
			}
		}
	}
	g.Explosions = newExplosions
}

func (g *Game) triggerExplosion(x, y float64) {
	g.ExplosionSound.Rewind()
	g.ExplosionSound.Play()
	g.Explosions = append(g.Explosions, Explosion{
		X:        x,
		Y:        y,
		Frame:    0,
		IsPlaying:true,
		MaxFrames: explosionFrames,
	})
}

func (g *Game) GameOver() {
	if g.State != StateGameOver {
		g.State = StateGameOver
		g.GameOverSound.Rewind()
		g.GameOverSound.Play()
	}
}

// Draw is the main game draw loop called each frame to render screen contents
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	switch g.State {
	case StateMenu:
		g.drawMenu(screen)
	case StatePlaying:
		g.drawPlaying(screen)
	case StateGameOver:
		g.drawPlaying(screen)
		g.drawGameOver(screen)
	}

	g.drawFPS(screen)
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	title := "ASTEROIDES DA PAOLA 💫"
	instruction1 := "Player 1 (Azul): Setas ← ↑ → ↓"
	instruction2 := "Player 2 (Vermelho): WASD"
	instruction3 := "Pressione ENTER para iniciar o jogo"

	ebitenutil.DebugPrintAt(screen, title, screenWidth/2-90, screenHeight/4)
	ebitenutil.DebugPrintAt(screen, instruction1, screenWidth/2-130, screenHeight/4+40)
	ebitenutil.DebugPrintAt(screen, instruction2, screenWidth/2-130, screenHeight/4+80)
	ebitenutil.DebugPrintAt(screen, instruction3, screenWidth/2-140, screenHeight/4+120)

	g.drawButton(screen, &g.StartButton)
}

func (g *Game) drawPlaying(screen *ebiten.Image) {
	g.Player1.draw(screen)
	g.Player2.draw(screen)

	for _, ast := range g.Asteroids {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(asteroidScale, asteroidScale)
		op.GeoM.Translate(ast.X, ast.Y)
		screen.DrawImage(ast.Img, op)
	}

	for _, e := range g.Explosions {
		if e.IsPlaying {
			g.drawExplosion(screen, e)
		}
	}

	g.drawScore(screen)
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	msg := "GAME OVER!"
	ebitenutil.DrawRect(screen, screenWidth/2-80, screenHeight/2-110, 160, 30, gameOverTextColor)
	ebitenutil.DebugPrintAt(screen, msg, screenWidth/2-70, screenHeight/2-100)

	score1Msg := fmt.Sprintf("Player 1 (Azul) Pontos: %d", g.Player1.Score)
	score2Msg := fmt.Sprintf("Player 2 (Vermelho) Pontos: %d", g.Player2.Score)
	restartMsg := "Pressione 'R' para reiniciar"

	ebitenutil.DebugPrintAt(screen, score1Msg, screenWidth/2-110, screenHeight/2-50)
	ebitenutil.DebugPrintAt(screen, score2Msg, screenWidth/2-110, screenHeight/2-20)
	ebitenutil.DebugPrintAt(screen, restartMsg, screenWidth/2-120, screenHeight/2+30)

	g.drawButton(screen, &g.RestartButton)
}

func (g *Game) drawButton(screen *ebiten.Image, btn *Button) {
	// Update hover
	mx, my := ebiten.CursorPosition()
	btn.IsHovered = float64(mx) >= btn.X && float64(mx) <= btn.X+btn.W &&
		float64(my) >= btn.Y && float64(my) <= btn.Y+btn.H

	col := buttonColor
	if btn.IsHovered {
		col = buttonHoverColor
	}

	ebitenutil.DrawRect(screen, btn.X, btn.Y, btn.W, btn.H, col)
	textX := int(btn.X + btn.W/2 - float64(len(btn.Label)*7)/2)
	textY := int(btn.Y + btn.H/2 - 10)
	ebitenutil.DebugPrintAt(screen, btn.Label, textX, textY)
}

func (g *Game) drawExplosion(screen *ebiten.Image, e Explosion) {
	alpha := 1 - float64(e.Frame)/float64(e.MaxFrames)
	radius := 10 + float64(e.Frame)*3
	col := color.RGBA{255, 180, 20, uint8(255 * alpha)}

	drawCircle(screen, e.X, e.Y, radius, col)
}

func (g *Game) drawScore(screen *ebiten.Image) {
	// Shadows for readability
	// Removed unused variable shadowColor to fix the compile error

	msg1 := fmt.Sprintf("Player 1 (Azul): %d", g.Player1.Score)
	msg2 := fmt.Sprintf("Player 2 (Vermelho): %d", g.Player2.Score)

	x1, y1 := 16, 16
	x2, y2 := screenWidth-170, 16

	ebitenutil.DebugPrintAt(screen, msg1, x1+1, y1+1)
	ebitenutil.DebugPrintAt(screen, msg2, x2+1, y2+1)

	ebitenutil.DebugPrintAt(screen, msg1, x1, y1)
	ebitenutil.DebugPrintAt(screen, msg2, x2, y2)
}

func (g *Game) drawFPS(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %0.0f", g.FPS), 10, screenHeight-24)
}

// Update player position according to input and difficulty
func (p *Player) update(difficulty float64) {
	if p.IsExploding {
		return
	}

	if ebiten.IsKeyPressed(p.Keys.Left) && p.X > 0 {
		p.X -= playerSpeed * difficulty
		if p.X < 0 {
			p.X = 0
		}
	}
	if ebiten.IsKeyPressed(p.Keys.Right) && p.X < screenWidth-p.Width {
		p.X += playerSpeed * difficulty
		if p.X > screenWidth-p.Width {
			p.X = screenWidth - p.Width
		}
	}
	if ebiten.IsKeyPressed(p.Keys.Up) && p.Y > 0 {
		p.Y -= playerSpeed * difficulty
		if p.Y < 0 {
			p.Y = 0
		}
	}
	if ebiten.IsKeyPressed(p.Keys.Down) && p.Y < screenHeight-p.Height {
		p.Y += playerSpeed * difficulty
		if p.Y > screenHeight-p.Height {
			p.Y = screenHeight - p.Height
		}
	}
}

// Draw player with color tint and scaling
func (p *Player) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(playerScale, playerScale)
	op.GeoM.Translate(p.X, p.Y)

	var img *ebiten.Image

	if p.IsExploding {
		// Draw explosion effect instead of player
		// (handled by Game.Draw explosions)
		return
	} else {
		img = p.Img
	}

	screen.DrawImage(tintImage(img, p.Color), op)
}

// Helper function to generate a simple player ship image procedurally
func generatePlayerImage() *ebiten.Image {
	const size = 64
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	// Clear background transparent
	for i := 0; i < len(img.Pix); i++ {
		img.Pix[i] = 0
	}

	// Draw simple triangular ship in white
	for y := 0; y < size; y++ {
		xCenter := size / 2
		width := y / 2
		for x := xCenter - width; x <= xCenter+width; x++ {
			img.Set(x, y, color.White)
		}
	}

	return ebiten.NewImageFromImage(img)
}

// Helper function to generate simple asteroid images (random gray circle)
func generateAsteroidImage(variant bool) *ebiten.Image {
	const size = 64
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	radius := size / 2

	grayBase := uint8(140)
	if variant {
		grayBase = 180
	}
	// Draw circle with random stones "texture"
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			dx := x - radius
			dy := y - radius
			dist := math.Sqrt(float64(dx*dx + dy*dy))
			if dist <= float64(radius) {
				shade := grayBase + uint8(rand.Intn(40))
				img.Set(x, y, color.RGBA{shade, shade, shade, 255})
			} else {
				img.Set(x, y, color.Transparent)
			}
		}
	}

	return ebiten.NewImageFromImage(img)
}

// circle drawing helper (filled circle)
func drawCircle(screen *ebiten.Image, x, y, radius float64, clr color.Color) {
	const segments = 64
	angleStep := 2 * math.Pi / segments
	for i := 0; i < segments; i++ {
		theta1 := float64(i) * angleStep
		theta2 := float64(i+1) * angleStep
		x1 := x + math.Cos(theta1)*radius
		y1 := y + math.Sin(theta1)*radius
		x2 := x + math.Cos(theta2)*radius
		y2 := y + math.Sin(theta2)*radius
		ebitenutil.DrawLine(screen, x, y, x1, y1, clr)
		ebitenutil.DrawLine(screen, x, y, x2, y2, clr)
		ebitenutil.DrawLine(screen, x1, y1, x2, y2, clr)
	}
}

// Simple rectangular collision detection
func collide(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	return x1 < x2+w2 &&
		x1+w1 > x2 &&
		y1 < y2+h2 &&
		y1+h1 > y2
}

// A simple function to tint an image with a given color
func tintImage(img *ebiten.Image, tint color.Color) *ebiten.Image {
	r, g, b, a := tint.RGBA()
	r8 := float64(r>>8) / 255
	g8 := float64(g>>8) / 255
	b8 := float64(b>>8) / 255
	a8 := float64(a>>8) / 255

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(r8, g8, b8, a8)
	w, h := img.Size()
	dst := ebiten.NewImage(w, h)
	dst.DrawImage(img, op)
	return dst
}

// Load audio player from PCM wave data in memory
func loadAudioPlayer(ctx *audio.Context, wavData []byte) *audio.Player {
	d, err := wav.Decode(ctx, bytes.NewReader(wavData))
	if err != nil {
		log.Fatalf("Failed to decode audio: %v", err)
	}
	p, err := audio.NewPlayer(ctx, d)
	if err != nil {
		log.Fatalf("Failed to create audio player: %v", err)
	}
	p.SetVolume(soundVolume)
	return p
}

// -- Sound generators below --

// generateExplosionSound creates a simple explosion sound waveform (white noise burst)
func generateExplosionSound() []byte {
	const sampleRate = 44100
	const length = sampleRate / 5 // 0.2 seconds
	buf := make([]byte, length*2*2) // stereo 16-bit little endian

	// Generate white noise burst with exponential decay envelope
	for i := 0; i < length; i++ {
		decay := math.Exp(-4 * float64(i) / float64(length))
		sample := int16((rand.Float64()*2 - 1) * decay * 30000)

		// write left channel
		buf[4*i] = byte(sample)
		buf[4*i+1] = byte(sample >> 8)

		// write right channel (same)
		buf[4*i+2] = byte(sample)
		buf[4*i+3] = byte(sample >> 8)
	}

	// Return as valid WAV
	return createWav(sampleRate, 2, 16, buf)
}

// generateStartSound creates a simple beep sound
func generateStartSound() []byte {
	const sampleRate = 44100
	const length = sampleRate / 10
	buf := make([]byte, length*2*2) // stereo

	freq := 440.0 // A4 tone
	for i := 0; i < length; i++ {
		val := int16(30000 * math.Sin(2*math.Pi*freq*float64(i)/float64(sampleRate)))
		buf[4*i] = byte(val)
		buf[4*i+1] = byte(val >> 8)
		buf[4*i+2] = byte(val)
		buf[4*i+3] = byte(val >> 8)
	}

	return createWav(sampleRate, 2, 16, buf)
}

// generateGameOverSound creates a short downward tone
func generateGameOverSound() []byte {
	const sampleRate = 44100
	const length = sampleRate / 8
	buf := make([]byte, length*2*2)

	for i := 0; i < length; i++ {
		// descending frequency from 700Hz to 200Hz
		freq := 700 - 500*float64(i)/float64(length)
		val := int16(30000 * math.Sin(2*math.Pi*freq*float64(i)/float64(sampleRate)))
		buf[4*i] = byte(val)
		buf[4*i+1] = byte(val >> 8)
		buf[4*i+2] = byte(val)
		buf[4*i+3] = byte(val >> 8)
	}

	return createWav(sampleRate, 2, 16, buf)
}

// createWav generates a minimal valid WAV header followed by raw PCM data (16bit) in little endian stereo
func createWav(sampleRate, channels, bitDepth int, pcm []byte) []byte {
	// WAV Header format
	header := make([]byte, 44)

	// ChunkID "RIFF"
	header[0] = 'R'
	header[1] = 'I'
	header[2] = 'F'
	header[3] = 'F'

	// ChunkSize (36 + SubChunk2Size) = 36 + pcm length
	chunkSize := 36 + len(pcm)
	header[4] = byte(chunkSize)
	header[5] = byte(chunkSize >> 8)
	header[6] = byte(chunkSize >> 16)
	header[7] = byte(chunkSize >> 24)

	// Format "WAVE"
	header[8] = 'W'
	header[9] = 'A'
	header[10] = 'V'
	header[11] = 'E'

	// Subchunk1ID "fmt "
	header[12] = 'f'
	header[13] = 'm'
	header[14] = 't'
	header[15] = ' '

	// Subchunk1Size 16 for PCM
	header[16] = 16
	header[17] = 0
	header[18] = 0
	header[19] = 0

	// AudioFormat PCM=1
	header[20] = 1
	header[21] = 0

	// NumChannels
	header[22] = byte(channels)
	header[23] = 0

	// SampleRate
	header[24] = byte(sampleRate)
	header[25] = byte(sampleRate >> 8)
	header[26] = byte(sampleRate >> 16)
	header[27] = byte(sampleRate >> 24)

	// ByteRate = SampleRate * NumChannels * BitsPerSample/8
	byteRate := sampleRate * channels * bitDepth / 8
	header[28] = byte(byteRate)
	header[29] = byte(byteRate >> 8)
	header[30] = byte(byteRate >> 16)
	header[31] = byte(byteRate >> 24)

	// BlockAlign = NumChannels * BitsPerSample/8
	blockAlign := channels * bitDepth / 8
	header[32] = byte(blockAlign)
	header[33] = 0

	// BitsPerSample
	header[34] = byte(bitDepth)
	header[35] = 0

	// Subchunk2ID "data"
	header[36] = 'd'
	header[37] = 'a'
	header[38] = 't'
	header[39] = 'a'

	// Subchunk2Size = NumSamples * NumChannels * BitsPerSample/8
	subChunk2Size := len(pcm)
	header[40] = byte(subChunk2Size)
	header[41] = byte(subChunk2Size >> 8)
	header[42] = byte(subChunk2Size >> 16)
	header[43] = byte(subChunk2Size >> 24)

	// Compose complete WAV
	return append(header, pcm...)
}
