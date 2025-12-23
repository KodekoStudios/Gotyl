package text

import (
	"strings"

	"github.com/fogleman/gg"
)

// MeasureStringWidthFallback measures the width of a string taking font fallback per-rune into account.
func MeasureStringWidthFallback(dc *gg.Context, s string, fonts []*Font) float64 {
	var w float64
	for _, r := range s {
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
		rw, _ := dc.MeasureString(string(r))
		w += rw
	}
	return w
}

// WrapTextFallback wraps the given text into multiple lines so each line's width
// does not exceed maxWidth. It favors splitting on spaces, and falls back to
// rune-level splitting for very long words or scripts without spaces.
func WrapTextFallback(dc *gg.Context, textStr string, fonts []*Font, maxWidth float64) []string {
	if maxWidth <= 0 {
		return []string{textStr}
	}

	// Use simple word-based wrapping first.
	words := strings.Fields(textStr)
	if len(words) == 0 {
		return []string{""}
	}

	spaceW := MeasureStringWidthFallback(dc, " ", fonts)
	var lines []string
	var current string
	var currentW float64

	appendCurrent := func() {
		if current != "" {
			lines = append(lines, current)
			current = ""
			currentW = 0
		}
	}

	// helper to split long word by runes
	splitLongWord := func(word string) {
		var buf []rune
		var bufW float64
		for _, r := range word {
			rw := MeasureStringWidthFallback(dc, string(r), fonts)
			if bufW+rw > maxWidth && len(buf) > 0 {
				lines = append(lines, string(buf))
				buf = []rune{r}
				bufW = rw
			} else {
				buf = append(buf, r)
				bufW += rw
			}
		}
		if len(buf) > 0 {
			lines = append(lines, string(buf))
		}
	}

	for _, w := range words {
		wW := MeasureStringWidthFallback(dc, w, fonts)
		if current == "" {
			if wW <= maxWidth {
				current = w
				currentW = wW
			} else {
				// too long for a single line: split by runes
				splitLongWord(w)
			}
			continue
		}

		// try to append word with a preceding space
		if currentW+spaceW+wW <= maxWidth {
			current = current + " " + w
			currentW += spaceW + wW
		} else {
			// flush current and start new
			appendCurrent()
			if wW <= maxWidth {
				current = w
				currentW = wW
			} else {
				splitLongWord(w)
			}
		}
	}

	appendCurrent()
	return lines
}
