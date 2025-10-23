package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/LLIEPJIOK/resume/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

const stateToken = "state-token"

type Service struct {
	svc *docs.Service
}

func New(ctx context.Context, cfg *config.Credentials) (*Service, error) {
	oauthConfig, err := google.ConfigFromJSON(cfg.Data, docs.DocumentsScope)
	if err != nil {
		return nil, fmt.Errorf("failed to get oauth config: %v", err)
	}

	client, err := getClient(ctx, oauthConfig, cfg.TokenPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get oauth client: %v", err)
	}

	srv, err := docs.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("failed to create docs service: %v", err)
	}

	return &Service{
		svc: srv,
	}, nil
}

func (s *Service) Documents() *docs.DocumentsService {
	return s.svc.Documents
}

func (s *Service) Document(documentID string) (*docs.Document, error) {
	doc, err := s.svc.Documents.Get(documentID).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get document: %v", err)
	}

	return doc, nil
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := f.Close(); err != nil {
			slog.Error("failed to close file", slog.Any("error", err))
		}
	}()

	token := &oauth2.Token{}

	err = json.NewDecoder(f).Decode(token)
	if err != nil {
		return nil, fmt.Errorf("failed to decode token: %v", err)
	}

	return token, nil
}

func saveToken(path string, token *oauth2.Token) error {
	slog.Info("Saving credentials", slog.String("path", path))

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			slog.Error("failed to close file", slog.Any("error", err))
		}
	}()

	json.NewEncoder(f).Encode(token)

	return nil
}

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL(stateToken)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string

	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("failed to read authorization code: %v", err)
	}

	token, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("failed to retrieve token from web: %v", err)
	}

	return token, nil
}

func getClient(ctx context.Context, config *oauth2.Config, tokenPath string) (*http.Client, error) {
	token, err := tokenFromFile(tokenPath)
	if err != nil {
		token, err = getTokenFromWeb(config)
		if err != nil {
			return nil, fmt.Errorf("failed to get token from web: %v", err)
		}

		if err := saveToken(tokenPath, token); err != nil {
			return nil, fmt.Errorf("failed to save token: %v", err)
		}
	}

	return config.Client(ctx, token), nil
}
