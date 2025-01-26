package signs

import (
	"cnc-svg-generator/pkg/svgutils"
	"fmt"
	"log"

	"github.com/tdewolff/canvas"
)

func drawRectangle(
	// ctx *canvas.Context,
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
	outerEdge := canvas.RoundedRectangle(width, height, 0.25)
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "outer-edge",
	})
	builder.EndGroup()

	// Add the rounded edge
	roundedEdge := canvas.RoundedRectangle(width-0.5, height-0.5, 0.2)
	roundedEdge = roundedEdge.Translate(0.25, 0.25)
	builder.StartGroup("Rounded Edge", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "rounded-edge",
	})
	builder.EndGroup()

	// Add the border
	borderOuter := canvas.RoundedRectangle(width-1.0, height-1.0, 0.15)
	borderOuter = borderOuter.Translate(0.5, 0.5)
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "my-custom-id",
	})
	borderInner := canvas.RoundedRectangle(width-1.25, height-1.25, 0.1)
	borderInner = borderInner.Translate(0.625, 0.625)
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
