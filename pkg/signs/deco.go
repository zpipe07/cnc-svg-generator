package signs

import (
	"cnc-svg-generator/pkg/fonts"
	"image/color"
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

func drawDeco(
	ctx *canvas.Context,
	width float64,
	height float64,
	foregroundColor color.RGBA,
	backgroundColor color.RGBA,
	lines []string,
) {
	ctx.SetStrokeColor(canvas.Black)
	ctx.SetStrokeWidth(0.025)
	ctx.SetFillColor(foregroundColor)
	outerEdge, err := canvas.ParseSVGPath("M 0.908768,31.446765 C 0.875547,31.44691 0.848537,31.420418 0.848537,31.387689 V 31.12425 C 0.849094,31.08683 0.834179,31.05079 0.807221,31.02442 L 0.610317,30.831796 C 0.584074,30.805915 0.548463,30.791372 0.51133,30.791372 H 0.274488 c -0.033117,0 -0.059964,-0.02645 -0.059964,-0.05908 v -8.974773 c -1.47e-4,-0.03273 0.026742,-0.05934 0.059964,-0.05934 h 0.234269 c 0.037815,6.78e-4 0.074297,-0.01375 0.101135,-0.04001 L 0.808308,21.46406 c 0.025779,-0.02579 0.040229,-0.06053 0.040229,-0.09672 v -0.265472 c 0,-0.01589 0.006366,-0.03031 0.0167218,-0.04093 0.0109112,-0.01119 0.0262519,-0.01815 0.0536598,-0.01815 H 6.103881 c 0.037211,4.2e-4 0.073057,-0.01379 0.099612,-0.03947 0.094778,-0.09152 0.2440221,-0.12094 0.387435,-0.08662 0.057473,0.01533 0.118492,-0.0069 0.1520071,-0.05539 0.095811,-0.138195 0.280362,-0.206655 0.476257,-0.181385 0.050567,0.0063 0.100591,-0.01506 0.130568,-0.05567 0.054649,-0.07392 0.167241,-0.17749 0.376076,-0.17749 0.207887,0 0.320405,0.102622 0.375335,0.176472 0.029978,0.04144 0.080717,0.06321 0.131936,0.05661 0.195214,-0.02499 0.37908,0.04309 0.474978,0.180508 0.03283,0.04829 0.092831,0.07103 0.150035,0.05699 0.143113,-0.03505 0.292412,-0.0065 0.387872,0.08389 0.026323,0.02658 0.062421,0.04156 0.100112,0.04156 h 5.1636789 c 0.04367,-1.89e-4 0.07073,0.02632 0.07073,0.05907 v 0.263422 c -5.61e-4,0.03743 0.01436,0.07347 0.04132,0.09985 l 0.197806,0.193493 c 0.02617,0.02535 0.0614,0.03955 0.09809,0.03955 h 0.236831 c 0.03312,0 0.05997,0.02645 0.05997,0.05908 v 8.974824 c 1.18e-4,0.03271 -0.02676,0.05929 -0.05997,0.05929 h -0.234426 c -0.03776,-6.33e-4 -0.07417,0.0138 -0.100966,0.04001 l -0.198489,0.19416 c -0.02574,0.02578 -0.04017,0.0605 -0.04017,0.09666 v 0.265484 c 0,0.03263 -0.02685,0.05908 -0.07038,0.05908 l -13.601357,0")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	outerBounds := outerEdge.Bounds()
	outerEdge = outerEdge.Translate(-outerBounds.X0, -outerBounds.Y0)
	outerEdge = outerEdge.Scale(1, -1)
	outerEdge = outerEdge.Translate(0, height)
	ctx.DrawPath(0, 0, outerEdge)

	ctx.SetStroke(nil)
	ctx.SetFillColor(backgroundColor)
	borderOuter, err := canvas.ParseSVGPath("m 14.40157,21.21908 v 0.279053 l 0.384801,0.376435 0.24925,-4.7e-5 -5e-5,2.868355 v 3.016954 2.855247 h -0.249268 l -0.384733,0.37638 v 0.279011 H 7.714521 1.027485 V 30.991457 L 0.64274,30.615077 H 0.393471 V 27.75983 24.742876 l -9.3e-4,-2.868355 0.248382,4.7e-5 0.3848,-0.376435 v -0.279055 l 5.226191,-2e-6 c 0.026775,-0.04548 0.094874,-0.118972 0.258125,-0.130022 l 0.94862,0.936577 c 0.00438,-0.0043 0.00896,-0.0083 0.013705,-0.0122 l -0.6523,-0.891281 c 0.103078,-0.232353 0.276781,-0.269653 0.279571,-0.272353 l 0.427643,1.12502 c 0.014554,-0.0084 0.029644,-0.01592 0.045167,-0.02232 L 7.510393,20.708228 c 0.00212,-8.3e-4 0.050761,-0.08309 0.211624,-0.08309 0.160878,0 0.209505,0.08226 0.211625,0.08309 l -0.061971,1.244211 c 0.01571,0.0065 0.030827,0.01406 0.045199,0.02246 l 0.427545,-1.1251 c 0.0028,0.0027 0.176478,0.04 0.27957,0.272353 l -0.652243,0.891193 c 0.00469,0.004 0.00928,0.008 0.013763,0.01218 l 0.948492,-0.936467 c 0.163264,0.01105 0.231349,0.08454 0.258138,0.130022 h 5.209435")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	borderOuter = borderOuter.Translate(-outerBounds.X0, -outerBounds.Y0)
	borderOuter = borderOuter.Scale(1, -1)
	borderOuter = borderOuter.Translate(0, height)
	ctx.DrawPath(0, 0, borderOuter)

	ctx.SetFillColor(foregroundColor)
	borderInner, err := canvas.ParseSVGPath("m 14.201674,21.416016 v 0.164061 l 0.502245,0.491332 h 0.13177 v 2.671467 3.016954 2.658312 h -0.13177 l -0.502245,0.491332 v 0.16406 H 7.714522 1.227367 v -0.16406 L 0.725122,30.418142 H 0.593353 V 27.75983 24.742876 22.071409 h 0.131769 l 0.502245,-0.491332 v -0.164061 h 5.246151 l 1.101031,0.748466 h 0.00937 0.016394 0.096817 0.051539 0.096819 0.016393 0.00937 L 8.97228,21.416016 h 5.229396")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %s", err)
	}
	borderInner = borderInner.Translate(-outerBounds.X0, -outerBounds.Y0)
	borderInner = borderInner.Scale(1, -1)
	borderInner = borderInner.Translate(0, height)
	ctx.DrawPath(0, 0, borderInner)

	// decorative lines
	ctx.SetFillColor(backgroundColor)
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
		path = path.Scale(1, -1)
		path = path.Translate(0, height)
		ctx.DrawPath(0, 0, path)
	}

	// draw text
	fontFamily := canvas.NewFontFamily("Limelight")
	if err := fontFamily.LoadFont(fonts.Limelight, 0, canvas.FontRegular); err != nil {
		log.Println("Failed to load Limelight font: ", err)
		panic("Font loading error")
	}

	for i, line := range lines {
		// draw text container
		ctx.SetFill(nil)

		// uncomment to draw text container border
		// ctx.SetStrokeWidth(0.1)
		// ctx.SetStrokeColor(canvas.Lightgray)

		var container *canvas.Path
		var containerX, containerY float64
		var containerBounds canvas.Rect
		if i == 0 {
			container = canvas.Rectangle(width-2, 1.5)
			containerBounds = container.Bounds()
			containerX = width/2 - containerBounds.W()/2
			containerY = height - 3.4
		} else if i == 1 {
			container = canvas.Rectangle(width-2, 3.0)
			containerBounds = container.Bounds()
			containerX = width/2 - containerBounds.W()/2
			containerY = height/2 - containerBounds.H()/2 - 0.5
		} else if i == 2 {
			container = canvas.Rectangle(width-2, 1.7)
			containerBounds = container.Bounds()
			containerX = width/2 - containerBounds.W()/2
			containerY = 0.8
		}
		ctx.DrawPath(containerX, containerY, container)

		// draw text
		ctx.SetStroke(nil)
		ctx.SetFillColor(backgroundColor)
		fontSize := 8 - float64(len(line))*0.25
		face := fontFamily.Face(fontSize, canvas.FontRegular, canvas.FontNormal)
		textPath, _, err := face.ToPath(line)
		if err != nil {
			log.Fatalf("Failed to convert text to path: %s", err)
		}
		textBounds := textPath.Bounds()

		// Calculate the scale factor to fit the path within the container
		scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/textBounds.H())
		textPath.Scale(scale, scale)

		// recalculate the bounds after scaling
		textBounds = textPath.Bounds()
		x := containerX + containerBounds.W()/2 - textBounds.W()/2
		y := containerY + containerBounds.H()/2 - textBounds.H()/2
		ctx.DrawPath(x, y, textPath)
	}
}
