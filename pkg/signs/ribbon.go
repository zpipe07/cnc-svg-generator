package signs

import (
	"cnc-svg-generator/pkg/svgutils"
	"fmt"
	"log"

	"github.com/tdewolff/canvas"
)

func drawExtraSmallRibbon(
	builder *svgutils.SVGBuilder,
	width float64,
	height float64,
	foregroundColor string,
	backgroundColor string,
	strokeOnly bool,
) {
	// add the outer edge
	outerEdge := canvas.RoundedRectangle(width, height, 0.25)
	builder.StartGroup("Outside", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill":         backgroundColor,
		"id":           "outside-path",
		"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
		"stroke-width": "0.025",
	})
	builder.EndGroup()

	// add the rounded edge
	roundedEdge := canvas.RoundedRectangle(width-0.5, height-0.5, 0.2)
	roundedEdge = roundedEdge.Translate(0.25, 0.25)
	builder.StartGroup("Round", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill":         foregroundColor,
		"id":           "round-path",
		"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
		"stroke-width": "0.025",
	})
	builder.EndGroup()

	// add the border
	// borderOuter := canvas.RoundedRectangle(width-1.0, height-1.0, 0.15)
	// borderOuter = borderOuter.Translate(0.5, 0.5)
	builder.StartGroup("Vcarve", map[string]string{})
	// builder.AddPath(borderOuter.ToSVG(), map[string]string{
	// 	"fill":         backgroundColor,
	// 	"id":           "border-outer-path",
	// 	"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
	// 	"stroke-width": "0.025",
	// })
	// borderInner := canvas.RoundedRectangle(width-1.2, height-1.2, 0.1)
	// borderInner = borderInner.Translate(0.6, 0.6)
	// builder.AddPath(borderInner.ToSVG(), map[string]string{
	// 	"fill":         foregroundColor,
	// 	"id":           "border-inner-path",
	// 	"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
	// 	"stroke-width": "0.025",
	// })
	builder.EndGroup()
}
func drawExtraSmallRibbonText(
	builder *svgutils.SVGBuilder,
	width float64,
	height float64,
	backgroundColor string,
	lines []string,
	fontFamily *canvas.FontFamily,
	strokeOnly bool,
) {
	if len(lines) == 0 {
		return
	}

	// draw text container
	container := canvas.Rectangle(width-1.25, height-1.5)
	containerX := 0.625
	containerY := 0.75
	container = container.Translate(containerX, containerY)
	container = container.Scale(1, -1)
	container = container.Translate(0, height)
	containerBounds := container.FastBounds()
	builder.AddPath(container.ToSVG(), map[string]string{
		"fill": "none",
		// "stroke":       "pink",
		"stroke-width": "0.025",
		"id":           "text-container-path",
	})

	// draw text
	fontSize := 20.0
	face := fontFamily.Face(fontSize, canvas.FontRegular, canvas.FontNormal)
	textPath, _, err := face.ToPath(lines[0])
	if err != nil {
		log.Fatalf("Failed to convert text to path: %s", err)
	}
	textBounds := textPath.FastBounds()
	scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/textBounds.H())
	textPath = textPath.Scale(scale, scale)

	// Check for descenders (characters that go below baseline)
	hasDescenders := false
	for _, char := range lines[0] {
		switch char {
		case 'g', 'j', 'p', 'q', 'y':
			hasDescenders = true
		}
	}

	// recalculate text position after scaling
	textBounds = textPath.FastBounds()
	// Add a small offset to shift text left
	xOffset := 0.3
	yOffset := 0.05
	if hasDescenders {
		yOffset = 1.0
	}
	x := containerX + (containerBounds.W()-textBounds.W())/2 - xOffset*scale
	y := containerY + (containerBounds.H()-textBounds.H())/2 + yOffset*scale
	textPath = textPath.Translate(x, y)
	textPath = textPath.Scale(1, -1)
	textPath = textPath.Translate(0, height)

	builder.AddPath(textPath.ToSVG(), map[string]string{
		"fill":         backgroundColor,
		"id":           "text-line-path",
		"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
		"stroke-width": "0.025",
	})
}
func drawExtraSmallVerticalRibbon(
	builder *svgutils.SVGBuilder,
	width float64,
	height float64,
	foregroundColor string,
	backgroundColor string,
	strokeOnly bool,
) {
	// add the outer edge
	outerEdge := canvas.RoundedRectangle(width, height, 0.25)
	builder.StartGroup("Outside", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill":         backgroundColor,
		"id":           "outside-path",
		"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
		"stroke-width": "0.025",
	})
	builder.EndGroup()

	// add the rounded edge
	roundedEdge := canvas.RoundedRectangle(width-0.5, height-0.5, 0.2)
	roundedEdge = roundedEdge.Translate(0.25, 0.25)
	builder.StartGroup("Round", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill":         foregroundColor,
		"id":           "round-path",
		"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
		"stroke-width": "0.025",
	})
	builder.EndGroup()

	// add the border
	// borderOuter := canvas.RoundedRectangle(width-1.0, height-1.0, 0.15)
	// borderOuter = borderOuter.Translate(0.5, 0.5)
	builder.StartGroup("Vcarve", map[string]string{})
	// builder.AddPath(borderOuter.ToSVG(), map[string]string{
	// 	"fill":         backgroundColor,
	// 	"id":           "border-outer-path",
	// 	"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
	// 	"stroke-width": "0.025",
	// })
	// borderInner := canvas.RoundedRectangle(width-1.2, height-1.2, 0.1)
	// borderInner = borderInner.Translate(0.6, 0.6)
	// builder.AddPath(borderInner.ToSVG(), map[string]string{
	// 	"fill":         foregroundColor,
	// 	"id":           "border-inner-path",
	// 	"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
	// 	"stroke-width": "0.025",
	// })
	builder.EndGroup()
}
func drawExtraSmallVerticalRibbonText(
	builder *svgutils.SVGBuilder,
	width float64,
	height float64,
	backgroundColor string,
	lines []string,
	fontFamily *canvas.FontFamily,
	strokeOnly bool,
	ctx *canvas.Context,
) {
	if len(lines) == 0 {
		return
	}

	// draw text container
	container := canvas.Rectangle(width-1.25, height-1.5)
	containerX := 0.625
	containerY := 0.75
	container = container.Translate(containerX, containerY)
	container = container.Scale(1, -1)
	container = container.Translate(0, height)
	containerBounds := container.FastBounds()
	builder.AddPath(container.ToSVG(), map[string]string{
		"fill":         "none",
		"stroke":       "pink",
		"stroke-width": "0.025",
		"id":           "text-container-path",
	})

	// draw text
	fontSize := 20.0
	face := fontFamily.Face(fontSize, canvas.FontRegular, canvas.FontNormal)

	// Convert each character to a path and stack them vertically
	chars := []rune(lines[0])

	// Create a group for all characters
	builder.StartGroup("TextLayer", map[string]string{
		"id": "text-layer",
	})

	// First pass - measure total height needed
	totalHeight := 0.0
	charPaths := make([]*canvas.Path, len(chars))
	charBounds := make([]canvas.Rect, len(chars))

	for i, char := range chars {
		if string(char) == "" {
			return
		}
		textPath, _, err := face.ToPath(string(char))
		if err != nil {
			log.Fatalf("Failed to convert character to path: %s", err)
		}
		charPaths[i] = textPath
		charBounds[i] = textPath.FastBounds()
		totalHeight += charBounds[i].H()
	}

	// Calculate spacing between characters
	spacing := charBounds[0].H() * 0.2 // 20% of character height
	totalHeight += spacing * float64(len(chars)-1)

	// Calculate scale to fit container width
	maxCharWidth := 0.0
	for _, bounds := range charBounds {
		if bounds.W() > maxCharWidth {
			maxCharWidth = bounds.W()
		}
	}
	scale := containerBounds.W() / (maxCharWidth * 1.5) // 1.5 gives some padding

	// Calculate starting Y position to center the text stack
	currentY := containerY + (containerBounds.H()-totalHeight*scale)/2

	// Second pass - render each character
	for i, charPath := range charPaths {
		// Scale the character
		scaledPath := charPath.Scale(scale, scale)
		scaledBounds := scaledPath.FastBounds()

		// Center horizontally and position vertically
		x := containerX + (containerBounds.W()-scaledBounds.W())/2
		// Change the y calculation to go from top to bottom
		y := currentY + float64(len(chars)-1-i)*(scaledBounds.H()+spacing*scale)

		// Apply transformations
		scaledPath = scaledPath.Translate(x, y)
		scaledPath = scaledPath.Scale(1, -1)
		scaledPath = scaledPath.Translate(0, height)
		// Rotate 90 degrees around the character's center point
		centerX := scaledPath.FastBounds().X0 + scaledPath.FastBounds().W()/2
		centerY := scaledPath.FastBounds().Y0 + scaledPath.FastBounds().H()/2
		scaledPath = scaledPath.Translate(-centerX, -centerY)
		scaledPath = scaledPath.Transform(canvas.Identity.Rotate(90))
		scaledPath = scaledPath.Translate(centerX, centerY)

		builder.AddPath(scaledPath.ToSVG(), map[string]string{
			"fill":         backgroundColor,
			"id":           fmt.Sprintf("text-char-%d", i),
			"stroke":       map[bool]string{true: "black", false: "none"}[strokeOnly],
			"stroke-width": "0.025",
		})
	}

	builder.EndGroup()
}

