package text

import (
	"sync"

	"gostyl/assets"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

type Font struct {
	Face font.Face
	SFNT *sfnt.Font
}

var (
	sfntCache = map[string]*sfnt.Font{}
	sfntMu    sync.RWMutex
)

func LoadFont(path string, size float64) (*Font, error) {
	// sfnt cache (font parsing)
	sfntMu.RLock()
	sf, ok := sfntCache[path]
	sfntMu.RUnlock()

	if !ok {
		data := assets.LoadFontBytes(path)

		var err error
		sf, err = sfnt.Parse(data)
		if err != nil {
			return nil, err
		}

		sfntMu.Lock()
		sfntCache[path] = sf
		sfntMu.Unlock()
	}

	face, err := opentype.NewFace(sf, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}

	return &Font{
		Face: face,
		SFNT: sf,
	}, nil
}
