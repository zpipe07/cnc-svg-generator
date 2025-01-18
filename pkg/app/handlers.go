package app

import (
	"cnc-svg-generator/pkg/svg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.Header("Content-Type", "image/svg+xml")

		productId := c.Query("productId")
		text := c.QueryArray("text")
		fontFamilyStr := c.Query("fontFamily")
		bgColorStr := c.Query("backgroundColor")
		fgColorStr := c.Query("foregroundColor")

		productConfig, err := svg.GetProductConfig(productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid productId."})
			return
		}
		log.Println("Product config: ", productConfig)

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
		fontFamily, err := svg.ParseFontFamily(fontFamilyStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fontFamily."})
			return
		}

		// Validate text
		if len(text) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The text parameter is required."})
			return
		}

		svgContent := svg.GenerateSVG(
			productConfig,
			text,
			fontFamily,
			foregroundColor,
			backgroundColor,
		)
		// c.Data(http.StatusOK, "image/svg+xml", []byte(svgContent))

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
