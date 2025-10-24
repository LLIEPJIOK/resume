package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/LLIEPJIOK/resume/internal/app/cli"
	"github.com/LLIEPJIOK/resume/internal/config"
)

const (
	CodeOk = iota
	CodeErrorFlag
	CodeErrorConfig
	CodeErrorUseCase
	CodeErrorConvert
	CodeErrorValidate

	defaultDocumentID     = ""
	defaultOutputFilePath = "resume.json"
)

func main() {
	var documentID string

	flag.StringVar(&documentID, "doc", defaultDocumentID, "Google Docs document ID (required)")
	flag.StringVar(&documentID, "d", defaultDocumentID, "Google Docs document ID (required)")
	flag.Parse()

	if documentID == "" {
		slog.Error("document ID is required")
		flag.Usage()
		os.Exit(CodeErrorFlag)
	}

	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(CodeErrorConfig)
	}

	uc, err := cli.New(ctx, &cfg.Credentials)
	if err != nil {
		slog.Error("failed to create use case", slog.Any("error", err))
		os.Exit(CodeErrorUseCase)
	}

	validationErrors, err := uc.ValidateResume(documentID)
	if err != nil {
		slog.Error("failed to validate resume", slog.Any("error", err))
		os.Exit(CodeErrorValidate)
	}

	if len(validationErrors) > 0 {
		slog.Warn("validation errors found")

		for field, errMsg := range validationErrors {
			fmt.Printf("%s: %s\n", field, errMsg)
		}
	} else {
		slog.Info("resume validation passed")
	}
}
