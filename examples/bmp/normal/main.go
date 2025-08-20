package main

import (
	"fmt"
	"imago/encodings/bmp"
	"os"
)

func main() {
	// read the input BMP file
	data, err := os.ReadFile("examples/inputs/500x500.bmp")
	if err != nil {
		panic(err)
	}

	// decode the BMP image from bytes
	img, err := bmp.Load(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("loaded BMP:", img.Width, "x", img.Height)

	// encode the image back to BMP format
	out, err := bmp.Save(img)
	if err != nil {
		panic(err)
	}

	// write the encoded BMP to a new file
	err = os.WriteFile("examples/outputs/500x500.bmp", out, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("ok :)")
}
