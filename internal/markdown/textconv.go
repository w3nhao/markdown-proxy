package markdown

import (
	"bytes"
	"fmt"
	"html"
	"strings"
)

// ConvertText converts plain text content to HTML with line anchors.
// Each line is wrapped in a <span id="Ln" class="source-line"> for
// individual line-level highlighting. The output is in a <pre><code> block.
func ConvertText(source []byte) []byte {
	lines := strings.Split(string(source), "\n")

	var buf bytes.Buffer
	buf.WriteString("<pre class=\"text-file\"><code>")
	for i, line := range lines {
		lineNum := i + 1
		fmt.Fprintf(&buf, `<span id="L%d" class="source-line">`, lineNum)
		buf.WriteString(html.EscapeString(line))
		buf.WriteString("</span>")
		if lineNum < len(lines) {
			buf.WriteByte('\n')
		}
	}
	buf.WriteString("</code></pre>\n")

	return buf.Bytes()
}
