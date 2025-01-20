package signs

import (
	"image/color"
	"os"

	"github.com/tdewolff/canvas"
)

func DrawSign(
	ctx *canvas.Context,
	// productConfig svg.ProductConfig,
	productId string,
	width float64,
	height float64,
	foregroundColor color.RGBA,
	backgroundColor color.RGBA,
	lines []string,
	fontFamily *canvas.FontFamily,
) {

	if productId == os.Getenv("ELLIPSE_PRODUCT_ID") {
		drawEllipse(ctx, width, height, foregroundColor, backgroundColor)
	} else if productId == os.Getenv("RECTANGLE_PRODUCT_ID") {

		drawRectangle(ctx, width, height, foregroundColor, backgroundColor, lines, fontFamily)
	} else if productId == os.Getenv("DECO_PRODUCT_ID") {
		drawDeco(ctx, width, height, foregroundColor, backgroundColor)
	} else if productId == os.Getenv("ALDER_PRODUCT_ID") {
		drawAlder(ctx, width, height, foregroundColor, backgroundColor, lines, fontFamily)
	} else if productId == os.Getenv("FLEUR_PRODUCT_ID") {
		drawFleur(ctx, width, height, foregroundColor, backgroundColor, lines, fontFamily)
	} else if productId == os.Getenv("CEZAR_PRODUCT_ID") {
		drawCezar(ctx, width, height, foregroundColor, backgroundColor, lines)
	} else if productId == os.Getenv("RECURSO_PRODUCT_ID") {
		drawRecurso(ctx, width, height, foregroundColor, backgroundColor, lines)
	} else {
		panic("Invalid product ID: " + productId)
	}
}
