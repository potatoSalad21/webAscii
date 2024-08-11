package ascii

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
)

func loadImg(data []byte) (*image.NRGBA, error) {
	imgData := bytes.NewReader(data)
	img, err := png.Decode(imgData)
	if err != nil {
		return nil, err
	}

	return img.(*image.NRGBA), nil
}

func grayscale(c color.Color) int {
	r, g, b, _ := c.RGBA()
	return int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

func avgPixel(img image.Image, x, y, w, h int) int {
	var (
		count int
		sum   int
		max   = img.Bounds().Max
	)

	for i := x; i < x+w && i < max.X; i++ {
		for j := y; y < y+h && j < max.Y; j++ {
			sum += grayscale(img.At(i, j))
			count++
		}
	}

	return sum / count
}

func Convert(data []byte) (string, error) {
	img, err := loadImg(data)
	if err != nil {
		return "", err
	}

	var frame string
	const (
		ramp   = "@#W$9a+=. "
		scaleX = 5
		scaleY = 10
	)

	max := img.Bounds().Max
	for y := 0; y < max.Y; y += scaleY {
		for x := 0; x < max.X; x += scaleX {
			avg := avgPixel(img, x, y, scaleX, scaleY)
			index := len(ramp) * avg / 65536
			frame += string(ramp[index])
		}
		frame += "\n"
	}

	return frame, nil
}
