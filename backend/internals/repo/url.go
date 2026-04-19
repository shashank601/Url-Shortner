package repo

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/shashank601/url-shortner/backend/internals/domain"
)

type UrlRepo struct {
	DB *pgxpool.Pool
    Redis *redis.Client
}

func NewUrlRepo(db *pgxpool.Pool, redis *redis.Client) *UrlRepo {
	return &UrlRepo{DB: db, Redis: redis}
}



func (r *UrlRepo) InsertUrlKey(ctx context.Context, url *domain.Url) (int, error) {
    query := `
        INSERT INTO urls (customer_id, original_url, short_code)
        VALUES ($1, $2, $3) 
        RETURNING id
    `

    var id int
    err := r.DB.QueryRow(ctx, query, url.CustomerID, url.OriginalUrl, url.ShortCode).Scan(&id)
    
    if err != nil {
        return 0, err
    }

	url.ID = id

	return id, nil
}

func (r *UrlRepo) GetUrlByKey(ctx context.Context, shortCode string) (*domain.Url, error) {
    query := `
        SELECT id, customer_id, original_url, short_code, is_active
        FROM urls
        WHERE short_code = $1
    `

    var url domain.Url
    err := r.DB.QueryRow(ctx, query, shortCode).Scan(
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
func (r *UrlRepo) SetUrlInCache(ctx context.Context, shortCode string, url *domain.Url) error {

    data, err := json.Marshal(url)
    if err != nil {
        return err
    }
    
    return r.Redis.Set(ctx, shortCode, data, 24*time.Hour).Err()
}

func (r *UrlRepo) GetUrlFromCache(ctx context.Context, shortCode string) (*domain.Url, error) {
    data, err := r.Redis.Get(ctx, shortCode).Bytes()
    if err != nil {
        return nil, err
    }

    var url domain.Url
    if err := json.Unmarshal(data, &url); err != nil {
        return nil, err
    }
    return &url, nil
}


func (r *UrlRepo) InsertClickEvent(ctx context.Context, evt *domain.ClickEvent) error {
    query := `
        INSERT INTO click_events (url_id, ip_address, user_agent, referrer)
        VALUES ($1, $2, $3, $4)
    `

    _, err := r.DB.Exec(ctx, query, evt.UrlID, evt.IP, evt.UserAgent, evt.Referrer)
    return err
}
