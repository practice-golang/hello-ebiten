package main // import "hello-ebiten"

import (
	"image"
	"image/color"
	"log"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	eb "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func drawCircle(x, y, r int) image.Image {
	dc := gg.NewContext(int(x+(r*2)), int(y+(r*2)))

	dc.Push()
	dc.SetRGBA(1, 0.5, 0, 1)
	dc.SetLineWidth(1)
	dc.DrawCircle(float64(r), float64(r), float64(r))
	dc.Fill()
	dc.Pop()

	return dc.Image()
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x11, 0x22, 0x33, 0xff})

	eb.DrawLine(screen, -50, -50, 50, 50, color.White)
	eb.DrawRect(screen, 150, 150, 50, 50, color.Black)
	eb.DebugPrint(screen, "Hello, World!")

	m := drawCircle(100, 100, 50)
	em := ebiten.NewImageFromImage(m)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, 0)
	opts.GeoM.Scale(1.5, 1)
	screen.DrawImage(em, opts)
}

func (g *Game) Layout(outsideW, outsideH int) (screenW, screenH int) {
	return 320, 240
}

func main() {
	app := &Game{}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, Ebiten!!")

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
