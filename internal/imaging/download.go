package imaging

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

var (
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	imageCache   = make(map[string]image.Image)
	imageCacheMu sync.RWMutex

	resizeCache   = make(map[string]image.Image)
	resizeCacheMu sync.RWMutex
)

// DrawResizedImage draws an image resized to (w, h) using a resize cache.
func DrawResizedImage(dc *gg.Context, img image.Image, x, y, w, h int) {
	key := fmt.Sprintf("%p:%dx%d", img, w, h)

	resizeCacheMu.RLock()
	if cached, ok := resizeCache[key]; ok {
		resizeCacheMu.RUnlock()
		dc.DrawImage(cached, x, y)
		return
	}
	resizeCacheMu.RUnlock()

	resized := imaging.Resize(img, w, h, imaging.Box)

	resizeCacheMu.Lock()
	resizeCache[key] = resized
	resizeCacheMu.Unlock()

	dc.DrawImage(resized, x, y)
}

// GetImageFromURL downloads and decodes an image from a URL.
// The result is cached in memory.
func GetImageFromURL(url string) image.Image {
	imageCacheMu.RLock()
	if img, ok := imageCache[url]; ok {
		imageCacheMu.RUnlock()
		return img
	}
	imageCacheMu.RUnlock()

	resp, err := httpClient.Get(url)
	if err != nil {
		log.Println("http error:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("bad status:", resp.Status)
		return nil
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		log.Println("read error:", err)
		return nil
	}

	img, _, err := image.Decode(bytes.NewReader(buf.Bytes()))
	if err != nil {
		log.Println("decode error:", err)
		return nil
	}

	imageCacheMu.Lock()
	imageCache[url] = img
	imageCacheMu.Unlock()

	return img
}
