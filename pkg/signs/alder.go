package signs

import (
	"image/color"
	"log"

	"github.com/tdewolff/canvas"
)

func drawAlder(
	ctx *canvas.Context,
	width float64,
	height float64,
	foregroundColor color.RGBA,
	backgroundColor color.RGBA,
	lines []string,
	fontFamily *canvas.FontFamily,
) {
	// draw shapes
	ctx.SetFillColor(foregroundColor)
	ctx.SetStrokeColor(backgroundColor)
	ctx.SetStrokeWidth(0.025)
	outerEdge := canvas.RoundedRectangle(width, height, 0.2)
	outerEdge = outerEdge.Translate(0, 0)
	ctx.DrawPath(0, 0, outerEdge)

	ctx.SetStroke(nil)
	ctx.SetFillColor(backgroundColor)
	borderOuter, err := canvas.ParseSVGPath("M 1.13 9.495 A 0.75 0.75 0 0 0 0.38 8.745 L 0.38 1.255 A 0.75 0.75 0 0 0 1.13 0.505 L 21.87 0.505 A 0.75 0.75 0 0 0 22.62 1.255 L 22.62 8.745 A 0.75 0.75 0 0 0 21.87 9.495 L 1.13 9.495 Z")
	if err != nil {
		log.Printf("Failed to parse SVG path: %s", err)
	}
	ctx.DrawPath(0, 0, borderOuter)

	ctx.SetFillColor(foregroundColor)
	borderInner, err := canvas.ParseSVGPath("M 1.309 9.295 A 0.95 0.95 0 0 0 0.58 8.566 L 0.58 1.434 A 0.95 0.95 0 0 0 1.309 0.705 L 21.691 0.705 A 0.95 0.95 0 0 0 22.42 1.434 L 22.42 8.566 A 0.95 0.95 0 0 0 21.691 9.295 L 1.309 9.295 Z")
	if err != nil {
		log.Printf("Failed to parse SVG path: %s", err)
	}
	ctx.DrawPath(0, 0, borderInner)

	// draw text
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
			container = canvas.Rectangle(4, 4)
			containerBounds = container.Bounds()
			containerX = 1
			containerY = height/2 - containerBounds.H()/2
		} else if i == 1 {
			container = canvas.Rectangle(width-7, 2.75)
			containerBounds = container.Bounds()
			containerX = width/2 - containerBounds.W()/2 + 2.25
			containerY = height/2 - containerBounds.H()/2 + 2.25
		} else if i == 2 {
			container = canvas.Rectangle(width-7, 2.75)
			containerBounds = container.Bounds()
			containerX = width/2 - containerBounds.W()/2 + 2.25
			containerY = height/2 - containerBounds.H()/2 - 2.25
		}
		ctx.DrawPath(containerX, containerY, container)

		// draw text
		ctx.SetStroke(nil)
		ctx.SetFillColor(backgroundColor)
		fontSize := 10 - float64(len(line))*0.45
		face := fontFamily.Face(fontSize, canvas.FontRegular, canvas.FontNormal)
		textPath, _, err := face.ToPath(line)
		if err != nil {
			log.Fatalf("Failed to convert text to path: %s", err)
		}
		textBounds := textPath.Bounds()

		// Calculate the scale factor to fit the path within the container
		scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/textBounds.H())
		textPath.Scale(scale, scale)

		// recalculate the bounds after scaling
		textBounds = textPath.Bounds()
		x := containerX + containerBounds.W()/2 - textBounds.W()/2
		y := containerY + containerBounds.H()/2 - textBounds.H()/2
		ctx.DrawPath(x, y, textPath)
	}

	// draw decoration
	horizontalLine := canvas.Rectangle(width/2, 0.25)
	ctx.DrawPath(6, height/2-0.125, horizontalLine)
}
