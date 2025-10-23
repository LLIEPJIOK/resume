package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/LLIEPJIOK/resume/internal/config"
	"github.com/LLIEPJIOK/resume/internal/infra/docs/document"
	"github.com/LLIEPJIOK/resume/internal/infra/docs/service"
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

	svc, err := service.New(ctx, &cfg.Credentials)
	if err != nil {
		slog.Error("failed to create docs service", slog.Any("error", err))
		os.Exit(CodeErrorDocsService)
	}

	doc, err := document.New(svc, "1ibwc7w2-Dkt2PVBXXzIAXkM0m3zeBNQAEPBuAOWF1Qc")
	if err != nil {
		slog.Error("failed to get document", slog.Any("error", err))
		os.Exit(CodeErrorGetDocument)
	}

	fmt.Println(doc.Title())
}
