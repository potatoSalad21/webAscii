package ascii

import (
    "os"
    "log"
    "image"
)

func loadImg(filePath string) *image.NRGBA {
    // TODO: instead of os IO, decode image from base64
    imgFile, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Cannot open file:", err)
    }
    defer imgFile.Close()

    img, _, err := image.Decode(imgFile)
    if err != nil {
        log.Fatal("Cannot decode file:", err)
    }

    return img.(*image.NRGBA)
}

func grayscale(img *image.NRGBA) {

}

func convert(string) {
    // load(), grayscale()

    // iterate over grayscaled img
    // convert to ascii

}
