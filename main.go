package main

import (
	_ "embed"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alderProductId := os.Getenv("ALDER_PRODUCT_ID")
	log.Println("Alder product ID: ", alderProductId)

	r := gin.Default()

	r.GET("/svg", func(c *gin.Context) {
		productId := c.Query("productId")
		text := c.QueryArray("text")
		fontFamilyStr := c.Query("fontFamily")
		bgColorStr := c.Query("backgroundColor")
		fgColorStr := c.Query("foregroundColor")

		productConfig, err := getProductConfig(productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid productId. Supported values are 'alder'."})
			return
		}
		log.Println("Product config: ", productConfig)

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

		// Parse font family
		fontFamily, err := parseFontFamily(fontFamilyStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fontFamily. Supported values are 'bungee' or 'limelight'."})
			return
		}

		// Validate text
		if len(text) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The text parameter is required."})
			return
		}

		svgContent := generateSVG(
			productConfig,
			text,
			fontFamily,
			foregroundColor,
			backgroundColor,
		)
		c.Data(http.StatusOK, "image/svg+xml", []byte(svgContent))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if no PORT is set
	}

	r.Run("0.0.0.0:8080") // Start the server on port 8080
}
