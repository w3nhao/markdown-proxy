package markdown

import (
	"fmt"

	"github.com/yuin/goldmark/ast"
	east "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"
)

// insertLineAnchors walks the AST and inserts <a id="Ln"></a> anchors
// for each source line that has corresponding content.
// It appends anchor HTML to the source slice and returns the modified source.
func insertLineAnchors(doc ast.Node, source []byte) []byte {
	lineStarts := buildLineStarts(source)
	seen := make(map[int]bool)

	type insertion struct {
		parent  ast.Node
		before  ast.Node // insert before this child; nil means append
		lineNum int
		inline  bool // true: RawHTML (inline), false: HTMLBlock (block)
	}
	var insertions []insertion

	ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		kind := node.Kind()

		// Paragraph and Heading: insert inline anchors for each text line
		if kind == ast.KindParagraph || kind == ast.KindHeading {
			for c := node.FirstChild(); c != nil; c = c.NextSibling() {
				if c.Kind() == ast.KindText {
					t := c.(*ast.Text)
					line := byteOffsetToLine(t.Segment.Start, lineStarts)
					if line > 0 && !seen[line] {
						seen[line] = true
						insertions = append(insertions, insertion{
							parent: node, before: c,
							lineNum: line, inline: true,
						})
					}
				}
			}
			return ast.WalkSkipChildren, nil
		}

		// TableCell: insert inline anchor for the cell's first text
		if kind == east.KindTableCell {
			line := firstChildTextLine(node, lineStarts)
			if line > 0 && !seen[line] {
				seen[line] = true
				insertions = append(insertions, insertion{
					parent: node, before: node.FirstChild(),
					lineNum: line, inline: true,
				})
			}
			return ast.WalkSkipChildren, nil
		}

		// FencedCodeBlock: insert block anchor before it (for the fence line)
		if kind == ast.KindFencedCodeBlock {
			lines := node.Lines()
			if lines.Len() > 0 {
				// First content line; the fence is on the line before
				contentLine := byteOffsetToLine(lines.At(0).Start, lineStarts)
				fenceLine := contentLine - 1
				if fenceLine < 1 {
					fenceLine = 1
				}
				if !seen[fenceLine] {
					seen[fenceLine] = true
					insertions = append(insertions, insertion{
						parent: node.Parent(), before: node,
						lineNum: fenceLine, inline: false,
					})
				}
			}
			return ast.WalkSkipChildren, nil
		}

		// CodeBlock (indented): insert block anchor before it
		if kind == ast.KindCodeBlock {
			lines := node.Lines()
			if lines.Len() > 0 {
				line := byteOffsetToLine(lines.At(0).Start, lineStarts)
				if line > 0 && !seen[line] {
					seen[line] = true
					insertions = append(insertions, insertion{
						parent: node.Parent(), before: node,
						lineNum: line, inline: false,
					})
				}
			}
			return ast.WalkSkipChildren, nil
		}

		// ThematicBreak (---): insert block anchor before it
		if kind == ast.KindThematicBreak {
			// ThematicBreak doesn't have Lines; estimate from surrounding nodes
			prev := node.PreviousSibling()
			if prev != nil {
				prevLine := lastNodeLine(prev, lineStarts)
				if prevLine > 0 {
					line := prevLine + 1
					// Skip blank lines
					for line <= len(lineStarts) && isBlankLine(line, lineStarts, source) {
						line++
					}
					if !seen[line] {
						seen[line] = true
						insertions = append(insertions, insertion{
							parent: node.Parent(), before: node,
							lineNum: line, inline: false,
						})
					}
				}
			}
			return ast.WalkSkipChildren, nil
		}

		// List, ListItem, Blockquote, Table, TableHeader, TableBody, TableRow:
		// These are container nodes; continue walking to process their children.
		return ast.WalkContinue, nil
	})

	// Apply insertions
	for _, ins := range insertions {
		anchorHTML := fmt.Sprintf(`<a id="L%d"></a>`, ins.lineNum)
		start := len(source)
		source = append(source, []byte(anchorHTML)...)
		end := len(source)

		if ins.inline {
			raw := ast.NewRawHTML()
			raw.Segments.Append(text.NewSegment(start, end))
			if ins.before != nil {
				ins.parent.InsertBefore(ins.parent, ins.before, raw)
			} else {
				ins.parent.AppendChild(ins.parent, raw)
			}
		} else {
			block := ast.NewHTMLBlock(ast.HTMLBlockType7)
			block.Lines().Append(text.NewSegment(start, end))
			if ins.before != nil {
				ins.parent.InsertBefore(ins.parent, ins.before, block)
			} else {
				ins.parent.AppendChild(ins.parent, block)
			}
		}
	}

	return source
}

// buildLineStarts returns byte offsets where each line begins.
// lineStarts[0] = offset of line 1, lineStarts[1] = offset of line 2, etc.
func buildLineStarts(source []byte) []int {
	starts := []int{0} // line 1 starts at offset 0
	for i, b := range source {
		if b == '\n' && i+1 <= len(source) {
			starts = append(starts, i+1)
		}
	}
	return starts
}

// byteOffsetToLine converts a byte offset to a 1-based line number.
func byteOffsetToLine(offset int, lineStarts []int) int {
	// Binary search for the line containing the offset
	lo, hi := 0, len(lineStarts)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if lineStarts[mid] <= offset {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo // 1-based line number (since lineStarts[0] = line 1)
}

// firstChildTextLine returns the line number of the first Text child node.
func firstChildTextLine(node ast.Node, lineStarts []int) int {
	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		if c.Kind() == ast.KindText {
			t := c.(*ast.Text)
			return byteOffsetToLine(t.Segment.Start, lineStarts)
		}
	}
	return 0
}

// lastNodeLine returns the last source line number of a node's content.
func lastNodeLine(node ast.Node, lineStarts []int) int {
	lines := node.Lines()
	if lines != nil && lines.Len() > 0 {
		lastSeg := lines.At(lines.Len() - 1)
		return byteOffsetToLine(lastSeg.Stop-1, lineStarts)
	}
	// Check children recursively
	var lastLine int
	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		l := lastNodeLine(c, lineStarts)
		if l > lastLine {
			lastLine = l
		}
	}
	// Check inline Text children
	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		if c.Kind() == ast.KindText {
			t := c.(*ast.Text)
			l := byteOffsetToLine(t.Segment.Stop-1, lineStarts)
			if l > lastLine {
				lastLine = l
			}
		}
	}
	return lastLine
}

// isBlankLine returns true if the given line number is a blank line.
func isBlankLine(line int, lineStarts []int, source []byte) bool {
	if line < 1 || line > len(lineStarts) {
		return false
	}
	start := lineStarts[line-1]
	var end int
	if line < len(lineStarts) {
		end = lineStarts[line]
	} else {
		end = len(source)
	}
	for i := start; i < end; i++ {
		if source[i] != ' ' && source[i] != '\t' && source[i] != '\n' && source[i] != '\r' {
			return false
		}
	}
	return true
}
