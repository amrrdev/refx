package url

import (
	"context"

	"github.com/amrrdev/refx/db"
	"github.com/amrrdev/refx/internal/database"
)

type Repository interface {
	GetLongUrl(ctx context.Context, shortUrl string) (string, error)
	GetByLongUrl(ctx context.Context, longUrl string) (string, error)
	CreateShortUrl(ctx context.Context, longUrl string, shortUrl string) (db.CreateShortCodeRow, error)
}

type repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) Repository {
	return &repository{db: db}
}

func (r *repository) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	return r.db.Queries.GetLongUrl(ctx, shortUrl)
}

func (r *repository) GetByLongUrl(ctx context.Context, longUrl string) (string, error) {
	return r.db.Queries.GetByLongUrl(ctx, longUrl)
}

func (r *repository) CreateShortUrl(ctx context.Context, longUrl string, shortUrl string) (db.CreateShortCodeRow, error) {
	return r.db.Queries.CreateShortCode(ctx, db.CreateShortCodeParams{
		LongUrl:   longUrl,
		ShortCode: shortUrl,
	})
}
