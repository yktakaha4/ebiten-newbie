package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	tick int64
}

func (g *Game) Update() error {
	g.tick += 1
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if (g.tick/10)%2 == 0 {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello, World! tick=%v", g.tick))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
