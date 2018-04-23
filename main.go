package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/sirupsen/logrus"
	"image/jpeg"
	"github.com/xfreakart/pasta_cutter/cutter"
)

var firstCutterSize = 0

func main() {
	imagePath := flag.String("file", "", "File name path")
	debugImages := flag.Bool("debug", false, "generate debug images")
	flag.Parse()

	if *imagePath == "" {
		logrus.Error("File name should not be empty")
		os.Exit(0)
	}
	fmt.Println("Using file:", *imagePath)

	fileImage, err := os.Open(*imagePath)
	if err != nil {
		logrus.Panicf("%v\n", err)
	}

	inputImage, err := jpeg.Decode(fileImage)
	if err != nil {
		logrus.Panicf("%v\n", err)
	}

	firstStepImage := cutter.CutHorizontaly(inputImage, *debugImages)
	lastStepImage := cutter.CutVertically(firstStepImage, *debugImages)

	cutter.SaveImage(lastStepImage, "output_final.png")
}
