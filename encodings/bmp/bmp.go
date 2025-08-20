package bmp

import (
	"bytes"
	"encoding/binary"
	"errors"

	"imago"
)

func Load(data []byte) (*imago.Image, error) {
	if len(data) < 54 || string(data[:2]) != "BM" {
		return nil, errors.New("invalid BMP header")
	}

	pxOffset := int(binary.LittleEndian.Uint32(data[10:14]))

	w := int(int32(binary.LittleEndian.Uint32(data[18:22])))
	h := int(int32(binary.LittleEndian.Uint32(data[22:26])))
	planes := binary.LittleEndian.Uint16(data[26:28])
	bpp := binary.LittleEndian.Uint16(data[28:30])
	compression := binary.LittleEndian.Uint32(data[30:34])

	if planes != 1 || bpp != 24 || compression != 0 {
		return nil, errors.New("only uncompressed 24-bit BMP supported")
	}

	row := ((bpp*uint16(w) + 31) / 32) * 4
	pixels := make([][]imago.Color, h)

	for y := range pixels {
		pixels[y] = make([]imago.Color, w)
		off := pxOffset + (h-1-y)*int(row)
		for x := range w {
			i := off + x*3
			if i+2 >= len(data) {
				return nil, errors.New("unexpected end of pixel data")
			}
			pixels[y][x] = imago.Color{R: data[i+2], G: data[i+1], B: data[i]}
		}
	}

	return &imago.Image{Width: w, Height: h, Pixels: pixels}, nil
}

func Save(img *imago.Image) ([]byte, error) {
	w, h := img.Width, img.Height
	row := ((24*w + 31) / 32) * 4
	dataSize := row * h
	size := 14 + 40 + dataSize

	buf := &bytes.Buffer{}
	buf.WriteString("BM")
	binary.Write(buf, binary.LittleEndian, uint32(size))
	binary.Write(buf, binary.LittleEndian, uint16(0))
	binary.Write(buf, binary.LittleEndian, uint16(0))
	binary.Write(buf, binary.LittleEndian, uint32(54))

	binary.Write(buf, binary.LittleEndian, uint32(40))
	binary.Write(buf, binary.LittleEndian, int32(w))
	binary.Write(buf, binary.LittleEndian, int32(h))
	binary.Write(buf, binary.LittleEndian, uint16(1))
	binary.Write(buf, binary.LittleEndian, uint16(24))
	binary.Write(buf, binary.LittleEndian, uint32(0))
	binary.Write(buf, binary.LittleEndian, uint32(dataSize))
	binary.Write(buf, binary.LittleEndian, int32(0))
	binary.Write(buf, binary.LittleEndian, int32(0))
	binary.Write(buf, binary.LittleEndian, uint32(0))
	binary.Write(buf, binary.LittleEndian, uint32(0))

	padding := make([]byte, row-w*3)

	for y := h - 1; y >= 0; y-- {
		for x := range w {
			c := img.Pixels[y][x]
			buf.WriteByte(c.B)
			buf.WriteByte(c.G)
			buf.WriteByte(c.R)
		}
		buf.Write(padding)
	}

	return buf.Bytes(), nil
}
