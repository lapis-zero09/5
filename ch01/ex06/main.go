package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	// color.RGBA{0, 0xff, 0, 0xff}, // green
	color.RGBA{255, 0, 0, 1},
	color.RGBA{255, 69, 0, 1},
	color.RGBA{255, 255, 0, 1},
	color.RGBA{0, 128, 0, 1},
	color.RGBA{0, 0, 255, 1},
	color.RGBA{75, 0, 130, 1},
	color.RGBA{238, 130, 238, 1},
}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			colorIdx := uint8(math.Exp(math.Abs(math.Sin(t)) * 2))
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIdx)
		}
		// fmt.Println(z)
		// break
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
