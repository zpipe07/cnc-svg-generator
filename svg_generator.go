package main

import (
	"bytes"
	"image/color"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/svg"
)

func generateSVG(
	productConfig ProductConfig,
	lines []string,
	fontFamily *canvas.FontFamily,
	foregroundColor color.RGBA,
	backgroundColor color.RGBA,
) string {

	width := productConfig.Width
	height := productConfig.Height

	// Create a new canvas with dynamic width and height
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	// Draw the appropriate shape around the text
	ctx.SetStrokeWidth(0.0125)

	drawSvg(
		ctx,
		productConfig,
		foregroundColor,
		backgroundColor,
		lines,
		fontFamily,
	)
	// drawText(ctx, lines, fontFamily, width, height, foregroundColor, backgroundColor)

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
