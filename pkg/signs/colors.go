package signs

// ColorMap contains the mapping of color names to their hex values
var ColorMap = map[string]string{
	"black": "#000000",
	"white": "#FFFFFF",
	"green": "#013801",
	"tan":   "#e5c498",
	"brown": "#552a0b",
	"blue":  "#0000e4",
	"grey":  "#D6DAD2",
}

// GetColor returns the hex value for a given color name
func GetColor(colorName string) string {
	if color, exists := ColorMap[colorName]; exists {
		return color
	}
	return ColorMap["black"] // Default to black if color not found
}
