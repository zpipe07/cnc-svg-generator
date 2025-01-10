package main

import (
	"bytes"
	"image/color"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/svg"
)

func generateSVG(
	lines []string,
	shape string,
	fontFamily *canvas.FontFamily,
	foregroundColor color.RGBA,
	backgroundColor color.RGBA,
) string {
	// Hardcoded dimensions in inches
	width := 15.0
	height := 11.0

	// Create a new canvas with dynamic width and height
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	// Draw the appropriate shape around the text
	ctx.SetStrokeWidth(0.0125)

	drawShape(ctx, width, height, shape, foregroundColor, backgroundColor)
	drawText(ctx, lines, fontFamily, width, height, foregroundColor, backgroundColor)

	// Export the canvas to an SVG
	var buf bytes.Buffer
	opts := &svg.Options{
		SizeUnits: "in",
	}
	svgCanvas := svg.New(&buf, width, height, opts)
	c.RenderTo(svgCanvas)
	svgCanvas.Close()

	return buf.String()
}
