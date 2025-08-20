package main

import (
	"log"
	"os"

	"imago"
	"imago/encodings/bmp"
)

func main() {
	// create a new 256x256 image
	img := imago.New(256, 256)

	// fill the image with red
	red := imago.Color{R: 255, G: 0, B: 0}
	img.Fill(red)

	// invert the colors (red becomes cyan)
	img.Invert()

	// convert the image to grayscale
	img.Grayscale()

	// save the image as a BMP file
	out, err := bmp.Save(img)
	if err != nil {
		log.Fatal("SaveBMP failed:", err)
	}

	// write the BMP data to a file
	err = os.WriteFile("examples/outputs/custom.bmp", out, 0644)
	if err != nil {
		log.Fatal("WriteFile failed:", err)
	}

	log.Println("ok :)")
}
