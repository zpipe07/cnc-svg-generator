package signs

import (
	"cnc-svg-generator/pkg/svgutils"
	"fmt"
	"log"

	"github.com/tdewolff/canvas"
)

func drawEllipse(
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
	outerEdge := canvas.Ellipse(width/2, height/2)
	outerEdge = outerEdge.Translate(width/2, height/2)
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "outer-edge",
	})
	builder.EndGroup()

	// Add the rounded edge
	roundedEdge := canvas.Ellipse(width/2-0.2, height/2-0.2)
	roundedEdge = roundedEdge.Translate(width/2, height/2)
	builder.StartGroup("Rounded Edge", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "rounded-edge",
	})
	builder.EndGroup()

	// Add the border
	borderOuter := canvas.Ellipse(width/2-0.5, height/2-0.5)
	borderOuter = borderOuter.Translate(width/2, height/2)
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "my-custom-id",
	})
	borderInner := canvas.Ellipse(width/2-0.7, height/2-0.7)
	borderInner = borderInner.Translate(width/2, height/2)
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "my-custom-id",
	})

	// draw text
	for i, line := range lines {
		// Create text container
		var container *canvas.Path
		var containerX, containerY float64
		var containerBounds canvas.Rect
		if i == 0 {
			container = canvas.Rectangle(width-7.5, 2.25)
			containerBounds = container.Bounds()
			containerX = width/2 - containerBounds.W()/2
			containerY = height/2 - containerBounds.H()/2 + 3.0
		} else if i == 1 {
			container = canvas.Rectangle(width-4.0, 3.5)
			containerBounds = container.Bounds()
			containerX = width/2 - containerBounds.W()/2
			containerY = height/2 - containerBounds.H()/2
		} else if i == 2 {
			container = canvas.Rectangle(width-7.5, 2.25)
			containerBounds = container.Bounds()
			containerX = width/2 - containerBounds.W()/2
			containerY = height/2 - containerBounds.H()/2 - 3.0
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
		totalHeight := ascent + descent

		// Calculate the scale factor to fit the path within the container
		scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/textBounds.H())
		textPath.Scale(scale, scale)

		// recalculate the bounds after scaling
		textBounds = textPath.Bounds()
		x := containerX + containerBounds.W()/2 - textBounds.W()/2
		y := containerY + containerBounds.H()/2 - (totalHeight/2)*scale + descent*scale
		textPath = textPath.Translate(x, y)
		textPath = textPath.Scale(1, -1)
		textPath = textPath.Translate(0, height)
		builder.AddPath(textPath.ToSVG(), map[string]string{
			"fill": backgroundColor,
			"id":   fmt.Sprintf("text-line-%d", i),
		})
	}
	builder.EndGroup()

	return builder.Close()
}
