---
name: AI Agent Documentation Setup
overview: Create a set of `.cursor/rules/*.mdc` files that give AI agents persistent context about the cnc-svg-generator project -- its architecture, conventions, and domain-specific patterns.
todos:
  - id: project-overview
    content: Create `.cursor/rules/project-overview.mdc` with always-apply project context
    status: completed
  - id: go-conventions
    content: "Create `.cursor/rules/go-conventions.mdc` with Go-specific standards (globs: **/*.go)"
    status: completed
  - id: sign-development
    content: "Create `.cursor/rules/sign-development.mdc` with sign/product development guide (globs: pkg/signs/*.go)"
    status: completed
  - id: svg-generation
    content: "Create `.cursor/rules/svg-generation.mdc` with SVG generation pipeline context (globs: pkg/svg/*.go)"
    status: completed
isProject: false
---

# AI Agent Documentation Setup

## Background

The codebase has **no existing `.cursor/rules/`** directory or `.cursorrules` file. There are also no `AGENTS.md` or `CLAUDE.md` files. The only documentation is a minimal [README.md](README.md). This means AI agents currently have zero persistent project context.

## Approach

Create focused `.cursor/rules/*.mdc` rule files following Cursor's recommended format (YAML frontmatter + concise markdown content, under 50 lines each). Each rule targets a specific concern.

## Rules to Create

### 1. `project-overview.mdc` (always apply)

Core context every agent conversation needs:

- Project purpose: Go HTTP API generating CNC-oriented SVGs for sign products
- Tech stack: Go 1.23, Gin, tdewolff/canvas, godotenv
- Layout: `cmd/server/` entry point, `pkg/app` (HTTP layer), `pkg/svg` (SVG generation), `pkg/signs` (per-product drawing), `pkg/fonts` (embedded fonts), `pkg/svgutils` (SVG string builder)
- Dev workflow: run with `air`, env vars in `.env`
- Deployment: staging auto-deploys on push to `main` (Render), production via GitHub Release + Render deploy hook
- Request flow: `GET /v1/api/svg` -> handler parses query params -> `svg.GenerateSVG` -> `signs.DrawSign` dispatches by product ID -> SVG string returned

### 2. `go-conventions.mdc` (globs: `**/*.go`)

Go-specific standards for this project:

- Use `pkg/` package layout; new features go in existing packages or new `pkg/` subpackages
- Entry point is `[cmd/server/main.go](cmd/server/main.go)` (used by Air); root `[main.go](main.go)` is a duplicate -- keep them in sync or remove root
- Error handling: avoid `panic` in request paths; return proper HTTP errors via Gin's `c.JSON`
- Fonts are embedded via `//go:embed` in `[pkg/fonts/fonts.go](pkg/fonts/fonts.go)`
- Environment variables configure product IDs and server binding (`HOST`, `PORT`)
- No test files exist yet; when adding tests, use standard `*_test.go` convention

### 3. `sign-development.mdc` (globs: `pkg/signs/*.go`)

Guide for adding or modifying sign/product types:

- Each sign is a separate file in `pkg/signs/` (e.g., `ellipse.go`, `rectangle.go`, `sora.go`)
- New signs need: a `draw<Name>(c *canvas.Context, params SignParams)` function, a new `case` in the `switch` in `[pkg/signs/signs.go](pkg/signs/signs.go)` `DrawSign`, and a corresponding `<NAME>_PRODUCT_ID` env var
- `SignParams` carries width, height, text lines, font, colors, strokeOnly
- Use `canvas.Context` methods for drawing (paths, text, fills, strokes)
- Some signs use `pkg/svgutils.SVGBuilder` for raw SVG string assembly instead of the canvas API
- Color handling: `pkg/signs/colors.go` and `pkg/svg/colors.go`

### 4. `svg-generation.mdc` (globs: `pkg/svg/*.go`)

Context for the SVG generation pipeline:

- `[pkg/svg/svg.go](pkg/svg/svg.go)` is the main orchestrator: creates a canvas, delegates to `signs.DrawSign`, then post-processes the SVG string
- Post-processing adds viewBox, xmlns attributes, and an optional drop-shadow CSS filter
- `modifyOrAddSVGAttributes` manipulates the raw SVG string after canvas export
- The response Content-Type is `image/svg+xml`
- `pkg/svg/attributes.go`, `colors.go`, `text.go` contain helpers
- `pkg/svg/product.go` has mostly commented-out legacy code

## File Structure

```
.cursor/
  rules/
    project-overview.mdc
    go-conventions.mdc
    sign-development.mdc
    svg-generation.mdc
```
