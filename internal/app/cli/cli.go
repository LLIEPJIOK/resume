package cli

import (
	"context"
	"fmt"

	"github.com/LLIEPJIOK/resume/internal/config"
	"github.com/LLIEPJIOK/resume/internal/domain/docs"
	docsService "github.com/LLIEPJIOK/resume/internal/infra/service/docs"
)

type UseCase struct {
	svc *docsService.Service
}

func New(ctx context.Context, credentials *config.Credentials) (*UseCase, error) {
	svc, err := docsService.New(ctx, credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to create docs service: %w", err)
	}

	return &UseCase{
		svc: svc,
	}, nil
}

func (uc *UseCase) ValidateResume(documentID string) (map[string]string, error) {
	doc, err := docs.New(uc.svc, documentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get document: %w", err)
	}

	resume, err := doc.ToResume()
	if err != nil {
		return nil, fmt.Errorf("failed to convert document to resume: %w", err)
	}

	return resume.Validate(), nil
}
