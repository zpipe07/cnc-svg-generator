package signs

import (
	"cnc-svg-generator/pkg/fonts"
	"cnc-svg-generator/pkg/svgutils"
	"fmt"
	"log"

	"github.com/tdewolff/canvas"
)

// Helper function to calculate the minimum of two values
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func drawMediumDeco(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// Add the outer edge
	outerEdge, err := canvas.ParseSVGPath("M 0.634 8.942 L 0.634 8.682 C 0.634 8.645 0.62 8.61 0.593 8.584 L 0.394 8.392 C 0.368 8.368 0.333 8.354 0.297 8.354 L 0.06 8.354 C 0.027 8.354 0 8.328 0 8.296 L 0 1.31 A 0.059 0.059 0 0 1 0.06 1.252 L 0.294 1.252 A 0.144 0.144 0 0 0 0.297 1.252 A 0.142 0.142 0 0 0 0.395 1.212 L 0.594 1.021 C 0.62 0.996 0.634 0.961 0.634 0.925 L 0.634 0.664 C 0.634 0.631 0.661 0.605 0.694 0.605 L 5.889 0.605 C 5.926 0.606 5.962 0.591 5.989 0.565 C 6.075 0.481 6.205 0.449 6.335 0.469 A 0.504 0.504 0 0 1 6.375 0.477 C 6.433 0.493 6.495 0.471 6.528 0.421 A 0.471 0.471 0 0 1 6.925 0.232 A 0.614 0.614 0 0 1 7.001 0.237 A 0.14 0.14 0 0 0 7.022 0.238 A 0.14 0.14 0 0 0 7.135 0.181 A 0.439 0.439 0 0 1 7.181 0.127 A 0.451 0.451 0 0 1 7.511 0 A 0.451 0.451 0 0 1 7.842 0.127 A 0.439 0.439 0 0 1 7.888 0.181 A 0.14 0.14 0 0 0 8 0.238 A 0.14 0.14 0 0 0 8.019 0.237 A 0.615 0.615 0 0 1 8.098 0.232 A 0.47 0.47 0 0 1 8.496 0.424 A 0.14 0.14 0 0 0 8.61 0.482 C 8.615 0.482 8.621 0.482 8.627 0.481 A 0.14 0.14 0 0 0 8.644 0.478 C 8.683 0.468 8.723 0.463 8.762 0.463 A 0.385 0.385 0 0 1 9.031 0.563 C 9.058 0.59 9.094 0.605 9.132 0.605 L 14.306 0.605 A 0.059 0.059 0 0 1 14.366 0.664 L 14.366 0.924 A 0.131 0.131 0 0 0 14.366 0.925 A 0.134 0.134 0 0 0 14.407 1.022 L 14.603 1.211 C 14.629 1.237 14.665 1.252 14.703 1.252 L 14.94 1.252 C 14.973 1.252 15 1.278 15 1.31 L 15 8.295 C 15 8.328 14.973 8.354 14.94 8.354 L 14.705 8.354 A 0.144 0.144 0 0 0 14.703 8.354 A 0.142 0.142 0 0 0 14.605 8.393 L 14.408 8.583 A 0.134 0.134 0 0 0 14.366 8.68 L 14.366 8.942 C 14.366 8.974 14.339 9 14.306 9 L 0.694 9 A 0.059 0.059 0 0 1 0.634 8.942 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	outerBounds := outerEdge.Bounds()
	outerEdge = outerEdge.Translate(-outerBounds.X0, -outerBounds.Y0)
	builder.StartGroup("Outside", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "outer-edge",
	})
	builder.EndGroup()

	// Add the border
	borderOuter, err := canvas.ParseSVGPath("M 0.813 8.551 L 0.428 8.18 L 0.179 8.18 L 0.179 5.276 L 0.179 4.482 L 0.179 1.426 L 0.428 1.426 L 0.813 1.055 L 0.813 0.779 L 6.038 0.779 A 0.219 0.219 0 0 1 6.077 0.731 A 0.339 0.339 0 0 1 6.297 0.647 L 7.245 1.598 C 7.249 1.594 7.254 1.59 7.259 1.586 L 6.607 0.681 C 6.71 0.445 6.883 0.407 6.886 0.405 L 7.314 1.547 C 7.328 1.538 7.343 1.53 7.359 1.524 L 7.297 0.261 A 0.174 0.174 0 0 1 7.343 0.219 A 0.298 0.298 0 0 1 7.508 0.177 A 0.298 0.298 0 0 1 7.674 0.219 A 0.174 0.174 0 0 1 7.72 0.261 L 7.658 1.524 C 7.674 1.53 7.689 1.538 7.703 1.547 L 8.131 0.405 A 0.405 0.405 0 0 1 8.308 0.521 A 0.57 0.57 0 0 1 8.41 0.681 L 7.758 1.586 C 7.763 1.59 7.767 1.594 7.772 1.598 L 8.72 0.647 A 0.339 0.339 0 0 1 8.94 0.731 A 0.219 0.219 0 0 1 8.978 0.779 L 14.187 0.779 L 14.187 1.055 L 14.572 1.426 L 14.821 1.426 L 14.821 4.482 L 14.821 5.276 L 14.821 8.18 L 14.572 8.18 L 14.187 8.551 L 14.187 8.824 L 7.5 8.822 L 0.813 8.824 L 0.813 8.551 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	borderOuter = borderOuter.Translate(-outerBounds.X0, -outerBounds.Y0)
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "my-custom-id",
	})
	borderInner, err := canvas.ParseSVGPath("M 7.5 8.619 L 13.987 8.624 L 13.987 8.468 L 14.489 7.983 L 14.621 7.983 L 14.621 5.273 L 14.621 4.479 L 14.621 1.617 L 14.489 1.617 L 13.987 1.133 L 13.987 0.976 L 8.758 0.976 L 7.657 1.736 L 7.647 1.736 L 7.631 1.736 L 7.534 1.736 L 7.483 1.736 L 7.386 1.736 L 7.369 1.736 L 7.36 1.736 L 6.259 0.976 L 1.013 0.976 L 1.013 1.133 L 0.511 1.617 L 0.379 1.617 L 0.379 4.479 L 0.379 5.273 L 0.379 7.983 L 0.511 7.983 L 1.013 8.468 L 1.013 8.624 L 7.5 8.619 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	borderInner = borderInner.Translate(-outerBounds.X0, -outerBounds.Y0)
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "my-custom-id",
	})

	// decorative lines
	paths := []string{
		"M 1.515 6.34 L 13.518 6.34 L 13.518 6.428 L 1.515 6.428 L 1.515 6.34 Z M 5.865 1.278 L 1.482 1.278 L 1.482 1.172 L 5.865 1.172 L 5.865 1.278 Z M 13.469 1.278 L 9.086 1.278 L 9.086 1.172 L 13.469 1.172 L 13.469 1.278 Z M 5.865 1.481 L 1.482 1.481 L 1.482 1.375 L 5.865 1.375 L 5.865 1.481 Z M 13.469 1.481 L 9.086 1.481 L 9.086 1.375 L 13.469 1.375 L 13.469 1.481 Z",
		// "M 6.11489,21.705244 H 1.732577 V 21.60063 H 6.11489 v 0.104614",
		// "M 6.11489,21.905256 H 1.732577 V 21.800643 H 6.11489 v 0.104613",
		// "m 2.275541,24.268541 h 10.98395 v 0.07983 H 2.275541 v -0.07983",
		// "m 1.769095,28.64164 h 11.996871 v 0.08721 H 1.769095 v -0.08721",
		// "M 13.717872,21.705244 H 9.335559 V 21.60063 h 4.382313 v 0.104614",
		// "M 13.717872,21.905256 H 9.335559 v -0.104613 h 4.382313 v 0.104613",
	}
	for _, pathStr := range paths {
		path, err := canvas.ParseSVGPath(pathStr)
		if err != nil {
			log.Printf("Failed to parse SVG path: %s", err)
			continue
		}
		path = path.Translate(-outerBounds.X0, -outerBounds.Y0)
		builder.AddPath(path.ToSVG(), map[string]string{
			"fill": backgroundColor,
			"id":   "my-custom-id",
		})
	}
}

