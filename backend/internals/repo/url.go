package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shashank601/url-shortner/backend/internals/domain"
)

type UrlRepo struct {
	DB *pgxpool.Pool
}

func NewUrlRepo(db *pgxpool.Pool) *UrlRepo {
	return &UrlRepo{DB: db}
}



func (r *UrlRepo) InsertUrlKey(ctx context.Context, url *domain.Url) (string, error) {
    query := `
        INSERT INTO urls (customer_id, original_url, short_code)
        VALUES ($1, $2, $3) 
        RETURNING short_code
    `

    var shortCode string
    err := r.DB.QueryRow(ctx, query, url.CustomerID, url.OriginalUrl, url.ShortCode).Scan(&shortCode)
    
    if err != nil {
        return "", err
    }

    return shortCode, nil
}

func (r *UrlRepo) GetUrlByKey(ctx context.Context, shortCode string, customerID int) (*domain.Url, error) {
    query := `
        SELECT id, customer_id, original_url, short_code, is_active
        FROM urls
        WHERE short_code = $1 AND customer_id = $2
    `

    var url domain.Url
    err := r.DB.QueryRow(ctx, query, shortCode, customerID).Scan(
        &url.ID,
        &url.CustomerID,
        &url.OriginalUrl,
        &url.ShortCode,
        &url.IsActive,
    )
    
    if err != nil {
        return nil, err
    }

    return &url, nil
}


