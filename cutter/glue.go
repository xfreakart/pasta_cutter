package cutter

import (
	"image"
	"os"
	"image/png"
)

const(
	OUTPUTFOLDER = "output_images/"
)




func SaveImage(img image.Image, name string) (err error) {
	outputFile, err := os.Create(OUTPUTFOLDER+name)
	png.Encode(outputFile, img)
	outputFile.Close()

	return err
}
