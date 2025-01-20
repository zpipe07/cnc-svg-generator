package signs

import (
	"image/color"
	"log"

	"github.com/tdewolff/canvas"
)

func drawEllipse(
	ctx *canvas.Context,
	width float64,
	height float64,
	foregroundColor color.RGBA,
	backgroundColor color.RGBA,
	lines []string,
	fontFamily *canvas.FontFamily,
) {
	// draw shapes
	ctx.SetStrokeColor(foregroundColor)
	ctx.SetStrokeWidth(0.01)
	ctx.SetFillColor(backgroundColor)
	outerEdge := canvas.Ellipse(width/2, height/2)
	outerEdge = outerEdge.Translate(width/2, height/2)
	ctx.DrawPath(0, 0, outerEdge)

	ctx.SetStroke(nil)
	ctx.SetFillColor(foregroundColor)
	roundedEdge := canvas.Ellipse(width/2-0.2, height/2-0.2)
	roundedEdge = roundedEdge.Translate(width/2, height/2)
	ctx.DrawPath(0, 0, roundedEdge)

	ctx.SetFillColor(backgroundColor)
	borderOuter := canvas.Ellipse(width/2-0.5, height/2-0.5)
	borderOuter = borderOuter.Translate(width/2, height/2)
	ctx.DrawPath(0, 0, borderOuter)

	ctx.SetFillColor(foregroundColor)
	borderInner := canvas.Ellipse(width/2-0.7, height/2-0.7)
	borderInner = borderInner.Translate(width/2, height/2)
	ctx.DrawPath(0, 0, borderInner)

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
