package imaging

import (
	"image"
	"sync"

	"github.com/fogleman/gg"
)

const (
	canvasW = 2280
	canvasH = 3480
)

var ctxPool = sync.Pool{
	New: func() any {
		img := image.NewRGBA(image.Rect(0, 0, canvasW, canvasH))
		return gg.NewContextForRGBA(img)
	},
}

func GetContext() *gg.Context {
	dc := ctxPool.Get().(*gg.Context)

	// Reset state
	dc.SetRGBA(0, 0, 0, 0)
	dc.Clear()
	dc.Identity()

	return dc
}

func PutContext(dc *gg.Context) {
	ctxPool.Put(dc)
}
