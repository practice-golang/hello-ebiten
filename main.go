package main // import "hello-ebiten"

import (
	"image"
	"image/color"
	"log"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	eb "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

var (
	emptyImage    = ebiten.NewImage(3, 3)
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

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

func drawLine(screen *ebiten.Image, x1, y1, x2, y2 float32, c color.Color) {
	var path vector.Path

	path.MoveTo(x1, y1)
	path.LineTo(x2, y2)

	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0x33 / float32(0xff)
		vs[i].ColorG = 0x66 / float32(0xff)
		vs[i].ColorB = 0xff / float32(0xff)
	}
	screen.DrawTriangles(vs, is, emptySubImage, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x11, 0x22, 0x33, 0xff})

	drawLine(screen, 100, 100, 150, 150, color.White)
	eb.DrawRect(screen, 150, 150, 50, 50, color.White)
	eb.DebugPrint(screen, "Hello, World!")

	c := drawCircle(0, 0, 25)
	c.At(100, 100)
	em := ebiten.NewImageFromImage(c)

	opts := &ebiten.DrawImageOptions{}
	// opts.GeoM.Translate(1, 1)
	// opts.GeoM.Scale(1.5, 1)
	// opts.GeoM.Scale(1, 1)
	// opts.GeoM.Rotate(90 * 3.14 / 360)
	// opts.GeoM.Translate(0, -1)

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
