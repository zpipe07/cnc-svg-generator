package signs

import (
	"cnc-svg-generator/pkg/svgutils"
	"fmt"
	"log"

	"github.com/tdewolff/canvas"
)

func drawMediumSora(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// Add the outer edge
	outerEdge, err := canvas.ParseSVGPath("M 0 8.75 A 0.25 0.25 0 0 0 0.25 9 L 14.75 9 A 0.25 0.25 0 0 0 15 8.75 L 15 2.25 C 15 1.653 14.21 1.081 12.803 0.659 A 13.532 13.532 0 0 0 11.192 0.292 C 10.076 0.102 8.805 0 7.5 0 C 5.511 0 3.603 0.237 2.197 0.659 A 8.02 8.02 0 0 0 1.391 0.945 C 0.587 1.283 0.109 1.683 0.017 2.101 A 0.689 0.689 0 0 0 0 2.25 L 0 8.75 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "outer-edge",
	})
	builder.EndGroup()

	// Add the rounded edge
	roundedEdge, err := canvas.ParseSVGPath("M 0.2 8.75 A 0.05 0.05 0 0 0 0.25 8.8 L 14.75 8.8 A 0.05 0.05 0 0 0 14.8 8.75 L 14.8 2.254 A 0.66 0.66 0 0 0 14.622 1.832 A 0.2 0.2 0 0 1 14.618 1.827 A 1.859 1.859 0 0 0 14.227 1.492 A 0.2 0.2 0 0 1 14.225 1.491 A 4.015 4.015 0 0 0 13.64 1.176 A 7.25 7.25 0 0 0 12.871 0.889 A 9.567 9.567 0 0 0 12.747 0.851 A 13.332 13.332 0 0 0 11.16 0.489 A 20.937 20.937 0 0 0 8.73 0.231 A 24.523 24.523 0 0 0 7.5 0.2 A 23.29 23.29 0 0 0 4.654 0.37 A 16.259 16.259 0 0 0 2.794 0.704 A 11.123 11.123 0 0 0 2.254 0.851 A 7.82 7.82 0 0 0 1.467 1.13 A 4.48 4.48 0 0 0 0.865 1.433 A 0.2 0.2 0 0 1 0.864 1.434 A 2.165 2.165 0 0 0 0.437 1.766 A 0.2 0.2 0 0 1 0.434 1.769 A 0.815 0.815 0 0 0 0.211 2.147 A 0.489 0.489 0 0 0 0.2 2.25 L 0.2 8.75 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.StartGroup("Round", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "rounded-edge",
	})
	builder.EndGroup()

	// Add the border
	borderOuter, err := canvas.ParseSVGPath("M 0.5 8.5 L 14.5 8.5 L 14.5 2.262 A 0.36 0.36 0 0 0 14.403 2.037 A 0.5 0.5 0 0 1 14.393 2.026 A 1.559 1.559 0 0 0 14.065 1.745 A 0.3 0.3 0 0 1 14.063 1.744 A 0.3 0.3 0 0 1 14.062 1.743 A 3.715 3.715 0 0 0 13.521 1.451 A 6.95 6.95 0 0 0 12.782 1.176 A 0.3 0.3 0 0 1 12.781 1.175 A 0.3 0.3 0 0 1 12.78 1.175 A 9.267 9.267 0 0 0 12.661 1.138 A 13.032 13.032 0 0 0 11.111 0.785 A 20.637 20.637 0 0 0 8.716 0.53 A 24.223 24.223 0 0 0 7.5 0.5 A 22.99 22.99 0 0 0 4.69 0.668 A 15.959 15.959 0 0 0 2.865 0.995 A 10.823 10.823 0 0 0 2.34 1.138 A 7.52 7.52 0 0 0 1.583 1.407 A 4.18 4.18 0 0 0 1.02 1.69 A 0.3 0.3 0 0 1 1.019 1.691 A 0.3 0.3 0 0 1 1.017 1.692 A 1.865 1.865 0 0 0 0.649 1.978 A 0.3 0.3 0 0 1 0.645 1.982 A 0.3 0.3 0 0 1 0.642 1.985 A 0.515 0.515 0 0 0 0.503 2.217 A 0.189 0.189 0 0 0 0.5 2.25 L 0.5 8.5 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "my-custom-id",
	})
	borderInner, err := canvas.ParseSVGPath("M 0.7 8.3 L 14.3 8.3 L 14.3 2.269 A 0.16 0.16 0 0 0 14.258 2.174 A 0.7 0.7 0 0 1 14.243 2.158 A 1.359 1.359 0 0 0 13.957 1.913 A 0.2 0.2 0 0 1 13.955 1.912 A 0.2 0.2 0 0 1 13.953 1.91 A 3.515 3.515 0 0 0 13.441 1.635 A 0.2 0.2 0 0 1 13.44 1.634 A 0.2 0.2 0 0 1 13.439 1.634 A 6.75 6.75 0 0 0 12.723 1.367 A 0.2 0.2 0 0 1 12.722 1.366 A 0.2 0.2 0 0 1 12.72 1.366 A 9.067 9.067 0 0 0 12.605 1.33 A 12.832 12.832 0 0 0 11.079 0.982 A 20.437 20.437 0 0 0 8.706 0.73 A 0.2 0.2 0 0 1 8.705 0.73 A 24.023 24.023 0 0 0 7.5 0.7 A 22.79 22.79 0 0 0 4.714 0.866 A 15.759 15.759 0 0 0 2.912 1.19 A 10.623 10.623 0 0 0 2.397 1.33 A 0.2 0.2 0 0 1 2.396 1.33 A 0.2 0.2 0 0 1 2.395 1.33 A 7.32 7.32 0 0 0 1.659 1.591 A 3.98 3.98 0 0 0 1.123 1.861 A 0.2 0.2 0 0 1 1.121 1.863 A 0.2 0.2 0 0 1 1.119 1.864 A 1.665 1.665 0 0 0 0.791 2.119 A 0.5 0.5 0 0 1 0.786 2.124 A 0.5 0.5 0 0 1 0.78 2.13 A 0.315 0.315 0 0 0 0.7 2.257 L 0.7 8.3 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "my-custom-id",
	})
}

