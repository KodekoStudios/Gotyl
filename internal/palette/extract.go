package palette

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type rgb struct {
	r, g, b float64
}

type cluster struct {
	center rgb
	count  int
}

func FromImage(img image.Image, k int) []color.Color {
	pixels := samplePixels(img, 5)
	clusters := kmeans(pixels, k, 15)

	sortClustersByCount(clusters)

	palette := make([]color.Color, 0, k)
	for _, c := range clusters {
		palette = append(palette, color.RGBA{
			R: uint8(c.center.r),
			G: uint8(c.center.g),
			B: uint8(c.center.b),
			A: 255,
		})
	}

	return palette
}

func DominantColor(img image.Image) color.Color {
	bounds := img.Bounds()
	hist := make(map[uint32]int)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x += 2 {
			r, g, b, _ := img.At(x, y).RGBA()

			r5 := uint32(r>>8) >> 3
			g5 := uint32(g>>8) >> 3
			b5 := uint32(b>>8) >> 3

			key := (r5 << 10) | (g5 << 5) | b5
			hist[key]++
		}
	}

	var maxKey uint32
	maxCount := 0
	for k, v := range hist {
		if v > maxCount {
			maxCount = v
			maxKey = k
		}
	}

	r := uint8(((maxKey >> 10) & 31) << 3)
	g := uint8(((maxKey >> 5) & 31) << 3)
	b := uint8((maxKey & 31) << 3)

	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func kmeans(points []rgb, k, iterations int) []cluster {
	rand.Seed(time.Now().UnixNano())

	clusters := make([]cluster, k)
	for i := 0; i < k; i++ {
		p := points[rand.Intn(len(points))]
		clusters[i].center = p
	}

	for it := 0; it < iterations; it++ {
		for i := range clusters {
			clusters[i].count = 0
		}

		sums := make([]rgb, k)

		for _, p := range points {
			idx := nearestCluster(p, clusters)
			sums[idx].r += p.r
			sums[idx].g += p.g
			sums[idx].b += p.b
			clusters[idx].count++
		}

		for i := range clusters {
			if clusters[i].count == 0 {
				continue
			}
			clusters[i].center = rgb{
				r: sums[i].r / float64(clusters[i].count),
				g: sums[i].g / float64(clusters[i].count),
				b: sums[i].b / float64(clusters[i].count),
			}
		}
	}

	return clusters
}

func nearestCluster(p rgb, clusters []cluster) int {
	minDist := math.MaxFloat64
	index := 0

	for i, c := range clusters {
		d := colorDistance(p, c.center)
		if d < minDist {
			minDist = d
			index = i
		}
	}

	return index
}

func colorDistance(a, b rgb) float64 {
	return (a.r-b.r)*(a.r-b.r) +
		(a.g-b.g)*(a.g-b.g) +
		(a.b-b.b)*(a.b-b.b)
}

func sortClustersByCount(clusters []cluster) {
	for i := 0; i < len(clusters)-1; i++ {
		for j := i + 1; j < len(clusters); j++ {
			if clusters[j].count > clusters[i].count {
				clusters[i], clusters[j] = clusters[j], clusters[i]
			}
		}
	}
}

func samplePixels(img image.Image, step int) []rgb {
	b := img.Bounds()
	pixels := make([]rgb, 0)

	for y := b.Min.Y; y < b.Max.Y; y += step {
		for x := b.Min.X; x < b.Max.X; x += step {
			r, g, b, _ := img.At(x, y).RGBA()
			pixels = append(pixels, rgb{
				r: float64(r >> 8),
				g: float64(g >> 8),
				b: float64(b >> 8),
			})
		}
	}

	return pixels
}

func ColorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("%02X%02X%02X", uint8(r>>8), uint8(g>>8), uint8(b>>8))
}
