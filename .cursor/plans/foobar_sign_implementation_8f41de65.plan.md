---
name: Foobar Sign Implementation
overview: Implement support for a new "Foobar" sign type using the NewSIGN_3-5-26.svg template, following the patterns established in rectangle.go, fleur.go, and deco.go. The sign will be referenced via FOOBAR_PRODUCT_ID env var (already in .env as 130).
todos: []
isProject: false
---

# Foobar Sign Implementation Plan

## SVG Template Analysis

The [NewSIGN_3-5-26.svg](NewSIGN_3-5-26.svg) template (15" x 11" viewBox) has three Inkscape layers:

| Layer       | Paths                  | Purpose                                                                                |
| ----------- | ---------------------- | -------------------------------------------------------------------------------------- |
| **Vcarve**  | 7 paths (lines 5-11)   | Cursive script text ("foobar") - **replace with dynamic text**                         |
| **Thru**    | 14 paths (lines 14-24) | Outer frame, border, horizontal dividers, corner circles                               |
| **Layer_2** | 35 paths (lines 28-60) | Block "FOOBAR" letters (28-41) - **replace**; decorative flourishes (42-60) - **keep** |

## Path Categorization

**outerEdge** (1 path): Thru layer path 19 - the full outer boundary rectangle with detailed corners (line 20)

**border** (4 paths): Thru layer paths 13-16 - left and right scalloped frame pieces (lines 14-17)

**decorations** (Thru): Paths 17, 18 (horizontal divider lines), 20-23 (four corner circles)

**decorations** (Layer_2): Paths 42-60 - decorative flourishes only (exclude 28-41 which are text)

**text**: Dynamically generated - replaces Vcarve (5-11) and Layer_2 FOOBAR (28-41)

## Implementation Steps

### 1. Create `pkg/signs/foobar.go`

- **drawFoobar()** - main entry matching `drawDeco`/`drawFleur` signature: `(width, height, foregroundColor, backgroundColor, lines []string, strokeOnly bool) string`
- **drawFoobarShapes()** - helper that draws outerEdge, border, decorations (mirrors `drawLargeFleur`/`drawMediumDeco` structure)
- **drawFoobarText()** - helper for dynamic text (mirrors `drawFleurText`/`drawText` in deco.go)
- Use `canvas.ParseSVGPath()` for each path, translate by `-outerBounds.X0, -outerBounds.Y0` to normalize origin (per fleur/deco)
- Groups: "Outside" (outerEdge), "Round"/"Thru" (border), "Vcarve" (decorations + text)
- Follow deco.go for `GetColor()` and strokeOnly handling

### 2. Path Extraction

Copy `d` attribute values from the SVG into string constants/slices:

- 1 outerEdge path from line 20
- 4 border paths from lines 14-17
- 6 Thru decoration paths (17, 18, 20-23)
- 19 Layer_2 decoration paths (42-60) - extract the `d` values

### 3. Text Layout

- Single text area for 1-3 lines (like rectangle, fleur, deco)
- Container: center of sign, approximate region y ~4.5-6.5 (where Vcarve script was)
- Use a font from [pkg/fonts/fonts.go](pkg/fonts/fonts.go) - e.g. Nephilm or Limelight to match deco
- Padding: `topPadding ~2.5`, `bottomPadding ~2.5`, `lineSpacing 0.5`
- Scale text to fit container, center horizontally and vertically

### 4. Register in `pkg/signs/signs.go`

Add case to the switch in `DrawSign()`:

```go
case os.Getenv("FOOBAR_PRODUCT_ID"):
    return drawFoobar(width, height, foregroundColor, backgroundColor, lines, strokeOnly)
```

Note: Foobar does not use `size` or `fontFamily` (similar to deco). If the switch requires it, pass empty string for size and nil for fontFamily where appropriate - check `drawDeco`/`drawRectangle` call sites.

### 5. Verify .env

`FOOBAR_PRODUCT_ID=130` is already present in [.env](.env) - no change needed.

## File Changes Summary

| File                  | Action                                    |
| --------------------- | ----------------------------------------- |
| `pkg/signs/foobar.go` | **Create** - new sign implementation      |
| `pkg/signs/signs.go`  | **Edit** - add FOOBAR case to switch      |
| `.env`                | No change (FOOBAR_PRODUCT_ID already set) |

## Key Code References

- **Path parsing and translation**: [fleur.go:22-33](pkg/signs/fleur.go) (outerEdge), [deco.go:28-48](pkg/signs/deco.go) (border)
- **Decorations loop**: [fleur.go:45-67](pkg/signs/fleur.go)
- **Text generation**: [deco.go:131-213](pkg/signs/deco.go) (drawText), [rectangle.go:146-239](pkg/signs/rectangle.go)
- **SVGBuilder usage**: [deco.go:219-237](pkg/signs/deco.go) (drawDeco main flow)

## Optional: Defs/Filters

Rectangle uses `rectangleDefs()` for groove-shadow and text-engrave filters. Foobar template paths are stroke-based; if depth effects are desired, add similar defs. For initial implementation, follow deco.go (no defs) for simplicity.
