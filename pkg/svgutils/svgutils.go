package svgutils

import (
	"bytes"
	"fmt"
	"log"
)

type SVGBuilder struct {
	buffer bytes.Buffer
	width  float64
	height float64
}

// NewSVGBuilder initializes a new SVGBuilder with specified dimensions.
func NewSVGBuilder(width, height float64) *SVGBuilder {
	log.Printf("Creating SVG with width: %f, height: %f", width, height)
	builder := &SVGBuilder{
		width:  width,
		height: height,
	}
	builder.buffer.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%fmm" height="%fmm">`, width, height))
	return builder
}

// StartGroup starts a new SVG group with optional attributes.
func (s *SVGBuilder) StartGroup(name string, attributes map[string]string) {
	log.Printf("Starting group: %s", name)
	s.buffer.WriteString(fmt.Sprintf(`<g inkscape:groupmode="layer" inkscape:label="%s"`, name))
	for key, value := range attributes {
		s.buffer.WriteString(fmt.Sprintf(` %s="%s"`, key, value))
	}
	s.buffer.WriteString(">")
}

// EndGroup ends the current SVG group.
func (s *SVGBuilder) EndGroup() {
	log.Printf("Ending group")
	s.buffer.WriteString("</g>")
}

// AddPath adds a path to the SVG with specified attributes.
func (s *SVGBuilder) AddPath(d string, attributes map[string]string) {
	log.Printf("Adding path: %s", d)
	s.buffer.WriteString(`<path d="` + d + `"`)
	for key, value := range attributes {
		s.buffer.WriteString(fmt.Sprintf(` %s="%s"`, key, value))
	}
	s.buffer.WriteString(`/>`)
}

// Close finalizes the SVG and returns it as a string.
func (s *SVGBuilder) Close() string {
	log.Printf("Closing SVG")
	s.buffer.WriteString("</svg>")
	return s.buffer.String()
}
