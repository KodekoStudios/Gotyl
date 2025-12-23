package imaging

import (
	"fmt"
	"image"
	"image/color"
	"sync"

	"gostyl/internal/palette"
)

var (
	scannableBaseCache  = map[string]image.Image{}
	scannableFinalCache = map[string]image.Image{}
	scannableMu         sync.RWMutex
)

func GetSpotifyScannableImage(
	uri string,
	theme palette.Theme,
	size int,
) image.Image {
	// Include theme colors in cache key
	key := fmt.Sprintf(
		"%s:%s:%s:%d",
		uri,
		palette.ColorToHex(theme.Background),
		palette.ColorToHex(theme.Text),
		size,
	)

	scannableMu.RLock()
	if img, ok := scannableFinalCache[key]; ok {
		scannableMu.RUnlock()
		return img
	}
	scannableMu.RUnlock()

	// Request scannable with transparent background and white bars
	url := fmt.Sprintf(
		"https://scannables.scdn.co/uri/plain/png/000000/white/%d/%s",
		size,
		uri,
	)

	base := getScannableBase(url)
	if base == nil {
		return nil
	}

	// Recolor: white bars -> theme.Text, black background -> transparent
	final := recolorScannable(base, theme.Text, theme.Background)

	scannableMu.Lock()
	scannableFinalCache[key] = final
	scannableMu.Unlock()

	return final
}

func getScannableBase(url string) image.Image {
	scannableMu.RLock()
	if img, ok := scannableBaseCache[url]; ok {
		scannableMu.RUnlock()
		return img
	}
	scannableMu.RUnlock()

	img := GetImageFromURL(url)
	if img == nil {
		return nil
	}

	scannableMu.Lock()
	scannableBaseCache[url] = img
	scannableMu.Unlock()

	return img
}

func recolorScannable(
	src image.Image,
	barColor color.Color,
	bgColor color.Color,
) image.Image {
	b := src.Bounds()
	dst := image.NewRGBA(b)

	r0, g0, b0, _ := barColor.RGBA()
	bars := color.RGBA{
		R: uint8(r0 >> 8),
		G: uint8(g0 >> 8),
		B: uint8(b0 >> 8),
		A: 255,
	}

	r1, g1, b1, _ := bgColor.RGBA()
	bg := color.RGBA{
		R: uint8(r1 >> 8),
		G: uint8(g1 >> 8),
		B: uint8(b1 >> 8),
		A: 255,
	}

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pr, pg, pb, _ := src.At(x, y).RGBA()

			// White pixels (bars) -> theme text color
			if pr > 60000 && pg > 60000 && pb > 60000 {
				dst.Set(x, y, bars)
			} else {
				// Everything else -> theme background
				dst.Set(x, y, bg)
			}
		}
	}

	return dst
}
