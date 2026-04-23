package markdown

import (
	"bytes"
	"html"
	"strings"
)

// PreprocessFrontmatter detects a YAML/TOML frontmatter block at the top of the
// source and replaces it with a rendered HTML info block so that goldmark does
// not misinterpret the delimiters (e.g. the closing "---" being read as a
// setext H2 underline).
//
// Recognized delimiters:
//   - YAML: "---" / "---"
//   - TOML: "+++" / "+++"
//
// If no frontmatter is present, the source is returned unchanged.
func PreprocessFrontmatter(source []byte) []byte {
	openDelim, ok := detectFrontmatterDelim(source)
	if !ok {
		return source
	}
	closeDelim := openDelim

	// Skip past the opening delimiter line.
	rest := source[len(openDelim):]
	if len(rest) > 0 && rest[0] == '\r' {
		rest = rest[1:]
	}
	if len(rest) > 0 && rest[0] == '\n' {
		rest = rest[1:]
	} else {
		return source
	}

	// Find the closing delimiter on its own line.
	closeKey := append([]byte{'\n'}, closeDelim...)
	idx := bytes.Index(rest, closeKey)
	for idx >= 0 {
		after := rest[idx+len(closeKey):]
		if len(after) == 0 || after[0] == '\n' || after[0] == '\r' {
			break
		}
		next := bytes.Index(after, closeKey)
		if next < 0 {
			idx = -1
			break
		}
		idx = idx + len(closeKey) + next
	}
	if idx < 0 {
		return source
	}

	fm := rest[:idx]
	tail := rest[idx+len(closeKey):]
	if len(tail) > 0 && tail[0] == '\r' {
		tail = tail[1:]
	}
	if len(tail) > 0 && tail[0] == '\n' {
		tail = tail[1:]
	}

	var buf bytes.Buffer
	buf.WriteString(`<div class="frontmatter"><pre class="frontmatter-body">`)
	buf.WriteString(html.EscapeString(strings.TrimRight(string(fm), "\r\n")))
	buf.WriteString("</pre></div>\n\n")
	buf.Write(tail)
	return buf.Bytes()
}

func detectFrontmatterDelim(source []byte) ([]byte, bool) {
	switch {
	case bytes.HasPrefix(source, []byte("---\n")), bytes.HasPrefix(source, []byte("---\r\n")):
		return []byte("---"), true
	case bytes.HasPrefix(source, []byte("+++\n")), bytes.HasPrefix(source, []byte("+++\r\n")):
		return []byte("+++"), true
	}
	return nil, false
}
