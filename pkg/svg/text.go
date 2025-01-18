package svg

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

func drawText(
	ctx *canvas.Context,
	lines []string,
	fontFamily *canvas.FontFamily,
	width float64,
	height float64,
	foregroundColor color.RGBA,
	backgroundColor color.RGBA,
) {

	yOffset := 0.0

	ctx.SetFillColor(backgroundColor)

	// Draw each line of text inside the shape
	for i, line := range lines {
		fontSize := 10 - float64(len(line))*0.45

		face := fontFamily.Face(fontSize, canvas.FontRegular, canvas.FontNormal)
		textPath, _, err := face.ToPath(line)
		if err != nil {
			panic("Failed to convert text to path: " + err.Error())
		}
		textBounds := textPath.Bounds()
		lineHeight := textBounds.H() + 0.75
		startY := ((height - (float64(len(lines)) * lineHeight)) / 2) - yOffset
		textX := (width - textBounds.W()) / 2
		textY := startY + lineHeight*float64(i+1) - textBounds.H() // Center the text vertically
		textPath = textPath.Translate(textX, textY)
		ctx.DrawPath(0, 0, textPath)
	}
}
