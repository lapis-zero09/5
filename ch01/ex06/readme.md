
## colorの選び方

定義域：0 < t < 10pi

```
x := math.Sin(t)
colorIdx := uint8(math.Exp(math.Abs(x) * 2))
```

f(t) = exp(2|Sin(t)|)とすると

値域：1 < f(x) < 8


よって，colorIdxの値域は，

1 <= colorIdx <= 7

paletteの1~7は七色の各色を指す．

```
var palette = []color.Color{
	color.RGBA{192, 192, 192, 0}, // 背景色gray
	color.RGBA{255, 0, 0, 0}, // red
	color.RGBA{255, 69, 0, 0}, // orange
	color.RGBA{255, 255, 0, 0}, // yellow
	color.RGBA{0, 128, 0, 0}, // green
	color.RGBA{0, 0, 255, 0}, // blue
	color.RGBA{75, 0, 130, 0}, // purple
	color.RGBA{238, 130, 238, 0}, // pink
}
```
