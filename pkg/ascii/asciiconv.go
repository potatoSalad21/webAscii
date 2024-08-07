package ascii

import (
	"bytes"
	"image"
	"image/png"
	"log"
)

func loadImg(data []byte) (*image.NRGBA, error) {
	imgData := bytes.NewReader(data)
	img, err := png.Decode(imgData) // ERROR: creates an empty png
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

	density := []string{"`", "'", "\"", ".", "~", "-", "=", "*", ":", ";", "/", "#"}
	bounds := img.Bounds()
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(y, x).RGBA()
			average := (float64(r) + float64(g) + float64(b)) / 3.0
			index := average * float64(len(density)-1) / 255
			log.Println("INDEX", index)
			log.Println("RGB:", r, g, b)
			//fmt.Printf(density[int(index)])

			// TODO: map grayscale to character
		}
		//fmt.Printf("\n")
	}
}
