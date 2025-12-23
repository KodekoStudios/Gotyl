package card

import (
	"bytes"
	"image/jpeg"
	"image/png"

	"gostyl/internal/imaging"
)

// RenderPNG renders the card as PNG with fast compression.
func RenderPNG(c *Card) ([]byte, error) {
	dc := imaging.GetContext()
	defer imaging.PutContext(dc)

	Draw(dc, c)

	var buf bytes.Buffer
	encoder := png.Encoder{CompressionLevel: png.NoCompression}
	if err := encoder.Encode(&buf, dc.Image()); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// RenderJPEG renders the card as JPEG (much faster than PNG).
func RenderJPEG(c *Card, quality int) ([]byte, error) {
	dc := imaging.GetContext()
	defer imaging.PutContext(dc)

	Draw(dc, c)

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, dc.Image(), &jpeg.Options{Quality: quality}); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
