package signs

import (
	"image/color"
	"log"

	"github.com/tdewolff/canvas"
)

func drawRectangle(ctx *canvas.Context,
	width float64,
	height float64,
	foregroundColor color.RGBA,
	backgroundColor color.RGBA,
	lines []string,
	fontFamily *canvas.FontFamily) {
	ctx.SetFillColor(backgroundColor)
	ctx.SetStrokeColor(foregroundColor)
	ctx.SetStrokeWidth(0.025)
	outerEdge := canvas.RoundedRectangle(width, height, 0.25)
	ctx.DrawPath(0, 0, outerEdge)

	ctx.SetStroke(nil)
	ctx.SetFillColor(foregroundColor)
	roundedEdge := canvas.RoundedRectangle(width-0.5, height-0.5, 0.2)
	ctx.DrawPath(0.25, 0.25, roundedEdge)

	ctx.SetFillColor(backgroundColor)
	borderOuter := canvas.RoundedRectangle(width-1.0, height-1.0, 0.15)
	ctx.DrawPath(0.5, 0.5, borderOuter)

	ctx.SetFillColor(foregroundColor)
	borderInner := canvas.RoundedRectangle(width-1.25, height-1.25, 0.1)
	ctx.DrawPath(0.625, 0.625, borderInner)

	// draw text
	ctx.SetFillColor(backgroundColor)
	for i, line := range lines {
		fontSize := 10 - float64(len(line))*0.45
		face := fontFamily.Face(fontSize, canvas.FontRegular, canvas.FontNormal)
		textPath, _, err := face.ToPath(line)
		if err != nil {
			log.Fatalf("Failed to convert text to path: %s", err)
		}
		textBounds := textPath.Bounds()
		var x, y float64
		if i == 0 {
			x = width/2 - textBounds.W()/2
			y = height/2 - textBounds.H()/2 + 2.75
		} else if i == 1 {
			x = width/2 - textBounds.W()/2
			y = height/2 - textBounds.H()/2
		} else if i == 2 {
			x = width/2 - textBounds.W()/2
			y = height/2 - textBounds.H()/2 - 2.75
		}
		ctx.DrawPath(x, y, textPath)
	}
}
