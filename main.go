package main

import (
	"bytes"
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/svg"
)

//go:embed fonts/Bungee-Regular.ttf
var bungeeFont []byte

func generateSVG(lines []string) string {
	fontFamily := canvas.NewFontFamily("Bungee")
	err := fontFamily.LoadFont(bungeeFont, 0, canvas.FontRegular)
	if err != nil {
		panic("Failed to load font: " + err.Error())
	}
	face := fontFamily.Face(44, canvas.FontRegular, canvas.FontNormal)

	// Calculate the total canvas height based on the number of lines
	lineHeight := 50.0                            // Adjust as needed for line spacing
	height := float64(len(lines))*lineHeight + 40 // Add extra padding
	width := 400.0                                // Fixed width, adjust if necessary

	// Create a new canvas with dynamic width and height
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	// Calculate dimensions of the single rounded rectangle
	rectX := 10.0
	rectY := 10.0
	rectWidth := width - 20.0
	rectHeight := height - 20.0

	// Draw the single rounded rectangle
	rectPath := canvas.RoundedRectangle(rectWidth, rectHeight, 10)
	rectPath = rectPath.Translate(rectX, rectY)
	ctx.SetStrokeColor(canvas.Black)
	ctx.SetStrokeWidth(1)
	ctx.SetFillColor(canvas.White)
	ctx.DrawPath(0, 0, rectPath)

	// Draw each line of text inside the rectangle
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
		svgContent := generateSVG(text)
		c.Data(http.StatusOK, "image/svg+xml", []byte(svgContent))
	})

	r.Run(":8080") // Start the server on port 8080
}
