package svg

import (
	"bytes"
	"encoding/xml"
)

// Define an XML structure for the SVG elements
type SVG struct {
	XMLName xml.Name   `xml:"svg"`
	Attrs   []xml.Attr `xml:",any,attr"` // Capture all attributes
	Content []byte     `xml:",innerxml"` // Preserve inner content
}

func modifyOrAddSVGAttributes(svgString string, attributes map[string]string) (string, error) {

	// Parse the SVG XML
	var svgData SVG
	err := xml.Unmarshal([]byte(svgString), &svgData)
	if err != nil {
		return "", err
	}

	// Update or add attributes
	for key, value := range attributes {
		found := false
		for i, attr := range svgData.Attrs {
			if attr.Name.Local == key {
				// Replace the value if the attribute exists
				svgData.Attrs[i].Value = value
				found = true
				break
			}
		}
		// Add the attribute if it doesn't exist
		if !found {
			svgData.Attrs = append(svgData.Attrs, xml.Attr{Name: xml.Name{Local: key}, Value: value})
		}
	}

	// Serialize the modified SVG back to a string
	var buf bytes.Buffer
	encoder := xml.NewEncoder(&buf)
	encoder.Indent("", "  ")
	err = encoder.Encode(svgData)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
