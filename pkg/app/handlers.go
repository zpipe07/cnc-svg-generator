package app

import (
	"cnc-svg-generator/pkg/fonts"
	"cnc-svg-generator/pkg/svg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tdewolff/canvas"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "weight tracker API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) GetSvg() gin.HandlerFunc {
	return func(c *gin.Context) {

		productId := c.Query("productId")
		widthStr := c.Query("width")
		heightStr := c.Query("height")
		text := c.QueryArray("text")
		fontFamilyStr := c.Query("fontFamily")
		bgColorStr := c.Query("backgroundColor")
		fgColorStr := c.Query("foregroundColor")

		// Parse colors
		backgroundColor, err := svg.ParseColor(bgColorStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid backgroundColor."})
			return
		}
		foregroundColor, err := svg.ParseColor(fgColorStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid foregroundColor."})
			return
		}

		// Parse font family
		var fontFamily *canvas.FontFamily
		if fontFamilyStr != "" {
			fontFamily, err = fonts.ParseFontFamily(fontFamilyStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fontFamily."})
				return
			}
		}

		// Validate text
		if len(text) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The text parameter is required."})
			return
		}

		width, err := strconv.ParseFloat(widthStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid width."})
			return
		}
		height, err := strconv.ParseFloat(heightStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid height."})
			return
		}

		svgContent := svg.GenerateSVG(
			productId,
			width,
			height,
			text,
			fontFamily,
			foregroundColor,
			backgroundColor,
		)

		c.Header("Content-Type", "image/svg+xml")
		c.String(http.StatusOK, svgContent)
	}
}

// func (s *Server) CreateUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Header("Content-Type", "application/json")

// 		var newUser api.NewUserRequest

// 		err := c.ShouldBindJSON(&newUser)

// 		if err != nil {
// 			log.Printf("handler error: %v", err)
// 			c.JSON(http.StatusBadRequest, nil)
// 			return
// 		}

// 		err = s.userService.New(newUser)

// 		if err != nil {
// 			log.Printf("service error: %v", err)
// 			c.JSON(http.StatusInternalServerError, nil)
// 			return
// 		}

// 		response := map[string]string{
// 			"status": "success",
// 			"data":   "new user created",
// 		}

// 		c.JSON(http.StatusOK, response)
// 	}
// }
