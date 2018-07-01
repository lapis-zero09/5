package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{
	color.RGBA{192, 192, 192, 0},
	color.RGBA{255, 0, 0, 0},
	color.RGBA{255, 69, 0, 0},
	color.RGBA{255, 255, 0, 0},
	color.RGBA{0, 128, 0, 0},
	color.RGBA{0, 0, 255, 0},
	color.RGBA{75, 0, 130, 0},
	color.RGBA{238, 130, 238, 0},
}

func Lissajous(out io.Writer, params map[string]float64) {
	cycles := int(params["cycles"])
	size := int(params["size"])
	nframes := int(params["nframes"])
	delay := int(params["delay"])
	res := params["res"]
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2.*math.Pi; t += res {
			x := math.Sin(t)
			colorIdx := uint8(math.Exp(math.Abs(x) * 2))
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5),
				colorIdx)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
