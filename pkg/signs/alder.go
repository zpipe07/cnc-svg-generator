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
	outerEdge := canvas.RoundedRectangle(width, height, 0.2)
	outerEdge = outerEdge.Translate(0, 0)
	ctx.DrawPath(0, 0, outerEdge)

	ctx.SetFillColor(backgroundColor)
	borderOuter, err := canvas.ParseSVGPath("M 1.13 9.495 A 0.75 0.75 0 0 0 0.38 8.745 L 0.38 1.255 A 0.75 0.75 0 0 0 1.13 0.505 L 21.87 0.505 A 0.75 0.75 0 0 0 22.62 1.255 L 22.62 8.745 A 0.75 0.75 0 0 0 21.87 9.495 L 1.13 9.495 Z")
	if err != nil {
		log.Printf("Failed to parse SVG path: %s", err)
	}
	// roundedEdge := canvas.RoundedRectangle(width-0.4, height-0.4, 0.1)
	// roundedEdge = roundedEdge.Translate(0.2, 0.2)
	ctx.DrawPath(0, 0, borderOuter)

	ctx.SetFillColor(foregroundColor)
	borderInner, err := canvas.ParseSVGPath("M 1.309 9.295 A 0.95 0.95 0 0 0 0.58 8.566 L 0.58 1.434 A 0.95 0.95 0 0 0 1.309 0.705 L 21.691 0.705 A 0.95 0.95 0 0 0 22.42 1.434 L 22.42 8.566 A 0.95 0.95 0 0 0 21.691 9.295 L 1.309 9.295 Z")
	if err != nil {
		log.Printf("Failed to parse SVG path: %s", err)
	}
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
			x = 2
			y = height/2 - textBounds.H()/2
		} else if i == 1 {
			x = 6
			y = height/2 - textBounds.H()/2 + 1.5
		} else if i == 2 {
			x = 6
			y = height/2 - textBounds.H()/2 - 1.5
		}
		// lineHeight := textBounds.H() + 0.75
		// startY := ((height - (float64(len(lines)) * lineHeight)) / 2) - yOffset
		// textX := (width - textBounds.W()) / 2
		// textY := startY + lineHeight*float64(i+1) - textBounds.H() // Center the text vertically
		// textPath = textPath.Translate(textX, textY)
		ctx.DrawPath(x, y, textPath)
	}

	// draw decoration
	horizontalLine := canvas.Rectangle(width/2, 0.25)
	ctx.DrawPath(6, height/2-0.05, horizontalLine)
}