func drawRibbon(
	size string,
	width float64,
	height float64,
	foregroundColor string,
	backgroundColor string,
	lines []string,
	fontFamily *canvas.FontFamily,
	strokeOnly bool,
	ctx *canvas.Context,
) string {
	if strokeOnly {
		foregroundColor = "transparent"
		backgroundColor = "transparent"
	} else {
		foregroundColor = GetColor(foregroundColor)
		backgroundColor = GetColor(backgroundColor)
	}

	if fontFamily == nil {
		log.Fatalf("Error: Font family is required but was not provided")
	}

	// Initialize SVG builder
	builder := svgutils.NewSVGBuilder(width, height)

	if size == "extra small" {
		drawExtraSmallRibbon(builder, width, height, foregroundColor, backgroundColor, strokeOnly)
		drawExtraSmallRibbonText(builder, width, height, backgroundColor, lines, fontFamily, strokeOnly)
	} else if size == "extra small vertical" {
		drawExtraSmallVerticalRibbon(builder, width, height, foregroundColor, backgroundColor, strokeOnly)
		drawExtraSmallVerticalRibbonText(builder, width, height, backgroundColor, lines, fontFamily, strokeOnly, ctx)
	} else {
		log.Fatalf("Invalid size: %s", size)
	}

	// builder.EndGroup()

	return builder.Close()
}
