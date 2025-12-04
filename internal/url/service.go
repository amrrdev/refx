package url

import (
	"context"
	"errors"
	"fmt"

	"github.com/amrrdev/refx/db"
	"github.com/jackc/pgx/v5"
)

var ErrShortNotFound = errors.New("short url not found")
var ErrLongAlreadyExists = errors.New("long url already exists")

type Service struct {
	repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{repository: repo}
}

func (s *Service) CreateShortUrl(ctx context.Context, longUrl string) (db.CreateShortCodeRow, error) {
	_, err := s.repository.GetByLongUrl(ctx, longUrl)
	if err == nil {
		return db.CreateShortCodeRow{}, ErrLongAlreadyExists
	}
	if err != pgx.ErrNoRows {
		return db.CreateShortCodeRow{}, fmt.Errorf("lookup failed: %w", err)
	}

	shortUrl := s.GenerateCode(longUrl)

	return s.repository.CreateShortUrl(ctx, longUrl, shortUrl)
}

func (s *Service) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	longUrl, err := s.repository.GetLongUrl(ctx, shortUrl)

	if err == pgx.ErrNoRows {
		return "", ErrShortNotFound
	}

	if err != nil {
		return "", fmt.Errorf("get long url failed: %w", err)
	}

	return longUrl, nil
}

func (s *Service) GenerateCode(longUrl string) string {
	return "123"
}