func drawLargeDeco(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// Add the outer edge
	outerEdge, err := canvas.ParseSVGPath("M 0.908768,31.446765 C 0.875547,31.44691 0.848537,31.420418 0.848537,31.387689 V 31.12425 C 0.849094,31.08683 0.834179,31.05079 0.807221,31.02442 L 0.610317,30.831796 C 0.584074,30.805915 0.548463,30.791372 0.51133,30.791372 H 0.274488 c -0.033117,0 -0.059964,-0.02645 -0.059964,-0.05908 v -8.974773 c -1.47e-4,-0.03273 0.026742,-0.05934 0.059964,-0.05934 h 0.234269 c 0.037815,6.78e-4 0.074297,-0.01375 0.101135,-0.04001 L 0.808308,21.46406 c 0.025779,-0.02579 0.040229,-0.06053 0.040229,-0.09672 v -0.265472 c 0,-0.01589 0.006366,-0.03031 0.0167218,-0.04093 0.0109112,-0.01119 0.0262519,-0.01815 0.0536598,-0.01815 H 6.103881 c 0.037211,4.2e-4 0.073057,-0.01379 0.099612,-0.03947 0.094778,-0.09152 0.2440221,-0.12094 0.387435,-0.08662 0.057473,0.01533 0.118492,-0.0069 0.1520071,-0.05539 0.095811,-0.138195 0.280362,-0.206655 0.476257,-0.181385 0.050567,0.0063 0.100591,-0.01506 0.130568,-0.05567 0.054649,-0.07392 0.167241,-0.17749 0.376076,-0.17749 0.207887,0 0.320405,0.102622 0.375335,0.176472 0.029978,0.04144 0.080717,0.06321 0.131936,0.05661 0.195214,-0.02499 0.37908,0.04309 0.474978,0.180508 0.03283,0.04829 0.092831,0.07103 0.150035,0.05699 0.143113,-0.03505 0.292412,-0.0065 0.387872,0.08389 0.026323,0.02658 0.062421,0.04156 0.100112,0.04156 h 5.1636789 c 0.04367,-1.89e-4 0.07073,0.02632 0.07073,0.05907 v 0.263422 c -5.61e-4,0.03743 0.01436,0.07347 0.04132,0.09985 l 0.197806,0.193493 c 0.02617,0.02535 0.0614,0.03955 0.09809,0.03955 h 0.236831 c 0.03312,0 0.05997,0.02645 0.05997,0.05908 v 8.974824 c 1.18e-4,0.03271 -0.02676,0.05929 -0.05997,0.05929 h -0.234426 c -0.03776,-6.33e-4 -0.07417,0.0138 -0.100966,0.04001 l -0.198489,0.19416 c -0.02574,0.02578 -0.04017,0.0605 -0.04017,0.09666 v 0.265484 c 0,0.03263 -0.02685,0.05908 -0.07038,0.05908 l -13.601357,0")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	outerBounds := outerEdge.Bounds()
	outerEdge = outerEdge.Translate(-outerBounds.X0, -outerBounds.Y0)
	builder.StartGroup("Outside", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "outer-edge",
	})
	builder.EndGroup()

	// Add the border
	borderOuter, err := canvas.ParseSVGPath("m 14.40157,21.21908 v 0.279053 l 0.384801,0.376435 0.24925,-4.7e-5 -5e-5,2.868355 v 3.016954 2.855247 h -0.249268 l -0.384733,0.37638 v 0.279011 H 7.714521 1.027485 V 30.991457 L 0.64274,30.615077 H 0.393471 V 27.75983 24.742876 l -9.3e-4,-2.868355 0.248382,4.7e-5 0.3848,-0.376435 v -0.279055 l 5.226191,-2e-6 c 0.026775,-0.04548 0.094874,-0.118972 0.258125,-0.130022 l 0.94862,0.936577 c 0.00438,-0.0043 0.00896,-0.0083 0.013705,-0.0122 l -0.6523,-0.891281 c 0.103078,-0.232353 0.276781,-0.269653 0.279571,-0.272353 l 0.427643,1.12502 c 0.014554,-0.0084 0.029644,-0.01592 0.045167,-0.02232 L 7.510393,20.708228 c 0.00212,-8.3e-4 0.050761,-0.08309 0.211624,-0.08309 0.160878,0 0.209505,0.08226 0.211625,0.08309 l -0.061971,1.244211 c 0.01571,0.0065 0.030827,0.01406 0.045199,0.02246 l 0.427545,-1.1251 c 0.0028,0.0027 0.176478,0.04 0.27957,0.272353 l -0.652243,0.891193 c 0.00469,0.004 0.00928,0.008 0.013763,0.01218 l 0.948492,-0.936467 c 0.163264,0.01105 0.231349,0.08454 0.258138,0.130022 h 5.209435")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	borderOuter = borderOuter.Translate(-outerBounds.X0, -outerBounds.Y0)
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "my-custom-id",
	})
	borderInner, err := canvas.ParseSVGPath("m 14.201674,21.416016 v 0.164061 l 0.502245,0.491332 h 0.13177 v 2.671467 3.016954 2.658312 h -0.13177 l -0.502245,0.491332 v 0.16406 H 7.714522 1.227367 v -0.16406 L 0.725122,30.418142 H 0.593353 V 27.75983 24.742876 22.071409 h 0.131769 l 0.502245,-0.491332 v -0.164061 h 5.246151 l 1.101031,0.748466 h 0.00937 0.016394 0.096817 0.051539 0.096819 0.016393 0.00937 L 8.97228,21.416016 h 5.229396")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	borderInner = borderInner.Translate(-outerBounds.X0, -outerBounds.Y0)
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "my-custom-id",
	})

	// decorative lines
	paths := []string{
		"M 6.11489,21.705244 H 1.732577 V 21.60063 H 6.11489 v 0.104614",
		"M 6.11489,21.905256 H 1.732577 V 21.800643 H 6.11489 v 0.104613",
		"m 2.275541,24.268541 h 10.98395 v 0.07983 H 2.275541 v -0.07983",
		"m 1.769095,28.64164 h 11.996871 v 0.08721 H 1.769095 v -0.08721",
		"M 13.717872,21.705244 H 9.335559 V 21.60063 h 4.382313 v 0.104614",
		"M 13.717872,21.905256 H 9.335559 v -0.104613 h 4.382313 v 0.104613",
	}
	for _, pathStr := range paths {
		path, err := canvas.ParseSVGPath(pathStr)
		if err != nil {
			log.Printf("Failed to parse SVG path: %s", err)
			continue
		}
		path = path.Translate(-outerBounds.X0, -outerBounds.Y0)
		builder.AddPath(path.ToSVG(), map[string]string{
			"fill": backgroundColor,
			"id":   "my-custom-id",
		})
	}
}

