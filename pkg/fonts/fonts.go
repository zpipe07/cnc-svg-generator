package fonts

import (
	_ "embed"
	"errors"
	"log"

	"github.com/tdewolff/canvas"
)

//go:embed ArbutusSlab-Regular.ttf
var ArbutusFont []byte

//go:embed Bungee-Regular.ttf
var bungeeFont []byte

//go:embed Danfo-Regular.ttf
var danfoFont []byte

//go:embed Limelight-Regular.ttf
var limelightFont []byte

//go:embed Shrikhand-Regular.ttf
var shrikhandFont []byte

//go:embed GreatVibes-Regular.ttf
var greatvibesFont []byte

//go:embed PlayfairDisplaySC-Regular.ttf
var playfairFont []byte

//go:embed Playball-Regular.ttf
var PlayballFont []byte

//go:embed AdventPro-Bold.ttf
var AdventPro []byte

//go:embed DMSerifDisplay-Regular.ttf
var DMSerif []byte

//go:embed Sancreek-Regular.ttf
var Sancreek []byte

//go:embed GermaniaOne-Regular.ttf
var GermaniaOne []byte

//go:embed XCompany.ttf
var XCompany []byte

//go:embed Forque.ttf
var Forque []byte

//go:embed Airstream.ttf
var Airstream []byte

//go:embed Gemola.ttf
var Gemola []byte

//go:embed UnicaOne-Regular.ttf
var UnicaOne []byte

//go:embed PoorRichardOpti.otf
var PoorRichard []byte

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
	case "adventpro":
		fontFamily = canvas.NewFontFamily("AdventPro")
		if err := fontFamily.LoadFont(AdventPro, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load AdventPro font: ", err)
			panic("Font loading error")
		}
	case "dmserif":
		fontFamily = canvas.NewFontFamily("DMSerif")
		if err := fontFamily.LoadFont(DMSerif, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load AdventPro font: ", err)
			panic("Font loading error")
		}
	case "sancreek":
		fontFamily = canvas.NewFontFamily("Sancreek")
		if err := fontFamily.LoadFont(Sancreek, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load AdventPro font: ", err)
			panic("Font loading error")
		}
	case "germaniaone":
		fontFamily = canvas.NewFontFamily("GermaniaOne")
		if err := fontFamily.LoadFont(GermaniaOne, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load AdventPro font: ", err)
			panic("Font loading error")
		}
	case "xcompany":
		fontFamily = canvas.NewFontFamily("XCompany")
		if err := fontFamily.LoadFont(XCompany, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load AdventPro font: ", err)
			panic("Font loading error")
		}
	case "forque":
		fontFamily = canvas.NewFontFamily("Forque")
		if err := fontFamily.LoadFont(Forque, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load AdventPro font: ", err)
			panic("Font loading error")
		}
	case "airstream":
		fontFamily = canvas.NewFontFamily("Airstream")
		if err := fontFamily.LoadFont(Airstream, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load AdventPro font: ", err)
			panic("Font loading error")
		}
	case "gemola":
		fontFamily = canvas.NewFontFamily("Gemola")
		if err := fontFamily.LoadFont(Gemola, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load AdventPro font: ", err)
			panic("Font loading error")
		}
	case "unicaone":
		fontFamily = canvas.NewFontFamily("UnicaOne")
		if err := fontFamily.LoadFont(UnicaOne, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load AdventPro font: ", err)
			panic("Font loading error")
		}
	case "poorrichard":
		fontFamily = canvas.NewFontFamily("PoorRichard")
		if err := fontFamily.LoadFont(PoorRichard, 0, canvas.FontRegular); err != nil {
			log.Println("Failed to load AdventPro font: ", err)
			panic("Font loading error")
		}
	default:
		// fontFamily = canvas.NewFontFamily("Bungee")
		// if err := fontFamily.LoadFont(bungeeFont, 0, canvas.FontRegular); err != nil {
		// 	log.Println("Failed to load Bungee font: ", err)
		// 	panic("Font loading error")
		// }
		log.Println("Invalid font family: ", fontFamilyStr)
		return nil, errors.New("invalid font family")
		// return nil, nil
	}
	return fontFamily, nil
}
