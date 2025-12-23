package text

import (
	"github.com/fogleman/gg"
)

// FitText tries to fit the text within maxWidth by reducing the font size step by step.
// It returns the loaded fonts at the best size (or minSize) and the text (truncated if necessary).
func FitText(
	dc *gg.Context,
	text string,
	fontPaths []string,
	initialSize, minSize, step, maxWidth float64,
) ([]*Font, string) {
	currentSize := initialSize

	for currentSize >= minSize {
		// Load fonts at current size
		var fonts []*Font
		for _, p := range fontPaths {
			f, err := LoadFont(p, currentSize)
			if err == nil {
				fonts = append(fonts, f)
			}
		}

		if len(fonts) == 0 {
			// Should not happen if paths are valid
			return nil, text
		}

		// Check if it fits
		if MeasureStringWidthFallback(dc, text, fonts) <= maxWidth {
			return fonts, text
		}

		// If not, reduce size
		currentSize -= step
	}

	// If we reached here, it didn't fit even at minSize (or close to it).
	// Use minSize and truncate.
	finalSize := minSize
	var fonts []*Font
	for _, p := range fontPaths {
		f, err := LoadFont(p, finalSize)
		if err == nil {
			fonts = append(fonts, f)
		}
	}

	truncated := TruncateTextFallback(dc, text, fonts, maxWidth)
	return fonts, truncated
}
