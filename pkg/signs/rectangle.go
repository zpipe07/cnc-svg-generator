package signs

import (
	"image/color"
	"log"

	"github.com/tdewolff/canvas"
)

func drawRectangle(
	ctx *canvas.Context,
	width float64,
	height float64,
	foregroundColor color.RGBA,
	backgroundColor color.RGBA,
	lines []string,
	fontFamily *canvas.FontFamily,
) {
	// draw shapes
	ctx.SetFillColor(backgroundColor)
	ctx.SetStrokeColor(foregroundColor)
	ctx.SetStrokeWidth(0.01)
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
		// draw text container
		ctx.SetFill(nil)

		// uncomment to draw text container border
		// ctx.SetStrokeWidth(0.1)
		// ctx.SetStrokeColor(canvas.Lightgray)

		var container *canvas.Path
		var containerX, containerY float64
		var containerBounds canvas.Rect
		if i == 0 {
			container = canvas.Rectangle(width-2.0, 2.0)
			containerBounds = container.Bounds()
			containerY = height/2 - containerBounds.H()/2 + 3.5
		} else if i == 1 {
			container = canvas.Rectangle(width-2.0, 4.0)
			containerBounds = container.Bounds()
			containerY = height/2 - containerBounds.H()/2
		} else if i == 2 {
			container = canvas.Rectangle(width-2.0, 2.0)
			containerBounds = container.Bounds()
			containerY = height/2 - containerBounds.H()/2 - 3.5
		}
		containerX = width/2 - containerBounds.W()/2
		ctx.DrawPath(containerX, containerY, container)

		// draw text
		ctx.SetStroke(nil)
		ctx.SetFillColor(backgroundColor)
		fontSize := 1.0
		face := fontFamily.Face(fontSize, canvas.FontRegular, canvas.FontNormal)
		textPath, _, err := face.ToPath(line)
		if err != nil {
			log.Fatalf("Failed to convert text to path: %s", err)
		}
		textBounds := textPath.Bounds()
		metrics := face.Metrics()
		ascent := metrics.Ascent
		descent := metrics.Descent
		log.Print("descent: ", descent)
		// totalHeight := ascent + descent

		// Calculate the scale factor to fit the path within the container
		scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/textBounds.H())
		// scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/ascent+descent)
		log.Print("scale: ", scale)
		textPath.Scale(scale, scale)

		// recalculate the bounds after scaling
		textBounds = textPath.Bounds()
		x := containerX +
			containerBounds.W()/2 -
			textBounds.W()/2
		y := containerY +
			containerBounds.H()/2 +
			descent*scale -
			((ascent+descent)/2)*scale

		ctx.DrawPath(x, y, textPath)
	}
}
