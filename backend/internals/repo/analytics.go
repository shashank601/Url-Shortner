package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shashank601/url-shortner/backend/internals/domain"
)

type AnalyticsRepo struct {
	DB *pgxpool.Pool
}

func NewAnalyticsRepo(db *pgxpool.Pool) *AnalyticsRepo {
	return &AnalyticsRepo{DB: db}
}

func (r *AnalyticsRepo) GetAnalytics(ctx context.Context, customerID int, urlID int) (*domain.Analytics, error) {

	// TODO: Implement a way to get these analytics
	result := &domain.Analytics{
		Browsers:  map[string]int{}, 
		OS:        map[string]int{}, 
		Referrers: map[string]int{}, 
	}

	query := `
		SELECT
			COUNT(*) AS total_clicks,
			COUNT(DISTINCT ce.ip_address) AS unique_clicks
		FROM click_events ce
		JOIN urls u ON u.id = ce.url_id
		WHERE ce.url_id = $1 AND u.customer_id = $2
	`
	if err := r.DB.QueryRow(ctx, query, urlID, customerID).Scan(&result.TotalClicks, &result.UniqueClicks); err != nil {
		return nil, err
	}

	return result, nil
}