func drawDeco(
	size string,
	width float64,
	height float64,
	foregroundColor string,
	backgroundColor string,
	lines []string,
) string {
	// Initialize SVG builder
	builder := svgutils.NewSVGBuilder(width, height)

	if size == "medium" {
		drawMediumDeco(builder, foregroundColor, backgroundColor)
	} else if size == "large" {
		drawLargeDeco(builder, foregroundColor, backgroundColor)
	} else {
		log.Fatalf("Invalid size: %s", size)
	}

	// draw text
	fontFamily := canvas.NewFontFamily("Limelight")
	if err := fontFamily.LoadFont(fonts.Limelight, 0, canvas.FontRegular); err != nil {
		log.Println("Failed to load Limelight font: ", err)
		panic("Font loading error")
	}

	numLines := len(lines)
	// func ensureLengthThree(slice []int) []int {
	// 	for len(slice) < 3 {
	// 			slice = append(slice, 0) // Append default value (0 for int)
	// 	}
	// 	return slice
	// }

	if numLines > 0 {
		// for len(lines) < 3 {
		// 	lines = append([]string{""}, lines...)
		// }
		// numLines = len(lines)

		// Calculate the total height for the text containers
		const topPadding = 2.25
		const bottomPadding = 0.75
		lineSpacing := 0.75
		availableHeight := height - (topPadding + bottomPadding) - float64(numLines-1)*lineSpacing

		// Determine the heights for each line
		var containerHeights []float64
		switch numLines {
		case 3:
			containerHeights = []float64{
				availableHeight * 0.2,  // First line
				availableHeight * 0.55, // Middle line (taller)
				availableHeight * 0.25, // Last line
			}
		case 2:
			containerHeights = []float64{
				availableHeight * 0.75, // First line (taller)
				availableHeight * 0.25, // Last line
			}
		default:
			// containerHeights = []float64{availableHeight} // Single line
			containerHeights = []float64{availableHeight}
		}

		// Calculate the starting y position for the topmost line
		currentY := height - topPadding

		for i, line := range lines {
			// containerHeight := containerHeights[i]
			log.Print("i: ", i)
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

	// for i, line := range lines {
	// 	// draw text container
	// 	var container *canvas.Path
	// 	var containerX, containerY float64
	// 	var containerBounds canvas.Rect
	// 	if i == 0 {
	// 		container = canvas.Rectangle(width-2, 1.5)
	// 		containerBounds = container.Bounds()
	// 		containerX = width/2 - containerBounds.W()/2
	// 		containerY = height - 3.4
	// 	} else if i == 1 {
	// 		container = canvas.Rectangle(width-2, 3.0)
	// 		containerBounds = container.Bounds()
	// 		containerX = width/2 - containerBounds.W()/2
	// 		containerY = height/2 - containerBounds.H()/2 - 0.5
	// 	} else if i == 2 {
	// 		container = canvas.Rectangle(width-2, 1.7)
	// 		containerBounds = container.Bounds()
	// 		containerX = width/2 - containerBounds.W()/2
	// 		containerY = 0.8
	// 	}
	// 	// *************************************************************
	// 	// ********** uncomment to draw text container border **********
	// 	// *************************************************************
	// 	container = container.Translate(containerX, containerY)
	// 	container = container.Scale(1, -1)
	// 	container = container.Translate(0, height)
	// 	builder.AddPath(container.ToSVG(), map[string]string{
	// 		"fill":         "none",
	// 		"stroke":       "pink",
	// 		"stroke-width": "0.025",
	// 		"id":           fmt.Sprintf("text-container-%d", i),
	// 	})

	// 	// draw text
	// 	fontSize := 1.0
	// 	face := fontFamily.Face(fontSize, canvas.FontRegular, canvas.FontNormal)
	// 	textPath, _, err := face.ToPath(line)
	// 	if err != nil {
	// 		log.Fatalf("Failed to convert text to path: %s", err)
	// 	}
	// 	textBounds := textPath.Bounds()

	// 	// Calculate the scale factor to fit the path within the container
	// 	scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/textBounds.H())
	// 	textPath.Scale(scale, scale)

	// 	// recalculate the bounds after scaling
	// 	textBounds = textPath.Bounds()
	// 	x := containerX + containerBounds.W()/2 - textBounds.W()/2
	// 	y := containerY + containerBounds.H()/2 - textBounds.H()/2
	// 	textPath = textPath.Translate(x, y)
	// 	textPath = textPath.Scale(1, -1)
	// 	textPath = textPath.Translate(0, height)
	// 	builder.AddPath(textPath.ToSVG(), map[string]string{
	// 		"fill": backgroundColor,
	// 		"id":   fmt.Sprintf("text-line-%d", i),
	// 	})
	// }

	builder.EndGroup()

	return builder.Close()
}
