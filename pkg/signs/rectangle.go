package signs

import (
	"cnc-svg-generator/pkg/svgutils"
	"fmt"
	"log"

	"github.com/tdewolff/canvas"
)

// rectangleDefs returns the SVG defs content (filters and gradients) for depth effects.
// Coordinates are in inches. Values tuned for clearly visible depth.
func rectangleDefs() string {
	return `
  <!-- Inner shadow for V-carved border groove: stronger recessed look -->
  <filter id="groove-shadow" x="-50%" y="-50%" width="200%" height="200%">
    <feComponentTransfer in="SourceAlpha">
      <feFuncA type="table" tableValues="1 0"/>
    </feComponentTransfer>
    <feGaussianBlur stdDeviation="0.08"/>
    <feOffset dx="0.06" dy="0.06" result="offsetblur"/>
    <feFlood flood-color="#000000" flood-opacity="0.65" result="color"/>
    <feComposite in2="offsetblur" operator="in"/>
    <feComposite in2="SourceAlpha" operator="in"/>
    <feMerge>
      <feMergeNode in="SourceGraphic"/>
      <feMergeNode/>
    </feMerge>
  </filter>

  <!-- Inner shadow for engraved text: more visible carved effect -->
  <filter id="text-engrave" x="-50%" y="-50%" width="200%" height="200%">
    <feComponentTransfer in="SourceAlpha">
      <feFuncA type="table" tableValues="1 0"/>
    </feComponentTransfer>
    <feGaussianBlur stdDeviation="0.06"/>
    <feOffset dx="0.04" dy="0.04" result="offsetblur"/>
    <feFlood flood-color="#000000" flood-opacity="0.7" result="color"/>
    <feComposite in2="offsetblur" operator="in"/>
    <feComposite in2="SourceAlpha" operator="in"/>
    <feMerge>
      <feMergeNode in="SourceGraphic"/>
      <feMergeNode/>
    </feMerge>
  </filter>

  <!-- Surface lighting gradient: stronger highlight and shade for depth -->
  <linearGradient id="surface-highlight" x1="0" y1="0" x2="0" y2="1">
    <stop offset="0%" stop-color="white" stop-opacity="0.18"/>
    <stop offset="50%" stop-color="white" stop-opacity="0.06"/>
    <stop offset="100%" stop-color="black" stop-opacity="0.12"/>
  </linearGradient>
`
}

