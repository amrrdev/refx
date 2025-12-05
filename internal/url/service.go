package url

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/amrrdev/refx/db"
	"github.com/amrrdev/refx/internal/redis"
	"github.com/amrrdev/refx/internal/snowflake"
	"github.com/jackc/pgx/v5"
)

var ErrShortNotFound = errors.New("short url not found")
var ErrLongAlreadyExists = errors.New("long url already exists")

type Service struct {
	repository Repository
	cache      *redis.RedisClient
	snowflake  *snowflake.Generator
}

func NewService(repo Repository, client *redis.RedisClient) *Service {
	return &Service{
		repository: repo,
		cache:      client,
		snowflake:  snowflake.New(1),
	}
}
func (s *Service) CreateShortUrl(ctx context.Context, longUrl string) (db.CreateShortCodeRow, error) {
	shortUrl := s.GenerateCode(longUrl)

	result, err := s.repository.CreateShortUrl(ctx, longUrl, shortUrl)
	if err != nil {
		return db.CreateShortCodeRow{}, fmt.Errorf("insert failed: %w", err)
	}

	if s.cache != nil {
		if err := s.cache.SetLongUrl(ctx, shortUrl, longUrl); err != nil {
			log.Println("redis cache set failed:", err)
		}
	}

	return result, nil
}

func (s *Service) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	longUrl, err := s.cache.GetLongUrl(ctx, shortUrl)
	if err == nil {
		return longUrl, nil
	}

	longUrl, err = s.repository.GetLongUrl(ctx, shortUrl)

	if err == pgx.ErrNoRows {
		return "", ErrShortNotFound
	}

	if err != nil {
		return "", fmt.Errorf("get long url failed: %w", err)
	}

	if err := s.cache.SetLongUrl(ctx, shortUrl, longUrl); err != nil {
		log.Println("redis cache failed:", err)
	}

	return longUrl, nil
}

func (s *Service) GenerateCode(longUrl string) string {
	id := s.snowflake.NextID()
	return EncodeBase62(id)
}
