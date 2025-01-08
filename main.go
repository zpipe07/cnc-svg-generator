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

func generateSVG(text string) string {

	// Create a new canvas with a defined width and height (e.g., 100x100 mm)
	c := canvas.New(200, 200)
	ctx := canvas.NewContext(c)

	// Draw text if provided
	if text != "" {
		ctx.SetFillColor(canvas.Black)
		fontFamily := canvas.NewFontFamily("Bungee")
		err := fontFamily.LoadFont(bungeeFont, 0, canvas.FontRegular)
		if err != nil {
			panic("Failed to load font: " + err.Error())
		}
		face := fontFamily.Face(44, canvas.FontRegular, canvas.FontNormal)

		// Convert text to a path
		textPath, _, err := face.ToPath(text)
		if err != nil {
			panic("Failed to convert text to path: " + err.Error())
		}
		textBounds := textPath.Bounds()
		textX := 100 - textBounds.W()/2
		textY := 100 - textBounds.H()/2

		// Move the text path to the desired position
		textPath = textPath.Translate(textX, textY)

		// Draw the text path
		ctx.SetFillColor(canvas.Black)
		ctx.DrawPath(0, 0, textPath)
	}

	// Export the canvas to an SVG
	var buf bytes.Buffer
	svgCanvas := svg.New(&buf, 200, 200, nil)
	c.RenderTo(svgCanvas)
	svgCanvas.Close()

	return buf.String()
}

func main() {
	r := gin.Default()

	r.GET("/svg", func(c *gin.Context) {
		text := c.Query("text")
		svgContent := generateSVG(text)
		c.Data(http.StatusOK, "image/svg+xml", []byte(svgContent))
	})

	r.Run(":8080") // Start the server on port 8080
}
