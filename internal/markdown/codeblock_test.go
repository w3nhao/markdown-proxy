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

func TestPreprocessCodeBlocks_Mermaid(t *testing.T) {
	input := []byte("```mermaid\ngraph TD\n  A-->B\n```\n")
	result := PreprocessCodeBlocks(input, "")
	s := string(result)

	if !strings.Contains(s, `class="mermaid"`) {
		t.Error("should render mermaid pre tag")
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
