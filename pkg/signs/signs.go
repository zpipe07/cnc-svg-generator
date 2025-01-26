package signs

import (
	"github.com/tdewolff/canvas"
)

func DrawSign(
	ctx *canvas.Context,
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
	return drawSesame(
		// ctx,
		width,
		height,
		foregroundColor,
		backgroundColor,
		lines,
		fontFamily,
	)
}
