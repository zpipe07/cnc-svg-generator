package signs

import (
	"cnc-svg-generator/pkg/svgutils"
	"fmt"
	"log"

	"github.com/tdewolff/canvas"
)

func drawLargeSesame(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// Add the outer edge
	outerEdge, err := canvas.ParseSVGPath("M 0 10.75 A 0.25 0.25 0 0 0 0.25 11 L 7.496 11 A 8.039 8.039 0 0 0 7.5 11 A 7.696 7.696 0 0 0 7.582 11 L 14.75 11 A 0.25 0.25 0 0 0 15 10.75 L 15 3 A 0.25 0.25 0 0 0 14.75 2.75 L 13.408 2.75 A 0.25 0.25 0 0 1 13.207 2.649 A 4.879 4.879 0 0 0 12.919 2.296 A 5.937 5.937 0 0 0 12.215 1.612 A 0.25 0.25 0 0 0 12.213 1.61 A 7.213 7.213 0 0 0 8.747 0.097 A 8.039 8.039 0 0 0 7.5 0 A 7.696 7.696 0 0 0 3.742 0.957 A 6.554 6.554 0 0 0 2.786 1.611 A 5.524 5.524 0 0 0 1.801 2.647 A 0.25 0.25 0 0 1 1.599 2.75 L 0.25 2.75 A 0.25 0.25 0 0 0 0 3 L 0 10.75 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "outer-edge",
	})
	builder.EndGroup()

	// Add the rounded edge
	roundedEdge, err := canvas.ParseSVGPath("M 0.2 10.75 A 0.05 0.05 0 0 0 0.25 10.8 L 7.496 10.8 A 0.2 0.2 0 0 1 7.498 10.8 A 0.2 0.2 0 0 1 7.5 10.8 A 0.2 0.2 0 0 1 7.501 10.8 A 7.496 7.496 0 0 0 7.581 10.8 A 0.2 0.2 0 0 1 7.582 10.8 L 14.75 10.8 A 0.05 0.05 0 0 0 14.8 10.75 L 14.8 3 A 0.05 0.05 0 0 0 14.75 2.95 L 13.408 2.95 A 0.45 0.45 0 0 1 13.046 2.768 A 4.679 4.679 0 0 0 12.77 2.43 A 0.2 0.2 0 0 1 12.768 2.428 A 5.737 5.737 0 0 0 12.088 1.766 A 7.013 7.013 0 0 0 8.718 0.295 A 0.2 0.2 0 0 1 8.716 0.295 A 7.839 7.839 0 0 0 7.5 0.2 A 0.2 0.2 0 0 1 7.499 0.2 A 7.496 7.496 0 0 0 3.84 1.132 A 6.354 6.354 0 0 0 2.913 1.765 A 0.2 0.2 0 0 1 2.912 1.766 A 5.324 5.324 0 0 0 1.963 2.765 A 0.45 0.45 0 0 1 1.599 2.95 L 0.25 2.95 A 0.05 0.05 0 0 0 0.2 3 L 0.2 10.75 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.StartGroup("Rounded Edge", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
	builder.EndGroup()

	// Add the border
	borderOuter, err := canvas.ParseSVGPath("M 0.5 10.5 L 7.496 10.5 A 0.3 0.3 0 0 1 7.499 10.5 A 0.3 0.3 0 0 1 7.503 10.5 A 7.196 7.196 0 0 0 7.579 10.5 A 0.3 0.3 0 0 1 7.58 10.5 A 0.3 0.3 0 0 1 7.582 10.5 L 14.5 10.5 L 14.5 3.25 L 13.408 3.25 A 0.75 0.75 0 0 1 12.805 2.947 A 4.379 4.379 0 0 0 12.547 2.63 A 0.3 0.3 0 0 1 12.544 2.627 A 0.3 0.3 0 0 1 12.542 2.625 A 5.437 5.437 0 0 0 11.899 1.999 A 6.713 6.713 0 0 0 8.674 0.592 A 0.3 0.3 0 0 1 8.672 0.591 A 0.3 0.3 0 0 1 8.669 0.591 A 7.539 7.539 0 0 0 7.5 0.5 A 0.3 0.3 0 0 1 7.499 0.5 A 0.3 0.3 0 0 1 7.497 0.5 A 7.196 7.196 0 0 0 3.986 1.394 A 6.054 6.054 0 0 0 3.104 1.997 A 0.3 0.3 0 0 1 3.102 1.998 A 0.3 0.3 0 0 1 3.101 1.999 A 5.024 5.024 0 0 0 2.205 2.942 A 0.75 0.75 0 0 1 1.599 3.25 L 0.5 3.25 L 0.5 10.5 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	borderInner, err := canvas.ParseSVGPath("M 0.7 10.3 L 7.496 10.3 A 0.5 0.5 0 0 1 7.5 10.3 A 0.2 0.2 0 0 1 7.504 10.3 A 6.996 6.996 0 0 0 7.578 10.3 A 0.2 0.2 0 0 1 7.58 10.3 A 0.2 0.2 0 0 1 7.582 10.3 L 14.3 10.3 L 14.3 3.45 L 13.408 3.45 A 0.95 0.95 0 0 1 12.645 3.066 A 4.179 4.179 0 0 0 12.398 2.764 A 0.2 0.2 0 0 1 12.394 2.76 A 0.2 0.2 0 0 1 12.391 2.756 A 5.237 5.237 0 0 0 11.772 2.154 A 6.513 6.513 0 0 0 8.645 0.79 A 0.2 0.2 0 0 1 8.641 0.789 A 0.2 0.2 0 0 1 8.638 0.789 A 7.339 7.339 0 0 0 7.5 0.7 A 0.2 0.2 0 0 1 7.498 0.7 A 0.2 0.2 0 0 1 7.496 0.7 A 6.996 6.996 0 0 0 4.083 1.569 A 5.854 5.854 0 0 0 3.231 2.152 A 0.2 0.2 0 0 1 3.229 2.153 A 0.2 0.2 0 0 1 3.227 2.155 A 4.824 4.824 0 0 0 2.367 3.06 A 0.95 0.95 0 0 1 1.599 3.45 L 0.7 3.45 L 0.7 10.3 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
}

func drawMediumSesame(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// Add the outer edge
	outerEdge, err := canvas.ParseSVGPath("M 0 8.75 A 0.25 0.25 0 0 0 0.25 9 L 7.465 9 A 9.558 9.558 0 0 0 7.5 9 A 9.329 9.329 0 0 0 7.578 9 L 14.75 9 A 0.25 0.25 0 0 0 15 8.75 L 15 2.5 A 0.25 0.25 0 0 0 14.75 2.25 L 13.385 2.25 A 0.25 0.25 0 0 1 13.199 2.167 A 4.961 4.961 0 0 0 12.369 1.426 A 6.023 6.023 0 0 0 12.214 1.318 C 10.964 0.474 9.268 0 7.5 0 A 9.329 9.329 0 0 0 4.361 0.53 A 7.143 7.143 0 0 0 2.786 1.318 A 5.193 5.193 0 0 0 1.8 2.168 A 0.25 0.25 0 0 1 1.615 2.25 L 0.25 2.25 A 0.25 0.25 0 0 0 0 2.5 L 0 8.75 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "outer-edge",
	})
	builder.EndGroup()

	// Add the rounded edge
	roundedEdge, err := canvas.ParseSVGPath("M 0.2 8.75 A 0.05 0.05 0 0 0 0.25 8.8 L 7.465 8.8 A 9.358 9.358 0 0 0 7.5 8.8 A 9.129 9.129 0 0 0 7.577 8.8 L 14.75 8.8 A 0.05 0.05 0 0 0 14.8 8.75 L 14.8 2.5 A 0.05 0.05 0 0 0 14.75 2.45 L 13.385 2.45 A 0.45 0.45 0 0 1 13.05 2.3 A 4.761 4.761 0 0 0 12.253 1.589 A 0.2 0.2 0 0 1 12.252 1.588 A 5.823 5.823 0 0 0 12.103 1.485 A 7.338 7.338 0 0 0 9.861 0.498 A 0.2 0.2 0 0 1 9.86 0.498 A 9.358 9.358 0 0 0 7.5 0.2 A 0.2 0.2 0 0 1 7.499 0.2 A 9.129 9.129 0 0 0 4.427 0.719 A 6.943 6.943 0 0 0 2.897 1.485 A 4.993 4.993 0 0 0 1.949 2.301 A 0.45 0.45 0 0 1 1.615 2.45 L 0.25 2.45 A 0.05 0.05 0 0 0 0.2 2.5 L 0.2 8.75 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.StartGroup("Rounded Edge", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
	builder.EndGroup()

	// Add the border
	borderOuter, err := canvas.ParseSVGPath("M 0.5 8.5 L 7.465 8.5 A 9.058 9.058 0 0 0 7.499 8.5 A 0.3 0.3 0 0 1 7.501 8.5 A 0.3 0.3 0 0 1 7.502 8.5 A 8.829 8.829 0 0 0 7.576 8.5 A 0.3 0.3 0 0 1 7.577 8.5 A 0.3 0.3 0 0 1 7.578 8.5 L 14.5 8.5 L 14.5 2.75 L 13.385 2.75 A 0.75 0.75 0 0 1 12.826 2.5 A 4.461 4.461 0 0 0 12.08 1.834 A 0.3 0.3 0 0 1 12.078 1.833 A 0.3 0.3 0 0 1 12.076 1.832 A 5.523 5.523 0 0 0 11.937 1.734 A 7.038 7.038 0 0 0 9.788 0.789 A 0.3 0.3 0 0 1 9.786 0.789 A 0.3 0.3 0 0 1 9.785 0.788 A 9.058 9.058 0 0 0 7.501 0.5 A 0.3 0.3 0 0 1 7.499 0.5 A 0.3 0.3 0 0 1 7.498 0.5 A 8.829 8.829 0 0 0 4.527 1.002 A 6.643 6.643 0 0 0 3.063 1.734 A 4.693 4.693 0 0 0 2.172 2.502 A 0.75 0.75 0 0 1 1.615 2.75 L 0.5 2.75 L 0.5 8.5 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	borderInner, err := canvas.ParseSVGPath("M 0.7 8.3 L 7.465 8.3 A 8.858 8.858 0 0 0 7.499 8.3 A 0.2 0.2 0 0 1 7.501 8.3 A 0.2 0.2 0 0 1 7.503 8.3 A 8.629 8.629 0 0 0 7.575 8.3 A 0.2 0.2 0 0 1 7.577 8.3 A 0.2 0.2 0 0 1 7.578 8.3 L 14.3 8.3 L 14.3 2.95 L 13.385 2.95 A 0.95 0.95 0 0 1 12.677 2.634 A 4.261 4.261 0 0 0 11.964 1.997 A 0.2 0.2 0 0 1 11.962 1.996 A 0.2 0.2 0 0 1 11.959 1.994 A 5.323 5.323 0 0 0 11.826 1.901 A 6.838 6.838 0 0 0 9.739 0.983 A 0.2 0.2 0 0 1 9.737 0.983 A 0.2 0.2 0 0 1 9.735 0.982 A 8.858 8.858 0 0 0 7.501 0.7 A 0.2 0.2 0 0 1 7.499 0.7 A 0.2 0.2 0 0 1 7.497 0.7 A 8.629 8.629 0 0 0 4.594 1.19 A 6.443 6.443 0 0 0 3.174 1.901 A 4.493 4.493 0 0 0 2.32 2.636 A 0.95 0.95 0 0 1 1.615 2.95 L 0.7 2.95 L 0.7 8.3 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
}

func drawSmallSesame(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// Add the outer edge
	outerEdge, err := canvas.ParseSVGPath("M 0 7.25 A 0.25 0.25 0 0 0 0.25 7.5 L 7.428 7.5 A 11.177 11.177 0 0 0 7.5 7.5 A 11.177 11.177 0 0 0 7.572 7.5 L 14.75 7.5 A 0.25 0.25 0 0 0 15 7.25 L 15 2.125 A 0.25 0.25 0 0 0 14.75 1.875 L 13.375 1.875 A 0.25 0.25 0 0 1 13.204 1.808 A 3.902 3.902 0 0 0 12.883 1.538 A 5.582 5.582 0 0 0 12.215 1.099 A 0.25 0.25 0 0 0 12.213 1.098 A 8.024 8.024 0 0 0 10.349 0.36 A 11.177 11.177 0 0 0 7.5 0 A 11.177 11.177 0 0 0 4.651 0.36 A 8.024 8.024 0 0 0 2.786 1.098 A 5.211 5.211 0 0 0 1.821 1.786 A 3.34 3.34 0 0 0 1.725 1.875 L 0.25 1.875 A 0.25 0.25 0 0 0 0 2.125 L 0 7.25 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "outer-edge",
	})
	builder.EndGroup()

	// Add the rounded edge
	roundedEdge, err := canvas.ParseSVGPath("M 0.2 7.25 A 0.05 0.05 0 0 0 0.25 7.3 L 7.428 7.3 A 10.977 10.977 0 0 0 7.499 7.3 A 10.977 10.977 0 0 0 7.572 7.3 L 14.75 7.3 A 0.05 0.05 0 0 0 14.8 7.25 L 14.8 2.125 A 0.05 0.05 0 0 0 14.75 2.075 L 13.375 2.075 A 0.45 0.45 0 0 1 13.067 1.953 A 3.702 3.702 0 0 0 12.763 1.698 A 0.2 0.2 0 0 1 12.761 1.697 A 5.382 5.382 0 0 0 12.117 1.273 A 7.824 7.824 0 0 0 10.299 0.553 A 10.977 10.977 0 0 0 7.501 0.2 A 0.2 0.2 0 0 1 7.499 0.2 A 10.977 10.977 0 0 0 4.702 0.553 A 7.824 7.824 0 0 0 2.883 1.274 A 5.011 5.011 0 0 0 1.955 1.934 A 0.2 0.2 0 0 1 1.954 1.935 A 3.14 3.14 0 0 0 1.864 2.019 A 0.2 0.2 0 0 1 1.725 2.075 L 0.25 2.075 A 0.05 0.05 0 0 0 0.2 2.125 L 0.2 7.25 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.StartGroup("Rounded Edge", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
	builder.EndGroup()

	// Add the border
	borderOuter, err := canvas.ParseSVGPath("M 0.5 7 L 7.428 7 A 10.677 10.677 0 0 0 7.498 7 A 0.3 0.3 0 0 1 7.5 7 A 0.3 0.3 0 0 1 7.502 7 A 10.677 10.677 0 0 0 7.571 7 L 14.5 7 L 14.5 2.375 L 13.375 2.375 A 0.75 0.75 0 0 1 12.862 2.172 A 3.402 3.402 0 0 0 12.582 1.937 A 0.3 0.3 0 0 1 12.581 1.936 A 0.3 0.3 0 0 1 12.579 1.935 A 5.082 5.082 0 0 0 11.973 1.536 A 7.524 7.524 0 0 0 10.224 0.844 A 10.677 10.677 0 0 0 7.502 0.5 A 0.3 0.3 0 0 1 7.5 0.5 A 0.3 0.3 0 0 1 7.498 0.5 A 10.677 10.677 0 0 0 4.777 0.844 A 7.524 7.524 0 0 0 3.028 1.536 A 4.711 4.711 0 0 0 2.156 2.157 A 0.3 0.3 0 0 1 2.155 2.158 A 0.3 0.3 0 0 1 2.154 2.159 A 2.84 2.84 0 0 0 2.072 2.235 A 0.5 0.5 0 0 1 1.725 2.375 L 0.5 2.375 L 0.5 7 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	borderInner, err := canvas.ParseSVGPath("M 0.7 6.8 L 7.428 6.8 A 0.2 0.2 0 0 1 7.429 6.8 A 0.2 0.2 0 0 1 7.43 6.8 A 10.477 10.477 0 0 0 7.498 6.8 A 0.2 0.2 0 0 1 7.5 6.8 A 0.2 0.2 0 0 1 7.502 6.8 A 10.477 10.477 0 0 0 7.57 6.8 A 0.2 0.2 0 0 1 7.571 6.8 A 0.2 0.2 0 0 1 7.572 6.8 L 14.3 6.8 L 14.3 2.575 L 13.375 2.575 A 0.95 0.95 0 0 1 12.725 2.318 A 3.202 3.202 0 0 0 12.462 2.097 A 0.2 0.2 0 0 1 12.46 2.095 A 0.2 0.2 0 0 1 12.458 2.094 A 4.882 4.882 0 0 0 11.875 1.711 A 7.324 7.324 0 0 0 10.174 1.038 A 10.477 10.477 0 0 0 7.502 0.7 A 0.2 0.2 0 0 1 7.5 0.7 A 0.2 0.2 0 0 1 7.498 0.7 A 10.477 10.477 0 0 0 4.828 1.037 A 7.324 7.324 0 0 0 3.124 1.711 A 4.511 4.511 0 0 0 2.29 2.306 A 0.2 0.2 0 0 1 2.288 2.307 A 0.2 0.2 0 0 1 2.287 2.308 A 2.64 2.64 0 0 0 2.211 2.379 A 0.7 0.7 0 0 1 1.725 2.575 L 0.7 2.575 L 0.7 6.8 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
}

func drawSesame(
	size string,
	width float64,
	height float64,
	foregroundColor string,
	backgroundColor string,
	lines []string,
	fontFamily *canvas.FontFamily,
) string {
	// Initialize SVG builder
	builder := svgutils.NewSVGBuilder(width, height)

	if size == "large" {
		drawLargeSesame(builder, foregroundColor, backgroundColor)
	} else if size == "medium" {
		drawMediumSesame(builder, foregroundColor, backgroundColor)
	} else if size == "small" {
		drawSmallSesame(builder, foregroundColor, backgroundColor)
	} else {
		log.Printf("Invalid size: %s", size)
	}

	// draw text
	numLines := len(lines)

	if numLines > 0 {
		// Calculate the total height for the text containers
		// const padding = 1.0 // Padding from top and bottom
		const topPadding = 1.75
		const bottomPadding = 1.0
		lineSpacing := 0.5 // Spacing between lines
		availableHeight := height - (topPadding + bottomPadding) - float64(numLines-1)*lineSpacing

		// Determine the dimensions for each line
		type ContainerDimensions struct {
			Width  float64
			Height float64
		}
		var containerDimensions []ContainerDimensions
		switch numLines {
		case 3:
			containerDimensions = []ContainerDimensions{
				{
					Height: availableHeight * 0.2,
					Width:  (width - 0.5) / 2,
				},
				{
					Height: availableHeight * 0.5,
					Width:  width - 2.0,
				},
				{
					Height: availableHeight * 0.3,
					Width:  width - 2.0,
				},
			}
		case 2:
			containerDimensions = []ContainerDimensions{
				{
					Height: availableHeight * 0.5,
					Width:  (width - 0.5) / 2,
				},
				{
					Height: availableHeight * 0.5,
					Width:  width - 2.0,
				},
			}
		default:
			containerDimensions = []ContainerDimensions{
				{
					Height: availableHeight,
					Width:  width * 3.0 / 5.0,
				},
			}
		}

		currentY := height - topPadding

		for i, line := range lines {
			containerDimensions := containerDimensions[i]
			container := canvas.Rectangle(
				containerDimensions.Width,
				containerDimensions.Height,
			)
			containerBounds := container.Bounds()
			containerX := width/2 - containerBounds.W()/2
			containerY := currentY - containerDimensions.Height

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
			currentY -= containerDimensions.Height + lineSpacing
		}
	}

	builder.EndGroup()

	// Finalize and return the SVG
	return builder.Close()
}
