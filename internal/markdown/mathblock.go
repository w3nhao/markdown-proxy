package markdown

import (
	"bytes"
	"regexp"
)

// singleLineMathRe matches a line containing only $$<content>$$ (with optional surrounding whitespace).
// This does NOT match inline usage like "text $$a=b$$ text" because the line must start with $$.
var singleLineMathRe = regexp.MustCompile(`^(\s*)\$\$(.+)\$\$\s*$`)

// fenceStartRe matches the start of a fenced code block (``` or ~~~).
var fenceStartRe = regexp.MustCompile(`^(\x60{3,}|~{3,})`)

// PreprocessMathBlocks ensures that $$ delimiters are on their own lines
// so that goldmark-mathjax can parse them correctly.
// It handles three cases:
//   - Single-line: $$content$$ → $$\ncontent\n$$
//   - Opening with content: $$content...\n → $$\ncontent...\n
//   - Closing with content: ...content$$ → ...content\n$$
//
// Blockquote prefixes ("> ") are stripped before detection and re-added
// to each output line, so math blocks inside blockquotes are handled.
// Fenced code blocks are skipped to avoid modifying code content.
// Only exactly two consecutive $ characters are treated as delimiters
// ($$$ or more are left for goldmark-mathjax to handle directly).
func PreprocessMathBlocks(source []byte) []byte {
	lines := bytes.Split(source, []byte("\n"))
	var result [][]byte
	var inFence bool
	var fenceMarker []byte
	var inMathBlock bool
	var mathBlockPrefix []byte // blockquote prefix captured at math block open

	for _, line := range lines {
		// Extract blockquote prefix (e.g. "> ", ">> ", "> > ")
		bqPrefix, body := splitBlockquotePrefix(line)

		if inFence {
			// Check if this line closes the fence
			if m := fenceStartRe.FindSubmatch(body); m != nil {
				marker := m[1]
				if marker[0] == fenceMarker[0] && len(marker) >= len(fenceMarker) {
					inFence = false
					fenceMarker = nil
				}
			}
			result = append(result, line)
			continue
		}

		// Check if this line opens a fenced code block
		if m := fenceStartRe.FindSubmatch(body); m != nil {
			inFence = true
			fenceMarker = m[1]
			result = append(result, line)
			continue
		}

		if inMathBlock {
			trimmed := bytes.TrimRight(body, " \t")
			if countConsecutiveDollarsAt(trimmed, len(trimmed)-2) == 2 {
				before := trimmed[:len(trimmed)-2]
				if len(bytes.TrimSpace(before)) > 0 {
					// Content before closing $$ → split
					result = append(result, prefixed(mathBlockPrefix, before))
					result = append(result, prefixed(mathBlockPrefix, []byte("$$")))
				} else {
					// Just $$ → proper closing delimiter
					result = append(result, line)
				}
				inMathBlock = false
				mathBlockPrefix = nil
			} else {
				result = append(result, line)
			}
			continue
		}

		// Check for single-line $$...$$ pattern (applied to body after stripping bqPrefix)
		if sm := singleLineMathRe.FindSubmatch(body); sm != nil {
			indent := sm[1]
			content := sm[2]
			result = append(result, prefixed(bqPrefix, append(append([]byte{}, indent...), []byte("$$")...)))
			result = append(result, prefixed(bqPrefix, append(append([]byte{}, indent...), content...)))
			result = append(result, prefixed(bqPrefix, append(append([]byte{}, indent...), []byte("$$")...)))
			continue
		}

		// Check for opening $$ (exactly 2 consecutive $)
		trimmed := bytes.TrimLeft(body, " \t")
		if countConsecutiveDollarsAt(trimmed, 0) == 2 {
			indent := body[:len(body)-len(trimmed)]
			content := trimmed[2:]
			if len(bytes.TrimSpace(content)) > 0 {
				// $$ followed by content → split
				result = append(result, prefixed(bqPrefix, append(append([]byte{}, indent...), []byte("$$")...)))
				result = append(result, prefixed(bqPrefix, append(append([]byte{}, indent...), content...)))
			} else {
				// Just $$ → opening delimiter
				result = append(result, line)
			}
			inMathBlock = true
			mathBlockPrefix = bqPrefix
			continue
		}

		result = append(result, line)
	}

	return bytes.Join(result, []byte("\n"))
}

// splitBlockquotePrefix extracts the leading blockquote markers from a line.
// For example, "> > text" returns ([]byte("> > "), []byte("text")).
// If there is no blockquote prefix, it returns (nil, line).
func splitBlockquotePrefix(line []byte) (prefix, body []byte) {
	i := 0
	for i < len(line) {
		// Skip optional leading spaces (up to 3, per CommonMark)
		spaces := 0
		for spaces < 3 && i+spaces < len(line) && line[i+spaces] == ' ' {
			spaces++
		}
		if i+spaces < len(line) && line[i+spaces] == '>' {
			i += spaces + 1
			// Skip one optional space after '>'
			if i < len(line) && line[i] == ' ' {
				i++
			}
		} else {
			break
		}
	}
	if i == 0 {
		return nil, line
	}
	return line[:i], line[i:]
}

// prefixed prepends a blockquote prefix to content, returning a new slice.
func prefixed(prefix, content []byte) []byte {
	if len(prefix) == 0 {
		return content
	}
	out := make([]byte, len(prefix)+len(content))
	copy(out, prefix)
	copy(out[len(prefix):], content)
	return out
}

// countConsecutiveDollarsAt counts how many consecutive '$' characters
// surround position pos. It returns the total length of the '$' run
// that includes the character at pos.
// If pos is out of range or the character at pos is not '$', it returns 0.
func countConsecutiveDollarsAt(b []byte, pos int) int {
	if pos < 0 || pos >= len(b) || b[pos] != '$' {
		return 0
	}
	// Find the start of the $ run
	start := pos
	for start > 0 && b[start-1] == '$' {
		start--
	}
	// Find the end of the $ run
	end := pos + 1
	for end < len(b) && b[end] == '$' {
		end++
	}
	return end - start
}
