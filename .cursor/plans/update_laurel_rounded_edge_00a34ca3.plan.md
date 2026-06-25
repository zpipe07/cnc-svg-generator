---
name: Update Laurel Rounded Edge
overview: Update the Laurel sign in `laurel.go` to match the new SVG design, replacing the notched edge with a rounded edge and updating all path coordinates.
todos:
  - id: update-outer-edge
    content: Replace outerEdge path data with the new outermost Round path (path 3 from new SVG)
    status: completed
  - id: remove-thru
    content: Remove the entire Thru group (corner circles code)
    status: completed
  - id: add-round-group
    content: Add Round group with 3 rounded-edge paths following fleur.go pattern
    status: completed
  - id: remove-borders
    content: Remove old borderPaths from Vcarve group
    status: completed
  - id: update-filigree
    content: Update all 10 filigree path constants (top/bottom outer + islands) with new coordinates
    status: completed
  - id: update-simple-paths
    content: Replace filligreeTopSimple and filligreeBottomSimple arrays with new decoration paths
    status: completed
  - id: update-horiz-line
    content: Update laurelHorizontalLinePath constant with new line data
    status: completed
isProject: false
---

# Update Laurel Sign to New Rounded-Edge Design

The new SVG (`NEW_SIGN_4-23-26.svg`) restructures the laurel sign from a notched-corner design to a rounded-edge design. The file contains two groups: `Vcarve` (24 decoration paths) and `Round` (4 rounded-edge paths). The current code has `Outside`, `Thru`, and `Vcarve` groups.

## Key structural changes in [pkg/signs/laurel.go](pkg/signs/laurel.go)

### 1. Replace the `outerEdge` path (Outside group)

The old outer edge path (notched corners) is replaced by the new outermost rounded path from the new SVG's Round group (path 3, coords ~0.0005 to ~14.9995). This path defines the sign boundary and is used for `outerBounds` coordinate translation.

### 2. Remove the `Thru` group entirely

The 4 corner circle paths (`cornerCirclePaths`) are no longer in the new design. Remove the entire Thru group block (lines 81-102).

### 3. Add a `Round` group (following [pkg/signs/fleur.go](pkg/signs/fleur.go) pattern)

Add 3 rounded-edge paths in a new `Round` group between the Outside and Vcarve groups:

- Path 1: Left decorative border (`M 0.471332,0.840719...`)
- Path 2: Full inner border frame (`M 14.798173,0.645544...`)
- Path 4: Right decorative border (`M 14.377594,0.940343...`)

Fill with `foregroundColor`, matching fleur.go's Round group pattern.

### 4. Remove old `borderPaths` from Vcarve group

The 4 border paths (lines 106-111) are replaced by the Round group paths above. Remove them entirely.

### 5. Update all filigree path constants

All 10 filigree path constants need updated coordinates from the new SVG:

- `laurelTopFiligreeOuter` (line 13)
- `laurelTopFiligreeIsland0-3` (lines 14-17)
- `laurelBottomFiligreeOuter` (line 28)
- `laurelBottomFiligreeIsland0-3` (lines 29-32)

### 6. Update simple decoration paths

Replace `filligreeTopSimple` (6 paths) and `filligreeBottomSimple` (6 paths) with the new SVG's decoration paths. The new design has:

- Top: 4 paths (right scroll, center diamond, left scroll, center flower) plus 2 new arc/ornament paths
- Bottom: Corresponding 6 mirrored paths

### 7. Update horizontal line path

Replace the `laurelHorizontalLinePath` constant with the new horizontal line data. The new design uses a slightly different line shape (simpler arrow endpoints, from x~2.4 to x~12.6).

## Color mapping

- Outside group: `foregroundColor` (unchanged)
- Round group: `foregroundColor` (matching fleur.go)
- Vcarve group: `backgroundColor` (unchanged)
