package convert

import (
	"github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/html"
    "github.com/gomarkdown/markdown/parser"
)

// ConvertToMarkdown takes a HTML file and converts it into a Markdown sheet
func ConvertToMarkdown(htmlContent string) string {
	// Parse HTML to Markdown
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	tdoc := p.Parse([]byte(htmlContent))

	// Render Markdown
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
}