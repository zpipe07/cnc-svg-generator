package signs_test

import (
	"strings"
	"testing"

	"cnc-svg-generator/pkg/fonts"
	"cnc-svg-generator/pkg/signs"

	"github.com/tdewolff/canvas"
)

func TestLaurelBorderDecorationsInOutput(t *testing.T) {
	ff := canvas.NewFontFamily("Marigold")
	if err := ff.LoadFont(fonts.Marigold, 0, canvas.FontRegular); err != nil {
		t.Fatal(err)
	}

	// Call drawLaurel via exported path - it's not exported. Use DrawSign with env.
	t.Setenv("LAUREL_PRODUCT_ID", "130")
	svg := signs.DrawSign(nil, "130", "", 15, 11, "black", "white", nil, ff, false)
	if strings.Count(svg, "border-path") != 2 {
		t.Fatalf("expected 2 border-path elements, got %d", strings.Count(svg, "border-path"))
	}
}