func drawLargeSora(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// Add the outer edge
	outerEdge, err := canvas.ParseSVGPath("M 0 10.75 A 0.25 0.25 0 0 0 0.25 11 L 14.75 11 A 0.25 0.25 0 0 0 15 10.75 L 15 2.75 A 1.224 1.224 0 0 0 14.715 1.999 C 14.462 1.672 14.046 1.365 13.484 1.092 A 7.384 7.384 0 0 0 12.803 0.806 A 11.704 11.704 0 0 0 11.052 0.328 C 10.24 0.168 9.352 0.064 8.431 0.021 A 20.337 20.337 0 0 0 7.5 0 C 5.511 0 3.603 0.29 2.197 0.806 A 6.973 6.973 0 0 0 1.325 1.189 C 0.776 1.481 0.387 1.809 0.178 2.154 A 1.143 1.143 0 0 0 0 2.75 L 0 10.75 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "outer-edge",
	})
	builder.EndGroup()

	// Add the rounded edge
	roundedEdge, err := canvas.ParseSVGPath("M 0.2 10.75 A 0.05 0.05 0 0 0 0.25 10.8 L 14.75 10.8 A 0.05 0.05 0 0 0 14.8 10.75 L 14.8 2.753 A 1.024 1.024 0 0 0 14.561 2.127 A 0.2 0.2 0 0 1 14.559 2.124 A 2.353 2.353 0 0 0 14.07 1.667 A 0.2 0.2 0 0 1 14.069 1.666 A 4.536 4.536 0 0 0 13.398 1.273 A 7.184 7.184 0 0 0 12.736 0.994 A 11.504 11.504 0 0 0 11.014 0.524 A 17.71 17.71 0 0 0 8.422 0.221 A 20.137 20.137 0 0 0 7.5 0.2 A 18.943 18.943 0 0 0 4.503 0.432 A 13.048 13.048 0 0 0 2.641 0.865 A 9.249 9.249 0 0 0 2.265 0.993 A 6.773 6.773 0 0 0 1.418 1.367 A 3.882 3.882 0 0 0 0.777 1.784 A 0.2 0.2 0 0 1 0.776 1.785 A 1.909 1.909 0 0 0 0.347 2.261 A 0.943 0.943 0 0 0 0.2 2.752 L 0.2 10.75 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.StartGroup("Rounded Edge", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "rounded-edge",
	})
	builder.EndGroup()

	// Add the border
	borderOuter, err := canvas.ParseSVGPath("M 0.5 10.5 L 14.5 10.5 L 14.5 2.758 A 0.724 0.724 0 0 0 14.331 2.32 A 0.3 0.3 0 0 1 14.328 2.316 A 0.3 0.3 0 0 1 14.325 2.312 A 2.053 2.053 0 0 0 13.899 1.913 A 0.3 0.3 0 0 1 13.897 1.912 A 0.3 0.3 0 0 1 13.896 1.911 A 4.236 4.236 0 0 0 13.269 1.544 A 0.3 0.3 0 0 1 13.269 1.543 A 0.3 0.3 0 0 1 13.268 1.543 A 6.884 6.884 0 0 0 12.634 1.276 A 11.204 11.204 0 0 0 10.958 0.819 A 17.41 17.41 0 0 0 8.409 0.521 A 19.837 19.837 0 0 0 7.5 0.5 A 18.643 18.643 0 0 0 4.55 0.728 A 12.748 12.748 0 0 0 2.731 1.151 A 8.949 8.949 0 0 0 2.369 1.275 A 0.3 0.3 0 0 1 2.367 1.275 A 0.3 0.3 0 0 1 2.366 1.276 A 6.473 6.473 0 0 0 1.557 1.632 A 0.3 0.3 0 0 1 1.555 1.633 A 3.582 3.582 0 0 0 0.965 2.018 A 0.3 0.3 0 0 1 0.964 2.02 A 0.3 0.3 0 0 1 0.962 2.021 A 1.609 1.609 0 0 0 0.6 2.422 A 0.643 0.643 0 0 0 0.5 2.754 L 0.5 10.5 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "my-custom-id",
	})
	borderInner, err := canvas.ParseSVGPath("M 0.7 10.3 L 14.3 10.3 L 14.3 2.762 A 0.524 0.524 0 0 0 14.178 2.448 A 0.5 0.5 0 0 1 14.173 2.442 A 0.5 0.5 0 0 1 14.169 2.437 A 1.853 1.853 0 0 0 13.784 2.077 A 0.2 0.2 0 0 1 13.782 2.076 A 0.2 0.2 0 0 1 13.78 2.074 A 4.036 4.036 0 0 0 13.183 1.724 A 0.2 0.2 0 0 1 13.182 1.724 A 0.2 0.2 0 0 1 13.181 1.723 A 6.684 6.684 0 0 0 12.566 1.464 A 11.004 11.004 0 0 0 10.921 1.016 A 17.21 17.21 0 0 0 8.401 0.721 A 0.2 0.2 0 0 1 8.4 0.721 A 0.2 0.2 0 0 1 8.399 0.721 A 19.637 19.637 0 0 0 7.5 0.7 A 18.443 18.443 0 0 0 4.581 0.926 A 12.548 12.548 0 0 0 2.791 1.342 A 8.749 8.749 0 0 0 2.437 1.463 A 0.2 0.2 0 0 1 2.436 1.463 A 0.2 0.2 0 0 1 2.434 1.464 A 6.273 6.273 0 0 0 1.65 1.809 A 0.2 0.2 0 0 1 1.649 1.81 A 0.2 0.2 0 0 1 1.648 1.81 A 3.382 3.382 0 0 0 1.091 2.174 A 0.2 0.2 0 0 1 1.088 2.176 A 0.2 0.2 0 0 1 1.086 2.178 A 1.409 1.409 0 0 0 0.769 2.529 A 0.443 0.443 0 0 0 0.7 2.757 L 0.7 10.3 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "my-custom-id",
	})
}

