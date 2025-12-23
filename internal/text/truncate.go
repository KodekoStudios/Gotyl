package text

import (
	"strings"

	"github.com/fogleman/gg"
)

// TruncateTextFallback truncates the text to fit within maxWidth, appending "..." if truncated.
// It uses the provided fonts for measurement.
func TruncateTextFallback(dc *gg.Context, text string, fonts []*Font, maxWidth float64) string {
	if maxWidth <= 0 {
		return ""
	}

	w := MeasureStringWidthFallback(dc, text, fonts)
	if w <= maxWidth {
		return text
	}

	ellipsis := "..."
	ellipsisW := MeasureStringWidthFallback(dc, ellipsis, fonts)

	// If even ellipsis doesn't fit, return empty (or just ellipsis if it fits barely, but simpler to return empty)
	if ellipsisW > maxWidth {
		return ""
	}

	targetWidth := maxWidth - ellipsisW

	// Iteratively remove rune from end
	// This is O(N^2) in worst case because of MeasureStringWidthFallback, but N is small (title length)
	// Binary search would be better for very long strings.
	runes := []rune(text)
	for len(runes) > 0 {
		runes = runes[:len(runes)-1]
		current := string(runes)
		if MeasureStringWidthFallback(dc, current, fonts) <= targetWidth {
			return strings.TrimSpace(current) + ellipsis
		}
	}

	return ellipsis
}
