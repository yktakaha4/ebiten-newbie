// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build example
// +build example

package main

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	gophersImage        *ebiten.Image
	gophersRenderTarget *ebiten.Image
)

func init() {
	// Decode image from a byte slice instead of a file so that
	// this example works in any working directory.
	// If you want to use a file, there are some options:
	// 1) Use os.Open and pass the file to the image decoder.
	//    This is a very regular way, but doesn't work on browsers.
	// 2) Use ebitenutil.OpenFile and pass the file to the image decoder.
	//    This works even on browsers.
	// 3) Use ebitenutil.NewImageFromFile to create an ebiten.Image directly from a file.
	//    This also works on browsers.
	img, _, err := image.Decode(bytes.NewReader(images.Gophers_jpg))
	if err != nil {
		log.Fatal(err)
	}
	gophersImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	mosaicRatio int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		if g.mosaicRatio < 255 {
			g.mosaicRatio++
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		if g.mosaicRatio > 1 {
			g.mosaicRatio--
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Shrink the image once.
	op := &ebiten.DrawImageOptions{}
	w, h := gophersImage.Size()
	gophersRenderTarget = ebiten.NewImage(w/g.mosaicRatio, h/g.mosaicRatio)
	gophersRenderTarget.DrawImage(gophersImage, op)

	// Enlarge the shrunk image.
	// The filter is the nearest filter, so the result will be mosaic.
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(g.mosaicRatio), float64(g.mosaicRatio))
	screen.DrawImage(gophersRenderTarget, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	initialMosaicRatio := 16
	w, h := gophersImage.Size()
	gophersRenderTarget = ebiten.NewImage(w/initialMosaicRatio, h/initialMosaicRatio)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Mosaic (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{
		initialMosaicRatio,
	}); err != nil {
		log.Fatal(err)
	}
}
