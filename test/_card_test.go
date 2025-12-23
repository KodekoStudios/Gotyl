package _card_test

import (
	"gostyl/core/palette"
	"image/color"
	"testing"

	"github.com/fogleman/gg"
)

func TestDrawBase(t *testing.T) {
	theme := palette.Dark

	dc := gg.NewContext(2280, 3480)
	dc.SetColor(theme.Background)
	dc.Clear()

	// Placeholder for cover image
	dc.DrawRectangle(140, 150, 2000, 2000)
	dc.SetColor(color.White)
	dc.Fill()

	// Placeholder for color palette swatches
	dc.DrawRectangle(140, 2250, 2000, 140)
	dc.SetColor(color.White)
	dc.Fill()

	// Placeholder for dominant color
	dc.DrawRectangle(0, 3400, 2280, 40)
	dc.SetColor(color.White)
	dc.Fill()

	err := dc.SavePNG("../../assets/base.png")
	if err != nil {
		t.Fatal(err)
	}
}
