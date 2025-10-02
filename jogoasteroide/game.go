package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)





type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StatePaused
	StateGameOver
)

type BulletPool struct {
	pool []Bullet
}

func (p *BulletPool) Get() *Bullet {
	if len(p.pool) > 0 {
		b := &p.pool[len(p.pool)-1]
		p.pool = p.pool[:len(p.pool)-1]
		return b
	}
	return &Bullet{}
}

func (p *BulletPool) Put(b *Bullet) {
	*b = Bullet{} // reset
	p.pool = append(p.pool, *b)
}

type ExplosionPool struct {
	pool []Explosion
}

func (p *ExplosionPool) Get() *Explosion {
	if len(p.pool) > 0 {
		e := &p.pool[len(p.pool)-1]
		p.pool = p.pool[:len(p.pool)-1]
		return e
	}
	return &Explosion{}
}

func (p *ExplosionPool) Put(e *Explosion) {
	*e = Explosion{} // reset
	p.pool = append(p.pool, *e)
}

type Game struct {
	player             Player
	bullets            []*Bullet
	asteroids          []Asteroid
	explosions         []*Explosion
	powerUps           []*PowerUp
	bulletPool         BulletPool
	explosionPool      ExplosionPool
	powerUpPool        PowerUpPool
	score              int
	highScore          int
	state              GameState
	frames             int
	fontFace           font.Face
	playerW            float64
	playerH            float64
	asteroidW          float64
	asteroidH          float64
	scale              float64
	message            string
	messageTimer       int
	currentMaxAsteroids int
}

func NewGame() *Game {
	fontFace := loadFont()
	g := &Game{
		fontFace: fontFace,
		state:    StateMenu,
	}

	var err error
	ImgPlayer, err = loadImage("nave.png")
	if err != nil {
		log.Fatalf("failed to load nave.png: %v", err)
	}
	boundsPlayer := ImgPlayer.Bounds()
	g.playerW = float64(boundsPlayer.Dx())
	g.playerW = 64
	g.playerH = 64

	ImgAsteroid, err = loadImage("asteroide.png")
	if err != nil {
		log.Fatalf("failed to load asteroide.png: %v", err)
	}
	boundsAst := ImgAsteroid.Bounds()
	g.asteroidW = float64(boundsAst.Dx())
	g.asteroidH = float64(boundsAst.Dy())

	ImgBullet = generateCircleImage(12, BulletColor)
	ImgExplosion = generateCircleImage(40, ExplosionColor)
	ImgHealthBg = ebiten.NewImage(200, 20)
	ImgHealthBg.Fill(color.RGBA{255, 0, 0, 255})
	ImgCooldownBg = ebiten.NewImage(200, 20)
	ImgCooldownBg.Fill(color.RGBA{0, 0, 255, 255})

	g.player = Player{
		position: Vector{ScreenWidth / 2, ScreenHeight / 2},
		width:    g.playerW,
		height:   g.playerH,
		img:      ImgPlayer,
		health:   3,
	}

	return g
}

func (g *Game) Reset() {
	g.player = Player{
		position:     Vector{ScreenWidth / 2, ScreenHeight / 2},
		width:        g.playerW,
		height:       g.playerH,
		img:          ImgPlayer,
		velocity:     Vector{0, 0},
		angle:        0,
		fireCooldown: 0,
		health:       3,
		shield:       0,
		rapidFire:    0,
		multiShot:    0,
	}
	g.bullets = make([]*Bullet, 0, MaxBullets)
	g.asteroids = make([]Asteroid, 0, MaxAsteroids+50)
	g.explosions = make([]*Explosion, 0, 20)
	g.powerUps = make([]*PowerUp, 0, 10)
	g.bulletPool = BulletPool{}
	g.explosionPool = ExplosionPool{}
	g.powerUpPool = PowerUpPool{}
	g.score = 0
	g.state = StatePlaying
	g.frames = 0
	g.currentMaxAsteroids = MaxAsteroids
	for i := 0; i < MaxAsteroids; i++ {
		g.spawnAsteroid()
	}
}

func (g *Game) spawnAsteroid() {
	minSize := 40.0
	maxSize := 96.0
	size := minSize + rand.Float64()*(maxSize-minSize)
	pos := Vector{X: rand.Float64() * float64(ScreenWidth), Y: rand.Float64()*float64(ScreenHeight)/4 - size}
	speedMultiplier := 1.0 + float64(g.score)/5000.0 // Increase speed with score
	vel := Vector{X: (rand.Float64()*2 - 1) * 1.5 * speedMultiplier, Y: rand.Float64()*2 + 1 * speedMultiplier}
	rotSpeed := (rand.Float64()*2 - 1) * 0.04
	g.asteroids = append(g.asteroids, Asteroid{position: pos, velocity: vel, size: size, rotSpeed: rotSpeed})
}

