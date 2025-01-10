package main

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/svg", func(c *gin.Context) {
		text := c.QueryArray("text")
		shape := c.DefaultQuery("shape", "rectangle")
		fontFamilyStr := c.Query("fontFamily")
		bgColorStr := c.Query("backgroundColor")
		fgColorStr := c.Query("foregroundColor")

		// Parse colors
		backgroundColor, err := parseColor(bgColorStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid backgroundColor. Supported values are 'white', 'black', 'red', 'blue', or 'green'."})
			return
		}
		foregroundColor, err := parseColor(fgColorStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid foregroundColor. Supported values are 'white', 'black', 'red', 'blue', or 'green'."})
			return
		}
		fontFamily, err := parseFontFamily(fontFamilyStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fontFamily. Supported values are 'bungee' or 'limelight'."})
			return
		}
		if shape != "ellipse" && shape != "rectangle" && shape != "deco" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shape. Supported values are 'ellipse', 'rectangle', or 'deco'."})
			return
		}
		if len(text) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The text parameter is required."})
			return
		}

		svgContent := generateSVG(text, shape, fontFamily, foregroundColor, backgroundColor)
		c.Data(http.StatusOK, "image/svg+xml", []byte(svgContent))
	})

	r.Run("localhost:8080") // Start the server on port 8080
}
