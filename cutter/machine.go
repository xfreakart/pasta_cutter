package cutter

import (
	"image"
	"github.com/disintegration/imaging"
	"golang.org/x/image/draw"
	"github.com/sirupsen/logrus"
)

const (
	MAXMULTIPLE = 50
)


func CutVertically(inputImage image.Image, debugEnabled bool) (image.Image) {
	imageHeight := inputImage.Bounds().Max.Y
	imageWidth := inputImage.Bounds().Max.X

	cutterSize := findBigestMultiple(imageHeight)
	factor := imageHeight / cutterSize

	outputImage := image.NewRGBA(image.Rect(0, 0, imageWidth*2, imageHeight/2))
	imageLeft := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight/2))
	imageRight := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight/2))

	logrus.Info("Vertical cutter size:", cutterSize)

	for i := 0; i < cutterSize; i++ {
		y0 := factor * i
		x0 := 0

		rectcropimg := imaging.Crop(inputImage, image.Rect(x0, y0, imageWidth, y0+factor))
		var offset image.Point

		if i%2 == 0 {
			offset = image.Pt(0, y0/2)
			draw.Draw(imageLeft, rectcropimg.Bounds().Add(offset), rectcropimg, image.ZP, draw.Over)
		} else {
			offset = image.Pt(0, y0/2-factor/2)
			draw.Draw(imageRight, rectcropimg.Bounds().Add(offset), rectcropimg, image.ZP, draw.Over)
		}
	}

	if debugEnabled {
		SaveImage(imageRight, "vertical_right.png")
		SaveImage(imageLeft, "vertical_left.png")
	}

	draw.Draw(outputImage, imageLeft.Bounds(), imageLeft, image.ZP, draw.Over)
	offset := image.Pt(imageWidth, 0)
	draw.Draw(outputImage, imageRight.Bounds().Add(offset), imageRight, image.ZP, draw.Over)

	return outputImage
}

func CutHorizontaly(inputImage image.Image, debugEnabled bool) (image.Image) {
	imageHeight := inputImage.Bounds().Max.Y
	imageWidth := inputImage.Bounds().Max.X

	cutterSize := findBigestMultiple(imageWidth)
	logrus.Info("Horitzontal cutter size:", cutterSize)

	outputImage := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	imageLeft := image.NewRGBA(image.Rect(0, 0, imageWidth/2, imageHeight))
	imageRight := image.NewRGBA(image.Rect(0, 0, imageWidth/2, imageHeight))

	factor := imageWidth / cutterSize

	for i := 0; i < cutterSize; i++ {
		x0 := factor * i

		rectcropimg := imaging.Crop(inputImage, image.Rect(x0, 0, x0+factor, imageHeight))

		var offset image.Point

		if i%2 == 0 {
			offset = image.Pt(x0/2, 0)
			draw.Draw(imageLeft, rectcropimg.Bounds().Add(offset), rectcropimg, image.ZP, draw.Over)
		} else {
			offset = image.Pt(x0/2-factor/2, 0)
			draw.Draw(imageRight, rectcropimg.Bounds().Add(offset), rectcropimg, image.ZP, draw.Over)
		}
	}

	draw.Draw(outputImage, imageLeft.Bounds(), imageLeft, image.ZP, draw.Over)

	offset := image.Pt(imageWidth/2, 0)
	draw.Draw(outputImage, imageRight.Bounds().Add(offset), imageRight, image.ZP, draw.Over)

	if debugEnabled {
		SaveImage(outputImage, "first_cut.png")
		SaveImage(imageLeft, "horizontal_left.png")
		SaveImage(imageRight, "horizontal_right.png")
	}

	return outputImage
}

func findBigestMultiple(num int) (maxmultiple int) {
	maxmultiple = 1

	for i := maxmultiple; i < MAXMULTIPLE; i ++ {
		if (num%i == 0) {
			if i > maxmultiple {
				maxmultiple = i
			}
		}
	}

	return
}
