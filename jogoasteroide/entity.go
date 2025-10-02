package main

import "github.com/hajimehoshi/ebiten/v2"

// Entity interface defines common behaviors for game entities.
type Entity interface {
	Update()
	Draw(screen *ebiten.Image)
}
