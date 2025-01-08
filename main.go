package main

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/svg"
)

func main() {
	r := gin.Default()

	r.GET("/svg", func(c *gin.Context) {
		svgContent := generateSVG()
		c.Data(http.StatusOK, "image/svg+xml", []byte(svgContent))
	})

	r.Run(":8080") // Start the server on port 8080
}

func generateSVG() string {
	// Create a new canvas with a defined width and height (e.g., 100x100 mm)
	c := canvas.New(100, 100)
	ctx := canvas.NewContext(c)

	// Draw some shapes on the canvas
	ctx.SetFillColor(canvas.Lightblue)
	// ctx.DrawRect(10, 10, 80, 80) // Draw a rectangle
	ctx.DrawPath(10, 10, canvas.Rectangle(80, 80)) // Draw a rectangle

	ctx.SetStrokeColor(canvas.Black)
	ctx.SetStrokeWidth(2)
	ctx.DrawPath(10, 10, canvas.Line(80, 80)) // Draw a diagonal line

	ctx.SetStrokeColor(canvas.Red)
	ctx.DrawPath(50, 50, canvas.Circle(30)) // Draw a circle

	// Export the canvas to an SVG
	var buf bytes.Buffer
	svgCanvas := svg.New(&buf, 100, 100, nil)
	c.RenderTo(svgCanvas)
	svgCanvas.Close()

	return buf.String()
}
