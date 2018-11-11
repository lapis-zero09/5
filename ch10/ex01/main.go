package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	format := flag.String("f", "jpeg", "output image format")
	flag.Parse()

	if err := fmtconv(*format, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func fmtconv(format string, in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		return fmt.Errorf("Input format = ", kind)
	}

	convmap := map[string]func(image.Image, io.Writer) error{
		"gif":  toGIF,
		"GIF":  toGIF,
		"jpg":  toJPEG,
		"JPG":  toJPEG,
		"jpeg": toJPEG,
		"JPEG": toJPEG,
		"png":  toPNG,
		"PNG":  toPNG,
	}
	encode, ok := convmap[format]
	if !ok {
		return fmt.Errorf("unsupported format: %v", format)
	}

	if err = encode(img, out); err != nil {
		return err
	}

	return nil
}

func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, nil)
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}
