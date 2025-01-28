package signs

import (
	"cnc-svg-generator/pkg/svgutils"
	"fmt"
	"log"

	"github.com/tdewolff/canvas"
)

func drawRectangle(
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
	numLines := len(lines)

	if numLines > 0 {
		// Calculate the total height for the text containers
		const padding = 1.0 // Padding from top and bottom
		lineSpacing := 0.5  // Spacing between lines
		availableHeight := height - 2*padding - float64(numLines-1)*lineSpacing

		// Determine the heights for each line
		var containerHeights []float64
		switch numLines {
		case 3:
			containerHeights = []float64{
				availableHeight * 0.25, // First line
				availableHeight * 0.5,  // Middle line (taller)
				availableHeight * 0.25, // Last line
			}
		case 2:
			containerHeights = []float64{
				availableHeight * 0.6, // First line (taller)
				availableHeight * 0.4, // Last line
			}
		default:
			containerHeights = []float64{availableHeight} // Single line
		}

		// Calculate the starting y position for the topmost line
		currentY := height - padding

		for i, line := range lines {
			containerHeight := containerHeights[i]
			container := canvas.Rectangle(width-2.0, containerHeight)
			containerBounds := container.Bounds()
			containerX := width/2 - containerBounds.W()/2
			containerY := currentY - containerHeight

			// Position the container
			container = container.Translate(containerX, containerY)
			container = container.Scale(1, -1)
			container = container.Translate(0, height)

			builder.AddPath(container.ToSVG(), map[string]string{
				"fill":         "none",
				"stroke":       "pink",
				"stroke-width": "0.025",
				"id":           fmt.Sprintf("text-container-%d", i),
			})

			// Draw text
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

			// Scale the text to fit within the container
			scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/textBounds.H())
			textPath.Scale(scale, scale)

			// Recalculate text position after scaling
			textBounds = textPath.Bounds()
			x := containerX +
				containerBounds.W()/2 -
				textBounds.W()/2
			y := containerY +
				containerBounds.H()/2 +
				descent*scale -
				((ascent + descent) / 2 * scale)
			textPath = textPath.Translate(x, y)
			textPath = textPath.Scale(1, -1)
			textPath = textPath.Translate(0, height)

			builder.AddPath(textPath.ToSVG(), map[string]string{
				"fill": backgroundColor,
				"id":   fmt.Sprintf("text-line-%d", i),
			})

			// Update currentY for the next line
			currentY -= containerHeight + lineSpacing
		}
	}

	builder.EndGroup()

	return builder.Close()
}
