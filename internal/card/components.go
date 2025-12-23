package card

import (
	"gostyl/assets"
	"gostyl/internal/imaging"
	"gostyl/internal/palette"
	"gostyl/internal/text"
	"image"

	"github.com/fogleman/gg"
)

// DrawCover downloads and draws the cover image at the default position.
// Returns the downloaded image for further processing (e.g. palette extraction).
func DrawCover(dc *gg.Context, url string) image.Image {
	img := imaging.GetImageFromURL(url)
	if img == nil {
		return nil
	}
	imaging.DrawResizedImage(dc, img, CoverX, CoverY, CoverSize, CoverSize)
	return img
}

// DrawIcons draws the theme icons at the default position.
func DrawIcons(dc *gg.Context, theme palette.Theme) {
	icons := assets.LoadIcon(theme.Icons)
	dc.DrawImage(icons, IconsX, IconsY)
}

// DrawScannable draws the Spotify scannable code at the default position.
func DrawScannable(dc *gg.Context, uri string, theme palette.Theme) {
	scannable := imaging.GetSpotifyScannableImage(uri, theme, ScannableSize)
	if scannable != nil {
		dc.DrawImage(scannable, ScannableX, ScannableY)
	}
}

// DrawPalette extracts and draws color palette swatches and a footer line.
func DrawPalette(dc *gg.Context, img image.Image) {
	if img == nil {
		return
	}

	paletteColors := palette.FromImage(img, PaletteCount)
	dominantColor := palette.DominantColor(img)

	// Draw swatches
	for i, col := range paletteColors {
		dc.SetColor(col)
		dc.DrawRectangle(
			PaletteX+float64(i)*PaletteSwatchWidth,
			PaletteY,
			PaletteSwatchWidth,
			PaletteSwatchHeight,
		)
		dc.Fill()
	}

	// Draw footer line
	dc.DrawRectangle(0, FooterY, float64(dc.Width()), FooterHeight)
	dc.SetColor(dominantColor)
	dc.Fill()
}

// DrawInfo draws the Title and Artist at the default position.
func DrawInfo(dc *gg.Context, title, artist, duration, releaseDate, label string, theme palette.Theme) {
	dc.SetColor(theme.Text)

	// Title
	// Max width for title: IconsX - InfoX - padding (e.g. 50)
	maxTitleWidth := IconsX - InfoX - 50.0

	titleFonts, fittedTitle := text.FitText(
		dc,
		title,
		[]string{FontBoldLatin, FontJapanese},
		FontSizeTitle,
		MinFontSizeTitle,
		5.0, // step
		maxTitleWidth,
	)

	text.DrawStringFallback(dc, fittedTitle, InfoX, InfoTitleY, titleFonts)

	// Artist
	// Max width for artist
	maxArtistWidth := float64(Width) - InfoX - InfoX

	artistFonts, fittedArtist := text.FitText(
		dc,
		artist,
		[]string{FontMediumLatin, FontJapanese},
		FontSizeArtist,
		MinFontSizeArtist,
		5.0, // step
		maxArtistWidth,
	)

	text.DrawStringFallback(dc, fittedArtist, InfoX, InfoTitleY+InfoGap, artistFonts)

	// Duration (right-aligned)
	latinDuration, _ := text.LoadFont(FontMediumLatin, FontSizeDuration)
	jpDuration, _ := text.LoadFont(FontJapanese, FontSizeDuration)

	text.DrawStringFallbackRTL(dc, duration, InfoDurationX, InfoDurationY, []*text.Font{latinDuration, jpDuration})

	// Label & ReleaseDate (right-aligned)
	latinLabel, _ := text.LoadFont(FontLatin, FontSizeLabel)
	jpLabel, _ := text.LoadFont(FontJapanese, FontSizeLabel)

	text.DrawStringFallbackRTL(dc, releaseDate, InfoLabelX, InfoLabelY, []*text.Font{latinLabel, jpLabel})
	text.DrawStringFallbackRTL(dc, label, InfoLabelX, InfoLabelY+InfoLabelDurationGap, []*text.Font{latinLabel, jpLabel})
}

// DrawLyrics draws the lyrics with wrapping at the default position.
func DrawLyrics(dc *gg.Context, lyrics []string, theme palette.Theme) {
	latin, _ := text.LoadFont(FontLatin, FontSizeLyrics)
	jp, _ := text.LoadFont(FontJapanese, FontSizeLyrics)
	fonts := []*text.Font{latin, jp}

	dc.SetColor(theme.Text)

	y := LyricsY
	linesDrawn := 0

	for _, l := range lyrics {
		if linesDrawn >= LyricsMaxLines {
			break
		}

		wrapped := text.WrapTextFallback(dc, l, fonts, LyricsMaxWidth)
		for _, wl := range wrapped {
			if linesDrawn >= LyricsMaxLines {
				break
			}
			text.DrawStringFallback(dc, wl, LyricsX, y, fonts)
			y += LyricsLineGap
			linesDrawn++
		}
	}
}
