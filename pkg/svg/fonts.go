package svg

import (
	_ "embed"
	"log"

	"github.com/tdewolff/canvas"
)

//go:embed fonts/ArbutusSlab-Regular.ttf
var ArbutusFont []byte

//go:embed fonts/Bungee-Regular.ttf
var bungeeFont []byte

//go:embed fonts/Danfo-Regular.ttf
var danfoFont []byte

//go:embed fonts/Limelight-Regular.ttf
var limelightFont []byte

//go:embed fonts/Shrikhand-Regular.ttf
var shrikhandFont []byte

//go:embed fonts/GreatVibes-Regular.ttf
var greatvibesFont []byte

//go:embed fonts/PlayfairDisplaySC-Regular.ttf
var playfairFont []byte

//go:embed fonts/Playball-Regular.ttf
var PlayballFont []byte

func ParseFontFamily(fontFamilyStr string) (*canvas.FontFamily, error) {
	var fontFamily *canvas.FontFamily
	switch fontFamilyStr {
	case "arbutus":
		fontFamily = canvas.NewFontFamily("Arbutus")
		if err := fontFamily.LoadFont(ArbutusFont, 0, canvas.FontRegular); err != nil {
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
		fontFamily = canvas.NewFontFamily("Bungee")
		if err := fontFamily.LoadFont(bungeeFont, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load Bungee font: ", err)
			panic("Font loading error")
		}
	case "greatvibes":
		fontFamily = canvas.NewFontFamily("GreatVibes")
		if err := fontFamily.LoadFont(greatvibesFont, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load GreatVibes font: ", err)
			panic("Font loading error")
		}
	case "playfair":
		fontFamily = canvas.NewFontFamily("Playfair")
		if err := fontFamily.LoadFont(playfairFont, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load Playfair font: ", err)
			panic("Font loading error")
		}
	default:
		// fontFamily = canvas.NewFontFamily("Bungee")
		// if err := fontFamily.LoadFont(bungeeFont, 0, canvas.FontRegular); err != nil {
		// 	log.Println("Failed to load Bungee font: ", err)
		// 	panic("Font loading error")
		// }
		// log.Println("Invalid font family: ", fontFamilyStr)
		// return nil, errors.New("invalid font family")
		return nil, nil
	}
	return fontFamily, nil
}
