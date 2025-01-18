package svg

import (
	_ "embed"
	"log"

	"github.com/tdewolff/canvas"
)

//go:embed fonts/ArbutusSlab-Regular.ttf
var arbutusFont []byte

//go:embed fonts/Bungee-Regular.ttf
var bungeeFont []byte

//go:embed fonts/Danfo-Regular.ttf
var danfoFont []byte

//go:embed fonts/Limelight-Regular.ttf
var limelightFont []byte

//go:embed fonts/Shrikhand-Regular.ttf
var shrikhandFont []byte

func ParseFontFamily(fontFamilyStr string) (*canvas.FontFamily, error) {
	var fontFamily *canvas.FontFamily
	switch fontFamilyStr {
	case "arbutus":
		fontFamily = canvas.NewFontFamily("Arbutus")
		if err := fontFamily.LoadFont(arbutusFont, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load Arbutus font: ", err)
			panic("Font loading error")
		}
	case "danfo":
		fontFamily = canvas.NewFontFamily("Danfo")
		if err := fontFamily.LoadFont(danfoFont, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load Danfo font: ", err)
			panic("Font loading error")
		}
	case "shrikhand":
		fontFamily = canvas.NewFontFamily("Shrikhand")
		if err := fontFamily.LoadFont(shrikhandFont, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load Shrikhand font: ", err)
			panic("Font loading error")
		}
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
	return fontFamily, nil
}
