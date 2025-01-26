package svg

import (
	"cnc-svg-generator/pkg/signs"
	"fmt"

	"github.com/tdewolff/canvas"
)

func GenerateSVG(
	productId string,
	width float64,
	height float64,
	lines []string,
	fontFamily *canvas.FontFamily,
	// foregroundColor color.RGBA,
	// backgroundColor color.RGBA,
	foregroundColor string,
	backgroundColor string,
) string {

	// Create a new canvas with dynamic width and height
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	// Draw the appropriate shape around the text
	ctx.SetStrokeWidth(0.0125)

	svgStr := signs.DrawSign(
		ctx,
		productId,
		width,
		height,
		foregroundColor,
		backgroundColor,
		lines,
		fontFamily,
	)

	// Export the canvas to an SVG
	// var buf bytes.Buffer
	// opts := &svg.Options{
	// 	SizeUnits: "in",
	// }
	// svgCanvas := svg.New(&buf, width, height, opts)
	// c.RenderTo(svgCanvas)
	// svgCanvas.Close()
	// svgString := buf.String()

	// Specify attributes to modify or add
	attributesToModify := map[string]string{
		"width":          "100%",
		"height":         "100%",
		"viewBox":        fmt.Sprintf("0 0 %f %f", width, height),
		"xmlns":          "http://www.w3.org/2000/svg",
		"xmlns:xlink":    "http://www.w3.org/1999/xlink",
		"filter":         "drop-shadow(rgba(0, 0, 0, 0.5) 0px 0px 2px)",
		"xmlns:inkscape": "http://www.inkscape.org/namespaces/inkscape",
	}
	// Modify or add attributes
	modifiedSVG, err := modifyOrAddSVGAttributes(svgStr, attributesToModify)
	if err != nil {
		fmt.Println("Error modifying SVG:", err)
	}

	return modifiedSVG
}