func (g *Game) spawnPowerUp() {
	pos := Vector{X: rand.Float64() * float64(ScreenWidth), Y: rand.Float64()*float64(ScreenHeight)/4 - 20}
	vel := Vector{X: (rand.Float64()*2 - 1) * 1.0, Y: rand.Float64()*1.5 + 0.5}
	powerType := PowerUpType(rand.Intn(4)) // Random type
	size := 20.0
	maxAge := PowerUpMaxAge // 10 seconds at 60fps
	pw := g.powerUpPool.Get()
	pw.position = pos
	pw.velocity = vel
	pw.powerType = powerType
	pw.size = size
	pw.age = 0
	pw.maxAge = maxAge
	g.powerUps = append(g.powerUps, pw)
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
	cooldown := FireCooldown
	if g.player.rapidFire > 0 {
		cooldown = 5
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.player.fireCooldown <= 0 && len(g.bullets) < MaxBullets {
		g.fireBullet()
		g.player.fireCooldown = cooldown
	}
	g.updateBullets()
	g.updateAsteroids()
	g.updateExplosions()
	g.updatePowerUps()
	for _, a := range g.asteroids {
		if circleCollision(g.player.position.X, g.player.position.Y, g.player.width/2, a.position.X, a.position.Y, a.size/2) {
			if g.player.shield <= 0 {
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
			} else {
				g.message = "Escudo protegeu!"
				g.messageTimer = 120
			}
			break
		}
	}
	// Progressive difficulty: increase max asteroids based on score
	g.currentMaxAsteroids = MaxAsteroids + g.score/1000
	if len(g.asteroids) < g.currentMaxAsteroids && g.frames%60 == 0 {
		g.spawnAsteroid()
	}
	if g.frames%500 == 0 {
		g.spawnPowerUp()
	}
	// Check powerup collection
	for i, p := range g.powerUps {
		if circleCollision(g.player.position.X, g.player.position.Y, g.player.width/2, p.position.X, p.position.Y, p.size/2) {
			g.applyPowerUp(p.powerType)
			g.powerUpPool.Put(p)
			g.powerUps = append(g.powerUps[:i], g.powerUps[i+1:]...)
			break
		}
	}
}

func (g *Game) fireBullet() {
	offsetX := math.Sin(g.player.angle) * g.player.height / 2
	offsetY := -math.Cos(g.player.angle) * g.player.height / 2
	bulletPos := Vector{X: g.player.position.X + offsetX, Y: g.player.position.Y + offsetY}
	if g.player.multiShot > 0 {
		angles := []float64{g.player.angle, g.player.angle - 0.2, g.player.angle + 0.2}
		for _, ang := range angles {
			bulletVel := Vector{X: math.Sin(ang) * BulletSpeed, Y: -math.Cos(ang) * BulletSpeed}
			b := g.bulletPool.Get()
			b.position = bulletPos
			b.velocity = bulletVel
			b.age = 0
			g.bullets = append(g.bullets, b)
		}
	} else {
		bulletVel := Vector{X: math.Sin(g.player.angle) * BulletSpeed, Y: -math.Cos(g.player.angle) * BulletSpeed}
		b := g.bulletPool.Get()
		b.position = bulletPos
		b.velocity = bulletVel
		b.age = 0
		g.bullets = append(g.bullets, b)
	}
}

func (g *Game) updateBullets() {
	active := g.bullets[:0]
	for _, b := range g.bullets {
		b.Update()
		if b.age > BulletMaxAge || b.IsOffScreen() {
			g.bulletPool.Put(b)
			continue
		}
		hit := false
		for j, a := range g.asteroids {
			if circleCollision(b.position.X, b.position.Y, 5, a.position.X, a.position.Y, a.size/2) {
				e := g.explosionPool.Get()
				e.position = a.position
				e.maxFrame = ExplosionFrames
				g.explosions = append(g.explosions, e)
				g.score += int(a.size) * 10
				if a.size > MinAsteroidSize {
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
		if hit {
			g.bulletPool.Put(b)
		} else {
			active = append(active, b)
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
	for _, e := range g.explosions {
		e.Update()
		if e.frame < e.maxFrame {
			active = append(active, e)
		} else {
			g.explosionPool.Put(e)
		}
	}
	g.explosions = active
}

func (g *Game) updatePowerUps() {
	active := g.powerUps[:0]
	for _, p := range g.powerUps {
		p.Update()
		if p.IsExpired() {
			g.powerUpPool.Put(p)
		} else {
			active = append(active, p)
		}
	}
	g.powerUps = active
}

func (g *Game) applyPowerUp(powerType PowerUpType) {
	switch powerType {
	case PowerUpShield:
		g.player.shield = 600 // 10 seconds
		g.message = "Escudo ativado!"
		g.messageTimer = 120
	case PowerUpRapidFire:
		g.player.rapidFire = 600
		g.message = "Tiro rápido ativado!"
		g.messageTimer = 120
	case PowerUpMultiShot:
		g.player.multiShot = 600
		g.message = "Tiro múltiplo ativado!"
		g.messageTimer = 120
	case PowerUpExtraLife:
		g.player.health++
		if g.player.health > 3 {
			g.player.health = 3
		}
		g.message = "Vida extra!"
		g.messageTimer = 120
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BgColor)
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
	y := ScreenHeight / 2
	text.Draw(screen, title, g.fontFace, ScreenWidth/2-len(title)*7, y-80, TextColor)
	text.Draw(screen, instr, g.fontFace, ScreenWidth/2-170, y-40, TextColor)
	text.Draw(screen, fmt.Sprintf("Melhor pontuação: %d", g.highScore), g.fontFace, ScreenWidth/2-100, y-120, TextColor)
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
	for _, p := range g.powerUps {
		p.Draw(screen)
	}
	text.Draw(screen, fmt.Sprintf("Pontos: %d", g.score), g.fontFace, 24, 40, TextColor)
	text.Draw(screen, fmt.Sprintf("Melhor: %d", g.highScore), g.fontFace, 24, 70, TextColor)

	// Draw health bar
	healthBarWidth := 200.0
	healthBarHeight := 20.0
	healthBarX := 24.0
	healthBarY := 100.0
	healthFg := ebiten.NewImage(int(healthBarWidth*float64(g.player.health)/3), int(healthBarHeight))
	healthFg.Fill(color.RGBA{0, 255, 0, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(healthBarX, healthBarY)
	screen.DrawImage(ImgHealthBg, op)
	screen.DrawImage(healthFg, op)

	// Draw cooldown bar
	cooldownBarWidth := 200.0
	cooldownBarHeight := 20.0
	cooldownBarX := 24.0
	cooldownBarY := 130.0
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(cooldownBarX, cooldownBarY)
	screen.DrawImage(ImgCooldownBg, op2)
	if g.player.fireCooldown > 0 {
		cooldownFg := ebiten.NewImage(int(cooldownBarWidth*float64(g.player.fireCooldown)/FireCooldown), int(cooldownBarHeight))
		cooldownFg.Fill(color.RGBA{255, 255, 0, 255})
		screen.DrawImage(cooldownFg, op2)
	}

	// Draw message if any
	if g.messageTimer > 0 {
		bounds := text.BoundString(g.fontFace, g.message)
		x := ScreenWidth/2 - bounds.Dx()/2
		y := ScreenHeight/2 - bounds.Dy()/2
		text.Draw(screen, g.message, g.fontFace, x, y, color.RGBA{255, 0, 0, 255})
		g.messageTimer--
	}
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	lines := []string{"🌟 FIM DE JOGO 🌟", fmt.Sprintf("Pontos finais: %d", g.score), fmt.Sprintf("Melhor pontuação: %d", g.highScore), "Pressione R para tentar novamente"}
	y := ScreenHeight / 2
	for i, line := range lines {
		bounds := text.BoundString(g.fontFace, line)
		x := ScreenWidth/2 - bounds.Dx()/2
		text.Draw(screen, line, g.fontFace, x, y+i*32, color.RGBA{255, 69, 0, 255})
	}
}

func (g *Game) drawPaused(screen *ebiten.Image) {
	text.Draw(screen, "PAUSADO - Pressione P para continuar", g.fontFace, ScreenWidth/2-150, ScreenHeight/2, TextColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.scale = math.Min(float64(outsideWidth)/ScreenWidth, float64(outsideHeight)/ScreenHeight)
	return outsideWidth, outsideHeight
}
