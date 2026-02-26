# asteroid game

full-featured 2d space shooter built with ebiten game engine.

## features

- player movement and shooting mechanics
- asteroid spawning and collision detection
- score tracking system
- explosion effects
- power-up system
- sprite rendering

## architecture

### core components

- **game.go**: main game loop and state management
- **player.go**: player entity, movement, and shooting
- **asteroid.go**: asteroid behavior and collision
- **bullet.go**: projectile physics
- **entity.go**: base entity interface
- **vector.go**: 2d vector math utilities
- **config.go**: game constants and configuration

### design patterns

- entity-component pattern for game objects
- vector math for physics calculations
- sprite-based rendering
- frame-independent movement

## dependencies

```bash
go get github.com/hajimehoshi/ebiten/v2
```

## controls

- **arrow keys**: move ship
- **space**: shoot
- **esc**: quit

## running

```bash
cd jogoasteroide
go run .
```

## building

```bash
# build executable
go build -o asteroid.exe

# run executable
./asteroid.exe
```

## technical details

### collision detection
uses circular bounding boxes for efficient collision checks between entities.

### rendering
leverages ebiten's 2d rendering pipeline with sprite transformations.

### game loop
implements fixed timestep for consistent physics across different frame rates.

## future improvements

- [ ] sound effects and music
- [ ] multiple levels with increasing difficulty
- [ ] high score persistence
- [ ] particle effects
- [ ] enemy ships
- [ ] boss battles

## learning outcomes

this project demonstrates:
- game loop architecture
- entity management
- collision detection algorithms
- sprite rendering and transformations
- input handling
- state management
- vector mathematics in games
