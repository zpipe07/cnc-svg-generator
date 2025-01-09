package main

import (
	"bytes"
	_ "embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/svg"
)

//go:embed fonts/Bungee-Regular.ttf
var bungeeFont []byte

//go:embed fonts/Limelight-Regular.ttf
var limelightFont []byte

func loadFontFamily(font string) *canvas.FontFamily {
	var fontFamily *canvas.FontFamily
	switch font {
	case "limelight":
		fontFamily = canvas.NewFontFamily("Limelight")
		if err := fontFamily.LoadFont(limelightFont, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load Limelight font: ", err)
			panic("Font loading error")
		}
	case "bungee":
		fallthrough
	default:
		fontFamily = canvas.NewFontFamily("Bungee")
		if err := fontFamily.LoadFont(bungeeFont, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load Bungee font: ", err)
			panic("Font loading error")
		}
	}
	return fontFamily
}

func drawShape(ctx *canvas.Context, width float64, height float64, shape string) {
	if shape == "ellipse" {
		ellipsePath := canvas.Ellipse(width/2, height/2)
		ellipsePath = ellipsePath.Translate(width/2, height/2)
		ctx.DrawPath(0, 0, ellipsePath)
	} else if shape == "rectangle" {
		rectPath := canvas.RoundedRectangle(width, height, 0.25)
		rectPath = rectPath.Translate(0, 0)
		ctx.DrawPath(0, 0, rectPath)
	} else if shape == "deco" {
		// Draw a decorative shape around the text
		svgStr := []string{
			"M 14.306 0.596 L 9.132 0.596 C 9.094 0.596 9.058 0.581 9.031 0.555 C 8.946 0.474 8.817 0.442 8.688 0.462 A 0.511 0.511 0 0 0 8.644 0.471 A 0.142 0.142 0 0 1 8.628 0.474 C 8.622 0.475 8.616 0.475 8.61 0.475 A 0.14 0.14 0 0 1 8.494 0.414 A 0.474 0.474 0 0 0 8.098 0.228 A 0.624 0.624 0 0 0 8.019 0.233 A 0.142 0.142 0 0 1 8 0.235 A 0.14 0.14 0 0 1 7.887 0.177 A 0.434 0.434 0 0 0 7.842 0.126 A 0.455 0.455 0 0 0 7.511 0 A 0.455 0.455 0 0 0 7.181 0.126 A 0.434 0.434 0 0 0 7.135 0.178 C 7.105 0.218 7.055 0.24 7.005 0.233 C 6.809 0.208 6.624 0.277 6.528 0.415 A 0.14 0.14 0 0 1 6.413 0.475 C 6.407 0.475 6.402 0.475 6.396 0.474 A 0.142 0.142 0 0 1 6.376 0.47 A 0.511 0.511 0 0 0 6.335 0.462 C 6.31 0.458 6.285 0.456 6.26 0.456 A 0.387 0.387 0 0 0 5.989 0.557 A 0.141 0.141 0 0 1 5.891 0.596 A 0.14 0.14 0 0 1 5.889 0.596 L 0.694 0.596 C 0.661 0.596 0.634 0.623 0.634 0.655 L 0.634 0.921 C 0.634 0.957 0.62 0.992 0.594 1.018 L 0.395 1.212 A 0.141 0.141 0 0 1 0.297 1.252 A 0.143 0.143 0 0 1 0.294 1.252 L 0.06 1.252 C 0.027 1.252 0 1.278 0 1.311 L 0 10.286 C 0 10.318 0.027 10.345 0.06 10.345 L 0.297 10.345 C 0.334 10.345 0.37 10.359 0.396 10.385 L 0.593 10.578 A 0.137 0.137 0 0 1 0.634 10.676 A 0.138 0.138 0 0 1 0.634 10.678 L 0.634 10.941 C 0.634 10.974 0.661 11 0.694 11 L 14.306 11 C 14.339 11 14.366 10.974 14.366 10.941 L 14.366 10.676 A 0.137 0.137 0 0 1 14.406 10.579 L 14.605 10.385 A 0.141 0.141 0 0 1 14.703 10.345 A 0.142 0.142 0 0 1 14.706 10.345 L 14.94 10.345 A 0.06 0.06 0 0 0 15 10.286 L 15 1.311 A 0.06 0.06 0 0 0 14.94 1.252 L 14.703 1.252 C 14.667 1.252 14.631 1.238 14.605 1.212 L 14.407 1.019 A 0.137 0.137 0 0 1 14.366 0.921 A 0.137 0.137 0 0 1 14.366 0.919 L 14.366 0.655 A 0.06 0.06 0 0 0 14.306 0.596 Z",
			"M 0.813 10.823 L 0.813 10.544 L 0.429 10.168 L 0.179 10.168 L 0.179 7.312 L 0.179 4.295 L 0.178 1.427 L 0.427 1.427 L 0.812 1.051 L 0.812 0.772 L 6.038 0.772 C 6.065 0.726 6.133 0.653 6.296 0.642 L 7.245 1.578 C 7.249 1.574 7.254 1.57 7.258 1.566 L 6.606 0.675 A 0.561 0.561 0 0 1 6.708 0.517 A 0.406 0.406 0 0 1 6.886 0.402 L 7.313 1.527 C 7.328 1.519 7.343 1.511 7.358 1.505 L 7.296 0.261 A 0.173 0.173 0 0 1 7.342 0.219 A 0.302 0.302 0 0 1 7.508 0.178 A 0.302 0.302 0 0 1 7.674 0.219 A 0.173 0.173 0 0 1 7.72 0.261 L 7.658 1.505 C 7.673 1.511 7.688 1.519 7.703 1.527 L 8.13 0.402 A 0.406 0.406 0 0 1 8.307 0.517 A 0.561 0.561 0 0 1 8.41 0.675 L 7.758 1.566 C 7.762 1.57 7.767 1.574 7.771 1.578 L 8.72 0.642 C 8.883 0.653 8.951 0.726 8.978 0.772 L 14.187 0.772 L 14.187 1.051 L 14.572 1.427 L 14.822 1.427 L 14.821 4.295 L 14.821 7.312 L 14.821 10.168 L 14.572 10.168 L 14.187 10.544 L 14.187 10.823 L 7.5 10.823 L 0.813 10.823 Z",
			"M 1.013 10.629 L 1.013 10.465 L 0.511 9.974 L 0.379 9.974 L 0.379 7.315 L 0.379 4.298 L 0.379 1.627 L 0.511 1.627 L 1.013 1.135 L 1.013 0.971 L 6.259 0.971 L 7.36 1.72 L 7.369 1.72 L 7.386 1.72 L 7.483 1.72 L 7.534 1.72 L 7.631 1.72 L 7.647 1.72 L 7.657 1.72 L 8.758 0.971 L 13.987 0.971 L 13.987 1.135 L 14.489 1.627 L 14.621 1.627 L 14.621 4.298 L 14.621 7.315 L 14.621 9.974 L 14.489 9.974 L 13.987 10.465 L 13.987 10.629 L 7.5 10.629 L 1.013 10.629 Z",
		}
		// iterate over each SVG path and draw it on the canvas
		for _, svgStr := range svgStr {
			decoPath, err := canvas.ParseSVGPath(svgStr)
			if err != nil {
				panic("Failed to parse SVG path: " + err.Error())
			}
			ctx.DrawPath(0, 0, decoPath)
		}
	} else {
		panic("Invalid shape: " + shape)
	}
}

