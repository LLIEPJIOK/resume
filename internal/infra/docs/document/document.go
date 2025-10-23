package document

import (
	"errors"
	"fmt"

	"github.com/LLIEPJIOK/resume/internal/infra/docs/service"
	"google.golang.org/api/docs/v1"
)

var ErrNoRows = errors.New("no rows in table data")

type Document struct {
	svc *service.Service
	doc *docs.Document
}

func New(svc *service.Service, documentID string) (*Document, error) {
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
