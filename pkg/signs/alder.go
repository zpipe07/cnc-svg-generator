package signs

import (
	"cnc-svg-generator/pkg/svgutils"
	"fmt"
	"log"

	"github.com/tdewolff/canvas"
)

func drawAlder(
	width float64,
	height float64,
	foregroundColor string,
	backgroundColor string,
	lines []string,
	fontFamily *canvas.FontFamily,
) string {
	// Initialize SVG builder
	builder := svgutils.NewSVGBuilder(width, height)

	// Add the outer edge
	outerEdge := canvas.RoundedRectangle(width, height, 0.2)
	outerEdge = outerEdge.Translate(0, 0)
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "outer-edge",
	})
	builder.EndGroup()

	// Add the border
	borderOuter, err := canvas.ParseSVGPath("M 1.13 9.495 A 0.75 0.75 0 0 0 0.38 8.745 L 0.38 1.255 A 0.75 0.75 0 0 0 1.13 0.505 L 21.87 0.505 A 0.75 0.75 0 0 0 22.62 1.255 L 22.62 8.745 A 0.75 0.75 0 0 0 21.87 9.495 L 1.13 9.495 Z")
	if err != nil {
		log.Printf("Failed to parse SVG path: %s", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "my-custom-id",
	})
	borderInner, err := canvas.ParseSVGPath("M 1.309 9.295 A 0.95 0.95 0 0 0 0.58 8.566 L 0.58 1.434 A 0.95 0.95 0 0 0 1.309 0.705 L 21.691 0.705 A 0.95 0.95 0 0 0 22.42 1.434 L 22.42 8.566 A 0.95 0.95 0 0 0 21.691 9.295 L 1.309 9.295 Z")
	if err != nil {
		log.Printf("Failed to parse SVG path: %s", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "my-custom-id",
	})

	// draw text
	for i, line := range lines {
		// draw text container
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
		// *************************************************************
		// ********** uncomment to draw text container border **********
		// *************************************************************
		container = container.Translate(containerX, containerY)
		container = container.Scale(1, -1)
		container = container.Translate(0, height)
		builder.AddPath(container.ToSVG(), map[string]string{
			"fill":         "none",
			"stroke":       "pink",
			"stroke-width": "0.025",
			"id":           fmt.Sprintf("text-container-%d", i),
		})

		// draw text
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

		// Calculate the scale factor to fit the path within the container
		scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/textBounds.H())
		textPath.Scale(scale, scale)

		// recalculate the bounds after scaling
		textBounds = textPath.Bounds()
		x := containerX
		y := containerY + containerBounds.H()/2 - textBounds.H()/2 + ascent + descent/2
		textPath = textPath.Translate(x, y)
		textPath = textPath.Scale(1, -1)
		textPath = textPath.Translate(0, height)
		builder.AddPath(textPath.ToSVG(), map[string]string{
			"fill": backgroundColor,
			"id":   fmt.Sprintf("text-line-%d", i),
		})
	}

	// draw decoration
	horizontalLine := canvas.Rectangle(width/2, 0.25)
	horizontalLine = horizontalLine.Translate(6, height/2-0.125)
	builder.AddPath(horizontalLine.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "horizontal-line",
	})

	builder.EndGroup()

	return builder.Close()
}