func generateSVG(lines []string, shape string, font string) string {
	// Hardcoded dimensions in inches
	width := 15.0
	height := 11.0

	fontFamily := loadFontFamily(font)

	// Create a new canvas with dynamic width and height
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	// Draw the appropriate shape around the text
	ctx.SetStrokeColor(canvas.Black)
	ctx.SetStrokeWidth(0.025)
	ctx.SetFillColor(canvas.White)

	drawShape(ctx, width, height, shape)

	// Draw each line of text inside the shape
	for i, line := range lines {
		fontSize := 10 - float64(len(line))*0.5
		lineHeight := 3.0
		face := fontFamily.Face(fontSize, canvas.FontRegular, canvas.FontNormal)
		textPath, _, err := face.ToPath(line)
		if err != nil {
			panic("Failed to convert text to path: " + err.Error())
		}
		textBounds := textPath.Bounds()
		textX := (width - textBounds.W()) / 2
		textY := lineHeight*float64(i) + lineHeight // Position each line below the previous
		textPath = textPath.Translate(textX, textY)
		ctx.DrawPath(0, 0, textPath)
	}

	// Export the canvas to an SVG
	var buf bytes.Buffer
	opts := &svg.Options{
		SizeUnits: "in",
	}
	svgCanvas := svg.New(&buf, width, height, opts)
	c.RenderTo(svgCanvas)
	svgCanvas.Close()

	return buf.String()
}

func main() {
	r := gin.Default()

	r.GET("/svg", func(c *gin.Context) {
		text := c.QueryArray("text")
		shape := c.DefaultQuery("shape", "rectangle")
		font := c.DefaultQuery("font", "bungee")

		if shape != "ellipse" && shape != "rectangle" && shape != "deco" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shape. Supported values are 'ellipse', 'rectangle', or 'deco'."})
			return
		}
		if font != "bungee" && font != "limelight" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid font. Supported values are 'bungee' or 'limelight'."})
			return
		}

		log.Printf("Shape: %s, Font: %s", shape, font)

		svgContent := generateSVG(text, shape, font)
		c.Data(http.StatusOK, "image/svg+xml", []byte(svgContent))
	})

	r.Run("localhost:8080") // Start the server on port 8080
}
