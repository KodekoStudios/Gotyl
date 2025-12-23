package text

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font/sfnt"
)

func supportsRune(f *sfnt.Font, r rune) bool {
	var buf sfnt.Buffer
	glyph, err := f.GlyphIndex(&buf, r)
	if err != nil {
		return false
	}
	return glyph != 0
}

func DrawStringFallback(
	dc *gg.Context,
	text string,
	x, y float64,
	fonts []*Font,
) {
	cursorX := x

	for _, r := range text {
		var used *Font
		for _, f := range fonts {
			if supportsRune(f.SFNT, r) {
				used = f
				break
			}
		}

		if used == nil {
			continue
		}

		dc.SetFontFace(used.Face)
		s := string(r)

		w, _ := dc.MeasureString(s)
		dc.DrawString(s, cursorX, y)
		cursorX += w
	}
}

// DrawStringFallbackRTL draws text right-aligned (text ends at x).
func DrawStringFallbackRTL(
	dc *gg.Context,
	text string,
	rightX, y float64,
	fonts []*Font,
) {
	// First measure total width
	width := MeasureStringWidthFallback(dc, text, fonts)
	// Then draw starting from (rightX - width)
	DrawStringFallback(dc, text, rightX-width, y, fonts)
}
