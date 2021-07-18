package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"math"
)

type Game struct{
	tick int64
}

func (g *Game) Update() error {
	g.tick += 1
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const speed = 10
	r := math.Abs(float64((g.tick * speed) % 511 - 255))

	screen.Fill(color.RGBA{R: uint8(r), A: 0xff})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Fill")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
