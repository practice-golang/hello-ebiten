package main // import "triangles"

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

type GameConfig struct {
	Width  int
	Height int
	Title  string
}

var (
	CFG = GameConfig{
		Width:  800,
		Height: 600,
		Title:  "Bugs?",
	}
)

var (
	emptyImage        = ebiten.NewImage(3, 3)
	debugCircleImage  *ebiten.Image
	emptyTextureImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
	face              font.Face
)

func NormalizeInt(v, fromMin, fromMax, toMin, toMax int) int {
	ratio := float64(toMax-toMin) / float64(fromMax-fromMin)
	return int(math.Round(float64(toMin) + ratio*float64(v-fromMin)))
}

func init() {
	emptyImage.Fill(color.White)
	img := image.NewRGBA(image.Rect(0, 0, 20, 20))
	gc := draw2dimg.NewGraphicContext(img)
	gc.SetStrokeColor(color.NRGBA{255, 0, 0, 255})
	gc.SetLineWidth(1)
	draw2dkit.Circle(gc, 10, 10, 7)
	gc.Stroke()
	debugCircleImage = ebiten.NewImageFromImage(img)
	emptyImage.Fill(colornames.Aquamarine)
	f, _ := truetype.Parse(goregular.TTF)
	face = truetype.NewFace(f, &truetype.Options{
		Size: 12,
		DPI:  72,
	})
}

type Location struct {
	X int
	Y int
}

type Circle struct {
	id     int
	size   int
	color  color.RGBA
	center Location
}

type Game struct {
	screenWidth, screenHeight int
	circles                   []Circle
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render non-vector parts first.
	for _, c := range g.circles {
		op := &ebiten.DrawImageOptions{
			GeoM:          ebiten.GeoM{},
			ColorM:        ebiten.ColorM{},
			CompositeMode: 0,
			Filter:        0,
		}
		op.GeoM.Translate(float64(c.center.X), float64(c.center.Y))
		screen.DrawImage(debugCircleImage, op)
		cord := fmt.Sprintf("%d, %.2f, %.2f, %.2f", c.id, float32(c.center.X), float32(c.center.Y), float32(c.size)/2)
		text.Draw(screen, cord, face, c.center.X, c.center.Y, colornames.White)
	}

	// Render vector parts second.
	for _, c := range g.circles {
		p := vector.Path{}
		p.Arc(float32(c.center.X), float32(c.center.Y), float32(c.size)/2, 0, 2*math.Pi, vector.Clockwise)
		filling, indicies := p.AppendVerticesAndIndicesForFilling(nil, nil)
		for i := range filling {
			filling[i].ColorR = float32(c.color.R) / 0xff
			filling[i].ColorG = float32(c.color.G) / 0xff
			filling[i].ColorB = float32(c.color.B) / 0xff
		}
		op := &ebiten.DrawTrianglesOptions{
			FillRule: ebiten.EvenOdd,
		}
		screen.DrawTriangles(filling, indicies, emptyTextureImage, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth = g.screenWidth
	screenHeight = g.screenHeight
	return
}

func NewGame(width, height, initNumberOfOrganisms int) *Game {
	g := &Game{
		screenWidth:  width,
		screenHeight: height,
		circles:      make([]Circle, 0, initNumberOfOrganisms),
	}
	for i := 0; i < initNumberOfOrganisms; i += 1 {
		g.circles = append(g.circles, NewCircle(i+1, width, height))
	}
	return g
}

func NewCircle(id, maxX, maxY int) Circle {
	o := Circle{
		id:     id,
		size:   12,
		color:  colornames.Azure,
		center: Location{rand.Intn(maxX), rand.Intn(maxY)},
	}
	return o
}

func main() {
	ebiten.SetWindowSize(CFG.Width, CFG.Height)
	ebiten.SetWindowTitle(CFG.Title)
	g := NewGame(CFG.Width, CFG.Height, 100)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}

}
