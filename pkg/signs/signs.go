package signs

import (
	"os"

	"github.com/tdewolff/canvas"
)

func DrawSign(
	// ctx *canvas.Context,
	// productConfig svg.ProductConfig,
	productId string,
	width float64,
	height float64,
	// foregroundColor color.RGBA,
	// backgroundColor color.RGBA,
	foregroundColor string,
	backgroundColor string,
	lines []string,
	fontFamily *canvas.FontFamily,
) string {

	// if productId == os.Getenv("ELLIPSE_PRODUCT_ID") {
	// 	return drawEllipse(ctx, width, height, foregroundColor, backgroundColor, lines, fontFamily)
	// } else if productId == os.Getenv("RECTANGLE_PRODUCT_ID") {
	// 	return drawRectangle(ctx, width, height, foregroundColor, backgroundColor, lines, fontFamily)
	// } else if productId == os.Getenv("DECO_PRODUCT_ID") {
	// 	return drawDeco(ctx, width, height, foregroundColor, backgroundColor, lines)
	// } else if productId == os.Getenv("ALDER_PRODUCT_ID") {
	// 	return drawAlder(ctx, width, height, foregroundColor, backgroundColor, lines, fontFamily)
	// } else if productId == os.Getenv("FLEUR_PRODUCT_ID") {
	// 	return drawFleur(ctx, width, height, foregroundColor, backgroundColor, lines, fontFamily)
	// } else if productId == os.Getenv("CEZAR_PRODUCT_ID") {
	// 	return drawCezar(ctx, width, height, foregroundColor, backgroundColor, lines)
	// } else if productId == os.Getenv("RECURSO_PRODUCT_ID") {
	// 	return drawRecurso(ctx, width, height, foregroundColor, backgroundColor, lines)
	// } else if productId == os.Getenv("SESAME_PRODUCT_ID") {
	// 	return drawSesame(ctx, width, height, foregroundColor, backgroundColor, lines, fontFamily)
	// } else {
	// 	panic("Invalid product ID: " + productId)
	// }
	switch productId {
	case os.Getenv("ELLIPSE_PRODUCT_ID"):
		return drawEllipse(width, height, foregroundColor, backgroundColor, lines, fontFamily)
	case os.Getenv("RECTANGLE_PRODUCT_ID"):
		return drawRectangle(width, height, foregroundColor, backgroundColor, lines, fontFamily)
	// case os.Getenv("DECO_PRODUCT_ID"):
	// 	drawDeco(width, height, foregroundColor, backgroundColor, lines)
	// case os.Getenv("ALDER_PRODUCT_ID"):
	// 	drawAlder(width, height, foregroundColor, backgroundColor, lines, fontFamily)
	// case os.Getenv("FLEUR_PRODUCT_ID"):
	// 	drawFleur(width, height, foregroundColor, backgroundColor, lines, fontFamily)
	// case os.Getenv("CEZAR_PRODUCT_ID"):
	// 	drawCezar(width, height, foregroundColor, backgroundColor, lines)
	// case os.Getenv("RECURSO_PRODUCT_ID"):
	// 	drawRecurso(width, height, foregroundColor, backgroundColor, lines)
	case os.Getenv("SESAME_PRODUCT_ID"):
		return drawSesame(width, height, foregroundColor, backgroundColor, lines, fontFamily)
	default:
		return "Invalid product ID: " + productId
	}
}
