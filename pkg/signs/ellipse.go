package signs

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

func drawEllipse(
	ctx *canvas.Context,
	width float64,
	height float64,
	foregroundColor color.RGBA,
	backgroundColor color.RGBA,
) {
	ctx.SetFillColor(backgroundColor)
	outerEdge := canvas.Ellipse(width/2, height/2)
	outerEdge = outerEdge.Translate(width/2, height/2)
	ctx.DrawPath(0, 0, outerEdge)

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
}
