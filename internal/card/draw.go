package card

import (
	"log"
	"time"

	"github.com/fogleman/gg"
)

func Draw(dc *gg.Context, c *Card) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("Draw took %s", elapsed)
	}()

	theme := c.Theme
	dc.SetColor(theme.Background)
	dc.Clear()

	img := DrawCover(dc, c.CoverImage)
	DrawIcons(dc, theme)
	if c.ShowScannable {
		DrawScannable(dc, c.URI, theme)
	}
	DrawPalette(dc, img)
	DrawInfo(dc, c.Title, c.Artist, c.Duration, c.ReleaseDate, c.Label, theme)
	DrawLyrics(dc, c.Lyrics, theme)
}
