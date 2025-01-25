package signs

import (
	"image/color"
	"log"

	"github.com/tdewolff/canvas"
)

func drawSesame(
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
	outerEdge, err := canvas.ParseSVGPath("M 0 10.75 A 0.25 0.25 0 0 0 0.25 11 L 7.496 11 A 8.039 8.039 0 0 0 7.5 11 A 7.696 7.696 0 0 0 7.582 11 L 14.75 11 A 0.25 0.25 0 0 0 15 10.75 L 15 3 A 0.25 0.25 0 0 0 14.75 2.75 L 13.408 2.75 A 0.25 0.25 0 0 1 13.207 2.649 A 4.879 4.879 0 0 0 12.919 2.296 A 5.937 5.937 0 0 0 12.215 1.612 A 0.25 0.25 0 0 0 12.213 1.61 A 7.213 7.213 0 0 0 8.747 0.097 A 8.039 8.039 0 0 0 7.5 0 A 7.696 7.696 0 0 0 3.742 0.957 A 6.554 6.554 0 0 0 2.786 1.611 A 5.524 5.524 0 0 0 1.801 2.647 A 0.25 0.25 0 0 1 1.599 2.75 L 0.25 2.75 A 0.25 0.25 0 0 0 0 3 L 0 10.75 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	outerBounds := outerEdge.Bounds()
	outerEdge = outerEdge.Translate(-outerBounds.X0, -outerBounds.Y0)
	outerEdge = outerEdge.Scale(1, -1)
	outerEdge = outerEdge.Translate(0, height)
	ctx.DrawPath(0, 0, outerEdge)

	ctx.SetStroke(nil)
	ctx.SetFillColor(foregroundColor)
	roundedEdge, err := canvas.ParseSVGPath("M 0.2 10.75 A 0.05 0.05 0 0 0 0.25 10.8 L 7.496 10.8 A 0.2 0.2 0 0 1 7.498 10.8 A 0.2 0.2 0 0 1 7.5 10.8 A 0.2 0.2 0 0 1 7.501 10.8 A 7.496 7.496 0 0 0 7.581 10.8 A 0.2 0.2 0 0 1 7.582 10.8 L 14.75 10.8 A 0.05 0.05 0 0 0 14.8 10.75 L 14.8 3 A 0.05 0.05 0 0 0 14.75 2.95 L 13.408 2.95 A 0.45 0.45 0 0 1 13.046 2.768 A 4.679 4.679 0 0 0 12.77 2.43 A 0.2 0.2 0 0 1 12.768 2.428 A 5.737 5.737 0 0 0 12.088 1.766 A 7.013 7.013 0 0 0 8.718 0.295 A 0.2 0.2 0 0 1 8.716 0.295 A 7.839 7.839 0 0 0 7.5 0.2 A 0.2 0.2 0 0 1 7.499 0.2 A 7.496 7.496 0 0 0 3.84 1.132 A 6.354 6.354 0 0 0 2.913 1.765 A 0.2 0.2 0 0 1 2.912 1.766 A 5.324 5.324 0 0 0 1.963 2.765 A 0.45 0.45 0 0 1 1.599 2.95 L 0.25 2.95 A 0.05 0.05 0 0 0 0.2 3 L 0.2 10.75 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	roundedEdge = roundedEdge.Scale(1, -1)
	roundedEdge = roundedEdge.Translate(0, height)
	ctx.DrawPath(0, 0, roundedEdge)

	ctx.SetFillColor(backgroundColor)
	borderOuter, err := canvas.ParseSVGPath("M 0.5 10.5 L 7.496 10.5 A 0.3 0.3 0 0 1 7.499 10.5 A 0.3 0.3 0 0 1 7.503 10.5 A 7.196 7.196 0 0 0 7.579 10.5 A 0.3 0.3 0 0 1 7.58 10.5 A 0.3 0.3 0 0 1 7.582 10.5 L 14.5 10.5 L 14.5 3.25 L 13.408 3.25 A 0.75 0.75 0 0 1 12.805 2.947 A 4.379 4.379 0 0 0 12.547 2.63 A 0.3 0.3 0 0 1 12.544 2.627 A 0.3 0.3 0 0 1 12.542 2.625 A 5.437 5.437 0 0 0 11.899 1.999 A 6.713 6.713 0 0 0 8.674 0.592 A 0.3 0.3 0 0 1 8.672 0.591 A 0.3 0.3 0 0 1 8.669 0.591 A 7.539 7.539 0 0 0 7.5 0.5 A 0.3 0.3 0 0 1 7.499 0.5 A 0.3 0.3 0 0 1 7.497 0.5 A 7.196 7.196 0 0 0 3.986 1.394 A 6.054 6.054 0 0 0 3.104 1.997 A 0.3 0.3 0 0 1 3.102 1.998 A 0.3 0.3 0 0 1 3.101 1.999 A 5.024 5.024 0 0 0 2.205 2.942 A 0.75 0.75 0 0 1 1.599 3.25 L 0.5 3.25 L 0.5 10.5 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	borderOuter = borderOuter.Scale(1, -1)
	borderOuter = borderOuter.Translate(0, height)
	ctx.DrawPath(0, 0, borderOuter)

	ctx.SetFillColor(foregroundColor)
	borderInner, err := canvas.ParseSVGPath("M 0.7 10.3 L 7.496 10.3 A 0.5 0.5 0 0 1 7.5 10.3 A 0.2 0.2 0 0 1 7.504 10.3 A 6.996 6.996 0 0 0 7.578 10.3 A 0.2 0.2 0 0 1 7.58 10.3 A 0.2 0.2 0 0 1 7.582 10.3 L 14.3 10.3 L 14.3 3.45 L 13.408 3.45 A 0.95 0.95 0 0 1 12.645 3.066 A 4.179 4.179 0 0 0 12.398 2.764 A 0.2 0.2 0 0 1 12.394 2.76 A 0.2 0.2 0 0 1 12.391 2.756 A 5.237 5.237 0 0 0 11.772 2.154 A 6.513 6.513 0 0 0 8.645 0.79 A 0.2 0.2 0 0 1 8.641 0.789 A 0.2 0.2 0 0 1 8.638 0.789 A 7.339 7.339 0 0 0 7.5 0.7 A 0.2 0.2 0 0 1 7.498 0.7 A 0.2 0.2 0 0 1 7.496 0.7 A 6.996 6.996 0 0 0 4.083 1.569 A 5.854 5.854 0 0 0 3.231 2.152 A 0.2 0.2 0 0 1 3.229 2.153 A 0.2 0.2 0 0 1 3.227 2.155 A 4.824 4.824 0 0 0 2.367 3.06 A 0.95 0.95 0 0 1 1.599 3.45 L 0.7 3.45 L 0.7 10.3 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	borderInner = borderInner.Scale(1, -1)
	borderInner = borderInner.Translate(0, height)
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
			container = canvas.Rectangle(width-8, 3.5)
			containerBounds = container.Bounds()
			containerX = width/2 - containerBounds.W()/2
			containerY = 6.0
			// continue
		} else if i == 1 {
			container = canvas.Rectangle(width-2, 3.0)
			containerBounds = container.Bounds()
			containerX = width/2 - containerBounds.W()/2
			containerY = 2.75
		} else if i == 2 {
			container = canvas.Rectangle(width-2, 1.5)
			containerBounds = container.Bounds()
			containerX = width/2 - containerBounds.W()/2
			containerY = 1.0
		}
		ctx.DrawPath(containerX, containerY, container)

		// draw text
		ctx.SetStroke(nil)
		ctx.SetFillColor(backgroundColor)
		fontSize := 1.0
		face := fontFamily.Face(fontSize, canvas.Black, canvas.FontRegular, canvas.FontNormal)
		textPath, _, err := face.ToPath(line)
		if err != nil {
			panic(err)
		}
		textBounds := textPath.Bounds()
		metrics := face.Metrics()
		ascent := metrics.Ascent
		descent := metrics.Descent
		totalHeight := ascent + descent

		// Calculate the scale factor to fit the path within the container
		scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/textBounds.H())
		textPath.Scale(scale, scale)

		// recalculate the bounds after scaling
		textBounds = textPath.Bounds()
		x := containerX + containerBounds.W()/2 - textBounds.W()/2
		y := containerY + containerBounds.H()/2 - (totalHeight/2)*scale + descent*scale

		ctx.DrawPath(x, y, textPath)
	}
}
