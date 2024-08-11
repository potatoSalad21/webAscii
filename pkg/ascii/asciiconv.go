package ascii

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
)

func loadImg(data []byte) (*image.NRGBA, error) {
	imgData := bytes.NewReader(data)
	img, err := png.Decode(imgData)
	if err != nil {
		return nil, err
	}

	return img.(*image.NRGBA), nil
}

func Convert(data []byte) {
	img, err := loadImg(data)
	if err != nil {
		log.Println("Cannot decode file:", err)
		return
	}

	// TODO: scale the image appropirately and send back
	ramp := "@#&0O[(;:,.'`  "
	bounds := img.Bounds()
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			grayscale := int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))

			index := len(ramp) * grayscale / 65536
			fmt.Printf(string(ramp[index]))
		}
		fmt.Println()
	}
}
