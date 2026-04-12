package markdown

import (
	"strings"
	"testing"
)

func TestPreprocessCodeBlocks_PlantUMLWithServer(t *testing.T) {
	input := []byte("```plantuml\n@startuml\nAlice -> Bob\n@enduml\n```\n")
	result := PreprocessCodeBlocks(input, "https://www.plantuml.com/plantuml")
	s := string(result)

	if !strings.Contains(s, "plantuml-container") {
		t.Error("should render PlantUML image when server is configured")
	}
	if !strings.Contains(s, "<img src=") {
		t.Error("should contain img tag")
	}
	if strings.Contains(s, "plantuml-notice") {
		t.Error("should not show notice when server is configured")
	}
}

func TestPreprocessCodeBlocks_PlantUMLWithoutServer(t *testing.T) {
	input := []byte("```plantuml\n@startuml\nAlice -> Bob\n@enduml\n```\n")
	result := PreprocessCodeBlocks(input, "")
	s := string(result)

	if !strings.Contains(s, "plantuml-notice") {
		t.Error("should show notice when server is not configured")
	}
	if !strings.Contains(s, "PlantUML rendering is disabled") {
		t.Error("notice should contain disabled message")
	}
	if !strings.Contains(s, "--plantuml-server") {
		t.Error("notice should mention --plantuml-server flag")
	}
	if !strings.Contains(s, "--configure") {
		t.Error("notice should mention --configure option")
	}
	if strings.Contains(s, "plantuml-container") {
		t.Error("should not render PlantUML image when server is not configured")
	}
}

func TestPreprocessCodeBlocks_SVG(t *testing.T) {
	input := []byte("```svg\n<svg><circle/></svg>\n```\n")
	result := PreprocessCodeBlocks(input, "")
	s := string(result)

	if !strings.Contains(s, "svg-container") {
		t.Error("should render SVG container")
	}
}

func TestPreprocessCodeBlocks_SVGWithBlankLines(t *testing.T) {
	input := []byte("```svg\n<svg xmlns=\"http://www.w3.org/2000/svg\">\n\n  <!-- comment -->\n  <rect width=\"100\" height=\"100\"/>\n\n  <circle cx=\"50\" cy=\"50\" r=\"40\"/>\n\n</svg>\n```\n")
	result := PreprocessCodeBlocks(input, "")
	s := string(result)

	if !strings.Contains(s, "svg-container") {
		t.Error("should render SVG container")
	}
	if strings.Contains(s, "\n\n") {
		t.Error("should not contain blank lines in SVG output (would break goldmark HTML block parsing)")
	}
	if !strings.Contains(s, "<rect") {
		t.Error("should preserve SVG content")
	}
	if !strings.Contains(s, "<circle") {
		t.Error("should preserve SVG content")
	}
}

func TestPreprocessCodeBlocks_Mermaid(t *testing.T) {
	input := []byte("```mermaid\ngraph TD\n  A-->B\n```\n")
	result := PreprocessCodeBlocks(input, "")
	s := string(result)

	if !strings.Contains(s, `class="mermaid"`) {
		t.Error("should render mermaid pre tag")
	}
}

func TestConvert_SVGWithBlankLines(t *testing.T) {
	input := []byte("# Title\n\n```svg\n<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 100 100\">\n\n  <rect width=\"100\" height=\"100\" fill=\"#fff\"/>\n\n  <text x=\"50\" y=\"50\">Hello</text>\n\n</svg>\n```\n\nParagraph after SVG.\n")
	result, err := Convert(input, "")
	if err != nil {
		t.Fatal(err)
	}
	s := string(result)

	if !strings.Contains(s, "svg-container") {
		t.Error("should contain svg-container div")
	}
	if !strings.Contains(s, "</svg>") {
		t.Error("should contain closing svg tag")
	}
	if !strings.Contains(s, "Paragraph after SVG.") {
		t.Error("should contain paragraph after SVG")
	}
	// The SVG should not be broken apart by goldmark
	if strings.Contains(s, "&lt;rect") || strings.Contains(s, "&lt;text") {
		t.Error("SVG elements should not be HTML-escaped (goldmark should treat as HTML block)")
	}
}

func TestPreprocessCodeBlocks_NoSpecialBlocks(t *testing.T) {
	input := []byte("```go\nfmt.Println(\"hello\")\n```\n")
	result := PreprocessCodeBlocks(input, "")

	if string(result) != string(input) {
		t.Error("should not modify non-special code blocks")
	}
}

func TestPreprocessCodeBlocks_MultiplePlantUMLBlocks(t *testing.T) {
	input := []byte("```plantuml\n@startuml\nA->B\n@enduml\n```\n\ntext\n\n```plantuml\n@startuml\nC->D\n@enduml\n```\n")
	result := PreprocessCodeBlocks(input, "")
	s := string(result)

	count := strings.Count(s, "plantuml-notice")
	if count != 2 {
		t.Errorf("should show 2 notices, got %d", count)
	}
}
