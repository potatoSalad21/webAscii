package ascii

import (
	"bytes"
	"image"
	"image/color"
	"image/png"

	"github.com/nfnt/resize"
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

func Convert(data []byte) (string, error) {
	src, err := loadImg(data)
	if err != nil {
		return "", err
	}
	resizedImg := resize.Resize(64, 48, src, resize.Lanczos2)

	var frame string
	const ramp = "@#W$9a+=. "
	//const ramp = " .=+a9$W#@"

	max := resizedImg.Bounds().Max
	for y := 0; y < max.Y; y++ {
		for x := 0; x < max.X; x++ {
			avg := grayscale(resizedImg.At(x, y))
			index := len(ramp) * avg / 65536
			frame += string(ramp[index])
		}
		frame += "\n"
	}

	return frame, nil
}
