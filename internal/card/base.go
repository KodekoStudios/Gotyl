package card

import "gostyl/internal/palette"

type Card struct {
	Theme palette.Theme

	// Spotify track URI, for scannable code
	URI string

	Title       string
	Artist      string
	Duration    string
	ReleaseDate string
	Label       string

	Lyrics []string

	CoverImage    string
	ShowScannable bool
}

// Canvas dimensions
const (
	Width  = 2280
	Height = 3480
)

// Layout positions
const (
	CoverX    = 120
	CoverY    = 120
	CoverSize = 2040

	IconsX = 1610
	IconsY = 2530

	ScannableX    = 90
	ScannableY    = 3220
	ScannableSize = 660

	PaletteX            = 120.0
	PaletteY            = 2240.0
	PaletteSwatchWidth  = 340.0
	PaletteSwatchHeight = 80.0
	PaletteCount        = 6

	FooterY      = 3440.0
	FooterHeight = 40.0

	InfoX      = 120.0
	InfoTitleY = 2550.0
	InfoGap    = 150.0

	InfoLabelDurationGap = 80.0

	InfoLabelX = 2160.0
	InfoLabelY = 3260

	InfoDurationX = 2160.0
	InfoDurationY = 2550.0

	LyricsX        = 120.0
	LyricsY        = 2860.0
	LyricsMaxWidth = 2000.0
	LyricsMaxLines = 4
	LyricsLineGap  = 100.0
)

// Font sizes
const (
	FontSizeTitle    = 160.0
	FontSizeArtist   = 120.0
	FontSizeArtistJP = 80.0
	FontSizeDuration = 90.0
	FontSizeLabel    = 60.0
	FontSizeLyrics   = 64.0

	MinFontSizeTitle  = 120.0
	MinFontSizeArtist = 100.0
)

// Font paths
const (
	FontLatin    = "fonts/JetBrainsMono-Regular.ttf"
	FontJapanese = "fonts/SarasaGothicJ-Regular.ttf"

	FontBoldLatin   = "fonts/JetBrainsMono-Bold.ttf"
	FontMediumLatin = "fonts/JetBrainsMono-Medium.ttf"
)
