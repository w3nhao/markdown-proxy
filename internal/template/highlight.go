package template

import (
	"bytes"
	"strings"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/styles"
)

// highlightCSS holds the combined syntax highlight CSS for all themes.
var highlightCSS string

func init() {
	var buf strings.Builder

	// GitHub/Simple themes: use "github" chroma style
	githubSyntax := generateSyntaxCSS("github")
	buf.WriteString(scopeCSS(githubSyntax, ".theme-github"))
	buf.WriteString("\n")
	buf.WriteString(scopeCSS(githubSyntax, ".theme-simple"))
	buf.WriteString("\n")

	// Dark theme: use "monokai" chroma style
	monokaiSyntax := generateSyntaxCSS("monokai")
	buf.WriteString(scopeCSS(monokaiSyntax, ".theme-dark"))
	buf.WriteString("\n")

	highlightCSS = buf.String()
}

// generateSyntaxCSS generates CSS from a named chroma style.
func generateSyntaxCSS(styleName string) string {
	style := styles.Get(styleName)
	formatter := chromahtml.New(chromahtml.WithClasses(true))

	var buf bytes.Buffer
	if err := formatter.WriteCSS(&buf, style); err != nil {
		return ""
	}
	return buf.String()
}

// scopeCSS prefixes CSS selectors with a theme class scope.
// Each line containing a CSS selector (starting with ".") gets the scope prepended.
func scopeCSS(css, scope string) string {
	var result strings.Builder
	for _, line := range strings.Split(css, "\n") {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		// Find the first "." which marks the CSS selector
		idx := strings.Index(line, ".")
		if idx >= 0 && strings.Contains(line, "{") {
			result.WriteString(line[:idx])
			result.WriteString(scope)
			result.WriteString(" ")
			result.WriteString(line[idx:])
		} else {
			result.WriteString(line)
		}
		result.WriteString("\n")
	}
	return result.String()
}
