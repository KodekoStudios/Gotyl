package palette

import "image/color"

type Theme struct {
	Background color.Color
	Text       color.Color
	Accent     color.Color

	Icons string
}
