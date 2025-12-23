package palette

import "image/color"

var Light = Theme{
	Background: color.RGBA{245, 245, 245, 255},
	Text:       color.RGBA{20, 20, 20, 255},
	Accent:     color.RGBA{60, 120, 255, 255},

	Icons: "icons/icons-light.png",
}