func drawSmallSora(
	builder *svgutils.SVGBuilder,
	foregroundColor string,
	backgroundColor string,
) {
	// Add the outer edge
	outerEdge, err := canvas.ParseSVGPath("M 0 7.25 A 0.25 0.25 0 0 0 0.25 7.5 L 14.75 7.5 A 0.25 0.25 0 0 0 15 7.25 L 15 1.875 A 0.628 0.628 0 0 0 14.806 1.451 C 14.492 1.112 13.803 0.799 12.803 0.549 A 15.865 15.865 0 0 0 11.192 0.243 C 10.076 0.085 8.805 0 7.5 0 C 5.763 0 4.089 0.151 2.756 0.423 A 13.185 13.185 0 0 0 2.197 0.549 A 9.22 9.22 0 0 0 1.417 0.778 C 0.502 1.095 0 1.479 0 1.875 L 0 7.25 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.StartGroup("Outer Edge", map[string]string{})
	builder.AddPath(outerEdge.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "outer-edge",
	})
	builder.EndGroup()

	// Add the rounded edge
	roundedEdge, err := canvas.ParseSVGPath("M 0.2 7.25 A 0.05 0.05 0 0 0 0.25 7.3 L 14.75 7.3 A 0.05 0.05 0 0 0 14.8 7.25 L 14.8 1.88 A 0.428 0.428 0 0 0 14.668 1.596 A 0.2 0.2 0 0 1 14.663 1.591 A 1.538 1.538 0 0 0 14.329 1.331 A 0.2 0.2 0 0 1 14.327 1.33 A 3.681 3.681 0 0 0 13.814 1.079 A 6.884 6.884 0 0 0 13.193 0.863 A 10.085 10.085 0 0 0 12.756 0.743 A 15.665 15.665 0 0 0 11.165 0.441 A 24.702 24.702 0 0 0 8.938 0.235 A 29.331 29.331 0 0 0 7.5 0.2 A 27.934 27.934 0 0 0 4.65 0.342 A 19.288 19.288 0 0 0 2.796 0.619 A 12.985 12.985 0 0 0 2.245 0.743 A 9.02 9.02 0 0 0 1.482 0.967 A 5.111 5.111 0 0 0 0.897 1.208 A 0.2 0.2 0 0 1 0.896 1.208 A 2.419 2.419 0 0 0 0.475 1.465 A 0.2 0.2 0 0 1 0.473 1.467 A 0.853 0.853 0 0 0 0.239 1.728 A 0.321 0.321 0 0 0 0.2 1.876 L 0.2 7.25 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.StartGroup("Rounded Edge", map[string]string{})
	builder.AddPath(roundedEdge.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "rounded-edge",
	})
	builder.EndGroup()

	// Add the border
	borderOuter, err := canvas.ParseSVGPath("M 0.5 7 L 14.5 7 L 14.5 1.892 A 0.128 0.128 0 0 0 14.461 1.813 A 0.5 0.5 0 0 1 14.447 1.8 A 1.238 1.238 0 0 0 14.179 1.59 A 0.3 0.3 0 0 1 14.177 1.589 A 0.3 0.3 0 0 1 14.175 1.588 A 3.381 3.381 0 0 0 13.704 1.357 A 0.3 0.3 0 0 1 13.702 1.357 A 6.584 6.584 0 0 0 13.109 1.151 A 9.785 9.785 0 0 0 12.684 1.035 A 15.365 15.365 0 0 0 11.125 0.738 A 24.402 24.402 0 0 0 8.924 0.535 A 29.031 29.031 0 0 0 7.5 0.5 A 27.634 27.634 0 0 0 4.68 0.64 A 18.988 18.988 0 0 0 2.855 0.913 A 12.685 12.685 0 0 0 2.318 1.034 A 8.72 8.72 0 0 0 1.579 1.251 A 4.811 4.811 0 0 0 1.028 1.478 A 0.3 0.3 0 0 1 1.026 1.478 A 0.3 0.3 0 0 1 1.025 1.479 A 2.119 2.119 0 0 0 0.657 1.704 A 0.3 0.3 0 0 1 0.654 1.706 A 0.3 0.3 0 0 1 0.65 1.709 A 0.553 0.553 0 0 0 0.502 1.873 A 0.021 0.021 0 0 0 0.5 1.88 L 0.5 7 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.StartGroup("Vcarve", map[string]string{})
	builder.AddPath(borderOuter.ToSVG(), map[string]string{
		"fill": backgroundColor,
		"id":   "my-custom-id",
	})
	borderInner, err := canvas.ParseSVGPath("M 0.7 6.8 L 14.3 6.8 L 14.3 1.935 A 1.038 1.038 0 0 0 14.079 1.763 A 0.2 0.2 0 0 1 14.076 1.762 A 0.2 0.2 0 0 1 14.073 1.76 A 3.181 3.181 0 0 0 13.63 1.543 A 0.2 0.2 0 0 1 13.629 1.543 A 0.2 0.2 0 0 1 13.627 1.542 A 6.384 6.384 0 0 0 13.052 1.343 A 0.2 0.2 0 0 1 13.051 1.342 A 0.2 0.2 0 0 1 13.05 1.342 A 9.585 9.585 0 0 0 12.636 1.229 A 15.165 15.165 0 0 0 11.097 0.937 A 24.202 24.202 0 0 0 8.915 0.734 A 28.831 28.831 0 0 0 7.5 0.7 A 27.434 27.434 0 0 0 4.7 0.839 A 18.788 18.788 0 0 0 2.894 1.109 A 12.485 12.485 0 0 0 2.366 1.228 A 8.52 8.52 0 0 0 1.644 1.44 A 4.611 4.611 0 0 0 1.115 1.657 A 0.2 0.2 0 0 1 1.113 1.658 A 0.2 0.2 0 0 1 1.111 1.659 A 1.919 1.919 0 0 0 0.778 1.863 A 0.2 0.2 0 0 1 0.773 1.866 A 0.2 0.2 0 0 1 0.768 1.87 A 0.353 0.353 0 0 0 0.7 1.936 L 0.7 6.8 Z")
	if err != nil {
		log.Fatalf("Failed to parse SVG path: %v", err)
	}
	builder.AddPath(borderInner.ToSVG(), map[string]string{
		"fill": foregroundColor,
		"id":   "my-custom-id",
	})
}

func drawSora(
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

	if size == "medium" {
		drawMediumSora(builder, foregroundColor, backgroundColor)
	} else if size == "large" {
		drawLargeSora(builder, foregroundColor, backgroundColor)
	} else if size == "small" {
		drawSmallSora(builder, foregroundColor, backgroundColor)
	} else {
		log.Fatalf("Invalid size: %s", size)
	}

	// Add the text
	numLines := len(lines)

	if numLines > 0 {
		// Calculate the total height for the text containers
		const topPadding = 2.0
		const bottomPadding = 1.0
		lineSpacing := 0.5
		availableHeight := height - (topPadding + bottomPadding) - float64(numLines-1)*lineSpacing

		// Determine the heights for each line
		var containerHeights []float64
		switch numLines {
		case 3:
			containerHeights = []float64{
				availableHeight * 0.2, // First line
				availableHeight * 0.5, // Middle line (taller)
				availableHeight * 0.3, // Last line
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
		currentY := height - topPadding

		for i, line := range lines {
			containerHeight := containerHeights[i]
			container := canvas.Rectangle(width-3.0, containerHeight)
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
