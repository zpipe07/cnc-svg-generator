package signs

import (
	"cnc-svg-generator/pkg/svgutils"
	"fmt"
	"log"

	"github.com/tdewolff/canvas"
)

func drawLargeDonnelly(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// outer edge
	outerEdge, err := canvas.ParseSVGPath("M 1.222 10.75 A 0.25 0.25 0 0 0 1.472 11 L 3.667 11 L 11.333 11 L 13.528 11 A 0.25 0.25 0 0 0 13.778 10.75 L 13.778 9.683 A 0.25 0.25 0 0 1 13.825 9.537 A 4.266 4.266 0 0 0 13.926 9.389 A 5.94 5.94 0 0 0 14.578 8.06 A 7.794 7.794 0 0 0 15 5.5 A 7.963 7.963 0 0 0 14.738 3.457 A 6.229 6.229 0 0 0 13.926 1.611 A 4.946 4.946 0 0 0 13.815 1.451 A 4.033 4.033 0 0 0 13.778 1.4 L 13.778 0.25 A 0.25 0.25 0 0 0 13.528 0 L 11.333 0 L 3.667 0 L 1.472 0 A 0.25 0.25 0 0 0 1.222 0.25 L 1.222 1.4 A 4.033 4.033 0 0 0 1.185 1.451 A 4.946 4.946 0 0 0 1.074 1.611 A 6.229 6.229 0 0 0 0.262 3.457 A 7.963 7.963 0 0 0 0 5.5 A 7.794 7.794 0 0 0 0.422 8.06 A 5.94 5.94 0 0 0 1.074 9.389 A 4.266 4.266 0 0 0 1.175 9.537 A 0.25 0.25 0 0 1 1.222 9.683 L 1.222 10.75 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	builder.EndGroup()

	// rounded edge
	roundedEdge, err := canvas.ParseSVGPath("M 1.422 10.75 A 0.05 0.05 0 0 0 1.472 10.8 L 3.667 10.8 L 11.333 10.8 L 13.528 10.8 A 0.05 0.05 0 0 0 13.578 10.75 L 13.578 9.683 A 0.45 0.45 0 0 1 13.662 9.421 A 4.066 4.066 0 0 0 13.759 9.28 A 5.74 5.74 0 0 0 14.389 7.995 A 7.594 7.594 0 0 0 14.8 5.501 A 0.2 0.2 0 0 1 14.8 5.5 A 7.763 7.763 0 0 0 14.544 3.508 A 0.2 0.2 0 0 1 14.544 3.507 A 6.029 6.029 0 0 0 13.759 1.721 A 4.746 4.746 0 0 0 13.653 1.568 A 0.2 0.2 0 0 1 13.652 1.567 A 3.833 3.833 0 0 0 13.617 1.518 A 0.2 0.2 0 0 1 13.578 1.4 L 13.578 0.25 A 0.05 0.05 0 0 0 13.528 0.2 L 11.333 0.2 L 3.667 0.2 L 1.472 0.2 A 0.05 0.05 0 0 0 1.422 0.25 L 1.422 1.4 A 0.2 0.2 0 0 1 1.383 1.518 A 3.833 3.833 0 0 0 1.348 1.567 A 0.2 0.2 0 0 1 1.347 1.568 A 4.746 4.746 0 0 0 1.241 1.721 A 6.029 6.029 0 0 0 0.456 3.507 A 0.2 0.2 0 0 1 0.456 3.508 A 7.763 7.763 0 0 0 0.2 5.5 A 0.2 0.2 0 0 1 0.2 5.501 A 7.594 7.594 0 0 0 0.611 7.995 A 5.74 5.74 0 0 0 1.241 9.279 A 4.066 4.066 0 0 0 1.338 9.421 A 0.45 0.45 0 0 1 1.422 9.683 L 1.422 10.75 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Rounded Edge", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
	builder.EndGroup()

	// add border
	borderOuter, err := canvas.ParseSVGPath("M 1.722 10.5 L 3.667 10.5 L 11.333 10.5 L 13.278 10.5 L 13.278 9.683 A 0.75 0.75 0 0 1 13.418 9.246 A 3.766 3.766 0 0 0 13.508 9.115 A 5.44 5.44 0 0 0 14.106 7.898 A 7.294 7.294 0 0 0 14.5 5.502 A 0.3 0.3 0 0 1 14.5 5.501 A 0.3 0.3 0 0 1 14.5 5.499 A 7.463 7.463 0 0 0 14.254 3.585 A 0.3 0.3 0 0 1 14.254 3.583 A 0.3 0.3 0 0 1 14.253 3.582 A 5.729 5.729 0 0 0 13.508 1.886 A 4.446 4.446 0 0 0 13.41 1.744 A 0.3 0.3 0 0 1 13.409 1.742 A 0.3 0.3 0 0 1 13.407 1.741 A 3.533 3.533 0 0 0 13.375 1.696 A 0.5 0.5 0 0 1 13.278 1.4 L 13.278 0.5 L 11.333 0.5 L 3.667 0.5 L 1.722 0.5 L 1.722 1.4 A 0.5 0.5 0 0 1 1.625 1.696 A 3.533 3.533 0 0 0 1.593 1.741 A 0.3 0.3 0 0 1 1.591 1.742 A 0.3 0.3 0 0 1 1.59 1.744 A 4.446 4.446 0 0 0 1.492 1.886 A 5.729 5.729 0 0 0 0.747 3.582 A 0.3 0.3 0 0 1 0.746 3.583 A 0.3 0.3 0 0 1 0.746 3.585 A 7.463 7.463 0 0 0 0.5 5.499 A 0.3 0.3 0 0 1 0.5 5.501 A 0.3 0.3 0 0 1 0.5 5.502 A 7.294 7.294 0 0 0 0.894 7.898 A 5.44 5.44 0 0 0 1.492 9.115 A 3.766 3.766 0 0 0 1.582 9.246 A 0.75 0.75 0 0 1 1.722 9.683 L 1.722 10.5 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	borderInner, err := canvas.ParseSVGPath("M 1.922 10.3 L 3.667 10.3 L 11.333 10.3 L 13.078 10.3 L 13.078 9.683 A 0.95 0.95 0 0 1 13.255 9.13 A 3.566 3.566 0 0 0 13.34 9.006 A 5.24 5.24 0 0 0 13.916 7.833 A 7.094 7.094 0 0 0 14.3 5.503 A 0.2 0.2 0 0 1 14.3 5.501 A 0.2 0.2 0 0 1 14.3 5.499 A 7.263 7.263 0 0 0 14.061 3.636 A 0.2 0.2 0 0 1 14.06 3.634 A 0.2 0.2 0 0 1 14.06 3.632 A 5.529 5.529 0 0 0 13.341 1.996 A 4.246 4.246 0 0 0 13.248 1.862 A 0.2 0.2 0 0 1 13.246 1.859 A 0.2 0.2 0 0 1 13.244 1.857 A 3.333 3.333 0 0 0 13.214 1.815 A 0.7 0.7 0 0 1 13.078 1.4 L 13.078 0.7 L 11.333 0.7 L 3.667 0.7 L 1.922 0.7 L 1.922 1.4 A 0.7 0.7 0 0 1 1.786 1.814 A 3.333 3.333 0 0 0 1.756 1.857 A 0.2 0.2 0 0 1 1.754 1.859 A 0.2 0.2 0 0 1 1.752 1.862 A 4.246 4.246 0 0 0 1.659 1.996 A 5.529 5.529 0 0 0 0.94 3.632 A 0.2 0.2 0 0 1 0.94 3.634 A 0.2 0.2 0 0 1 0.939 3.636 A 7.263 7.263 0 0 0 0.7 5.499 A 0.2 0.2 0 0 1 0.7 5.501 A 0.2 0.2 0 0 1 0.7 5.503 A 7.094 7.094 0 0 0 1.084 7.833 A 5.24 5.24 0 0 0 1.659 9.005 A 3.566 3.566 0 0 0 1.744 9.129 A 0.95 0.95 0 0 1 1.922 9.683 L 1.922 10.3 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
}

func drawMedium(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// outer edge
	outerEdge, err := canvas.ParseSVGPath("M 1 8.75 A 0.25 0.25 0 0 0 1.25 9 L 3 9 L 12 9 L 13.75 9 A 0.25 0.25 0 0 0 14 8.75 L 14 7.855 A 3.49 3.49 0 0 0 14.121 7.682 A 4.86 4.86 0 0 0 14.655 6.595 A 6.377 6.377 0 0 0 15 4.5 A 6.515 6.515 0 0 0 14.785 2.829 A 5.096 5.096 0 0 0 14.121 1.318 A 4.046 4.046 0 0 0 14.03 1.187 A 3.3 3.3 0 0 0 14 1.145 L 14 0.25 A 0.25 0.25 0 0 0 13.75 0 L 12 0 L 3 0 L 1.25 0 A 0.25 0.25 0 0 0 1 0.25 L 1 1.145 A 3.3 3.3 0 0 0 0.97 1.187 A 4.046 4.046 0 0 0 0.879 1.318 A 5.096 5.096 0 0 0 0.215 2.829 A 6.515 6.515 0 0 0 0 4.5 A 6.377 6.377 0 0 0 0.345 6.595 A 4.86 4.86 0 0 0 0.879 7.682 A 3.49 3.49 0 0 0 1 7.855 L 1 8.75 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	builder.EndGroup()

	// rounded edge
	roundedEdge, err := canvas.ParseSVGPath("M 1.2 8.75 A 0.05 0.05 0 0 0 1.25 8.8 L 3 8.8 L 12 8.8 L 13.75 8.8 A 0.05 0.05 0 0 0 13.8 8.75 L 13.8 7.855 A 0.2 0.2 0 0 1 13.84 7.736 A 3.29 3.29 0 0 0 13.954 7.573 A 4.66 4.66 0 0 0 14.466 6.53 A 6.177 6.177 0 0 0 14.8 4.501 A 0.2 0.2 0 0 1 14.8 4.5 A 6.315 6.315 0 0 0 14.592 2.88 A 0.2 0.2 0 0 1 14.592 2.879 A 4.896 4.896 0 0 0 13.954 1.428 A 3.846 3.846 0 0 0 13.868 1.304 A 0.2 0.2 0 0 1 13.867 1.303 A 3.1 3.1 0 0 0 13.839 1.264 A 0.2 0.2 0 0 1 13.8 1.145 L 13.8 0.25 A 0.05 0.05 0 0 0 13.75 0.2 L 12 0.2 L 3 0.2 L 1.25 0.2 A 0.05 0.05 0 0 0 1.2 0.25 L 1.2 1.145 A 0.2 0.2 0 0 1 1.161 1.264 A 3.1 3.1 0 0 0 1.133 1.303 A 0.2 0.2 0 0 1 1.132 1.304 A 3.846 3.846 0 0 0 1.046 1.428 A 4.896 4.896 0 0 0 0.408 2.879 A 0.2 0.2 0 0 1 0.408 2.88 A 6.315 6.315 0 0 0 0.2 4.5 A 0.2 0.2 0 0 1 0.2 4.501 A 6.177 6.177 0 0 0 0.534 6.53 A 4.66 4.66 0 0 0 1.046 7.572 A 3.29 3.29 0 0 0 1.16 7.736 A 0.2 0.2 0 0 1 1.2 7.855 L 1.2 8.75 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Rounded Edge", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
	builder.EndGroup()

	// add border
	borderOuter, err := canvas.ParseSVGPath("M 1.5 8.5 L 3 8.5 L 12 8.5 L 13.5 8.5 L 13.5 7.855 A 0.5 0.5 0 0 1 13.599 7.557 A 2.99 2.99 0 0 0 13.703 7.408 A 4.36 4.36 0 0 0 14.182 6.432 A 5.877 5.877 0 0 0 14.5 4.502 A 0.3 0.3 0 0 1 14.5 4.501 A 0.3 0.3 0 0 1 14.5 4.499 A 6.015 6.015 0 0 0 14.302 2.956 A 0.3 0.3 0 0 1 14.302 2.955 A 0.3 0.3 0 0 1 14.301 2.953 A 4.596 4.596 0 0 0 13.704 1.593 A 3.546 3.546 0 0 0 13.625 1.481 A 0.3 0.3 0 0 1 13.624 1.479 A 0.3 0.3 0 0 1 13.623 1.477 A 2.8 2.8 0 0 0 13.597 1.441 A 0.5 0.5 0 0 1 13.5 1.145 L 13.5 0.5 L 12 0.5 L 3 0.5 L 1.5 0.5 L 1.5 1.145 A 0.5 0.5 0 0 1 1.403 1.441 A 2.8 2.8 0 0 0 1.377 1.477 A 0.3 0.3 0 0 1 1.376 1.479 A 0.3 0.3 0 0 1 1.375 1.481 A 3.546 3.546 0 0 0 1.296 1.593 A 4.596 4.596 0 0 0 0.699 2.953 A 0.3 0.3 0 0 1 0.698 2.955 A 0.3 0.3 0 0 1 0.698 2.956 A 6.015 6.015 0 0 0 0.5 4.499 A 0.3 0.3 0 0 1 0.5 4.501 A 0.3 0.3 0 0 1 0.5 4.502 A 5.877 5.877 0 0 0 0.818 6.432 A 4.36 4.36 0 0 0 1.297 7.407 A 2.99 2.99 0 0 0 1.401 7.557 A 0.5 0.5 0 0 1 1.5 7.855 L 1.5 8.5 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	borderInner, err := canvas.ParseSVGPath("M 1.7 8.3 L 3 8.3 L 12 8.3 L 13.3 8.3 L 13.3 7.855 A 0.7 0.7 0 0 1 13.439 7.437 A 2.79 2.79 0 0 0 13.536 7.299 A 4.16 4.16 0 0 0 13.993 6.367 A 5.677 5.677 0 0 0 14.3 4.503 A 0.2 0.2 0 0 1 14.3 4.501 A 0.2 0.2 0 0 1 14.3 4.499 A 5.815 5.815 0 0 0 14.108 3.007 A 0.2 0.2 0 0 1 14.108 3.005 A 0.2 0.2 0 0 1 14.107 3.003 A 4.396 4.396 0 0 0 13.537 1.703 A 3.346 3.346 0 0 0 13.463 1.598 A 0.2 0.2 0 0 1 13.462 1.595 A 0.2 0.2 0 0 1 13.46 1.593 A 2.6 2.6 0 0 0 13.436 1.56 A 0.7 0.7 0 0 1 13.3 1.145 L 13.3 0.7 L 12 0.7 L 3 0.7 L 1.7 0.7 L 1.7 1.145 A 0.7 0.7 0 0 1 1.564 1.56 A 2.6 2.6 0 0 0 1.54 1.593 A 0.2 0.2 0 0 1 1.538 1.595 A 0.2 0.2 0 0 1 1.537 1.598 A 3.346 3.346 0 0 0 1.463 1.703 A 4.396 4.396 0 0 0 0.893 3.003 A 0.2 0.2 0 0 1 0.892 3.005 A 0.2 0.2 0 0 1 0.892 3.007 A 5.815 5.815 0 0 0 0.7 4.499 A 0.2 0.2 0 0 1 0.7 4.501 A 0.2 0.2 0 0 1 0.7 4.503 A 5.677 5.677 0 0 0 1.007 6.367 A 4.16 4.16 0 0 0 1.464 7.298 A 2.79 2.79 0 0 0 1.561 7.437 A 0.7 0.7 0 0 1 1.7 7.855 L 1.7 8.3 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
}

func drawSmall(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// outer edge
	outerEdge, err := canvas.ParseSVGPath("M 0.833 7.25 A 0.25 0.25 0 0 0 1.083 7.5 L 2.5 7.5 L 12.5 7.5 L 13.917 7.5 A 0.25 0.25 0 0 0 14.167 7.25 L 14.167 6.546 A 2.909 2.909 0 0 0 14.268 6.402 C 14.737 5.698 15 4.745 15 3.75 A 5.429 5.429 0 0 0 14.821 2.357 A 4.247 4.247 0 0 0 14.268 1.098 A 3.372 3.372 0 0 0 14.192 0.989 A 2.75 2.75 0 0 0 14.167 0.954 L 14.167 0.25 A 0.25 0.25 0 0 0 13.917 0 L 12.5 0 L 2.5 0 L 1.083 0 A 0.25 0.25 0 0 0 0.833 0.25 L 0.833 0.954 A 2.75 2.75 0 0 0 0.808 0.989 A 3.372 3.372 0 0 0 0.732 1.098 A 4.247 4.247 0 0 0 0.179 2.357 A 5.429 5.429 0 0 0 0 3.75 C 0 4.745 0.263 5.698 0.732 6.402 A 2.909 2.909 0 0 0 0.833 6.546 L 0.833 7.25 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	builder.EndGroup()

	// rounded edge
	roundedEdge, err := canvas.ParseSVGPath("M 1.033 7.25 A 0.05 0.05 0 0 0 1.083 7.3 L 2.5 7.3 L 12.5 7.3 L 13.917 7.3 A 0.05 0.05 0 0 0 13.967 7.25 L 13.967 6.546 A 0.2 0.2 0 0 1 14.006 6.427 A 2.709 2.709 0 0 0 14.1 6.292 A 3.85 3.85 0 0 0 14.523 5.431 A 5.114 5.114 0 0 0 14.8 3.751 A 0.2 0.2 0 0 1 14.8 3.75 A 5.229 5.229 0 0 0 14.628 2.408 A 0.2 0.2 0 0 1 14.627 2.407 A 4.047 4.047 0 0 0 14.101 1.208 A 3.172 3.172 0 0 0 14.03 1.107 A 0.2 0.2 0 0 1 14.029 1.105 A 2.55 2.55 0 0 0 14.006 1.073 A 0.2 0.2 0 0 1 13.967 0.954 L 13.967 0.25 A 0.05 0.05 0 0 0 13.917 0.2 L 12.5 0.2 L 2.5 0.2 L 1.083 0.2 A 0.05 0.05 0 0 0 1.033 0.25 L 1.033 0.954 A 0.2 0.2 0 0 1 0.994 1.073 A 2.55 2.55 0 0 0 0.971 1.105 A 0.2 0.2 0 0 1 0.97 1.107 A 3.172 3.172 0 0 0 0.899 1.208 A 4.047 4.047 0 0 0 0.373 2.407 A 0.2 0.2 0 0 1 0.372 2.408 A 5.229 5.229 0 0 0 0.2 3.75 A 0.2 0.2 0 0 1 0.2 3.751 A 5.114 5.114 0 0 0 0.477 5.431 A 3.85 3.85 0 0 0 0.899 6.292 A 2.709 2.709 0 0 0 0.994 6.427 A 0.2 0.2 0 0 1 1.033 6.546 L 1.033 7.25 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Rounded Edge", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
	builder.EndGroup()

	// add border
	borderOuter, err := canvas.ParseSVGPath("M 1.333 7 L 2.5 7 L 12.5 7 L 13.667 7 L 13.667 6.546 A 0.5 0.5 0 0 1 13.766 6.248 A 2.409 2.409 0 0 0 13.849 6.128 A 3.55 3.55 0 0 0 14.24 5.333 A 4.814 4.814 0 0 0 14.5 3.752 A 0.3 0.3 0 0 1 14.5 3.751 A 0.3 0.3 0 0 1 14.5 3.749 A 4.929 4.929 0 0 0 14.338 2.485 A 0.3 0.3 0 0 1 14.337 2.483 A 0.3 0.3 0 0 1 14.337 2.482 A 3.747 3.747 0 0 0 13.85 1.373 A 2.872 2.872 0 0 0 13.787 1.283 A 0.3 0.3 0 0 1 13.786 1.281 A 0.3 0.3 0 0 1 13.784 1.279 A 2.25 2.25 0 0 0 13.764 1.251 A 0.5 0.5 0 0 1 13.667 0.954 L 13.667 0.5 L 12.5 0.5 L 2.5 0.5 L 1.333 0.5 L 1.333 0.954 A 0.5 0.5 0 0 1 1.236 1.251 A 2.25 2.25 0 0 0 1.216 1.279 A 0.3 0.3 0 0 1 1.214 1.281 A 0.3 0.3 0 0 1 1.213 1.283 A 2.872 2.872 0 0 0 1.15 1.373 A 3.747 3.747 0 0 0 0.663 2.482 A 0.3 0.3 0 0 1 0.663 2.483 A 0.3 0.3 0 0 1 0.662 2.485 A 4.929 4.929 0 0 0 0.5 3.749 A 0.3 0.3 0 0 1 0.5 3.751 A 0.3 0.3 0 0 1 0.5 3.752 A 4.814 4.814 0 0 0 0.76 5.333 A 3.55 3.55 0 0 0 1.15 6.127 A 2.409 2.409 0 0 0 1.234 6.248 A 0.5 0.5 0 0 1 1.333 6.546 L 1.333 7 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	borderInner, err := canvas.ParseSVGPath("M 1.533 6.8 L 2.5 6.8 L 12.5 6.8 L 13.467 6.8 L 13.467 6.546 A 0.7 0.7 0 0 1 13.605 6.128 A 2.209 2.209 0 0 0 13.682 6.018 A 3.35 3.35 0 0 0 14.051 5.268 A 4.614 4.614 0 0 0 14.3 3.753 A 0.2 0.2 0 0 1 14.3 3.751 A 0.2 0.2 0 0 1 14.3 3.749 A 4.729 4.729 0 0 0 14.144 2.535 A 0.2 0.2 0 0 1 14.144 2.533 A 0.2 0.2 0 0 1 14.143 2.532 A 3.547 3.547 0 0 0 13.683 1.483 A 2.672 2.672 0 0 0 13.625 1.4 A 0.2 0.2 0 0 1 13.623 1.397 A 0.2 0.2 0 0 1 13.621 1.395 A 2.05 2.05 0 0 0 13.603 1.369 A 0.7 0.7 0 0 1 13.467 0.954 L 13.467 0.7 L 12.5 0.7 L 2.5 0.7 L 1.533 0.7 L 1.533 0.954 A 0.7 0.7 0 0 1 1.397 1.369 A 2.05 2.05 0 0 0 1.379 1.395 A 0.2 0.2 0 0 1 1.377 1.397 A 0.2 0.2 0 0 1 1.375 1.4 A 2.672 2.672 0 0 0 1.317 1.483 A 3.547 3.547 0 0 0 0.857 2.532 A 0.2 0.2 0 0 1 0.856 2.533 A 0.2 0.2 0 0 1 0.856 2.535 A 4.729 4.729 0 0 0 0.7 3.749 A 0.2 0.2 0 0 1 0.7 3.751 A 0.2 0.2 0 0 1 0.7 3.753 A 4.614 4.614 0 0 0 0.949 5.268 A 3.35 3.35 0 0 0 1.317 6.017 A 2.209 2.209 0 0 0 1.395 6.128 A 0.7 0.7 0 0 1 1.533 6.546 L 1.533 6.8 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
}

func drawExtraSmall(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// outer edge
	outerEdge, err := canvas.ParseSVGPath("M 0.389 3.25 A 0.25 0.25 0 0 0 0.639 3.5 L 1.167 3.5 L 13.833 3.5 L 14.361 3.5 A 0.25 0.25 0 0 0 14.611 3.25 L 14.611 3.055 A 1.357 1.357 0 0 0 14.658 2.987 A 1.89 1.89 0 0 0 14.866 2.565 A 2.48 2.48 0 0 0 15 1.75 A 2.534 2.534 0 0 0 14.917 1.1 A 1.982 1.982 0 0 0 14.658 0.513 A 1.573 1.573 0 0 0 14.623 0.462 A 1.283 1.283 0 0 0 14.611 0.445 L 14.611 0.25 A 0.25 0.25 0 0 0 14.361 0 L 13.833 0 L 1.167 0 L 0.639 0 A 0.25 0.25 0 0 0 0.389 0.25 L 0.389 0.445 A 1.283 1.283 0 0 0 0.377 0.462 A 1.573 1.573 0 0 0 0.342 0.513 A 1.982 1.982 0 0 0 0.083 1.1 A 2.534 2.534 0 0 0 0 1.75 A 2.48 2.48 0 0 0 0.134 2.565 A 1.89 1.89 0 0 0 0.342 2.987 A 1.357 1.357 0 0 0 0.389 3.055 L 0.389 3.25 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Outside", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	builder.EndGroup()

	// rounded edge
	roundedEdge, err := canvas.ParseSVGPath("M 0.589 3.25 A 0.05 0.05 0 0 0 0.639 3.3 L 1.167 3.3 L 13.833 3.3 L 14.361 3.3 A 0.05 0.05 0 0 0 14.411 3.25 L 14.411 3.055 A 0.2 0.2 0 0 1 14.451 2.935 A 1.157 1.157 0 0 0 14.491 2.878 A 1.69 1.69 0 0 0 14.677 2.5 A 2.28 2.28 0 0 0 14.8 1.751 A 0.2 0.2 0 0 1 14.8 1.75 A 2.334 2.334 0 0 0 14.723 1.151 A 0.2 0.2 0 0 1 14.723 1.15 A 1.782 1.782 0 0 0 14.491 0.623 A 1.373 1.373 0 0 0 14.461 0.579 A 0.2 0.2 0 0 1 14.46 0.578 A 1.083 1.083 0 0 0 14.45 0.564 A 0.2 0.2 0 0 1 14.411 0.445 L 14.411 0.25 A 0.05 0.05 0 0 0 14.361 0.2 L 13.833 0.2 L 1.167 0.2 L 0.639 0.2 A 0.05 0.05 0 0 0 0.589 0.25 L 0.589 0.445 A 0.2 0.2 0 0 1 0.55 0.564 A 1.083 1.083 0 0 0 0.54 0.578 A 0.2 0.2 0 0 1 0.539 0.579 A 1.373 1.373 0 0 0 0.509 0.623 A 1.782 1.782 0 0 0 0.277 1.15 A 0.2 0.2 0 0 1 0.277 1.151 A 2.334 2.334 0 0 0 0.2 1.749 A 0.2 0.2 0 0 1 0.2 1.751 A 2.28 2.28 0 0 0 0.323 2.5 A 1.69 1.69 0 0 0 0.509 2.877 A 1.157 1.157 0 0 0 0.549 2.935 A 0.2 0.2 0 0 1 0.589 3.055 L 0.589 3.25 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Round", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
	builder.EndGroup()

	// add border
	borderOuter, err := canvas.ParseSVGPath("M 0.886 3 L 1.167 3 L 13.833 3 L 14.114 3 A 0.5 0.5 0 0 1 14.21 2.756 A 0.857 0.857 0 0 0 14.24 2.714 A 1.39 1.39 0 0 0 14.393 2.402 A 1.98 1.98 0 0 0 14.5 1.752 A 0.3 0.3 0 0 1 14.5 1.75 A 0.3 0.3 0 0 1 14.5 1.749 A 2.034 2.034 0 0 0 14.433 1.227 A 0.3 0.3 0 0 1 14.433 1.226 A 0.3 0.3 0 0 1 14.432 1.225 A 1.482 1.482 0 0 0 14.241 0.788 A 1.073 1.073 0 0 0 14.218 0.755 A 0.3 0.3 0 0 1 14.217 0.753 A 0.3 0.3 0 0 1 14.215 0.751 A 0.783 0.783 0 0 0 14.208 0.742 A 0.5 0.5 0 0 1 14.114 0.5 L 13.833 0.5 L 1.167 0.5 L 0.886 0.5 A 0.5 0.5 0 0 1 0.792 0.742 A 0.783 0.783 0 0 0 0.785 0.751 A 0.3 0.3 0 0 1 0.783 0.753 A 0.3 0.3 0 0 1 0.782 0.755 A 1.073 1.073 0 0 0 0.759 0.788 A 1.482 1.482 0 0 0 0.568 1.225 A 0.3 0.3 0 0 1 0.567 1.226 A 0.3 0.3 0 0 1 0.567 1.227 A 2.034 2.034 0 0 0 0.5 1.749 A 0.3 0.3 0 0 1 0.5 1.75 A 0.3 0.3 0 0 1 0.5 1.752 A 1.98 1.98 0 0 0 0.607 2.402 A 1.39 1.39 0 0 0 0.76 2.713 A 0.857 0.857 0 0 0 0.79 2.756 A 0.5 0.5 0 0 1 0.886 3 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	borderInner, err := canvas.ParseSVGPath("M 1.041 2.8 L 1.167 2.8 L 13.833 2.8 L 13.959 2.8 A 0.7 0.7 0 0 1 14.05 2.637 A 0.657 0.657 0 0 0 14.073 2.604 A 1.19 1.19 0 0 0 14.204 2.337 A 1.78 1.78 0 0 0 14.3 1.753 A 0.2 0.2 0 0 1 14.3 1.751 A 0.2 0.2 0 0 1 14.3 1.748 A 1.834 1.834 0 0 0 14.24 1.278 A 0.2 0.2 0 0 1 14.239 1.276 A 0.2 0.2 0 0 1 14.239 1.274 A 1.282 1.282 0 0 0 14.074 0.898 A 0.873 0.873 0 0 0 14.056 0.872 A 0.2 0.2 0 0 1 14.054 0.87 A 0.2 0.2 0 0 1 14.052 0.867 A 0.583 0.583 0 0 0 14.047 0.86 A 0.7 0.7 0 0 1 13.959 0.7 L 13.833 0.7 L 1.167 0.7 L 1.041 0.7 A 0.7 0.7 0 0 1 0.953 0.86 A 0.583 0.583 0 0 0 0.948 0.867 A 0.2 0.2 0 0 1 0.946 0.87 A 0.2 0.2 0 0 1 0.944 0.872 A 0.873 0.873 0 0 0 0.926 0.898 A 1.282 1.282 0 0 0 0.761 1.274 A 0.2 0.2 0 0 1 0.761 1.276 A 0.2 0.2 0 0 1 0.76 1.278 A 1.834 1.834 0 0 0 0.7 1.748 A 0.2 0.2 0 0 1 0.7 1.751 A 0.2 0.2 0 0 1 0.7 1.753 A 1.78 1.78 0 0 0 0.796 2.337 A 1.19 1.19 0 0 0 0.927 2.603 A 0.657 0.657 0 0 0 0.95 2.637 A 0.7 0.7 0 0 1 1.041 2.8 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
}

func drawExtraSmallVertical(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// outer edge
	outerEdge, err := canvas.ParseSVGPath("M 3.25 14.611 A 0.25 0.25 0 0 0 3.5 14.361 L 3.5 13.833 L 3.5 1.167 L 3.5 0.639 A 0.25 0.25 0 0 0 3.25 0.389 L 3.055 0.389 A 1.357 1.357 0 0 0 2.987 0.342 A 1.89 1.89 0 0 0 2.565 0.134 A 2.48 2.48 0 0 0 1.75 0 A 2.534 2.534 0 0 0 1.1 0.083 A 1.982 1.982 0 0 0 0.513 0.342 A 1.573 1.573 0 0 0 0.462 0.377 A 1.283 1.283 0 0 0 0.445 0.389 L 0.25 0.389 A 0.25 0.25 0 0 0 0 0.639 L 0 1.167 L 0 13.833 L 0 14.361 A 0.25 0.25 0 0 0 0.25 14.611 L 0.445 14.611 A 1.283 1.283 0 0 0 0.462 14.623 A 1.573 1.573 0 0 0 0.513 14.658 A 1.982 1.982 0 0 0 1.1 14.917 A 2.534 2.534 0 0 0 1.75 15 A 2.48 2.48 0 0 0 2.565 14.866 A 1.89 1.89 0 0 0 2.987 14.658 A 1.357 1.357 0 0 0 3.055 14.611 L 3.25 14.611 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Outside", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	builder.EndGroup()

	// rounded edge
	roundedEdge, err := canvas.ParseSVGPath("M 3.25 14.411 A 0.05 0.05 0 0 0 3.3 14.361 L 3.3 13.833 L 3.3 1.167 L 3.3 0.639 A 0.05 0.05 0 0 0 3.25 0.589 L 3.055 0.589 A 0.2 0.2 0 0 1 2.935 0.549 A 1.157 1.157 0 0 0 2.878 0.509 A 1.69 1.69 0 0 0 2.5 0.323 A 2.28 2.28 0 0 0 1.751 0.2 A 0.2 0.2 0 0 1 1.75 0.2 A 2.334 2.334 0 0 0 1.151 0.277 A 0.2 0.2 0 0 1 1.15 0.277 A 1.782 1.782 0 0 0 0.623 0.509 A 1.373 1.373 0 0 0 0.579 0.539 A 0.2 0.2 0 0 1 0.578 0.54 A 1.083 1.083 0 0 0 0.564 0.55 A 0.2 0.2 0 0 1 0.445 0.589 L 0.25 0.589 A 0.05 0.05 0 0 0 0.2 0.639 L 0.2 1.167 L 0.2 13.833 L 0.2 14.361 A 0.05 0.05 0 0 0 0.25 14.411 L 0.445 14.411 A 0.2 0.2 0 0 1 0.564 14.45 A 1.083 1.083 0 0 0 0.578 14.46 A 0.2 0.2 0 0 1 0.579 14.461 A 1.373 1.373 0 0 0 0.623 14.491 A 1.782 1.782 0 0 0 1.15 14.723 A 0.2 0.2 0 0 1 1.151 14.723 A 2.334 2.334 0 0 0 1.749 14.8 A 0.2 0.2 0 0 1 1.751 14.8 A 2.28 2.28 0 0 0 2.5 14.677 A 1.69 1.69 0 0 0 2.877 14.491 A 1.157 1.157 0 0 0 2.935 14.451 A 0.2 0.2 0 0 1 3.055 14.411 L 3.25 14.411 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Round", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
	builder.EndGroup()

	// add border
	borderOuter, err := canvas.ParseSVGPath("M 3 14.114 L 3 13.833 L 3 1.167 L 3 0.886 A 0.5 0.5 0 0 1 2.756 0.79 A 0.857 0.857 0 0 0 2.714 0.76 A 1.39 1.39 0 0 0 2.402 0.607 A 1.98 1.98 0 0 0 1.752 0.5 A 0.3 0.3 0 0 1 1.75 0.5 A 0.3 0.3 0 0 1 1.749 0.5 A 2.034 2.034 0 0 0 1.227 0.567 A 0.3 0.3 0 0 1 1.226 0.567 A 0.3 0.3 0 0 1 1.225 0.568 A 1.482 1.482 0 0 0 0.788 0.759 A 1.073 1.073 0 0 0 0.755 0.782 A 0.3 0.3 0 0 1 0.753 0.783 A 0.3 0.3 0 0 1 0.751 0.785 A 0.783 0.783 0 0 0 0.742 0.792 A 0.5 0.5 0 0 1 0.5 0.886 L 0.5 1.167 L 0.5 13.833 L 0.5 14.114 A 0.5 0.5 0 0 1 0.742 14.208 A 0.783 0.783 0 0 0 0.751 14.215 A 0.3 0.3 0 0 1 0.753 14.217 A 0.3 0.3 0 0 1 0.755 14.218 A 1.073 1.073 0 0 0 0.788 14.241 A 1.482 1.482 0 0 0 1.225 14.432 A 0.3 0.3 0 0 1 1.226 14.433 A 0.3 0.3 0 0 1 1.227 14.433 A 2.034 2.034 0 0 0 1.749 14.5 A 0.3 0.3 0 0 1 1.75 14.5 A 0.3 0.3 0 0 1 1.752 14.5 A 1.98 1.98 0 0 0 2.402 14.393 A 1.39 1.39 0 0 0 2.713 14.24 A 0.857 0.857 0 0 0 2.756 14.21 A 0.5 0.5 0 0 1 3 14.114 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
	})
	borderInner, err := canvas.ParseSVGPath("M 2.8 13.959 L 2.8 13.833 L 2.8 1.167 L 2.8 1.041 A 0.7 0.7 0 0 1 2.637 0.95 A 0.657 0.657 0 0 0 2.604 0.927 A 1.19 1.19 0 0 0 2.337 0.796 A 1.78 1.78 0 0 0 1.753 0.7 A 0.2 0.2 0 0 1 1.751 0.7 A 0.2 0.2 0 0 1 1.748 0.7 A 1.834 1.834 0 0 0 1.278 0.76 A 0.2 0.2 0 0 1 1.276 0.761 A 0.2 0.2 0 0 1 1.274 0.761 A 1.282 1.282 0 0 0 0.898 0.926 A 0.873 0.873 0 0 0 0.872 0.944 A 0.2 0.2 0 0 1 0.87 0.946 A 0.2 0.2 0 0 1 0.867 0.948 A 0.583 0.583 0 0 0 0.86 0.953 A 0.7 0.7 0 0 1 0.7 1.041 L 0.7 1.167 L 0.7 13.833 L 0.7 13.959 A 0.7 0.7 0 0 1 0.86 14.047 A 0.583 0.583 0 0 0 0.867 14.052 A 0.2 0.2 0 0 1 0.87 14.054 A 0.2 0.2 0 0 1 0.872 14.056 A 0.873 0.873 0 0 0 0.898 14.074 A 1.282 1.282 0 0 0 1.274 14.239 A 0.2 0.2 0 0 1 1.276 14.239 A 0.2 0.2 0 0 1 1.278 14.24 A 1.834 1.834 0 0 0 1.748 14.3 A 0.2 0.2 0 0 1 1.751 14.3 A 0.2 0.2 0 0 1 1.753 14.3 A 1.78 1.78 0 0 0 2.337 14.204 A 1.19 1.19 0 0 0 2.603 14.073 A 0.657 0.657 0 0 0 2.637 14.05 A 0.7 0.7 0 0 1 2.8 13.959 Z")
	if err != nil {
		log.Fatalf("Error parsing SVG path: %v", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
	})
}

func drawDonnelly(
	size string,
	width float64,
	height float64,
	foregroundColor string,
	backgroundColor string,
	lines []string,
	fontFamily *canvas.FontFamily,
) string {
	// Initialize SVG builder
	builder := svgutils.NewSVGBuilder(width, height)

	if size == "large" {
		drawLargeDonnelly(builder, foregroundColor, backgroundColor)
	} else if size == "medium" {
		drawMedium(builder, foregroundColor, backgroundColor)
	} else if size == "small" {
		drawSmall(builder, foregroundColor, backgroundColor)
	} else if size == "extra small" {
		drawExtraSmall(builder, foregroundColor, backgroundColor)
	} else if size == "extra small vertical" {
		drawExtraSmallVertical(builder, foregroundColor, backgroundColor)
	} else {
		// log.Fatalf("Invalid size: %s", size)
		log.Printf("Invalid size: %s", size)
	}

	// add text
	numLines := len(lines)

	if numLines > 0 {
		// calculate total height of text containers
		const padding = 1.0
		const lineSpacing = 0.5
		availableHeight := height - 2*padding - float64(numLines-1)*lineSpacing

		// determine heights for each line
		var containerHeights []float64
		switch numLines {
		case 3:
			containerHeights = []float64{
				availableHeight * 0.25, // First line
				availableHeight * 0.5,  // Middle line (taller)
				availableHeight * 0.25, // Last line
			}
		case 2:
			containerHeights = []float64{
				availableHeight * 0.7, // First line (taller)
				availableHeight * 0.3, // Last line
			}
		default:
			containerHeights = []float64{availableHeight} // Single line
		}

		// Calculate the starting y position for the topmost line
		currentY := height - padding

		for i, line := range lines {
			containerHeight := containerHeights[i]
			container := canvas.Rectangle(width-4.25, containerHeight)
			containerBounds := container.Bounds()
			containerX := width/2 - containerBounds.W()/2
			containerY := currentY - containerHeight

			// Position the container
			container = container.Translate(containerX, containerY)
			container = container.Scale(1, -1)
			container = container.Translate(0, height)

			builder.AddPath(container.ToSVG(), map[string]string{
				"fill":         "none",
				"stroke":       "pink",
				"stroke-width": "0.025",
				"id":           fmt.Sprintf("text-container-%d", i),
			})

			// Draw text
			fontSize := 1.0
			face := fontFamily.Face(fontSize, canvas.FontRegular, canvas.FontNormal)
			textPath, _, err := face.ToPath(line)
			if err != nil {
				log.Fatalf("Failed to convert text to path: %s", err)
			}
			textBounds := textPath.Bounds()
			metrics := face.Metrics()
			ascent := metrics.Ascent
			descent := metrics.Descent

			// Scale the text to fit within the container
			scale := min(containerBounds.W()/textBounds.W(), containerBounds.H()/textBounds.H())
			textPath.Scale(scale, scale)

			// Recalculate text position after scaling
			textBounds = textPath.Bounds()
			x := containerX +
				containerBounds.W()/2 -
				textBounds.W()/2
			y := containerY +
				containerBounds.H()/2 +
				descent*scale -
				((ascent + descent) / 2 * scale)
			textPath = textPath.Translate(x, y)
			textPath = textPath.Scale(1, -1)
			textPath = textPath.Translate(0, height)

			builder.AddPath(textPath.ToSVG(), map[string]string{
				"fill": backgroundColor,
				"id":   fmt.Sprintf("text-line-%d", i),
			})

			// Update currentY for the next line
			currentY -= containerHeight + lineSpacing
		}
	}

	builder.EndGroup()

	return builder.Close()
}
