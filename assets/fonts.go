package assets

import (
	"embed"
	"sync"
)

//go:embed fonts/*.ttf
var fontsFS embed.FS

var (
	fontBytesCache = map[string][]byte{}
	fontBytesMu    sync.RWMutex
)

func LoadFontBytes(path string) []byte {
	fontBytesMu.RLock()
	if b, ok := fontBytesCache[path]; ok {
		fontBytesMu.RUnlock()
		return b
	}
	fontBytesMu.RUnlock()

	data, err := fontsFS.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fontBytesMu.Lock()
	fontBytesCache[path] = data
	fontBytesMu.Unlock()

	return data
}
