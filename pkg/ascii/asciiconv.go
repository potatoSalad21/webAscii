package ascii

import (
	"bytes"
	"image"
	"image/png"
	"log"
)

func loadImg(data []byte) *image.NRGBA {
	imgData := bytes.NewReader(data)
	img, err := png.Decode(imgData)
	if err != nil {
		log.Fatal("Cannot decode file:", err)
	}
	log.Println("Succesfully loaded to png")

	return img.(*image.NRGBA)
}

func grayscale(img *image.NRGBA) {

}

func Convert(data []byte) {
	loadImg(data)
	// load(), grayscale()

	// iterate over grayscaled img
	// convert to ascii

}
