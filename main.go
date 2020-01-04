package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func main() {

	imageRead := gocv.IMRead("hotdog1.jpg", gocv.IMReadAnyColor) // fullscreen image
	xmlFile := "haarcascade_hotdog1.xml"

	// open display window
	// window := gocv.NewWindow("Hot Dog Detect")
	// defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize hot dogs
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	imageRead.CopyTo(&img)
	if img.Empty() {
		return
	}

	// detect faces
	rects := classifier.DetectMultiScale(img)
	fmt.Printf("found %d hot dogs\n", len(rects))

	// draw a rectangle around each hot dog on the original image,
	// along with text identifying as "Hot Dog"
	for _, r := range rects {
		gocv.Rectangle(&img, r, blue, 2)

		size := gocv.GetTextSize("Hot Dog", gocv.FontHersheyPlain, 1.2, 2)
		pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
		gocv.PutText(&img, "Hot Dog", pt, gocv.FontHersheyPlain, 1.0, blue, 1)
	}

	// write the image to disk
	gocv.IMWrite("OutputImage.jpg", img)

}
