package cutter

import (
	"image"
	"os"
	"image/png"
	"github.com/sirupsen/logrus"
)

const (
	OUTPUTFOLDER = "output_images/"
)

func SaveImage(img image.Image, name string) (err error) {
	outputFile, err := os.Create(OUTPUTFOLDER + name)
	png.Encode(outputFile, img)
	outputFile.Close()
	logrus.Info("Saving file:", OUTPUTFOLDER+name)
	return err

}