func drawRectangle(
	width float64,
	height float64,
	foregroundColor string,
	backgroundColor string,
	lines []string,
	fontFamily *canvas.FontFamily,
	strokeOnly bool,
) string {
	if !strokeOnly {
		foregroundColor = GetColor(foregroundColor)
		backgroundColor = GetColor(backgroundColor)
	}

	// Initialize SVG builder
	builder := svgutils.NewSVGBuilder(width, height)

	if !strokeOnly {
		builder.AddDefs(rectangleDefs())
	}

	// Add the outer edge
	outerEdge := canvas.RoundedRectangle(width, height, 0.25)
	outerEdge = outerEdge.ReplaceArcs()
	builder.StartGroup("Outer Edge", map[string]string{})
	outerEdgeAttrs := map[string]string{
		"fill":         map[bool]string{true: "none", false: backgroundColor}[strokeOnly],
		"id":           "outer-edge",
		"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
		"stroke-width": "0.025",
	}
	if !strokeOnly {
		outerEdgeAttrs["filter"] = "url(#sign-shadow)"
	}
	builder.AddPath(outerEdge.ToSVG(), outerEdgeAttrs)
	builder.EndGroup()

	// Add the rounded edge
	roundedEdge := canvas.RoundedRectangle(width-0.5, height-0.5, 0.2)
	roundedEdge = roundedEdge.Translate(0.25, 0.25)
	roundedEdge = roundedEdge.ReplaceArcs()
	roundedEdgePath := roundedEdge.ToSVG()
	builder.StartGroup("Rounded Edge", map[string]string{})
	builder.AddPath(roundedEdgePath, map[string]string{
		"fill":         map[bool]string{true: "none", false: foregroundColor}[strokeOnly],
		"id":           "rounded-edge",
		"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
		"stroke-width": "0.025",
	})
	if !strokeOnly {
		builder.AddPath(roundedEdgePath, map[string]string{
			"fill":   "url(#surface-highlight)",
			"id":     "rounded-edge-highlight",
			"stroke": "none",
		})
	}
	builder.EndGroup()

	// Add the border
	borderOuter := canvas.RoundedRectangle(width-1.0, height-1.0, 0.15)
	borderOuter = borderOuter.Translate(0.5, 0.5)
	borderOuter = borderOuter.ReplaceArcs()
	builder.StartGroup("Vcarve", map[string]string{})
	borderOuterAttrs := map[string]string{
		"fill":         map[bool]string{true: "none", false: backgroundColor}[strokeOnly],
		"id":           "my-custom-id",
		"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
		"stroke-width": "0.025",
	}
	if !strokeOnly {
		borderOuterAttrs["filter"] = "url(#groove-shadow)"
	}
	builder.AddPath(borderOuter.ToSVG(), borderOuterAttrs)
	borderInner := canvas.RoundedRectangle(width-1.25, height-1.25, 0.1)
	borderInner = borderInner.Translate(0.625, 0.625)
	borderInner = borderInner.ReplaceArcs()
	borderInnerAttrs := map[string]string{
		"fill":         map[bool]string{true: "none", false: foregroundColor}[strokeOnly],
		"id":           "my-custom-id",
		"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
		"stroke-width": "0.025",
	}
	if !strokeOnly {
		borderInnerAttrs["filter"] = "url(#groove-shadow)"
	}
	builder.AddPath(borderInner.ToSVG(), borderInnerAttrs)

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
				"fill":   "none",
				"stroke": "none",
				// "stroke":       "pink",
				// "stroke-width": "0.025",
				"id": fmt.Sprintf("text-container-%d", i),
			})

			// Draw text
			fontSize := 1.0
			face := fontFamily.Face(fontSize, canvas.FontRegular, canvas.FontNormal)
			textPath, _, err := face.ToPath(line)
			if err != nil {
				log.Fatalf("Failed to convert text to path: %s", err)
			}
			textBounds := textPath.Bounds()

			// Scale the text to fit within the container
			scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/textBounds.H())
			textPath = textPath.Scale(scale, scale)

			// Recalculate text bounds after scaling to get actual rendered dimensions
			textBounds = textPath.Bounds()

			// Calculate the center of the container
			centerX := containerX + containerBounds.W()/2
			centerY := containerY + containerBounds.H()/2

			// Calculate the center of the text's actual bounding box
			// This uses the real rendered glyph bounds, not font metrics,
			// which ensures proper centering regardless of font design choices
			textCenterX := textBounds.X0 + textBounds.W()/2
			textCenterY := textBounds.Y0 + textBounds.H()/2

			// Translate so text center aligns with container center
			x := centerX - textCenterX
			y := centerY - textCenterY

			textPath = textPath.Translate(x, y)
			textPath = textPath.Scale(1, -1)
			textPath = textPath.Translate(0, height)

			textAttrs := map[string]string{
				"fill":         map[bool]string{true: "none", false: backgroundColor}[strokeOnly],
				"id":           fmt.Sprintf("text-line-%d", i),
				"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
				"stroke-width": "0.025",
			}
			if !strokeOnly {
				textAttrs["filter"] = "url(#text-engrave)"
			}
			builder.AddPath(textPath.ToSVG(), textAttrs)

			// Update currentY for the next line
			currentY -= containerHeight + lineSpacing
		}
	}

	builder.EndGroup()

	return builder.Close()
}
