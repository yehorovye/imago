package imago

// rgb color struct.
type Color struct {
	R, G, B uint8
}

// image struct, with width, height and pixels.
type Image struct {
	Width, Height int
	Pixels        [][]Color // x, y
}

// creates a new `Image` instance.
func New(width, height int) *Image {
	pixels := make([][]Color, height)
	for y := range pixels {
		pixels[y] = make([]Color, width)
	}
	return &Image{Width: width, Height: height, Pixels: pixels}
}

// returns specified coordinates pixel color.
func (img *Image) At(x, y int) Color {
	if x < 0 || y < 0 || x >= img.Width || y >= img.Height {
		return Color{}
	}
	return img.Pixels[y][x]
}

// sets the pixel color at specified coordinates.
func (img *Image) Set(x, y int, c Color) {
	if x < 0 || y < 0 || x >= img.Width || y >= img.Height {
		return
	}
	img.Pixels[y][x] = c
}

// fills the image with the specified color.
func (img *Image) Fill(c Color) {
	for y := range img.Pixels {
		for x := range img.Pixels[y] {
			img.Pixels[y][x] = c
		}
	}
}

// inverts the image.
func (img *Image) Invert() {
	for y := range img.Pixels {
		for x := range img.Pixels[y] {
			p := img.Pixels[y][x]
			img.Pixels[y][x] = Color{R: 255 - p.R, G: 255 - p.G, B: 255 - p.B}
		}
	}
}

// turns the image gray. *sad*.
func (img *Image) Grayscale() {
	for y := range img.Pixels {
		for x := range img.Pixels[y] {
			p := img.Pixels[y][x]
			avg := uint8((uint16(p.R) + uint16(p.G) + uint16(p.B)) / 3)
			img.Pixels[y][x] = Color{R: avg, G: avg, B: avg}
		}
	}
}
