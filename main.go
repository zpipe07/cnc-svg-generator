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

func generateSVG(lines []string, shape string) string {
	fontFamily := canvas.NewFontFamily("Bungee")
	err := fontFamily.LoadFont(bungeeFont, 0, canvas.FontRegular)
	if err != nil {
		panic("Failed to load font: " + err.Error())
	}
	face := fontFamily.Face(44, canvas.FontRegular, canvas.FontNormal)

	// Calculate the total canvas height based on the number of lines
	lineHeight := 25.0                            // Adjust as needed for line spacing
	height := float64(len(lines))*lineHeight + 40 // Add extra padding
	width := 200.0                                // Fixed width, adjust if necessary

	log.Println("Canvas width: ", width)
	log.Println("Canvas height: ", height)

	// Create a new canvas with dynamic width and height
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	// Draw the appropriate shape around the text
	ctx.SetStrokeColor(canvas.Black)
	ctx.SetStrokeWidth(1)
	ctx.SetFillColor(canvas.White)

	if shape == "ellipse" {
		ellipsePath := canvas.Ellipse(width/2, height/2)
		// ellipsePath = ellipsePath.Translate(100, 50)
		ellipsePath = ellipsePath.Translate(width/2, height/2)
		ctx.DrawPath(0, 0, ellipsePath)
	} else if shape == "rectangle" {
		// Default to rectangle if shape is not "ellipse"
		rectPath := canvas.RoundedRectangle(width, height, 10)
		rectPath = rectPath.Translate(0, 0)
		ctx.DrawPath(0, 0, rectPath)
	} else {
		panic("Invalid shape: " + shape)
	}

	// Draw each line of text inside the shape
	ctx.SetFillColor(canvas.Black)
	for i, line := range lines {
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
	svgCanvas := svg.New(&buf, width, height, nil)
	c.RenderTo(svgCanvas)
	svgCanvas.Close()

	return buf.String()
}

func main() {
	r := gin.Default()

	r.GET("/svg", func(c *gin.Context) {
		text := c.QueryArray("text")
		shape := c.DefaultQuery("shape", "rectangle")
		log.Println("Shape: ", shape)
		svgContent := generateSVG(text, shape)
		c.Data(http.StatusOK, "image/svg+xml", []byte(svgContent))
	})

	r.Run("localhost:8080") // Start the server on port 8080
}
