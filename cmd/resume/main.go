package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/LLIEPJIOK/resume/internal/config"
	"github.com/LLIEPJIOK/resume/internal/domain/docs"
	docsService "github.com/LLIEPJIOK/resume/internal/infra/service/docs"
)

const (
	CodeOk               = 0
	CodeErrorConfig      = 1
	CodeErrorDocsService = 2
	CodeErrorGetDocument = 3
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(CodeErrorConfig)
	}

	svc, err := docsService.New(ctx, &cfg.Credentials)
	if err != nil {
		slog.Error("failed to create docs service", slog.Any("error", err))
		os.Exit(CodeErrorDocsService)
	}

	doc, err := docs.New(svc, "1lvYb-rtVvKA__maOFTRC_ogIFLB80Jx1KTWDPauhpFE")
	if err != nil {
		slog.Error("failed to get document", slog.Any("error", err))
		os.Exit(CodeErrorGetDocument)
	}

	fmt.Println(doc.Title())

	r, err := doc.ToResume()
	if err != nil {
		slog.Error("failed to convert document to resume", slog.Any("error", err))
		os.Exit(1)
	}

	data, err := json.Marshal(r)
	if err != nil {
		slog.Error("failed to marshal resume to json", slog.Any("error", err))
		os.Exit(1)
	}

	if err := os.WriteFile("resume.json", data, 0600); err != nil {
		slog.Error("failed to write resume to file", slog.Any("error", err))
		os.Exit(1)
	}
}
