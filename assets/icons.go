package assets

import (
	"bytes"
	"embed"
	"image"
	"image/png"
	"sync"
)

//go:embed icons/*.png
var iconsFS embed.FS

var (
	iconCache = map[string]image.Image{}
	iconMu    sync.RWMutex
)

func LoadIcon(path string) image.Image {
	// Fast path
	iconMu.RLock()
	img, ok := iconCache[path]
	iconMu.RUnlock()
	if ok {
		return img
	}

	// Slow path
	data, err := iconsFS.ReadFile(path)
	if err != nil {
		panic(err)
	}

	img, err = png.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	iconMu.Lock()
	iconCache[path] = img
	iconMu.Unlock()

	return img
}
