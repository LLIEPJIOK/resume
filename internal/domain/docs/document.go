package docs

import (
	"errors"
	"fmt"
	"strings"

	docsService "github.com/LLIEPJIOK/resume/internal/infra/service/docs"
	"google.golang.org/api/docs/v1"
)

var ErrNoRows = errors.New("no rows in table data")

type Document struct {
	svc *docsService.Service
	doc *docs.Document
}

func New(svc *docsService.Service, documentID string) (*Document, error) {
	doc, err := svc.Document(documentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get document: %v", err)
	}

	return &Document{
		svc: svc,
		doc: doc,
	}, nil
}

func (d *Document) Title() string {
	return d.doc.Title
}

func (d *Document) Parse() (*ParsedDocument, error) {
	if d == nil || d.doc == nil {
		return nil, fmt.Errorf("invalid document")
	}

	parsed := &ParsedDocument{
		Title: d.doc.Title,
	}

	for _, elem := range d.doc.Body.Content {
		switch {
		case elem.Paragraph != nil:
			text := extractText(elem.Paragraph.Elements)
			if strings.TrimSpace(text) != "" {
				parsed.Content = append(parsed.Content, ContentItem{
					Type: TypeParagraph,
					Text: strings.TrimSpace(text),
				})
			}

		case elem.Table != nil:
			table := extractTable(elem.Table)
			if len(table.Rows) > 0 {
				parsed.Content = append(parsed.Content, ContentItem{
					Type:  TypeTable,
					Table: table,
				})
			}
		}
	}

	return parsed, nil
}

func (d *Document) ToResume() (*Resume, error) {
	pr, err := d.Parse()
	if err != nil {
		return nil, fmt.Errorf("failed to parse document: %w", err)
	}

	return pr.ToResume()
}

func extractText(elements []*docs.ParagraphElement) string {
	var text string

	for _, el := range elements {
		if el.TextRun != nil && el.TextRun.Content != "" {
			text += el.TextRun.Content
		}
	}

	return text
}

func extractTable(table *docs.Table) *ParsedTable {
	result := &ParsedTable{}

	for _, row := range table.TableRows {
		var cells []string

		for _, cell := range row.TableCells {
			var cellText string

			for _, content := range cell.Content {
				if content.Paragraph != nil {
					cellText += extractText(content.Paragraph.Elements)
				}
			}

			cellText = strings.TrimSpace(cellText)
			if cellText != "" {
				cells = append(cells, cellText)
			}
		}

		result.Rows = append(result.Rows, cells)
	}

	return result
}
