package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

func loadFont() font.Face {
	return basicfont.Face7x13
}

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

func circleCollision(x1, y1, r1, x2, y2, r2 float64) bool {
	dx := x1 - x2
	dy := y1 - y2
	return dx*dx+dy*dy < (r1+r2)*(r1+r2)
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatalf("failed to load %s: %v", path, err)
	}
	return img
}
